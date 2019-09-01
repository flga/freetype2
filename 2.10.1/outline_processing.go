package freetype2

// #include <stdlib.h>
// #include <ft2build.h>
// #include FT_IMAGE_H
// #include FT_OUTLINE_H
// #include FT_BBOX_H
//
// int OutlineMoveToCallback(const FT_Vector* to, void* user);
// int OutlineLineToCallback(const FT_Vector* to, void* user);
// int OutlineConicToCallback(const FT_Vector* control, const FT_Vector* to, void* user);
// int OutlineCubicToCallback(const FT_Vector* control1, const FT_Vector* control2, const FT_Vector* to, void* user);
import "C"

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/flga/freetype2/fixed"
)

// Outline describes an outline to the scan-line converter.
//
// The B/W rasterizer only checks bit 2 in the tags array for the first point of
// each contour. The drop-out mode as given with OutlineIgnoreDropouts,
// OutlineSmartDropouts, and OutlineIncludeStubs in flags is then overridden.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline
type Outline struct {
	ptr         *C.FT_Outline `deep:"-"`
	l           *Library      `deep:"-"`
	userCreated bool

	// The outline's point coordinates.
	Points []Vector
	// The type of each outline point.
	//
	// If bit 0 is unset, the point is ‘off’ the curve, i.e., a Bezier control
	// point, while it is ‘on’ if set.
	//
	// Bit 1 is meaningful for ‘off’ points only. If set, it indicates a
	// third-order Bezier arc control point; and a second-order control point if
	// unset.
	//
	// If bit 2 is set, bits 5-7 contain the drop-out mode (as defined in the
	// OpenType specification; the value is the same as the argument to the
	// ‘SCANMODE’ instruction).
	//
	// Bits 3 and 4 are reserved for internal purposes.
	Tags []byte
	// The end point of each contour within the outline. For example, the first
	// contour is defined by the points ‘0’ to contours[0], the second one is
	// defined by the points contours[0]+1 to contours[1], etc.
	Contours []int16
	// A set of bit flags used to characterize the outline and give hints to the
	// scan-converter and hinter on how to convert/grid-fit it.
	Flags OutlineFlag
}

// NewOutline creates a new outline of a given size.
//
// Points must be smaller than or equal to 0xFFFF (65535).
// Contours must be in the range 0 to points.
//
// Note that the new outline will not necessarily be freed, when destroying the
// library, by Library.Free.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_new
func (l *Library) NewOutline(points, contours int) (*Outline, error) {
	if l == nil || l.ptr == nil {
		return nil, ErrInvalidLibraryHandle
	}

	if points < 0 || contours < 0 {
		return nil, ErrInvalidArgument
	}

	if points > 0xffff {
		return nil, ErrArrayTooLarge
	}

	if contours > points {
		return nil, ErrInvalidArgument
	}

	var outline C.FT_Outline
	if err := getErr(C.FT_Outline_New(l.ptr, C.uint(points), C.int(contours), &outline)); err != nil {
		return nil, err
	}

	ret := &Outline{ptr: &outline, l: l, userCreated: true}
	ret.reload()

	l.dealloc = append(l.dealloc, func() { ret.Free() })

	return ret, nil
}

// Free destroys an outline created with NewOutline.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_done
func (o *Outline) Free() error {
	if o == nil || o.ptr == nil {
		return nil
	}

	if !o.userCreated {
		return nil
	}

	if o.l == nil || o.l.ptr == nil {
		return ErrInvalidLibraryHandle
	}

	if err := getErr(C.FT_Outline_Done(o.l.ptr, o.ptr)); err != nil {
		return err
	}

	*o = Outline{}
	return nil
}

func (o *Outline) reload() {
	if o == nil {
		return
	}

	if o.ptr == nil {
		*o = Outline{}
		return
	}

	if o.ptr.n_points > 0 {
		if len(o.Points) != int(o.ptr.n_points) {
			o.Points = make([]Vector, o.ptr.n_points)
		}

		ptr := (*[(1<<31 - 1) / C.sizeof_FT_Vector]C.FT_Vector)(unsafe.Pointer(o.ptr.points))[:o.ptr.n_points:o.ptr.n_points]
		for i := range o.Points {
			o.Points[i] = Vector{X: Pos(ptr[i].x), Y: Pos(ptr[i].y)}
		}
	}

	if o.ptr.n_points > 0 {
		if len(o.Tags) != int(o.ptr.n_points) {
			o.Tags = make([]byte, o.ptr.n_points)
		}
		ptr := (*[(1<<31 - 1) / C.sizeof_char]C.char)(unsafe.Pointer(o.ptr.tags))[:o.ptr.n_points:o.ptr.n_points]
		for i := range o.Tags {
			o.Tags[i] = byte(ptr[i])
		}
	}

	if o.ptr.n_contours > 0 {
		if len(o.Contours) != int(o.ptr.n_contours) {
			o.Contours = make([]int16, o.ptr.n_contours)
		}
		ptr := (*[(1<<31 - 1) / C.sizeof_short]C.short)(unsafe.Pointer(o.ptr.contours))[:o.ptr.n_contours:o.ptr.n_contours]
		for i := range o.Contours {
			o.Contours[i] = int16(ptr[i])
		}
	}

	o.Flags = OutlineFlag(o.ptr.flags)
}

// CopyTo copies an outline into another one.
// Both objects must have the same sizes (number of points & number of contours)
// when this function is called.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_copy
func (o *Outline) CopyTo(target *Outline) error {
	if o == nil || o.ptr == nil {
		return ErrInvalidOutline
	}

	if target == nil || target.ptr == nil {
		return ErrInvalidOutline
	}

	if o.ptr.n_points != target.ptr.n_points || o.ptr.n_contours != target.ptr.n_contours {
		return ErrInvalidArgument
	}

	if err := getErr(C.FT_Outline_Copy(o.ptr, target.ptr)); err != nil {
		return err
	}

	target.reload()
	return nil
}

// Translate applies a simple translation to the points of an outline.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_translate
func (o *Outline) Translate(x, y Pos) {
	if o == nil || o.ptr == nil {
		return
	}

	C.FT_Outline_Translate(o.ptr, C.FT_Pos(x), C.FT_Pos(y))
	o.reload()
}

// Transform applies a simple 2x2 matrix to all of an outline's points.
// Useful for applying rotations, slanting, flipping, etc.
//
// You can use Translate if you need to translate the outline's points.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_transform
func (o *Outline) Transform(m Matrix) {
	if o == nil || o.ptr == nil {
		return
	}

	cm := C.FT_Matrix{
		xx: C.FT_Fixed(m.Xx),
		xy: C.FT_Fixed(m.Xy),
		yx: C.FT_Fixed(m.Yx),
		yy: C.FT_Fixed(m.Yy),
	}

	C.FT_Outline_Transform(o.ptr, &cm)
	o.reload()
}

// Embolden emboldens an outline.
//
// The new outline will be at most 4 times strength pixels wider and higher.
// You may think of the left and bottom borders as unchanged.
//
// Negative strength values to reduce the outline thickness are possible also.
//
// The used algorithm to increase or decrease the thickness of the glyph doesn't
// change the number of points; this means that certain situations like acute
// angles or intersections are sometimes handled incorrectly.
//
// If you need ‘better’ metrics values you should call CBox or BBox.
//
// To get meaningful results, font scaling values must be set with functions
// like Face.SetCharSize before calling GlyphSlot.RenderGlyph.
//
//	TODO:
//	face.LoadGlyph(index, LoadDefault)
//	slot := face.GlyphSlot()
//	if slot.Format == GlyphFormatOutline {
//		slot.Glyph.Embolden(strength)
//	}
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_embolden
func (o *Outline) Embolden(strength fixed.Int26_6) error {
	if o == nil || o.ptr == nil {
		return ErrInvalidOutline
	}

	if err := getErr(C.FT_Outline_Embolden(o.ptr, C.FT_Pos(strength))); err != nil {
		return err
	}

	o.reload()
	return nil
}

// EmboldenXY emboldens an outline.
//
//The new outline will be xStrength pixels wider and yStrength pixels higher.
// Otherwise, it is similar to Embolden, which uses the same strength in both
// directions.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_emboldenxy
func (o *Outline) EmboldenXY(xStrength, yStrength Pos) error {
	if o == nil || o.ptr == nil {
		return ErrInvalidOutline
	}

	if err := getErr(C.FT_Outline_EmboldenXY(o.ptr, C.FT_Pos(xStrength), C.FT_Pos(yStrength))); err != nil {
		return err
	}

	o.reload()
	return nil
}

// Reverse reverses the drawing direction of an outline.
// This is used to ensure consistent fill conventions for mirrored glyphs.
//
// This function toggles the bit flag OutlineReverseFill in the outline's flags
// field.
//
// It shouldn't be used by a normal client application, unless it knows what it
// is doing.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_reverse
func (o *Outline) Reverse() {
	if o == nil || o.ptr == nil {
		return
	}

	C.FT_Outline_Reverse(o.ptr)
	o.reload()
}

// Check checks the contents of an outline descriptor.
//
// An outline with no points or with a single point only is also valid.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_check
func (o *Outline) Check() bool {
	if o == nil || o.ptr == nil {
		return false
	}

	return C.FT_Outline_Check(o.ptr) == 0
}

// CBox returns an outline's ‘control box’.
// The control box encloses all the outline's points, including Bezier control
// points. Though it coincides with the exact bounding box for most glyphs, it
// can be slightly larger in some situations (like when rotating an outline that
// contains Bezier outside arcs).
//
// Computing the control box is very fast, while getting the bounding box can
// take much more time as it needs to walk over all segments and arcs in the
// outline. To get the latter, you can use the ‘ftbbox’ component, which is
// dedicated to this single task.
//
// See Glyph.CBox for a discussion of tricky fonts.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_get_cbox
func (o *Outline) CBox() BBox {
	if o == nil || o.ptr == nil {
		return BBox{}
	}

	var cbox C.FT_BBox
	C.FT_Outline_Get_CBox(o.ptr, &cbox)

	return BBox{
		XMin: Pos(cbox.xMin),
		XMax: Pos(cbox.xMax),
		YMin: Pos(cbox.yMin),
		YMax: Pos(cbox.yMax),
	}
}

// BBox computes the exact bounding box of an outline. This is slower than
// computing the control box. However, it uses an advanced algorithm that returns
// very quickly when the two boxes coincide. Otherwise, the outline Bezier arcs
// are traversed to extract their extrema.
//
// If the font is tricky and the glyph has been loaded with LoadNoScale, the
// resulting BBox is meaningless. To get reasonable values for the BBox it is
// necessary to load the glyph at a large ppem value (so that the hinting
// instructions can properly shift and scale the subglyphs), then extracting the
// BBox, which can be eventually converted back to font units.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_get_bbox
func (o *Outline) BBox() BBox {
	if o == nil || o.ptr == nil {
		return BBox{}
	}

	var bbox C.FT_BBox
	C.FT_Outline_Get_BBox(o.ptr, &bbox)

	return BBox{
		XMin: Pos(bbox.xMin),
		XMax: Pos(bbox.xMax),
		YMin: Pos(bbox.yMin),
		YMax: Pos(bbox.yMax),
	}
}

// Bitmap renders the outline within a bitmap.
// The outline's image is simply OR-ed to the target bitmap.
//
// This function does not create the bitmap, it only renders an outline image
// within the one you pass to it! Consequently, the various fields in it should
// be set accordingly.
//
// It will use the raster corresponding to the default glyph format.
//
// The value of the NumGrays field in b is ignored. If you select the gray-level
// rasterizer, and you want less than 256 gray levels, you have to use Render
// directly.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_get_bitmap
// TODO
// func (o *Outline) Bitmap(b *Bitmap) error {
// 	if o == nil || o.ptr == nil {
// 		return ErrInvalidOutline
// 	}

// 	if b == nil || b.ptr == nil {
// 		return ErrInvalidArgument
// 	}

// 	if o.l == nil || o.l.ptr == nil {
// 		return ErrInvalidLibraryHandle
// 	}

// 	if err := getErr(C.FT_Outline_Get_Bitmap(o.l.ptr, o.ptr, b.ptr)); err != nil {
// 		return err
// 	}

// 	b.reload()
// 	return nil
// }

// FT_Outline_Render TODO
// FT_Outline_Decompose TODO

// OutlineDecomposer is used during outline decomposition in order to emit
// segments, conic, and cubic Beziers.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_funcs
type OutlineDecomposer interface {
	// The ‘move to’ emitter.
	MoveTo(Vector) error
	// The segment emitter.
	LineTo(Vector) error
	// The second-order Bezier arc emitter.
	ConicTo(control, to Vector) error
	// The third-order Bezier arc emitter.
	CubicTo(control1, control2, to Vector) error
}

type decomposerTable struct {
	sync.Mutex
	idx   uintptr
	table map[uintptr]OutlineDecomposer
}

func (m *decomposerTable) acquire(d OutlineDecomposer, shift int, delta Pos) (*C.FT_Outline_Funcs, uintptr, error) {
	m.Lock()
	var idx uintptr
	start := m.idx
	for {
		if _, occupied := m.table[m.idx]; occupied {
			m.idx++
			if m.idx == start {
				m.Unlock()
				return nil, 0, errors.New("overflow")
			}

			continue
		}

		idx = m.idx
		m.table[idx] = d
		break
	}
	m.Unlock()

	return &C.FT_Outline_Funcs{
		move_to:  (*[0]byte)(C.OutlineMoveToCallback),
		line_to:  (*[0]byte)(C.OutlineLineToCallback),
		conic_to: (*[0]byte)(C.OutlineConicToCallback),
		cubic_to: (*[0]byte)(C.OutlineCubicToCallback),
		shift:    C.int(shift),
		delta:    C.FT_Pos(delta),
	}, idx, nil
}

func (m *decomposerTable) valueOf(idx uintptr) OutlineDecomposer {
	m.Lock()
	v := m.table[idx]
	m.Unlock()
	return v
}
func (m *decomposerTable) release(idx uintptr) {
	m.Lock()
	m.table[idx] = nil
	m.Unlock()
}

var decomposers = &decomposerTable{table: make(map[uintptr]OutlineDecomposer)}

// Decompose walks over an outline's structure to decompose it into individual
// segments and Bezier arcs. This function also emits ‘move to’ operations to
// indicate the start of new contours in the outline.
//
// Shift is the shift that is applied to coordinates before they are sent to the
// emitter.
//
// Delta is the delta that is applied to coordinates before they are sent to the
// emitter, but after the shift.
//
// The point coordinates sent to the emitters are the transformed version of the
// original coordinates (this is important for high accuracy during
// scan-conversion). The transformation is simple:
//
//	x' := (x << shift) - delta
//	y' := (y << shift) - delta
//
// Set the values of shift and delta to 0 to get the original point coordinates.
//
// A contour that contains a single point only is represented by a ‘move to’
// operation followed by ‘line to’ to the same point. In most cases, it is best
// to filter this out before using the outline for stroking purposes (otherwise
// it would result in a visible dot when round caps are used).
//
// Similarly, the function returns success for an empty outline also (doing
// nothing, this is, not calling any emitter); if necessary, you should filter
// this out, too.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_decompose
func (o *Outline) Decompose(decomposer OutlineDecomposer, shift int, delta Pos) error {
	if o == nil || o.ptr == nil {
		return ErrInvalidOutline
	}

	funcs, handle, err := decomposers.acquire(decomposer, shift, delta)
	if err != nil {
		return err
	}

	return getErr(C.FT_Outline_Decompose(o.ptr, funcs, unsafe.Pointer(handle)))
}

// Orientation is used to describe an outline's contour orientation.
//
// The TrueType and PostScript specifications use different conventions to
// determine whether outline contours should be filled or unfilled.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_orientation
type Orientation uint

const (
	// OrientationNone the orientation cannot be determined.
	// That is, different parts of the glyph have different orientation.
	OrientationNone Orientation = C.FT_ORIENTATION_NONE
	// OrientationTruetype according to the TrueType specification, clockwise
	// contours must be filled, and counter-clockwise ones must be unfilled.
	OrientationTruetype Orientation = C.FT_ORIENTATION_TRUETYPE
	// OrientationPostscript according to the PostScript specification,
	// counter-clockwise contours must be filled, and clockwise ones must be unfilled.
	OrientationPostscript Orientation = C.FT_ORIENTATION_POSTSCRIPT
	// OrientationFillRight is identical to OrientationTruetype, but is used to
	// remember that in TrueType, everything that is to the right of the drawing
	// direction of a contour must be filled.
	OrientationFillRight Orientation = C.FT_ORIENTATION_FILL_RIGHT
	// OrientationFillLeft is identical to OrientationPostscript, but is used to
	// remember that in PostScript, everything that is to the left of the
	// drawing direction of a contour must be filled.
	OrientationFillLeft Orientation = C.FT_ORIENTATION_FILL_LEFT
)

// Orientation analyzes a glyph outline and tries to compute its fill orientation.
// This is done by integrating the total area covered by the outline.
// The positive integral corresponds to the clockwise orientation and OrientationPostscript
// is returned. The negative integral corresponds to the counter-clockwise
// orientation and OrientationTruetype is returned.
//
// Note that it will return OrientationTruetype for outlines with 0 points.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_get_orientation
func (o *Outline) Orientation() Orientation {
	if o == nil || o.ptr == nil {
		return OrientationNone
	}

	return Orientation(C.FT_Outline_Get_Orientation(o.ptr))
}

// OutlineFlag is a list of bit-field constants used for the flags in an Outline flags field.
//
// The flags OutlineIgnoreDropouts, OutlineSmartDropouts, and OutlineIncludeStubs
// are ignored by the smooth rasterizer.
//
// There exists a second mechanism to pass the drop-out mode to the B/W rasterizer;
// see the tags field in Outline.
//
// Please refer to the description of the ‘SCANTYPE’ instruction in the OpenType
// specification (in file ttinst1.doc) how simple drop-outs, smart drop-outs,
// and stubs are defined.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_xxx
type OutlineFlag uint

const (
	// OutlineNone is reserved.
	OutlineNone OutlineFlag = C.FT_OUTLINE_NONE

	// OutlineOwner if set, this flag indicates that the outline's field arrays
	// (i.e., points, flags, and contours) are ‘owned’ by the outline object,
	// and should thus be freed when it is destroyed.
	OutlineOwner OutlineFlag = C.FT_OUTLINE_OWNER
	// OutlineEvenOddFill by default, outlines are filled using the non-zero
	// winding rule. If set to 1, the outline will be filled using the even-odd
	// fill rule (only works with the smooth rasterizer).
	OutlineEvenOddFill OutlineFlag = C.FT_OUTLINE_EVEN_ODD_FILL
	// OutlineReverseFill by default, outside contours of an outline are
	// oriented in clock-wise direction, as defined in the TrueType
	// specification. This flag is set if the outline uses the opposite
	// direction (typically for Type 1 fonts). This flag is ignored by the scan
	// converter.
	OutlineReverseFill OutlineFlag = C.FT_OUTLINE_REVERSE_FILL
	// OutlineIgnoreDropouts by default, the scan converter will try to detect
	// drop-outs in an outline and correct the glyph bitmap to ensure consistent
	// shape continuity. If set, this flag hints the scan-line converter to
	// ignore such cases.
	OutlineIgnoreDropouts OutlineFlag = C.FT_OUTLINE_IGNORE_DROPOUTS
	// OutlineSmartDropouts select smart dropout control. If unset, use simple
	// dropout control. Ignored if OutlineIgnoreDropouts is set.
	OutlineSmartDropouts OutlineFlag = C.FT_OUTLINE_SMART_DROPOUTS
	// OutlineIncludeStubs if set, turn pixels on for ‘stubs’, otherwise exclude
	// them. Ignored if OutlineIgnoreDropouts is set.
	OutlineIncludeStubs OutlineFlag = C.FT_OUTLINE_INCLUDE_STUBS
	// OutlineHighPrecision indicates that the scan-line converter should try to
	// convert this outline to bitmaps with the highest possible quality.
	// It is typically set for small character sizes. Note that this is only a
	// hint that might be completely ignored by a given scan-converter.
	OutlineHighPrecision OutlineFlag = C.FT_OUTLINE_HIGH_PRECISION
	// OutlineSinglePass is set to force a given scan-converter to only use a
	// single pass over the outline to render a bitmap glyph image. Normally,
	// it is set for very large character sizes. It is only a hint that might be
	// completely ignored by a given scan-converter.
	OutlineSinglePass OutlineFlag = C.FT_OUTLINE_SINGLE_PASS
)

func (x OutlineFlag) String() string {
	// the maximum concatenated len, at the time of writing, is 97.
	s := make([]byte, 0, 97)

	if x&OutlineOwner == OutlineOwner {
		s = append(s, []byte("Owner|")...)
	}
	if x&OutlineEvenOddFill == OutlineEvenOddFill {
		s = append(s, []byte("EvenOddFill|")...)
	}
	if x&OutlineReverseFill == OutlineReverseFill {
		s = append(s, []byte("ReverseFill|")...)
	}
	if x&OutlineIgnoreDropouts == OutlineIgnoreDropouts {
		s = append(s, []byte("IgnoreDropouts|")...)
	}
	if x&OutlineSmartDropouts == OutlineSmartDropouts {
		s = append(s, []byte("SmartDropouts|")...)
	}
	if x&OutlineIncludeStubs == OutlineIncludeStubs {
		s = append(s, []byte("IncludeStubs|")...)
	}
	if x&OutlineHighPrecision == OutlineHighPrecision {
		s = append(s, []byte("HighPrecision|")...)
	}
	if x&OutlineSinglePass == OutlineSinglePass {
		s = append(s, []byte("SinglePass|")...)
	}

	if len(s) == 0 {
		return ""
	}

	return string(s[:len(s)-1]) // trim the leading |
}
