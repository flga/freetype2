package freetype2

// #include <stdlib.h>
// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

func makeVec(v *C.FT_Vector) Vector {
	if v == nil {
		return Vector{}
	}

	return Vector{
		X: Pos(v.x),
		Y: Pos(v.y),
	}
}

//export OutlineMoveToCallback
func OutlineMoveToCallback(to *C.FT_Vector, user uintptr) C.int {
	d := decomposers.get(user)
	if d == nil {
		return 1
	}
	if err := d.MoveTo(makeVec(to)); err != nil {
		return 1
	}
	return 0
}

//export OutlineLineToCallback
func OutlineLineToCallback(to *C.FT_Vector, user uintptr) C.int {
	d := decomposers.get(user)
	if d == nil {
		return 1
	}
	if err := d.LineTo(makeVec(to)); err != nil {
		return 1
	}
	return 0
}

//export OutlineConicToCallback
func OutlineConicToCallback(control, to *C.FT_Vector, user uintptr) C.int {
	d := decomposers.get(user)
	if d == nil {
		return 1
	}
	if err := d.ConicTo(makeVec(control), makeVec(to)); err != nil {
		return 1
	}
	return 0
}

//export OutlineCubicToCallback
func OutlineCubicToCallback(control1, control2, to *C.FT_Vector, user uintptr) C.int {
	d := decomposers.get(user)
	if d == nil {
		return 1
	}
	if err := d.CubicTo(makeVec(control1), makeVec(control2), makeVec(to)); err != nil {
		return 1
	}
	return 0
}
