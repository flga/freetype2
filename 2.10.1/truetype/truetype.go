// Package truetype provides definitions of some basic tables specific to TrueType and OpenType
package truetype

// #cgo windows LDFLAGS: -lfreetype2
// #cgo !static,!windows pkg-config: freetype2
//
// #cgo linux,386,static CFLAGS: -I${SRCDIR}/../linux_386/include/freetype2 -Werror -Wall -Wextra -Wno-unused-parameter
// #cgo linux,386,static,!harfbuzz LDFLAGS: -L${SRCDIR}/../linux_386/lib -lfreetype -lbz2 -lpng16 -lz -lm
// #cgo linux,386,static,harfbuzz LDFLAGS: -L${SRCDIR}/../linux_386/lib -lfreetypehb -lharfbuzz -lfreetypehb -lbz2 -lpng16 -lz -lm
// #cgo linux,386,static,harfbuzz,subset LDFLAGS: -L${SRCDIR}/../linux_386/lib -lfreetypehb -lharfbuzz -lharfbuzz-subset -lbz2 -lpng16 -lz -lm
//
// #cgo linux,amd64,static CFLAGS: -I${SRCDIR}/../linux_amd64/include/freetype2 -Werror -Wall -Wextra -Wno-unused-parameter
// #cgo linux,amd64,static,!harfbuzz LDFLAGS: -L${SRCDIR}/../linux_amd64/lib -lfreetype -lbz2 -lpng16 -lz -lm
// #cgo linux,amd64,static,harfbuzz LDFLAGS: -L${SRCDIR}/../linux_amd64/lib -lfreetypehb -lharfbuzz -lfreetypehb -lbz2 -lpng16 -lz -lm
// #cgo linux,amd64,static,harfbuzz,subset LDFLAGS: -L${SRCDIR}/../linux_amd64/lib -lfreetypehb -lharfbuzz -lharfbuzz-subset -lbz2 -lpng16 -lz -lm
//
// #cgo darwin,386,static CFLAGS: -I${SRCDIR}/../darwin_386/include/freetype2 -Werror -Wall -Wextra -Wno-unused-parameter
// #cgo darwin,386,static,!harfbuzz LDFLAGS: -L${SRCDIR}/../darwin_386/lib -lfreetype -lbz2 -lpng16 -lz -lm
// #cgo darwin,386,static,harfbuzz LDFLAGS: -L${SRCDIR}/../darwin_386/lib -lfreetypehb -lharfbuzz -lfreetypehb -lbz2 -lpng16 -lz -lm
// #cgo darwin,386,static,harfbuzz,subset LDFLAGS: -L${SRCDIR}/../darwin_386/lib -lfreetypehb -lharfbuzz -lharfbuzz-subset -lbz2 -lpng16 -lz -lm
//
// #cgo darwin,amd64,static CFLAGS: -I${SRCDIR}/../darwin_amd64/include/freetype2 -Werror -Wall -Wextra -Wno-unused-parameter
// #cgo darwin,amd64,static,!harfbuzz LDFLAGS: -L${SRCDIR}/../darwin_amd64/lib -lfreetype -lbz2 -lpng16 -lz -lm
// #cgo darwin,amd64,static,harfbuzz LDFLAGS: -L${SRCDIR}/../darwin_amd64/lib -lfreetypehb -lharfbuzz -lfreetypehb -lbz2 -lpng16 -lz -lm
// #cgo darwin,amd64,static,harfbuzz,subset LDFLAGS: -L${SRCDIR}/../darwin_amd64/lib -lfreetypehb -lharfbuzz -lharfbuzz-subset -lbz2 -lpng16 -lz -lm
//
// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_TRUETYPE_TABLES_H
// #include FT_TRUETYPE_IDS_H
//
// #define my_explicit_uint32_TT_UCR_GENERAL_PUNCTUATION (unsigned long) TT_UCR_GENERAL_PUNCTUATION
// #define my_explicit_uint32_TT_UCR_ARABIC_PRESENTATION_FORMS_A (unsigned long) TT_UCR_ARABIC_PRESENTATION_FORMS_A
// #define my_explicit_uint32_TT_UCR_NEW_TAI_LUE (unsigned long) TT_UCR_NEW_TAI_LUE
import (
	"C"
)
import (
	"time"

	"github.com/flga/freetype2/2.10.1/fixed"
)

// Header models a TrueType font header table. All fields follow the OpenType specification.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_header
type Header struct {
	TableVersion fixed.Int16_16
	FontRevision fixed.Int16_16

	CheckSumAdjust int32
	MagicNumber    int32

	Flags      uint16
	UnitsPerEM uint16

	Created  time.Time
	Modified time.Time

	XMin int16
	YMin int16
	XMax int16
	YMax int16

	MacStyle      uint16
	LowestRecPPEM uint16

	FontDirection    int16
	IndexToLocFormat int16
	GlyphDataFormat  int16
}

// HoriHeader models a TrueType horizontal header, the ‘hhea’ table, as well as
// the corresponding horizontal metrics table, ‘hmtx’.
//
// NOTE: For an OpenType variation font, the values of the following fields
// can change after a call to SetVarDesignCoordinates (and friends) if the font
// contains an ‘MVAR’ table: CaretSlopeRise, CaretSlopeRun, and CaretOffset.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_horiheader
type HoriHeader struct {
	// The table version.
	Version fixed.Int16_16
	// The font's ascender, i.e., the distance from the baseline to the top-most
	// of all glyph points found in the font.
	// This value is invalid in many fonts, as it is usually set by the font
	// designer, and often reflects only a portion of the glyphs found in the
	// font (maybe ASCII).
	// You should use the sTypoAscender field of the ‘OS/2’ table instead if you
	// want the correct one.
	Ascender int16
	// The font's descender, i.e., the distance from the baseline to the
	// bottom-most of all glyph points found in the font. It is negative.
	// This value is invalid in many fonts, as it is usually set by the font
	// designer, and often reflects only a portion of the glyphs found in the
	// font (maybe ASCII).
	// You should use the sTypoDescender field of the ‘OS/2’ table instead if
	// you want the correct one.
	Descender int16
	// The font's line gap, i.e., the distance to add to the ascender and
	// descender to get the BTB, i.e., the baseline-to-baseline distance for
	// the font.
	LineGap int16

	// The maximum of all advance widths found in the font. It can be used to
	// compute the maximum width of an arbitrary string of text.
	AdvanceWidthMax uint16

	// The minimum left side bearing of all glyphs within the font.
	MinLeftSideBearing int16
	// The minimum right side bearing of all glyphs within the font.
	MinRightSideBearing int16
	// The maximum horizontal extent (i.e., the ‘width’ of a glyph's bounding
	// box) for all glyphs in the font.
	XMaxExtent int16
	// The rise coefficient of the cursor's slope of the cursor (slope=rise/run).
	// For an OpenType variation font, the value of this field can change after
	// a call to SetVarDesignCoordinates (and friends) if the font contains an
	// ‘MVAR’ table.
	CaretSlopeRise int16
	// The run coefficient of the cursor's slope.
	// For an OpenType variation font, the value of this field can change after
	// a call to SetVarDesignCoordinates (and friends) if the font contains an
	// ‘MVAR’ table.
	CaretSlopeRun int16
	// The cursor's offset for slanted fonts.
	// For an OpenType variation font, the value of this field can change after
	// a call to SetVarDesignCoordinates (and friends) if the font contains an
	// ‘MVAR’ table.
	CaretOffset int16

	// Always 0
	MetricDataFormat int16
	// Number of HMetrics entries in the ‘hmtx’ table -- this value can be
	// smaller than the total number of glyphs in the font.
	NumberOfHMetrics uint16

	// void*      long_metrics;
	// void*      short_metrics;

}

// VertHeader models a TrueType vertical header, the ‘vhea’ table, as well as
// the corresponding vertical metrics table, ‘vmtx’.
//
// NOTE: For an OpenType variation font, the values of the following fields can
// change after a call to SetVarDesignCoordinates (and friends) if the font
// contains an ‘MVAR’ table: Ascender, Descender, LineGap, CaretSlopeRise,
// CaretSlopeRun, and  CaretOffset.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_vertheader
type VertHeader struct {
	// The table version.
	Version fixed.Int16_16
	// The font's ascender, i.e., the distance from the baseline to the top-most
	// of all glyph points found in the font.
	// This value is invalid in many fonts, as it is usually set by the font
	// designer, and often reflects only a portion of the glyphs found in the
	// font (maybe ASCII).
	// You should use the sTypoAscender field of the ‘OS/2’ table instead if you
	// want the correct one.
	Ascender int16
	// The font's descender, i.e., the distance from the baseline to the
	// bottom-most of all glyph points found in the font. It is negative.
	// This value is invalid in many fonts, as it is usually set by the font
	// designer, and often reflects only a portion of the glyphs found in the
	// font (maybe ASCII).
	// You should use the sTypoDescender field of the ‘OS/2’ table instead if
	// you want the correct one.
	Descender int16
	// The font's line gap, i.e., the distance to add to the ascender and
	// descender to get the BTB, i.e., the baseline-to-baseline distance for the
	// font.
	LineGap int16

	// The maximum of all advance heights found in the font. It can be used to
	// compute the maximum height of an arbitrary string of text.
	AdvanceHeightMax uint16

	// The minimum top side bearing of all glyphs within the font.
	MinTopSideBearing int16
	// The minimum bottom side bearing of all glyphs within the font.
	MinBottomSideBearing int16
	// The maximum vertical extent (i.e., the ‘height’ of a glyph's bounding
	// box) for all glyphs in the font.
	YMaxExtent int16
	// The rise coefficient of the cursor's slope of the cursor (slope=rise/run).
	CaretSlopeRise int16
	// The run coefficient of the cursor's slope.
	CaretSlopeRun int16
	// The cursor's offset for slanted fonts.
	CaretOffset int16

	// Always 0.
	MetricDataFormat int16
	// Number of VMetrics entries in the ‘vmtx’ table -- this value can be
	// smaller than the total number of glyphs in the font.
	NumberOfVMetrics uint16

	//   void*      long_metrics;
	//   void*      short_metrics;
}

// OS2 models a TrueType ‘OS/2’ table. All fields comply to the OpenType
// specification.
//
// Note that we now support old Mac fonts that do not include an ‘OS/2’ table.
// In this case, the version field is always set to 0xFFFF.
//
// NOTE: For an OpenType variation font, the values of the following fields can
// change after a call to SetVarDesignCoordinates (and friends) if the font
// contains an ‘MVAR’ table: SCapHeight, STypoAscender, STypoDescender,
// STypoLineGap, SxHeight, UsWinAscent, UsWinDescent, YStrikeoutPosition,
// YStrikeoutSize, YSubscriptXOffset, YSubScriptXSize, YSubscriptYOffset,
// YSubscriptYSize, YSuperscriptXOffset, YSuperscriptXSize, YSuperscriptYOffset,
// and YSuperscriptYSize.
//
// Possible values for bits in the ulUnicodeRangeX fields are given by the TODO: TT_UCR_XXX macros.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_os2
type OS2 struct {
	Version       uint16 // 0x0001 - more or 0xFFFF
	XAvgCharWidth int16
	UsWeightClass uint16
	UsWidthClass  uint16
	FsType        uint16
	// The value of the following fields can change after a call to
	// SetVarDesignCoordinates (and friends) if the font contains an ‘MVAR’
	// table.
	YSubscriptXSize     int16
	YSubscriptYSize     int16
	YSubscriptXOffset   int16
	YSubscriptYOffset   int16
	YSuperscriptXSize   int16
	YSuperscriptYSize   int16
	YSuperscriptXOffset int16
	YSuperscriptYOffset int16
	YStrikeoutSize      int16
	YStrikeoutPosition  int16
	SFamilyClass        int16

	Panose [10]byte

	UlUnicodeRange1 UCRMask // Bits 0-31
	UlUnicodeRange2 UCRMask // Bits 32-63
	UlUnicodeRange3 UCRMask // Bits 64-95
	UlUnicodeRange4 UCRMask // Bits 96-127

	AchVendID [4]int8

	FsSelection      uint16
	UsFirstCharIndex uint16
	UsLastCharIndex  uint16

	// The value of the following fields can change after a call to
	// SetVarDesignCoordinates (and friends) if the font contains an ‘MVAR’
	// table.
	STypoAscender  int16
	STypoDescender int16
	STypoLineGap   int16
	UsWinAscent    uint16
	UsWinDescent   uint16

	// only version 1 and higher
	UlCodePageRange1 uint32 // Bits 0-31
	UlCodePageRange2 uint32 // Bits 32-63

	// only version 2 and higher
	// This value can change after a call to SetVarDesignCoordinates (and
	// friends) if the font contains an ‘MVAR’ table.
	SxHeight int16
	// This value can change after a call to SetVarDesignCoordinates (and
	// friends) if the font contains an ‘MVAR’ table.
	SCapHeight    int16
	UsDefaultChar uint16
	UsBreakChar   uint16
	UsMaxContext  uint16

	// only version 5 and higher
	UsLowerOpticalPointSize uint16 // in twips (1/20th points)
	UsUpperOpticalPointSize uint16 // in twips (1/20th points)
}

// Postscript models a TrueType ‘post’ table. All fields comply to the OpenType
// specification. This struct does not reference a font's PostScript glyph names;
// use GetGlyphName to retrieve them.
//
// NOTE: For an OpenType variation font, the values of the following fields can
// change after a call to SetVarDesignCoordinates (and friends) if the font
// contains an ‘MVAR’ table: UnderlinePosition and UnderlineThickness.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_postscript
type Postscript struct {
	FormatType  fixed.Int16_16
	ItalicAngle fixed.Int16_16
	// This value can change after a call to SetVarDesignCoordinates (and
	// friends) if the font contains an ‘MVAR’ table.
	UnderlinePosition int16
	// This value can change after a call to SetVarDesignCoordinates (and
	// friends) if the font contains an ‘MVAR’ table.
	UnderlineThickness int16
	IsFixedPitch       uint32
	MinMemType42       uint32
	MaxMemType42       uint32
	MinMemType1        uint32
	MaxMemType1        uint32
}

// PCLT models a TrueType ‘PCLT’ table. All fields comply to the OpenType
// specification.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_pclt
type PCLT struct {
	Version             fixed.Int16_16
	FontNumber          uint32
	Pitch               uint16
	XHeight             uint16
	Style               uint16
	TypeFamily          uint16
	CapHeight           uint16
	SymbolSet           uint16
	TypeFace            [16]int8
	CharacterComplement [8]int8
	FileName            [6]int8
	StrokeWeight        int8
	WidthType           int8
	SerifStyle          byte
}

// MaxProfile models the (‘maxp’) table and contains many max values, which can
// be used to pre-allocate arrays for speeding up glyph loading and hinting.
//
// NOTE: This struct is only used during font loading.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_maxprofile
type MaxProfile struct {
	// The version number.
	Version fixed.Int16_16
	// The number of glyphs in this TrueType font.
	NumGlyphs uint16
	// The maximum number of points in a non-composite TrueType glyph.
	// See also MaxCompositePoints.
	MaxPoints uint16
	// The maximum number of contours in a non-composite TrueType glyph.
	// See also MaxCompositeContours.
	MaxContours uint16
	// The maximum number of points in a composite TrueType glyph.
	// See also MaxPoints.
	MaxCompositePoints uint16
	// The maximum number of contours in a composite TrueType glyph.
	// See also MaxContours.
	MaxCompositeContours uint16
	// The maximum number of zones used for glyph hinting.
	MaxZones uint16
	// The maximum number of points in the twilight zone used for glyph hinting.
	MaxTwilightPoints uint16
	// The maximum number of elements in the storage area used for glyph hinting.
	MaxStorage uint16
	// The maximum number of function definitions in the TrueType bytecode for
	// this font.
	MaxFunctionDefs uint16
	// The maximum number of instruction definitions in the TrueType bytecode
	// for this font.
	MaxInstructionDefs uint16
	// The maximum number of stack elements used during bytecode interpretation.
	MaxStackElements uint16
	// The maximum number of TrueType opcodes used for glyph hinting.
	MaxSizeOfInstructions uint16
	// The maximum number of simple (i.e., non-composite) glyphs in a composite
	// glyph.
	MaxComponentElements uint16
	// The maximum nesting depth of composite glyphs.
	MaxComponentDepth uint16
}

// CmapFormat is an enumeration of the cmap format values
//
// See https://docs.microsoft.com/en-us/typography/opentype/spec/cmap
type CmapFormat int

const (
	// ByteEncodingTable is the Apple standard character to glyph index mapping table.
	ByteEncodingTable CmapFormat = 0

	// HighByteMappingThroughTable is useful for the national character code standards used for Japanese, Chinese, and
	// Korean characters. These code standards use a mixed 8-/16-bit encoding, in which certain byte values signal the
	// first byte of a 2-byte character (but these values are also legal as the second byte of a 2-byte character).
	HighByteMappingThroughTable CmapFormat = 2
	// SegmentMappingToDeltaValues is the standard character-to-glyph-index mapping table for the Windows platform for
	// fonts that support Unicode BMP characters.
	SegmentMappingToDeltaValues CmapFormat = 4
	// TrimmedTableMapping should be used when character codes for a font fall into a single contiguous range. This
	// results in what is termed a dense mapping. Two-byte fonts that are not densely mapped (due to their multiple
	// contiguous ranges) should use Format 4.
	TrimmedTableMapping CmapFormat = 6
	// Mixed16And32bitCoverage is similar to HighByteMappingThroughTable, in that it provides for mixed-length character
	// codes. Instead of allowing for 8 and 16-bit character codes, however, it allows for 16 and 32-bit character codes.
	Mixed16And32bitCoverage CmapFormat = 8
	//TrimmedArray is similar to format TrimmedTableMapping, in that it defines a trimmed array for a tight range of
	// character codes. It differs, however, in that is uses 32-bit character codes
	TrimmedArray CmapFormat = 10
	// SegmentedCoverage is the standard character-to-glyph-index mapping table for the Windows platform for fonts
	// supporting Unicode supplementary-plane characters (U+10000 to U+10FFFF).
	SegmentedCoverage CmapFormat = 12
	// ManyToOneRangeMappings provides for situations in which the same glyph is used for hundreds or even thousands of
	// consecutive characters spanning across multiple ranges of the code space.
	ManyToOneRangeMappings CmapFormat = 13
	// UnicodeVariationSequences specifies the Unicode Variation Sequences (UVSes) supported by the font. A Variation
	// Sequence, according to the Unicode Standard, comprises a base character followed by a variation selector.
	// For example, <U+82A6, U+E0101>.
	UnicodeVariationSequences CmapFormat = 14
)

// PlatformID is an enum for the PlatformID in CharMap and SfntName structs.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_platform_xxx
type PlatformID int

const (
	// PlatformAppleUnicode is used by Apple to indicate a Unicode character map and/or name entry. See AppleEncodingIDs
	// for corresponding values. Note that name entries in this format are coded as big-endian UCS-2 character codes only.
	PlatformAppleUnicode PlatformID = C.TT_PLATFORM_APPLE_UNICODE
	// PlatformMacintosh is used by Apple to indicate a MacOS-specific charmap and/or name entry. See MacEncodingIDs for
	// corresponding values. Note that most TrueType fonts contain an Apple roman charmap to be usable on MacOS systems
	// (even if they contain a Microsoft charmap as well).
	PlatformMacintosh PlatformID = C.TT_PLATFORM_MACINTOSH
	// PlatformMicrosoft is used by Microsoft to indicate Windows-specific charmaps. See MicrosoftEncodingIDs for
	// corresponding values. Note that most fonts contain a Unicode charmap using (PlatformMicrosoft, EncodingMSIDUnicodeCS).
	PlatformMicrosoft PlatformID = C.TT_PLATFORM_MICROSOFT
	// PlatformCustom is used to indicate application-specific charmaps.
	PlatformCustom PlatformID = C.TT_PLATFORM_CUSTOM
	//PlatformAdobe isn't part of any font format specification, but is used by FreeType to report Adobe-specific
	// charmaps in a CharMap struct. See AdobeEncodingIDs.
	PlatformAdobe PlatformID = C.TT_PLATFORM_ADOBE
)

func (p PlatformID) String() string {
	switch p {
	case PlatformAppleUnicode:
		return "AppleUnicode"
	case PlatformMacintosh:
		return "Macintosh"
	case PlatformMicrosoft:
		return "Microsoft"
	case PlatformCustom:
		return "Custom"
	case PlatformAdobe:
		return "Adobe"
	default:
		return "Unknown"
	}
}

// EncodingID is an enum for EncodingID in CharMap and SfntName structs.
type EncodingID int

// EncodingIDs for PlatformAppleUnicode
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_apple_id_xxx
const (
	// Unicode version 1.0.
	AppleEncodingDefault EncodingID = C.TT_APPLE_ID_DEFAULT
	// Unicode 1.1; specifies Hangul characters starting at U+34xx.
	AppleEncodingUnicode1_1 EncodingID = C.TT_APPLE_ID_UNICODE_1_1
	// Unicode 2.0 and beyond (UTF-16 BMP only).
	AppleEncodingUnicode2_0 EncodingID = C.TT_APPLE_ID_UNICODE_2_0
	// Unicode 3.1 and beyond, using UTF-32.
	AppleEncodingUnicode32 EncodingID = C.TT_APPLE_ID_UNICODE_32
	// From Adobe, not Apple. Not a normal cmap. Specifies variations on a real cmap.
	AppleEncodingVariantSelector EncodingID = C.TT_APPLE_ID_VARIANT_SELECTOR
	// Used for fallback fonts that provide complete Unicode coverage with a type 13 cmap.
	AppleEncodingFullUnicode EncodingID = C.TT_APPLE_ID_FULL_UNICODE
)

// EncodingIDs for PlatformMacintosh
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_mac_id_xxx
const (
	MacEncodingRoman              EncodingID = C.TT_MAC_ID_ROMAN
	MacEncodingJapanese           EncodingID = C.TT_MAC_ID_JAPANESE
	MacEncodingTraditionalChinese EncodingID = C.TT_MAC_ID_TRADITIONAL_CHINESE
	MacEncodingKorean             EncodingID = C.TT_MAC_ID_KOREAN
	MacEncodingArabic             EncodingID = C.TT_MAC_ID_ARABIC
	MacEncodingHebrew             EncodingID = C.TT_MAC_ID_HEBREW
	MacEncodingGreek              EncodingID = C.TT_MAC_ID_GREEK
	MacEncodingRussian            EncodingID = C.TT_MAC_ID_RUSSIAN
	MacEncodingRsymbol            EncodingID = C.TT_MAC_ID_RSYMBOL
	MacEncodingDevanagari         EncodingID = C.TT_MAC_ID_DEVANAGARI
	MacEncodingGurmukhi           EncodingID = C.TT_MAC_ID_GURMUKHI
	MacEncodingGujarati           EncodingID = C.TT_MAC_ID_GUJARATI
	MacEncodingOriya              EncodingID = C.TT_MAC_ID_ORIYA
	MacEncodingBengali            EncodingID = C.TT_MAC_ID_BENGALI
	MacEncodingTamil              EncodingID = C.TT_MAC_ID_TAMIL
	MacEncodingTelugu             EncodingID = C.TT_MAC_ID_TELUGU
	MacEncodingKannada            EncodingID = C.TT_MAC_ID_KANNADA
	MacEncodingMalayalam          EncodingID = C.TT_MAC_ID_MALAYALAM
	MacEncodingSinhalese          EncodingID = C.TT_MAC_ID_SINHALESE
	MacEncodingBurmese            EncodingID = C.TT_MAC_ID_BURMESE
	MacEncodingKhmer              EncodingID = C.TT_MAC_ID_KHMER
	MacEncodingThai               EncodingID = C.TT_MAC_ID_THAI
	MacEncodingLaotian            EncodingID = C.TT_MAC_ID_LAOTIAN
	MacEncodingGeorgian           EncodingID = C.TT_MAC_ID_GEORGIAN
	MacEncodingArmenian           EncodingID = C.TT_MAC_ID_ARMENIAN
	MacEncodingMaldivian          EncodingID = C.TT_MAC_ID_MALDIVIAN
	MacEncodingSimplifiedChinese  EncodingID = C.TT_MAC_ID_SIMPLIFIED_CHINESE
	MacEncodingTibetan            EncodingID = C.TT_MAC_ID_TIBETAN
	MacEncodingMongolian          EncodingID = C.TT_MAC_ID_MONGOLIAN
	MacEncodingGeez               EncodingID = C.TT_MAC_ID_GEEZ
	MacEncodingSlavic             EncodingID = C.TT_MAC_ID_SLAVIC
	MacEncodingVietnamese         EncodingID = C.TT_MAC_ID_VIETNAMESE
	MacEncodingSindhi             EncodingID = C.TT_MAC_ID_SINDHI
	MacEncodingUninterp           EncodingID = C.TT_MAC_ID_UNINTERP
)

// EncodingIDs for PlatformMicrosoft
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_ms_id_xxx
const (
	// Microsoft symbol encoding. See EncodingMsSymbol
	MicrosoftEncodingSymbolCs EncodingID = C.TT_MS_ID_SYMBOL_CS
	// Microsoft WGL4 charmap, matching Unicode. See EncodingUnicode.
	MicrosoftEncodingUnicodeCs EncodingID = C.TT_MS_ID_UNICODE_CS
	// Shift JIS Japanese encoding. See EncodingSjis.
	MicrosoftEncodingSjis EncodingID = C.TT_MS_ID_SJIS
	// Chinese encodings as used in the People's Republic of China (PRC). This means the encodings GB 2312 and its
	// supersets GBK and GB 18030. See  EncodingPrc.
	MicrosoftEncodingPrc EncodingID = C.TT_MS_ID_PRC
	// Traditional Chinese as used in Taiwan and Hong Kong. See EncodingBig5.
	MicrosoftEncodingBig5 EncodingID = C.TT_MS_ID_BIG_5
	// Korean Extended Wansung encoding. See EncodingWansung.
	MicrosoftEncodingWansung EncodingID = C.TT_MS_ID_WANSUNG
	// Korean Johab encoding. See EncodingJohab.
	MicrosoftEncodingJohab EncodingID = C.TT_MS_ID_JOHAB
	// UCS-4 or UTF-32 charmaps. This has been added to the OpenType specification version 1.4 (mid-2001).
	MicrosoftEncodingUCS4 EncodingID = C.TT_MS_ID_UCS_4
)

// EncodingIDs for PlatformAdobe
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_adobe_id_xxx
const (
	AdobeEncodingStandard EncodingID = C.TT_ADOBE_ID_STANDARD
	AdobeEncodingExpert   EncodingID = C.TT_ADOBE_ID_EXPERT
	AdobeEncodingCustom   EncodingID = C.TT_ADOBE_ID_CUSTOM
	AdobeEncodingLatin1   EncodingID = C.TT_ADOBE_ID_LATIN_1
)

// LanguageID is an enum of possible values of the language identifier field in the name records of the
// SFNT ‘name’ table.
type LanguageID int

// LangIDs for PlatformMacintosh
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_mac_langid_xxx
const (
	MacLangEnglish                   LanguageID = C.TT_MAC_LANGID_ENGLISH
	MacLangFrench                    LanguageID = C.TT_MAC_LANGID_FRENCH
	MacLangGerman                    LanguageID = C.TT_MAC_LANGID_GERMAN
	MacLangItalian                   LanguageID = C.TT_MAC_LANGID_ITALIAN
	MacLangDutch                     LanguageID = C.TT_MAC_LANGID_DUTCH
	MacLangSwedish                   LanguageID = C.TT_MAC_LANGID_SWEDISH
	MacLangSpanish                   LanguageID = C.TT_MAC_LANGID_SPANISH
	MacLangDanish                    LanguageID = C.TT_MAC_LANGID_DANISH
	MacLangPortuguese                LanguageID = C.TT_MAC_LANGID_PORTUGUESE
	MacLangNorwegian                 LanguageID = C.TT_MAC_LANGID_NORWEGIAN
	MacLangHebrew                    LanguageID = C.TT_MAC_LANGID_HEBREW
	MacLangJapanese                  LanguageID = C.TT_MAC_LANGID_JAPANESE
	MacLangArabic                    LanguageID = C.TT_MAC_LANGID_ARABIC
	MacLangFinnish                   LanguageID = C.TT_MAC_LANGID_FINNISH
	MacLangGreek                     LanguageID = C.TT_MAC_LANGID_GREEK
	MacLangIcelandic                 LanguageID = C.TT_MAC_LANGID_ICELANDIC
	MacLangMaltese                   LanguageID = C.TT_MAC_LANGID_MALTESE
	MacLangTurkish                   LanguageID = C.TT_MAC_LANGID_TURKISH
	MacLangCroatian                  LanguageID = C.TT_MAC_LANGID_CROATIAN
	MacLangChineseTraditional        LanguageID = C.TT_MAC_LANGID_CHINESE_TRADITIONAL
	MacLangUrdu                      LanguageID = C.TT_MAC_LANGID_URDU
	MacLangHindi                     LanguageID = C.TT_MAC_LANGID_HINDI
	MacLangThai                      LanguageID = C.TT_MAC_LANGID_THAI
	MacLangKorean                    LanguageID = C.TT_MAC_LANGID_KOREAN
	MacLangLithuanian                LanguageID = C.TT_MAC_LANGID_LITHUANIAN
	MacLangPolish                    LanguageID = C.TT_MAC_LANGID_POLISH
	MacLangHungarian                 LanguageID = C.TT_MAC_LANGID_HUNGARIAN
	MacLangEstonian                  LanguageID = C.TT_MAC_LANGID_ESTONIAN
	MacLangLettish                   LanguageID = C.TT_MAC_LANGID_LETTISH
	MacLangSaamisk                   LanguageID = C.TT_MAC_LANGID_SAAMISK
	MacLangFaeroese                  LanguageID = C.TT_MAC_LANGID_FAEROESE
	MacLangFarsi                     LanguageID = C.TT_MAC_LANGID_FARSI
	MacLangRussian                   LanguageID = C.TT_MAC_LANGID_RUSSIAN
	MacLangChineseSimplified         LanguageID = C.TT_MAC_LANGID_CHINESE_SIMPLIFIED
	MacLangFlemish                   LanguageID = C.TT_MAC_LANGID_FLEMISH
	MacLangIrish                     LanguageID = C.TT_MAC_LANGID_IRISH
	MacLangAlbanian                  LanguageID = C.TT_MAC_LANGID_ALBANIAN
	MacLangRomanian                  LanguageID = C.TT_MAC_LANGID_ROMANIAN
	MacLangCzech                     LanguageID = C.TT_MAC_LANGID_CZECH
	MacLangSlovak                    LanguageID = C.TT_MAC_LANGID_SLOVAK
	MacLangSlovenian                 LanguageID = C.TT_MAC_LANGID_SLOVENIAN
	MacLangYiddish                   LanguageID = C.TT_MAC_LANGID_YIDDISH
	MacLangSerbian                   LanguageID = C.TT_MAC_LANGID_SERBIAN
	MacLangMacedonian                LanguageID = C.TT_MAC_LANGID_MACEDONIAN
	MacLangBulgarian                 LanguageID = C.TT_MAC_LANGID_BULGARIAN
	MacLangUkrainian                 LanguageID = C.TT_MAC_LANGID_UKRAINIAN
	MacLangByelorussian              LanguageID = C.TT_MAC_LANGID_BYELORUSSIAN
	MacLangUzbek                     LanguageID = C.TT_MAC_LANGID_UZBEK
	MacLangKazakh                    LanguageID = C.TT_MAC_LANGID_KAZAKH
	MacLangAzerbaijani               LanguageID = C.TT_MAC_LANGID_AZERBAIJANI
	MacLangAzerbaijaniCyrillicScript LanguageID = C.TT_MAC_LANGID_AZERBAIJANI_CYRILLIC_SCRIPT
	MacLangAzerbaijaniArabicScript   LanguageID = C.TT_MAC_LANGID_AZERBAIJANI_ARABIC_SCRIPT
	MacLangArmenian                  LanguageID = C.TT_MAC_LANGID_ARMENIAN
	MacLangGeorgian                  LanguageID = C.TT_MAC_LANGID_GEORGIAN
	MacLangMoldavian                 LanguageID = C.TT_MAC_LANGID_MOLDAVIAN
	MacLangKirghiz                   LanguageID = C.TT_MAC_LANGID_KIRGHIZ
	MacLangTajiki                    LanguageID = C.TT_MAC_LANGID_TAJIKI
	MacLangTurkmen                   LanguageID = C.TT_MAC_LANGID_TURKMEN
	MacLangMongolian                 LanguageID = C.TT_MAC_LANGID_MONGOLIAN
	MacLangMongolianMongolianScript  LanguageID = C.TT_MAC_LANGID_MONGOLIAN_MONGOLIAN_SCRIPT
	MacLangMongolianCyrillicScript   LanguageID = C.TT_MAC_LANGID_MONGOLIAN_CYRILLIC_SCRIPT
	MacLangPashto                    LanguageID = C.TT_MAC_LANGID_PASHTO
	MacLangKurdish                   LanguageID = C.TT_MAC_LANGID_KURDISH
	MacLangKashmiri                  LanguageID = C.TT_MAC_LANGID_KASHMIRI
	MacLangSindhi                    LanguageID = C.TT_MAC_LANGID_SINDHI
	MacLangTibetan                   LanguageID = C.TT_MAC_LANGID_TIBETAN
	MacLangNepali                    LanguageID = C.TT_MAC_LANGID_NEPALI
	MacLangSanskrit                  LanguageID = C.TT_MAC_LANGID_SANSKRIT
	MacLangMarathi                   LanguageID = C.TT_MAC_LANGID_MARATHI
	MacLangBengali                   LanguageID = C.TT_MAC_LANGID_BENGALI
	MacLangAssamese                  LanguageID = C.TT_MAC_LANGID_ASSAMESE
	MacLangGujarati                  LanguageID = C.TT_MAC_LANGID_GUJARATI
	MacLangPunjabi                   LanguageID = C.TT_MAC_LANGID_PUNJABI
	MacLangOriya                     LanguageID = C.TT_MAC_LANGID_ORIYA
	MacLangMalayalam                 LanguageID = C.TT_MAC_LANGID_MALAYALAM
	MacLangKannada                   LanguageID = C.TT_MAC_LANGID_KANNADA
	MacLangTamil                     LanguageID = C.TT_MAC_LANGID_TAMIL
	MacLangTelugu                    LanguageID = C.TT_MAC_LANGID_TELUGU
	MacLangSinhalese                 LanguageID = C.TT_MAC_LANGID_SINHALESE
	MacLangBurmese                   LanguageID = C.TT_MAC_LANGID_BURMESE
	MacLangKhmer                     LanguageID = C.TT_MAC_LANGID_KHMER
	MacLangLao                       LanguageID = C.TT_MAC_LANGID_LAO
	MacLangVietnamese                LanguageID = C.TT_MAC_LANGID_VIETNAMESE
	MacLangIndonesian                LanguageID = C.TT_MAC_LANGID_INDONESIAN
	MacLangTagalog                   LanguageID = C.TT_MAC_LANGID_TAGALOG
	MacLangMalayRomanScript          LanguageID = C.TT_MAC_LANGID_MALAY_ROMAN_SCRIPT
	MacLangMalayArabicScript         LanguageID = C.TT_MAC_LANGID_MALAY_ARABIC_SCRIPT
	MacLangAmharic                   LanguageID = C.TT_MAC_LANGID_AMHARIC
	MacLangTigrinya                  LanguageID = C.TT_MAC_LANGID_TIGRINYA
	MacLangGalla                     LanguageID = C.TT_MAC_LANGID_GALLA
	MacLangSomali                    LanguageID = C.TT_MAC_LANGID_SOMALI
	MacLangSwahili                   LanguageID = C.TT_MAC_LANGID_SWAHILI
	MacLangRuanda                    LanguageID = C.TT_MAC_LANGID_RUANDA
	MacLangRundi                     LanguageID = C.TT_MAC_LANGID_RUNDI
	MacLangChewa                     LanguageID = C.TT_MAC_LANGID_CHEWA
	MacLangMalagasy                  LanguageID = C.TT_MAC_LANGID_MALAGASY
	MacLangEsperanto                 LanguageID = C.TT_MAC_LANGID_ESPERANTO
	MacLangWelsh                     LanguageID = C.TT_MAC_LANGID_WELSH
	MacLangBasque                    LanguageID = C.TT_MAC_LANGID_BASQUE
	MacLangCatalan                   LanguageID = C.TT_MAC_LANGID_CATALAN
	MacLangLatin                     LanguageID = C.TT_MAC_LANGID_LATIN
	MacLangQuechua                   LanguageID = C.TT_MAC_LANGID_QUECHUA
	MacLangGuarani                   LanguageID = C.TT_MAC_LANGID_GUARANI
	MacLangAymara                    LanguageID = C.TT_MAC_LANGID_AYMARA
	MacLangTatar                     LanguageID = C.TT_MAC_LANGID_TATAR
	MacLangUighur                    LanguageID = C.TT_MAC_LANGID_UIGHUR
	MacLangDzongkha                  LanguageID = C.TT_MAC_LANGID_DZONGKHA
	MacLangJavanese                  LanguageID = C.TT_MAC_LANGID_JAVANESE
	MacLangSundanese                 LanguageID = C.TT_MAC_LANGID_SUNDANESE
	MacLangGalician                  LanguageID = C.TT_MAC_LANGID_GALICIAN
	MacLangAfrikaans                 LanguageID = C.TT_MAC_LANGID_AFRIKAANS
	MacLangBreton                    LanguageID = C.TT_MAC_LANGID_BRETON
	MacLangInuktitut                 LanguageID = C.TT_MAC_LANGID_INUKTITUT
	MacLangScottishGaelic            LanguageID = C.TT_MAC_LANGID_SCOTTISH_GAELIC
	MacLangManxGaelic                LanguageID = C.TT_MAC_LANGID_MANX_GAELIC
	MacLangIrishGaelic               LanguageID = C.TT_MAC_LANGID_IRISH_GAELIC
	MacLangTongan                    LanguageID = C.TT_MAC_LANGID_TONGAN
	MacLangGreekPolytonic            LanguageID = C.TT_MAC_LANGID_GREEK_POLYTONIC
	MacLangGreelandic                LanguageID = C.TT_MAC_LANGID_GREELANDIC
	MacLangAzerbaijaniRomanScript    LanguageID = C.TT_MAC_LANGID_AZERBAIJANI_ROMAN_SCRIPT
)

// LangIDs for PlatformMicrosoft
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_ms_langid_xxx
const (
	MicrosoftLangArabicSaudiArabia           LanguageID = C.TT_MS_LANGID_ARABIC_SAUDI_ARABIA
	MicrosoftLangArabicIraq                  LanguageID = C.TT_MS_LANGID_ARABIC_IRAQ
	MicrosoftLangArabicEgypt                 LanguageID = C.TT_MS_LANGID_ARABIC_EGYPT
	MicrosoftLangArabicLibya                 LanguageID = C.TT_MS_LANGID_ARABIC_LIBYA
	MicrosoftLangArabicAlgeria               LanguageID = C.TT_MS_LANGID_ARABIC_ALGERIA
	MicrosoftLangArabicMorocco               LanguageID = C.TT_MS_LANGID_ARABIC_MOROCCO
	MicrosoftLangArabicTunisia               LanguageID = C.TT_MS_LANGID_ARABIC_TUNISIA
	MicrosoftLangArabicOman                  LanguageID = C.TT_MS_LANGID_ARABIC_OMAN
	MicrosoftLangArabicYemen                 LanguageID = C.TT_MS_LANGID_ARABIC_YEMEN
	MicrosoftLangArabicSyria                 LanguageID = C.TT_MS_LANGID_ARABIC_SYRIA
	MicrosoftLangArabicJordan                LanguageID = C.TT_MS_LANGID_ARABIC_JORDAN
	MicrosoftLangArabicLebanon               LanguageID = C.TT_MS_LANGID_ARABIC_LEBANON
	MicrosoftLangArabicKuwait                LanguageID = C.TT_MS_LANGID_ARABIC_KUWAIT
	MicrosoftLangArabicUae                   LanguageID = C.TT_MS_LANGID_ARABIC_UAE
	MicrosoftLangArabicBahrain               LanguageID = C.TT_MS_LANGID_ARABIC_BAHRAIN
	MicrosoftLangArabicQatar                 LanguageID = C.TT_MS_LANGID_ARABIC_QATAR
	MicrosoftLangBulgarianBulgaria           LanguageID = C.TT_MS_LANGID_BULGARIAN_BULGARIA
	MicrosoftLangCatalanCatalan              LanguageID = C.TT_MS_LANGID_CATALAN_CATALAN
	MicrosoftLangChineseTaiwan               LanguageID = C.TT_MS_LANGID_CHINESE_TAIWAN
	MicrosoftLangChinesePrc                  LanguageID = C.TT_MS_LANGID_CHINESE_PRC
	MicrosoftLangChineseHongKong             LanguageID = C.TT_MS_LANGID_CHINESE_HONG_KONG
	MicrosoftLangChineseSingapore            LanguageID = C.TT_MS_LANGID_CHINESE_SINGAPORE
	MicrosoftLangChineseMacao                LanguageID = C.TT_MS_LANGID_CHINESE_MACAO
	MicrosoftLangCzechCzechRepublic          LanguageID = C.TT_MS_LANGID_CZECH_CZECH_REPUBLIC
	MicrosoftLangDanishDenmark               LanguageID = C.TT_MS_LANGID_DANISH_DENMARK
	MicrosoftLangGermanGermany               LanguageID = C.TT_MS_LANGID_GERMAN_GERMANY
	MicrosoftLangGermanSwitzerland           LanguageID = C.TT_MS_LANGID_GERMAN_SWITZERLAND
	MicrosoftLangGermanAustria               LanguageID = C.TT_MS_LANGID_GERMAN_AUSTRIA
	MicrosoftLangGermanLuxembourg            LanguageID = C.TT_MS_LANGID_GERMAN_LUXEMBOURG
	MicrosoftLangGermanLiechtenstein         LanguageID = C.TT_MS_LANGID_GERMAN_LIECHTENSTEIN
	MicrosoftLangGreekGreece                 LanguageID = C.TT_MS_LANGID_GREEK_GREECE
	MicrosoftLangEnglishUnitedStates         LanguageID = C.TT_MS_LANGID_ENGLISH_UNITED_STATES
	MicrosoftLangEnglishUnitedKingdom        LanguageID = C.TT_MS_LANGID_ENGLISH_UNITED_KINGDOM
	MicrosoftLangEnglishAustralia            LanguageID = C.TT_MS_LANGID_ENGLISH_AUSTRALIA
	MicrosoftLangEnglishCanada               LanguageID = C.TT_MS_LANGID_ENGLISH_CANADA
	MicrosoftLangEnglishNewZealand           LanguageID = C.TT_MS_LANGID_ENGLISH_NEW_ZEALAND
	MicrosoftLangEnglishIreland              LanguageID = C.TT_MS_LANGID_ENGLISH_IRELAND
	MicrosoftLangEnglishSouthAfrica          LanguageID = C.TT_MS_LANGID_ENGLISH_SOUTH_AFRICA
	MicrosoftLangEnglishJamaica              LanguageID = C.TT_MS_LANGID_ENGLISH_JAMAICA
	MicrosoftLangEnglishCaribbean            LanguageID = C.TT_MS_LANGID_ENGLISH_CARIBBEAN
	MicrosoftLangEnglishBelize               LanguageID = C.TT_MS_LANGID_ENGLISH_BELIZE
	MicrosoftLangEnglishTrinidad             LanguageID = C.TT_MS_LANGID_ENGLISH_TRINIDAD
	MicrosoftLangEnglishZimbabwe             LanguageID = C.TT_MS_LANGID_ENGLISH_ZIMBABWE
	MicrosoftLangEnglishPhilippines          LanguageID = C.TT_MS_LANGID_ENGLISH_PHILIPPINES
	MicrosoftLangEnglishIndia                LanguageID = C.TT_MS_LANGID_ENGLISH_INDIA
	MicrosoftLangEnglishMalaysia             LanguageID = C.TT_MS_LANGID_ENGLISH_MALAYSIA
	MicrosoftLangEnglishSingapore            LanguageID = C.TT_MS_LANGID_ENGLISH_SINGAPORE
	MicrosoftLangSpanishSpainTraditionalSort LanguageID = C.TT_MS_LANGID_SPANISH_SPAIN_TRADITIONAL_SORT
	MicrosoftLangSpanishMexico               LanguageID = C.TT_MS_LANGID_SPANISH_MEXICO
	MicrosoftLangSpanishSpainModernSort      LanguageID = C.TT_MS_LANGID_SPANISH_SPAIN_MODERN_SORT
	MicrosoftLangSpanishGuatemala            LanguageID = C.TT_MS_LANGID_SPANISH_GUATEMALA
	MicrosoftLangSpanishCostaRica            LanguageID = C.TT_MS_LANGID_SPANISH_COSTA_RICA
	MicrosoftLangSpanishPanama               LanguageID = C.TT_MS_LANGID_SPANISH_PANAMA
	MicrosoftLangSpanishDominicanRepublic    LanguageID = C.TT_MS_LANGID_SPANISH_DOMINICAN_REPUBLIC
	MicrosoftLangSpanishVenezuela            LanguageID = C.TT_MS_LANGID_SPANISH_VENEZUELA
	MicrosoftLangSpanishColombia             LanguageID = C.TT_MS_LANGID_SPANISH_COLOMBIA
	MicrosoftLangSpanishPeru                 LanguageID = C.TT_MS_LANGID_SPANISH_PERU
	MicrosoftLangSpanishArgentina            LanguageID = C.TT_MS_LANGID_SPANISH_ARGENTINA
	MicrosoftLangSpanishEcuador              LanguageID = C.TT_MS_LANGID_SPANISH_ECUADOR
	MicrosoftLangSpanishChile                LanguageID = C.TT_MS_LANGID_SPANISH_CHILE
	MicrosoftLangSpanishUruguay              LanguageID = C.TT_MS_LANGID_SPANISH_URUGUAY
	MicrosoftLangSpanishParaguay             LanguageID = C.TT_MS_LANGID_SPANISH_PARAGUAY
	MicrosoftLangSpanishBolivia              LanguageID = C.TT_MS_LANGID_SPANISH_BOLIVIA
	MicrosoftLangSpanishElSalvador           LanguageID = C.TT_MS_LANGID_SPANISH_EL_SALVADOR
	MicrosoftLangSpanishHonduras             LanguageID = C.TT_MS_LANGID_SPANISH_HONDURAS
	MicrosoftLangSpanishNicaragua            LanguageID = C.TT_MS_LANGID_SPANISH_NICARAGUA
	MicrosoftLangSpanishPuertoRico           LanguageID = C.TT_MS_LANGID_SPANISH_PUERTO_RICO
	MicrosoftLangSpanishUnitedStates         LanguageID = C.TT_MS_LANGID_SPANISH_UNITED_STATES
	MicrosoftLangFinnishFinland              LanguageID = C.TT_MS_LANGID_FINNISH_FINLAND
	MicrosoftLangFrenchFrance                LanguageID = C.TT_MS_LANGID_FRENCH_FRANCE
	MicrosoftLangFrenchBelgium               LanguageID = C.TT_MS_LANGID_FRENCH_BELGIUM
	MicrosoftLangFrenchCanada                LanguageID = C.TT_MS_LANGID_FRENCH_CANADA
	MicrosoftLangFrenchSwitzerland           LanguageID = C.TT_MS_LANGID_FRENCH_SWITZERLAND
	MicrosoftLangFrenchLuxembourg            LanguageID = C.TT_MS_LANGID_FRENCH_LUXEMBOURG
	MicrosoftLangFrenchMonaco                LanguageID = C.TT_MS_LANGID_FRENCH_MONACO
	MicrosoftLangHebrewIsrael                LanguageID = C.TT_MS_LANGID_HEBREW_ISRAEL
	MicrosoftLangHungarianHungary            LanguageID = C.TT_MS_LANGID_HUNGARIAN_HUNGARY
	MicrosoftLangIcelandicIceland            LanguageID = C.TT_MS_LANGID_ICELANDIC_ICELAND
	MicrosoftLangItalianItaly                LanguageID = C.TT_MS_LANGID_ITALIAN_ITALY
	MicrosoftLangItalianSwitzerland          LanguageID = C.TT_MS_LANGID_ITALIAN_SWITZERLAND
	MicrosoftLangJapaneseJapan               LanguageID = C.TT_MS_LANGID_JAPANESE_JAPAN
	MicrosoftLangKoreanKorea                 LanguageID = C.TT_MS_LANGID_KOREAN_KOREA
	MicrosoftLangDutchNetherlands            LanguageID = C.TT_MS_LANGID_DUTCH_NETHERLANDS
	MicrosoftLangDutchBelgium                LanguageID = C.TT_MS_LANGID_DUTCH_BELGIUM
	MicrosoftLangNorwegianNorwayBokmal       LanguageID = C.TT_MS_LANGID_NORWEGIAN_NORWAY_BOKMAL
	MicrosoftLangNorwegianNorwayNynorsk      LanguageID = C.TT_MS_LANGID_NORWEGIAN_NORWAY_NYNORSK
	MicrosoftLangPolishPoland                LanguageID = C.TT_MS_LANGID_POLISH_POLAND
	MicrosoftLangPortugueseBrazil            LanguageID = C.TT_MS_LANGID_PORTUGUESE_BRAZIL
	MicrosoftLangPortuguesePortugal          LanguageID = C.TT_MS_LANGID_PORTUGUESE_PORTUGAL
	MicrosoftLangRomanshSwitzerland          LanguageID = C.TT_MS_LANGID_ROMANSH_SWITZERLAND
	MicrosoftLangRomanianRomania             LanguageID = C.TT_MS_LANGID_ROMANIAN_ROMANIA
	MicrosoftLangRussianRussia               LanguageID = C.TT_MS_LANGID_RUSSIAN_RUSSIA
	MicrosoftLangCroatianCroatia             LanguageID = C.TT_MS_LANGID_CROATIAN_CROATIA
	MicrosoftLangSerbianSerbiaLatin          LanguageID = C.TT_MS_LANGID_SERBIAN_SERBIA_LATIN
	MicrosoftLangSerbianSerbiaCyrillic       LanguageID = C.TT_MS_LANGID_SERBIAN_SERBIA_CYRILLIC
	MicrosoftLangCroatianBosniaHerzegovina   LanguageID = C.TT_MS_LANGID_CROATIAN_BOSNIA_HERZEGOVINA
	MicrosoftLangBosnianBosniaHerzegovina    LanguageID = C.TT_MS_LANGID_BOSNIAN_BOSNIA_HERZEGOVINA
	MicrosoftLangSerbianBosniaHerzLatin      LanguageID = C.TT_MS_LANGID_SERBIAN_BOSNIA_HERZ_LATIN
	MicrosoftLangSerbianBosniaHerzCyrillic   LanguageID = C.TT_MS_LANGID_SERBIAN_BOSNIA_HERZ_CYRILLIC
	MicrosoftLangBosnianBosniaHerzCyrillic   LanguageID = C.TT_MS_LANGID_BOSNIAN_BOSNIA_HERZ_CYRILLIC
	MicrosoftLangSlovakSlovakia              LanguageID = C.TT_MS_LANGID_SLOVAK_SLOVAKIA
	MicrosoftLangAlbanianAlbania             LanguageID = C.TT_MS_LANGID_ALBANIAN_ALBANIA
	MicrosoftLangSwedishSweden               LanguageID = C.TT_MS_LANGID_SWEDISH_SWEDEN
	MicrosoftLangSwedishFinland              LanguageID = C.TT_MS_LANGID_SWEDISH_FINLAND
	MicrosoftLangThaiThailand                LanguageID = C.TT_MS_LANGID_THAI_THAILAND
	MicrosoftLangTurkishTurkey               LanguageID = C.TT_MS_LANGID_TURKISH_TURKEY
	MicrosoftLangUrduPakistan                LanguageID = C.TT_MS_LANGID_URDU_PAKISTAN
	MicrosoftLangIndonesianIndonesia         LanguageID = C.TT_MS_LANGID_INDONESIAN_INDONESIA
	MicrosoftLangUkrainianUkraine            LanguageID = C.TT_MS_LANGID_UKRAINIAN_UKRAINE
	MicrosoftLangBelarusianBelarus           LanguageID = C.TT_MS_LANGID_BELARUSIAN_BELARUS
	MicrosoftLangSlovenianSlovenia           LanguageID = C.TT_MS_LANGID_SLOVENIAN_SLOVENIA
	MicrosoftLangEstonianEstonia             LanguageID = C.TT_MS_LANGID_ESTONIAN_ESTONIA
	MicrosoftLangLatvianLatvia               LanguageID = C.TT_MS_LANGID_LATVIAN_LATVIA
	MicrosoftLangLithuanianLithuania         LanguageID = C.TT_MS_LANGID_LITHUANIAN_LITHUANIA
	MicrosoftLangTajikTajikistan             LanguageID = C.TT_MS_LANGID_TAJIK_TAJIKISTAN
	MicrosoftLangVietnameseVietNam           LanguageID = C.TT_MS_LANGID_VIETNAMESE_VIET_NAM
	MicrosoftLangArmenianArmenia             LanguageID = C.TT_MS_LANGID_ARMENIAN_ARMENIA
	MicrosoftLangAzeriAzerbaijanLatin        LanguageID = C.TT_MS_LANGID_AZERI_AZERBAIJAN_LATIN
	MicrosoftLangAzeriAzerbaijanCyrillic     LanguageID = C.TT_MS_LANGID_AZERI_AZERBAIJAN_CYRILLIC
	MicrosoftLangBasqueBasque                LanguageID = C.TT_MS_LANGID_BASQUE_BASQUE
	MicrosoftLangUpperSorbianGermany         LanguageID = C.TT_MS_LANGID_UPPER_SORBIAN_GERMANY
	MicrosoftLangLowerSorbianGermany         LanguageID = C.TT_MS_LANGID_LOWER_SORBIAN_GERMANY
	MicrosoftLangMacedonianMacedonia         LanguageID = C.TT_MS_LANGID_MACEDONIAN_MACEDONIA
	MicrosoftLangSetswanaSouthAfrica         LanguageID = C.TT_MS_LANGID_SETSWANA_SOUTH_AFRICA
	MicrosoftLangIsixhosaSouthAfrica         LanguageID = C.TT_MS_LANGID_ISIXHOSA_SOUTH_AFRICA
	MicrosoftLangIsizuluSouthAfrica          LanguageID = C.TT_MS_LANGID_ISIZULU_SOUTH_AFRICA
	MicrosoftLangAfrikaansSouthAfrica        LanguageID = C.TT_MS_LANGID_AFRIKAANS_SOUTH_AFRICA
	MicrosoftLangGeorgianGeorgia             LanguageID = C.TT_MS_LANGID_GEORGIAN_GEORGIA
	MicrosoftLangFaeroeseFaeroeIslands       LanguageID = C.TT_MS_LANGID_FAEROESE_FAEROE_ISLANDS
	MicrosoftLangHindiIndia                  LanguageID = C.TT_MS_LANGID_HINDI_INDIA
	MicrosoftLangMalteseMalta                LanguageID = C.TT_MS_LANGID_MALTESE_MALTA
	MicrosoftLangSamiNorthernNorway          LanguageID = C.TT_MS_LANGID_SAMI_NORTHERN_NORWAY
	MicrosoftLangSamiNorthernSweden          LanguageID = C.TT_MS_LANGID_SAMI_NORTHERN_SWEDEN
	MicrosoftLangSamiNorthernFinland         LanguageID = C.TT_MS_LANGID_SAMI_NORTHERN_FINLAND
	MicrosoftLangSamiLuleNorway              LanguageID = C.TT_MS_LANGID_SAMI_LULE_NORWAY
	MicrosoftLangSamiLuleSweden              LanguageID = C.TT_MS_LANGID_SAMI_LULE_SWEDEN
	MicrosoftLangSamiSouthernNorway          LanguageID = C.TT_MS_LANGID_SAMI_SOUTHERN_NORWAY
	MicrosoftLangSamiSouthernSweden          LanguageID = C.TT_MS_LANGID_SAMI_SOUTHERN_SWEDEN
	MicrosoftLangSamiSkoltFinland            LanguageID = C.TT_MS_LANGID_SAMI_SKOLT_FINLAND
	MicrosoftLangSamiInariFinland            LanguageID = C.TT_MS_LANGID_SAMI_INARI_FINLAND
	MicrosoftLangIrishIreland                LanguageID = C.TT_MS_LANGID_IRISH_IRELAND
	MicrosoftLangMalayMalaysia               LanguageID = C.TT_MS_LANGID_MALAY_MALAYSIA
	MicrosoftLangMalayBruneiDarussalam       LanguageID = C.TT_MS_LANGID_MALAY_BRUNEI_DARUSSALAM
	MicrosoftLangKazakhKazakhstan            LanguageID = C.TT_MS_LANGID_KAZAKH_KAZAKHSTAN
	MicrosoftLangKyrgyzKyrgyzstan            LanguageID = C.TT_MS_LANGID_KYRGYZ_KYRGYZSTAN
	MicrosoftLangKiswahiliKenya              LanguageID = C.TT_MS_LANGID_KISWAHILI_KENYA
	MicrosoftLangTurkmenTurkmenistan         LanguageID = C.TT_MS_LANGID_TURKMEN_TURKMENISTAN
	MicrosoftLangUzbekUzbekistanLatin        LanguageID = C.TT_MS_LANGID_UZBEK_UZBEKISTAN_LATIN
	MicrosoftLangUzbekUzbekistanCyrillic     LanguageID = C.TT_MS_LANGID_UZBEK_UZBEKISTAN_CYRILLIC
	MicrosoftLangTatarRussia                 LanguageID = C.TT_MS_LANGID_TATAR_RUSSIA
	MicrosoftLangBengaliIndia                LanguageID = C.TT_MS_LANGID_BENGALI_INDIA
	MicrosoftLangBengaliBangladesh           LanguageID = C.TT_MS_LANGID_BENGALI_BANGLADESH
	MicrosoftLangPunjabiIndia                LanguageID = C.TT_MS_LANGID_PUNJABI_INDIA
	MicrosoftLangGujaratiIndia               LanguageID = C.TT_MS_LANGID_GUJARATI_INDIA
	MicrosoftLangOdiaIndia                   LanguageID = C.TT_MS_LANGID_ODIA_INDIA
	MicrosoftLangTamilIndia                  LanguageID = C.TT_MS_LANGID_TAMIL_INDIA
	MicrosoftLangTeluguIndia                 LanguageID = C.TT_MS_LANGID_TELUGU_INDIA
	MicrosoftLangKannadaIndia                LanguageID = C.TT_MS_LANGID_KANNADA_INDIA
	MicrosoftLangMalayalamIndia              LanguageID = C.TT_MS_LANGID_MALAYALAM_INDIA
	MicrosoftLangAssameseIndia               LanguageID = C.TT_MS_LANGID_ASSAMESE_INDIA
	MicrosoftLangMarathiIndia                LanguageID = C.TT_MS_LANGID_MARATHI_INDIA
	MicrosoftLangSanskritIndia               LanguageID = C.TT_MS_LANGID_SANSKRIT_INDIA
	MicrosoftLangMongolianMongolia           LanguageID = C.TT_MS_LANGID_MONGOLIAN_MONGOLIA
	MicrosoftLangMongolianPrc                LanguageID = C.TT_MS_LANGID_MONGOLIAN_PRC
	MicrosoftLangTibetanPrc                  LanguageID = C.TT_MS_LANGID_TIBETAN_PRC
	MicrosoftLangWelshUnitedKingdom          LanguageID = C.TT_MS_LANGID_WELSH_UNITED_KINGDOM
	MicrosoftLangKhmerCambodia               LanguageID = C.TT_MS_LANGID_KHMER_CAMBODIA
	MicrosoftLangLaoLaos                     LanguageID = C.TT_MS_LANGID_LAO_LAOS
	MicrosoftLangGalicianGalician            LanguageID = C.TT_MS_LANGID_GALICIAN_GALICIAN
	MicrosoftLangKonkaniIndia                LanguageID = C.TT_MS_LANGID_KONKANI_INDIA
	MicrosoftLangSyriacSyria                 LanguageID = C.TT_MS_LANGID_SYRIAC_SYRIA
	MicrosoftLangSinhalaSriLanka             LanguageID = C.TT_MS_LANGID_SINHALA_SRI_LANKA
	MicrosoftLangInuktitutCanada             LanguageID = C.TT_MS_LANGID_INUKTITUT_CANADA
	MicrosoftLangInuktitutCanadaLatin        LanguageID = C.TT_MS_LANGID_INUKTITUT_CANADA_LATIN
	MicrosoftLangAmharicEthiopia             LanguageID = C.TT_MS_LANGID_AMHARIC_ETHIOPIA
	MicrosoftLangTamazightAlgeria            LanguageID = C.TT_MS_LANGID_TAMAZIGHT_ALGERIA
	MicrosoftLangNepaliNepal                 LanguageID = C.TT_MS_LANGID_NEPALI_NEPAL
	MicrosoftLangFrisianNetherlands          LanguageID = C.TT_MS_LANGID_FRISIAN_NETHERLANDS
	MicrosoftLangPashtoAfghanistan           LanguageID = C.TT_MS_LANGID_PASHTO_AFGHANISTAN
	MicrosoftLangFilipinoPhilippines         LanguageID = C.TT_MS_LANGID_FILIPINO_PHILIPPINES
	MicrosoftLangDhivehiMaldives             LanguageID = C.TT_MS_LANGID_DHIVEHI_MALDIVES
	MicrosoftLangHausaNigeria                LanguageID = C.TT_MS_LANGID_HAUSA_NIGERIA
	MicrosoftLangYorubaNigeria               LanguageID = C.TT_MS_LANGID_YORUBA_NIGERIA
	MicrosoftLangQuechuaBolivia              LanguageID = C.TT_MS_LANGID_QUECHUA_BOLIVIA
	MicrosoftLangQuechuaEcuador              LanguageID = C.TT_MS_LANGID_QUECHUA_ECUADOR
	MicrosoftLangQuechuaPeru                 LanguageID = C.TT_MS_LANGID_QUECHUA_PERU
	MicrosoftLangSesothoSaLeboaSouthAfrica   LanguageID = C.TT_MS_LANGID_SESOTHO_SA_LEBOA_SOUTH_AFRICA
	MicrosoftLangBashkirRussia               LanguageID = C.TT_MS_LANGID_BASHKIR_RUSSIA
	MicrosoftLangLuxembourgishLuxembourg     LanguageID = C.TT_MS_LANGID_LUXEMBOURGISH_LUXEMBOURG
	MicrosoftLangGreenlandicGreenland        LanguageID = C.TT_MS_LANGID_GREENLANDIC_GREENLAND
	MicrosoftLangIgboNigeria                 LanguageID = C.TT_MS_LANGID_IGBO_NIGERIA
	MicrosoftLangYiPrc                       LanguageID = C.TT_MS_LANGID_YI_PRC
	MicrosoftLangMapudungunChile             LanguageID = C.TT_MS_LANGID_MAPUDUNGUN_CHILE
	MicrosoftLangMohawkMohawk                LanguageID = C.TT_MS_LANGID_MOHAWK_MOHAWK
	MicrosoftLangBretonFrance                LanguageID = C.TT_MS_LANGID_BRETON_FRANCE
	MicrosoftLangUighurPrc                   LanguageID = C.TT_MS_LANGID_UIGHUR_PRC
	MicrosoftLangMaoriNewZealand             LanguageID = C.TT_MS_LANGID_MAORI_NEW_ZEALAND
	MicrosoftLangOccitanFrance               LanguageID = C.TT_MS_LANGID_OCCITAN_FRANCE
	MicrosoftLangCorsicanFrance              LanguageID = C.TT_MS_LANGID_CORSICAN_FRANCE
	MicrosoftLangAlsatianFrance              LanguageID = C.TT_MS_LANGID_ALSATIAN_FRANCE
	MicrosoftLangYakutRussia                 LanguageID = C.TT_MS_LANGID_YAKUT_RUSSIA
	MicrosoftLangKicheGuatemala              LanguageID = C.TT_MS_LANGID_KICHE_GUATEMALA
	MicrosoftLangKinyarwandaRwanda           LanguageID = C.TT_MS_LANGID_KINYARWANDA_RWANDA
	MicrosoftLangWolofSenegal                LanguageID = C.TT_MS_LANGID_WOLOF_SENEGAL
	MicrosoftLangDariAfghanistan             LanguageID = C.TT_MS_LANGID_DARI_AFGHANISTAN
)

// NameID is the ‘name’ identifier field in the name records of an SFNT ‘name’ table.
// NameID values are platform independent.
type NameID int

// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_name_id_xxx
const (
	NameIDCopyright            NameID = C.TT_NAME_ID_COPYRIGHT
	NameIDFontFamily           NameID = C.TT_NAME_ID_FONT_FAMILY
	NameIDFontSubfamily        NameID = C.TT_NAME_ID_FONT_SUBFAMILY
	NameIDUniqueID             NameID = C.TT_NAME_ID_UNIQUE_ID
	NameIDFullName             NameID = C.TT_NAME_ID_FULL_NAME
	NameIDVersionString        NameID = C.TT_NAME_ID_VERSION_STRING
	NameIDPsName               NameID = C.TT_NAME_ID_PS_NAME
	NameIDTrademark            NameID = C.TT_NAME_ID_TRADEMARK
	NameIDManufacturer         NameID = C.TT_NAME_ID_MANUFACTURER
	NameIDDesigner             NameID = C.TT_NAME_ID_DESIGNER
	NameIDDescription          NameID = C.TT_NAME_ID_DESCRIPTION
	NameIDVendorURL            NameID = C.TT_NAME_ID_VENDOR_URL
	NameIDDesignerURL          NameID = C.TT_NAME_ID_DESIGNER_URL
	NameIDLicense              NameID = C.TT_NAME_ID_LICENSE
	NameIDLicenseURL           NameID = C.TT_NAME_ID_LICENSE_URL
	NameIDTypographicFamily    NameID = C.TT_NAME_ID_TYPOGRAPHIC_FAMILY
	NameIDTypographicSubfamily NameID = C.TT_NAME_ID_TYPOGRAPHIC_SUBFAMILY
	NameIDMacFullName          NameID = C.TT_NAME_ID_MAC_FULL_NAME
	NameIDSampleText           NameID = C.TT_NAME_ID_SAMPLE_TEXT
	NameIDCidFindfontName      NameID = C.TT_NAME_ID_CID_FINDFONT_NAME
	NameIDWwsFamily            NameID = C.TT_NAME_ID_WWS_FAMILY
	NameIDWwsSubfamily         NameID = C.TT_NAME_ID_WWS_SUBFAMILY
	NameIDLightBackground      NameID = C.TT_NAME_ID_LIGHT_BACKGROUND
	NameIDDarkBackground       NameID = C.TT_NAME_ID_DARK_BACKGROUND
	NameIDVariationsPrefix     NameID = C.TT_NAME_ID_VARIATIONS_PREFIX
)

// UCRMask is a bit mask value for the UlUnicodeRangeX fields in an SFNT ‘OS/2’ table.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_ucr_xxx
type UCRMask uint32

// UlUnicodeRange1
const (
	// Bit 0   Basic Latin                                     U+0020-U+007E
	UCRBasicLatin UCRMask = C.TT_UCR_BASIC_LATIN
	// Bit 1   C1 Controls and Latin-1 Supplement              U+0080-U+00FF
	UCRLatin1Supplement UCRMask = C.TT_UCR_LATIN1_SUPPLEMENT
	// Bit 2   Latin Extended-A                                U+0100-U+017F
	UCRLatinExtendedA UCRMask = C.TT_UCR_LATIN_EXTENDED_A
	// Bit 3   Latin Extended-B                                U+0180-U+024F
	UCRLatinExtendedB UCRMask = C.TT_UCR_LATIN_EXTENDED_B
	// Bit 4   IPA Extensions                                  U+0250-U+02AF
	//         Phonetic Extensions                             U+1D00-U+1D7F
	//         Phonetic Extensions Supplement                  U+1D80-U+1DBF
	UCRIpaExtensions UCRMask = C.TT_UCR_IPA_EXTENSIONS
	// Bit 5   Spacing Modifier Letters                        U+02B0-U+02FF
	//         Modifier Tone Letters                           U+A700-U+A71F
	UCRSpacingModifier UCRMask = C.TT_UCR_SPACING_MODIFIER
	// Bit 6   Combining Diacritical Marks                     U+0300-U+036F
	//         Combining Diacritical Marks Supplement          U+1DC0-U+1DFF
	UCRCombiningDiacriticalMarks UCRMask = C.TT_UCR_COMBINING_DIACRITICAL_MARKS
	// Bit 7   Greek and Coptic                                U+0370-U+03FF
	UCRGreek UCRMask = C.TT_UCR_GREEK
	// Bit 8   Coptic                                          U+2C80-U+2CFF
	UCRCoptic UCRMask = C.TT_UCR_COPTIC
	// Bit 9   Cyrillic                                        U+0400-U+04FF
	//         Cyrillic Supplement                             U+0500-U+052F
	//         Cyrillic Extended-A                             U+2DE0-U+2DFF
	//         Cyrillic Extended-B                             U+A640-U+A69F
	UCRCyrillic UCRMask = C.TT_UCR_CYRILLIC
	// Bit 10  Armenian                                        U+0530-U+058F
	UCRArmenian UCRMask = C.TT_UCR_ARMENIAN
	// Bit 11  Hebrew                                          U+0590-U+05FF
	UCRHebrew UCRMask = C.TT_UCR_HEBREW
	// Bit 12  Vai                                             U+A500-U+A63F
	UCRVai UCRMask = C.TT_UCR_VAI
	// Bit 13  Arabic                                          U+0600-U+06FF
	//         Arabic Supplement                               U+0750-U+077F
	UCRArabic UCRMask = C.TT_UCR_ARABIC
	// Bit 14  NKo                                             U+07C0-U+07FF
	UCRNko UCRMask = C.TT_UCR_NKO
	// Bit 15  Devanagari                                      U+0900-U+097F
	UCRDevanagari UCRMask = C.TT_UCR_DEVANAGARI
	// Bit 16  Bengali                                         U+0980-U+09FF
	UCRBengali UCRMask = C.TT_UCR_BENGALI
	// Bit 17  Gurmukhi                                        U+0A00-U+0A7F
	UCRGurmukhi UCRMask = C.TT_UCR_GURMUKHI
	// Bit 18  Gujarati                                        U+0A80-U+0AFF
	UCRGujarati UCRMask = C.TT_UCR_GUJARATI
	// Bit 19  Oriya                                           U+0B00-U+0B7F
	UCROriya UCRMask = C.TT_UCR_ORIYA
	// Bit 20  Tamil                                           U+0B80-U+0BFF
	UCRTamil UCRMask = C.TT_UCR_TAMIL
	// Bit 21  Telugu                                          U+0C00-U+0C7F
	UCRTelugu UCRMask = C.TT_UCR_TELUGU
	// Bit 22  Kannada                                         U+0C80-U+0CFF
	UCRKannada UCRMask = C.TT_UCR_KANNADA
	// Bit 23  Malayalam                                       U+0D00-U+0D7F
	UCRMalayalam UCRMask = C.TT_UCR_MALAYALAM
	// Bit 24  Thai                                            U+0E00-U+0E7F
	UCRThai UCRMask = C.TT_UCR_THAI
	// Bit 25  Lao                                             U+0E80-U+0EFF
	UCRLao UCRMask = C.TT_UCR_LAO
	// Bit 26  Georgian                                        U+10A0-U+10FF
	//         Georgian Supplement                             U+2D00-U+2D2F
	UCRGeorgian UCRMask = C.TT_UCR_GEORGIAN
	// Bit 27  Balinese                                        U+1B00-U+1B7F
	UCRBalinese UCRMask = C.TT_UCR_BALINESE
	// Bit 28  Hangul Jamo                                     U+1100-U+11FF
	UCRHangulJamo UCRMask = C.TT_UCR_HANGUL_JAMO
	// Bit 29  Latin Extended Additional                       U+1E00-U+1EFF
	//         Latin Extended-C                                U+2C60-U+2C7F
	//         Latin Extended-D                                U+A720-U+A7FF
	UCRLatinExtendedAdditional UCRMask = C.TT_UCR_LATIN_EXTENDED_ADDITIONAL
	// Bit 30  Greek Extended                                  U+1F00-U+1FFF
	UCRGreekExtended UCRMask = C.TT_UCR_GREEK_EXTENDED
	// Bit 31  General Punctuation                             U+2000-U+206F
	//         Supplemental Punctuation                        U+2E00-U+2E7F
	UCRGeneralPunctuation UCRMask = C.my_explicit_uint32_TT_UCR_GENERAL_PUNCTUATION
)

// UlUnicodeRange2
const (
	// Bit 32  Superscripts And Subscripts                     U+2070-U+209F
	UCRSuperscriptsSubscripts UCRMask = C.TT_UCR_SUPERSCRIPTS_SUBSCRIPTS
	// Bit 33  Currency Symbols                                U+20A0-U+20CF
	UCRCurrencySymbols UCRMask = C.TT_UCR_CURRENCY_SYMBOLS
	// Bit 34  Combining Diacritical Marks For Symbols         U+20D0-U+20FF
	UCRCombiningDiacriticalMarksSymb UCRMask = C.TT_UCR_COMBINING_DIACRITICAL_MARKS_SYMB
	// Bit 35  Letterlike Symbols                              U+2100-U+214F
	UCRLetterlikeSymbols UCRMask = C.TT_UCR_LETTERLIKE_SYMBOLS
	// Bit 36  Number Forms                                    U+2150-U+218F
	UCRNumberForms UCRMask = C.TT_UCR_NUMBER_FORMS
	// Bit 37  Arrows                                          U+2190-U+21FF
	//         Supplemental Arrows-A                           U+27F0-U+27FF
	//         Supplemental Arrows-B                           U+2900-U+297F
	//         Miscellaneous Symbols and Arrows                U+2B00-U+2BFF
	UCRArrows UCRMask = C.TT_UCR_ARROWS
	// Bit 38  Mathematical Operators                          U+2200-U+22FF
	//         Supplemental Mathematical Operators             U+2A00-U+2AFF
	//         Miscellaneous Mathematical Symbols-A            U+27C0-U+27EF
	//         Miscellaneous Mathematical Symbols-B            U+2980-U+29FF
	UCRMathematicalOperators UCRMask = C.TT_UCR_MATHEMATICAL_OPERATORS
	// Bit 39  Miscellaneous Technical                         U+2300-U+23FF
	UCRMiscellaneousTechnical UCRMask = C.TT_UCR_MISCELLANEOUS_TECHNICAL
	// Bit 40  Control Pictures                                U+2400-U+243F
	UCRControlPictures UCRMask = C.TT_UCR_CONTROL_PICTURES
	// Bit 41  Optical Character Recognition                   U+2440-U+245F
	UCROcr UCRMask = C.TT_UCR_OCR
	// Bit 42  Enclosed Alphanumerics                          U+2460-U+24FF
	UCREnclosedAlphanumerics UCRMask = C.TT_UCR_ENCLOSED_ALPHANUMERICS
	// Bit 43  Box Drawing                                     U+2500-U+257F
	UCRBoxDrawing UCRMask = C.TT_UCR_BOX_DRAWING
	// Bit 44  Block Elements                                  U+2580-U+259F
	UCRBlockElements UCRMask = C.TT_UCR_BLOCK_ELEMENTS
	// Bit 45  Geometric Shapes                                U+25A0-U+25FF
	UCRGeometricShapes UCRMask = C.TT_UCR_GEOMETRIC_SHAPES
	// Bit 46  Miscellaneous Symbols                           U+2600-U+26FF
	UCRMiscellaneousSymbols UCRMask = C.TT_UCR_MISCELLANEOUS_SYMBOLS
	// Bit 47  Dingbats                                        U+2700-U+27BF
	UCRDingbats UCRMask = C.TT_UCR_DINGBATS
	// Bit 48  CJK Symbols and Punctuation                     U+3000-U+303F
	UCRCjkSymbols UCRMask = C.TT_UCR_CJK_SYMBOLS
	// Bit 49  Hiragana                                        U+3040-U+309F
	UCRHiragana UCRMask = C.TT_UCR_HIRAGANA
	// Bit 50  Katakana                                        U+30A0-U+30FF
	//         Katakana Phonetic Extensions                    U+31F0-U+31FF
	UCRKatakana UCRMask = C.TT_UCR_KATAKANA
	// Bit 51  Bopomofo                                        U+3100-U+312F
	//         Bopomofo Extended                               U+31A0-U+31BF
	UCRBopomofo UCRMask = C.TT_UCR_BOPOMOFO
	// Bit 52  Hangul Compatibility Jamo                       U+3130-U+318F
	UCRHangulCompatibilityJamo UCRMask = C.TT_UCR_HANGUL_COMPATIBILITY_JAMO
	// Bit 53  Phags-Pa                                        U+A840-U+A87F
	UCRCjkMisc UCRMask = C.TT_UCR_CJK_MISC
	// Bit 54  Enclosed CJK Letters and Months                 U+3200-U+32FF
	UCREnclosedCjkLettersMonths UCRMask = C.TT_UCR_ENCLOSED_CJK_LETTERS_MONTHS
	// Bit 55  CJK Compatibility                               U+3300-U+33FF
	UCRCjkCompatibility UCRMask = C.TT_UCR_CJK_COMPATIBILITY
	// Bit 56  Hangul Syllables                                U+AC00-U+D7A3
	UCRHangul UCRMask = C.TT_UCR_HANGUL
	// Bit 57  High Surrogates                                 U+D800-U+DB7F
	//         High Private Use Surrogates                     U+DB80-U+DBFF
	//         Low Surrogates                                  U+DC00-U+DFFF
	// According to OpenType specs v.1.3+, setting bit 57 implies that there is
	// at least one codepoint beyond the Basic Multilingual Plane that is
	// supported by this font. So it really means >=           U+10000.
	UCRSurrogates UCRMask = C.TT_UCR_SURROGATES
	UCRNonPlane0  UCRMask = C.TT_UCR_NON_PLANE_0
	// Bit 58  Phoenician                                      U+10900-U+1091F
	UCRPhoenician UCRMask = C.TT_UCR_PHOENICIAN
	// Bit 59  CJK Unified Ideographs                          U+4E00-U+9FFF
	//         CJK Radicals Supplement                         U+2E80-U+2EFF
	//         Kangxi Radicals                                 U+2F00-U+2FDF
	//         Ideographic Description Characters              U+2FF0-U+2FFF
	//         CJK Unified Ideographs Extension A              U+3400-U+4DB5
	//         CJK Unified Ideographs Extension B              U+20000-U+2A6DF
	//         Kanbun                                          U+3190-U+319F
	UCRCjkUnifiedIdeographs UCRMask = C.TT_UCR_CJK_UNIFIED_IDEOGRAPHS
	// Bit 60  Private Use                                     U+E000-U+F8FF
	UCRPrivateUse UCRMask = C.TT_UCR_PRIVATE_USE
	// Bit 61  CJK Strokes                                     U+31C0-U+31EF
	//         CJK Compatibility Ideographs                    U+F900-U+FAFF
	//         CJK Compatibility Ideographs Supplement         U+2F800-U+2FA1F
	UCRCjkCompatibilityIdeographs UCRMask = C.TT_UCR_CJK_COMPATIBILITY_IDEOGRAPHS
	// Bit 62  Alphabetic Presentation Forms                   U+FB00-U+FB4F
	UCRAlphabeticPresentationForms UCRMask = C.TT_UCR_ALPHABETIC_PRESENTATION_FORMS
	// Bit 63  Arabic Presentation Forms-A                     U+FB50-U+FDFF
	UCRArabicPresentationFormsA UCRMask = C.my_explicit_uint32_TT_UCR_ARABIC_PRESENTATION_FORMS_A
)

// UlUnicodeRange3
const (
	// Bit 64  Combining Half Marks                            U+FE20-U+FE2F
	UCRCombiningHalfMarks UCRMask = C.TT_UCR_COMBINING_HALF_MARKS
	// Bit 65  Vertical forms                                  U+FE10-U+FE1F
	//         CJK Compatibility Forms                         U+FE30-U+FE4F
	UCRCjkCompatibilityForms UCRMask = C.TT_UCR_CJK_COMPATIBILITY_FORMS
	// Bit 66  Small Form Variants                             U+FE50-U+FE6F
	UCRSmallFormVariants UCRMask = C.TT_UCR_SMALL_FORM_VARIANTS
	// Bit 67  Arabic Presentation Forms-B                     U+FE70-U+FEFE
	UCRArabicPresentationFormsB UCRMask = C.TT_UCR_ARABIC_PRESENTATION_FORMS_B
	// Bit 68  Halfwidth and Fullwidth Forms                   U+FF00-U+FFEF
	UCRHalfwidthFullwidthForms UCRMask = C.TT_UCR_HALFWIDTH_FULLWIDTH_FORMS
	// Bit 69  Specials                                        U+FFF0-U+FFFD
	UCRSpecials UCRMask = C.TT_UCR_SPECIALS
	// Bit 70  Tibetan                                         U+0F00-U+0FFF
	UCRTibetan UCRMask = C.TT_UCR_TIBETAN
	// Bit 71  Syriac                                          U+0700-U+074F
	UCRSyriac UCRMask = C.TT_UCR_SYRIAC
	// Bit 72  Thaana                                          U+0780-U+07BF
	UCRThaana UCRMask = C.TT_UCR_THAANA
	// Bit 73  Sinhala                                         U+0D80-U+0DFF
	UCRSinhala UCRMask = C.TT_UCR_SINHALA
	// Bit 74  Myanmar                                         U+1000-U+109F
	UCRMyanmar UCRMask = C.TT_UCR_MYANMAR
	// Bit 75  Ethiopic                                        U+1200-U+137F
	//         Ethiopic Supplement                             U+1380-U+139F
	//         Ethiopic Extended                               U+2D80-U+2DDF
	UCREthiopic UCRMask = C.TT_UCR_ETHIOPIC
	// Bit 76  Cherokee                                        U+13A0-U+13FF
	UCRCherokee UCRMask = C.TT_UCR_CHEROKEE
	// Bit 77  Unified Canadian Aboriginal Syllabics           U+1400-U+167F
	UCRCanadianAboriginalSyllabics UCRMask = C.TT_UCR_CANADIAN_ABORIGINAL_SYLLABICS
	// Bit 78  Ogham                                           U+1680-U+169F
	UCROgham UCRMask = C.TT_UCR_OGHAM
	// Bit 79  Runic                                           U+16A0-U+16FF
	UCRRunic UCRMask = C.TT_UCR_RUNIC
	// Bit 80  Khmer                                           U+1780-U+17FF
	//         Khmer Symbols                                   U+19E0-U+19FF
	UCRKhmer UCRMask = C.TT_UCR_KHMER
	// Bit 81  Mongolian                                       U+1800-U+18AF
	UCRMongolian UCRMask = C.TT_UCR_MONGOLIAN
	// Bit 82  Braille Patterns                                U+2800-U+28FF
	UCRBraille UCRMask = C.TT_UCR_BRAILLE
	// Bit 83  Yi Syllables                                    U+A000-U+A48F
	//         Yi Radicals                                     U+A490-U+A4CF
	UCRYi UCRMask = C.TT_UCR_YI
	// Bit 84  Tagalog                                         U+1700-U+171F
	//         Hanunoo                                         U+1720-U+173F
	//         Buhid                                           U+1740-U+175F
	//         Tagbanwa                                        U+1760-U+177F
	UCRPhilippine UCRMask = C.TT_UCR_PHILIPPINE
	// Bit 85  Old Italic                                      U+10300-U+1032F
	UCROldItalic UCRMask = C.TT_UCR_OLD_ITALIC
	// Bit 86  Gothic                                          U+10330-U+1034F
	UCRGothic UCRMask = C.TT_UCR_GOTHIC
	// Bit 87  Deseret                                         U+10400-U+1044F
	UCRDeseret UCRMask = C.TT_UCR_DESERET
	// Bit 88  Byzantine Musical Symbols                       U+1D000-U+1D0FF
	//         Musical Symbols                                 U+1D100-U+1D1FF
	//         Ancient Greek Musical Notation                  U+1D200-U+1D24F
	UCRMusicalSymbols UCRMask = C.TT_UCR_MUSICAL_SYMBOLS
	// Bit 89  Mathematical Alphanumeric Symbols               U+1D400-U+1D7FF
	UCRMathAlphanumericSymbols UCRMask = C.TT_UCR_MATH_ALPHANUMERIC_SYMBOLS
	// Bit 90  Private Use (plane 15)                          U+F0000-U+FFFFD
	//         Private Use (plane 16)                          U+100000-U+10FFFD
	UCRPrivateUseSupplementary UCRMask = C.TT_UCR_PRIVATE_USE_SUPPLEMENTARY
	// Bit 91  Variation Selectors                             U+FE00-U+FE0F
	//         Variation Selectors Supplement                  U+E0100-U+E01EF
	UCRVariationSelectors UCRMask = C.TT_UCR_VARIATION_SELECTORS
	// Bit 92  Tags                                            U+E0000-U+E007F
	UCRTags UCRMask = C.TT_UCR_TAGS
	// Bit 93  Limbu                                           U+1900-U+194F
	UCRLimbu UCRMask = C.TT_UCR_LIMBU
	// Bit 94  Tai Le                                          U+1950-U+197F
	UCRTaiLe UCRMask = C.TT_UCR_TAI_LE
	// Bit 95  New Tai Lue                                     U+1980-U+19DF
	UCRNewTaiLue UCRMask = C.my_explicit_uint32_TT_UCR_NEW_TAI_LUE
)

// UlUnicodeRange4
const (
	// Bit 96  Buginese                                        U+1A00-U+1A1F
	UCRBuginese UCRMask = C.TT_UCR_BUGINESE
	// Bit 97  Glagolitic                                      U+2C00-U+2C5F
	UCRGlagolitic UCRMask = C.TT_UCR_GLAGOLITIC
	// Bit 98  Tifinagh                                        U+2D30-U+2D7F
	UCRTifinagh UCRMask = C.TT_UCR_TIFINAGH
	// Bit 99  Yijing Hexagram Symbols                         U+4DC0-U+4DFF
	UCRYijing UCRMask = C.TT_UCR_YIJING
	// Bit 100 Syloti Nagri                                    U+A800-U+A82F
	UCRSylotiNagri UCRMask = C.TT_UCR_SYLOTI_NAGRI
	// Bit 101 Linear B Syllabary                              U+10000-U+1007F
	//         Linear B Ideograms                              U+10080-U+100FF
	//         Aegean Numbers                                  U+10100-U+1013F
	UCRLinearB UCRMask = C.TT_UCR_LINEAR_B
	// Bit 102 Ancient Greek Numbers                           U+10140-U+1018F
	UCRAncientGreekNumbers UCRMask = C.TT_UCR_ANCIENT_GREEK_NUMBERS
	// Bit 103 Ugaritic                                        U+10380-U+1039F
	UCRUgaritic UCRMask = C.TT_UCR_UGARITIC
	// Bit 104 Old Persian                                     U+103A0-U+103DF
	UCROldPersian UCRMask = C.TT_UCR_OLD_PERSIAN
	// Bit 105 Shavian                                         U+10450-U+1047F
	UCRShavian UCRMask = C.TT_UCR_SHAVIAN
	// Bit 106 Osmanya                                         U+10480-U+104AF
	UCROsmanya UCRMask = C.TT_UCR_OSMANYA
	// Bit 107 Cypriot Syllabary                               U+10800-U+1083F
	UCRCypriotSyllabary UCRMask = C.TT_UCR_CYPRIOT_SYLLABARY
	// Bit 108 Kharoshthi                                      U+10A00-U+10A5F
	UCRKharoshthi UCRMask = C.TT_UCR_KHAROSHTHI
	// Bit 109 Tai Xuan Jing Symbols                           U+1D300-U+1D35F
	UCRTaiXuanJing UCRMask = C.TT_UCR_TAI_XUAN_JING
	// Bit 110 Cuneiform                                       U+12000-U+123FF
	//         Cuneiform Numbers and Punctuation               U+12400-U+1247F
	UCRCuneiform UCRMask = C.TT_UCR_CUNEIFORM
	// Bit 111 Counting Rod Numerals                           U+1D360-U+1D37F
	UCRCountingRodNumerals UCRMask = C.TT_UCR_COUNTING_ROD_NUMERALS
	// Bit 112 Sundanese                                       U+1B80-U+1BBF
	UCRSundanese UCRMask = C.TT_UCR_SUNDANESE
	// Bit 113 Lepcha                                          U+1C00-U+1C4F
	UCRLepcha UCRMask = C.TT_UCR_LEPCHA
	// Bit 114 Ol Chiki                                        U+1C50-U+1C7F
	UCROlChiki UCRMask = C.TT_UCR_OL_CHIKI
	// Bit 115 Saurashtra                                      U+A880-U+A8DF
	UCRSaurashtra UCRMask = C.TT_UCR_SAURASHTRA
	// Bit 116 Kayah Li                                        U+A900-U+A92F
	UCRKayahLi UCRMask = C.TT_UCR_KAYAH_LI
	// Bit 117 Rejang                                          U+A930-U+A95F
	UCRRejang UCRMask = C.TT_UCR_REJANG
	// Bit 118 Cham                                            U+AA00-U+AA5F
	UCRCham UCRMask = C.TT_UCR_CHAM
	// Bit 119 Ancient Symbols                                 U+10190-U+101CF
	UCRAncientSymbols UCRMask = C.TT_UCR_ANCIENT_SYMBOLS
	// Bit 120 Phaistos Disc                                   U+101D0-U+101FF
	UCRPhaistosDisc UCRMask = C.TT_UCR_PHAISTOS_DISC
	// Bit 121 Carian                                          U+102A0-U+102DF
	//         Lycian                                          U+10280-U+1029F
	//         Lydian                                          U+10920-U+1093F
	UCROldAnatolian UCRMask = C.TT_UCR_OLD_ANATOLIAN
	// Bit 122 Domino Tiles                                    U+1F030-U+1F09F
	//         Mahjong Tiles                                   U+1F000-U+1F02F
	UCRGameTiles UCRMask = C.TT_UCR_GAME_TILES
)
