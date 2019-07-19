package freetype2

import (
	"io"
	"path/filepath"
)

func testdata(parts ...string) string {
	d := []string{"..", "testdata"}
	d = append(d, parts...)
	return filepath.Join(d...)
}

type zeroReader struct{}

func (n zeroReader) Read(_ []byte) (int, error) {
	return 0, io.EOF
}
