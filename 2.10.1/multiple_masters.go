package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_MULTIPLE_MASTERS_H
import "C"
import (
	"unsafe"

	"github.com/flga/freetype2/fixed"
)

// VarAxisFlag is a list of bit flags.
type VarAxisFlag uint

// VarAxisFlagHidden the variation axis should not be exposed to user interfaces.
const VarAxisFlagHidden VarAxisFlag = C.FT_VAR_AXIS_FLAG_HIDDEN

// MMAxis models a given axis in design space for Multiple Masters fonts.
//
// This structure can't be used for TrueType GX or OpenType variation fonts.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_mm_axis
type MMAxis struct {
	Name string
	Min  int
	Max  int
}

// MultiMaster models the axes and space of a Multiple Masters font.
//
// This structure can't be used for TrueType GX or OpenType variation fonts.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_multi_master
type MultiMaster struct {
	// Number of axes. Cannot exceed 4.
	NumAxis uint
	// Number of designs; should be normally 2^num_axis even though the Type 1
	// specification strangely allows for intermediate designs to be present.
	// This number cannot exceed 16.
	NumDesigns uint
	// A table of axis descriptors.
	Axis []MMAxis
}

// VarAxisTag is the tag field of an axis record.
//
// The values of are not limited to the constants exposed, they are merely the
// tags that are defined in the spec, other values are also valid.
type VarAxisTag uint

const (
	//VarAxisTagItal is used to vary between non-italic and italic
	VarAxisTagItal VarAxisTag = 0x6974616c // (ital) - Italic
	// VarAxisTagOpsz is used to vary design to suit different text sizes.
	VarAxisTagOpsz VarAxisTag = 0x6f70737a // (opsz) - Optical size
	// VarAxisTagSlnt is used to vary between upright and slanted text.
	VarAxisTagSlnt VarAxisTag = 0x736c6e74 // (slnt) - Slant
	// VarAxisTagWdth is used to vary width of text from narrower to wider.
	VarAxisTagWdth VarAxisTag = 0x77647468 // (wdth) - Width
	// VarAxisTagWght is used to vary stroke thicknesses or other design details
	// to give variation from lighter to blacker.
	VarAxisTagWght VarAxisTag = 0x77676874 // (wght) - Weight
)

// VarAxis models a given axis in design space for Multiple Masters, TrueType GX,
// and OpenType variation fonts.
//
// The fields Min, Def, and Max are 16.16 fractional values for TrueType GX and
// OpenType variation fonts. For Adobe MM fonts, the values are integers.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_var_axis
type VarAxis struct {
	// The axis's name. Not always meaningful for TrueType GX or OpenType
	// variation fonts.
	Name string

	// The axis's minimum design coordinate.
	Min fixed.Int16_16
	// The axis's default design coordinate. FreeType computes meaningful default
	// values for Adobe MM fonts.
	Def fixed.Int16_16
	// The axis's maximum design coordinate.
	Max fixed.Int16_16

	// The axis's tag (the equivalent to ‘name’ for TrueType GX and OpenType
	// variation fonts). FreeType provides default values for Adobe MM fonts if
	// possible.
	Tag VarAxisTag
	// The axis name entry in the font's ‘name’ table. This is another (and often
	// better) version of the ‘name’ field for TrueType GX or OpenType variation
	// fonts. Not meaningful for Adobe MM fonts.
	Strid uint
	// The ‘flags’ field of an OpenType Variation Axis Record.
	// Not meaningful for Adobe MM fonts (it's always zero).
	Flags VarAxisFlag
}

// VarNamedStyle models a named instance in a TrueType GX or OpenType variation
// font.
//
// This structure can't be used for Adobe MM fonts.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_var_named_style
type VarNamedStyle struct {
	// The design coordinates for this instance. This is an array with one entry
	// for each axis.
	Coords []fixed.Int16_16
	// The entry in ‘name’ table identifying this instance.
	Strid uint
	// The entry in ‘name’ table identifying a PostScript name for this instance.
	// Value 0xFFFF indicates a missing entry.
	Psid uint
}

// MMVar models the axes and space of an Adobe MM, TrueType GX, or OpenType
// variation font.
//
// Some fields are specific to one format and not to the others.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_mm_var
type MMVar struct {
	// The number of axes. The maximum value is 4 for Adobe MM fonts; no limit
	// in TrueType GX or OpenType variation fonts.
	NumAxis uint
	// The number of designs; should be normally 2^NumAxis for Adobe MM fonts.
	// Not meaningful for TrueType GX or OpenType variation fonts (where every
	// glyph could have a different number of designs).
	NumDesigns uint
	// The number of named styles; a ‘named style’ is a tuple of design
	// coordinates that has a string ID (in the ‘name’ table) associated with it.
	// The font can tell the user that, for example, [Weight=1.5,Width=1.1] is
	// ‘Bold’. Another name for ‘named style’ is ‘named instance’.
	//
	// For Adobe Multiple Masters fonts, this value is always zero because the
	// format does not support named styles.
	NumNamedstyles uint
	// An axis descriptor table. TrueType GX and OpenType variation fonts contain
	// slightly more data than Adobe MM fonts.
	Axis []VarAxis
	// A named style (instance) table. Only meaningful for TrueType GX and
	// OpenType variation fonts.
	Namedstyle []VarNamedStyle
}

// MultiMaster retrieves a variation descriptor of a given Adobe MM font.
//
// This function can't be used with TrueType GX or OpenType variation fonts.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_multi_master
func (f *Face) MultiMaster() (MultiMaster, error) {
	if f == nil || f.ptr == nil {
		return MultiMaster{}, ErrInvalidFaceHandle
	}

	var master C.FT_Multi_Master
	if err := getErr(C.FT_Get_Multi_Master(f.ptr, &master)); err != nil {
		return MultiMaster{}, err
	}

	ret := MultiMaster{
		NumAxis:    uint(master.num_axis),
		NumDesigns: uint(master.num_designs),
		Axis:       make([]MMAxis, master.num_axis),
	}

	for i := range ret.Axis {
		v := master.axis[i]
		ret.Axis[i] = MMAxis{
			Name: C.GoString(v.name),
			Min:  int(v.minimum),
			Max:  int(v.maximum),
		}
	}

	return ret, nil
}

// MMVar retrieves a variation descriptor for a given font.
//
// This function works with all supported variation formats.
//
// It allocates a data structure, which the user must deallocate with a call
// to Free after use.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_mm_var
func (f *Face) MMVar() (*MMVar, error) {
	if f == nil || f.ptr == nil {
		return nil, ErrInvalidFaceHandle
	}
	if f.lib == nil || f.lib.ptr == nil {
		return nil, ErrInvalidFaceHandle
	}

	var master *C.FT_MM_Var
	if err := getErr(C.FT_Get_MM_Var(f.ptr, &master)); err != nil {
		return nil, err
	}
	defer C.FT_Done_MM_Var(f.lib.ptr, master)

	var numDesigns uint
	if master.num_designs != ^C.uint(0) {
		numDesigns = uint(master.num_designs)
	}
	ret := &MMVar{
		NumAxis:        uint(master.num_axis),
		NumDesigns:     numDesigns,
		NumNamedstyles: uint(master.num_namedstyles),
	}

	if master.num_axis > 0 && master.axis != nil {
		ret.Axis = make([]VarAxis, master.num_axis)

		ptr := (*[(1<<31 - 1) / C.sizeof_FT_Var_Axis]C.FT_Var_Axis)(unsafe.Pointer(master.axis))[:master.num_axis:master.num_axis]
		for i := range ret.Axis {
			var flags C.FT_UInt
			C.FT_Get_Var_Axis_Flags(master, C.uint(i), &flags)

			var strid uint
			if ptr[i].strid != ^C.uint(0) {
				strid = uint(ptr[i].strid)
			}

			ret.Axis[i] = VarAxis{
				Name:  C.GoString(ptr[i].name),
				Min:   fixed.Int16_16(ptr[i].minimum),
				Def:   fixed.Int16_16(ptr[i].def),
				Max:   fixed.Int16_16(ptr[i].maximum),
				Tag:   VarAxisTag(uint(ptr[i].tag)),
				Strid: strid,
				Flags: VarAxisFlag(flags),
			}
		}
	}

	if master.num_namedstyles > 0 && master.namedstyle != nil {
		ret.Namedstyle = make([]VarNamedStyle, master.num_namedstyles)

		ptr := (*[(1<<31 - 1) / C.sizeof_FT_Var_Named_Style]C.FT_Var_Named_Style)(unsafe.Pointer(master.namedstyle))[:master.num_namedstyles:master.num_namedstyles]
		numCoords := master.num_axis
		for i := range ret.Namedstyle {
			ret.Namedstyle[i] = VarNamedStyle{
				Coords: make([]fixed.Int16_16, numCoords),
				Strid:  uint(ptr[i].strid),
				Psid:   uint(ptr[i].psid),
			}

			coordsptr := (*[(1<<31 - 1) / C.sizeof_FT_Fixed]C.FT_Fixed)(unsafe.Pointer(ptr[i].coords))[:numCoords:numCoords]
			for j := range ret.Namedstyle[i].Coords {
				ret.Namedstyle[i].Coords[j] = fixed.Int16_16(coordsptr[j])
			}
		}
	}

	return ret, nil
}

// SetMMDesignCoords sets the mm design coordinates.
// For Adobe MM fonts, choose an interpolated font design through design coordinates.
//
// This function can't be used with TrueType GX or OpenType variation fonts.
//
// To reset all axes to the default values, call the function with coords set to
// nil (or an empty slice).
//
// If len(coords) is larger than zero, this function sets the FaceFlagVariation
// bit in the Flags field. If len(coords) is zero, this bit flag gets unset.
//
// If len(coords) is smaller than the number of axes, it will use default values
// for the remaining axes. Excess values will be ignored.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_mm_design_coordinates
func (f *Face) SetMMDesignCoords(coords []int) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	var ccoords *C.FT_Long
	if len(coords) > 0 {
		length := C.ulong(len(coords))
		block := C.calloc(C.size_t(length), C.sizeof_FT_Long)
		ptr := (*[(1<<31 - 1) / C.sizeof_FT_Long]C.FT_Long)(block)[:length:length]
		for i, v := range coords {
			ptr[i] = C.FT_Long(v)
		}
		ccoords = (*C.FT_Long)(block)
		defer free(block)
	}

	return getErr(C.FT_Set_MM_Design_Coordinates(f.ptr, C.uint(len(coords)), ccoords))
}

// SetVarDesignCoords sets the var design coordinates.
// Choose an interpolated font design through design coordinates.
//
// This function works with all supported variation formats.
//
// To reset all axes to the default values, call the function with with coords
// set to nil (or an empty slice).
//
// If len(coords) is larger than zero, this function sets the FaceFlagVariation
// bit in the Flags field. If len(coords) is zero, this bit flag gets unset.
//
// If len(coords) is smaller than the number of axes, it will use default values
// for the remaining axes. Excess values will be ignored.
//
// ‘Default values’ means the currently selected named instance (or the base
// font if no named instance is selected).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_var_design_coordinates
func (f *Face) SetVarDesignCoords(coords []fixed.Int16_16) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	var ccoords *C.FT_Fixed
	if len(coords) > 0 {
		length := C.ulong(len(coords))
		block := C.calloc(C.size_t(length), C.sizeof_FT_Fixed)
		ptr := (*[(1<<31 - 1) / C.sizeof_FT_Fixed]C.FT_Fixed)(block)[:length:length]
		for i, v := range coords {
			ptr[i] = C.FT_Fixed(v)
		}
		ccoords = (*C.FT_Fixed)(block)
		defer free(block)
	}

	return getErr(C.FT_Set_Var_Design_Coordinates(f.ptr, C.uint(len(coords)), ccoords))
}

// VarDesignCoords returns the design coordinates of the currently selected
// interpolated font.
//
// This function works with all supported variation formats.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_var_design_coordinates
func (f *Face) VarDesignCoords() ([]fixed.Int16_16, error) {
	if f == nil || f.ptr == nil {
		return nil, ErrInvalidFaceHandle
	}

	var master *C.FT_MM_Var
	if err := getErr(C.FT_Get_MM_Var(f.ptr, &master)); err != nil {
		return nil, err
	}

	length := master.num_axis
	if length <= 0 {
		return nil, nil
	}

	block := C.calloc(C.size_t(length), C.sizeof_FT_Fixed)
	defer free(block)

	if err := getErr(C.FT_Get_Var_Design_Coordinates(f.ptr, length, (*C.FT_Fixed)(block))); err != nil {
		return nil, err
	}

	ptr := (*[(1<<31 - 1) / C.sizeof_FT_Fixed]C.FT_Fixed)(block)[:length:length]
	ret := make([]fixed.Int16_16, length)
	for i := range ret {
		ret[i] = fixed.Int16_16(ptr[i])
	}

	return ret, nil
}

// SetMMBlendCoords sets the mm blend coordinates.
// Choose an interpolated font design through normalized blend coordinates.
//
// This function works with all supported variation formats.
//
// The design coordinates array (each element must be between 0 and 1.0 for Adobe
// MM fonts, and between -1.0 and 1.0 for TrueType GX and OpenType variation fonts).
//
// To reset all axes to the default values, call the function with with coords
// set to nil (or an empty slice).
//
// If len(coords) is larger than zero, this function sets the FaceFlagVariation
// bit in the Flags field. If len(coords) is zero, this bit flag gets unset.
//
// If len(coords) is smaller than the number of axes, it will use default values
// for the remaining axes. Excess values will be ignored.
//
// ‘Default values’ means the currently selected named instance (or the base
// font if no named instance is selected).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_mm_blend_coordinates
func (f *Face) SetMMBlendCoords(coords []fixed.Int16_16) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	var ccoords *C.FT_Fixed
	if len(coords) > 0 {
		length := C.ulong(len(coords))
		block := C.calloc(C.size_t(length), C.sizeof_FT_Fixed)
		ptr := (*[(1<<31 - 1) / C.sizeof_FT_Fixed]C.FT_Fixed)(block)[:length:length]
		for i, v := range coords {
			ptr[i] = C.FT_Fixed(v)
		}
		ccoords = (*C.FT_Fixed)(block)
		defer free(block)
	}

	return getErr(C.FT_Set_MM_Blend_Coordinates(f.ptr, C.uint(len(coords)), ccoords))
}

// MMBlendCoords returns the normalized blend coordinates of the currently
// selected interpolated font.
//
// This function works with all supported variation formats.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_mm_blend_coordinates
func (f *Face) MMBlendCoords() ([]fixed.Int16_16, error) {
	if f == nil || f.ptr == nil {
		return nil, ErrInvalidFaceHandle
	}

	var master *C.FT_MM_Var
	if err := getErr(C.FT_Get_MM_Var(f.ptr, &master)); err != nil {
		return nil, err
	}

	length := master.num_axis
	if length <= 0 {
		return nil, nil
	}

	block := C.calloc(C.size_t(length), C.sizeof_FT_Fixed)
	defer free(block)

	if err := getErr(C.FT_Get_MM_Blend_Coordinates(f.ptr, length, (*C.FT_Fixed)(block))); err != nil {
		return nil, err
	}

	ptr := (*[(1<<31 - 1) / C.sizeof_FT_Fixed]C.FT_Fixed)(block)[:length:length]
	ret := make([]fixed.Int16_16, length)
	for i := range ret {
		ret[i] = fixed.Int16_16(ptr[i])
	}

	return ret, nil
}

// SetVarBlendCoords is an alias of SetMMBlendCoords.
func (f *Face) SetVarBlendCoords(coords []fixed.Int16_16) error {
	return f.SetMMBlendCoords(coords)
}

// VarBlendCoords is an alias of MMBlendCoords.
func (f *Face) VarBlendCoords() ([]fixed.Int16_16, error) {
	return f.MMBlendCoords()
}

// SetMMWeightVector sets the mm weight vector.
// For Adobe MM fonts, choose an interpolated font design by directly setting
// the weight vector.
//
// This function can't be used with TrueType GX or OpenType variation fonts.
//
// Adobe Multiple Master fonts limit the number of designs, and thus the length
// of the weight vector to 16.
//
// If vec is null or empty, the weight vector array is reset to the default values.
//
// The Adobe documentation also states that the values in the WeightVector array
// must total 1.0 ± 0.001. In practice this does not seem to be enforced, so is
// not enforced here, either.
//
// If len(vec) is larger than the number of designs, the extra values are ignored.
// If it is less than the number of designs, the remaining values are set to zero.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_mm_weightvector
func (f *Face) SetMMWeightVector(vec []fixed.Int16_16) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	var cvec *C.FT_Fixed
	if len(vec) > 0 {
		length := C.ulong(len(vec))
		block := C.calloc(C.size_t(length), C.sizeof_FT_Fixed)
		ptr := (*[(1<<31 - 1) / C.sizeof_FT_Fixed]C.FT_Fixed)(block)[:length:length]
		for i, v := range vec {
			ptr[i] = C.FT_Fixed(v)
		}
		cvec = (*C.FT_Fixed)(block)
		defer free(block)
	}

	return getErr(C.FT_Set_MM_WeightVector(f.ptr, C.uint(len(vec)), cvec))
}

// MMWeightVector retrieves the current weight vector of the font for Adobe MM
// fonts.
//
// This function can't be used with TrueType GX or OpenType variation fonts.
//
// Adobe Multiple Master fonts limit the number of designs, and thus the length
// of the WeightVector to 16.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_mm_weightvector
func (f *Face) MMWeightVector() ([]fixed.Int16_16, error) {
	if f == nil || f.ptr == nil {
		return nil, ErrInvalidFaceHandle
	}

	var master *C.FT_MM_Var
	if err := getErr(C.FT_Get_MM_Var(f.ptr, &master)); err != nil {
		return nil, err
	}

	length := master.num_designs
	if length <= 0 {
		return nil, nil
	}

	block := C.calloc(C.size_t(length), C.sizeof_FT_Fixed)
	defer free(block)

	clength := length
	if err := getErr(C.FT_Get_MM_WeightVector(f.ptr, &clength, (*C.FT_Fixed)(block))); err != nil {
		return nil, err
	}

	ptr := (*[(1<<31 - 1) / C.sizeof_FT_Fixed]C.FT_Fixed)(block)[:length:length]
	ret := make([]fixed.Int16_16, length)
	for i := range ret {
		ret[i] = fixed.Int16_16(ptr[i])
	}

	return ret, nil
}

// SetNamedInstance sets the current named instance.
//
// The index of the requested instance starts with value 1.
// If set to value 0, FreeType switches to font access without a named instance.
//
// The function uses this index to set bits 16-30 of the face's NamedIndex field.
// It also resets any variation applied to the font, and the FaceFlagVariation
// bit of the face's flags gets reset to zero.
//
// For Adobe MM fonts (which don't have named instances) this function simply
// resets the current face to the default instance.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_named_instance
func (f *Face) SetNamedInstance(idx int) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	return getErr(C.FT_Set_Named_Instance(f.ptr, C.uint(idx)))
}
