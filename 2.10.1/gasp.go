package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_GASP_H
import "C"

// GaspFlag is a list of values and/or bit-flags returned by the Face.Gasp method.
//
// The bit-flags GaspDoGridfit and GaspDoGray are to be used for standard font
// rasterization only. Independently of that, GaspSymmetricSmoothing and
// GaspSymmetricGridfit are to be used if ClearType is enabled (and
// GaspDoGridfit and GaspDoGray are consequently ignored).
//
// ‘ClearType’ is Microsoft's implementation of LCD rendering, partly protected
// by patents.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-gasp_table.html#ft_gasp_xxx
type GaspFlag uint

const (
	// GaspFlagDoGridfit grid-fitting and hinting should be performed at the
	// specified ppem. This really means TrueType bytecode interpretation.
	// If this bit is not set, no hinting gets applied.
	GaspFlagDoGridfit GaspFlag = C.FT_GASP_DO_GRIDFIT
	// GaspFlagDoGray anti-aliased rendering should be performed at the specified
	// ppem. If not set, do monochrome rendering.
	GaspFlagDoGray GaspFlag = C.FT_GASP_DO_GRAY
	// GaspFlagSymmetricGridfit grid-fitting must be used with ClearType's
	// symmetric smoothing.
	GaspFlagSymmetricGridfit GaspFlag = C.FT_GASP_SYMMETRIC_GRIDFIT
	// GaspFlagSymmetricSmoothing if set, smoothing along multiple axes must be
	// used with ClearType.
	GaspFlagSymmetricSmoothing GaspFlag = C.FT_GASP_SYMMETRIC_SMOOTHING
)

// GaspFlags returns the rasterizer behaviour flags from the font's ‘gasp’ table
// corresponding to a given character pixel size for a TrueType or OpenType font.
//
// If the font does not have a gasp table, it returns 0 and false.
//
// If you want to use the MM functionality of OpenType variation fonts (i.e.,
// using SetVarDesignCoordinates and friends), call this function after setting
// an instance since the return values can change.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-gasp_table.html#ft_get_gasp
func (f *Face) GaspFlags(verticalPPem int) (flags GaspFlag, ok bool) {
	if f == nil || f.ptr == nil {
		return 0, false
	}

	v := C.FT_Get_Gasp(f.ptr, C.uint(verticalPPem))
	if v == C.FT_GASP_NO_TABLE {
		return 0, false
	}

	return GaspFlag(v), true
}
