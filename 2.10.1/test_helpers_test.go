package freetype2

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/flga/freetype2/fixed"
	"github.com/flga/freetype2/internal/deep"
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
func twemojiMozilla() (testface, error) {
	return openFace(testdata("twemoji-colr", "TwemojiMozilla.ttf"))
}
func notoColorEmoji() (testface, error) {
	return openFace(testdata("noto-color-emoji", "NotoColorEmoji.ttf"))
}
func chromacheckColr() (testface, error) {
	return openFace(testdata("chromacheck", "chromacheck-colr.woff"))
}
func nimbusMono() (testface, error) {
	return openFace(testdata("nimbus", "NimbusMonoPS-Regular.pfa"))
}
func gohuBdf() (testface, error) {
	return openFace(testdata("gohu", "gohufont-11.bdf"))
}
func gohuPcf() (testface, error) {
	return openFace(testdata("gohu", "gohufont-11.pcf"))
}
func amelia() (testface, error) {
	return openFace(testdata("amelia", "Amelia.pfr"))
}
func bitout() (testface, error) {
	return openFace(testdata("bitout", "bitout.fon"))
}
func faceFromPath(p string) func() (testface, error) {
	return func() (testface, error) { return openFace(testdata(p)) }
}

type testglyph struct {
	Glyph
	f testface
	l *Library
}

func (g testglyph) Free() {
	if g.Glyph != nil {
		g.Glyph.Free()
	}
	g.f.Free()
	g.l.Free()
}

func newTestGlyph(facefn func() (testface, error), char rune, flags LoadFlag, size fixed.Int26_6, dpi uint) func() (testglyph, error) {
	return func() (testglyph, error) {
		face, err := facefn()
		if err != nil {
			return testglyph{}, err
		}

		if face.Flags()&FaceFlagFixedSizes > 0 {
			if err := face.SelectSize(0); err != nil {
				face.Free()
				return testglyph{}, err
			}
		} else {
			if err := face.SetCharSize(size, size, dpi, dpi); err != nil {
				face.Free()
				return testglyph{}, err
			}
		}

		if err := face.LoadChar(char, flags); err != nil {
			face.Free()
			return testglyph{}, err
		}

		slot := face.GlyphSlot()
		got, err := slot.Glyph()
		if err != nil {
			face.Free()
			return testglyph{}, err
		}

		return testglyph{Glyph: got}, nil
	}
}

func newEmptyGlyph(format GlyphFormat) func() (testglyph, error) {
	return func() (testglyph, error) {
		l, err := NewLibrary()
		if err != nil {
			return testglyph{}, err
		}

		g, err := l.NewGlyph(format)
		if err != nil {
			l.Free()
			return testglyph{}, err
		}

		return testglyph{Glyph: g, l: l}, nil
	}
}

func newNilGlyph() (testglyph, error) {
	return testglyph{}, nil
}

func newZeroGlyph(typ ...GlyphFormat) func() (testglyph, error) {
	return func() (testglyph, error) {
		t := GlyphFormatBitmap
		if len(typ) > 0 {
			t = typ[0]
		}
		switch t {
		case GlyphFormatBitmap:
			return testglyph{Glyph: &BitmapGlyph{}}, nil
		case GlyphFormatOutline:
			return testglyph{Glyph: &OutlineGlyph{}}, nil
		default:
			return testglyph{}, nil
		}
	}
}

func diff(a, b interface{}) []string {
	return deep.Equal(a, b, deep.WithCompareUnexportedFields())
}
