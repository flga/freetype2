package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

// FaceFlags is a list of bit flags of a given face.
// They inform client applications of properties of the corresponding face.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_face_flag_xxx
type FaceFlags int

const (
	// FaceFlagScalable the face contains outline glyphs. Note that a face can contain bitmap strikes also, i.e., a
	// face can have both this flag and FaceFlagFixedSizes set.
	FaceFlagScalable FaceFlags = C.FT_FACE_FLAG_SCALABLE

	// FaceFlagFixedSizes the face contains bitmap strikes. See also the Face.NumFixedSizes() and Face.AvailableSizes().
	FaceFlagFixedSizes FaceFlags = C.FT_FACE_FLAG_FIXED_SIZES

	// FaceFlagFixedWidth the face contains fixed-width characters (like Courier, Lucida, MonoType, etc.).
	FaceFlagFixedWidth FaceFlags = C.FT_FACE_FLAG_FIXED_WIDTH

	// FaceFlagSfnt the face uses the SFNT storage scheme. For now, this means TrueType and OpenType.
	FaceFlagSfnt FaceFlags = C.FT_FACE_FLAG_SFNT

	// FaceFlagHorizontal the face contains horizontal glyph metrics. This should be set for all common formats.
	FaceFlagHorizontal FaceFlags = C.FT_FACE_FLAG_HORIZONTAL

	// FaceFlagVertical the face contains vertical glyph metrics. This is only available in some formats, not all of them.
	FaceFlagVertical FaceFlags = C.FT_FACE_FLAG_VERTICAL

	// FaceFlagKerning the face contains kerning information. If set, the kerning distance can be retrieved using the
	// GetKerning() method.
	// Otherwise the function always returnS the vector (0,0). Note that FreeType doesn't handle kerning data from the
	// SFNT ‘GPOS’ table (as present in many OpenType fonts).
	FaceFlagKerning FaceFlags = C.FT_FACE_FLAG_KERNING

	// FaceFlagMultipleMasters the face contains multiple masters and is capable of interpolating between them.
	// Supported formats are Adobe MM, TrueType GX, and OpenType variation fonts.
	//
	// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html
	FaceFlagMultipleMasters FaceFlags = C.FT_FACE_FLAG_MULTIPLE_MASTERS

	// FaceFlagGlyphNames the face contains glyph names, which can be retrieved using GetGlyphName().
	// Note that some TrueType fonts contain broken glyph name tables. Use HasPSGlyphNames() when needed.
	FaceFlagGlyphNames FaceFlags = C.FT_FACE_FLAG_GLYPH_NAMES

	// FaceFlagHinter the font driver has a hinting machine of its own. For example, with TrueType fonts, it makes sense
	// to use data from the SFNT ‘gasp’ table only if the native TrueType hinting engine (with the bytecode interpreter)
	// is available and active.
	FaceFlagHinter FaceFlags = C.FT_FACE_FLAG_HINTER

	// FaceFlagCidKeyed The face is CID-keyed. In that case, the face is not accessed by glyph indices but by CID values.
	// For subsetted CID-keyed fonts this has the consequence that not all index values are a valid argument to
	// LoadGlyph(). Only the CID values for which corresponding glyphs in the subsetted font exist make LoadGlyph()
	// return successfully; in all other cases you get an ErrInvalidArgument error.
	//
	// Note that CID-keyed fonts that are in an SFNT wrapper (this is, all OpenType/CFF fonts) don't have this flag set
	// since the glyphs are accessed in the normal way (using contiguous indices); the ‘CID-ness’ isn't visible to the
	// application.
	FaceFlagCidKeyed FaceFlags = C.FT_FACE_FLAG_CID_KEYED

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
	FaceFlagTricky FaceFlags = C.FT_FACE_FLAG_TRICKY

	// FaceFlagColor the face has color glyph tables. See LoadColor for more information.
	// [Since 2.5.1].
	FaceFlagColor FaceFlags = C.FT_FACE_FLAG_COLOR

	// FaceFlagVariation [Since 2.9] set if the current face (or named instance) has been altered with
	// SetMMDesignCoordinates, SetVarDesignCoordinates, or SetVarBlend_Coordinates.
	// This flag is unset by a call to SetNamedInstance.
	FaceFlagVariation FaceFlags = C.FT_FACE_FLAG_VARIATION
)

func (x FaceFlags) String() string {
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

// StyleFlags is a list of bit flags to indicate the style of a given face.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_style_flag_xxx
type StyleFlags int

const (
	// StyleFlagItalic the face style is italic or oblique
	StyleFlagItalic StyleFlags = C.FT_STYLE_FLAG_ITALIC
	// StyleFlagBold the face is bold
	StyleFlagBold StyleFlags = C.FT_STYLE_FLAG_BOLD
)

func (x StyleFlags) String() string {
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

// FaceIndex is the index of a face in a given font file. It holds two different values.
// Bits 0-15 are the index of the face in the font file (starting with value 0). Set it to 0 if there is only one face
// in the font file.
//
// [Since 2.6.1] Bits 16-30 are relevant to GX and OpenType variation fonts only, specifying the named instance index
// for the current face index (starting with value 1; value 0 makes FreeType ignore named instances).
// For non-variation fonts, bits 16-30 are ignored. Assuming that you want to access the third named instance in face 4,
// the value should be set to 0x00030004. If you want to access face 4 without variation handling, simply set it to 4.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_open_face (face_index argument)
type FaceIndex int

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

// Index returns the index of the given face.
func (f *Face) Index() FaceIndex {
	if f == nil || f.ptr == nil {
		return 0
	}
	return FaceIndex(f.ptr.face_index)
}

// FaceFlags returns the set of bit flags that give important information about the face.
func (f *Face) FaceFlags() FaceFlags {
	if f == nil || f.ptr == nil {
		return 0
	}
	return FaceFlags(f.ptr.face_flags)
}

// HasFlag reports whether the face has the given flag.
func (f *Face) HasFlag(flag FaceFlags) bool { return f.FaceFlags()&flag > 0 }

// StyleFlags returns the set of bit flags indicating the style of the face; see StyleFlags for details.
//
// [Since 2.6.1] Bits 16-30 hold the number of named instances available for the current face if we have a GX or
// OpenType variation (sub)font. Bit 31 is always zero (this is, it is always a positive value). Note that a variation
// font has always at least one named instance, namely the default instance.
func (f *Face) StyleFlags() StyleFlags {
	if f == nil || f.ptr == nil {
		return 0
	}
	return StyleFlags(f.ptr.style_flags)
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
