package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"
import (
	"fmt"
	"io"
	"io/ioutil"
	"unsafe"
)

// Version contains version information
type Version struct {
	Major, Minor, Patch int
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Library is a handle to a FreeType library instance. Each ‘library’ is completely independent from the others; it is
// the ‘root’ of a set of objects like fonts, faces, sizes, etc.
//
// Library is NOT safe for concurrent use.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_library
type Library struct {
	ptr   C.FT_Library `deep:"-"`
	faces []*Face
}

// NewLibrary creates a new Library instance.
// The set of modules that are registered by this function is determined at build time.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_init_freetype
func NewLibrary() (*Library, error) {
	var ft C.FT_Library
	if err := getErr(C.FT_Init_FreeType(&ft)); err != nil {
		return nil, err
	}
	return &Library{ptr: ft}, nil
}

// Free releases the library instance, destroying it and all of its children, including resources, drivers, faces,
// sizes, etc.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_done_freetype
func (l *Library) Free() error {
	if l == nil || l.ptr == nil {
		return nil
	}

	if err := getErr(C.FT_Done_FreeType(l.ptr)); err != nil {
		return err
	}

	for _, f := range l.faces {
		f.freeInternal()
		l.faces = l.faces[1:]
	}

	l.ptr = nil
	return nil
}

// Version reports the version of the FreeType library being used.
// This is useful when dynamically linking to the library.
//
// The reason why this function takes a library argument is because certain programs implement library initialization in
// a custom way that doesn't use InitFreeType.
//
// In such cases, the library version might not be available before the library object has been created.
func (l *Library) Version() Version {
	if l == nil {
		return Version{}
	}

	var major, minor, patch C.FT_Int
	C.FT_Library_Version(l.ptr, &major, &minor, &patch)
	return Version{
		Major: int(major),
		Minor: int(minor),
		Patch: int(patch),
	}
}

// NewFace creates a new face from the given io.Reader.
// Beware, the data will be read, all at once, into memory.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_new_memory_face
func (l *Library) NewFace(r io.Reader, index, namedInstanceIndex int) (*Face, error) {
	if l == nil || l.ptr == nil {
		return nil, ErrInvalidLibraryHandle
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, ErrUnknownFileFormat
	}

	cdata := C.CBytes(data)
	free := func() { free(cdata) }

	var face C.FT_Face
	if err := getErr(C.FT_New_Memory_Face(
		l.ptr,
		(*C.uchar)(cdata),
		C.long(len(data)),
		C.FT_Long(index&0xFFFF|namedInstanceIndex<<16),
		&face,
	)); err != nil {
		free()
		return nil, err
	}

	f := &Face{ptr: face, lib: l, dealloc: []func(){free}}
	l.faces = append(l.faces, f)
	return f, nil
}

// NewFaceFromPath creates a new face from the given path.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_new_face
func (l *Library) NewFaceFromPath(path string, index, namedInstanceIndex int) (*Face, error) {
	if l == nil || l.ptr == nil {
		return nil, ErrInvalidLibraryHandle
	}

	var face C.FT_Face
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	if err := getErr(C.FT_New_Face(
		l.ptr,
		cpath,
		C.FT_Long(index&0xFFFF|namedInstanceIndex<<16),
		&face,
	)); err != nil {
		return nil, err
	}

	f := &Face{ptr: face, lib: l}
	l.faces = append(l.faces, f)
	return f, nil
}
