package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_TRUETYPE_TABLES_H
import "C"
import (
	"unsafe"

	"github.com/flga/freetype2/2.10.1/fixed"
)

// GlyphIndex is the index of the glyph in the font file. For CID-keyed fonts
// (either in PS or in CFF format) it specifies the CID value.
type GlyphIndex uint

// MissingGlyph is the GlyphIndex for the undefined char code.
const MissingGlyph GlyphIndex = 0

// Face models a given typeface, in a given style.
//
// A Face object can only be safely used from one goroutine at a time. Similarly, creation and destruction of a Face
// with the same Library object can only be done from one goroutine at a time. On the other hand, functions like
// LoadGlyph and its siblings are thread-safe and do not need the lock to be held as long as the same Face object is not
// used from multiple goroutines at the same time.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_face
type Face struct {
	ptr     C.FT_Face
	dealloc []func()
}

// Free discards the face, as well as all of its child slots and sizes.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_done_face
func (f *Face) Free() error {
	if f == nil || f.ptr == nil {
		return nil
	}

	if err := getErr(C.FT_Done_Face(f.ptr)); err != nil {
		return err
	}
	f.freeInternal()
	return nil
}

func (f *Face) freeInternal() {
	if f == nil || f.ptr == nil {
		return
	}

	for _, fn := range f.dealloc {
		fn()
	}

	f.ptr = nil
}

// NumFaces returns the number of faces in the font file. Some font formats can have multiple faces in a single font file.
func (f *Face) NumFaces() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.num_faces)
}

// Index returns the index of the given face in the font file.
func (f *Face) Index() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.face_index) & 0xFFFF
}

// NamedIndex returns the named instance index for the current face index
// (starting with value 1; value 0 makes FreeType ignore named instances).
func (f *Face) NamedIndex() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.face_index) >> 16
}

// Flags returns the set of bit flags that give important information about the face.
func (f *Face) Flags() FaceFlag {
	if f == nil || f.ptr == nil {
		return 0
	}
	return FaceFlag(f.ptr.face_flags)
}

// HasFlag reports whether the face has the given flag.
func (f *Face) HasFlag(flag FaceFlag) bool { return f.Flags()&flag > 0 }

// Style returns the set of bit flags indicating the style of the face; see StyleFlag for details.
//
// [Since 2.6.1] Bits 16-30 hold the number of named instances available for the current face if we have a GX or
// OpenType variation (sub)font. Bit 31 is always zero (this is, it is always a positive value). Note that a variation
// font has always at least one named instance, namely the default instance.
func (f *Face) Style() StyleFlag {
	if f == nil || f.ptr == nil {
		return 0
	}
	return StyleFlag(f.ptr.style_flags)
}

// HasStyle reports whether the face has the given style flag.
func (f *Face) HasStyle(flag StyleFlag) bool { return f.Style()&flag > 0 }

// NumNamedInstances reports the number of available named instances available
// for the current face if we have a GX or OpenType variation (sub)font.
func (f *Face) NumNamedInstances() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.style_flags) >> 16
}

// NumGlyphs returns the number of glyphs in the face. If the face is scalable and has sbits (see NumFixedSizes), it is
// set to the number of outline glyphs.
//
// For CID-keyed fonts (not in an SFNT wrapper) this value gives the highest CID used in the font.
func (f *Face) NumGlyphs() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.num_glyphs)
}

// FamilyName returns the face's family name. This is an ASCII string, usually in English, that describes the typeface's
// family (like ‘Times New Roman’, ‘Bodoni’, ‘Garamond’, etc). This is a least common denominator used to list fonts.
// Some formats (TrueType & OpenType) provide localized and Unicode versions of this string.
// Applications should use the format-specific interface to access them.
// The returned value can be empty (e.g., in fonts embedded in a PDF file).
//
// In case the font doesn't provide a specific family name entry, FreeType tries to synthesize one, deriving it from
// other name entries.
func (f *Face) FamilyName() string {
	if f == nil || f.ptr == nil {
		return ""
	}
	return C.GoString(f.ptr.family_name)
}

// StyleName returns the face's style name. This is an ASCII string, usually in English, that describes the typeface's
// style (like ‘Italic’, ‘Bold’, ‘Condensed’, etc).
// Not all font formats provide a style name.
// Some formats provide localized and Unicode versions of this string.
// Applications should use the format-specific interface to access them.
func (f *Face) StyleName() string {
	if f == nil || f.ptr == nil {
		return ""
	}
	return C.GoString(f.ptr.style_name)
}

// NumFixedSizes reports the number of bitmap strikes in the face. Even if the face is scalable, there might still be
// bitmap strikes, which are called ‘sbits’ in that case.
func (f *Face) NumFixedSizes() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.num_fixed_sizes)
}

// AvailableSizes returns a copy of the list of BitmapSize for all bitmap strikes in the face.
// It returns nil if there is no bitmap strike.
//
// Note that FreeType tries to sanitize the strike data since they are sometimes sloppy or incorrect, but this can
// easily fail.
func (f *Face) AvailableSizes() []BitmapSize {
	if f == nil || f.ptr == nil {
		return nil
	}

	n := int(f.ptr.num_fixed_sizes)
	if n == 0 {
		return nil
	}

	ret := make([]BitmapSize, n)
	ptr := (*[(1<<31 - 1) / C.sizeof_FT_Bitmap_Size]C.FT_Bitmap_Size)(unsafe.Pointer(f.ptr.available_sizes))[:n:n]
	for i := range ret {
		ret[i] = BitmapSize{
			Height: int(ptr[i].height),
			Width:  int(ptr[i].width),
			Size:   fixed.Int26_6(ptr[i].size),
			XPpem:  fixed.Int26_6(ptr[i].x_ppem),
			YPpem:  fixed.Int26_6(ptr[i].y_ppem),
		}
	}
	return ret
}

// NumCharMaps reports the number of charmaps in the face.
func (f *Face) NumCharMaps() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.num_charmaps)
}

// CharMaps returns a copy of the charmaps of the face.
func (f *Face) CharMaps() []CharMap {
	if f == nil || f.ptr == nil {
		return nil
	}

	n := int(f.ptr.num_charmaps)
	if n == 0 {
		return nil
	}

	ret := make([]CharMap, n)
	for i, v := range f.charmaps() {
		ret[i] = newCharMap(v)
	}
	return ret
}

func (f *Face) charmaps() []C.FT_CharMap {
	return (*[(1<<31 - 1) / C.sizeof_FT_CharMap]C.FT_CharMap)(unsafe.Pointer(f.ptr.charmaps))[:f.ptr.num_charmaps:f.ptr.num_charmaps]
}

// BBox returns a copy of the font bounding box. Coordinates are expressed in font units (see UnitsPerEM).
// The box is large enough to contain any glyph from the font. Thus, bbox.YMax can be seen as the ‘maximum ascender’,
// and bbox.YMin as the ‘minimum descender’. Only relevant for scalable formats.
//
// Note that the bounding box might be off by (at least) one pixel for hinted fonts.
// See SizeMetrics for further discussion.
func (f *Face) BBox() BBox {
	if f == nil || f.ptr == nil {
		return BBox{}
	}

	return BBox{
		XMin: Pos(f.ptr.bbox.xMin),
		YMin: Pos(f.ptr.bbox.yMin),
		XMax: Pos(f.ptr.bbox.xMax),
		YMax: Pos(f.ptr.bbox.yMax),
	}
}

// UnitsPerEM reports the number of font units per EM square for this face. This is typically 2048 for TrueType fonts,
// and 1000 for Type 1 fonts.
// Only relevant for scalable formats.
func (f *Face) UnitsPerEM() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.units_per_EM)
}

// Ascender reports the typographic ascender of the face, expressed in font units. For font formats not having this
// information, it is set to bbox.yMax.
// Only relevant for scalable formats.
func (f *Face) Ascender() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.ascender)
}

// Descender reports the typographic descender of the face, expressed in font units. For font formats not having this
// information, it is set to bbox.yMin. Note that this field is negative for values below the baseline.
// Only relevant for scalable formats.
func (f *Face) Descender() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.descender)
}

// Height reports the vertical distance between two consecutive baselines, expressed in font units.
// It is always positive.
// Only relevant for scalable formats.
//
// If you want the global glyph height, use  ascender - descender.
func (f *Face) Height() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.height)
}

// MaxAdvanceWidth reports the maximum advance width, in font units, for all glyphs in this face. This can be used to
// make word wrapping computations faster.
// Only relevant for scalable formats.
func (f *Face) MaxAdvanceWidth() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.max_advance_width)
}

// MaxAdvanceHeight reports the maximum advance height, in font units, for all glyphs in this face. This is only
// relevant for vertical layouts, and is set to height for fonts that do not provide vertical metrics.
// Only relevant for scalable formats.
func (f *Face) MaxAdvanceHeight() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.max_advance_height)
}

// UnderlinePosition reports the position, in font units, of the underline line for this face. It is the center of the
// underlining stem.
// Only relevant for scalable formats.
func (f *Face) UnderlinePosition() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.underline_position)
}

// UnderlineThickness reports the thickness, in font units, of the underline for this face.
// Only relevant for scalable formats.
func (f *Face) UnderlineThickness() int {
	if f == nil || f.ptr == nil {
		return 0
	}
	return int(f.ptr.underline_thickness)
}

// Glyph returns a copy of the current contents, if any, of the face's glyph slot.
func (f *Face) Glyph() GlyphSlot {
	if f == nil || f.ptr == nil {
		return GlyphSlot{}
	}

	return newGlyphSlot(f.ptr.glyph)
}

// I don't know how to test this yet
//
// // Glyphs returns a copy of the current contents, if any, of the face's glyph
// // slot, following the linked list.
// func (f *Face) Glyphs() []GlyphSlot {
// 	if f == nil || f.ptr == nil {
// 		return nil
// 	}

// 	var ret []GlyphSlot
// 	ptr := f.ptr.glyph
// 	for ptr.next != nil {
// 		ret = append(ret, newGlyphSlot(ptr))
// 		ptr = ptr.next
// 	}
// 	return ret
// }

// Size returns a copy of the current active size for this face.
func (f *Face) Size() Size {
	if f == nil || f.ptr == nil {
		return Size{}
	}

	return newSize(f.ptr.size)
}

// ActiveCharMap returns a copy of the active charmap.
// If there is no active charmap, it returns the zero value and false.
func (f *Face) ActiveCharMap() (CharMap, bool) {
	if f == nil || f.ptr == nil {
		return CharMap{}, false
	}

	active := f.ptr.charmap
	if active == nil {
		return CharMap{}, false
	}

	return newCharMap(active), true
}

// PostscriptName returns the ASCII PostScript name of a given face, if
// available. This only works with PostScript, TrueType, and OpenType fonts.
//
// NOTE:
// For variation fonts, this string changes if you select a different instance,
// and you have to call PostscriptName again to retrieve it. FreeType follows
// Adobe TechNote #5902, ‘Generating PostScript Names for Fonts Using OpenType
// Font Variations’. https://download.macromedia.com/pub/developer/opentype/tech-notes/5902.AdobePSNameGeneration.html
//
// [Since 2.9] Special PostScript names for named instances are only returned if
// the named instance is set with SetNamedInstance (and the font has
// corresponding entries in its ‘fvar’ table). If face.HasFlag(FaceFlagVariation)
// returns true, the algorithmically derived PostScript name is provided, not
// looking up special entries for named instances.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_postscript_name
func (f *Face) PostscriptName() string {
	if f == nil || f.ptr == nil {
		return ""
	}

	return C.GoString(C.FT_Get_Postscript_Name(f.ptr))
}

// SetCharSize sets the character size for the face.
//
// If either the character width or height is zero, it is set equal to the other value.
// If either the horizontal or vertical resolution is zero, it is set equal to the other value.
//
// Don't use this function if you are using the FreeType cache API.
//
// NOTE:
// While this function allows fractional points as input values, the resulting ppem value for the given resolution
// is always rounded to the nearest integer.
// A character width or height smaller than 1pt is set to 1pt; if both resolution values are zero, they are set to 72dpi.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_set_char_size
func (f *Face) SetCharSize(nominalWidth, nominalHeight fixed.Int26_6, horzDPI, vertDPI uint) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	return getErr(C.FT_Set_Char_Size(f.ptr,
		C.FT_F26Dot6(nominalWidth),
		C.FT_F26Dot6(nominalHeight),
		C.FT_UInt(horzDPI),
		C.FT_UInt(vertDPI),
	))
}

// SetPixelSizes sets the face pixel size.
//
// Don't use this function if you are using the FreeType cache API.
//
// NOTE:
// You should not rely on the resulting glyphs matching or being constrained to this pixel size. Refer to FT_Request_Size to understand how requested sizes relate to actual sizes.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_set_pixel_sizes
func (f *Face) SetPixelSizes(width, height uint) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	return getErr(C.FT_Set_Pixel_Sizes(f.ptr,
		C.FT_UInt(width),
		C.FT_UInt(height),
	))
}

// SizeRequestType is an enumeration of the supported size request types, i.e.,
// what input size (in font units) maps to the requested output size (in pixels,
// as computed from the arguments of SizeRequest).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_size_request_type
type SizeRequestType int

const (
	// SizeRequestTypeNominal the UnitsPerEM method of Face is used to determine both scaling values.
	//
	// This is the standard scaling found in most applications. In particular, use this size request type for TrueType
	// fonts if they provide optical scaling or something similar. Note, however, that UnitsPerEM is a rather abstract
	// value which bears no relation to the actual size of the glyphs in a font.
	SizeRequestTypeNominal SizeRequestType = C.FT_SIZE_REQUEST_TYPE_NOMINAL
	// SizeRequestTypeRealDim the sum of the ascender and (minus of) the descender fields of Face is used to determine
	// both scaling values.
	SizeRequestTypeRealDim SizeRequestType = C.FT_SIZE_REQUEST_TYPE_REAL_DIM
	// SizeRequestTypeBBox the width and height of the BBox field of Face are used to determine the horizontal and
	// vertical scaling value, respectively.
	SizeRequestTypeBBox SizeRequestType = C.FT_SIZE_REQUEST_TYPE_BBOX
	// SizeRequestTypeCell the MaxAdvanceWidth field of Face is used to determine the horizontal scaling value; the
	// vertical scaling valueis determined the same way as SizeRequestTypeRealDim does. Finally, both scaling values are
	// set to the smaller one. This type is useful if you want to specify the font size for, say, a window of a given
	// dimension and 80x24 cells.
	SizeRequestTypeCell SizeRequestType = C.FT_SIZE_REQUEST_TYPE_CELL
	// SizeRequestTypeScales specify the scaling values directly.
	SizeRequestTypeScales SizeRequestType = C.FT_SIZE_REQUEST_TYPE_SCALES
)

func (s SizeRequestType) String() string {
	switch s {
	case SizeRequestTypeNominal:
		return "Nominal"
	case SizeRequestTypeRealDim:
		return "RealDim"
	case SizeRequestTypeBBox:
		return "BBox"
	case SizeRequestTypeCell:
		return "Cell"
	case SizeRequestTypeScales:
		return "Scales"
	default:
		return "Unknown"
	}
}

// SizeRequest models a size request.
//
// If type is SizeRequestTypeScales, width and height are interpreted directly as 16.16 fractional scaling values,
// without any further modification, and both horiResolution and vertResolution are ignored.
//
// NOTE:
// If width is zero, the horizontal scaling value is set equal to the vertical scaling value, and vice versa.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_size_requestrec
type SizeRequest struct {
	// See SizeRequestType
	Type SizeRequestType
	// The desired width, given as a 26.6 fractional point value (with 72pt = 1in).
	Width fixed.Int26_6
	// The desired height, given as a 26.6 fractional point value (with 72pt = 1in).
	Height fixed.Int26_6
	// The horizontal resolution (dpi, i.e., pixels per inch). If set to zero,
	// width is treated as a 26.6 fractional pixel value, which gets internally
	// rounded to an integer.
	HoriResolution uint
	// The vertical resolution (dpi, i.e., pixels per inch). If set to zero,
	// height is treated as a 26.6 fractional pixel value, which gets internally
	// rounded to an integer.
	VertResolution uint
}

// RequestSize resizes the scale of the active Size object.
//
// Don't use this function if you are using the FreeType cache API.
//
// Although drivers may select the bitmap strike matching the request, you should not rely on this if you intend to
// select a particular bitmap strike. Use SelectSize instead in that case.
//
// The relation between the requested size and the resulting glyph size is dependent entirely on how the size is defined
// in the source face. The font designer chooses the final size of each glyph relative to this size. For more
// information refer to ‘https://www.freetype.org/freetype2/docs/glyphs/glyphs-2.html’.
//
// Contrary to SetCharSize, this function doesn't have special code to normalize zero-valued widths, heights, or
// resolutions (which lead to errors in most cases).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_request_size
func (f *Face) RequestSize(req SizeRequest) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	ptr := (*C.FT_Size_RequestRec)(C.calloc(1, C.sizeof_struct_FT_Size_RequestRec_))
	ptr._type = C.FT_Size_Request_Type(req.Type)
	ptr.width = C.FT_Long(req.Width)
	ptr.height = C.FT_Long(req.Height)
	ptr.horiResolution = C.FT_UInt(req.HoriResolution)
	ptr.vertResolution = C.FT_UInt(req.VertResolution)
	defer free(unsafe.Pointer(ptr))

	// creq := C.FT_Size_RequestRec{
	// 	_type:          C.FT_Size_Request_Type(req.Type),
	// 	width:          C.FT_Long(req.Width),
	// 	height:         C.FT_Long(req.Height),
	// 	horiResolution: C.FT_UInt(req.HoriResolution),
	// 	vertResolution: C.FT_UInt(req.VertResolution),
	// }

	return getErr(C.FT_Request_Size(f.ptr, ptr))
}

// SelectSize selects a bitmap strike.
// To be more precise, this function sets the scaling factors of the active Size object in a face so that bitmaps from
// this particular strike are taken by LoadGlyph and friends.
//
// Don't use this function if you are using the FreeType cache API.
//
// NOTE:
// For bitmaps embedded in outline fonts it is common that only a subset of the available glyphs at a given ppem value
// is available. FreeType silently uses outlines if there is no bitmap for a given glyph index.
//
// For GX and OpenType variation fonts, a bitmap strike makes sense only if the default instance is active (this is, no
// glyph variation takes place); otherwise, FreeType simply ignores bitmap strikes. The same is true for all named
// instances that are different from the default instance.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_select_size
func (f *Face) SelectSize(strikeIndex int) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	return getErr(C.FT_Select_Size(f.ptr, C.FT_Int(strikeIndex)))
}

// SetTransform sets the transformation that is applied to glyph images when they are loaded into a glyph slot through
// LoadGlyph.
//
// NOTE:
// The transformation is only applied to scalable image formats after the glyph has been loaded. It means that hinting
// is unaltered by the transformation and is performed on the character size given in the last call to SetCharSize or
// SetPixelSizes.
//
// Note that this also transforms the face.glyph.advance field, but not the values in face.glyph.metrics.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_set_transform
func (f *Face) SetTransform(matrix Matrix, delta Vector) {
	if f == nil || f.ptr == nil {
		return
	}

	var cmatrix *C.FT_Matrix
	if matrix != (Matrix{}) {
		cmatrix = (*C.FT_Matrix)(C.calloc(1, C.sizeof_struct_FT_Matrix_))
		cmatrix.xx = C.FT_Fixed(matrix.Xx)
		cmatrix.xy = C.FT_Fixed(matrix.Xy)
		cmatrix.yx = C.FT_Fixed(matrix.Yx)
		cmatrix.yy = C.FT_Fixed(matrix.Yy)
		defer free(unsafe.Pointer(cmatrix)) // FT_Set_Transform makes a copy
	}

	var cdelta *C.FT_Vector
	if delta != (Vector{}) {
		cdelta = (*C.FT_Vector)(C.calloc(1, C.sizeof_struct_FT_Vector_))
		cdelta.x = C.FT_Pos(delta.X)
		cdelta.y = C.FT_Pos(delta.Y)
		defer free(unsafe.Pointer(cdelta)) // FT_Set_Transform makes a copy
	}

	C.FT_Set_Transform(f.ptr, cmatrix, cdelta)
}

// LoadGlyph loads a glyph into the glyph slot of the face.
//
// The loaded glyph may be transformed. See SetTransform for the details.
//
// For subsetted CID-keyed fonts, ErrInvalidArgument is returned for invalid CID values (this is, for CID values that
// don't have a corresponding glyph in the font). See the discussion of the FaceFlagCidKeyed flag for more details.
//
// If you receive ErrGlyphTooBig, try getting the glyph outline at EM size, then scale it manually and fill it as a
// graphics operation.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_load_glyph
func (f *Face) LoadGlyph(idx GlyphIndex, flags LoadFlag) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	return getErr(C.FT_Load_Glyph(f.ptr, C.FT_UInt(idx), C.FT_Int32(flags)))
}

// CharIndex returns the glyph index of a given character code. This function
// uses the currently selected charmap to do the mapping.
//
// NOTE:
// If you use FreeType to manipulate the contents of font files directly, be
// aware that the glyph index returned by this function doesn't always
// correspond to the internal indices used within the file. This is done to
// ensure that value 0 always corresponds to the ‘missing glyph’. If the first
// glyph is not named ‘.notdef’, then for Type 1 and Type 42 fonts, ‘.notdef’
// will be moved into the glyph ID 0 position, and whatever was there will be
// moved to the position ‘.notdef’ had. For Type 1 fonts, if there is no
// ‘.notdef’ glyph at all, then one will be created at index 0 and whatever was
// there will be moved to the last index -- Type 42 fonts are considered invalid
// under this condition.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_char_index
func (f *Face) CharIndex(r rune) GlyphIndex {
	if f == nil || f.ptr == nil {
		return 0
	}

	return GlyphIndex(C.FT_Get_Char_Index(f.ptr, C.ulong(r)))
}

// FirstChar returns the first character code in the current charmap of a given
// face, together with its corresponding glyph index.
//
// NOTE:
// You should use this function together with NextChar to parse all character
// codes available in a given charmap. The code should look like this: (TODO)
//	FT_ULong  charcode;
//	FT_UInt   gindex;
//
//
//	charcode = FT_Get_First_Char( face, &gindex );
//	while ( gindex != 0 )
//	{
//	... do something with (charcode,gindex) pair ...
//
//	charcode = FT_Get_Next_Char( face, charcode, &gindex );
//	}
//
// Be aware that character codes can have values up to 0xFFFFFFFF; this might
// happen for non-Unicode or malformed cmaps. However, even with regular Unicode
// encoding, so-called ‘last resort fonts’ (using SFNT cmap format 13, see
// function CharMap.Format) normally have entries for all Unicode characters up
// to 0x1FFFFF, which can cause a lot of iterations.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_first_char
func (f *Face) FirstChar() (rune, GlyphIndex) {
	if f == nil || f.ptr == nil {
		return 0, 0
	}

	var idx C.uint
	r := C.FT_Get_First_Char(f.ptr, &idx)
	return rune(r), GlyphIndex(idx)
}

// NextChar returns the next character code in the current charmap of a given
// face following current, as well as the corresponding glyph index.
//
// GlyphIndex is set to 0 when there are no more codes in the charmap.
//
// NOTE:
// You should use this function with FirstChar to walk over all character codes
// available in a given charmap. See the note for that function for a simple
// code example.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_next_char
func (f *Face) NextChar(current rune) (rune, GlyphIndex) {
	if f == nil || f.ptr == nil {
		return 0, 0
	}

	var idx C.uint
	r := C.FT_Get_Next_Char(f.ptr, C.ulong(current), &idx)
	return rune(r), GlyphIndex(idx)
}

// IndexOf returns the glyph index of a given glyph name.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_name_index
func (f *Face) IndexOf(glyphName string) GlyphIndex {
	if f == nil || f.ptr == nil {
		return 0
	}

	cstr := C.CString(glyphName)
	defer free(unsafe.Pointer(cstr))

	return GlyphIndex(C.FT_Get_Name_Index(f.ptr, cstr))
}

// LoadChar loads a glyph into the glyph slot of a face object, accessed by its
// character code.
// NOTE:
// This function simply calls CharIndex and LoadGlyph.
//
// Many fonts contain glyphs that can't be loaded by this function since its
// glyph indices are not listed in any of the font's charmaps.
//
// If no active cmap is set up (i.e., face.CharMap is zero), the call to
// CharIndex is omitted, and the function behaves identically to LoadGlyph.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_load_char
func (f *Face) LoadChar(r rune, flags LoadFlag) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	return getErr(C.FT_Load_Char(f.ptr, C.ulong(r), C.FT_Int32(flags)))
}

// TODO: FT_Render_Glyph

// Kern returns the kerning vector between two glyphs of the same face.
//
// The return vector is either in font units, fractional pixels (26.6 format),
// or pixels for scalable formats, and in pixels for fixed-sizes formats.
//
// NOTE:
// Only horizontal layouts (left-to-right & right-to-left) are supported by this
// method. Other layouts, or more sophisticated kernings, are out of the scope
// of this API function -- they can be implemented through format-specific
// interfaces.
//
// Kerning for OpenType fonts implemented in a ‘GPOS’ table is not supported;
// use HasFlag(FaceFlagKerning) to find out whether a font has data that can be
// extracted with Kerning().
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_kerning
func (f *Face) Kern(left, right GlyphIndex, mode KerningMode) (Vector, error) {
	if f == nil || f.ptr == nil {
		return Vector{}, ErrInvalidFaceHandle
	}

	var vec C.FT_Vector
	if err := getErr(C.FT_Get_Kerning(f.ptr, C.uint(left), C.uint(right), C.uint(mode), &vec)); err != nil {
		return Vector{}, err
	}

	return Vector{X: Pos(vec.x), Y: Pos(vec.y)}, nil
}

// TODO: FT_Get_Track_Kerning

// GlyphName returns the ASCII name of a given glyph in a face. This only works
// for those faces where face.HasFlag(FaceFlagGlyphNames) is true.
//
// NOTE:
// An error is returned if the face doesn't provide glyph names or if the glyph
// index is invalid.
//
// Be aware that FreeType reorders glyph indices internally so that glyph index
// 0 always corresponds to the ‘missing glyph’ (called ‘.notdef’).
//
// This function always returns an error if the config macro
// FT_CONFIG_OPTION_NO_GLYPH_NAMES is not defined in  ftoption.h.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_glyph_name
func (f *Face) GlyphName(idx GlyphIndex) (string, error) {
	if f == nil || f.ptr == nil {
		return "", ErrInvalidFaceHandle
	}

	// In all cases of failure, the first byte of buffer is set to 0 to indicate an empty name.
	// The glyph name is truncated to fit within the buffer if it is too long. The returned string is always zero-terminated.
	buf := (*C.char)(C.calloc(1024, C.sizeof_char))
	defer free(unsafe.Pointer(buf))

	if err := getErr(C.FT_Get_Glyph_Name(f.ptr, C.uint(idx), (C.FT_Pointer)(buf), 1024)); err != nil {
		return "", err
	}

	return C.GoString(buf), nil
}

// SelectCharMap selects a given charmap by its encoding tag.
// It returns an error if no charmap in the face corresponds to the encoding queried.
//
// Because many fonts contain more than a single cmap for Unicode encoding, this function has some special code to
// select the one that covers Unicode best (‘best’ in the sense that a UCS-4 cmap is preferred to a UCS-2 cmap).
// It is thus preferable to use SetCharMap in this case.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_select_charmap
func (f *Face) SelectCharMap(enc Encoding) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	return getErr(C.FT_Select_Charmap(f.ptr, C.FT_Encoding(enc)))
}

// SetCharMap marks the given charmap as active for character code to glyph index mapping.
//
// It returns an error if the charmap is not part of the face (i.e., if it is not listed in the CharMaps() table).
// It also fails if an OpenType type 14 charmap is selected (which doesn't map character codes to glyph indices at all).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_set_charmap
func (f *Face) SetCharMap(c CharMap) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	if !c.valid {
		return ErrInvalidCharMapHandle
	}

	maps := f.charmaps()
	if c.index < 0 || c.index >= len(maps) {
		return ErrInvalidCharMapHandle
	}

	return getErr(C.FT_Set_Charmap(f.ptr, maps[c.index]))
}

// FSTypeFlags returns the fsType flags for a font.
//
// NOTE:
// Use this function rather than directly reading the FsType field in the
// FontInfo struct, which is only guaranteed to return the correct results for
// Type 1 fonts.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_fstype_flags
func (f *Face) FSTypeFlags() FSTypeFlag {
	if f == nil || f.ptr == nil {
		return 0
	}

	return FSTypeFlag(C.FT_Get_FSType_Flags(f.ptr))
}

// SubGlyphInfo contains info about a subglyph.
// The values of Arg1, Arg2, and Transform must be interpreted depending on the
// flags present in Flags. See the OpenType specification for details.
//
// https://docs.microsoft.com/en-us/typography/opentype/spec/glyf#composite-glyph-description
type SubGlyphInfo struct {
	Index     GlyphIndex
	Flags     SubGlyphFlag
	Arg1      int
	Arg2      int
	Transform Matrix
}

// SubGlyphInfo retrieves a description of a given subglyph. Only use it if
// glyph.Format is GlyphFormatComposite; an error is returned otherwise.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_subglyph_info
func (f *Face) SubGlyphInfo(idx int) (SubGlyphInfo, error) {
	if f == nil || f.ptr == nil {
		return SubGlyphInfo{}, ErrInvalidFaceHandle
	}

	if f.ptr.glyph == nil || C.uint(idx) >= f.ptr.glyph.num_subglyphs {
		return SubGlyphInfo{}, ErrInvalidArgument
	}

	var index C.FT_Int
	var flags C.FT_UInt
	var arg1 C.FT_Int
	var arg2 C.FT_Int
	var transform C.FT_Matrix
	if err := getErr(C.FT_Get_SubGlyph_Info(f.ptr.glyph, C.uint(idx), &index, &flags, &arg1, &arg2, &transform)); err != nil {
		return SubGlyphInfo{}, err
	}

	return SubGlyphInfo{
		Index: GlyphIndex(index),
		Flags: SubGlyphFlag(flags),
		Arg1:  int(arg1),
		Arg2:  int(arg2),
		Transform: Matrix{
			Xx: fixed.Int16_16(transform.xx),
			Xy: fixed.Int16_16(transform.xy),
			Yx: fixed.Int16_16(transform.yx),
			Yy: fixed.Int16_16(transform.yy),
		},
	}, nil
}
