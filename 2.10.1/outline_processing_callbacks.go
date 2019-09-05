package freetype2

// #include <stdlib.h>
// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"
import "unsafe"

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
	d := decomposers.valueOf(user)
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
	d := decomposers.valueOf(user)
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
	d := decomposers.valueOf(user)
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
	d := decomposers.valueOf(user)
	if d == nil {
		return 1
	}
	if err := d.CubicTo(makeVec(control1), makeVec(control2), makeVec(to)); err != nil {
		return 1
	}
	return 0
}

//export OutlineRenderSpanFunc
func OutlineRenderSpanFunc(y, count C.int, cspans *C.FT_Span, user uintptr) {
	fn := spanFuncs.valueOf(user)
	if fn == nil {
		return
	}

	spans := make([]Span, count)
	ptr := (*[(1<<31 - 1) / C.sizeof_FT_Span]C.FT_Span)(unsafe.Pointer(cspans))[:count:count]
	for i := range spans {
		spans[i] = Span{
			X:        int16(ptr[i].x),
			Len:      uint16(ptr[i].len),
			Coverage: uint8(ptr[i].coverage),
		}
	}

	fn(int(y), spans)
}
