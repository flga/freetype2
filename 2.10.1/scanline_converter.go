package freetype2

// #include <stdlib.h>
// #include <ft2build.h>
// #include FT_IMAGE_H
import (
	"C"
)

// Raster is an opaque handle (pointer) to a raster object. Each object can be
// used independently to convert an outline into a bitmap or pixmap.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-raster.html#ft_raster
type Raster struct {
	ptr C.FT_Raster
}

// Span models a single span of gray pixels when rendering an anti-aliased bitmap.
//
// Span is used by the span drawing callback type SpanFunc (TODO) that takes the y
// coordinate of the span as a parameter.
//
// The coverage value is always between 0 and 255. If you want less gray values,
// the callback function has to reduce them.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-raster.html#ft_span
type Span struct {
	// The span's horizontal start position.
	X int16
	// The span's length in pixels.
	Len uint16
	// The span color/coverage, ranging from 0 (background) to 255 (foreground).
	Coverage uint8
}

// SpanFunc is used as a call-back by the anti-aliased renderer in order to let
// client applications draw themselves the gray pixel spans on each scan line.
//
// This callback allows client applications to directly render the gray spans of
// the anti-aliased bitmap to any kind of surfaces.
//
// This can be used to write anti-aliased outlines directly to a given
// background bitmap, and even perform translucency.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-raster.html#ft_spanfunc
type SpanFunc func(upwardY int, spans []Span)

// RasterParams holds the parameters used by a raster's render function, passed
// as an argument to Outline.Render.
//
// An anti-aliased glyph bitmap is drawn if the RasterFlagAA bit flag is set in
// the flags field, otherwise a monochrome bitmap is generated.
//
// If the RasterFlagDirect bit flag is set in flags, the raster will call the
// GraySpans callback to draw gray pixel spans. This allows direct composition
// over a pre-existing bitmap through user-provided callbacks to perform the
// span drawing and composition. Not supported by the monochrome rasterizer.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-raster.html#ft_raster_params
type RasterParams struct {
	// The target bitmap.
	Target *Bitmap
	// The rendering flags.
	Flags RasterFlag
	// The gray span drawing callback.
	GraySpans SpanFunc
	// An optional clipping box. It is only used in direct rendering mode. Note
	// that coordinates here should be expressed in integer pixels (and not in
	// 26.6 fixed-point units).
	ClipBox BBox
}

// RasterFlag is a list of bit flag constants as used in the flags field
// of RasterParams.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-raster.html#ft_raster_flag_xxx
type RasterFlag uint

const (
	// RasterFlagDefault this value is 0.
	RasterFlagDefault RasterFlag = C.FT_RASTER_FLAG_DEFAULT
	// RasterFlagAA indicates that an anti-aliased glyph image should be
	// generated. Otherwise, it will be monochrome (1-bit).
	RasterFlagAA RasterFlag = C.FT_RASTER_FLAG_AA
	// RasterFlagDirect indicates direct rendering. In this mode, client
	// applications must provide their own span callback. This lets them
	// directly draw or compose over an existing bitmap. If this bit is not set,
	// the target pixmap's buffer must be zeroed before rendering and the output
	// will be clipped to its size.
	//
	// Direct rendering is only possible with anti-aliased glyphs.
	RasterFlagDirect RasterFlag = C.FT_RASTER_FLAG_DIRECT
	// RasterFlagClip is only used in direct rendering mode. If set, the output
	// will be clipped to a box specified in the ClipBox field of RasterParams.
	// Otherwise, the ClipBox is effectively set to the bounding box and all
	// spans are generated.
	RasterFlagClip RasterFlag = C.FT_RASTER_FLAG_CLIP
)
