package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_TRUETYPE_TABLES_H
import (
	"C"
)
import (
	"github.com/flga/freetype2/2.10.1/fixed"
	"github.com/flga/freetype2/2.10.1/truetype"
)

// Pos is used to store vectorial coordinates. Depending on the context, these can represent distances in integer
// font units, or 16.16, or 26.6 fixed-point pixel coordinates.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pos
type Pos int32

// Vector models a 2D vector
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_vector
type Vector struct {
	X, Y Pos
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

func newBBox(b C.FT_BBox) BBox {
	return BBox{
		XMin: Pos(b.xMin),
		YMin: Pos(b.yMin),
		XMax: Pos(b.xMax),
		YMax: Pos(b.yMax),
	}
}

// Matrix stores a 2x2 matrix. Coefficients are in 16.16 fixed-point format.
//
// The computation performed is:
// x' = x*xx + y*xy
// y' = x*yx + y*yy
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_matrix
type Matrix struct {
	Xx, Xy fixed.Int16_16
	Yx, Yy fixed.Int16_16
}

// FWord is a distance in original font units.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_fword
type FWord int16

// UFWord is an unsigned distance in original font units.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_ufword
type UFWord uint16

// UnitVector stores a 2D vector unit vector.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_unitvector
type UnitVector struct {
	X, Y fixed.Int2_14
}

// Bitmap represents a bitmap or pixmap to the raster.
// Note that we now manage pixmaps of various depths through the PixelMode field.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bitmap
type Bitmap struct {
	// The number of bitmap rows.
	Rows int

	// The number of pixels in bitmap row.
	Width int

	// The pitch's absolute value is the number of bytes taken by one bitmap row,
	// including padding. However, the pitch is positive when the bitmap has a
	// ‘down’ flow, and negative when it has an ‘up’ flow. In all cases, the
	// pitch is an offset to add to a bitmap pointer in order to go down one row.
	//
	// Note that ‘padding’ means the alignment of a bitmap to a byte border,
	// and FreeType functions normally align to the smallest possible integer
	// value.
	//
	// For the B/W rasterizer, pitch is always an even number.
	//
	// To change the pitch of a bitmap (say, to make it a multiple of 4), use
	// Convert().
	Pitch int

	// Pixel data
	Buffer []byte

	// This field is only used with PixelModeGray; it gives the number of gray
	// levels used in the bitmap.
	NumGrays int

	// The pixel mode, i.e., how pixel bits are stored. See PixelMode for
	// possible values.
	PixelMode PixelMode
}

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
	Size fixed.Int26_6
	// The horizontal ppem (nominal width).
	XPpem fixed.Int26_6
	// The vertical ppem (nominal height).
	YPpem fixed.Int26_6
}

func newBitmapSize(b C.FT_Bitmap_Size) BitmapSize {
	return BitmapSize{
		Height: int(b.height),
		Width:  int(b.width),
		Size:   fixed.Int26_6(b.size),
		XPpem:  fixed.Int26_6(b.x_ppem),
		YPpem:  fixed.Int26_6(b.y_ppem),
	}
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
	Language truetype.LanguageID

	// An Encoding tag identifying the charmap. Use this with SelectCharmap.
	Encoding Encoding

	// An ID number describing the platform for the following encoding ID.
	// This comes directly from the TrueType specification and gets emulated for
	// other formats.
	PlatformID truetype.PlatformID

	// A platform-specific encoding number. This also comes from the TrueType
	// specification and gets emulated similarly.
	EncodingID truetype.EncodingID

	// The index into Face.CharMaps
	index int

	// not user created
	valid bool
}

// Index reports the index into the array of character maps within the face to which c belongs.
// If c is not a valid charmap, it will return 0 and false.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_charmap_index
func (c CharMap) Index() (idx int, ok bool) {
	if !c.valid {
		return 0, false
	}

	return c.index, true
}

func newCharMap(c C.FT_CharMap) CharMap {
	if c == nil {
		return CharMap{}
	}

	return CharMap{
		Format:     int(C.FT_Get_CMap_Format(c)),
		Language:   truetype.LanguageID(C.FT_Get_CMap_Language_ID(c)),
		Encoding:   Encoding(c.encoding),
		PlatformID: truetype.PlatformID(c.platform_id),
		EncodingID: truetype.EncodingID(c.encoding_id),
		index:      int(C.FT_Get_Charmap_Index(c)),
		valid:      true,
	}
}

// Size models a face scaled to a given character size.
//
// A Face has one active Size object that is used by functions like LoadGlyph to determine the scaling transformation
// that in turn is used to load and hint glyphs and metrics.
//
// You can use SetCharSize, SetPixelSizes, RequestSize or even SelectSize to change the content (i.e., the scaling
// values) of the active Size.
//
// You can use NewSize to create additional size objects for a given Face, but they won't be used by other functions
// until you activate it through ActivateSize. Only one size can be activated at any given time per face.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_size
type Size struct {
	SizeMetrics
}

func newSize(s C.FT_Size) Size {
	if s == nil {
		return Size{}
	}

	return Size{
		SizeMetrics{
			XPpem:      int(s.metrics.x_ppem),
			YPpem:      int(s.metrics.y_ppem),
			XScale:     fixed.Int16_16(s.metrics.x_scale),
			YScale:     fixed.Int16_16(s.metrics.y_scale),
			Ascender:   fixed.Int26_6(s.metrics.ascender),
			Descender:  fixed.Int26_6(s.metrics.descender),
			Height:     fixed.Int26_6(s.metrics.height),
			MaxAdvance: fixed.Int26_6(s.metrics.max_advance),
		},
	}
}

// SizeMetrics contains the metrics of a size object.
//
// The scaling values, if relevant, are determined first during a size changing operation. The remaining fields are then
// set by the driver. For scalable formats, they are usually set to scaled values of the corresponding fields in Face.
// Some values like ascender or descender are rounded for historical reasons; more precise values (for outline fonts)
// can be derived by scaling the corresponding Face values manually, with code similar to the following.
//	TODO: scaled_ascender = FT_MulFix( face->ascender, size_metrics->y_scale );
//
// Note that due to glyph hinting and the selected rendering mode these values are usually not exact; consequently, they
// must be treated as unreliable with an error margin of at least one pixel!
//
// Indeed, the only way to get the exact metrics is to render all glyphs. As this would be a definite performance hit,
// it is up to client applications to perform such computations.
//
// SizeMetrics is valid for bitmap fonts also.
//
// TrueType fonts with native bytecode hinting
//
// All applications that handle TrueType fonts with native hinting must be aware that TTFs expect different rounding of
// vertical font dimensions. The application has to cater for this, especially if it wants to rely on a TTF's vertical
// data (for example, to properly align box characters vertically).
//
// Only the application knows in advance that it is going to use native hinting for TTFs! FreeType, on the other hand,
// selects the hinting mode not at the time of creating a Size object but much later, namely while calling LoadGlyph.
//
// Here is some pseudo code that illustrates a possible solution. (TODO)
//	font_format = FT_Get_Font_Format( face );
//
//	if ( !strcmp( font_format, "TrueType" ) &&
//		do_native_bytecode_hinting         )
//	{
//	ascender  = ROUND( FT_MulFix( face->ascender,
//									size_metrics->y_scale ) );
//	descender = ROUND( FT_MulFix( face->descender,
//									size_metrics->y_scale ) );
//	}
//	else
//	{
//	ascender  = size_metrics->ascender;
//	descender = size_metrics->descender;
//	}
//
//	height      = size_metrics->height;
//	max_advance = size_metrics->max_advance;
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_size_metrics
type SizeMetrics struct {
	// The width of the scaled EM square in pixels, hence the term ‘ppem’
	// (pixels per EM). It is also referred to as ‘nominal width’.
	XPpem int

	// The height of the scaled EM square in pixels, hence the term ‘ppem’
	// (pixels per EM). It is also referred to as ‘nominal height’.
	YPpem int

	// A 16.16 fractional scaling value to convert horizontal metrics from font
	// units to 26.6 fractional pixels. Only relevant for scalable font formats.
	XScale fixed.Int16_16

	// A 16.16 fractional scaling value to convert vertical metrics from font
	// units to 26.6 fractional pixels. Only relevant for scalable font formats.
	YScale fixed.Int16_16

	// The ascender in 26.6 fractional pixels, rounded up to an integer value.
	// See Face for the details.
	Ascender fixed.Int26_6

	// The descender in 26.6 fractional pixels, rounded down to an integer value.
	// See Face for the details.
	Descender fixed.Int26_6

	// The height in 26.6 fractional pixels, rounded to an integer value.
	// See Face for the details.
	Height fixed.Int26_6

	// The maximum advance width in 26.6 fractional pixels, rounded to an
	// integer value. See Face for the details.
	MaxAdvance fixed.Int26_6
}

// SizeRequestType is an enumeration of the supported size request types, i.e.,
// what input size (in font units) maps to the requested output size (in pixels,
// as computed from the arguments of SizeRequest).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_size_request_type
type SizeRequestType int

const (
	// SizeRequestTypeNominal the UnitsPerEM method of Face is used to determine both scaling values.
	//
	// This is the standard scaling found in most applications. In particular, use this size request type for TrueType
	// fonts if they provide optical scaling or something similar. Note, however, that UnitsPerEM is a rather abstract
	// value which bears no relation to the actual size of the glyphs in a font.
	SizeRequestTypeNominal SizeRequestType = C.FT_SIZE_REQUEST_TYPE_NOMINAL
	// SizeRequestTypeRealDim the sum of the ascender and (minus of) the descender fields of Face is used to determine
	// both scaling values.
	SizeRequestTypeRealDim SizeRequestType = C.FT_SIZE_REQUEST_TYPE_REAL_DIM
	// SizeRequestTypeBBox the width and height of the BBox field of Face are used to determine the horizontal and
	// vertical scaling value, respectively.
	SizeRequestTypeBBox SizeRequestType = C.FT_SIZE_REQUEST_TYPE_BBOX
	// SizeRequestTypeCell the MaxAdvanceWidth field of Face is used to determine the horizontal scaling value; the
	// vertical scaling valueis determined the same way as SizeRequestTypeRealDim does. Finally, both scaling values are
	// set to the smaller one. This type is useful if you want to specify the font size for, say, a window of a given
	// dimension and 80x24 cells.
	SizeRequestTypeCell SizeRequestType = C.FT_SIZE_REQUEST_TYPE_CELL
	// SizeRequestTypeScales specify the scaling values directly.
	SizeRequestTypeScales SizeRequestType = C.FT_SIZE_REQUEST_TYPE_SCALES
)

func (s SizeRequestType) String() string {
	switch s {
	case SizeRequestTypeNominal:
		return "Nominal"
	case SizeRequestTypeRealDim:
		return "RealDim"
	case SizeRequestTypeBBox:
		return "BBox"
	case SizeRequestTypeCell:
		return "Cell"
	case SizeRequestTypeScales:
		return "Scales"
	default:
		return "Unknown"
	}
}

// SizeRequest models a size request.
//
// If type is SizeRequestTypeScales, width and height are interpreted directly as 16.16 fractional scaling values,
// without any further modification, and both horiResolution and vertResolution are ignored.
//
// NOTE:
// If width is zero, the horizontal scaling value is set equal to the vertical scaling value, and vice versa.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_size_requestrec
type SizeRequest struct {
	// See SizeRequestType
	Type SizeRequestType
	// The desired width, given as a 26.6 fractional point value (with 72pt = 1in).
	Width fixed.Int26_6
	// The desired height, given as a 26.6 fractional point value (with 72pt = 1in).
	Height fixed.Int26_6
	// The horizontal resolution (dpi, i.e., pixels per inch). If set to zero,
	// width is treated as a 26.6 fractional pixel value, which gets internally
	// rounded to an integer.
	HoriResolution uint
	// The vertical resolution (dpi, i.e., pixels per inch). If set to zero, height is treated as a 26.6 fractional pixel value, which gets internally rounded to an integer.
	VertResolution uint
}
