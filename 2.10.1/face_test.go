package freetype2

import (
	"testing"

	"github.com/flga/freetype2/2.10.1/fixed"
	"github.com/flga/freetype2/2.10.1/truetype"
)

func TestFaceFlags_String(t *testing.T) {
	var x FaceFlag
	if got, want := x.String(), ""; got != want {
		t.Errorf("FaceFlags.String(0) = %v, want %v", got, want)
	}

	x = FaceFlagColor
	if got, want := x.String(), "Color"; got != want {
		t.Errorf("FaceFlags.String(FaceFlagColor) = %v, want %v", got, want)
	}

	x = FaceFlagKerning | FaceFlagCidKeyed
	if got, want := x.String(), "Kerning|CidKeyed"; got != want {
		t.Errorf("FaceFlags.String(FaceFlagKerning | FaceFlagCidKeyed) = %v, want %v", got, want)
	}

	x = FaceFlagVertical | FaceFlagMultipleMasters | FaceFlagHinter
	if got, want := x.String(), "Vertical|MultipleMasters|Hinter"; got != want {
		t.Errorf("FaceFlags.String(FaceFlagVertical | FaceFlagMultipleMasters | FaceFlagHinter) = %v, want %v", got, want)
	}

	x = FaceFlagScalable | FaceFlagFixedSizes | FaceFlagFixedWidth | FaceFlagSfnt | FaceFlagHorizontal |
		FaceFlagVertical | FaceFlagKerning | FaceFlagMultipleMasters | FaceFlagGlyphNames | FaceFlagHinter |
		FaceFlagCidKeyed | FaceFlagTricky | FaceFlagColor | FaceFlagVariation
	if got, want := x.String(), "Scalable|FixedSizes|FixedWidth|Sfnt|Horizontal|Vertical|Kerning|MultipleMasters|GlyphNames|Hinter|CidKeyed|Tricky|Color|Variation"; got != want {
		t.Errorf("FaceFlags.String(FaceFlagScalable | FaceFlagFixedSizes | FaceFlagFixedWidth | FaceFlagSfnt | FaceFlagHorizontal | FaceFlagVertical | FaceFlagKerning | FaceFlagMultipleMasters | FaceFlagGlyphNames | FaceFlagHinter | FaceFlagCidKeyed | FaceFlagTricky | FaceFlagColor | FaceFlagVariation) = %v, want %v", got, want)
	}
}

func TestStyleFlags_String(t *testing.T) {
	var x StyleFlag

	if got, want := x.String(), ""; got != want {
		t.Errorf("StyleFlags.String(0) = %v, want %v", got, want)
	}

	x = StyleFlagItalic
	if got, want := x.String(), "Italic"; got != want {
		t.Errorf("StyleFlags.String(StyleFlagItalic) = %v, want %v", got, want)
	}

	x = StyleFlagBold
	if got, want := x.String(), "Bold"; got != want {
		t.Errorf("StyleFlags.String(StyleFlagBold) = %v, want %v", got, want)
	}

	x = StyleFlagItalic | StyleFlagBold
	if got, want := x.String(), "Italic|Bold"; got != want {
		t.Errorf("StyleFlags.String(StyleFlagItalic | StyleFlagBold) = %v, want %v", got, want)
	}
}

func TestFaceFree(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to init lib: %s", err)
	}
	defer l.Free()

	var called bool
	sentinel := func() { called = true }

	f, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open face: %s", err)
	}
	f.dealloc = append(f.dealloc, sentinel)

	if err := f.Free(); err != nil {
		t.Fatalf("unable to free face: %s", err)
	}
	if f.ptr != nil {
		t.Fatalf("Free should set ptr to nil")
	}
	if called != true {
		t.Fatalf("Free should call every function in dealoc")
	}
	if err := f.Free(); err != nil {
		t.Fatalf("Free on an already freed face should be a noop, got: %s", err)
	}
}

func TestFaceProps(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to init lib: %s", err)
	}
	defer l.Free()

	goRegular, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goRegular.Free()

	goBold, err := l.NewFaceFromPath(testdata("go", "Go-Bold.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goBold.Free()

	goItalic, err := l.NewFaceFromPath(testdata("go", "Go-Italic.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goItalic.Free()

	goBoldItalic, err := l.NewFaceFromPath(testdata("go", "Go-Bold-Italic.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goBoldItalic.Free()

	goMono, err := l.NewFaceFromPath(testdata("go", "Go-Mono.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goMono.Free()

	bungeeColorWin, err := l.NewFaceFromPath(testdata("bungee", "BungeeColor-Regular_colr_Windows.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer bungeeColorWin.Free()

	bungeeColorMac, err := l.NewFaceFromPath(testdata("bungee", "BungeeColor-Regular_sbix_MacOS.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer bungeeColorMac.Free()

	bungeeLayersReg, err := l.NewFaceFromPath(testdata("bungee", "BungeeLayers-Regular.otf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer bungeeLayersReg.Free()

	type charmapdata struct {
		format   int
		platform truetype.PlatformID
		encoding truetype.EncodingID
		language truetype.LanguageID
		active   bool
	}

	tests := []struct {
		name               string
		face               *Face
		family             string
		style              string
		numFaces           int
		numNamedInstances  int
		faceIdx            int
		namedIdx           int
		bold               bool
		italic             bool
		sfntWrapped        bool
		scalable           bool
		fixedSize          bool
		horizontal         bool
		vertical           bool
		fixedWidth         bool
		glyphNames         bool
		emSize             int
		globalBBox         BBox
		ascent             int
		descent            int
		textHeight         int
		glyphCount         int
		numCharmaps        int
		charmaps           []charmapdata
		numSizes           int
		avaliableSizes     []BitmapSize
		maxAdvanceWidth    int
		maxAdvanceHeight   int
		underlinePosition  int
		underlineThickness int
	}{
		{
			name:              "goRegular",
			face:              goRegular,
			family:            "Go",
			style:             "Regular",
			numFaces:          1,
			numNamedInstances: 0,
			faceIdx:           0,
			namedIdx:          0,
			bold:              false,
			italic:            false,
			sfntWrapped:       true,
			scalable:          true,
			fixedSize:         false,
			horizontal:        true,
			vertical:          false,
			fixedWidth:        false,
			glyphNames:        true,
			emSize:            2048,
			globalBBox:        BBox{XMin: -440, YMin: -543, XMax: 2160, YMax: 2118},
			ascent:            1935,
			descent:           -432,
			textHeight:        2367,
			glyphCount:        666,
			numCharmaps:       3,
			charmaps: []charmapdata{
				{format: 4, platform: 0, encoding: 3, language: 0, active: false},
				{format: 6, platform: 1, encoding: 0, language: 0, active: false},
				{format: 4, platform: 3, encoding: 1, language: 0, active: true},
			},
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    2240,
			maxAdvanceHeight:   2367,
			underlinePosition:  -300,
			underlineThickness: 50,
		},
		{
			name:              "goBold",
			face:              goBold,
			family:            "Go",
			style:             "Bold",
			numFaces:          1,
			numNamedInstances: 0,
			faceIdx:           0,
			namedIdx:          0,
			bold:              true,
			italic:            false,
			sfntWrapped:       true,
			scalable:          true,
			fixedSize:         false,
			horizontal:        true,
			vertical:          false,
			fixedWidth:        false,
			glyphNames:        true,
			emSize:            2048,
			globalBBox:        BBox{XMin: -452, YMin: -432, XMax: 2190, YMax: 2193},
			ascent:            1935,
			descent:           -432,
			textHeight:        2367,
			glyphCount:        666,
			numCharmaps:       3,
			charmaps: []charmapdata{
				{format: 4, platform: 0, encoding: 3, language: 0, active: false},
				{format: 6, platform: 1, encoding: 0, language: 0, active: false},
				{format: 4, platform: 3, encoding: 1, language: 0, active: true},
			},
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    2283,
			maxAdvanceHeight:   2367,
			underlinePosition:  -300,
			underlineThickness: 100,
		},
		{
			name:              "goItalic",
			face:              goItalic,
			family:            "Go",
			style:             "Italic",
			numFaces:          1,
			numNamedInstances: 0,
			faceIdx:           0,
			namedIdx:          0,
			bold:              false,
			italic:            true,
			sfntWrapped:       true,
			scalable:          true,
			fixedSize:         false,
			horizontal:        true,
			vertical:          false,
			fixedWidth:        false,
			glyphNames:        true,
			emSize:            2048,
			globalBBox:        BBox{XMin: -436, YMin: -543, XMax: 2276, YMax: 2118},
			ascent:            1935,
			descent:           -432,
			textHeight:        2367,
			glyphCount:        666,
			numCharmaps:       3,
			charmaps: []charmapdata{
				{format: 4, platform: 0, encoding: 3, language: 0, active: false},
				{format: 6, platform: 1, encoding: 0, language: 0, active: false},
				{format: 4, platform: 3, encoding: 1, language: 0, active: true},
			},
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    2262,
			maxAdvanceHeight:   2367,
			underlinePosition:  -300,
			underlineThickness: 50,
		},
		{
			name:              "goBoldItalic",
			face:              goBoldItalic,
			family:            "Go",
			style:             "Bold Italic",
			numFaces:          1,
			numNamedInstances: 0,
			faceIdx:           0,
			namedIdx:          0,
			bold:              true,
			italic:            true,
			sfntWrapped:       true,
			scalable:          true,
			fixedSize:         false,
			horizontal:        true,
			vertical:          false,
			fixedWidth:        false,
			glyphNames:        true,
			emSize:            2048,
			globalBBox:        BBox{XMin: -459, YMin: -432, XMax: 2300, YMax: 2193},
			ascent:            1935,
			descent:           -432,
			textHeight:        2367,
			glyphCount:        666,
			numCharmaps:       3,
			charmaps: []charmapdata{
				{format: 4, platform: 0, encoding: 3, language: 0, active: false},
				{format: 6, platform: 1, encoding: 0, language: 0, active: false},
				{format: 4, platform: 3, encoding: 1, language: 0, active: true},
			},
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    2283,
			maxAdvanceHeight:   2367,
			underlinePosition:  -350,
			underlineThickness: 100,
		},
		{
			name:              "goMono",
			face:              goMono,
			family:            "Go Mono",
			style:             "Regular",
			numFaces:          1,
			numNamedInstances: 0,
			faceIdx:           0,
			namedIdx:          0,
			bold:              false,
			italic:            false,
			sfntWrapped:       true,
			scalable:          true,
			fixedSize:         false,
			horizontal:        true,
			vertical:          false,
			fixedWidth:        true,
			glyphNames:        true,
			emSize:            2048,
			globalBBox:        BBox{XMin: 0, YMin: -432, XMax: 1229, YMax: 2227},
			ascent:            1935,
			descent:           -432,
			textHeight:        2367,
			glyphCount:        666,
			numCharmaps:       3,
			charmaps: []charmapdata{
				{format: 4, platform: 0, encoding: 3, language: 0, active: false},
				{format: 6, platform: 1, encoding: 0, language: 0, active: false},
				{format: 4, platform: 3, encoding: 1, language: 0, active: true},
			},
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    1229,
			maxAdvanceHeight:   2367,
			underlinePosition:  -300,
			underlineThickness: 50,
		},
		{
			name:              "bungeeColorWin",
			face:              bungeeColorWin,
			family:            "Bungee Color",
			style:             "Regular",
			numFaces:          1,
			numNamedInstances: 0,
			faceIdx:           0,
			namedIdx:          0,
			bold:              false,
			italic:            false,
			sfntWrapped:       true,
			scalable:          true,
			fixedSize:         false,
			horizontal:        true,
			vertical:          false,
			fixedWidth:        false,
			glyphNames:        true,
			emSize:            1000,
			globalBBox:        BBox{XMin: -49, YMin: -362, XMax: 1393, YMax: 1138},
			ascent:            860,
			descent:           -140,
			textHeight:        1200,
			glyphCount:        868,
			numCharmaps:       3,
			charmaps: []charmapdata{
				{format: 4, platform: 0, encoding: 3, language: 0, active: false},
				{format: 6, platform: 1, encoding: 0, language: 0, active: false},
				{format: 4, platform: 3, encoding: 1, language: 0, active: true},
			},
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    1417,
			maxAdvanceHeight:   1200,
			underlinePosition:  0,
			underlineThickness: 0,
		},
		{
			name:              "bungeeColorMac",
			face:              bungeeColorMac,
			family:            "Bungee Color",
			style:             "Regular",
			numNamedInstances: 0,
			numFaces:          1,
			faceIdx:           0,
			namedIdx:          0,
			bold:              false,
			italic:            false,
			sfntWrapped:       true,
			scalable:          false,
			fixedSize:         true,
			horizontal:        true,
			vertical:          false,
			fixedWidth:        false,
			glyphNames:        true,
			emSize:            0,
			globalBBox:        BBox{},
			ascent:            0,
			descent:           0,
			textHeight:        0,
			glyphCount:        868,
			numCharmaps:       3,
			charmaps: []charmapdata{
				{format: 4, platform: 0, encoding: 3, language: 0, active: false},
				{format: 6, platform: 1, encoding: 0, language: 0, active: false},
				{format: 4, platform: 3, encoding: 1, language: 0, active: true},
			},
			numSizes: 9,
			avaliableSizes: []BitmapSize{
				{Height: 24, Width: 13, Size: 20 << 6, XPpem: 20 << 6, YPpem: 20 << 6},
				{Height: 38, Width: 21, Size: 32 << 6, XPpem: 32 << 6, YPpem: 32 << 6},
				{Height: 48, Width: 27, Size: 40 << 6, XPpem: 40 << 6, YPpem: 40 << 6},
				{Height: 86, Width: 48, Size: 72 << 6, XPpem: 72 << 6, YPpem: 72 << 6},
				{Height: 115, Width: 64, Size: 96 << 6, XPpem: 96 << 6, YPpem: 96 << 6},
				{Height: 153, Width: 85, Size: 128 << 6, XPpem: 128 << 6, YPpem: 128 << 6},
				{Height: 307, Width: 171, Size: 256 << 6, XPpem: 256 << 6, YPpem: 256 << 6},
				{Height: 614, Width: 342, Size: 512 << 6, XPpem: 512 << 6, YPpem: 512 << 6},
				{Height: 1228, Width: 683, Size: 1024 << 6, XPpem: 1024 << 6, YPpem: 1024 << 6},
			},
			maxAdvanceWidth:    0,
			maxAdvanceHeight:   0,
			underlinePosition:  0,
			underlineThickness: 0,
		},
		{
			name:              "bungeeLayersReg",
			face:              bungeeLayersReg,
			family:            "Bungee Layers",
			style:             "Regular",
			numFaces:          1,
			numNamedInstances: 0,
			faceIdx:           0,
			namedIdx:          0,
			bold:              false,
			italic:            false,
			sfntWrapped:       true,
			scalable:          true,
			fixedSize:         false,
			horizontal:        true,
			vertical:          false,
			fixedWidth:        false,
			glyphNames:        true,
			emSize:            1000,
			globalBBox:        BBox{XMin: -607, YMin: -915, XMax: 1943, YMax: 1635},
			ascent:            860,
			descent:           -140,
			textHeight:        1200,
			glyphCount:        1075,
			numCharmaps:       4,
			charmaps: []charmapdata{
				{format: 4, platform: 0, encoding: 3, language: 0, active: false},
				{format: 6, platform: 1, encoding: 0, language: 0, active: false},
				{format: 4, platform: 3, encoding: 1, language: 0, active: true},
				{format: -1, platform: 7, encoding: 0, language: 0, active: false},
			},
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    1417,
			maxAdvanceHeight:   1200,
			underlinePosition:  0,
			underlineThickness: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.face.FamilyName(); got != tt.family {
				t.Errorf("FamilyName() want %s, got %s", tt.family, got)
			}

			if got := tt.face.StyleName(); got != tt.style {
				t.Errorf("StyleName() want %s, got %s", tt.style, got)
			}

			if got := tt.face.NumFaces(); got != tt.numFaces {
				t.Errorf("NumFaces() want %v, got %v", tt.numFaces, got)
			}

			if got := tt.face.NumNamedInstances(); got != tt.numNamedInstances {
				t.Errorf("NumNamedInstances() want %v, got %v", tt.numNamedInstances, got)
			}

			if got := tt.face.Index(); got != tt.faceIdx {
				t.Errorf("Index() want %v, got %v", tt.faceIdx, got)
			}

			if got := tt.face.NamedIndex(); got != tt.namedIdx {
				t.Errorf("NamedIndex() want %v, got %v", tt.namedIdx, got)
			}

			testFlag := func(v bool, f FaceFlag) {
				if v && !tt.face.HasFlag(f) {
					t.Errorf("Flags() face should have flag %s", f)
				}
				if !v && tt.face.HasFlag(f) {
					t.Errorf("Flags() face should not have flag %s", f)
				}
			}

			testStyleFlag := func(v bool, f StyleFlag) {
				if v && !tt.face.HasStyle(f) {
					t.Errorf("StyleFlags() face should have flag %s", f)
				}
				if !v && tt.face.HasStyle(f) {
					t.Errorf("StyleFlags() face should not have flag %s", f)
				}
			}

			testStyleFlag(tt.bold, StyleFlagBold)
			testStyleFlag(tt.italic, StyleFlagItalic)
			testFlag(tt.sfntWrapped, FaceFlagSfnt)
			testFlag(tt.scalable, FaceFlagScalable)
			testFlag(tt.fixedSize, FaceFlagFixedSizes)
			testFlag(tt.horizontal, FaceFlagHorizontal)
			testFlag(tt.vertical, FaceFlagVertical)
			testFlag(tt.fixedWidth, FaceFlagFixedWidth)
			testFlag(tt.glyphNames, FaceFlagGlyphNames)

			if got := tt.face.UnitsPerEM(); got != tt.emSize {
				t.Errorf("UnitsPerEM() want %v, got %v", tt.emSize, got)
			}

			if got := tt.face.BBox(); got != tt.globalBBox {
				t.Errorf("BBox() want %v, got %v", tt.globalBBox, got)
			}

			if got := tt.face.Ascender(); got != tt.ascent {
				t.Errorf("Ascender() want %v, got %v", tt.ascent, got)
			}

			if got := tt.face.Descender(); got != tt.descent {
				t.Errorf("Descender() want %v, got %v", tt.descent, got)
			}

			if got := tt.face.Height(); got != tt.textHeight {
				t.Errorf("Height() want %v, got %v", tt.textHeight, got)
			}

			if got := tt.face.NumGlyphs(); got != tt.glyphCount {
				t.Errorf("GlyphCount() want %v, got %v", tt.glyphCount, got)
			}

			if got := tt.face.NumCharMaps(); got != tt.numCharmaps {
				t.Errorf("NumCharMaps() want %v, got %v", tt.numCharmaps, got)
			}

			faceCharmaps := tt.face.CharMaps()
			for i, cmap := range tt.charmaps {
				active, _ := tt.face.ActiveCharMap()
				got := charmapdata{
					format:   faceCharmaps[i].Format,
					platform: faceCharmaps[i].PlatformID,
					encoding: faceCharmaps[i].EncodingID,
					language: faceCharmaps[i].Language,
					active:   active == faceCharmaps[i],
				}

				if got != cmap {
					t.Errorf("CharMaps(%d) want %v, got %v", i, cmap, got)
				}
			}

			if got := tt.face.NumFixedSizes(); got != tt.numSizes {
				t.Errorf("NumFixedSizes() want %v, got %v", tt.numSizes, got)
			}

			faceSizes := tt.face.AvailableSizes()
			for i, want := range tt.avaliableSizes {
				if got := faceSizes[i]; want != got {
					t.Errorf("AvailableSizes(%d) want %v, got %v", i, want, got)
				}
			}

			if got := tt.face.MaxAdvanceWidth(); got != tt.maxAdvanceWidth {
				t.Errorf("MaxAdvanceWidth() want %v, got %v", tt.maxAdvanceWidth, got)
			}
			if got := tt.face.MaxAdvanceHeight(); got != tt.maxAdvanceHeight {
				t.Errorf("MaxAdvanceHeight() want %v, got %v", tt.maxAdvanceHeight, got)
			}
			if got := tt.face.UnderlinePosition(); got != tt.underlinePosition {
				t.Errorf("UnderlinePosition() want %v, got %v", tt.underlinePosition, got)
			}
			if got := tt.face.UnderlineThickness(); got != tt.underlineThickness {
				t.Errorf("UnderlineThickness() want %v, got %v", tt.underlineThickness, got)
			}
		})
	}
}

func TestFaceZeroVal(t *testing.T) {
	var f *Face
	if got, want := f.NumFaces(), 0; got != want {
		t.Errorf("NumFaces() = %v, want %v", got, want)
	}
	if got, want := f.Index(), 0; got != want {
		t.Errorf("Index() = %v, want %v", got, want)
	}
	if got, want := f.NamedIndex(), 0; got != want {
		t.Errorf("NamedIndex() = %v, want %v", got, want)
	}
	if got, want := f.Flags(), FaceFlag(0); got != want {
		t.Errorf("Flags() = %v, want %v", got, want)
	}
	if got, want := f.HasFlag(FaceFlagHorizontal), false; got != want {
		t.Errorf("Flags() = %v, want %v", got, want)
	}
	if got, want := f.Style(), StyleFlag(0); got != want {
		t.Errorf("Style() = %v, want %v", got, want)
	}
	if got, want := f.HasStyle(StyleFlagBold), false; got != want {
		t.Errorf("Style() = %v, want %v", got, want)
	}
	if got, want := f.NumNamedInstances(), 0; got != want {
		t.Errorf("NumNamedInstances() = %v, want %v", got, want)
	}
	if got, want := f.NumGlyphs(), 0; got != want {
		t.Errorf("NumGlyphs() = %v, want %v", got, want)
	}
	if got, want := f.FamilyName(), ""; got != want {
		t.Errorf("FamilyName() = %v, want %v", got, want)
	}
	if got, want := f.StyleName(), ""; got != want {
		t.Errorf("StyleName() = %v, want %v", got, want)
	}
	if got, want := f.NumFixedSizes(), 0; got != want {
		t.Errorf("NumFixedSizes() = %v, want %v", got, want)
	}
	if got, want := f.AvailableSizes(), 0; len(got) != want {
		t.Errorf("AvailableSizes() len = %v, want len %v", got, want)
	}
	if got, want := f.NumCharMaps(), 0; got != want {
		t.Errorf("NumCharMaps() = %v, want %v", got, want)
	}
	if got, want := f.CharMaps(), 0; len(got) != want {
		t.Errorf("CharMaps() len = %v, want len %v", got, want)
	}
	if got, want := f.BBox(), (BBox{}); got != want {
		t.Errorf("BBox() = %v, want %v", got, want)
	}
	if got, want := f.UnitsPerEM(), 0; got != want {
		t.Errorf("UnitsPerEM() = %v, want %v", got, want)
	}
	if got, want := f.Ascender(), 0; got != want {
		t.Errorf("Ascender() = %v, want %v", got, want)
	}
	if got, want := f.Descender(), 0; got != want {
		t.Errorf("Descender() = %v, want %v", got, want)
	}
	if got, want := f.Height(), 0; got != want {
		t.Errorf("Height() = %v, want %v", got, want)
	}
	if got, want := f.MaxAdvanceWidth(), 0; got != want {
		t.Errorf("MaxAdvanceWidth() = %v, want %v", got, want)
	}
	if got, want := f.MaxAdvanceHeight(), 0; got != want {
		t.Errorf("MaxAdvanceHeight() = %v, want %v", got, want)
	}
	if got, want := f.UnderlinePosition(), 0; got != want {
		t.Errorf("UnderlinePosition() = %v, want %v", got, want)
	}
	if got, want := f.UnderlineThickness(), 0; got != want {
		t.Errorf("UnderlineThickness() = %v, want %v", got, want)
	}
	if got, want := f.Size(), (Size{}); got != want {
		t.Errorf("Size() = %v, want %v", got, want)
	}
	{
		want, wantOk := CharMap{}, false
		if got, gotOk := f.ActiveCharMap(); got != want || gotOk != wantOk {
			t.Errorf("ActiveCharMap() = %v, %v, want %v, %v", got, gotOk, want, wantOk)
		}
	}
}

func TestFace_SelectCharMap(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %s", err)
	}
	defer l.Free()

	goRegular, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goRegular.Free()

	bungeeLayersReg, err := l.NewFaceFromPath(testdata("bungee", "BungeeLayers-Regular.otf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer bungeeLayersReg.Free()

	tests := []struct {
		name    string
		face    *Face
		enc     Encoding
		want    CharMap
		wantErr error
	}{
		{
			name: "nil face",
			face: nil,
			enc:  EncodingNone,
			want: CharMap{
				Format:     0,
				Language:   0,
				Encoding:   EncodingNone,
				PlatformID: 0,
				EncodingID: 0,
				index:      0,
				valid:      false,
			},
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name: "go regular, unicode",
			face: goRegular,
			enc:  EncodingUnicode,
			want: CharMap{
				Format:     4,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: truetype.MicrosoftEncodingUnicodeCs,
				index:      2,
				valid:      true,
			},
			wantErr: nil,
		},
		{
			name: "go regular, apple roman",
			face: goRegular,
			enc:  EncodingAppleRoman,
			want: CharMap{
				Format:     6,
				Language:   truetype.MacLangEnglish,
				Encoding:   EncodingAppleRoman,
				PlatformID: truetype.PlatformMacintosh,
				EncodingID: truetype.MacEncodingRoman,
				index:      1,
				valid:      true,
			},
			wantErr: nil,
		},
		{
			name:    "go regular, adobe latin1",
			face:    goRegular,
			enc:     EncodingAdobeLatin1,
			want:    CharMap{},
			wantErr: ErrInvalidArgument,
		},
		{
			name: "bungee layers regular, adobe standard",
			face: goRegular,
			enc:  EncodingAdobeStandard,
			want: CharMap{
				Format:     0,
				Language:   0,
				Encoding:   EncodingNone,
				PlatformID: 0,
				EncodingID: 0,
				index:      0,
				valid:      false,
			},
			wantErr: ErrInvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.face.testClearCharmap()
			if err := tt.face.SelectCharMap(tt.enc); err != tt.wantErr {
				t.Errorf("%q.SelectCharMap(%s) error = %v, wantErr %v", tt.face.FamilyName(), tt.enc, err, tt.wantErr)
			}

			if got, _ := tt.face.ActiveCharMap(); got != tt.want {
				t.Errorf("%q.SelectCharMap(%s) got charmap = %v, want %v", tt.face.FamilyName(), tt.enc, got, tt.want)
			}
		})
	}
}

func TestFace_SetCharMap(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %s", err)
	}
	defer l.Free()

	goRegular, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goRegular.Free()

	bungeeLayersReg, err := l.NewFaceFromPath(testdata("bungee", "BungeeLayers-Regular.otf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer bungeeLayersReg.Free()

	goRegMaps := goRegular.CharMaps()
	bungeeLayersRegMaps := bungeeLayersReg.CharMaps()

	tests := []struct {
		name    string
		face    *Face
		cmap    CharMap
		want    CharMap
		wantErr error
	}{
		{
			name:    "nil face",
			face:    nil,
			cmap:    CharMap{},
			want:    CharMap{},
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name:    "invalid charmap",
			face:    goRegular,
			cmap:    CharMap{},
			want:    CharMap{},
			wantErr: ErrInvalidCharMapHandle,
		},
		{
			name: "out of bounds charmap",
			face: goRegular,
			cmap: CharMap{
				valid: true,
				index: 999,
			},
			want:    CharMap{},
			wantErr: ErrInvalidCharMapHandle,
		},
		{
			name:    "go regular, cmap 0",
			face:    goRegular,
			cmap:    goRegMaps[0],
			want:    goRegMaps[0],
			wantErr: nil,
		},
		{
			name:    "go regular, cmap 1",
			face:    goRegular,
			cmap:    goRegMaps[1],
			want:    goRegMaps[1],
			wantErr: nil,
		},
		{
			name:    "go regular, cmap 2",
			face:    goRegular,
			cmap:    goRegMaps[2],
			want:    goRegMaps[2],
			wantErr: nil,
		},
		{
			name:    "bungee layers regular, cmap 0",
			face:    bungeeLayersReg,
			cmap:    bungeeLayersRegMaps[0],
			want:    bungeeLayersRegMaps[0],
			wantErr: nil,
		},
		{
			name:    "bungee layers regular, cmap 1",
			face:    bungeeLayersReg,
			cmap:    bungeeLayersRegMaps[1],
			want:    bungeeLayersRegMaps[1],
			wantErr: nil,
		},
		{
			name:    "bungee layers regular, cmap 2",
			face:    bungeeLayersReg,
			cmap:    bungeeLayersRegMaps[2],
			want:    bungeeLayersRegMaps[2],
			wantErr: nil,
		},
		{
			name:    "bungee layers regular, cmap 3",
			face:    bungeeLayersReg,
			cmap:    bungeeLayersRegMaps[3],
			want:    bungeeLayersRegMaps[3],
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.face.testClearCharmap()
			if err := tt.face.SetCharMap(tt.cmap); err != tt.wantErr {
				t.Errorf("%q.SetCharMap(%v) error = %v, wantErr %v", tt.face.FamilyName(), tt.cmap, err, tt.wantErr)
			}

			if got, _ := tt.face.ActiveCharMap(); got != tt.want {
				t.Errorf("%q.SetCharMap(%v) got charmap = %v, want %v", tt.face.FamilyName(), tt.cmap, got, tt.want)
			}
		})
	}
}

func TestFace_SetCharSize(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %s", err)
	}
	defer l.Free()

	goRegular, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goRegular.Free()

	bungeeColorMac, err := l.NewFaceFromPath(testdata("bungee", "BungeeColor-Regular_sbix_MacOS.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer bungeeColorMac.Free()

	type args struct {
		nominalWidth  fixed.Int26_6
		nominalHeight fixed.Int26_6
		horzDPI       uint
		vertDPI       uint
	}
	tests := []struct {
		name     string
		font     *Face
		args     args
		wantSize Size
		wantErr  error
	}{
		{
			name:     "nil face",
			font:     nil,
			args:     args{},
			wantSize: Size{},
			wantErr:  ErrInvalidFaceHandle,
		},
		{
			name: "go regular",
			font: goRegular,
			args: args{
				nominalWidth:  20 << 6,
				nominalHeight: 20 << 6,
				horzDPI:       72,
				vertDPI:       72,
			},
			wantSize: Size{
				SizeMetrics{
					XPpem:      20,
					YPpem:      20,
					XScale:     40960,
					YScale:     40960,
					Ascender:   1216,
					Descender:  -320,
					Height:     1472,
					MaxAdvance: 1408,
				},
			},
			wantErr: nil,
		},
		{
			name: "bungee color mac, first size",
			font: bungeeColorMac,
			args: args{
				nominalWidth:  20 << 6,
				nominalHeight: 20 << 6,
				horzDPI:       72,
				vertDPI:       72,
			},
			wantSize: Size{
				SizeMetrics{
					XPpem:      20,
					YPpem:      20,
					XScale:     83886,
					YScale:     83886,
					Ascender:   1101,
					Descender:  -179,
					Height:     1536,
					MaxAdvance: 1814,
				},
			},
			wantErr: nil,
		},
		{
			name: "bungee color mac, second size",
			font: bungeeColorMac,
			args: args{
				nominalWidth:  32 << 6,
				nominalHeight: 32 << 6,
				horzDPI:       72,
				vertDPI:       72,
			},
			wantSize: Size{
				SizeMetrics{
					XPpem:      32,
					YPpem:      32,
					XScale:     134218,
					YScale:     134218,
					Ascender:   1761,
					Descender:  -287,
					Height:     2458,
					MaxAdvance: 2902,
				},
			},
			wantErr: nil,
		},
		{
			name: "bungee color mac, < first size",
			font: bungeeColorMac,
			args: args{
				nominalWidth:  19 << 6,
				nominalHeight: 19 << 6,
				horzDPI:       72,
				vertDPI:       72,
			},
			wantSize: Size{
				SizeMetrics{
					XScale: 1 << 16,
					YScale: 1 << 16,
				},
			},
			wantErr: ErrInvalidPixelSize,
		},
		{
			name: "bungee color mac, > first size",
			font: bungeeColorMac,
			args: args{
				nominalWidth:  21 << 6,
				nominalHeight: 21 << 6,
				horzDPI:       72,
				vertDPI:       72,
			},
			wantSize: Size{
				SizeMetrics{
					XScale: 1 << 16,
					YScale: 1 << 16,
				},
			},
			wantErr: ErrInvalidPixelSize,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.font.SetCharSize(tt.args.nominalWidth, tt.args.nominalHeight, tt.args.horzDPI, tt.args.vertDPI); err != tt.wantErr {
				t.Errorf("Face.SetCharSize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := tt.font.Size(); got != tt.wantSize {
				t.Errorf("Face.SetCharSize() %v, want %v", got, tt.wantSize)
			}
		})
	}
}
