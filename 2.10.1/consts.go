package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_TRUETYPE_TABLES_H
import (
	"C"
)

// Encoding is an enumeration to specify character sets supported by charmaps.
// Used in the SelectCharMap API function.
//
// Despite the name, this enumeration lists specific character repertoires (i.e., charsets), and not text encoding
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
// Use GetBDFCharsetID to find out which encoding is really present. If, for example, the csRegistry field is ‘KOI8’ and
// the csEncoding field is ‘R’, the font is encoded in KOI8-R.
//
// EncodingNone is always set (with a single exception) by the winfonts driver.
// Use GetWinFNTHeader and examine the charset field of the WinFNTHeader struct to find out which encoding is really
// present. For example, TODO: WinFNTIDs.CP1251 (204) means Windows code page 1251 (for Russian).
//
// EncodingNone is set if PlatformID is PlatformMacintosh and EncodingID is not MacEncodingRoman (otherwise it is
// set to EncodingAppleRoman).
//
// If PlatformID is PlatformMacintosh, use the function CharMapLanguage to query the Mac language ID that may be needed
// to be able to distinguish Apple encoding variants.
// See https://www.unicode.org/Public/MAPPINGS/VENDORS/APPLE/Readme.txt to get an idea how to do that. Basically, if the
// language ID is 0, don't use it, otherwise subtract 1 from the language ID. Then examine EncodingID.
// If, for example, EncodingID is MacEncodingRoman and the language ID (minus 1) is MacLangGreek, it is the
// Greek encoding, not Roman.
// MacEncodingArabic with MacLangFarsi means the Farsi variant the Arabic encoding.
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

func (e Encoding) String() string {
	switch e {
	case EncodingNone:
		return "None"
	case EncodingMsSymbol:
		return "MsSymbol"
	case EncodingUnicode:
		return "Unicode"
	case EncodingSJIS:
		return "SJIS"
	case EncodingPRC:
		return "PRC"
	case EncodingBig5:
		return "Big5"
	case EncodingWansung:
		return "Wansung"
	case EncodingJohab:
		return "Johab"
	case EncodingAdobeStandard:
		return "AdobeStandard"
	case EncodingAdobeExpert:
		return "AdobeExpert"
	case EncodingAdobeCustom:
		return "AdobeCustom"
	case EncodingAdobeLatin1:
		return "AdobeLatin1"
	case EncodingAppleRoman:
		return "AppleRoman"
	default:
		return "Unknown"
	}
}

// PixelMode is an enumeration type used to describe the format of pixels in a given bitmap. Note that additional
// formats may be added in the future.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pixel_mode
type PixelMode int

const (
	// PixelModeMono is a monochrome bitmap, using 1 bit per pixel.
	// Note that pixels are stored in most-significant order (MSB), which means that the left-most pixel in a byte has
	// value 128.
	PixelModeMono PixelMode = C.FT_PIXEL_MODE_MONO

	// PixelModeGray is an 8-bit bitmap, generally used to represent anti-aliased glyph images. Each pixel is stored in
	// one byte. Note that the number of ‘gray’ levels is stored in the NumGrays field of the Bitmap type
	// (it generally is 256).
	PixelModeGray PixelMode = C.FT_PIXEL_MODE_GRAY

	// PixelModeGray2 is a 2-bit per pixel bitmap, used to represent embedded anti-aliased bitmaps in font files
	// according to the OpenType spec.
	// We haven't found a single font using this format, however.
	PixelModeGray2 PixelMode = C.FT_PIXEL_MODE_GRAY2

	// PixelModeGray4 is a 4-bit per pixel bitmap, representing embedded anti-aliased bitmaps in font files according to
	// the OpenType spec.
	// We haven't found a single font using this format, however.
	PixelModeGray4 PixelMode = C.FT_PIXEL_MODE_GRAY4

	// PixelModeLCD is an 8-bit bitmap, representing RGB or BGR decimated glyph images used for display on LCD displays;
	// the bitmap is three times wider than the original glyph image. See also RenderModeLCD.
	PixelModeLCD PixelMode = C.FT_PIXEL_MODE_LCD

	// PixelModeLCDV is an 8-bit bitmap, representing RGB or BGR decimated glyph images used for display on rotated LCD
	// displays; the bitmap is three times taller than the original glyph image. See also RenderModeLCDV.
	PixelModeLCDV PixelMode = C.FT_PIXEL_MODE_LCD_V

	// PixelModeBGRA is an image with four 8-bit channels per pixel, representing a color image (such as emoticons) with
	// alpha channel.
	// For each pixel, the format is BGRA, which means, the blue channel comes first in memory. The color channels are
	// pre-multiplied and in the sRGB colorspace. For example, full red at half-translucent opacity will be represented
	// as ‘00,00,80,80’, not ‘00,00,FF,80’. See also  LoadColor.
	PixelModeBGRA PixelMode = C.FT_PIXEL_MODE_BGRA
)

func (p PixelMode) String() string {
	switch p {
	case PixelModeMono:
		return "Mono"
	case PixelModeGray:
		return "Gray"
	case PixelModeGray2:
		return "Gray2"
	case PixelModeGray4:
		return "Gray4"
	case PixelModeLCD:
		return "LCD"
	case PixelModeLCDV:
		return "LCDV"
	case PixelModeBGRA:
		return "BGRA"
	default:
		return "Unknown"
	}
}

// BitsPerPixel reports the number of bits needed for a single pixel.
func (p PixelMode) BitsPerPixel() int {
	switch p {
	case PixelModeMono:
		return 1
	case PixelModeGray:
		return 8
	case PixelModeGray2:
		return 2
	case PixelModeGray4:
		return 4
	case PixelModeLCD:
		return 8
	case PixelModeLCDV:
		return 8
	case PixelModeBGRA:
		return 8 * 4
	default:
		return 0
	}
}

// GlyphFormat is an enumeration type used to describe the format of a given glyph image. Note that this version of
// FreeType only supports two image formats, even though future font drivers will be able to register their own format.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_glyph_format
type GlyphFormat int

const (
	// GlyphFormatComposite the glyph image is a composite of several other images. This format is only used with
	// LoadNoRecurse, and is used to report compound glyphs (like accented characters).
	GlyphFormatComposite GlyphFormat = C.FT_GLYPH_FORMAT_COMPOSITE

	// GlyphFormatBitmap the glyph image is a bitmap, and can be described as a Bitmap. You generally need to access the
	// bitmap field of the GlyphSlot structure to read it.
	GlyphFormatBitmap GlyphFormat = C.FT_GLYPH_FORMAT_BITMAP

	// GlyphFormatOutline the glyph image is a vectorial outline made of line segments and Bezier arcs; it can be
	// described as an Outline; you generally want to access the outline field of the GlyphSlot structure to read it.
	GlyphFormatOutline GlyphFormat = C.FT_GLYPH_FORMAT_OUTLINE

	// GlyphFormatPlotter the glyph image is a vectorial path with no inside and outside contours. Some Type 1 fonts,
	// like those in the Hershey family, contain glyphs in this format. These are described as Outline, but FreeType
	// isn't currently capable of rendering them correctly.
	GlyphFormatPlotter GlyphFormat = C.FT_GLYPH_FORMAT_PLOTTER
)

func (g GlyphFormat) String() string {
	switch g {
	case GlyphFormatComposite:
		return "Composite"
	case GlyphFormatBitmap:
		return "Bitmap"
	case GlyphFormatOutline:
		return "Outline"
	case GlyphFormatPlotter:
		return "Plotter"
	default:
		return "Unknown"
	}
}
