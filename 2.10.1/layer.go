package freetype2

// #include <stdlib.h>
// #include <ft2build.h>
// #include FT_FREETYPE_H
//
// FT_Error GetColorGlyphLayers(FT_Face face, FT_UInt base_glyph, FT_UInt** ptr, int* len) {
// 	if (face == NULL) {
// 		return FT_Err_Invalid_Face_Handle;
// 	}
//
// 	FT_LayerIterator iterator;
// 	iterator.p = NULL;
//
// 	FT_UInt layer_glyph_index;
// 	FT_UInt layer_color_index;
//
// 	if (!FT_Get_Color_Glyph_Layer(face, base_glyph, &layer_glyph_index, &layer_color_index, &iterator)) {
// 		return 0;
// 	}
//
// 	if (iterator.num_layers <= 0) {
// 		return 0;
// 	}
//
// 	FT_UInt* innerptr = (FT_UInt*)calloc(iterator.num_layers, sizeof(FT_UInt) * 2);
// 	if (innerptr == NULL) {
// 		return FT_Err_Out_Of_Memory;
// 	}
//
// 	int i = 0;
// 	do {
// 		innerptr[i]   = layer_glyph_index;
// 		innerptr[i+1] = layer_color_index;
//		i+=2;
// 	} while (FT_Get_Color_Glyph_Layer(face, base_glyph, &layer_glyph_index, &layer_color_index, &iterator));
//
//	*ptr = innerptr;
//	*len = iterator.num_layers * 2;
//
// 	return 0;
// }
import "C"
import "unsafe"

// ColorLayer models the layer of a COLR record
type ColorLayer struct {
	GlyphIndex, ColorIndex int
}

// GetColorGlyphLayers is the same as ForEachColorGlyphLayer, but it will
// eagearly allocate a slice with a single CGO call, instead of the N calls of
// its counterpart.
//
// The slice has an upper bound of uint16 elements.
func (f *Face) GetColorGlyphLayers(baseGlyph GlyphIndex) []ColorLayer {
	if f == nil || f.ptr == nil {
		return nil
	}

	var length C.int
	var data *C.FT_UInt

	if err := getErr(C.GetColorGlyphLayers(f.ptr, C.uint(baseGlyph), &data, &length)); err != nil {
		panic(err)
	}

	if data == nil {
		return nil
	}

	ptr := (*[(1<<31 - 1) / C.sizeof_FT_UInt]C.FT_UInt)(unsafe.Pointer(data))[:length:length]
	ret := make([]ColorLayer, length/2)
	for i := range ret {
		ret[i] = ColorLayer{
			GlyphIndex: int(ptr[i*2]),
			ColorIndex: int(ptr[i*2+1]),
		}
	}

	return ret
}

// ForEachColorGlyphLayer is an interface to the ‘COLR’ table in OpenType fonts to
// iteratively retrieve the colored glyph layers associated with the current
// glyph slot. https://docs.microsoft.com/en-us/typography/opentype/spec/colr
//
// It returns the number of layers.
//
// The callback will be called for each layer, receiving the glyph index and the
// color index into the font face's color palette of the current layer.
// Note that the color index 0xFFFF is special; it doesn't reference a palette
// entry but indicates that the text foreground color should be used instead
// (to be set up by the application outside of FreeType).
//
// To stop iteration, return true from the callback.
//
// The color palette can be retrieved with SelectPalette.
//
// The glyph layer data for a given glyph index, if present, provides an
// alternative, multi-colour glyph representation: Instead of rendering the
// outline or bitmap with the given glyph index, glyphs with the indices and
// colors returned by this function are rendered layer by layer.
//
// The returned elements are ordered in the z direction from bottom to top; the
// 'n'th element should be rendered with the associated palette color and
// blended on top of the already rendered layers (elements 0, 1, ..., n-1).
//
// NOTE:
// This function is necessary if you want to handle glyph layers by yourself.
// In particular, functions that operate with Glyph objects (like GetGlyph or
// GlyphToBitmap) don't have access to this information. TODO: revisit these names, they're not implemented yet, might change
//
// Note that RenderGlyph is able to handle colored glyph layers automatically if
// the LoadColor flag is passed to a previous call to LoadGlyph.
// [This is an experimental feature.]
//
// example TODO
//	FT_Color*         palette;
//	FT_LayerIterator  iterator;
//
//	FT_Bool  have_layers;
//	FT_UInt  layer_glyph_index;
//	FT_UInt  layer_color_index;
//
//
//	error = FT_Palette_Select( face, palette_index, &palette );
//	if ( error )
//	palette = NULL;
//
//	iterator.p  = NULL;
//	have_layers = FT_Get_Color_Glyph_Layer( face,
//											glyph_index,
//											&layer_glyph_index,
//											&layer_color_index,
//											&iterator );
//
//	if ( palette && have_layers )
//	{
//	do
//	{
//		FT_Color  layer_color;
//
//
//		if ( layer_color_index == 0xFFFF )
//		layer_color = text_foreground_color;
//		else
//		layer_color = palette[layer_color_index];
//
//		// Load and render glyph `layer_glyph_index', then
//		// blend resulting pixmap (using color `layer_color')
//		// with previously created pixmaps.
//
//	} while ( FT_Get_Color_Glyph_Layer( face,
//										glyph_index,
//										&layer_glyph_index,
//										&layer_color_index,
//										&iterator ) );
//	}
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-layer_management.html#ft_get_color_glyph_layer
func (f *Face) ForEachColorGlyphLayer(baseGlyph GlyphIndex, fn func(layerGlyphIndex, layerColorIndex int) (done bool)) int {
	if f == nil || f.ptr == nil {
		return 0
	}

	var (
		iterator        C.FT_LayerIterator
		layerGlyphIndex C.FT_UInt
		layerColorIndex C.FT_UInt
	)

	hasLayers := C.FT_Get_Color_Glyph_Layer(f.ptr, C.uint(baseGlyph), &layerGlyphIndex, &layerColorIndex, &iterator) == 1
	numLayers := int(iterator.num_layers)

	if !hasLayers || fn == nil {
		return numLayers
	}

	// do
	if fn(int(layerGlyphIndex), int(layerColorIndex)) {
		return numLayers
	}
	// while
	for C.FT_Get_Color_Glyph_Layer(f.ptr, C.uint(baseGlyph), &layerGlyphIndex, &layerColorIndex, &iterator) == 1 {
		if fn(int(layerGlyphIndex), int(layerColorIndex)) {
			return numLayers
		}
	}

	return numLayers
}
