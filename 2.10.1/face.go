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

// LoadFlag is a list of bit field constants for LoadGlyph to indicate what
// kind of operations to perform during glyph loading.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_load_xxx
type LoadFlag int32

const (
	// LoadDefault is used as the default glyph load operation. In this case, the following happens:
	//
	// 1 - FreeType looks for a bitmap for the glyph corresponding to the face's current size. If one is found, the function
	// returns. The bitmap data can be accessed from the glyph slot (see note below).
	// 2 - If no embedded bitmap is searched for or found, FreeType looks for a scalable outline. If one is found, it is
	// loaded from the font file, scaled to device pixels, then ‘hinted’ to the pixel grid in order to optimize it. The
	// outline data can be accessed from the glyph slot (see note below).
	//
	// Note that by default the glyph loader doesn't render outlines into bitmaps. The following flags are used to modify
	// this default behaviour to more specific and useful cases.
	LoadDefault LoadFlag = C.FT_LOAD_DEFAULT
	// LoadNoScale don't scale the loaded outline glyph but keep it in font units.
	//
	// This flag implies LoadNoHinting and LoadNoBitmap, and unsets LoadRender.
	//
	// If the font is ‘tricky’ (see FaceFlagTricky for more), using LoadNoScale usually yields meaningless outlines
	// because the subglyphs must be scaled and positioned with hinting instructions. This can be solved by loading the
	// font without LoadNoScale and setting the character size to face.UnitsPerEM().
	LoadNoScale LoadFlag = C.FT_LOAD_NO_SCALE
	// LoadNoHinting disables hinting. This generally generates ‘blurrier’ bitmap glyphs when the glyph are rendered
	// in any of the anti-aliased modes. See also the note below.
	//
	// This flag is implied by LoadNoScale.
	LoadNoHinting LoadFlag = C.FT_LOAD_NO_HINTING
	// LoadRender call RenderGlyph after the glyph is loaded. By default, the glyph is rendered in RenderModeNormal mode.
	// This can be overridden by any LoadTarget or LoadMonochrome.
	//
	// This flag is unset by LoadNoScale.
	LoadRender LoadFlag = C.FT_LOAD_RENDER
	// LoadNoBitmap ignores bitmap strikes when loading. Bitmap-only fonts ignore this flag.
	//
	// LoadNoScale always sets this flag.
	LoadNoBitmap LoadFlag = C.FT_LOAD_NO_BITMAP
	// LoadVerticalLayout load the glyph for vertical text layout. In particular, the advance value in the GlyphSlot is
	// set to the VertAdvance value of the metrics field.
	//
	// If the face does not have FaceFlagVertical, you shouldn't use this flag currently. Reason is that in this case
	// vertical metrics get synthesized, and those values are not always consistent across various font formats.
	LoadVerticalLayout LoadFlag = C.FT_LOAD_VERTICAL_LAYOUT
	// LoadForceAutohint prefer the auto-hinter over the font's native hinter. See also the note below.
	LoadForceAutohint LoadFlag = C.FT_LOAD_FORCE_AUTOHINT
	// LoadPedantic makes the font driver perform pedantic verifications during glyph loading and hinting. This is
	// mostly used to detect broken glyphs in fonts. By default, FreeType tries to handle broken fonts also.
	//
	// In particular, errors from the TrueType bytecode engine are not passed to the application if this flag is not set;
	// this might result in partially hinted or distorted glyphs in case a glyph's bytecode is buggy.
	LoadPedantic LoadFlag = C.FT_LOAD_PEDANTIC
	// LoadNoRecurse don't load composite glyphs recursively. Instead, the font driver fills the NumSubglyph and
	// Subglyphs values of the glyph slot; it also sets glyph.Format to GlyphFormatComposite. The description of
	// subglyphs can then be accessed with GetSubGlyphInfo.
	//
	// Don't use this flag for retrieving metrics information since some font drivers only return rudimentary data.
	//
	// This flag implies LoadNoScale and LoadIgnoreTransform.
	LoadNoRecurse LoadFlag = C.FT_LOAD_NO_RECURSE
	// LoadIgnoreTransform ignore the transform matrix set by SetTransform.
	LoadIgnoreTransform LoadFlag = C.FT_LOAD_IGNORE_TRANSFORM
	// LoadMonochrome is used with LoadRender to indicate that you want to render an outline glyph to a 1-bit monochrome
	// bitmap glyph, with 8 pixels packed into each byte of the bitmap data.
	//
	// Note that this has no effect on the hinting algorithm used. You should rather use LoadTargetMono so that the
	// monochrome-optimized hinting algorithm is used.
	LoadMonochrome LoadFlag = C.FT_LOAD_MONOCHROME
	// LoadLinearDesign keep LinearHoriAdvance and LinearVertAdvance fields of GlyphSlot in font units. See GlyphSlot
	// for details.
	LoadLinearDesign LoadFlag = C.FT_LOAD_LINEAR_DESIGN
	// LoadNoAutohint disables the auto-hinter. See also the note below.
	LoadNoAutohint LoadFlag = C.FT_LOAD_NO_AUTOHINT
	// LoadColor loads colored glyphs. There are slight differences depending on the font format.
	//
	// [Since 2.5] Load embedded color bitmap images. The resulting color bitmaps, if available, will have the
	// PixelModeBGRA format, with pre-multiplied color channels. If the flag is not set and color bitmaps are found,
	// they are converted to 256-level gray bitmaps, using the PixelModeGray format.
	//
	// [Since 2.10, experimental] If the glyph index contains an entry in the face's ‘COLR’ table with a ‘CPAL’ palette
	// table (as defined in the OpenType specification), make RenderGlyph provide a default blending of the color glyph
	// layers associated with the glyph index, using the same bitmap format as embedded color bitmap images. This is
	// mainly for convenience; for full control of color layers use GetColorGlyphLayer and FreeType's color functions
	// like PaletteSelect instead of setting LoadColor for rendering so that the client application can handle blending
	// by itself.
	LoadColor LoadFlag = C.FT_LOAD_COLOR
	// LoadComputeMetrics [Since 2.6.1] Compute glyph metrics from the glyph data, without the use of bundled metrics
	// tables (for example, the ‘hdmx’ table in TrueType fonts). This flag is mainly used by font validating or font
	// editing applications, which need to ignore, verify, or edit those tables.
	//
	// Currently, this flag is only implemented for TrueType fonts.
	LoadComputeMetrics LoadFlag = C.FT_LOAD_COMPUTE_METRICS
	// LoadBitmapMetricsOnly [Since 2.7.1] request loading of the metrics and bitmap image information of a (possibly
	// embedded) bitmap glyph without allocating or copying the bitmap image data itself. No effect if the target glyph
	// is not a bitmap image.
	//
	// This flag unsets LoadRender.
	LoadBitmapMetricsOnly LoadFlag = C.FT_LOAD_BITMAP_METRICS_ONLY
)

func (x LoadFlag) String() string {
	// the maximum concatenated len, at the time of writing, is 180.
	s := make([]byte, 0, 180)

	if x == LoadDefault {
		return "Default"
	}

	if x&LoadNoScale > 0 {
		s = append(s, []byte("NoScale|")...)
	}
	if x&LoadNoHinting > 0 {
		s = append(s, []byte("NoHinting|")...)
	}
	if x&LoadRender > 0 {
		s = append(s, []byte("Render|")...)
	}
	if x&LoadNoBitmap > 0 {
		s = append(s, []byte("NoBitmap|")...)
	}
	if x&LoadVerticalLayout > 0 {
		s = append(s, []byte("VerticalLayout|")...)
	}
	if x&LoadForceAutohint > 0 {
		s = append(s, []byte("ForceAutohint|")...)
	}
	if x&LoadPedantic > 0 {
		s = append(s, []byte("Pedantic|")...)
	}
	if x&LoadNoRecurse > 0 {
		s = append(s, []byte("NoRecurse|")...)
	}
	if x&LoadIgnoreTransform > 0 {
		s = append(s, []byte("IgnoreTransform|")...)
	}
	if x&LoadMonochrome > 0 {
		s = append(s, []byte("Monochrome|")...)
	}
	if x&LoadLinearDesign > 0 {
		s = append(s, []byte("LinearDesign|")...)
	}
	if x&LoadNoAutohint > 0 {
		s = append(s, []byte("NoAutohint|")...)
	}
	if x&LoadColor > 0 {
		s = append(s, []byte("Color|")...)
	}
	if x&LoadComputeMetrics > 0 {
		s = append(s, []byte("ComputeMetrics|")...)
	}
	if x&LoadBitmapMetricsOnly > 0 {
		s = append(s, []byte("BitmapMetricsOnly|")...)
	}
	if len(s) == 0 {
		return ""
	}

	return string(s[:len(s)-1]) // trim the leading |
}

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

func (f *Face) charmaps() []C.FT_CharMap {
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

	charmaps := f.charmaps()
	if c.index < 0 || c.index >= len(charmaps) {
		return ErrInvalidCharMapHandle
	}

	var charmap C.FT_CharMap
	for i, cmap := range charmaps {
		if i == c.index {
			charmap = cmap
			break
		}
	}

	return getErr(C.FT_Set_Charmap(f.ptr, charmap))
}
