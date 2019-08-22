package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_SIZES_H
import "C"

// NewSize creates a new Size.
//
// You need to call ActivateSize in order to select the new size for upcoming
// calls to SetPixelSizes, SetChar_Size, LoadGlyph, LoadChar, etc.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-sizes_management.html#ft_new_size
func (f *Face) NewSize() (*Size, error) {
	if f == nil || f.ptr == nil {
		return nil, ErrInvalidFaceHandle
	}

	var size C.FT_Size
	if err := getErr(C.FT_New_Size(f.ptr, &size)); err != nil {
		return nil, err
	}

	ret := newSize(size)
	f.dealloc = append(f.dealloc, func() {
		ret.ptr = nil
	})

	return ret, nil
}

// ActivateSize activates the given size.
//
// Even though it is possible to create several size objects for a given face
// (see NewSize for details), functions like LoadGlyph or LoadChar only use the
// one that has been activated last to determine the ‘current character pixel size’.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-sizes_management.html#ft_activate_size
func (f *Face) ActivateSize(s *Size) error {
	if f == nil || f.ptr == nil {
		return ErrInvalidFaceHandle
	}

	if s == nil || s.ptr == nil {
		return ErrInvalidSizeHandle
	}

	return getErr(C.FT_Activate_Size(s.ptr))
}
