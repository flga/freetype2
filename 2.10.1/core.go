package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"
import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"unsafe"
)

var (
	errInvalidLib = errors.New("library must not be nil")
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
	ptr C.FT_Library
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

// FaceIndex is the index of a face in a given font file. It holds two different values.
// Bits 0-15 are the index of the face in the font file (starting with value 0). Set it to 0 if there is only one face
// in the font file.
//
// [Since 2.6.1] Bits 16-30 are relevant to GX and OpenType variation fonts only, specifying the named instance index
// for the current face index (starting with value 1; value 0 makes FreeType ignore named instances).
// For non-variation fonts, bits 16-30 are ignored. Assuming that you want to access the third named instance in face 4,
// the value should be set to 0x00030004. If you want to access face 4 without variation handling, simply set it to 4.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_open_face (face_index argument)
type FaceIndex int

// Face models a given typeface, in a given style.
//
// A Face object can only be safely used from one goroutine at a time. Similarly, creation and destruction of a Face
// with the same Library object can only be done from one goroutine at a time. On the other hand, functions like
// LoadGlyph and its siblings are thread-safe and do not need the lock to be held as long as the same Face object is not
// used from multiple goroutines at the same time.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_face
type Face struct {
	ptr     C.FT_Face
	dealloc []func()
}

// NewFace creates a new face from the given io.Reader.
// Beware, the data will be read, all at once, into memory.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_new_memory_face
func NewFace(l *Library, r io.Reader, idx FaceIndex) (*Face, error) {
	if l == nil || l.ptr == nil {
		return nil, errInvalidLib
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	cdata := C.CBytes(data)
	free := func() { C.free(cdata) }

	var face C.FT_Face
	if err := getErr(C.FT_New_Memory_Face(l.ptr, (*C.uchar)(cdata), C.long(len(data)), C.FT_Long(idx), &face)); err != nil {
		free()
		return nil, err
	}

	return &Face{ptr: face, dealloc: []func(){free}}, nil
}

// NewFaceFromPath creates a new face from the given path.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_new_face
func NewFaceFromPath(l *Library, path string, idx FaceIndex) (*Face, error) {
	if l == nil || l.ptr == nil {
		return nil, errInvalidLib
	}

	var face C.FT_Face
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	if err := getErr(C.FT_New_Face(l.ptr, cpath, C.FT_Long(idx), &face)); err != nil {
		return nil, err
	}

	return &Face{ptr: face}, nil
}

// Free discards the face, as well as all of its child slots and sizes.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_done_face
func (f *Face) Free() error {
	if f == nil || f.ptr == nil {
		return nil
	}

	if err := getErr(C.FT_Done_Face(f.ptr)); err != nil {
		return err
	}
	for _, fn := range f.dealloc {
		fn()
	}
	f.ptr = nil
	return nil
}

// FamilyName returns the face's family name. This is an ASCII string, usually in English, that describes the typeface's
// family (like ‘Times New Roman’, ‘Bodoni’, ‘Garamond’, etc). This is a least common denominator used to list fonts.
// Some formats (TrueType & OpenType) provide localized and Unicode versions of this string.
// Applications should use the format-specific interface to access them.
// The returned value can be empty (e.g., in fonts embedded in a PDF file).
//
// In case the font doesn't provide a specific family name entry, FreeType tries to synthesize one, deriving it from
// other name entries.
func (f *Face) FamilyName() string {
	if f == nil || f.ptr == nil {
		return ""
	}
	return C.GoString(f.ptr.family_name)
}

// StyleName returns the face's style name. This is an ASCII string, usually in English, that describes the typeface's
// style (like ‘Italic’, ‘Bold’, ‘Condensed’, etc).
// Not all font formats provide a style name.
// Some formats provide localized and Unicode versions of this string.
// Applications should use the format-specific interface to access them.
func (f *Face) StyleName() string {
	if f == nil || f.ptr == nil {
		return ""
	}
	return C.GoString(f.ptr.style_name)
}
