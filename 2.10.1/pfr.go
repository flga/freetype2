package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_PFR_H
import "C"
import (
	"github.com/flga/freetype2/fixed"
)

// PFRMetrics models the outline and metrics resolutions of a given PFR face.
type PFRMetrics struct {
	// This is equivalent to Face.UnitsPerEM() for non-PFR fonts.
	OutlineResolution uint
	// This is equivalent to OutlineResolution for non-PFR fonts.
	MetricsResolution uint
	// A 16.16 fixed-point number used to scale distance expressed in metrics
	// units to device subpixels. This is equivalent to Face.Size.XScale, but
	// for metrics only.
	MetricsXScale fixed.Int16_16
	// Same as MetricsXScale but for the vertical direction.
	MetricsYScale fixed.Int16_16
}

// PFRMetrics returns the outline and metrics resolutions of a given PFR face.
//
// If the input face is not a PFR, this function will return an error. However,
// in all cases, it will return valid values.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-pfr_fonts.html#ft_get_pfr_metrics
func (f *Face) PFRMetrics() (PFRMetrics, error) {
	if f == nil || f.ptr == nil {
		return PFRMetrics{}, ErrInvalidFaceHandle
	}

	var outlineResolution, metricsResolution C.FT_UInt
	var metricsXScale, metricsYScale C.FT_Fixed
	err := getErr(C.FT_Get_PFR_Metrics(
		f.ptr,
		&outlineResolution,
		&metricsResolution,
		&metricsXScale,
		&metricsYScale,
	))

	return PFRMetrics{
		OutlineResolution: uint(outlineResolution),
		MetricsResolution: uint(metricsResolution),
		MetricsXScale:     fixed.Int16_16(metricsXScale),
		MetricsYScale:     fixed.Int16_16(metricsYScale),
	}, err
}

// PFRKerning returns the kerning pair corresponding to two glyphs in a PFR face.
// The distance is expressed in metrics units, unlike the result of Kern.
//
// This function always return distances in original PFR metrics units.
// This is unlike Kern with the KerningUnscaled mode, which always returns
// distances converted to outline units.
//
// You can use the MetricsXScale and MetricsYScale values of PFRMetrics to scale
// these to device subpixels.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-pfr_fonts.html#ft_get_pfr_kerning
func (f *Face) PFRKerning(left, right GlyphIndex) (Vector, error) {
	if f == nil || f.ptr == nil {
		return Vector{}, ErrInvalidFaceHandle
	}

	var vector C.FT_Vector
	if err := getErr(C.FT_Get_PFR_Kerning(f.ptr, C.uint(left), C.uint(right), &vector)); err != nil {
		return Vector{}, err
	}

	return Vector{X: Pos(vector.x), Y: Pos(vector.y)}, nil
}

// PFRAdvance returns a given glyph advance, expressed in original metrics units,
// from a PFR font.
//
// You can use the MetricsXScale or MetricsYScale values of PFRMetrics to convert
// the advance to device subpixels (i.e., 1/64th of pixels).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-pfr_fonts.html#ft_get_pfr_advance
func (f *Face) PFRAdvance(idx GlyphIndex) (Pos, error) {
	if f == nil || f.ptr == nil {
		return 0, ErrInvalidFaceHandle
	}

	var advance C.FT_Pos
	if err := getErr(C.FT_Get_PFR_Advance(f.ptr, C.uint(idx), &advance)); err != nil {
		return 0, err
	}

	return Pos(advance), nil
}
