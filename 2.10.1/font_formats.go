package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_FONT_FORMATS_H
import "C"

// FontFormat is an enumeration of the possible font formats.
type FontFormat int

// FontFormat values
const (
	FontFormatNone FontFormat = iota
	FontFormatTrueType
	FontFormatType1
	FontFormatBDF
	FontFormatPCF
	FontFormatType42
	FontFormatCIDType1
	FontFormatCFF
	FontFormatPFR
	FontFormatWindowsFNT
)

func (f FontFormat) String() string {
	switch f {
	case FontFormatTrueType:
		return "TrueType"
	case FontFormatType1:
		return "Type 1"
	case FontFormatBDF:
		return "BDF"
	case FontFormatPCF:
		return "PCF"
	case FontFormatType42:
		return "Type 42"
	case FontFormatCIDType1:
		return "CID Type 1"
	case FontFormatCFF:
		return "CFF"
	case FontFormatPFR:
		return "PFR"
	case FontFormatWindowsFNT:
		return "Windows FNT"
	default:
		return ""
	}
}

// FontFormat returns the format of the face.
//
// The return value is suitable to be used as an X11 FONT_PROPERTY.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-font_formats.html#ft_get_font_format
func (f *Face) FontFormat() FontFormat {
	if f == nil || f.ptr == nil {
		return FontFormatNone
	}

	switch C.GoString(C.FT_Get_Font_Format(f.ptr)) {
	case "TrueType":
		return FontFormatTrueType
	case "Type 1":
		return FontFormatType1
	case "BDF":
		return FontFormatBDF
	case "PCF":
		return FontFormatPCF
	case "Type 42":
		return FontFormatType42
	case "CID Type 1":
		return FontFormatCIDType1
	case "CFF":
		return FontFormatCFF
	case "PFR":
		return FontFormatPFR
	case "Windows FNT":
		return FontFormatWindowsFNT
	default:
		return FontFormatNone
	}
}
