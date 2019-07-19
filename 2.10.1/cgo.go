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
// #include <stdlib.h>
// #include <ft2build.h>
// #include FT_FREETYPE_H
// const char* getFreetypeError(FT_Error err)
// {
//     #undef __FTERRORS_H__
//     #define FT_ERRORDEF( e, v, s )  case e: return s;
//     #define FT_ERROR_START_LIST     switch (err) {
//     #define FT_ERROR_END_LIST       }
//     #include FT_ERRORS_H
//     return "Unknown error";
// }
import "C"
import (
	"errors"
)

func getErr(code C.int) error {
	if code == 0 {
		return nil
	}
	return errors.New(C.GoString(C.getFreetypeError(code)))
}
