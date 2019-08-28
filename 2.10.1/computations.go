package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_GLYPH_H
// #include FT_TRIGONOMETRY_H
import "C"
import "github.com/flga/freetype2/fixed"

// Transform transforms a single vector through a 2x2 matrix
//
// The result is undefined if either vector or matrix is invalid.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_transform
func (v Vector16_16) Transform(m Matrix) Vector16_16 {
	cv := C.FT_Vector{
		x: C.long(v.X),
		y: C.long(v.Y),
	}
	cm := C.FT_Matrix{
		xx: C.long(m.Xx),
		xy: C.long(m.Xy),
		yx: C.long(m.Yx),
		yy: C.long(m.Yy),
	}

	C.FT_Vector_Transform(&cv, &cm)
	return Vector16_16{
		X: fixed.Int16_16(cv.x),
		Y: fixed.Int16_16(cv.y),
	}
}

// Mul perform the matrix operation a*b.
//
// The result is undefined if either a or b is zero.
//
// Since the function uses wrap-around arithmetic, results become meaningless if
// the arguments are very large.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_matrix_multiply
func (a Matrix) Mul(b Matrix) Matrix {
	ca := C.FT_Matrix{
		xx: C.long(a.Xx),
		xy: C.long(a.Xy),
		yx: C.long(a.Yx),
		yy: C.long(a.Yy),
	}
	cb := C.FT_Matrix{
		xx: C.long(b.Xx),
		xy: C.long(b.Xy),
		yx: C.long(b.Yx),
		yy: C.long(b.Yy),
	}

	C.FT_Matrix_Multiply(&ca, &cb)
	return Matrix{
		Xx: fixed.Int16_16(cb.xx),
		Xy: fixed.Int16_16(cb.xy),
		Yx: fixed.Int16_16(cb.yx),
		Yy: fixed.Int16_16(cb.yy),
	}
}

// Invert inverts the matrix.
// It returns an error if it can't be inverted.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_matrix_invert
func (a Matrix) Invert() (Matrix, error) {
	ca := C.FT_Matrix{
		xx: C.long(a.Xx),
		xy: C.long(a.Xy),
		yx: C.long(a.Yx),
		yy: C.long(a.Yy),
	}
	if err := getErr(C.FT_Matrix_Invert(&ca)); err != nil {
		return Matrix{}, err
	}

	return Matrix{
		Xx: fixed.Int16_16(ca.xx),
		Xy: fixed.Int16_16(ca.xy),
		Yx: fixed.Int16_16(ca.yx),
		Yy: fixed.Int16_16(ca.yy),
	}, nil
}

// Angle is used to model angle values in FreeType. Note that the angle is a
// 16.16 fixed-point value expressed in degrees.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_angle
type Angle fixed.Int16_16

const (
	// AnglePI is the angle pi expressed in FT_Angle units.
	AnglePI Angle = C.FT_ANGLE_PI
	// Angle2PI is the angle 2*pi expressed in FT_Angle units.
	Angle2PI Angle = C.FT_ANGLE_2PI
	// AnglePI2 is the angle pi/2 expressed in FT_Angle units.
	AnglePI2 Angle = C.FT_ANGLE_PI2
	// AnglePI4 is the angle pi/4 expressed in FT_Angle units.
	AnglePI4 Angle = C.FT_ANGLE_PI4
)

// Sin returns the sinus of a given angle in fixed-point format.
// If you need both the sinus and cosinus for a given angle, use Angle.UnitVector.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_sin
func (a Angle) Sin() fixed.Int16_16 {
	return fixed.Int16_16(C.FT_Sin(C.FT_Angle(a)))
}

// Cos returns the cosinus of a given angle in fixed-point format.
// If you need both the sinus and cosinus for a given angle, use Angle.UnitVector.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_cos
func (a Angle) Cos() fixed.Int16_16 {
	return fixed.Int16_16(C.FT_Cos(C.FT_Angle(a)))
}

// Tan returns the tangent of a given angle in fixed-point format.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_tan
func (a Angle) Tan() fixed.Int16_16 {
	return fixed.Int16_16(C.FT_Tan(C.FT_Angle(a)))
}

// Atan2 returns the arc-tangent corresponding to a given vector (x,y) in the 2d
// plane.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_atan2
func Atan2(x, y fixed.Int16_16) Angle {
	return Angle(C.FT_Atan2(C.long(x), C.long(y)))
}

// Diff returns the difference between two angles. The result is always
// constrained to the ]-PI..PI] interval.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_angle_diff
func (a Angle) Diff(b Angle) Angle {
	return Angle(C.FT_Angle_Diff(C.FT_Angle(a), C.FT_Angle(b)))
}

// UnitVector returns the unit vector corresponding to a given angle, ie
// x: cos(angle), y: sin(angle).
//
// This function is useful to retrieve both the sinus and cosinus of a given
// angle quickly.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_unit
func (a Angle) UnitVector() Vector {
	var vec C.FT_Vector
	C.FT_Vector_Unit(&vec, C.FT_Angle(a))
	return Vector{
		X: Pos(vec.x),
		Y: Pos(vec.y),
	}
}

// Rotate rotates the vector by the given angle.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_rotate
func (v Vector) Rotate(a Angle) Vector {
	cv := C.FT_Vector{
		x: C.long(v.X),
		y: C.long(v.Y),
	}

	C.FT_Vector_Rotate(&cv, C.FT_Angle(a))
	return Vector{
		X: Pos(cv.x),
		Y: Pos(cv.y),
	}
}

// Length returns the length of a given vector, expressed in the same units as
// the original vector coordinates.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_length
func (v Vector) Length() Pos {
	cv := C.FT_Vector{
		x: C.long(v.X),
		y: C.long(v.Y),
	}

	return Pos(C.FT_Vector_Length(&cv))
}

// Polarize computes both the length and angle of the vector.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_polarize
func (v Vector) Polarize() (length Pos, angle Angle) {
	cv := C.FT_Vector{
		x: C.long(v.X),
		y: C.long(v.Y),
	}

	var clength C.FT_Fixed
	var cangle C.FT_Angle
	C.FT_Vector_Polarize(&cv, &clength, &cangle)
	return Pos(clength), Angle(cangle)
}

// VectorFromPolar computes vector coordinates from a length and angle.
func VectorFromPolar(length fixed.Int16_16, angle Angle) Vector {
	var cv C.FT_Vector
	C.FT_Vector_From_Polar(&cv, C.FT_Fixed(length), C.FT_Angle(angle))
	return Vector{
		X: Pos(cv.x),
		Y: Pos(cv.y),
	}
}
