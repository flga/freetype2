package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_TRUETYPE_TABLES_H
import "C"
import (
	"unsafe"
)

// FaceFlag is a list of bit flags of a given face.
// They inform client applications of properties of the corresponding face.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_face_flag_xxx
type FaceFlag int

const (
	// FaceFlagScalable the face contains outline glyphs. Note that a face can contain bitmap strikes also, i.e., a
	// face can have both this flag and FaceFlagFixedSizes set.
	FaceFlagScalable FaceFlag = C.FT_FACE_FLAG_SCALABLE

	// FaceFlagFixedSizes the face contains bitmap strikes. See also the Face.NumFixedSizes() and Face.AvailableSizes().
	FaceFlagFixedSizes FaceFlag = C.FT_FACE_FLAG_FIXED_SIZES

	// FaceFlagFixedWidth the face contains fixed-width characters (like Courier, Lucida, MonoType, etc.).
	FaceFlagFixedWidth FaceFlag = C.FT_FACE_FLAG_FIXED_WIDTH

	// FaceFlagSfnt the face uses the SFNT storage scheme. For now, this means TrueType and OpenType.
	FaceFlagSfnt FaceFlag = C.FT_FACE_FLAG_SFNT

	// FaceFlagHorizontal the face contains horizontal glyph metrics. This should be set for all common formats.
	FaceFlagHorizontal FaceFlag = C.FT_FACE_FLAG_HORIZONTAL

	// FaceFlagVertical the face contains vertical glyph metrics. This is only available in some formats, not all of them.
	FaceFlagVertical FaceFlag = C.FT_FACE_FLAG_VERTICAL

	// FaceFlagKerning the face contains kerning information. If set, the kerning distance can be retrieved using the
	// GetKerning() method.
	// Otherwise the function always returnS the vector (0,0). Note that FreeType doesn't handle kerning data from the
	// SFNT ‘GPOS’ table (as present in many OpenType fonts).
	FaceFlagKerning FaceFlag = C.FT_FACE_FLAG_KERNING

	// FaceFlagMultipleMasters the face contains multiple masters and is capable of interpolating between them.
	// Supported formats are Adobe MM, TrueType GX, and OpenType variation fonts.
	//
	// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html
	FaceFlagMultipleMasters FaceFlag = C.FT_FACE_FLAG_MULTIPLE_MASTERS

	// FaceFlagGlyphNames the face contains glyph names, which can be retrieved using GetGlyphName().
	// Note that some TrueType fonts contain broken glyph name tables. Use HasPSGlyphNames() when needed.
	FaceFlagGlyphNames FaceFlag = C.FT_FACE_FLAG_GLYPH_NAMES

	// FaceFlagHinter the font driver has a hinting machine of its own. For example, with TrueType fonts, it makes sense
	// to use data from the SFNT ‘gasp’ table only if the native TrueType hinting engine (with the bytecode interpreter)
	// is available and active.
	FaceFlagHinter FaceFlag = C.FT_FACE_FLAG_HINTER

	// FaceFlagCidKeyed The face is CID-keyed. In that case, the face is not accessed by glyph indices but by CID values.
	// For subsetted CID-keyed fonts this has the consequence that not all index values are a valid argument to
	// LoadGlyph(). Only the CID values for which corresponding glyphs in the subsetted font exist make LoadGlyph()
	// return successfully; in all other cases you get an ErrInvalidArgument error.
	//
	// Note that CID-keyed fonts that are in an SFNT wrapper (this is, all OpenType/CFF fonts) don't have this flag set
	// since the glyphs are accessed in the normal way (using contiguous indices); the ‘CID-ness’ isn't visible to the
	// application.
	FaceFlagCidKeyed FaceFlag = C.FT_FACE_FLAG_CID_KEYED

	// FaceFlagTricky the face is ‘tricky’, this is, it always needs the font format's native hinting engine to get a
	// reasonable result.
	// A typical example is the old Chinese font mingli.ttf (but not mingliu.ttc) that uses TrueType bytecode
	// instructions to move and scale all of its subglyphs.
	//
	// It is not possible to auto-hint such fonts using LoadForceAutohint; it will also ignore LoadNoHinting. You have
	// to set both LoadNoHinting and LoadNoAutohint to really disable hinting; however, you probably never want this
	// except for demonstration purposes.
	//
	// Currently, there are about a dozen TrueType fonts in the list of tricky fonts; they are hard-coded in file ttobjs.c.
	FaceFlagTricky FaceFlag = C.FT_FACE_FLAG_TRICKY

	// FaceFlagColor the face has color glyph tables. See LoadColor for more information.
	// [Since 2.5.1].
	FaceFlagColor FaceFlag = C.FT_FACE_FLAG_COLOR

	// FaceFlagVariation [Since 2.9] set if the current face (or named instance) has been altered with
	// SetMMDesignCoordinates, SetVarDesignCoordinates, or SetVarBlend_Coordinates.
	// This flag is unset by a call to SetNamedInstance.
	FaceFlagVariation FaceFlag = C.FT_FACE_FLAG_VARIATION
)

func (x FaceFlag) String() string {
	s := make([]byte, 0, 130) // len = sum of all the strings below.

	if x&FaceFlagScalable > 0 {
		s = append(s, []byte("Scalable|")...)
	}
	if x&FaceFlagFixedSizes > 0 {
		s = append(s, []byte("FixedSizes|")...)
	}
	if x&FaceFlagFixedWidth > 0 {
		s = append(s, []byte("FixedWidth|")...)
	}
	if x&FaceFlagSfnt > 0 {
		s = append(s, []byte("Sfnt|")...)
	}
	if x&FaceFlagHorizontal > 0 {
		s = append(s, []byte("Horizontal|")...)
	}
	if x&FaceFlagVertical > 0 {
		s = append(s, []byte("Vertical|")...)
	}
	if x&FaceFlagKerning > 0 {
		s = append(s, []byte("Kerning|")...)
	}
	if x&FaceFlagMultipleMasters > 0 {
		s = append(s, []byte("MultipleMasters|")...)
	}
	if x&FaceFlagGlyphNames > 0 {
		s = append(s, []byte("GlyphNames|")...)
	}
	if x&FaceFlagHinter > 0 {
		s = append(s, []byte("Hinter|")...)
	}
	if x&FaceFlagCidKeyed > 0 {
		s = append(s, []byte("CidKeyed|")...)
	}
	if x&FaceFlagTricky > 0 {
		s = append(s, []byte("Tricky|")...)
	}
	if x&FaceFlagColor > 0 {
		s = append(s, []byte("Color|")...)
	}
	if x&FaceFlagVariation > 0 {
		s = append(s, []byte("Variation|")...)
	}

	if len(s) == 0 {
		return ""
	}
	return string(s[:len(s)-1]) // trim the leading |
}

// StyleFlag is a list of bit flags to indicate the style of a given face.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_style_flag_xxx
type StyleFlag int

const (
	// StyleFlagItalic the face style is italic or oblique
	StyleFlagItalic StyleFlag = C.FT_STYLE_FLAG_ITALIC
	// StyleFlagBold the face is bold
	StyleFlagBold StyleFlag = C.FT_STYLE_FLAG_BOLD
)

func (x StyleFlag) String() string {
	s := make([]byte, 0, 12) // len = sum of all the strings below.

	if x&StyleFlagItalic > 0 {
		s = append(s, []byte("Italic|")...)
	}
	if x&StyleFlagBold > 0 {
		s = append(s, []byte("Bold|")...)
	}

	if len(s) == 0 {
		return ""
	}
	return string(s[:len(s)-1]) // trim the leading |
}

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
		ret[i] = newBitmapSize(ptr[i])
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
	ptr := (*[(1<<31 - 1) / C.sizeof_FT_CharMap]C.FT_CharMap)(unsafe.Pointer(f.ptr.charmaps))[:n:n]
	for i := range ret {
		ret[i] = newCharMap(ptr[i])
	}
	return ret
}

func (f *Face) testCCharMaps() []C.FT_CharMap {
	if f == nil || f.ptr == nil {
		return nil
	}

	n := int(f.ptr.num_charmaps)
	if n == 0 {
		return nil
	}

	return (*[(1<<31 - 1) / C.sizeof_FT_CharMap]C.FT_CharMap)(unsafe.Pointer(f.ptr.charmaps))[:n:n]
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

	return newBBox(f.ptr.bbox)
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

// TODO: GLYPH

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

func (f *Face) testClearCharmap() {
	if f == nil || f.ptr == nil {
		return
	}
	f.ptr.charmap = nil
}
