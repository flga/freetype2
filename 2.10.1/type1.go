package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_TYPE1_TABLES_H
import (
	"C"
)

import (
	"github.com/flga/freetype2/fixed"
)

// PSFontInfo models a Type 1 or Type 2 FontInfo dictionary.
// Note that for Multiple Master fonts, each instance has its own FontInfo dictionary.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#ps_fontinforec
type PSFontInfo struct {
	Version            string
	Notice             string
	FullName           string
	FamilyName         string
	Weight             string
	ItalicAngle        int
	IsFixedPitch       bool
	UnderlinePosition  int
	UnderlineThickness int
}

// PSPrivate models a Type 1 or Type 2 private dictionary.
// Note that for Multiple Master fonts, each instance has its own Private dictionary.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#ps_privaterec
type PSPrivate struct {
	UniqueID int
	LenIV    int

	NumBlueValues       byte
	NumOtherBlues       byte
	NumFamilyBlues      byte
	NumFamilyOtherBlues byte

	BlueValues [14]int16
	OtherBlues [10]int16

	FamilyBlues      [14]int16
	FamilyOtherBlues [10]int16

	BlueScale fixed.Int16_16
	BlueShift int
	BlueFuzz  int

	StandardWidth  uint16
	StandardHeight uint16

	NumSnapWidths  byte
	NumSnapHeights byte
	ForceBold      bool
	RoundStemUp    bool

	SnapWidths  [13]int16 /* including std width  */
	SnapHeights [13]int16 /* including std height */

	ExpansionFactor fixed.Int16_16

	LanguageGroup int
	Password      int

	MinFeature [2]int16
}

// CIDFaceDict represents data in a CID top-level dictionary. In most cases,
// they are part of the font's ‘/FDArray’ array. Within a CID font file, such
// (internal) subfont dictionaries are enclosed by ‘%ADOBeginFontDict’ and
// ‘%ADOEndFontDict’ comments.
//
// Note that CID_FaceDictRec misses a field for the ‘/FontName’ keyword,
// specifying the subfont's name (the top-level font name is given by the
// ‘/CIDFontName’ keyword). This is an oversight, but it doesn't limit the ‘cid’
// font module's functionality because FreeType neither needs this entry nor
// gives access to CID subfonts.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#cid_facedictrec
type CIDFaceDict struct {
	PrivateDict PSPrivate

	LenBuildchar       uint
	ForceboldThreshold fixed.Int16_16
	StrokeWidth        Pos
	ExpansionFactor    fixed.Int16_16
	PaintType          byte
	FontType           byte
	FontMatrix         Matrix
	FontOffset         Vector

	NumSubrs      uint
	SubrmapOffset uint
	SdBytes       int
}

// CIDFaceInfo represents CID Face information.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#cid_faceinforec
type CIDFaceInfo struct {
	CidFontName string
	CidVersion  fixed.Int16_16
	CidFontType int

	Registry   string
	Ordering   string
	Supplement int

	FontInfo PSFontInfo
	FontBbox BBox
	UIDBase  uint

	NumXUID int
	XUID    [16]uint

	CidmapOffset uint
	FdBytes      int
	GdBytes      int
	CidCount     uint

	NumDicts  int
	FontDicts CIDFaceDict

	DataOffset uint
}

// HasPSGlyphNames reports whether a given face provides reliable PostScript
// glyph names. This is similar to using the FT_HAS_GLYPH_NAMES macro, except
// that certain fonts (mostly TrueType) contain incorrect glyph name tables.
//
// When this function returns true, the caller is sure that the glyph names
// returned by GetGlyphName are reliable.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#ft_has_ps_glyph_names
func (f *Face) HasPSGlyphNames() bool {
	if f == nil || f.ptr == nil {
		return false
	}

	v := C.FT_Has_PS_Glyph_Names(f.ptr)
	return v == 1
}

// PSFontInfo returns the PSFontInfo corresponding to a given PostScript font.
//
// If the font's format is not PostScript-based, it will return ErrInvalidArgument.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#ft_get_ps_font_info
func (f *Face) PSFontInfo() (PSFontInfo, error) {
	if f == nil || f.ptr == nil {
		return PSFontInfo{}, ErrInvalidFaceHandle
	}

	var info C.PS_FontInfoRec
	if err := getErr(C.FT_Get_PS_Font_Info(f.ptr, &info)); err != nil {
		return PSFontInfo{}, err
	}

	ret := PSFontInfo{
		Version:            C.GoString(info.version),
		Notice:             C.GoString(info.notice),
		FullName:           C.GoString(info.full_name),
		FamilyName:         C.GoString(info.family_name),
		Weight:             C.GoString(info.weight),
		ItalicAngle:        int(info.italic_angle),
		UnderlinePosition:  int(info.underline_position),
		UnderlineThickness: int(info.underline_thickness),
	}
	if info.is_fixed_pitch == 1 {
		ret.IsFixedPitch = true
	}

	return ret, nil
}

// PSPrivate returns the PSPrivate corresponding to a given PostScript font.
//
// If the font's format is not PostScript-based, it will return ErrInvalidArgument.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#ft_get_ps_font_private
func (f *Face) PSPrivate() (PSPrivate, error) {
	if f == nil || f.ptr == nil {
		return PSPrivate{}, ErrInvalidFaceHandle
	}

	var info C.PS_PrivateRec
	if err := getErr(C.FT_Get_PS_Font_Private(f.ptr, &info)); err != nil {
		return PSPrivate{}, err
	}

	ret := PSPrivate{
		UniqueID:            int(info.unique_id),
		LenIV:               int(info.lenIV),
		NumBlueValues:       byte(info.num_blue_values),
		NumOtherBlues:       byte(info.num_other_blues),
		NumFamilyBlues:      byte(info.num_family_blues),
		NumFamilyOtherBlues: byte(info.num_family_other_blues),
		BlueScale:           fixed.Int16_16(info.blue_scale),
		BlueShift:           int(info.blue_shift),
		BlueFuzz:            int(info.blue_fuzz),
		StandardWidth:       uint16(info.standard_width[0]),
		StandardHeight:      uint16(info.standard_height[0]),
		NumSnapWidths:       byte(info.num_snap_widths),
		NumSnapHeights:      byte(info.num_snap_heights),
		ExpansionFactor:     fixed.Int16_16(info.expansion_factor),
		LanguageGroup:       int(info.language_group),
		Password:            int(info.password),
	}
	for i, v := range info.blue_values {
		ret.BlueValues[i] = int16(v)
	}
	for i, v := range info.other_blues {
		ret.OtherBlues[i] = int16(v)
	}
	for i, v := range info.family_blues {
		ret.FamilyBlues[i] = int16(v)
	}
	for i, v := range info.family_other_blues {
		ret.FamilyOtherBlues[i] = int16(v)
	}
	for i, v := range info.snap_widths {
		ret.SnapWidths[i] = int16(v)
	}
	for i, v := range info.snap_heights {
		ret.SnapHeights[i] = int16(v)
	}
	for i, v := range info.min_feature {
		ret.MinFeature[i] = int16(v)
	}
	if info.force_bold == 1 {
		ret.ForceBold = true
	}
	if info.round_stem_up == 1 {
		ret.RoundStemUp = true
	}

	return ret, nil
}

// T1BlendFlags is a set of flags used to indicate which fields are present in a
// given blend dictionary (font info or private). Used to support Multiple
// Masters fonts.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#t1_blend_flags
type T1BlendFlags uint

// required fields in a FontInfo blend dictionary
const (
	T1BlendUnderlinePosition  T1BlendFlags = C.T1_BLEND_UNDERLINE_POSITION
	T1BlendUnderlineThickness T1BlendFlags = C.T1_BLEND_UNDERLINE_THICKNESS
	T1BlendItalicAngle        T1BlendFlags = C.T1_BLEND_ITALIC_ANGLE
)

// required fields in a Private blend dictionary
const (
	T1BlendBlueValues       T1BlendFlags = C.T1_BLEND_BLUE_VALUES
	T1BlendOtherBlues       T1BlendFlags = C.T1_BLEND_OTHER_BLUES
	T1BlendStandardWidth    T1BlendFlags = C.T1_BLEND_STANDARD_WIDTH
	T1BlendStandardHeight   T1BlendFlags = C.T1_BLEND_STANDARD_HEIGHT
	T1BlendStemSnapWidths   T1BlendFlags = C.T1_BLEND_STEM_SNAP_WIDTHS
	T1BlendStemSnapHeights  T1BlendFlags = C.T1_BLEND_STEM_SNAP_HEIGHTS
	T1BlendBlueScale        T1BlendFlags = C.T1_BLEND_BLUE_SCALE
	T1BlendBlueShift        T1BlendFlags = C.T1_BLEND_BLUE_SHIFT
	T1BlendFamilyBlues      T1BlendFlags = C.T1_BLEND_FAMILY_BLUES
	T1BlendFamilyOtherBlues T1BlendFlags = C.T1_BLEND_FAMILY_OTHER_BLUES
	T1BlendForceBold        T1BlendFlags = C.T1_BLEND_FORCE_BOLD
)

// T1EncodingType is an enumeration describing the ‘Encoding’ entry in a Type 1
// dictionary.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#t1_encodingtype
type T1EncodingType uint

// T1EncodingType values
const (
	T1EncodingTypeNone      T1EncodingType = C.T1_ENCODING_TYPE_NONE
	T1EncodingTypeArray     T1EncodingType = C.T1_ENCODING_TYPE_ARRAY
	T1EncodingTypeStandard  T1EncodingType = C.T1_ENCODING_TYPE_STANDARD
	T1EncodingTypeISOLatin1 T1EncodingType = C.T1_ENCODING_TYPE_ISOLATIN1
	T1EncodingTypeExpert    T1EncodingType = C.T1_ENCODING_TYPE_EXPERT
)

// PSDictKey is an enumeration used in calls to GetPSFontValue to identify the
// Type 1 dictionary entry to retrieve.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-type1_tables.html#ps_dict_keys
type PSDictKey uint

// conventionally in the font dictionary
const (
	PSDictFontType       PSDictKey = C.PS_DICT_FONT_TYPE
	PSDictFontMatrix     PSDictKey = C.PS_DICT_FONT_MATRIX
	PSDictFontBbox       PSDictKey = C.PS_DICT_FONT_BBOX
	PSDictPaintType      PSDictKey = C.PS_DICT_PAINT_TYPE
	PSDictFontName       PSDictKey = C.PS_DICT_FONT_NAME
	PSDictUniqueID       PSDictKey = C.PS_DICT_UNIQUE_ID
	PSDictNumCharStrings PSDictKey = C.PS_DICT_NUM_CHAR_STRINGS
	PSDictCharStringKey  PSDictKey = C.PS_DICT_CHAR_STRING_KEY
	PSDictCharString     PSDictKey = C.PS_DICT_CHAR_STRING
	PSDictEncodingType   PSDictKey = C.PS_DICT_ENCODING_TYPE
	PSDictEncodingEntry  PSDictKey = C.PS_DICT_ENCODING_ENTRY
)

// conventionally in the font Private dictionary
const (
	PSDictNumSubrs            PSDictKey = C.PS_DICT_NUM_SUBRS
	PSDictSubr                PSDictKey = C.PS_DICT_SUBR
	PSDictStdHw               PSDictKey = C.PS_DICT_STD_HW
	PSDictStdVw               PSDictKey = C.PS_DICT_STD_VW
	PSDictNumBlueValues       PSDictKey = C.PS_DICT_NUM_BLUE_VALUES
	PSDictBlueValue           PSDictKey = C.PS_DICT_BLUE_VALUE
	PSDictBlueFuzz            PSDictKey = C.PS_DICT_BLUE_FUZZ
	PSDictNumOtherBlues       PSDictKey = C.PS_DICT_NUM_OTHER_BLUES
	PSDictOtherBlue           PSDictKey = C.PS_DICT_OTHER_BLUE
	PSDictNumFamilyBlues      PSDictKey = C.PS_DICT_NUM_FAMILY_BLUES
	PSDictFamilyBlue          PSDictKey = C.PS_DICT_FAMILY_BLUE
	PSDictNumFamilyOtherBlues PSDictKey = C.PS_DICT_NUM_FAMILY_OTHER_BLUES
	PSDictFamilyOtherBlue     PSDictKey = C.PS_DICT_FAMILY_OTHER_BLUE
	PSDictBlueScale           PSDictKey = C.PS_DICT_BLUE_SCALE
	PSDictBlueShift           PSDictKey = C.PS_DICT_BLUE_SHIFT
	PSDictNumStemSnapH        PSDictKey = C.PS_DICT_NUM_STEM_SNAP_H
	PSDictStemSnapH           PSDictKey = C.PS_DICT_STEM_SNAP_H
	PSDictNumStemSnapV        PSDictKey = C.PS_DICT_NUM_STEM_SNAP_V
	PSDictStemSnapV           PSDictKey = C.PS_DICT_STEM_SNAP_V
	PSDictForceBold           PSDictKey = C.PS_DICT_FORCE_BOLD
	PSDictRndStemUp           PSDictKey = C.PS_DICT_RND_STEM_UP
	PSDictMinFeature          PSDictKey = C.PS_DICT_MIN_FEATURE
	PSDictLenIv               PSDictKey = C.PS_DICT_LEN_IV
	PSDictPassword            PSDictKey = C.PS_DICT_PASSWORD
	PSDictLanguageGroup       PSDictKey = C.PS_DICT_LANGUAGE_GROUP
)

// conventionally in the font FontInfo dictionary
const (
	PSDictVersion            PSDictKey = C.PS_DICT_VERSION
	PSDictNotice             PSDictKey = C.PS_DICT_NOTICE
	PSDictFullName           PSDictKey = C.PS_DICT_FULL_NAME
	PSDictFamilyName         PSDictKey = C.PS_DICT_FAMILY_NAME
	PSDictWeight             PSDictKey = C.PS_DICT_WEIGHT
	PSDictIsFixedPitch       PSDictKey = C.PS_DICT_IS_FIXED_PITCH
	PSDictUnderlinePosition  PSDictKey = C.PS_DICT_UNDERLINE_POSITION
	PSDictUnderlineThickness PSDictKey = C.PS_DICT_UNDERLINE_THICKNESS
	PSDictFsType             PSDictKey = C.PS_DICT_FS_TYPE
	PSDictItalicAngle        PSDictKey = C.PS_DICT_ITALIC_ANGLE
)
