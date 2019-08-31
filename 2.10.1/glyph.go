package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_GLYPH_H
import "C"

import (
	"unsafe"

	"github.com/flga/freetype2/fixed"
)

// BBoxMode controls how the values of Glyph.CBox() are returned.
type BBoxMode uint

const (
	// GlyphBBoxUnscaled unscaled font units.
	GlyphBBoxUnscaled BBoxMode = C.FT_GLYPH_BBOX_UNSCALED
	// GlyphBBoxSubpixels unfitted 26.6 coordinates.
	GlyphBBoxSubpixels BBoxMode = C.FT_GLYPH_BBOX_SUBPIXELS
	// GlyphBBoxGridfit grid-fitted 26.6 coordinates.
	GlyphBBoxGridfit BBoxMode = C.FT_GLYPH_BBOX_GRIDFIT
	// GlyphBBoxTruncate coordinates in integer pixels.
	GlyphBBoxTruncate BBoxMode = C.FT_GLYPH_BBOX_TRUNCATE
	// GlyphBBoxPixels grid-fitted pixel coordinates.
	GlyphBBoxPixels BBoxMode = C.FT_GLYPH_BBOX_PIXELS
)

// Glyph is used to model generic glyph images.
//
// Glyph objects are not owned by the library. You must thus release them manually
// (through Free()) before calling Library.Free().
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph
type Glyph interface {
	getptr() C.FT_Glyph
	reload()
	reset()

	// Free destroys the glyph
	//
	// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_done_glyph
	Free()
	// Format returns the format of the glyph's image.
	// You can use it to get the underlying type of Glyph.
	Format() GlyphFormat
	// Advance returns a 16.16 vector that gives the glyph's advance width.
	Advance() Vector16_16
	// Copy copies a glyph image. Note that the created Glyph must be released
	// with Free().
	//
	// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_copy
	Copy() (Glyph, error)
	// Transform applies a transformation to a glyph image if its format is scalable.
	// The 2x2 transformation matrix is also applied to the glyph's advance vector.
	//
	// If the format is not scalable it returns ErrInvalidGlyphFormat.
	//
	// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_transform
	Transform(matrix Matrix, delta Vector26_6) error
	// CBox returns the glyph's ‘control box’. The control box encloses all the
	// outline's points, including Bezier control points. Though it coincides
	// with the exact bounding box for most glyphs, it can be slightly larger in
	// some situations (like when rotating an outline that contains Bezier
	// outside arcs).
	//
	// Computing the control box is very fast, while getting the bounding box can
	// take much more time as it needs to walk over all segments and arcs in the
	// outline. To get the latter, you can use the ‘ftbbox’ component, which is
	// dedicated to this single task.
	//
	// Coordinates are expressed in 1/64th of pixels if it is grid-fitted.
	//
	// Coordinates are relative to the glyph origin, using the y upwards convention.
	//
	// If the glyph has been loaded with LoadNoScale, mode must be set to
	// GlyphBBoxUnscaled to get unscaled font units in 26.6 pixel format.
	// The value GlyphBBoxSubpixels is another name for this constant.
	//
	// If the font is tricky and the glyph has been loaded with LoadNoScale, the
	// resulting CBox is meaningless. To get reasonable values for the CBox it
	// is necessary to load the glyph at a large ppem value (so that the hinting
	// instructions can properly shift and scale the subglyphs), then extracting
	// the CBox, which can be eventually converted back to font units.
	//
	// Note that the maximum coordinates are exclusive, which means that one can
	// compute the width and height of the glyph image (be it in integer or 26.6
	// pixels) as:
	//	width  := bbox.XMax - bbox.XMin;
	//	height := bbox.YMax - bbox.YMin;
	//
	// Note also that for 26.6 coordinates, if mode is set to GlyphBBoxGridfit,
	// the coordinates will also be grid-fitted, which corresponds to:
	//	bbox.XMin = floor(bbox.XMin);
	//	bbox.YMin = floor(bbox.YMin);
	//	bbox.XMax = ceil(bbox.XMax);
	//	bbox.YMax = ceil(bbox.YMax);
	//
	// To get the bbox in pixel coordinates, use GlyphBBoxTruncate.
	//
	// To get the bbox in grid-fitted pixel coordinates, use GlyphBBoxPixels.
	//
	// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_get_cbox
	CBox(mode BBoxMode) BBox
	// ToBitmap converts the glyph to a bitmap glyph object.
	//
	// The destroy argument indicates wether the original glyph image should be
	// destroyed by this function. It is never destroyed in case of error.
	//
	// The receiver will not be modified, but it might be freed if destroy is true,
	// which will render it unusable.
	//
	// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_to_bitmap
	ToBitmap(mode RenderMode, origin Vector26_6, destroy bool) (*BitmapGlyph, error)
}

func newGlyph(g C.FT_Glyph) (Glyph, error) {
	if g == nil {
		return nil, nil
	}

	var ret Glyph
	switch g.format {
	case C.FT_GLYPH_FORMAT_BITMAP:
		ret = &BitmapGlyph{ptr: g}
		ret.reload()
		return ret, nil
	case C.FT_GLYPH_FORMAT_OUTLINE:
		ret = &OutlineGlyph{ptr: g}
		ret.reload()
		return ret, nil
	}

	return nil, ErrInvalidGlyphFormat
}

func glyphFree(g Glyph) {
	if g == nil || g.getptr() == nil {
		return
	}

	C.FT_Done_Glyph(g.getptr())
	g.reset()
}

func glyphCopy(g Glyph) (Glyph, error) {
	if g == nil || g.getptr() == nil {
		return nil, ErrInvalidArgument
	}

	var target C.FT_Glyph
	if err := getErr(C.FT_Glyph_Copy(g.getptr(), &target)); err != nil {
		return nil, err
	}

	return newGlyph(target)
}

func glyphTransform(g Glyph, matrix Matrix, delta Vector26_6) error {
	if g == nil || g.getptr() == nil {
		return ErrInvalidArgument
	}

	var cmatrix *C.FT_Matrix
	if matrix != (Matrix{}) {
		cmatrix = &C.FT_Matrix{
			xx: C.FT_Fixed(matrix.Xx),
			xy: C.FT_Fixed(matrix.Xy),
			yx: C.FT_Fixed(matrix.Yx),
			yy: C.FT_Fixed(matrix.Yy),
		}
	}

	var cdelta *C.FT_Vector
	if delta != (Vector26_6{}) {
		cdelta = &C.FT_Vector{
			x: C.FT_Pos(delta.X),
			y: C.FT_Pos(delta.Y),
		}
	}
	if err := getErr(C.FT_Glyph_Transform(g.getptr(), cmatrix, cdelta)); err != nil {
		return err
	}
	g.reload()

	return nil
}

func glyphCBox(g Glyph, mode BBoxMode) BBox {
	if g == nil || g.getptr() == nil {
		return BBox{}
	}

	var acbox C.FT_BBox
	C.FT_Glyph_Get_CBox(g.getptr(), C.uint(mode), &acbox)

	return BBox{
		XMin: Pos(acbox.xMin),
		YMin: Pos(acbox.yMin),
		XMax: Pos(acbox.xMax),
		YMax: Pos(acbox.yMax),
	}
}

func glyphToBitmap(g Glyph, mode RenderMode, origin Vector26_6, destroy bool) (*BitmapGlyph, error) {
	if g == nil || g.getptr() == nil {
		return nil, ErrInvalidArgument
	}

	var corigin *C.FT_Vector
	if origin != (Vector26_6{}) {
		corigin = &C.FT_Vector{
			x: C.FT_Pos(origin.X),
			y: C.FT_Pos(origin.Y),
		}
	}

	var cdestroy C.FT_Bool
	if destroy {
		cdestroy = 1
	}

	target := g.getptr()
	if err := getErr(C.FT_Glyph_To_Bitmap(&target, C.FT_Render_Mode(mode), corigin, cdestroy)); err != nil {
		return nil, err
	}

	if destroy {
		g.reset()
	}
	g.reload()

	ret := &BitmapGlyph{ptr: target}
	ret.reload()
	return ret, nil
}

var _ Glyph = &BitmapGlyph{}

// BitmapGlyph models a bitmap glyph image.
//
// You can cast a Glyph to *BitmapGlyph if you have glyph.Format() == GlyphFormatBitmap.
// This lets you access the bitmap's contents easily.
//
// Glyph objects are not owned by the library. You must thus release them manually
// (through Free()) before calling Library.Free().
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_bitmapglyphrec
type BitmapGlyph struct {
	ptr     C.FT_Glyph `deep:"-"`
	format  GlyphFormat
	advance Vector16_16

	// The left-side bearing, i.e., the horizontal distance from the current pen
	// position to the left border of the glyph bitmap.
	Left int
	// The top-side bearing, i.e., the vertical distance from the current pen
	// position to the top border of the glyph bitmap. This distance is positive
	// for upwards y!
	Top int
	// A descriptor for the bitmap.
	Bitmap *Bitmap
}

func (g *BitmapGlyph) getptr() C.FT_Glyph {
	if g == nil {
		return nil
	}

	return g.ptr
}

func (g *BitmapGlyph) reload() {
	if g == nil {
		return
	}

	if g.ptr == nil {
		*g = BitmapGlyph{}
		return
	}

	if g.ptr.format != C.FT_GLYPH_FORMAT_BITMAP {
		panic("ptr is not a bitmap")
	}

	ptr := (C.FT_BitmapGlyph)(unsafe.Pointer(g.ptr))
	*g = BitmapGlyph{
		ptr:    g.ptr,
		format: GlyphFormat(g.ptr.format),
		advance: Vector16_16{
			X: fixed.Int16_16(g.ptr.advance.x),
			Y: fixed.Int16_16(g.ptr.advance.y),
		},
		Left:   int(ptr.left),
		Top:    int(ptr.top),
		Bitmap: &Bitmap{ptr: &ptr.bitmap},
	}
	g.Bitmap.reload()
}

func (g *BitmapGlyph) reset() {
	if g == nil {
		return
	}
	if g.Bitmap != nil {
		*g.Bitmap = Bitmap{}
	}
	*g = BitmapGlyph{}
}

// Free destroys the glyph
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_done_glyph
func (g *BitmapGlyph) Free() {
	glyphFree(g)
}

// Format returns the format of the glyph's image.
// You can use it to get the underlying type of Glyph.
func (g *BitmapGlyph) Format() GlyphFormat {
	if g == nil || g.ptr == nil {
		return 0
	}
	return g.format
}

// Advance returns a 16.16 vector that gives the glyph's advance width.
func (g *BitmapGlyph) Advance() Vector16_16 {
	if g == nil || g.ptr == nil {
		return Vector16_16{}
	}
	return g.advance
}

// Copy copies a glyph image. Note that the created Glyph must be released
// with Free().
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_copy
func (g *BitmapGlyph) Copy() (Glyph, error) {
	return glyphCopy(g)
}

// Transform applies a transformation to a glyph image if its format is scalable.
// The 2x2 transformation matrix is also applied to the glyph's advance vector.
//
// If the format is not scalable it returns ErrInvalidGlyphFormat.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_transform
func (g *BitmapGlyph) Transform(matrix Matrix, delta Vector26_6) error {
	return glyphTransform(g, matrix, delta)
}

// CBox returns the glyph's ‘control box’. The control box encloses all the
// outline's points, including Bezier control points. Though it coincides
// with the exact bounding box for most glyphs, it can be slightly larger in
// some situations (like when rotating an outline that contains Bezier
// outside arcs).
//
// Computing the control box is very fast, while getting the bounding box can
// take much more time as it needs to walk over all segments and arcs in the
// outline. To get the latter, you can use the ‘ftbbox’ component, which is
// dedicated to this single task.
//
// Coordinates are expressed in 1/64th of pixels if it is grid-fitted.
//
// Coordinates are relative to the glyph origin, using the y upwards convention.
//
// If the glyph has been loaded with LoadNoScale, mode must be set to
// GlyphBBoxUnscaled to get unscaled font units in 26.6 pixel format.
// The value GlyphBBoxSubpixels is another name for this constant.
//
// If the font is tricky and the glyph has been loaded with LoadNoScale, the
// resulting CBox is meaningless. To get reasonable values for the CBox it
// is necessary to load the glyph at a large ppem value (so that the hinting
// instructions can properly shift and scale the subglyphs), then extracting
// the CBox, which can be eventually converted back to font units.
//
// Note that the maximum coordinates are exclusive, which means that one can
// compute the width and height of the glyph image (be it in integer or 26.6
// pixels) as:
//	width  := bbox.XMax - bbox.XMin;
//	height := bbox.YMax - bbox.YMin;
//
// Note also that for 26.6 coordinates, if mode is set to GlyphBBoxGridfit,
// the coordinates will also be grid-fitted, which corresponds to:
//	bbox.XMin = floor(bbox.XMin);
//	bbox.YMin = floor(bbox.YMin);
//	bbox.XMax = ceil(bbox.XMax);
//	bbox.YMax = ceil(bbox.YMax);
//
// To get the bbox in pixel coordinates, use GlyphBBoxTruncate.
//
// To get the bbox in grid-fitted pixel coordinates, use GlyphBBoxPixels.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_get_cbox
func (g *BitmapGlyph) CBox(mode BBoxMode) BBox {
	return glyphCBox(g, mode)
}

// ToBitmap converts the glyph to a bitmap glyph object.
//
// The destroy argument indicates wether the original glyph image should be
// destroyed by this function. It is never destroyed in case of error.
//
// The receiver will not be modified, but it might be freed if destroy is true,
// which will render it unusable.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_to_bitmap
func (g *BitmapGlyph) ToBitmap(mode RenderMode, origin Vector26_6, destroy bool) (*BitmapGlyph, error) {
	return glyphToBitmap(g, mode, origin, destroy)
}

var _ Glyph = &OutlineGlyph{}

// OutlineGlyph models an outline (vectorial) glyph image.
//
// You can cast a Glyph to *OutlineGlyph if you have glyph.Format() == GlyphFormatOutline.
// This lets you access the outline's contents easily.
//
// As the outline is extracted from a glyph slot, its coordinates are expressed
// normally in 26.6 pixels, unless the flag LoadNoScale was used in LoadGlyph or
// LoadChar.
//
// Glyph objects are not owned by the library. You must thus release them manually
// (through Free()) before calling Library.Free().
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_outlineglyphrec
type OutlineGlyph struct {
	ptr     C.FT_Glyph `deep:"-"`
	format  GlyphFormat
	advance Vector16_16

	// A descriptor for the outline.
	Outline *Outline
}

func (g *OutlineGlyph) getptr() C.FT_Glyph {
	if g == nil {
		return nil
	}

	return g.ptr
}

func (g *OutlineGlyph) reload() {
	if g == nil {
		return
	}

	if g.ptr == nil {
		*g = OutlineGlyph{}
		return
	}

	if g.ptr.format != C.FT_GLYPH_FORMAT_OUTLINE {
		panic("ptr is not an outline")
	}

	ptr := (C.FT_OutlineGlyph)(unsafe.Pointer(g.ptr))
	*g = OutlineGlyph{
		ptr:    g.ptr,
		format: GlyphFormat(g.ptr.format),
		advance: Vector16_16{
			X: fixed.Int16_16(g.ptr.advance.x),
			Y: fixed.Int16_16(g.ptr.advance.y),
		},
		Outline: &Outline{ptr: &ptr.outline},
	}
	g.Outline.reload()
}

func (g *OutlineGlyph) reset() {
	if g == nil {
		return
	}
	if g.Outline != nil {
		*g.Outline = Outline{}
	}
	*g = OutlineGlyph{}
}

// Free destroys the glyph
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_done_glyph
func (g *OutlineGlyph) Free() {
	glyphFree(g)
}

// Format returns the format of the glyph's image.
// You can use it to get the underlying type of Glyph.
func (g *OutlineGlyph) Format() GlyphFormat {
	if g == nil || g.ptr == nil {
		return 0
	}
	return g.format
}

// Advance returns a 16.16 vector that gives the glyph's advance width.
func (g *OutlineGlyph) Advance() Vector16_16 {
	if g == nil || g.ptr == nil {
		return Vector16_16{}
	}
	return g.advance
}

// Copy copies a glyph image. Note that the created Glyph must be released
// with Free().
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_copy
func (g *OutlineGlyph) Copy() (Glyph, error) {
	return glyphCopy(g)
}

// Transform applies a transformation to a glyph image if its format is scalable.
// The 2x2 transformation matrix is also applied to the glyph's advance vector.
//
// If the format is not scalable it returns ErrInvalidGlyphFormat.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_transform
func (g *OutlineGlyph) Transform(matrix Matrix, delta Vector26_6) error {
	return glyphTransform(g, matrix, delta)
}

// CBox returns the glyph's ‘control box’. The control box encloses all the
// outline's points, including Bezier control points. Though it coincides
// with the exact bounding box for most glyphs, it can be slightly larger in
// some situations (like when rotating an outline that contains Bezier
// outside arcs).
//
// Computing the control box is very fast, while getting the bounding box can
// take much more time as it needs to walk over all segments and arcs in the
// outline. To get the latter, you can use the ‘ftbbox’ component, which is
// dedicated to this single task.
//
// Coordinates are expressed in 1/64th of pixels if it is grid-fitted.
//
// Coordinates are relative to the glyph origin, using the y upwards convention.
//
// If the glyph has been loaded with LoadNoScale, mode must be set to
// GlyphBBoxUnscaled to get unscaled font units in 26.6 pixel format.
// The value GlyphBBoxSubpixels is another name for this constant.
//
// If the font is tricky and the glyph has been loaded with LoadNoScale, the
// resulting CBox is meaningless. To get reasonable values for the CBox it
// is necessary to load the glyph at a large ppem value (so that the hinting
// instructions can properly shift and scale the subglyphs), then extracting
// the CBox, which can be eventually converted back to font units.
//
// Note that the maximum coordinates are exclusive, which means that one can
// compute the width and height of the glyph image (be it in integer or 26.6
// pixels) as:
//	width  := bbox.XMax - bbox.XMin;
//	height := bbox.YMax - bbox.YMin;
//
// Note also that for 26.6 coordinates, if mode is set to GlyphBBoxGridfit,
// the coordinates will also be grid-fitted, which corresponds to:
//	bbox.XMin = floor(bbox.XMin);
//	bbox.YMin = floor(bbox.YMin);
//	bbox.XMax = ceil(bbox.XMax);
//	bbox.YMax = ceil(bbox.YMax);
//
// To get the bbox in pixel coordinates, use GlyphBBoxTruncate.
//
// To get the bbox in grid-fitted pixel coordinates, use GlyphBBoxPixels.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_get_cbox
func (g *OutlineGlyph) CBox(mode BBoxMode) BBox {
	return glyphCBox(g, mode)
}

// ToBitmap converts the glyph to a bitmap glyph object.
//
// The destroy argument indicates wether the original glyph image should be
// destroyed by this function. It is never destroyed in case of error.
//
// The receiver will not be modified, but it might be freed if destroy is true,
// which will render it unusable.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_glyph_to_bitmap
func (g *OutlineGlyph) ToBitmap(mode RenderMode, origin Vector26_6, destroy bool) (*BitmapGlyph, error) {
	return glyphToBitmap(g, mode, origin, destroy)
}

// NewGlyph creates a new empty glyph image. Note that the created Glyph must be
// released with Free.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_new_glyph
func (l *Library) NewGlyph(format GlyphFormat) (Glyph, error) {
	if l == nil || l.ptr == nil {
		return nil, ErrInvalidArgument
	}

	var aglyph C.FT_Glyph
	if err := getErr(C.FT_New_Glyph(l.ptr, C.FT_Glyph_Format(format), &aglyph)); err != nil {
		return nil, err
	}

	return newGlyph(aglyph)
}

// Glyph extracts a glyph image from the slot. Note that the created Glyph object
// must be released with Free.
//
// Because glyph.Advance().X and glyph.Advance().Y are 16.16 fixed-point numbers,
// slot.Advance.X and slot.Advance.Y (which are in 26.6 fixed-point format) must
// be in the range ]-32768;32768[.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-glyph_management.html#ft_get_glyph
func (s *GlyphSlot) Glyph() (Glyph, error) {
	if s == nil || s.ptr == nil {
		return nil, ErrInvalidSlotHandle
	}

	var aglyph C.FT_Glyph
	if err := getErr(C.FT_Get_Glyph(s.ptr, &aglyph)); err != nil {
		return nil, err
	}

	return newGlyph(aglyph)
}
