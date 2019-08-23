package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_BDF_H
//
// const char* bdfPropAtom(BDF_PropertyRec v) { return v.u.atom; }
// FT_Int32 bdfPropInteger(BDF_PropertyRec v) { return v.u.integer; }
// FT_UInt32 bdfPropCardinal(BDF_PropertyRec v) { return v.u.cardinal; }
//
import "C"
import (
	"unsafe"
)

// BDFPropertyType is a list of BDF property types.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bdf_fonts.html#bdf_propertytype
type BDFPropertyType uint

const (
	// BDFPropertyTypeNone value 0 is used to indicate a missing property.
	BDFPropertyTypeNone BDFPropertyType = C.BDF_PROPERTY_TYPE_NONE
	// BDFPropertyTypeAtom property is a string atom.
	BDFPropertyTypeAtom BDFPropertyType = C.BDF_PROPERTY_TYPE_ATOM
	// BDFPropertyTypeInteger property is a 32-bit signed integer.
	BDFPropertyTypeInteger BDFPropertyType = C.BDF_PROPERTY_TYPE_INTEGER
	// BDFPropertyTypeCardinal property is a 32-bit unsigned integer.
	BDFPropertyTypeCardinal BDFPropertyType = C.BDF_PROPERTY_TYPE_CARDINAL
)

// BDFProperty models a given BDF/PCF property.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bdf_fonts.html#bdf_propertyrec
type BDFProperty struct {
	Type BDFPropertyType
	// The atom string, if Type is BDFPropertyTypeAtom.
	Atom string
	// A signed integer, if Type is BDFPropertyTypeInteger.
	Integer int32
	// An unsigned integer, if Type is BDFPropertyTypeCardinal.
	Cardinal uint32
}

// BDFCharsetID retrieves a BDF font character set identity, according to the
// BDF specification.
//
// This function only works with BDF faces, returning an error otherwise.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bdf_fonts.html#ft_get_bdf_charset_id
func (f *Face) BDFCharsetID() (encoding, registry string, err error) {
	if f == nil || f.ptr == nil {
		return "", "", ErrInvalidFaceHandle
	}

	var acharsetEncoding, acharsetRegistry *C.char
	if err := getErr(C.FT_Get_BDF_Charset_ID(f.ptr, &acharsetEncoding, &acharsetRegistry)); err != nil {
		return "", "", err
	}

	return C.GoString(acharsetEncoding), C.GoString(acharsetRegistry), nil
}

// BDFProperty retrieves a BDF property from a BDF or PCF font file.
//
// This function works with BDF and PCF fonts. It returns an error otherwise.
// It also returns an error if the property is not in the font.
//
// A ‘property’ is a either key-value pair within the
// STARTPROPERTIES ... ENDPROPERTIES block of a BDF font or a key-value pair
// from the info.props array within a FontRec structure of a PCF font.
//
// Integer properties are always stored as ‘signed’ within PCF fonts;
// consequently, BDFPropertyTypeCardinal is a possible return value for BDF
// fonts only.
//
// In case of error, BDFProperty.Type is always set to BDFPropertyTypeNone.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bdf_fonts.html#ft_get_bdf_property
func (f *Face) BDFProperty(name string) (BDFProperty, error) {
	if f == nil || f.ptr == nil {
		return BDFProperty{}, ErrInvalidFaceHandle
	}

	cname := C.CString(name)
	defer free(unsafe.Pointer(cname))

	var aproperty C.BDF_PropertyRec
	if err := getErr(C.FT_Get_BDF_Property(f.ptr, cname, &aproperty)); err != nil {
		return BDFProperty{}, err
	}

	ret := BDFProperty{
		Type: BDFPropertyType(aproperty._type),
	}

	switch ret.Type {
	case BDFPropertyTypeAtom:
		ret.Atom = C.GoString(C.bdfPropAtom(aproperty))
	case BDFPropertyTypeInteger:
		ret.Integer = int32(C.bdfPropInteger(aproperty))
	case BDFPropertyTypeCardinal:
		ret.Cardinal = uint32(C.bdfPropCardinal(aproperty))
	}

	return ret, nil
}
