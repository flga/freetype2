package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_COLOR_H
import "C"
import (
	"image/color"
	"unsafe"
)

// PaletteFlag is a list of bit field constants used in the Flags slice of the
// Palette struct to indicate for which background a palette with a given index
// is usable.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-color_management.html#ft_palette_xxx
type PaletteFlag uint16

const (
	// PaletteForLightBackground the palette is appropriate to use when
	// displaying the font on a light background such as white.
	PaletteForLightBackground PaletteFlag = C.FT_PALETTE_FOR_LIGHT_BACKGROUND
	// PaletteForDarkBackground the palette is appropriate to use when
	// displaying the font on a dark background such as black.
	PaletteForDarkBackground PaletteFlag = C.FT_PALETTE_FOR_DARK_BACKGROUND
)

// PaletteData holds the data of the ‘CPAL’ table.
//
// Use GetSfntName to map name IDs and entry name IDs to name strings.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-color_management.html#ft_palette_data
type PaletteData struct {
	// The number of palettes.
	NumPalettes int
	// Read-only slice of palette name IDs, corresponding to entries like
	// ‘dark’ or ‘light’ in the font's ‘name’ table.
	// An empty name ID in the ‘CPAL’ table gets represented as value 0xFFFF.
	// Nil if the font's ‘CPAL’ table doesn't contain appropriate data.
	NameIDs []uint16
	// Read-only slice of palette flags. Possible values are an ORed combination
	// of PaletteForLightBackground and PaletteForDarkBackground.
	// Nil if the font's ‘CPAL’ table doesn't contain appropriate data.
	Flags []PaletteFlag
	// The number of entries in a single palette. All palettes have the same size.
	NumPaletteEntries int
	// Read-only slice of palette entry name IDs. In each palette, entries with
	// the same index have the same function. For example, index 0 might
	// correspond to string ‘outline’ in the font's ‘name’ table to indicate
	// that this palette entry is used for outlines, index 1 might correspond to
	// ‘fill’ to indicate the filling color palette entry, etc.
	//
	// An empty entry name ID in the ‘CPAL’ table gets represented as value 0xFFFF.
	// Nil if the font's ‘CPAL’ table doesn't contain appropriate data.
	EntryNameIDs []uint16
}

func newPaletteData(c C.FT_Palette_Data) PaletteData {
	var nameIDs []uint16
	if c.num_palettes > 0 && c.palette_name_ids != nil {
		nameIDs = make([]uint16, c.num_palettes)
		ptr := (*[(1<<31 - 1) / C.sizeof_FT_UShort]C.FT_UShort)(unsafe.Pointer(c.palette_name_ids))[:c.num_palettes:c.num_palettes]
		for i := range nameIDs {
			nameIDs[i] = uint16(ptr[i])
		}
	}

	var flags []PaletteFlag
	if c.num_palettes > 0 && c.palette_flags != nil {
		flags = make([]PaletteFlag, c.num_palettes)
		ptr := (*[(1<<31 - 1) / C.sizeof_FT_UShort]C.FT_UShort)(unsafe.Pointer(c.palette_flags))[:c.num_palettes:c.num_palettes]
		for i := range flags {
			flags[i] = PaletteFlag(ptr[i])
		}
	}

	var entryNameIDs []uint16
	if c.num_palette_entries > 0 && c.palette_entry_name_ids != nil {
		entryNameIDs = make([]uint16, c.num_palette_entries)
		ptr := (*[(1<<31 - 1) / C.sizeof_FT_UShort]C.FT_UShort)(unsafe.Pointer(c.palette_entry_name_ids))[:c.num_palette_entries:c.num_palette_entries]
		for i := range entryNameIDs {
			entryNameIDs[i] = uint16(ptr[i])
		}
	}

	return PaletteData{
		NumPalettes:       int(c.num_palettes),
		NameIDs:           nameIDs,
		Flags:             flags,
		NumPaletteEntries: int(c.num_palette_entries),
		EntryNameIDs:      entryNameIDs,
	}
}

// PaletteData retrieve the face's color palette data.
//
// It will return an error if the config macro TT_CONFIG_OPTION_COLOR_LAYERS is
// not defined in ftoption.h.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-color_management.html#ft_palette_data_get
func (f *Face) PaletteData() (PaletteData, error) {
	if f == nil || f.ptr == nil {
		return PaletteData{}, ErrInvalidFaceHandle
	}

	var palette C.FT_Palette_Data
	if err := getErr(C.FT_Palette_Data_Get(f.ptr, &palette)); err != nil {
		return PaletteData{}, err
	}

	return newPaletteData(palette), nil
}

// SelectPalette activates a palette for rendering color glyphs.
//
// It also provides a callback which can be used to modify the palette's color
// entries. Modified values will persist until SelectPalette is called again,
// in which case mutate will be called with the original values of the CPAL table.
//
// It will returns an error if the config macro TT_CONFIG_OPTION_COLOR_LAYERS is
// not defined in ftoption.h.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-color_management.html#ft_palette_select
func (f *Face) SelectPalette(idx int, mutate func(palette []color.RGBA)) ([]color.RGBA, error) {
	if f == nil || f.ptr == nil {
		return nil, ErrInvalidFaceHandle
	}

	var cpalette *C.FT_Color
	if err := getErr(C.FT_Palette_Select(f.ptr, C.ushort(idx), &cpalette)); err != nil {
		return nil, err
	}

	var data C.FT_Palette_Data
	if err := getErr(C.FT_Palette_Data_Get(f.ptr, &data)); err != nil {
		return nil, err
	}

	length := int(data.num_palette_entries)
	palette := make([]color.RGBA, length)
	ptr := (*[(1<<31 - 1) / C.sizeof_FT_Color]C.FT_Color)(unsafe.Pointer(cpalette))[:length:length]
	for i := range palette {
		palette[i] = color.RGBA{
			R: byte(ptr[i].red),
			G: byte(ptr[i].green),
			B: byte(ptr[i].blue),
			A: byte(ptr[i].alpha),
		}
	}

	if mutate != nil {
		mutate(palette)
		for i, c := range palette {
			ptr[i] = C.FT_Color{
				red:   C.uchar(c.R),
				green: C.uchar(c.G),
				blue:  C.uchar(c.B),
				alpha: C.uchar(c.A),
			}
		}
	}

	return palette, nil
}

// SetPaletteForeground sets palette index 0xFFFF to c.
// ‘COLR’ uses this to indicate a ‘text foreground color’.
//
// If this function isn't called, the text foreground color is set to white
// opaque (BGRA value 0xFFFFFFFF) if PaletteForDarkBackground is present for the
// current palette, and black opaque (BGRA value 0x000000FF) otherwise, including
// the case that no palette types are available in the ‘CPAL’ table.
//
// This function always returns an error if the config macro
// TT_CONFIG_OPTION_COLOR_LAYERS is not defined in ftoption.h.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-color_management.html#ft_palette_set_foreground_color
func (f *Face) SetPaletteForeground(c color.RGBA) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	return getErr(C.FT_Palette_Set_Foreground_Color(f.ptr, C.FT_Color{
		red:   C.uchar(c.R),
		green: C.uchar(c.G),
		blue:  C.uchar(c.B),
		alpha: C.uchar(c.A),
	}))
}
