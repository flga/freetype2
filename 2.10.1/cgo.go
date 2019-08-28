// Package freetype2 provides Go bindings to the FreeType project
package freetype2

// #cgo windows LDFLAGS: -lfreetype2
// #cgo !static,!windows pkg-config: freetype2
//
// #cgo linux,386,static CFLAGS: -I${SRCDIR}/linux_386/include/freetype2 -Werror -Wall -Wextra -Wno-unused-parameter
// #cgo linux,386,static,!harfbuzz LDFLAGS: -L${SRCDIR}/linux_386/lib -lfreetype -lbz2 -lpng16 -lz -lm
// #cgo linux,386,static,harfbuzz LDFLAGS: -L${SRCDIR}/linux_386/lib -lfreetypehb -lharfbuzz -lfreetypehb -lbz2 -lpng16 -lz -lm
// #cgo linux,386,static,harfbuzz,subset LDFLAGS: -L${SRCDIR}/linux_386/lib -lfreetypehb -lharfbuzz -lharfbuzz-subset -lbz2 -lpng16 -lz -lm
//
// #cgo linux,amd64,static CFLAGS: -I${SRCDIR}/linux_amd64/include/freetype2 -Werror -Wall -Wextra -Wno-unused-parameter
// #cgo linux,amd64,static,!harfbuzz LDFLAGS: -L${SRCDIR}/linux_amd64/lib -lfreetype -lbz2 -lpng16 -lz -lm
// #cgo linux,amd64,static,harfbuzz LDFLAGS: -L${SRCDIR}/linux_amd64/lib -lfreetypehb -lharfbuzz -lfreetypehb -lbz2 -lpng16 -lz -lm
// #cgo linux,amd64,static,harfbuzz,subset LDFLAGS: -L${SRCDIR}/linux_amd64/lib -lfreetypehb -lharfbuzz -lharfbuzz-subset -lbz2 -lpng16 -lz -lm
//
// #cgo darwin,386,static CFLAGS: -I${SRCDIR}/darwin_386/include/freetype2 -Werror -Wall -Wextra -Wno-unused-parameter
// #cgo darwin,386,static,!harfbuzz LDFLAGS: -L${SRCDIR}/darwin_386/lib -lfreetype -lbz2 -lpng16 -lz -lm
// #cgo darwin,386,static,harfbuzz LDFLAGS: -L${SRCDIR}/darwin_386/lib -lfreetypehb -lharfbuzz -lfreetypehb -lbz2 -lpng16 -lz -lm
// #cgo darwin,386,static,harfbuzz,subset LDFLAGS: -L${SRCDIR}/darwin_386/lib -lfreetypehb -lharfbuzz -lharfbuzz-subset -lbz2 -lpng16 -lz -lm
//
// #cgo darwin,amd64,static CFLAGS: -I${SRCDIR}/darwin_amd64/include/freetype2 -Werror -Wall -Wextra -Wno-unused-parameter
// #cgo darwin,amd64,static,!harfbuzz LDFLAGS: -L${SRCDIR}/darwin_amd64/lib -lfreetype -lbz2 -lpng16 -lz -lm
// #cgo darwin,amd64,static,harfbuzz LDFLAGS: -L${SRCDIR}/darwin_amd64/lib -lfreetypehb -lharfbuzz -lfreetypehb -lbz2 -lpng16 -lz -lm
// #cgo darwin,amd64,static,harfbuzz,subset LDFLAGS: -L${SRCDIR}/darwin_amd64/lib -lfreetypehb -lharfbuzz -lharfbuzz-subset -lbz2 -lpng16 -lz -lm
//
// #include <stdlib.h>
// #include <ft2build.h>
// #include FT_FREETYPE_H
//
// FT_UInt32* test_makeList(int elems) {
// 	if (elems <= 0) {
// 		return NULL;
// 	}
//
// 	FT_UInt32* ptr = (FT_UInt32*)calloc(elems+1, sizeof(FT_UInt32));
// 	for(int i = 0; i < elems; ++i) {
// 		ptr[i] = i+1;
// 	}
//
// 	return ptr;
// }
import "C"

import (
	"sync"
	"unsafe"
)

var free = func(v unsafe.Pointer) {
	C.free(v)
}

var mockFreeMu sync.Mutex

func mockFree(fn func()) (restore func()) {
	mockFreeMu.Lock()
	orig := free
	free = func(v unsafe.Pointer) {
		fn()
		orig(v)
	}

	return func() {
		free = orig
		mockFreeMu.Unlock()
	}
}

func testMakeList(n int) *C.FT_UInt32 {
	return C.test_makeList(C.int(n))
}

func sliceFromZeroTerminatedUint32(head *C.FT_UInt32) []rune {
	if head == nil {
		return nil
	}

	var ret []rune
	ptr := (*[(1<<31 - 1) / C.sizeof_FT_UInt32]C.FT_UInt32)(unsafe.Pointer(head))[:]
	for i := 0; ptr[i] != 0; i++ {
		ret = append(ret, rune(ptr[i]))
	}
	return ret
}
