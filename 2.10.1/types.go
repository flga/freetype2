package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"
import (
	"fmt"

	"golang.org/x/image/math/fixed"
)

// Encoding is an enumeration to specify character sets supported by charmaps.
// Used in the SelectCharmap API function.
//
// Despite the name, this enumeration lists specific character repertories (i.e., charsets), and not text encoding
// methods (e.g., UTF-8, UTF-16, etc.).
//
// Other encodings might be defined in the future.
//
// By default, FreeType enables a Unicode charmap and tags it with EncodingUnicode when it is either provided or can be
// generated from PostScript glyph name dictionaries in the font file. All other encodings are considered legacy and
// tagged only if explicitly defined in the font file. Otherwise, EncodingNone is used.
//
// EncodingNone is set by the BDF and PCF drivers if the charmap is neither Unicode nor ISO-8859-1 (otherwise it is set
// to EncodingUnicode).
// Use TODO: FT_Get_BDF_Charset_ID to find out which encoding is really present. If, for example, the TODO: cs_registry
// field is ‘KOI8’ and the  TODO: cs_encoding field is ‘R’, the font is encoded in KOI8-R.
//
// EncodingNone is always set (with a single exception) by the winfonts driver.
// Use TODO: FT_Get_WinFNT_Header and examine the charset field of the TODO: FT_WinFNT_HeaderRec structure to find out
// which encoding is really present. For example, TODO: FT_WinFNT_ID_CP1251 (204) means Windows code page 1251
// (for Russian).
//
// EncodingNone is set if TODO: platform_id is TODO: TT_PLATFORM_MACINTOSH and TODO: encoding_id is not
// TODO: TT_MAC_ID_ROMAN (otherwise it is set to EncodingAppleRoman).
//
// If TODO: platform_id is TODO: TT_PLATFORM_MACINTOSH, use the function TODO: FT_Get_CMap_Language_ID to query the Mac
// language ID that may be needed to be able to distinguish Apple encoding variants.
// See https://www.unicode.org/Public/MAPPINGS/VENDORS/APPLE/Readme.txt to get an idea how to do that. Basically, if the
// language ID is 0, don't use it, otherwise subtract 1 from the language ID. Then examine TODO: encoding_id.
// If, for example, TODO: encoding_id is TODO: TT_MAC_ID_ROMAN and the language ID (minus 1) is TODO: TT_MAC_LANGID_GREEK,
// it is the Greek encoding, not Roman.
// TODO: TT_MAC_ID_ARABIC with TODO: TT_MAC_LANGID_FARSI means the Farsi variant
// the Arabic encoding.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_encoding
type Encoding int

const (
	// EncodingNone is reserved for all formats except BDF, PCF, and Windows FNT
	EncodingNone Encoding = C.FT_ENCODING_NONE

	// EncodingMsSymbol is the Microsoft Symbol encoding, used to encode mathematical symbols and wingdings. This
	// encoding uses character codes from the PUA (Private Unicode Area) in the range U+F020-U+F0FF.
	//
	// For more information, see https://www.microsoft.com/typography/otspec/recom.htm#non-standard-symbol-fonts,
	// http://www.kostis.net/charsets/symbol.htm, and http://www.kostis.net/charsets/wingding.htm
	EncodingMsSymbol Encoding = C.FT_ENCODING_MS_SYMBOL

	// EncodingUnicode is the Unicode character set. This value covers all versions of the Unicode repertoire, including
	// ASCII and Latin-1. Most fonts include a Unicode charmap, but not all of them.
	//
	// For example, if you want to access Unicode value U+1F028 (and the font contains it), use value 0x1F028 as the
	// input value for GetCharIndex.
	EncodingUnicode Encoding = C.FT_ENCODING_UNICODE

	// EncodingSJIS is the Shift JIS encoding for Japanese.
	//
	// More info at https://en.wikipedia.org/wiki/Shift_JIS.
	EncodingSJIS Encoding = C.FT_ENCODING_SJIS

	// EncodingPRC corresponds to encoding systems mainly for Simplified Chinese as used in People's Republic of China
	// (PRC). The encoding layout is based on GB 2312 and its supersets GBK and GB 18030.
	EncodingPRC Encoding = C.FT_ENCODING_PRC

	// EncodingBig5 corresponds to an encoding system for Traditional Chinese as used in Taiwan and Hong Kong.
	EncodingBig5 Encoding = C.FT_ENCODING_BIG5

	// EncodingWansung corresponds to the Korean encoding system known as Extended Wansung (MS Windows code page 949).
	//
	// For more information see‘https://www.unicode.org/Public/MAPPINGS/VENDORS/MICSFT/WindowsBestFit/bestfit949.txt
	EncodingWansung Encoding = C.FT_ENCODING_WANSUNG

	// EncodingJohab is the Korean standard character set (KS C 5601-1992), which corresponds to MS Windows code page
	// 1361. This character set includes all possible Hangul character combinations.
	EncodingJohab Encoding = C.FT_ENCODING_JOHAB

	// EncodingAdobeStandard is the Adobe Standard encoding, as found in Type 1, CFF, and OpenType/CFF fonts.
	// It is limited to 256 character codes.
	EncodingAdobeStandard Encoding = C.FT_ENCODING_ADOBE_STANDARD

	// EncodingAdobeExpert is the Adobe Expert encoding, as found in Type 1, CFF, and OpenType/CFF fonts.
	// It is limited to 256 character codes.
	EncodingAdobeExpert Encoding = C.FT_ENCODING_ADOBE_EXPERT

	// EncodingAdobeCustom corresponds to a custom encoding, as found in Type 1, CFF, and OpenType/CFF fonts.
	// It is limited to 256 character codes.
	EncodingAdobeCustom Encoding = C.FT_ENCODING_ADOBE_CUSTOM

	// EncodingAdobeLatin1 corresponds to a Latin-1 encoding as defined in a Type 1 PostScript font.
	// It is limited to 256 character codes.
	EncodingAdobeLatin1 Encoding = C.FT_ENCODING_ADOBE_LATIN_1

	// EncodingAppleRoman is the Apple roman encoding. Many TrueType and OpenType fonts contain a charmap for this 8-bit
	// encoding, since older versions of Mac OS are able to use it.
	EncodingAppleRoman Encoding = C.FT_ENCODING_APPLE_ROMAN
)

// Int26_6 is a signed 26.6 fixed-point type used for vectorial pixel coordinates.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_f26dot6
type Int26_6 = fixed.Int26_6

// Int16_16 is a signed 16.16 fixed-point number.
//
// The integer part ranges from -32768 to 32767, inclusive. The fractional part has 16 bits of precision.
// For example, the number one-and-a-quarter is Int16_16(1<<6 + 1<<14).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_fixed
type Int16_16 int32

func (x Int16_16) String() string {
	const shift, mask = 16, 1<<16 - 1
	if x >= 0 {
		return fmt.Sprintf("%d:%d", int32(x>>shift), int32(x&mask))
	}
	x = -x
	if x >= 0 {
		return fmt.Sprintf("-%d:%d", int32(x>>shift), int32(x&mask))
	}
	return "-32768:0" // The minimum value is -(1<<15).
}

// Floor returns the greatest integer value less than or equal to x.
//
// Its return type is int, not Int16_16.
func (x Int16_16) Floor() int { return int((x) >> 16) }

// Round returns the nearest integer value to x. Ties are rounded up.
//
// Its return type is int, not Int16_16.
func (x Int16_16) Round() int { return int((x + 1<<15) >> 16) }

// Ceil returns the least integer value greater than or equal to x.
//
// Its return type is int, not Int16_16.
func (x Int16_16) Ceil() int { return int((x + 1<<16 - 1) >> 16) }

// Mul returns x*y in 16.16 fixed-point arithmetic.
func (x Int16_16) Mul(y Int16_16) Int16_16 {
	return Int16_16((int64(x)*int64(y) + 1<<15) >> 16)
}

// F32 converts the underlying value to float32.
func (x Int16_16) F32() float32 {
	return float32(x) / float32(1<<16)
}

// F64 converts the underlying value to float64.
func (x Int16_16) F64() float64 {
	return float64(x) / float64(1<<16)
}

// Pos is used to store vectorial coordinates. Depending on the context, these can represent distances in integer
// font units, or 16.16, or 26.6 fixed-point pixel coordinates.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pos
type Pos int32

// BitmapSize models the metrics of a bitmap strike (i.e., a set of glyphs for a given point size and resolution) in a
// bitmap font.
//
// NOTE:
// Windows FNT: The nominal size given in a FNT font is not reliable. If the driver finds it incorrect, it sets Size to
// some calculated values, and XPpem and YPpem to the pixel width and height given in the font, respectively.
//
// NOTE:
// TrueType embedded bitmaps: size, width, and height values are not contained in the bitmap strike itself. They are
// computed from the global font parameters.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_bitmap_size
type BitmapSize struct {
	// The vertical distance, in pixels, between two consecutive baselines.
	// It is always positive.
	Height int
	// The average width, in pixels, of all glyphs in the strike.
	Width int
	// The nominal size of the strike. This field is not very useful.
	Size Int26_6
	// The horizontal ppem (nominal width).
	XPpem Int26_6
	// The vertical ppem (nominal height).
	YPpem Int26_6
}

// CharMap is used to translate character codes in a given encoding into glyph indexes for its parent's face.
// Some font formats may provide several charmaps per font.
//
// Each face object owns zero or more charmaps, but only one of them can be ‘active’, providing the data used by
// GetCharIndex or LoadChar.
//
// NOTE:
// When a new face is created (either through NewFace or OpenFace), the library looks for a Unicode
// charmap within the list and automatically activates it. If there is no Unicode charmap, FreeType doesn't set an
// ‘active’ charmap.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_charmap
type CharMap struct {
	// Format of the CharMap
	Format int

	// Language id
	Language LanguageID

	// An Encoding tag identifying the charmap. Use this with SelectCharmap.
	Encoding Encoding

	// An ID number describing the platform for the following encoding ID.
	// This comes directly from the TrueType specification and gets emulated for
	// other formats.
	PlatformID PlatformID

	// A platform-specific encoding number. This also comes from the TrueType
	// specification and gets emulated similarly.
	EncodingID EncodingID

	// The index into Face.CharMaps
	index int

	// not user created
	ours bool
}

// BBox holds an outline's bounding box, i.e., the coordinates of its extrema in the horizontal and vertical directions.
//
// The bounding box is specified with the coordinates of the lower left and the upper right corner. In PostScript, those
// values are often called (llx,lly) and (urx,ury), respectively.
//
// If YMin is negative, this value gives the glyph's descender. Otherwise, the glyph doesn't descend below the baseline.
// Similarly, if YMax is positive, this value gives the glyph's ascender.
//
// XMin gives the horizontal distance from the glyph's origin to the left edge of the glyph's bounding box. If XMin is
// negative, the glyph extends to the left of the origin.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bbox
type BBox struct {
	// The horizontal minimum (left-most).
	XMin Pos
	// The vertical minimum (bottom-most).
	YMin Pos
	// The horizontal maximum (right-most).
	XMax Pos
	// The vertical maximum (top-most).
	YMax Pos
}
