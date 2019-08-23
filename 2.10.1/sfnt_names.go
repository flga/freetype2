package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_SFNT_NAMES_H
import "C"

import (
	"bytes"
	"encoding/binary"
	"errors"
	"unicode/utf16"
	"unsafe"

	"github.com/flga/freetype2/2.10.1/truetype"
)

// ErrUnableToDecode occurs when the length of SfntLangTag is not even.
var ErrUnableToDecode = errors.New("unable to decode sfnt lang tag")

// SfntName models an SFNT ‘name’ table entry.
//
// Please refer to the TrueType or OpenType specification for more details.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_sfntname
type SfntName struct {
	// The platform ID for Name.
	PlatformID truetype.PlatformID
	// The encoding ID for Name.
	EncodingID truetype.EncodingID
	// The language ID for Name.
	// Registered OpenType values for LanguageID are always smaller than 0x8000;
	// values equal or larger than 0x8000 usually indicate a language tag string
	// (introduced in OpenType version 1.6). Use function SfntLangTag with
	// LanguageID as its argument to retrieve the associated language tag.
	LanguageID truetype.LanguageID
	// An identifier for Name
	NameID truetype.NameID
	// The ‘name’ string. Note that its format differs depending on the
	// (platform,encoding) pair, being either a string of bytes (without a
	// terminating NULL byte) or containing UTF-16BE entities. TODO: can this be a string?
	Name string
}

// https://github.com/fonttools/fonttools/blob/63fb3fb881440b636267dc43082aab134e40f8f9/Lib/fontTools/ttLib/tables/_n_a_m_e.py#L325
func (s SfntName) isUnicode() bool {
	if s.PlatformID == truetype.PlatformAppleUnicode {
		return true
	}

	if s.PlatformID == truetype.PlatformMicrosoft {
		switch s.EncodingID {
		case truetype.MicrosoftEncodingSymbolCs,
			truetype.MicrosoftEncodingUnicodeCs,
			truetype.MicrosoftEncodingUCS4:
			return true

		}

		return false
	}

	return false
}

// SfntNameCount returns the number of strings in the ‘name’ table.
//
// It returns 0 if the config macro TT_CONFIG_OPTION_SFNT_NAMES is not defined.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_get_sfnt_name_count
func (f *Face) SfntNameCount() int {
	if f == nil || f.ptr == nil {
		return 0
	}

	return int(C.FT_Get_Sfnt_Name_Count(f.ptr))
}

// SfntName retrieves a string of the SFNT ‘name’ table for a given index.
//
// Use SfntNameCount to get the total number of available ‘name’ table entries,
// then do a loop until you get the right platform, encoding, and name ID.
//
// ‘name’ table format 1 entries can use language tags also, see SfntLangTag.
//
// It returns an error if the config macro TT_CONFIG_OPTION_SFNT_NAMES is not
// defined.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_get_sfnt_name
func (f *Face) SfntName(idx int) (SfntName, error) {
	if f == nil || f.ptr == nil {
		return SfntName{}, ErrInvalidArgument
	}

	if idx < 0 {
		return SfntName{}, ErrInvalidArgument
	}

	var aname C.FT_SfntName
	if err := getErr(C.FT_Get_Sfnt_Name(f.ptr, C.uint(idx), &aname)); err != nil {
		return SfntName{}, err
	}

	ret := SfntName{
		PlatformID: truetype.PlatformID(aname.platform_id),
		EncodingID: truetype.EncodingID(aname.encoding_id),
		LanguageID: truetype.LanguageID(aname.language_id),
		NameID:     truetype.NameID(aname.name_id),
	}

	name := C.GoBytes(unsafe.Pointer(aname.string), C.int(aname.string_len))
	if ret.isUnicode() {
		dec, err := decodeUTF16BE(name)
		if err != nil {
			ret.Name = string(name)
		}
		ret.Name = string(dec)
	} else {
		ret.Name = string(name)
	}

	return ret, nil
}

// SfntLangTag is the language tag string, encoded in UTF8.
//
// Please refer to the TrueType or OpenType specification for more details.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_sfntlangtag
type SfntLangTag string

// SfntLangTag retrieves the language tag associated with a language ID of an
// SFNT ‘name’ table entry.
//
// Only ‘name’ table format 1 supports language tags. For format 0 tables, it
// returns ErrInvalidTable. For invalid format 1 language ID values,
// ErrInvalidArgument is returned.
//
// It returns an error if the config macro TT_CONFIG_OPTION_SFNT_NAMES is not
// defined.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_get_sfnt_langtag
func (f *Face) SfntLangTag(id truetype.LanguageID) (SfntLangTag, error) {
	if f == nil || f.ptr == nil {
		return "", ErrInvalidArgument
	}

	var alangTag C.FT_SfntLangTag
	if err := getErr(C.FT_Get_Sfnt_LangTag(f.ptr, C.uint(id), &alangTag)); err != nil {
		return "", err
	}

	dec, err := decodeUTF16BE(C.GoBytes(unsafe.Pointer(alangTag.string), C.int(alangTag.string_len)))
	if err != nil {
		return "", err
	}

	return SfntLangTag(dec), nil
}

func decodeUTF16BE(data []byte) ([]rune, error) {
	if len(data)%2 != 0 {
		return nil, ErrUnableToDecode
	}

	str := make([]uint16, len(data)/2)
	if err := binary.Read(bytes.NewReader(data), binary.BigEndian, str); err != nil {
		return nil, err
	}

	return utf16.Decode(str), nil
}
