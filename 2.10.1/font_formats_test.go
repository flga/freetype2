package freetype2

import "testing"

func TestFace_FontFormat(t *testing.T) {
	tests := []struct {
		path string
		want FontFormat
	}{
		{path: testdata("noto-color-emoji", "NotoColorEmoji.ttf"), want: FontFormatTrueType},
		{path: testdata("amelia", "Amelia.pfr"), want: FontFormatPFR},
		{path: testdata("bungee", "BungeeColor-Regular_colr_Windows.ttf"), want: FontFormatTrueType},
		{path: testdata("bungee", "BungeeColor-Regular_sbix_MacOS.ttf"), want: FontFormatTrueType},
		{path: testdata("bungee", "BungeeColor-Regular_svg.ttf"), want: FontFormatTrueType},
		{path: testdata("bungee", "Bungee-Regular.otf"), want: FontFormatCFF},
		{path: testdata("bungee", "BungeeColor-Regular_svg.woff"), want: FontFormatTrueType},
		{path: testdata("bungee", "BungeeLayers-Regular.otf"), want: FontFormatCFF},
		{path: testdata("chromacheck", "chromacheck-cbdt.woff"), want: FontFormatTrueType},
		{path: testdata("chromacheck", "chromacheck-sbix.woff"), want: FontFormatTrueType},
		{path: testdata("chromacheck", "chromacheck-colr.woff"), want: FontFormatTrueType},
		{path: testdata("chromacheck", "chromacheck-svg.woff"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Italic.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Mono-Italic.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Mono-Bold-Italic.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Smallcaps-Italic.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Medium.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Bold.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Mono-Bold.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Mono.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Regular.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Smallcaps.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Medium-Italic.ttf"), want: FontFormatTrueType},
		{path: testdata("go", "Go-Bold-Italic.ttf"), want: FontFormatTrueType},
		{path: testdata("noto sans jp", "NotoSansJP-Light.otf"), want: FontFormatCFF},
		{path: testdata("noto sans jp", "NotoSansJP-Bold.otf"), want: FontFormatCFF},
		{path: testdata("noto sans jp", "NotoSansJP-Regular.otf"), want: FontFormatCFF},
		{path: testdata("noto sans jp", "NotoSansJP-Black.otf"), want: FontFormatCFF},
		{path: testdata("noto sans jp", "NotoSansJP-Medium.otf"), want: FontFormatCFF},
		{path: testdata("noto sans jp", "NotoSansJP-Thin.otf"), want: FontFormatCFF},
		{path: testdata("arimo", "Arimo-Regular.ttf"), want: FontFormatTrueType},
		{path: testdata("bitout", "bitout.fon"), want: FontFormatWindowsFNT},
		{path: testdata("nimbus", "NimbusMonoPS-Regular.pfa"), want: FontFormatType1},
		{path: testdata("gohu", "gohufont-11.bdf"), want: FontFormatBDF},
		{path: testdata("gohu", "gohufont-11.pcf"), want: FontFormatPCF},
		{path: testdata("twemoji-colr", "TwemojiMozilla.ttf"), want: FontFormatTrueType},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			l, err := NewLibrary()
			if err != nil {
				t.Fatalf("unable to load lib: %s", err)
			}
			defer l.Free()

			face, err := l.NewFaceFromPath(tt.path, 0, 0)
			if err != nil {
				t.Fatalf("unable to load face: %s", err)
			}
			defer face.Free()

			if got := face.FontFormat(); got != tt.want {
				t.Errorf("Face.FontFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
