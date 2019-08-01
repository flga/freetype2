package freetype2

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/go-test/deep"
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

type testface struct {
	*Face
	l *Library
}

func (t testface) Free() error {
	t.Face.Free()
	t.l.Free()
	return nil
}

func openFace(path string) (testface, error) {
	l, err := NewLibrary()
	if err != nil {
		return testface{}, fmt.Errorf("unable to initialize library: %v", err)
	}
	face, err := l.NewFaceFromPath(path, 0, 0)
	if err != nil {
		l.Free()
		return testface{}, fmt.Errorf("unable to open font: %s", err)
	}
	return testface{face, l}, nil
}
func nilFace() (testface, error) {
	return testface{}, nil
}
func goRegular() (testface, error) {
	return openFace(testdata("go", "Go-Regular.ttf"))
}
func goBold() (testface, error) {
	return openFace(testdata("go", "Go-Bold.ttf"))
}
func goItalic() (testface, error) {
	return openFace(testdata("go", "Go-Italic.ttf"))
}
func goBoldItalic() (testface, error) {
	return openFace(testdata("go", "Go-Bold-Italic.ttf"))
}
func goMono() (testface, error) {
	return openFace(testdata("go", "Go-Mono.ttf"))
}
func bungeeColorWin() (testface, error) {
	return openFace(testdata("bungee", "BungeeColor-Regular_colr_Windows.ttf"))
}
func bungeeColorMac() (testface, error) {
	return openFace(testdata("bungee", "BungeeColor-Regular_sbix_MacOS.ttf"))
}
func bungeeLayersReg() (testface, error) {
	return openFace(testdata("bungee", "BungeeLayers-Regular.otf"))
}
func notoSansJpReg() (testface, error) {
	return openFace(testdata("noto sans jp", "NotoSansJP-Regular.otf"))
}
func notoSansJpBold() (testface, error) {
	return openFace(testdata("noto sans jp", "NotoSansJP-Bold.otf"))
}
func arimoRegular() (testface, error) {
	return openFace(testdata("arimo", "Arimo-Regular.ttf"))
}

func diff(a, b interface{}) []string {
	orig := deep.CompareUnexportedFields
	defer func() {
		deep.CompareUnexportedFields = orig
	}()
	deep.CompareUnexportedFields = true
	return deep.Equal(a, b)
}
