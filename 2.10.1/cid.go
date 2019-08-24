package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_CID_H
import "C"

// CIDRegistryOrderingSupplement retrieves the Registry/Ordering/Supplement
// triple (also known as the "R/O/S") from a CID-keyed font.
//
// This function only works with CID faces, returning an error otherwise.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-cid_fonts.html#ft_get_cid_registry_ordering_supplement
func (f *Face) CIDRegistryOrderingSupplement() (registry, ordering string, supplement int, err error) {
	if f == nil || f.ptr == nil {
		return "", "", 0, ErrInvalidArgument
	}

	var cregistry, cordering *C.char
	var csupplement C.int
	if err := getErr(C.FT_Get_CID_Registry_Ordering_Supplement(f.ptr, &cregistry, &cordering, &csupplement)); err != nil {
		return "", "", 0, err
	}

	return C.GoString(cregistry), C.GoString(cordering), int(csupplement), nil
}

// IsInternallyCIDKeyed retrieves the type of the input face, CID keyed or not.
// In contrast to the FT_IS_CID_KEYED macro this function returns successfully
// also for CID-keyed fonts in an SFNT wrapper.
//
// This function only works with CID faces and OpenType fonts, returning false
// otherwise.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-cid_fonts.html#ft_get_cid_is_internally_cid_keyed
func (f *Face) IsInternallyCIDKeyed() bool {
	if f == nil || f.ptr == nil {
		return false
	}

	var isCid C.FT_Bool
	if err := getErr(C.FT_Get_CID_Is_Internally_CID_Keyed(f.ptr, &isCid)); err != nil {
		return false
	}

	if isCid == 1 {
		return true
	}

	return false
}

// CIDFromGlyphIndex retrieves the CID of the input glyph index.
//
// This function only works with CID faces and OpenType fonts, returning an error otherwise.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-cid_fonts.html#ft_get_cid_from_glyph_index
func (f *Face) CIDFromGlyphIndex(idx GlyphIndex) (uint, error) {
	if f == nil || f.ptr == nil {
		return 0, ErrInvalidArgument
	}

	var cid C.uint
	if err := getErr(C.FT_Get_CID_From_Glyph_Index(f.ptr, C.uint(idx), &cid)); err != nil {
		return 0, err
	}

	return uint(cid), nil
}
