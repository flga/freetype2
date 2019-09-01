package freetype2

// #include <stdlib.h>
// #include <ft2build.h>
// #include FT_ADVANCES_H
import "C"

// Advance retrieves the advance value of a given glyph outline.
//
// Flags are used to determine what kind of advances you need.
//
// If scaling is performed (based on the value of flags), the advance value is
// in 16.16 format. Otherwise, it is in font units.
//
// If LoadVerticalLayout is set, this is the vertical advance corresponding to a
// vertical layout. Otherwise, it is the horizontal advance in a horizontal layout.
//
// This function may fail if you use AdvanceFlagFastOnly and if the corresponding
// font backend doesn't have a quick way to retrieve the advances.
//
// A scaled advance is returned in 16.16 format but isn't transformed by the
// affine transformation specified by SetTransform.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-quick_advance.html#ft_get_advance
func (f *Face) Advance(idx GlyphIndex, flags LoadFlag) (Pos, error) {
	if f == nil || f.ptr == nil {
		return 0, ErrInvalidFaceHandle
	}

	var ret C.FT_Fixed
	if err := getErr(C.FT_Get_Advance(f.ptr, C.FT_UInt(idx), C.FT_Int32(flags), &ret)); err != nil {
		return 0, err
	}
	return Pos(ret), nil
}

// Advances retrieves the advance values of several glyph outlines.
//
// Flags are used to determine what kind of advances you need.
//
// If scaling is performed (based on the value of flags), the advance values are
// in 16.16 format. Otherwise, they are in font units.
//
// If LoadVerticalLayout is set, these are the vertical advances corresponding to
// a vertical layout. Otherwise, they are the horizontal advances in a horizontal
// layout.
//
// This function may fail if you use AdvanceFlagFastOnly and if the corresponding
// font backend doesn't have a quick way to retrieve the advances.
//
// Scaled advancea are returned in 16.16 format but aren't transformed by the
// affine transformation specified by SetTransform.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-quick_advance.html#ft_get_advances
func (f *Face) Advances(startIdx GlyphIndex, numGlyphs int, flags LoadFlag) ([]Pos, error) {
	if f == nil || f.ptr == nil {
		return nil, ErrInvalidFaceHandle
	}

	if numGlyphs <= 0 {
		return nil, nil
	}

	block := C.calloc(C.size_t(numGlyphs), C.sizeof_FT_Fixed)
	defer free(block)
	if err := getErr(C.FT_Get_Advances(f.ptr, C.FT_UInt(startIdx), C.FT_UInt(numGlyphs), C.FT_Int32(flags), (*C.FT_Fixed)(block))); err != nil {
		return nil, err
	}

	ptr := (*[(1<<31 - 1) / C.sizeof_FT_Fixed]C.FT_Fixed)(block)[:numGlyphs:numGlyphs]
	ret := make([]Pos, numGlyphs)
	for i := range ret {
		ret[i] = Pos(ptr[i])
	}
	return ret, nil
}

// AdvanceFlagFastOnly is a bit-flag to be OR-ed with the flags parameter of the
// GetAdvance and GetAdvances face methods.
//
// If set, it indicates that you want these functions to fail if the corresponding
// hinting mode or font driver doesn't allow for very quick advance computation.
//
// Typically, glyphs that are either unscaled, unhinted, bitmapped, or
// light-hinted can have their advance width computed very quickly.
//
// Normal and bytecode hinted modes that require loading, scaling, and hinting of
// the glyph outline, are extremely slow by comparison.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-quick_advance.html#ft_advance_flag_fast_only
const AdvanceFlagFastOnly LoadFlag = C.FT_ADVANCE_FLAG_FAST_ONLY
