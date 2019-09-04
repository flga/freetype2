package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_SIZES_H
// #include FT_TRUETYPE_TABLES_H
import (
	"C"
)
import (
	"unsafe"

	"github.com/flga/freetype2/2.10.1/truetype"
	"github.com/flga/freetype2/fixed"
)

// Tag is a typedef for 32-bit tags (as used in the SFNT format).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_tag
type Tag uint32

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

// Vector26_6 models a 2D vector
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_vector
type Vector26_6 struct {
	X, Y fixed.Int26_6
}

// Vector16_16 models a 2D vector
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_vector
type Vector16_16 struct {
	X, Y fixed.Int16_16
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
	ptr         *C.FT_Bitmap `deep:"-"`
	l           *Library     `deep:"-"`
	userCreated bool

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

func (b *Bitmap) reload() {
	if b == nil {
		return
	}

	if b.ptr == nil {
		*b = Bitmap{}
	}

	b.Rows = int(b.ptr.rows)
	b.Width = int(b.ptr.width)
	b.Pitch = int(b.ptr.pitch) //TODO: perhaps split pitch and flow?
	b.NumGrays = int(b.ptr.num_grays)
	b.PixelMode = PixelMode(b.ptr.pixel_mode)

	if b.ptr.buffer != nil {
		pitch := b.Pitch
		if pitch < 0 {
			pitch = -pitch
		}

		length := pitch * b.Rows
		b.Buffer = make([]byte, length)

		buf := (*[(1<<31 - 1) / C.sizeof_uchar]C.uchar)(unsafe.Pointer(b.ptr.buffer))[:length:length]
		bpp := b.PixelMode.BytesPerPixel()
		for i := range b.Buffer {
			if i%pitch/bpp >= b.Width { // set padding bytes explicitly to 0
				b.Buffer[i] = 0
			} else {
				b.Buffer[i] = byte(buf[i])
			}
		}
	} else {
		b.Buffer = nil
	}
}

// BitmapSize models the metrics of a bitmap strike (i.e., a set of glyphs for a given point size and resolution) in a
// bitmap font.
//
// Windows FNT: The nominal size given in a FNT font is not reliable. If the driver finds it incorrect, it sets Size to
// some calculated values, and XPpem and YPpem to the pixel width and height given in the font, respectively.
//
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

// CharMap is used to translate character codes in a given encoding into glyph indexes for its parent's face.
// Some font formats may provide several charmaps per font.
//
// Each face object owns zero or more charmaps, but only one of them can be ‘active’, providing the data used by
// GetCharIndex or LoadChar.
//
// When a new face is created (either through NewFace or OpenFace), the library looks for a Unicode
// charmap within the list and automatically activates it. If there is no Unicode charmap, FreeType doesn't set an
// ‘active’ charmap.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_charmap
type CharMap struct {
	// Format of the CharMap
	Format truetype.CmapFormat

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
		Format:     truetype.CmapFormat(C.FT_Get_CMap_Format(c)),
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
	ptr C.FT_Size `deep:"-"`
	SizeMetrics
}

func newSize(s C.FT_Size) *Size {
	if s == nil {
		return nil
	}

	return &Size{
		ptr: s,
		SizeMetrics: SizeMetrics{
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

// Free discards the size object.
//
// Note that Face.Free() automatically discards all size objects allocated with
// Face.NewSize().
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-sizes_management.html#ft_done_size
func (s *Size) Free() error {
	if s == nil || s.ptr == nil {
		return nil
	}

	err := getErr(C.FT_Done_Size(s.ptr))
	s.ptr = nil
	return err
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

// GlyphSlot is a container where individual glyphs can be loaded, be they in
// outline or bitmap format.
//
// If LoadGlyph is called with default flags (see LoadDefault) the glyph image
// is loaded in the glyph slot in its native format (e.g., an outline glyph for
// TrueType and Type 1 formats). [Since 2.9] The prospective bitmap metrics are
// calculated according to LoadTarget and other flags even for the outline
// glyph, even if LoadRender is not set.
//
// This image can later be converted into a bitmap by calling RenderGlyph. This
// function searches the current renderer for the native image's format,
// then invokes it.
//
// The renderer is in charge of transforming the native image through the slot's
// face transformation fields, then converting it into a bitmap that is returned
// in Bitmap().
//
// Note that BitmapLeft and BitmapTop are also used to specify the position of
// the bitmap relative to the current pen position (e.g., coordinates (0,0) on
// the baseline). Of course, Format is also changed to GlyphFormatBitmap.
//
// Here is a small pseudo code fragment that shows how to use LsbDelta and
// RsbDelta to do fractional positioning of glyphs:
//
//	FT_GlyphSlot  slot     = face->glyph;                                <- TODO
//	FT_Pos        origin_x = 0;
//
//
//	for all glyphs do
//	<load glyph with `FT_Load_Glyph'>
//
//	FT_Outline_Translate( slot->outline, origin_x & 63, 0 );
//
//	<save glyph image, or render glyph, or ...>
//
//	<compute kern between current and next glyph
//		and add it to `origin_x'>
//
//	origin_x += slot->advance.x;
//	origin_x += slot->lsb_delta - slot->rsb_delta;
//	endfor
//
// Here is another small pseudo code fragment that shows how to use LsbDelta and
// RsbDelta to improve integer positioning of glyphs:
//
//	FT_GlyphSlot  slot           = face->glyph;                          <- TODO
//	FT_Pos        origin_x       = 0;
//	FT_Pos        prev_rsb_delta = 0;
//
//
//	for all glyphs do
//	<compute kern between current and previous glyph
//		and add it to `origin_x'>
//
//	<load glyph with `FT_Load_Glyph'>
//
//	if ( prev_rsb_delta - slot->lsb_delta >  32 )
//		origin_x -= 64;
//	else if ( prev_rsb_delta - slot->lsb_delta < -31 )
//		origin_x += 64;
//
//	prev_rsb_delta = slot->rsb_delta;
//
//	<save glyph image, or render glyph, or ...>
//
//	origin_x += slot->advance.x;
//	endfor
//
// If you use strong auto-hinting, you must apply these delta values! Otherwise
// you will experience far too large inter-glyph spacing at small rendering
// sizes in most cases. Note that it doesn't harm to use the above code for
// other hinting modes also, since the delta values are zero then.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_glyphslot
type GlyphSlot struct {
	ptr C.FT_GlyphSlot `deep:"-"`
	// The glyph index passed as an argument to LoadGlyph while
	// initializing the glyph slot.
	GlyphIndex GlyphIndex
	// The metrics of the last loaded glyph in the slot. The returned values
	// depend on the last load flags (see the LoadGlyph API function) and can be
	// expressed either in 26.6 fractional pixels or font units.
	//
	// Note that even when the glyph image is transformed, the metrics are not.
	Metrics GlyphMetrics
	// The advance width of the unhinted glyph. Its value is expressed in 16.16
	// fractional pixels, unless LoadLinearDesign is set when loading the glyph.
	// This field can be important to perform correct WYSIWYG layout.
	// Only relevant for outline glyphs.
	LinearHoriAdvance fixed.Int16_16
	// The advance height of the unhinted glyph. Its value is expressed in 16.16
	// fractional pixels, unless LoadLinearDesign is set when loading the glyph.
	// This field can be important to perform correct WYSIWYG layout.
	// Only relevant for outline glyphs.
	LinearVertAdvance fixed.Int16_16
	// This shorthand is, depending on LoadIgnoreTransform, the transformed
	// (hinted) advance width for the glyph, in 26.6 fractional pixel format.
	// As specified with LoadVerticalLayout, it uses either the horiAdvance or
	// the vertAdvance value of metrics field.
	Advance Vector26_6
	// This field indicates the format of the image contained in the glyph slot.
	// Typically GlyphFormatBitmap, GlyphFormatOutline, or GlyphFormatComposite,
	// but other values are possible.
	Format GlyphFormat
	// This field is used as a bitmap descriptor.
	Bitmap *Bitmap
	// The bitmap's left bearing expressed in integer pixels.
	BitmapLeft int
	// The bitmap's top bearing expressed in integer pixels. This is the
	// distance from the baseline to the top-most glyph scanline, upwards y
	// coordinates being positive.
	BitmapTop int
	// The outline descriptor for the current glyph image if its format is
	// GlyphFormatOutline. Once a glyph is loaded, outline can be transformed,
	// distorted, emboldened, etc.
	//
	// [Since 2.10.1] If LoadNoScale is set, outline coordinates of OpenType
	// variation fonts for a selected instance are internally handled as 26.6
	// fractional font units but returned as (rounded) integers, as expected.
	// To get unrounded font units, don't use LoadNoScale but load the glyph
	// with LoadNoHinting and scale it, using the font's UnitsPerEM value as the
	// ppem.
	Outline *Outline
	// The number of subglyphs in a composite glyph. This field is only valid
	// for the composite glyph format that should normally only be loaded with
	// the LoadNoRecurse flag.
	NumSubglyphs int
	// The difference between hinted and unhinted left side bearing while
	// auto-hinting is active. Zero otherwise.
	LsbDelta Pos
	// The difference between hinted and unhinted right side bearing while
	// auto-hinting is active. Zero otherwise.
	RsbDelta Pos
}

func (g *GlyphSlot) reload() {
	if g == nil {
		return
	}

	if g.ptr == nil {
		*g = GlyphSlot{}
	}

	g.GlyphIndex = GlyphIndex(g.ptr.glyph_index)
	g.Metrics = GlyphMetrics{
		Width:        Pos(g.ptr.metrics.width),
		Height:       Pos(g.ptr.metrics.height),
		HoriBearingX: Pos(g.ptr.metrics.horiBearingX),
		HoriBearingY: Pos(g.ptr.metrics.horiBearingY),
		HoriAdvance:  Pos(g.ptr.metrics.horiAdvance),
		VertBearingX: Pos(g.ptr.metrics.vertBearingX),
		VertBearingY: Pos(g.ptr.metrics.vertBearingY),
		VertAdvance:  Pos(g.ptr.metrics.vertAdvance),
	}
	g.LinearHoriAdvance = fixed.Int16_16(g.ptr.linearHoriAdvance)
	g.LinearVertAdvance = fixed.Int16_16(g.ptr.linearVertAdvance)
	g.Advance = Vector26_6{
		X: fixed.Int26_6(g.ptr.advance.x),
		Y: fixed.Int26_6(g.ptr.advance.y),
	}
	g.Format = GlyphFormat(g.ptr.format)
	g.BitmapLeft = int(g.ptr.bitmap_left)
	g.BitmapTop = int(g.ptr.bitmap_top)
	g.NumSubglyphs = int(g.ptr.num_subglyphs)
	g.LsbDelta = Pos(g.ptr.lsb_delta)
	g.RsbDelta = Pos(g.ptr.rsb_delta)

	if g.Bitmap == nil {
		g.Bitmap = &Bitmap{}
	}
	g.Bitmap.ptr = &g.ptr.bitmap
	g.Bitmap.reload()

	if g.Outline == nil {
		g.Outline = &Outline{}
	}
	g.Outline.ptr = &g.ptr.outline
	g.Outline.reload()
}

// SubGlyphInfo contains info about a subglyph.
//
// The values of Arg1, Arg2, and Transform must be interpreted depending on the
// flags present in Flags. See the OpenType specification for details.
//
// https://docs.microsoft.com/en-us/typography/opentype/spec/glyf#composite-glyph-description
type SubGlyphInfo struct {
	Index     GlyphIndex
	Flags     SubGlyphFlag
	Arg1      int
	Arg2      int
	Transform Matrix
}

// SubGlyphInfo retrieves a description of a given subglyph. Only use it if
// glyph.Format is GlyphFormatComposite; an error is returned otherwise.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_get_subglyph_info
func (g *GlyphSlot) SubGlyphInfo(idx int) (SubGlyphInfo, error) {
	if g == nil || g.ptr == nil {
		return SubGlyphInfo{}, ErrInvalidArgument
	}

	if g.Format != GlyphFormatComposite {
		return SubGlyphInfo{}, ErrInvalidArgument
	}

	if C.uint(idx) >= g.ptr.num_subglyphs {
		return SubGlyphInfo{}, ErrInvalidArgument
	}

	var index C.FT_Int
	var flags C.FT_UInt
	var arg1 C.FT_Int
	var arg2 C.FT_Int
	var transform C.FT_Matrix
	if err := getErr(C.FT_Get_SubGlyph_Info(g.ptr, C.uint(idx), &index, &flags, &arg1, &arg2, &transform)); err != nil {
		return SubGlyphInfo{}, err
	}

	return SubGlyphInfo{
		Index: GlyphIndex(index),
		Flags: SubGlyphFlag(flags),
		Arg1:  int(arg1),
		Arg2:  int(arg2),
		Transform: Matrix{
			Xx: fixed.Int16_16(transform.xx),
			Xy: fixed.Int16_16(transform.xy),
			Yx: fixed.Int16_16(transform.yx),
			Yy: fixed.Int16_16(transform.yy),
		},
	}, nil
}

// RenderGlyph converts a given glyph image to a bitmap. It does so by inspecting
// the glyph image format, finding the relevant renderer, and invoking it.
//
// If RenderModeNormal is used, a previous call of Face.LoadGlyph with flag
// LoadColor makes RenderGlyph provide a default blending of colored glyph layers
// associated with the current glyph slot (provided the font contains such layers)
// instead of rendering the glyph slot's outline. This is an experimental feature;
// see LoadColor for more information.
//
// To get meaningful results, font scaling values must be set with functions like
// Face.SetCharSize before calling RenderGlyph.
//
// When FreeType outputs a bitmap of a glyph, it really outputs an alpha coverage
// map. If a pixel is completely covered by a filled-in outline, the bitmap
// contains 0xFF at that pixel, meaning that 0xFF/0xFF fraction of that pixel is
// covered, meaning the pixel is 100% black (or 0% bright). If a pixel is only
// 50% covered (value 0x80), the pixel is made 50% black (50% bright or a middle
// shade of grey). 0% covered means 0% black (100% bright or white).
//
// On high-DPI screens like on smartphones and tablets, the pixels are so small
// that their chance of being completely covered and therefore completely black
// are fairly good. On the low-DPI screens, however, the situation is different.
// The pixels are too large for most of the details of a glyph and shades of
// gray are the norm rather than the exception.
//
// This is relevant because all our screens have a second problem: they are not
// linear. 1 + 1 is not 2. Twice the value does not result in twice the brightness.
// When a pixel is only 50% covered, the coverage map says 50% black, and this
// translates to a pixel value of 128 when you use 8 bits per channel (0-255).
// However, this does not translate to 50% brightness for that pixel on our sRGB
// and gamma 2.2 screens. Due to their non-linearity, they dwell longer in the
// darks and only a pixel value of about 186 results in 50% brightness -- 128
// ends up too dark on both bright and dark backgrounds. The net result is that
// dark text looks burnt-out, pixely and blotchy on bright background, bright
// text too frail on dark backgrounds, and colored text on colored background
// (for example, red on green) seems to have dark halos or ‘dirt’ around it. The
// situation is especially ugly for diagonal stems like in ‘w’ glyph shapes where
// the quality of FreeType's anti-aliasing depends on the correct display of
// grays. On high-DPI screens where smaller, fully black pixels reign supreme,
// this doesn't matter, but on our low-DPI screens with all the gray shades, it
// does. 0% and 100% brightness are the same things in linear and non-linear
// space, just all the shades in-between aren't.
//
// The blending function for placing text over a background is
//	dst := alpha * src + (1 - alpha) * dst
// which is known as the OVER operator.
//
// To correctly composite an antialiased pixel of a glyph onto a surface:
//
// 1 - take the foreground and background colors (e.g., in sRGB space) and apply
// gamma to get them in a linear space
//
// 2 - use OVER to blend the two linear colors using the glyph pixel as the
// alpha value (remember, the glyph bitmap is an alpha coverage bitmap)
//
// 3 - apply inverse gamma to the blended pixel and write it back to the image.
//
// Internal testing at Adobe found that a target inverse gamma of 1.8 for step 3
// gives good results across a wide range of displays with an sRGB gamma curve or
// a similar one.
//
// This process can cost performance. There is an approximation that does not
// need to know about the background color; see https://bel.fi/alankila/lcd/ and
// https://bel.fi/alankila/lcd/alpcor.html for details.
//
// ATTENTION: Linear blending is even more important when dealing with
// subpixel-rendered glyphs to prevent color-fringing! A subpixel-rendered
// glyph must first be filtered with a filter that gives equal weight to the
// three color primaries and does not exceed a sum of 0x100, see section
// ‘Subpixel Rendering’. Then the only difference to gray linear blending is that
// subpixel-rendered linear blending is done 3 times per pixel: red foreground
// subpixel to red background subpixel and so on for green and blue.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_render_glyph
func (g *GlyphSlot) RenderGlyph(mode RenderMode) error {
	if g == nil || g.ptr == nil {
		return ErrInvalidArgument
	}

	if err := getErr(C.FT_Render_Glyph(g.ptr, C.FT_Render_Mode(mode))); err != nil {
		return err
	}

	g.reload()
	return nil
}

// GlyphMetrics models the metrics of a single glyph. The values are expressed
// in 26.6 fractional pixel format; if the flag LoadNoScale has been used while
// loading the glyph, values are expressed in font units instead.
//
// If not disabled with LoadNoHinting, the values represent dimensions of the
// hinted glyph (in case hinting is applicable).
//
// Stroking a glyph with an outside border does not increase HoriAdvance or
// VertAdvance; you have to manually adjust these values to account for the
// added width and height.
//
// FreeType doesn't use the ‘VORG’ table data for CFF fonts because it doesn't
// have an interface to quickly retrieve the glyph height. The y coordinate of
// the vertical origin can be simply computed as vertBearingY + height after
// loading a glyph.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_glyph_metrics
type GlyphMetrics struct {
	// The glyph's width.
	Width Pos
	// The glyph's height.
	Height Pos

	// Left side bearing for horizontal layout.
	HoriBearingX Pos
	// Top side bearing for horizontal layout.
	HoriBearingY Pos
	// Advance width for horizontal layout.
	HoriAdvance Pos

	// Left side bearing for vertical layout.
	VertBearingX Pos
	// Top side bearing for vertical layout. Larger positive values mean further
	// below the vertical glyph origin.
	VertBearingY Pos
	// Advance height for vertical layout. Positive values mean the glyph has a
	// positive advance downward.
	VertAdvance Pos
}
