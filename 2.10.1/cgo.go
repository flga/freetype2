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
import "C"

import (
	"sync"
	"unsafe"
)

var free = func(v unsafe.Pointer) {
	C.free(v)
}

type freeBehaviour bool

const (
	actuallyFreeItAfter freeBehaviour = true
	iWannaFreeItMyself  freeBehaviour = false
)

var mockFreeMu sync.Mutex

func mockFree(fn func(v unsafe.Pointer), actuallyFreeIt freeBehaviour) (restore func()) {
	mockFreeMu.Lock()
	orig := free
	free = func(v unsafe.Pointer) {
		fn(v)
		if actuallyFreeIt {
			orig(v)
		}
	}

	return func() {
		free = orig
		mockFreeMu.Unlock()
	}
}
