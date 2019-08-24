package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_WINFONTS_H
import "C"

// WinFntID is a list of valid values for the charset byte in WinFNTHeader.
//
// Exact mapping tables for the various ‘cpXXXX’ encodings (except for ‘cp1361’)
// can be found at ‘ftp://ftp.unicode.org/Public/’ in the MAPPINGS/VENDORS/MICSFT/WINDOWS
// subdirectory. ‘cp1361’ is roughly a superset of MAPPINGS/OBSOLETE/EASTASIA/KSC/JOHAB.TXT.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-winfnt_fonts.html#ft_winfnt_id_xxx
type WinFntID uint

const (
	// WinFntIDCP1252 is ANSI encoding. A superset of ISO 8859-1.
	WinFntIDCP1252 WinFntID = C.FT_WinFNT_ID_CP1252
	// WinFntIDDEFAULT is used for font enumeration and font creation as a
	// ‘don't care’ value. Valid font files don't contain this value. When
	// querying for information about the character set of the font that is
	// currently selected into a specified device context, this return value
	// (of the related Windows API) simply denotes failure.
	WinFntIDDEFAULT WinFntID = C.FT_WinFNT_ID_DEFAULT
	// WinFntIDSYMBOL there is no known mapping table available.
	WinFntIDSYMBOL WinFntID = C.FT_WinFNT_ID_SYMBOL
	// WinFntIDMAC Mac Roman encoding.
	WinFntIDMAC WinFntID = C.FT_WinFNT_ID_MAC
	// WinFntIDCP932 is A superset of Japanese Shift-JIS (with minor deviations).
	WinFntIDCP932 WinFntID = C.FT_WinFNT_ID_CP932
	// WinFntIDCP949 is a superset of Korean Hangul KS C 5601-1987 (with
	// different ordering and minor deviations).
	WinFntIDCP949 WinFntID = C.FT_WinFNT_ID_CP949
	// WinFntIDCP1361 is Korean (Johab).
	WinFntIDCP1361 WinFntID = C.FT_WinFNT_ID_CP1361
	// WinFntIDCP936 is a superset of simplified Chinese GB 2312-1980 (with
	// different ordering and minor deviations).
	WinFntIDCP936 WinFntID = C.FT_WinFNT_ID_CP936
	// WinFntIDCP950 is a superset of traditional Chinese Big 5 ETen (with
	// different ordering and minor deviations).
	WinFntIDCP950 WinFntID = C.FT_WinFNT_ID_CP950
	// WinFntIDCP1253 is a superset of Greek ISO 8859-7 (with minor modifications).
	WinFntIDCP1253 WinFntID = C.FT_WinFNT_ID_CP1253
	// WinFntIDCP1254 is a superset of Turkish ISO 8859-9.
	WinFntIDCP1254 WinFntID = C.FT_WinFNT_ID_CP1254
	// WinFntIDCP1258 is for Vietnamese. This encoding doesn't cover all
	// necessary characters.
	WinFntIDCP1258 WinFntID = C.FT_WinFNT_ID_CP1258
	// WinFntIDCP1255 is a superset of Hebrew ISO 8859-8 (with some modifications).
	WinFntIDCP1255 WinFntID = C.FT_WinFNT_ID_CP1255
	// WinFntIDCP1256 is a superset of Arabic ISO 8859-6 (with different ordering).
	WinFntIDCP1256 WinFntID = C.FT_WinFNT_ID_CP1256
	// WinFntIDCP1257 is a superset of Baltic ISO 8859-13 (with some deviations).
	WinFntIDCP1257 WinFntID = C.FT_WinFNT_ID_CP1257
	// WinFntIDCP1251 is a superset of Russian ISO 8859-5 (with different ordering).
	WinFntIDCP1251 WinFntID = C.FT_WinFNT_ID_CP1251
	// WinFntIDCP874 is a superset of Thai TIS 620 and ISO 8859-11.
	WinFntIDCP874 WinFntID = C.FT_WinFNT_ID_CP874
	// WinFntIDCP1250 is a superset of East European ISO 8859-2 (with slightly
	// different ordering).
	WinFntIDCP1250 WinFntID = C.FT_WinFNT_ID_CP1250
	// WinFntIDOEM from Michael Poettgen <michael@poettgen.de>:
	// The ‘Windows Font Mapping’ article says that WinFntIDOEM is used for the
	// charset of vector fonts, like modern.fon, roman.fon, and script.fon on Windows.
	//
	// The ‘CreateFont’ documentation says: The WinFntIDOEM value specifies a
	// character set that is operating-system dependent.
	//
	// The ‘IFIMETRICS’ documentation from the ‘Windows Driver Development Kit’
	// says: This font supports an OEM-specific character set.
	// The OEM character set is system dependent.
	//
	// In general OEM, as opposed to ANSI (i.e., ‘cp1252’), denotes the second
	// default codepage that most international versions of Windows have.
	// It is one of the OEM codepages from https://docs.microsoft.com/en-us/windows/desktop/intl/code-page-identifiers,
	// and is used for the ‘DOS boxes’, to support legacy applications.
	// A German Windows version for example usually uses ANSI codepage 1252 and
	// OEM codepage 850.
	WinFntIDOEM WinFntID = C.FT_WinFNT_ID_OEM
)

// WinFntHeader models Windows FNT Header info.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-winfnt_fonts.html#ft_winfnt_headerrec
type WinFntHeader struct {
	Version              uint16
	FileSize             uint64
	Copyright            [60]byte
	FileType             uint16
	NominalPointSize     uint16
	VerticalResolution   uint16
	HorizontalResolution uint16
	Ascent               uint16
	InternalLeading      uint16
	ExternalLeading      uint16
	Italic               byte
	Underline            byte
	StrikeOut            byte
	Weight               uint16
	Charset              byte
	PixelWidth           uint16
	PixelHeight          uint16
	PitchAndFamily       byte
	AvgWidth             uint16
	MaxWidth             uint16
	FirstChar            byte
	LastChar             byte
	DefaultChar          byte
	BreakChar            byte
	BytesPerRow          uint16
	DeviceOffset         uint64
	FaceNameOffset       uint64
	BitsPointer          uint64
	BitsOffset           uint64
	Reserved             byte
	Flags                uint64
	ASpace               uint16
	BSpace               uint16
	CSpace               uint16
	ColorTableOffset     uint16
	Reserved1            [4]uint64
}

// WinFntHeader retrieves a Windows FNT font info header.
//
// This function only works with Windows FNT faces, returning an error otherwise.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-winfnt_fonts.html#ft_get_winfnt_header
func (f *Face) WinFntHeader() (WinFntHeader, error) {
	if f == nil || f.ptr == nil {
		return WinFntHeader{}, ErrInvalidFaceHandle
	}

	var header C.FT_WinFNT_HeaderRec
	if err := getErr(C.FT_Get_WinFNT_Header(f.ptr, &header)); err != nil {
		return WinFntHeader{}, err
	}

	ret := WinFntHeader{
		Version:              uint16(header.version),
		FileSize:             uint64(header.file_size),
		FileType:             uint16(header.file_type),
		NominalPointSize:     uint16(header.nominal_point_size),
		VerticalResolution:   uint16(header.vertical_resolution),
		HorizontalResolution: uint16(header.horizontal_resolution),
		Ascent:               uint16(header.ascent),
		InternalLeading:      uint16(header.internal_leading),
		ExternalLeading:      uint16(header.external_leading),
		Italic:               byte(header.italic),
		Underline:            byte(header.underline),
		StrikeOut:            byte(header.strike_out),
		Weight:               uint16(header.weight),
		Charset:              byte(header.charset),
		PixelWidth:           uint16(header.pixel_width),
		PixelHeight:          uint16(header.pixel_height),
		PitchAndFamily:       byte(header.pitch_and_family),
		AvgWidth:             uint16(header.avg_width),
		MaxWidth:             uint16(header.max_width),
		FirstChar:            byte(header.first_char),
		LastChar:             byte(header.last_char),
		DefaultChar:          byte(header.default_char),
		BreakChar:            byte(header.break_char),
		BytesPerRow:          uint16(header.bytes_per_row),
		DeviceOffset:         uint64(header.device_offset),
		FaceNameOffset:       uint64(header.face_name_offset),
		BitsPointer:          uint64(header.bits_pointer),
		BitsOffset:           uint64(header.bits_offset),
		Reserved:             byte(header.reserved),
		Flags:                uint64(header.flags),
		ASpace:               uint16(header.A_space),
		BSpace:               uint16(header.B_space),
		CSpace:               uint16(header.C_space),
		ColorTableOffset:     uint16(header.color_table_offset),
	}

	for i, v := range header.copyright {
		ret.Copyright[i] = byte(v)
	}
	for i, v := range header.reserved1 {
		ret.Reserved1[i] = uint64(v)
	}

	return ret, nil
}
