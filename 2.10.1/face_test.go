package freetype2

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"unsafe"

	"github.com/flga/freetype2/2.10.1/fixed"
	"github.com/flga/freetype2/2.10.1/truetype"
	"github.com/go-test/deep"
)

func TestFaceFlag_String(t *testing.T) {
	var x FaceFlag
	if got, want := x.String(), ""; got != want {
		t.Errorf("FaceFlag.String(0) = %v, want %v", got, want)
	}

	x = FaceFlagColor
	if got, want := x.String(), "Color"; got != want {
		t.Errorf("FaceFlag.String(FaceFlagColor) = %v, want %v", got, want)
	}

	x = FaceFlagKerning | FaceFlagCidKeyed
	if got, want := x.String(), "Kerning|CidKeyed"; got != want {
		t.Errorf("FaceFlag.String(FaceFlagKerning | FaceFlagCidKeyed) = %v, want %v", got, want)
	}

	x = FaceFlagVertical | FaceFlagMultipleMasters | FaceFlagHinter
	if got, want := x.String(), "Vertical|MultipleMasters|Hinter"; got != want {
		t.Errorf("FaceFlag.String(FaceFlagVertical | FaceFlagMultipleMasters | FaceFlagHinter) = %v, want %v", got, want)
	}

	x = FaceFlagScalable | FaceFlagFixedSizes | FaceFlagFixedWidth | FaceFlagSfnt | FaceFlagHorizontal |
		FaceFlagVertical | FaceFlagKerning | FaceFlagMultipleMasters | FaceFlagGlyphNames | FaceFlagHinter |
		FaceFlagCidKeyed | FaceFlagTricky | FaceFlagColor | FaceFlagVariation
	if got, want := x.String(), "Scalable|FixedSizes|FixedWidth|Sfnt|Horizontal|Vertical|Kerning|MultipleMasters|GlyphNames|Hinter|CidKeyed|Tricky|Color|Variation"; got != want {
		t.Errorf("FaceFlag.String(FaceFlagScalable | FaceFlagFixedSizes | FaceFlagFixedWidth | FaceFlagSfnt | FaceFlagHorizontal | FaceFlagVertical | FaceFlagKerning | FaceFlagMultipleMasters | FaceFlagGlyphNames | FaceFlagHinter | FaceFlagCidKeyed | FaceFlagTricky | FaceFlagColor | FaceFlagVariation) = %v, want %v", got, want)
	}
}

func TestStyleFlag_String(t *testing.T) {
	var x StyleFlag

	if got, want := x.String(), ""; got != want {
		t.Errorf("StyleFlag.String(0) = %v, want %v", got, want)
	}

	x = StyleFlagItalic
	if got, want := x.String(), "Italic"; got != want {
		t.Errorf("StyleFlag.String(StyleFlagItalic) = %v, want %v", got, want)
	}

	x = StyleFlagBold
	if got, want := x.String(), "Bold"; got != want {
		t.Errorf("StyleFlag.String(StyleFlagBold) = %v, want %v", got, want)
	}

	x = StyleFlagItalic | StyleFlagBold
	if got, want := x.String(), "Italic|Bold"; got != want {
		t.Errorf("StyleFlag.String(StyleFlagItalic | StyleFlagBold) = %v, want %v", got, want)
	}
}

func TestLoadFlag_String(t *testing.T) {
	var x LoadFlag
	if got, want := x.String(), "Default"; got != want {
		t.Errorf("LoadFlag.String(0) = %v, want %v", got, want)
	}

	x = 1 << 30
	if got, want := x.String(), ""; got != want {
		t.Errorf("LoadFlag.String(1 << 30) = %v, want %v", got, want)
	}

	x = LoadDefault
	if got, want := x.String(), "Default"; got != want {
		t.Errorf("LoadFlag.String(LoadDefault) = %v, want %v", got, want)
	}

	x = LoadColor
	if got, want := x.String(), "Color"; got != want {
		t.Errorf("LoadFlag.String(LoadColor) = %v, want %v", got, want)
	}

	x = LoadMonochrome | LoadNoAutohint
	if got, want := x.String(), "Monochrome|NoAutohint"; got != want {
		t.Errorf("LoadFlag.String(LoadMonochrome | LoadNoAutohint) = %v, want %v", got, want)
	}

	x = LoadIgnoreTransform | LoadColor | LoadComputeMetrics
	if got, want := x.String(), "IgnoreTransform|Color|ComputeMetrics"; got != want {
		t.Errorf("LoadFlag.String(LoadIgnoreTransform | LoadColor | LoadComputeMetrics) = %v, want %v", got, want)
	}

	x = LoadDefault | LoadNoScale | LoadNoHinting | LoadRender | LoadNoBitmap | LoadVerticalLayout | LoadForceAutohint |
		LoadPedantic | LoadNoRecurse | LoadIgnoreTransform | LoadMonochrome | LoadLinearDesign | LoadNoAutohint |
		LoadColor | LoadComputeMetrics | LoadBitmapMetricsOnly
	if got, want := x.String(), "NoScale|NoHinting|Render|NoBitmap|VerticalLayout|ForceAutohint|Pedantic|NoRecurse|IgnoreTransform|Monochrome|LinearDesign|NoAutohint|Color|ComputeMetrics|BitmapMetricsOnly"; got != want {
		t.Errorf("LoadFlag.String(LoadDefault | LoadNoScale | LoadNoHinting | LoadRender | LoadNoBitmap | LoadVerticalLayout | LoadForceAutohint | LoadPedantic | LoadNoRecurse | LoadIgnoreTransform | LoadMonochrome | LoadLinearDesign | LoadNoAutohint | LoadColor | LoadComputeMetrics | LoadBitmapMetricsOnly) = %v, want %v", got, want)
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
	tests := []struct {
		name               string
		face               func() (testface, error)
		family             string
		style              string
		postscript         string
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
		charmaps           []CharMap
		activeCharmap      CharMap
		activeOk           bool
		numSizes           int
		avaliableSizes     []BitmapSize
		maxAdvanceWidth    int
		maxAdvanceHeight   int
		underlinePosition  int
		underlineThickness int
	}{
		{
			name:               "nil face",
			face:               nilFace,
			family:             "",
			style:              "",
			postscript:         "",
			numFaces:           0,
			numNamedInstances:  0,
			faceIdx:            0,
			namedIdx:           0,
			bold:               false,
			italic:             false,
			sfntWrapped:        false,
			scalable:           false,
			fixedSize:          false,
			horizontal:         false,
			vertical:           false,
			fixedWidth:         false,
			glyphNames:         false,
			emSize:             0,
			globalBBox:         BBox{},
			ascent:             0,
			descent:            0,
			textHeight:         0,
			glyphCount:         0,
			numCharmaps:        0,
			charmaps:           nil,
			activeCharmap:      CharMap{},
			activeOk:           false,
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    0,
			maxAdvanceHeight:   0,
			underlinePosition:  0,
			underlineThickness: 0,
		},
		{
			name:              "goRegular",
			face:              goRegular,
			family:            "Go",
			style:             "Regular",
			postscript:        "GoRegular",
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
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.TrimmedTableMapping, Language: 0, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: 0, index: 1, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 2, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentMappingToDeltaValues,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 1,
				index:      2,
				valid:      true,
			},
			activeOk:           true,
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
			postscript:        "Go-Bold",
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
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.TrimmedTableMapping, Language: 0, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: 0, index: 1, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 2, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentMappingToDeltaValues,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 1,
				index:      2,
				valid:      true,
			},
			activeOk:           true,
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
			postscript:        "Go-Italic",
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
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.TrimmedTableMapping, Language: 0, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: 0, index: 1, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 2, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentMappingToDeltaValues,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 1,
				index:      2,
				valid:      true,
			},
			activeOk:           true,
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
			postscript:        "Go-BoldItalic",
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
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.TrimmedTableMapping, Language: 0, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: 0, index: 1, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 2, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentMappingToDeltaValues,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 1,
				index:      2,
				valid:      true,
			},
			activeOk:           true,
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
			postscript:        "GoMono",
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
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.TrimmedTableMapping, Language: 0, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: 0, index: 1, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 2, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentMappingToDeltaValues,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 1,
				index:      2,
				valid:      true,
			},
			activeOk:           true,
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
			postscript:        "BungeeColor-Regular",
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
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.TrimmedTableMapping, Language: 0, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: 0, index: 1, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 2, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentMappingToDeltaValues,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 1,
				index:      2,
				valid:      true,
			},
			activeOk:           true,
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
			postscript:        "BungeeColor-Regular",
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
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.TrimmedTableMapping, Language: 0, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: 0, index: 1, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 2, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentMappingToDeltaValues,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 1,
				index:      2,
				valid:      true,
			},
			activeOk: true,
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
			postscript:        "BungeeLayers-Regular",
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
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.TrimmedTableMapping, Language: 0, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: 0, index: 1, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 2, valid: true},
				{Format: -1, Language: 0, Encoding: EncodingAdobeStandard, PlatformID: truetype.PlatformAdobe, EncodingID: 0, index: 3, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentMappingToDeltaValues,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 1,
				index:      2,
				valid:      true,
			},
			activeOk:           true,
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    1417,
			maxAdvanceHeight:   1200,
			underlinePosition:  0,
			underlineThickness: 0,
		},
		{
			name:              "notoSansJpReg",
			face:              notoSansJpReg,
			family:            "Noto Sans JP",
			style:             "Regular",
			postscript:        "NotoSansJP-Regular",
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
			vertical:          true,
			fixedWidth:        false,
			glyphNames:        false,
			emSize:            1000,
			globalBBox:        BBox{XMin: -1002, YMin: -1048, XMax: 2928, YMax: 1808},
			ascent:            1160,
			descent:           -320,
			textHeight:        1480,
			glyphCount:        17802,
			numCharmaps:       5,
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.SegmentedCoverage, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 4, index: 1, valid: true},
				{Format: truetype.UnicodeVariationSequences, Language: truetype.UVSLang, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 5, index: 2, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 3, valid: true},
				{Format: truetype.SegmentedCoverage, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 10, index: 4, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentedCoverage,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 10,
				index:      4,
				valid:      true,
			},
			numSizes:           0,
			activeOk:           true,
			avaliableSizes:     nil,
			maxAdvanceWidth:    3000,
			maxAdvanceHeight:   3000,
			underlinePosition:  -150,
			underlineThickness: 50,
		},
		{
			name:              "notoSansJpBold",
			face:              notoSansJpBold,
			family:            "Noto Sans JP",
			style:             "Bold",
			postscript:        "NotoSansJP-Bold",
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
			vertical:          true,
			fixedWidth:        false,
			glyphNames:        false,
			emSize:            1000,
			globalBBox:        BBox{XMin: -1013, YMin: -1046, XMax: 2926, YMax: 1806},
			ascent:            1160,
			descent:           -320,
			textHeight:        1480,
			glyphCount:        17802,
			numCharmaps:       5,
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 3, index: 0, valid: true},
				{Format: truetype.SegmentedCoverage, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 4, index: 1, valid: true},
				{Format: truetype.UnicodeVariationSequences, Language: truetype.UVSLang, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: 5, index: 2, valid: true},
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 3, valid: true},
				{Format: truetype.SegmentedCoverage, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 10, index: 4, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentedCoverage,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 10,
				index:      4,
				valid:      true,
			},
			activeOk:           true,
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    3000,
			maxAdvanceHeight:   3000,
			underlinePosition:  -150,
			underlineThickness: 50,
		},
		{
			name:              "arimoRegular",
			face:              arimoRegular,
			family:            "Arimo",
			style:             "Regular",
			postscript:        "Arimo",
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
			globalBBox:        BBox{XMin: -1114, YMin: -621, XMax: 2666, YMax: 2007},
			ascent:            1854,
			descent:           -434,
			textHeight:        2355,
			glyphCount:        2584,
			numCharmaps:       1,
			charmaps: []CharMap{
				{Format: truetype.SegmentMappingToDeltaValues, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: 1, index: 0, valid: true},
			},
			activeCharmap: CharMap{
				Format:     truetype.SegmentMappingToDeltaValues,
				Language:   0,
				Encoding:   EncodingUnicode,
				PlatformID: truetype.PlatformMicrosoft,
				EncodingID: 1,
				index:      0,
				valid:      true,
			},
			activeOk:           true,
			numSizes:           0,
			avaliableSizes:     nil,
			maxAdvanceWidth:    2740,
			maxAdvanceHeight:   2355,
			underlinePosition:  -292,
			underlineThickness: 150,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			testFlag := func(v bool, f FaceFlag) {
				if v && !face.HasFlag(f) {
					t.Errorf("Flags() face should have flag %s", f)
				}
				if !v && face.HasFlag(f) {
					t.Errorf("Flags() face should not have flag %s", f)
				}
			}

			testStyleFlag := func(v bool, f StyleFlag) {
				if v && !face.HasStyle(f) {
					t.Errorf("StyleFlags() face should have flag %s", f)
				}
				if !v && face.HasStyle(f) {
					t.Errorf("StyleFlags() face should not have flag %s", f)
				}
			}

			testProp := func(name string, got, want interface{}) {
				switch reflect.TypeOf(got).Kind() {
				case reflect.Slice,
					reflect.Struct,
					reflect.Array,
					reflect.Interface,
					reflect.Map:
					if diff := deep.Equal(got, want); diff != nil {
						t.Errorf("%s() = %v", name, diff)
					}
				default:
					if !reflect.DeepEqual(got, want) {
						t.Errorf("%s() = %v, want %v", name, got, want)
					}
				}
			}

			testProp("FamilyName", face.FamilyName(), tt.family)
			testProp("StyleName", face.StyleName(), tt.style)
			testProp("PostscriptName", face.PostscriptName(), tt.postscript)
			testProp("NumFaces", face.NumFaces(), tt.numFaces)
			testProp("NumNamedInstances", face.NumNamedInstances(), tt.numNamedInstances)
			testProp("Index", face.Index(), tt.faceIdx)
			testProp("NamedIndex", face.NamedIndex(), tt.namedIdx)
			testStyleFlag(tt.bold, StyleFlagBold)
			testStyleFlag(tt.italic, StyleFlagItalic)
			testFlag(tt.sfntWrapped, FaceFlagSfnt)
			testFlag(tt.scalable, FaceFlagScalable)
			testFlag(tt.fixedSize, FaceFlagFixedSizes)
			testFlag(tt.horizontal, FaceFlagHorizontal)
			testFlag(tt.vertical, FaceFlagVertical)
			testFlag(tt.fixedWidth, FaceFlagFixedWidth)
			testFlag(tt.glyphNames, FaceFlagGlyphNames)
			testProp("UnitsPerEM", face.UnitsPerEM(), tt.emSize)
			testProp("BBox", face.BBox(), tt.globalBBox)
			testProp("Ascender", face.Ascender(), tt.ascent)
			testProp("Descender", face.Descender(), tt.descent)
			testProp("Height", face.Height(), tt.textHeight)
			testProp("NumGlyphs", face.NumGlyphs(), tt.glyphCount)
			testProp("NumCharMaps", face.NumCharMaps(), tt.numCharmaps)
			testProp("CharMaps", face.CharMaps(), tt.charmaps)
			active, activeOk := face.ActiveCharMap()
			testProp("ActiveCharMap", active, tt.activeCharmap)
			testProp("ActiveCharMap", activeOk, tt.activeOk)
			testProp("NumFixedSizes", face.NumFixedSizes(), tt.numSizes)
			testProp("AvailableSizes", face.AvailableSizes(), tt.avaliableSizes)
			testProp("Size", face.Size(), Size{})
			testProp("MaxAdvanceWidth", face.MaxAdvanceWidth(), tt.maxAdvanceWidth)
			testProp("MaxAdvanceHeight", face.MaxAdvanceHeight(), tt.maxAdvanceHeight)
			testProp("UnderlinePosition", face.UnderlinePosition(), tt.underlinePosition)
			testProp("UnderlineThickness", face.UnderlineThickness(), tt.underlineThickness)
			testProp("Glyph", face.Glyph(), GlyphSlot{})
		})
	}
}

func TestFace_SelectCharMap(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		enc     Encoding
		want    CharMap
		wantErr error
	}{
		{
			name: "nil face",
			face: nilFace,
			enc:  EncodingNone,
			want: CharMap{
				Format:     truetype.ByteEncodingTable,
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
				Format:     truetype.SegmentMappingToDeltaValues,
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
				Format:     truetype.TrimmedTableMapping,
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
				Format:     truetype.ByteEncodingTable,
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
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if err := face.SelectCharMap(tt.enc); err != tt.wantErr {
				t.Errorf("Face.SelectCharMap(%s) error = %v, wantErr %v", tt.enc, err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}

			if got, _ := face.ActiveCharMap(); got != tt.want {
				t.Errorf("Face.SelectCharMap(%s) got charmap = %v, want %v", tt.enc, got, tt.want)
			}
		})
	}
}

func TestFace_SetCharMap(t *testing.T) {
	var goRegMaps, bungeeLayersRegMaps []CharMap
	{
		goRegularFace, err := goRegular()
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		goRegMaps = goRegularFace.CharMaps()
		goRegularFace.Free()

		bungeeLayersRegFace, err := bungeeLayersReg()
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		bungeeLayersRegMaps = bungeeLayersRegFace.CharMaps()
		bungeeLayersRegFace.Free()
	}

	tests := []struct {
		name    string
		face    func() (testface, error)
		cmap    CharMap
		want    CharMap
		wantErr error
	}{
		{
			name:    "nil face",
			face:    nilFace,
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
			name: "negative charmap",
			face: goRegular,
			cmap: CharMap{
				valid: true,
				index: -1,
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
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if err := face.SetCharMap(tt.cmap); err != tt.wantErr {
				t.Errorf("%q.SetCharMap(%v) error = %v, wantErr %v", face.FamilyName(), tt.cmap, err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}

			if got, _ := face.ActiveCharMap(); got != tt.want {
				t.Errorf("%q.SetCharMap(%v) got charmap = %v, want %v", face.FamilyName(), tt.cmap, got, tt.want)
			}
		})
	}
}

func TestFace_SetCharSize(t *testing.T) {
	type args struct {
		nominalWidth  fixed.Int26_6
		nominalHeight fixed.Int26_6
		horzDPI       uint
		vertDPI       uint
	}
	tests := []struct {
		name     string
		face     func() (testface, error)
		args     args
		wantSize Size
		wantErr  error
	}{
		{
			name:     "nil face",
			face:     nilFace,
			args:     args{},
			wantSize: Size{},
			wantErr:  ErrInvalidFaceHandle,
		},
		{
			name: "go regular",
			face: goRegular,
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
			face: bungeeColorMac,
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
			face: bungeeColorMac,
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
			face: bungeeColorMac,
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
			face: bungeeColorMac,
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
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if err := face.SetCharSize(tt.args.nominalWidth, tt.args.nominalHeight, tt.args.horzDPI, tt.args.vertDPI); err != tt.wantErr {
				t.Errorf("Face.SetCharSize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := face.Size(); got != tt.wantSize {
				t.Errorf("Face.SetCharSize() %v, want %v", got, tt.wantSize)
			}
		})
	}
}

func TestFace_SetPixelSizes(t *testing.T) {
	type args struct {
		width  uint
		height uint
	}
	tests := []struct {
		name     string
		face     func() (testface, error)
		args     args
		wantSize Size
		wantErr  error
	}{
		{
			name:     "nil face",
			face:     nilFace,
			args:     args{},
			wantSize: Size{},
			wantErr:  ErrInvalidFaceHandle,
		},
		{
			name: "go regular",
			face: goRegular,
			args: args{
				width:  20,
				height: 20,
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
			face: bungeeColorMac,
			args: args{
				width:  20,
				height: 20,
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
			face: bungeeColorMac,
			args: args{
				width:  32,
				height: 32,
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
			face: bungeeColorMac,
			args: args{
				width:  19,
				height: 19,
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
			face: bungeeColorMac,
			args: args{
				width:  21,
				height: 21,
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
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if err := face.SetPixelSizes(tt.args.width, tt.args.height); err != tt.wantErr {
				t.Errorf("Face.SetPixelSizes() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := face.Size(); got != tt.wantSize {
				t.Errorf("Face.SetPixelSizes() %v, want %v", got, tt.wantSize)
			}
		})
	}
}

func TestFace_RequestSize(t *testing.T) {
	tests := []struct {
		name     string
		face     func() (testface, error)
		req      SizeRequest
		wantSize Size
		wantErr  error
	}{
		{
			name:     "nil face",
			face:     nilFace,
			req:      SizeRequest{},
			wantSize: Size{},
			wantErr:  ErrInvalidFaceHandle,
		},
		{
			name: "go regular nominal",
			face: goRegular,
			req: SizeRequest{
				Type:           SizeRequestTypeNominal,
				Width:          20 << 6,
				Height:         20 << 6,
				HoriResolution: 72,
				VertResolution: 72,
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
			name: "go regular real dim",
			face: goRegular,
			req: SizeRequest{
				Type:           SizeRequestTypeRealDim,
				Width:          20 << 6,
				Height:         20 << 6,
				HoriResolution: 72,
				VertResolution: 72,
			},
			wantSize: Size{
				SizeMetrics{
					XPpem:      17,
					YPpem:      17,
					XScale:     35440,
					YScale:     35440,
					Ascender:   1088,
					Descender:  -256,
					Height:     1280,
					MaxAdvance: 1216,
				},
			},
			wantErr: nil,
		},
		{
			name: "go regular bbox",
			face: goRegular,
			req: SizeRequest{
				Type:           SizeRequestTypeBBox,
				Width:          20 << 6,
				Height:         20 << 6,
				HoriResolution: 72,
				VertResolution: 72,
			},
			wantSize: Size{
				SizeMetrics{
					XPpem:      16,
					YPpem:      15,
					XScale:     32264,
					YScale:     31524,
					Ascender:   960,
					Descender:  -256,
					Height:     1152,
					MaxAdvance: 1088,
				},
			},
			wantErr: nil,
		},
		{
			name: "go regular cell",
			face: goRegular,
			req: SizeRequest{
				Type:           SizeRequestTypeCell,
				Width:          20 << 6,
				Height:         20 << 6,
				HoriResolution: 72,
				VertResolution: 72,
			},
			wantSize: Size{
				SizeMetrics{
					XPpem:      17,
					YPpem:      17,
					XScale:     35440,
					YScale:     35440,
					Ascender:   1088,
					Descender:  -256,
					Height:     1280,
					MaxAdvance: 1216,
				},
			},
			wantErr: nil,
		},
		{
			name: "go regular scales",
			face: goRegular,
			req: SizeRequest{
				Type:           SizeRequestTypeScales,
				Width:          20 << 6,
				Height:         20 << 6,
				HoriResolution: 72,
				VertResolution: 72,
			},
			wantSize: Size{
				SizeMetrics{
					XPpem:      1,
					YPpem:      1,
					XScale:     1280,
					YScale:     1280,
					Ascender:   64,
					Descender:  -64,
					Height:     64,
					MaxAdvance: 64,
				},
			},
			wantErr: nil,
		},
		{
			name: "go regular invalid ppem",
			face: goRegular,
			req: SizeRequest{
				Type:           SizeRequestTypeNominal,
				Width:          20 << 6,
				Height:         20 << 6,
				HoriResolution: 1,
				VertResolution: 1,
			},
			wantSize: Size{
				SizeMetrics{
					XPpem:      0,
					YPpem:      0,
					XScale:     576,
					YScale:     576,
					Ascender:   64,
					Descender:  -64,
					Height:     0,
					MaxAdvance: 0,
				},
			},
			wantErr: ErrInvalidPPem,
		},
		{
			name: "bungee color mac, first size, nominal",
			face: bungeeColorMac,
			req: SizeRequest{
				Type:           SizeRequestTypeNominal,
				Width:          20 << 6,
				Height:         20 << 6,
				HoriResolution: 72,
				VertResolution: 72,
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
			name: "bungee color mac, second size, nominal",
			face: bungeeColorMac,
			req: SizeRequest{
				Type:           SizeRequestTypeNominal,
				Width:          32 << 6,
				Height:         32 << 6,
				HoriResolution: 72,
				VertResolution: 72,
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
			name: "bungee color mac, real dim",
			face: bungeeColorMac,
			req: SizeRequest{
				Type:           SizeRequestTypeRealDim,
				Width:          32 << 6,
				Height:         32 << 6,
				HoriResolution: 72,
				VertResolution: 72,
			},
			wantSize: Size{
				SizeMetrics{
					XScale: 1 << 16,
					YScale: 1 << 16,
				},
			},
			wantErr: ErrUnimplementedFeature,
		},
		{
			name: "bungee color mac, bbox",
			face: bungeeColorMac,
			req: SizeRequest{
				Type:           SizeRequestTypeBBox,
				Width:          32 << 6,
				Height:         32 << 6,
				HoriResolution: 72,
				VertResolution: 72,
			},
			wantSize: Size{
				SizeMetrics{
					XScale: 1 << 16,
					YScale: 1 << 16,
				},
			},
			wantErr: ErrUnimplementedFeature,
		},
		{
			name: "bungee color mac, cell",
			face: bungeeColorMac,
			req: SizeRequest{
				Type:           SizeRequestTypeCell,
				Width:          32 << 6,
				Height:         32 << 6,
				HoriResolution: 72,
				VertResolution: 72,
			},
			wantSize: Size{
				SizeMetrics{
					XScale: 1 << 16,
					YScale: 1 << 16,
				},
			},
			wantErr: ErrUnimplementedFeature,
		},
		{
			name: "bungee color mac, scales",
			face: bungeeColorMac,
			req: SizeRequest{
				Type:           SizeRequestTypeScales,
				Width:          32 << 6,
				Height:         32 << 6,
				HoriResolution: 72,
				VertResolution: 72,
			},
			wantSize: Size{
				SizeMetrics{
					XScale: 1 << 16,
					YScale: 1 << 16,
				},
			},
			wantErr: ErrUnimplementedFeature,
		},
		{
			name: "bungee color mac, < first size",
			face: bungeeColorMac,
			req: SizeRequest{
				Type:           SizeRequestTypeNominal,
				Width:          19 << 6,
				Height:         19 << 6,
				HoriResolution: 72,
				VertResolution: 72,
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
			face: bungeeColorMac,
			req: SizeRequest{
				Type:           SizeRequestTypeNominal,
				Width:          21 << 6,
				Height:         21 << 6,
				HoriResolution: 72,
				VertResolution: 72,
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
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if err := face.RequestSize(tt.req); err != tt.wantErr {
				t.Errorf("Face.RequestSize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := face.Size(); got != tt.wantSize {
				t.Errorf("Face.RequestSize() %v, want %v", got, tt.wantSize)
			}
		})
	}
}

func TestFace_RequestSize_Free(t *testing.T) {
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

	var freed bool
	defer mockFree(func(_ unsafe.Pointer) {
		freed = true
	}, actuallyFreeItAfter)()
	if err := goRegular.RequestSize(SizeRequest{
		Type:           SizeRequestTypeNominal,
		Width:          20 << 6,
		Height:         20 << 6,
		HoriResolution: 72,
		VertResolution: 72,
	}); err != nil {
		t.Fatalf("unable to request size: %s", err)
	}

	if !freed {
		t.Errorf("free() was not called")
	}
}

func TestFace_SelectSize(t *testing.T) {
	tests := []struct {
		name     string
		face     func() (testface, error)
		idx      int
		wantSize Size
		wantErr  error
	}{
		{
			name:     "nil face",
			face:     nilFace,
			idx:      0,
			wantSize: Size{},
			wantErr:  ErrInvalidFaceHandle,
		},
		{
			name:     "go regular 0",
			face:     goRegular,
			idx:      0,
			wantSize: Size{},
			wantErr:  ErrInvalidFaceHandle,
		},
		{
			name:     "go regular 1",
			face:     goRegular,
			idx:      1,
			wantSize: Size{},
			wantErr:  ErrInvalidFaceHandle,
		},
		{
			name: "bungee color mac 0",
			face: bungeeColorMac,
			idx:  0,
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
			name: "bungee color mac 1",
			face: bungeeColorMac,
			idx:  1,
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
			name: "bungee color mac 2",
			face: bungeeColorMac,
			idx:  2,
			wantSize: Size{
				SizeMetrics{
					XPpem:      40,
					YPpem:      40,
					XScale:     167772,
					YScale:     167772,
					Ascender:   2202,
					Descender:  -358,
					Height:     3072,
					MaxAdvance: 3628,
				},
			},
			wantErr: nil,
		},
		{
			name: "bungee color mac 3",
			face: bungeeColorMac,
			idx:  3,
			wantSize: Size{
				SizeMetrics{
					XPpem:      0x48,
					YPpem:      0x48,
					XScale:     301990,
					YScale:     301990,
					Ascender:   3963,
					Descender:  -645,
					Height:     5530,
					MaxAdvance: 6530,
				},
			},
			wantErr: nil,
		},
		{
			name: "bungee color mac 4",
			face: bungeeColorMac,
			idx:  4,
			wantSize: Size{
				SizeMetrics{
					XPpem:      0x60,
					YPpem:      0x60,
					XScale:     402653,
					YScale:     402653,
					Ascender:   5284,
					Descender:  -860,
					Height:     7373,
					MaxAdvance: 8706,
				},
			},
			wantErr: nil,
		},
		{
			name: "bungee color mac 5",
			face: bungeeColorMac,
			idx:  5,
			wantSize: Size{
				SizeMetrics{
					XPpem:      0x80,
					YPpem:      0x80,
					XScale:     536871,
					YScale:     536871,
					Ascender:   7045,
					Descender:  -1147,
					Height:     9830,
					MaxAdvance: 11608,
				},
			},
			wantErr: nil,
		},
		{
			name: "bungee color mac 6",
			face: bungeeColorMac,
			idx:  6,
			wantSize: Size{
				SizeMetrics{
					XPpem:      0x100,
					YPpem:      0x100,
					XScale:     1073742,
					YScale:     1073742,
					Ascender:   14090,
					Descender:  -2294,
					Height:     19661,
					MaxAdvance: 23216,
				},
			},
			wantErr: nil,
		},
		{
			name: "bungee color mac 7",
			face: bungeeColorMac,
			idx:  7,
			wantSize: Size{
				SizeMetrics{
					XPpem:      0x200,
					YPpem:      0x200,
					XScale:     2147484,
					YScale:     2147484,
					Ascender:   28180,
					Descender:  -4588,
					Height:     39322,
					MaxAdvance: 46432,
				},
			},
			wantErr: nil,
		},
		{
			name: "bungee color mac 8",
			face: bungeeColorMac,
			idx:  8,
			wantSize: Size{
				SizeMetrics{
					XPpem:      0x400,
					YPpem:      0x400,
					XScale:     4294967,
					YScale:     4294967,
					Ascender:   56361,
					Descender:  -9175,
					Height:     78643,
					MaxAdvance: 92865,
				},
			},
			wantErr: nil,
		},
		{
			name:     "bungee color mac 9",
			face:     bungeeColorMac,
			idx:      9,
			wantSize: Size{},
			wantErr:  ErrInvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if err := face.SelectSize(tt.idx); err != tt.wantErr {
				t.Errorf("Face.SelectSize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := face.Size(); got != tt.wantSize {
				t.Errorf("Face.SelectSize() %v, want %v", got, tt.wantSize)
			}
		})
	}
}

func TestFace_Glyph(t *testing.T) {
	tests := []struct {
		name  string
		setup func(f *Face) error
		face  func() (testface, error)
		want  GlyphSlot
	}{
		{
			name: "nill face",
			face: nilFace,
			want: GlyphSlot{},
		},
		{
			name: "go regular",
			face: goRegular,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return f.LoadGlyph(0x24, LoadRender|LoadColor)
			},
			want: GlyphSlot{
				GlyphIndex: 0x24,
				Metrics: GlyphMetrics{
					Width:        640,
					Height:       704,
					HoriBearingX: 0,
					HoriBearingY: 704,
					HoriAdvance:  576,
				},
				LinearHoriAdvance: 611968,
				LinearVertAdvance: 884352,
				Advance: Vector26_6{
					X: 576,
					Y: 0,
				},
				Format: GlyphFormatBitmap,
				Bitmap: Bitmap{
					Rows:      0xb,
					Width:     0xa,
					Pitch:     10,
					Buffer:    goRegularGlyphBuf(0x24, 0),
					NumGrays:  0x100,
					PixelMode: PixelModeGray,
				},
				BitmapLeft: 0,
				BitmapTop:  11,
				Outline: Outline{
					Points: []Vector{
						{X: 0x00000008, Y: 0x00000000}, {X: 0x000000fe, Y: 0x000002c0}, {X: 0x00000159, Y: 0x000002c0},
						{X: 0x0000024b, Y: 0x00000000}, {X: 0x000001e8, Y: 0x00000000}, {X: 0x000001a5, Y: 0x000000b6},
						{X: 0x000000a1, Y: 0x000000b6}, {X: 0x0000005e, Y: 0x00000000}, {X: 0x000000bb, Y: 0x000000ff},
						{X: 0x0000018c, Y: 0x000000ff}, {X: 0x00000124, Y: 0x00000235},
					},
					Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
					Contours: []int16{0x0007, 0x000a},
					Flags:    0x00000130,
				},
				NumSubglyphs: 0,
				LsbDelta:     0,
				RsbDelta:     0,
			},
		},
		{
			name: "bungee color mac",
			face: bungeeColorMac,
			setup: func(f *Face) error {
				if err := f.SelectSize(0); err != nil {
					return fmt.Errorf("unable to select size: %v", err)
				}

				return f.LoadGlyph(0x2b, LoadRender|LoadColor)
			},
			want: GlyphSlot{
				GlyphIndex: 0x2b,
				Metrics: GlyphMetrics{
					Width:        832,
					Height:       960,
					HoriBearingX: 0,
					HoriBearingY: 960,
					HoriAdvance:  896,
				},
				LinearHoriAdvance: 0,
				LinearVertAdvance: 0,
				Advance: Vector26_6{
					X: 896,
					Y: 0,
				},
				Format: GlyphFormatBitmap,
				Bitmap: Bitmap{
					Rows:      0xf,
					Width:     0xd,
					Pitch:     52,
					Buffer:    bungeeColorMacGlyphBuf(0x2b),
					NumGrays:  0x100,
					PixelMode: PixelModeBGRA,
				},
				BitmapLeft:   0,
				BitmapTop:    15,
				Outline:      Outline{},
				NumSubglyphs: 0,
				LsbDelta:     0,
				RsbDelta:     0,
			},
		},
		{
			name: "noto sans jp reg",
			face: notoSansJpReg,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return f.LoadGlyph(0x22, LoadRender|LoadColor)
			},
			want: GlyphSlot{
				GlyphIndex: 0x22,
				Metrics: GlyphMetrics{
					Width:        576,
					Height:       704,
					HoriBearingX: 0,
					HoriBearingY: 704,
					HoriAdvance:  576,
					VertBearingX: -320,
					VertBearingY: 128,
					VertAdvance:  896,
				},
				LinearHoriAdvance: 556923,
				LinearVertAdvance: 917500,
				Advance: Vector26_6{
					X: 576,
					Y: 0,
				},
				Format: GlyphFormatBitmap,
				Bitmap: Bitmap{
					Rows:      0xb,
					Width:     0x9,
					Pitch:     9,
					Buffer:    notoSansJpRegGlyphBuf(0x22),
					NumGrays:  0x100,
					PixelMode: PixelModeGray,
				},
				BitmapLeft: 0,
				BitmapTop:  11,
				Outline: Outline{
					Points: []Vector{
						{X: 0x000000ab, Y: 0x00000102}, {X: 0x000000cb, Y: 0x00000178}, {X: 0x000000e2, Y: 0x000001cf},
						{X: 0x000000f8, Y: 0x00000222}, {X: 0x0000010c, Y: 0x0000027c}, {X: 0x00000110, Y: 0x0000027c},
						{X: 0x00000126, Y: 0x00000223}, {X: 0x0000013a, Y: 0x000001cf}, {X: 0x00000152, Y: 0x00000178},
						{X: 0x00000172, Y: 0x00000102}, {X: 0x000001c5, Y: 0x00000000}, {X: 0x0000021d, Y: 0x00000000},
						{X: 0x0000013e, Y: 0x000002c0}, {X: 0x000000e1, Y: 0x000002c0}, {X: 0x00000002, Y: 0x00000000},
						{X: 0x00000056, Y: 0x00000000}, {X: 0x00000096, Y: 0x000000c0}, {X: 0x00000186, Y: 0x000000c0},
					},
					Tags:     []byte{0x01, 0x01, 0x02, 0x02, 0x01, 0x01, 0x02, 0x02, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
					Contours: []int16{0x0009, 0x0011},
					Flags:    0x00000104,
				},
				NumSubglyphs: 0,
				LsbDelta:     0,
				RsbDelta:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if tt.setup != nil {
				if err := tt.setup(face.Face); err != nil {
					t.Fatalf("Face.Glyph() setup error: %v", err)
				}
			}

			got := face.Glyph()

			// in this case, vertical metrics are unreliable
			if !face.HasFlag(FaceFlagVertical) {
				got.Metrics.VertAdvance, tt.want.Metrics.VertAdvance = -1, -1
				got.Metrics.VertBearingX, tt.want.Metrics.VertBearingX = -1, -1
				got.Metrics.VertBearingY, tt.want.Metrics.VertBearingY = -1, -1
			}

			// in this case, horizontal metrics are unreliable
			if !face.HasFlag(FaceFlagHorizontal) {
				got.Metrics.HoriAdvance, tt.want.Metrics.HoriAdvance = -1, -1
				got.Metrics.HoriBearingX, tt.want.Metrics.HoriBearingX = -1, -1
				got.Metrics.HoriBearingY, tt.want.Metrics.HoriBearingY = -1, -1
			}

			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestFace_SetTransform(t *testing.T) {
	goRegular, err := goRegular()
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goRegular.Free()

	// There's no way to test without including freetype's internal headers,
	// as such the tests are pretty superficial.

	t.Run("nil face", func(t *testing.T) {
		var f *Face
		var freeCount int
		defer mockFree(func(v unsafe.Pointer) {
			freeCount++
		}, actuallyFreeItAfter)()

		f.SetTransform(Matrix{}, Vector{})
		if want := 0; freeCount != want {
			t.Errorf("Face.SetTransform() free count %d, want %d", freeCount, want)
		}
	})

	t.Run("non 0 matrix", func(t *testing.T) {
		var freeCount int
		defer mockFree(func(v unsafe.Pointer) {
			freeCount++
		}, actuallyFreeItAfter)()

		goRegular.SetTransform(Matrix{Xx: 1}, Vector{})
		if want := 1; freeCount != want {
			t.Errorf("Face.SetTransform() free count %d, want %d", freeCount, want)
		}
	})

	t.Run("non 0 vec", func(t *testing.T) {
		var freeCount int
		defer mockFree(func(v unsafe.Pointer) {
			freeCount++
		}, actuallyFreeItAfter)()

		goRegular.SetTransform(Matrix{}, Vector{X: 1})
		if want := 1; freeCount != want {
			t.Errorf("Face.SetTransform() free count %d, want %d", freeCount, want)
		}
	})

	t.Run("non 0 matrix and vec", func(t *testing.T) {
		var freeCount int
		defer mockFree(func(v unsafe.Pointer) {
			freeCount++
		}, actuallyFreeItAfter)()

		goRegular.SetTransform(Matrix{Xx: 1}, Vector{X: 1})
		if want := 2; freeCount != want {
			t.Errorf("Face.SetTransform() free count %d, want %d", freeCount, want)
		}
	})
}

func TestRotate90Deg(t *testing.T) {
	goRegular, err := goRegular()
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goRegular.Free()
	deg := float64(90)
	angle := deg / 360.0 * math.Pi * 2.0
	m := Matrix{
		Xx: (fixed.Int16_16)(math.Cos(angle) * 0x10000),
		Xy: (fixed.Int16_16)(-math.Sin(angle) * 0x10000),
		Yx: (fixed.Int16_16)(math.Sin(angle) * 0x10000),
		Yy: (fixed.Int16_16)(math.Cos(angle) * 0x10000),
	}
	goRegular.SetTransform(m, Vector{})

	if err := goRegular.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	if err := goRegular.LoadGlyph(0x24, LoadRender|LoadColor); err != nil {
		t.Fatalf("unable to load glyph: %v", err)
	}

	got := goRegular.Glyph()
	want := GlyphSlot{
		GlyphIndex: 0x24,
		Metrics: GlyphMetrics{
			Width:        640,
			Height:       704,
			HoriBearingX: 0,
			HoriBearingY: 704,
			HoriAdvance:  576,
		},
		LinearHoriAdvance: 611968,
		LinearVertAdvance: 884352,
		Advance: Vector26_6{
			X: 0,
			Y: 576,
		},
		Format: GlyphFormatBitmap,
		Bitmap: Bitmap{
			Rows:      0xa,
			Width:     0xb,
			Pitch:     11,
			Buffer:    goRegularGlyphBuf(0x24, deg),
			NumGrays:  0x100,
			PixelMode: PixelModeGray,
		},
		BitmapLeft: -11,
		BitmapTop:  10,
		Outline: Outline{
			Points: []Vector{
				{X: 0, Y: 8}, {X: -704, Y: 254}, {X: -704, Y: 345}, {X: 0, Y: 587}, {X: 0, Y: 488}, {X: -182, Y: 421},
				{X: -182, Y: 161}, {X: 0, Y: 94}, {X: -255, Y: 187}, {X: -255, Y: 396}, {X: -565, Y: 292},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{0x0007, 0x000a},
			Flags:    0x00000130,
		},
		NumSubglyphs: 0,
		LsbDelta:     0,
		RsbDelta:     0,
	}

	// go regular is not a vertical font, vertical metrics are unreliable
	got.Metrics.VertAdvance, want.Metrics.VertAdvance = -1, -1
	got.Metrics.VertBearingX, want.Metrics.VertBearingX = -1, -1
	got.Metrics.VertBearingY, want.Metrics.VertBearingY = -1, -1

	if diff := deep.Equal(got, want); diff != nil {
		t.Error(diff)
	}
}

func TestFace_LoadGlyph(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		setup   func(*Face) error
		idx     GlyphIndex
		flags   LoadFlag
		wantErr error
	}{
		{
			name:    "nil face",
			face:    nilFace,
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name:    "go regular, no size set",
			face:    goRegular,
			idx:     0x24,
			flags:   LoadRender | LoadColor,
			wantErr: ErrInvalidSizeHandle,
		},
		{
			name: "go regular",
			face: goRegular,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			idx:     0x24,
			flags:   LoadRender | LoadColor,
			wantErr: nil,
		},
		{
			name:    "noto sans jp reg, no size",
			face:    notoSansJpReg,
			idx:     0x22,
			flags:   LoadRender | LoadColor,
			wantErr: ErrInvalidSizeHandle,
		},
		{
			name: "noto sans jp reg, horizontal",
			face: notoSansJpReg,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			idx:     0x22,
			flags:   LoadRender | LoadColor,
			wantErr: nil,
		},
		{
			name: "noto sans jp reg, vertical",
			face: notoSansJpReg,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			idx:     0x22,
			flags:   LoadRender | LoadColor | LoadVerticalLayout,
			wantErr: nil,
		},
		{
			name:    "bungee color mac, no size",
			face:    bungeeColorMac,
			idx:     0x2b,
			flags:   LoadRender | LoadColor,
			wantErr: ErrInvalidSizeHandle,
		},
		{
			name: "bungee color mac",
			face: bungeeColorMac,
			setup: func(f *Face) error {
				if err := f.SelectSize(0); err != nil {
					return fmt.Errorf("unable to select size: %v", err)
				}

				return nil
			},
			idx:     0x2b,
			flags:   LoadRender | LoadColor,
			wantErr: nil,
		},
		{
			name:    "bungee color windows, no size",
			face:    bungeeColorWin,
			idx:     0x2b,
			flags:   LoadRender | LoadColor,
			wantErr: ErrInvalidSizeHandle,
		},
		{
			name: "bungee color windows",
			face: bungeeColorWin,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			idx:     0x2b,
			flags:   LoadRender | LoadColor,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if tt.setup != nil {
				if err := tt.setup(face.Face); err != nil {
					t.Fatalf("Face.LoadGlyph() setup error: %v", err)
				}
			}
			if err := face.LoadGlyph(tt.idx, tt.flags); err != tt.wantErr {
				t.Errorf("Face.LoadGlyph() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFace_CharIndex(t *testing.T) {
	tests := []struct {
		name string
		face func() (testface, error)
		r    rune
		want GlyphIndex
	}{
		{
			name: "nil face",
			face: nilFace,
			r:    0,
			want: 0,
		},
		{
			name: "goRegular",
			face: goRegular,
			r:    'A',
			want: 0x24,
		},
		{
			name: "notoSansJpReg",
			face: notoSansJpReg,
			r:    'A',
			want: 0x22,
		},
		{
			name: "bungeeColorMac",
			face: bungeeColorMac,
			r:    'A',
			want: 0x2b,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if got := face.CharIndex(tt.r); got != tt.want {
				t.Errorf("Face.CharIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFace_FirstChar(t *testing.T) {
	tests := []struct {
		name     string
		face     func() (testface, error)
		wantRune rune
		wantIdx  GlyphIndex
	}{
		{
			name:     "nil face",
			face:     nilFace,
			wantRune: 0,
			wantIdx:  0,
		},
		{
			name:     "goRegular",
			face:     goRegular,
			wantRune: 0,
			wantIdx:  1,
		},
		{
			name:     "bungeeColorWin",
			face:     bungeeColorWin,
			wantRune: 32,
			wantIdx:  10,
		},
		{
			name:     "bungeeColorMac",
			face:     bungeeColorMac,
			wantRune: 32,
			wantIdx:  10,
		},
		{
			name:     "notoSansJpReg",
			face:     notoSansJpReg,
			wantRune: 0,
			wantIdx:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if gotRune, gotIdx := face.FirstChar(); gotRune != tt.wantRune || gotIdx != tt.wantIdx {
				t.Errorf("Face.FirstChar() got = %v, %v, want %v, %v", gotRune, gotIdx, tt.wantRune, tt.wantIdx)
			}
		})
	}
}

func TestFace_NextChar(t *testing.T) {
	tests := []struct {
		name     string
		face     func() (testface, error)
		current  rune
		wantRune rune
	}{
		{
			name:     "nil face",
			face:     nilFace,
			current:  0,
			wantRune: 0,
		},
		{
			name:     "goRegular",
			face:     goRegular,
			current:  ' ',
			wantRune: '!',
		},
		{
			name:     "bungeeColorWin",
			face:     bungeeColorWin,
			current:  ' ',
			wantRune: '!',
		},
		{
			name:     "bungeeColorMac",
			face:     bungeeColorMac,
			current:  ' ',
			wantRune: '!',
		},
		{
			name:     "notoSansJpReg",
			face:     notoSansJpReg,
			current:  ' ',
			wantRune: '!',
		},
		{
			name:     "arimoRegular",
			face:     arimoRegular,
			current:  ' ',
			wantRune: '!',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			wantIdx := face.CharIndex(tt.wantRune)
			if gotRune, gotIdx := face.NextChar(tt.current); gotRune != tt.wantRune || gotIdx != wantIdx {
				t.Errorf("Face.NextChar(%d) got = %v, %v, want %v, %v", tt.current, gotRune, gotIdx, tt.wantRune, wantIdx)
			}
		})
	}
}

func TestFace_IndexOf(t *testing.T) {
	tests := []struct {
		name      string
		face      func() (testface, error)
		glyphName string
		want      GlyphIndex
	}{
		{
			name:      "nil face",
			face:      nilFace,
			glyphName: "",
			want:      0,
		},
		{
			name:      "goRegular",
			face:      goRegular,
			glyphName: "A",
			want:      0x24,
		},
		{
			name:      "bungeeColorWin",
			face:      bungeeColorWin,
			glyphName: "A",
			want:      0x2b,
		},
		{
			name:      "bungeeColorMac",
			face:      bungeeColorMac,
			glyphName: "A",
			want:      0x2b,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if got := face.IndexOf(tt.glyphName); got != tt.want {
				t.Errorf("Face.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFace_LoadChar(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		setup   func(*Face) error
		r       rune
		flags   LoadFlag
		wantErr error
	}{
		{
			name:    "nil face",
			face:    nilFace,
			r:       0,
			flags:   0,
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name:    "go regular, no size set",
			face:    goRegular,
			r:       'A',
			flags:   LoadRender | LoadColor,
			wantErr: ErrInvalidSizeHandle,
		},
		{
			name: "go regular",
			face: goRegular,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			r:       'A',
			flags:   LoadRender | LoadColor,
			wantErr: nil,
		},
		{
			name:    "noto sans jp reg, no size",
			face:    notoSansJpReg,
			r:       'A',
			flags:   LoadRender | LoadColor,
			wantErr: ErrInvalidSizeHandle,
		},
		{
			name: "noto sans jp reg, horizontal",
			face: notoSansJpReg,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			r:       'A',
			flags:   LoadRender | LoadColor,
			wantErr: nil,
		},
		{
			name: "noto sans jp reg, vertical",
			face: notoSansJpReg,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			r:       'A',
			flags:   LoadRender | LoadColor | LoadVerticalLayout,
			wantErr: nil,
		},
		{
			name:    "bungee color mac, no size",
			face:    bungeeColorMac,
			r:       'A',
			flags:   LoadRender | LoadColor,
			wantErr: ErrInvalidSizeHandle,
		},
		{
			name: "bungee color mac",
			face: bungeeColorMac,
			setup: func(f *Face) error {
				if err := f.SelectSize(0); err != nil {
					return fmt.Errorf("unable to select size: %v", err)
				}

				return nil
			},
			r:       'A',
			flags:   LoadRender | LoadColor,
			wantErr: nil,
		},
		{
			name:    "bungee color windows, no size",
			face:    bungeeColorWin,
			r:       'A',
			flags:   LoadRender | LoadColor,
			wantErr: ErrInvalidSizeHandle,
		},
		{
			name: "bungee color windows",
			face: bungeeColorWin,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			r:       'A',
			flags:   LoadRender | LoadColor,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if tt.setup != nil {
				if err := tt.setup(face.Face); err != nil {
					t.Fatalf("Face.LoadChar() setup error: %v", err)
				}
			}

			if err := face.LoadChar(tt.r, tt.flags); err != tt.wantErr {
				t.Errorf("Face.LoadChar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFace_Kerning(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		setup   func(*Face) error
		left    GlyphIndex
		right   GlyphIndex
		mode    KerningMode
		want    Vector
		wantErr error
	}{
		{
			name:    "nil face",
			face:    nilFace,
			want:    Vector{},
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name: "goRegular",
			face: goRegular,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			left:    0x24,
			right:   0x39,
			mode:    KerningModeDefault,
			want:    Vector{},
			wantErr: nil,
		},
		{
			name: "arimoRegular, default",
			face: arimoRegular,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			left:    0x24,
			right:   0x39,
			mode:    KerningModeDefault,
			want:    Vector{X: -64, Y: 0},
			wantErr: nil,
		},
		{
			name: "arimoRegular, unfitted",
			face: arimoRegular,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			left:    0x24,
			right:   0x39,
			mode:    KerningModeUnfitted,
			want:    Vector{X: -67, Y: 0},
			wantErr: nil,
		},
		{
			name: "arimoRegular, unscaled",
			face: arimoRegular,
			setup: func(f *Face) error {
				if err := f.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					return fmt.Errorf("unable to set char size: %v", err)
				}

				return nil
			},
			left:    0x24,
			right:   0x39,
			mode:    KerningModeUnscaled,
			want:    Vector{X: -152, Y: 0},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if tt.setup != nil {
				if err := tt.setup(face.Face); err != nil {
					t.Fatalf("Face.LoadChar() setup error: %v", err)
				}
			}

			if got, err := face.Kern(tt.left, tt.right, tt.mode); got != tt.want || err != tt.wantErr {
				t.Errorf("Face.Kerning() = %v, %v, want %v, %v", got, err, tt.want, tt.wantErr)
				return
			}
		})
	}
}

func TestFace_GlyphName(t *testing.T) {
	t.Run("free", func(t *testing.T) {
		face, err := goRegular()
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		var freed bool
		defer mockFree(func(_ unsafe.Pointer) {
			freed = true
		}, actuallyFreeItAfter)()

		if _, err := face.GlyphName(0); err != nil {
			t.Errorf("Face.GlyphName() error = %v", err)
		}

		if !freed {
			t.Errorf("Face.GlyphName() free was not")
		}
	})

	t.Run("free on error", func(t *testing.T) {
		face, err := goRegular()
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		var freed bool
		defer mockFree(func(_ unsafe.Pointer) {
			freed = true
		}, actuallyFreeItAfter)()

		wantErr := ErrBbxTooBig
		defer mockGetErr(func(_ int) error {
			return ErrBbxTooBig
		})()

		if _, err := face.GlyphName(0); err != wantErr {
			t.Errorf("Face.GlyphName() error = %v, want %v", err, wantErr)
		}

		if !freed {
			t.Errorf("Face.GlyphName() free was not")
		}
	})

	tests := []struct {
		name    string
		face    func() (testface, error)
		idx     GlyphIndex
		want    string
		wantErr error
	}{
		{
			name:    "nil face",
			face:    nilFace,
			idx:     0,
			want:    "",
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name:    "goRegular .notdef",
			face:    goRegular,
			idx:     0x00,
			want:    ".notdef",
			wantErr: nil,
		},
		{
			name:    "goRegular A",
			face:    goRegular,
			idx:     0x24,
			want:    "A",
			wantErr: nil,
		},
		{
			name:    "goRegular question",
			face:    goRegular,
			idx:     0x22,
			want:    "question",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if got, err := face.GlyphName(tt.idx); got != tt.want || err != tt.wantErr {
				t.Errorf("Face.GlyphName() = %q, %v, want %v, %v", got, err, tt.want, tt.wantErr)
				return
			}
		})
	}
}

func TestFace_FSTypeFlags(t *testing.T) {
	tests := []struct {
		name string
		face func() (testface, error)
		want FSTypeFlag
	}{
		{
			name: "nil face",
			face: nilFace,
			want: FsTypeFlagInstallableEmbedding,
		},
		{
			name: "goRegular",
			face: goRegular,
			want: FsTypeFlagInstallableEmbedding,
		},
		{
			name: "bungeeColorMac",
			face: bungeeColorMac,
			want: 1,
		},
		{
			name: "bungeeColorWin",
			face: bungeeColorWin,
			want: 1,
		},
		{
			name: "notoSansJpReg",
			face: notoSansJpReg,
			want: FsTypeFlagInstallableEmbedding,
		},
		{
			name: "arimoRegular",
			face: arimoRegular,
			want: FsTypeFlagEditableEmbedding,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if got := face.FSTypeFlags(); got != tt.want {
				t.Errorf("Face.FSTypeFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSTypeFlag_String(t *testing.T) {
	var x FSTypeFlag
	if got, want := x.String(), "InstallableEmbedding"; got != want {
		t.Errorf("FSTypeFlag.String() = %v, want %v", got, want)
	}

	x = FsTypeFlagRestrictedLicenseEmbedding
	if got, want := x.String(), "RestrictedLicenseEmbedding"; got != want {
		t.Errorf("FSTypeFlag.String(FsTypeFlagRestrictedLicenseEmbedding) = %v, want %v", got, want)
	}

	x = FsTypeFlagRestrictedLicenseEmbedding | FsTypeFlagPreviewAndPrintEmbedding
	if got, want := x.String(), "RestrictedLicenseEmbedding|PreviewAndPrintEmbedding"; got != want {
		t.Errorf("FSTypeFlag.String(FsTypeFlagRestrictedLicenseEmbedding | FsTypeFlagPreviewAndPrintEmbedding) = %v, want %v", got, want)
	}

	x = FsTypeFlagInstallableEmbedding | FsTypeFlagRestrictedLicenseEmbedding |
		FsTypeFlagPreviewAndPrintEmbedding | FsTypeFlagEditableEmbedding |
		FsTypeFlagNoSubsetting | FsTypeFlagBitmapEmbeddingOnly
	if got, want := x.String(), "RestrictedLicenseEmbedding|PreviewAndPrintEmbedding|EditableEmbedding|NoSubsetting|BitmapEmbeddingOnly"; got != want {
		t.Errorf("FSTypeFlag.String(FsTypeFlagInstallableEmbedding | FsTypeFlagRestrictedLicenseEmbedding | FsTypeFlagPreviewAndPrintEmbedding | FsTypeFlagEditableEmbedding | FsTypeFlagNoSubsetting | FsTypeFlagBitmapEmbeddingOnly) = %v, want %v", got, want)
	}
}

func TestRenderMode_String(t *testing.T) {
	tests := []struct {
		name string
		r    RenderMode
		want string
	}{
		{name: "Normal", r: RenderModeNormal, want: "Normal"},
		{name: "Light", r: RenderModeLight, want: "Light"},
		{name: "Mono", r: RenderModeMono, want: "Mono"},
		{name: "LCD", r: RenderModeLCD, want: "LCD"},
		{name: "LCDV", r: RenderModeLCDV, want: "LCDV"},
		{name: "Unknown", r: 90102, want: "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("RenderMode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFace_SubGlyphInfo(t *testing.T) {
	t.Skip("I could not find a font with composite glyphs")
}

func TestSubGlyphFlag_String(t *testing.T) {
	var x SubGlyphFlag
	if got, want := x.String(), ""; got != want {
		t.Errorf("SubGlyphFlag.String() = %v, want %v", got, want)
	}

	x = SubGlyphFlagXyScale
	if got, want := x.String(), "XyScale"; got != want {
		t.Errorf("SubGlyphFlag.String(SubGlyphFlagXyScale) = %v, want %v", got, want)
	}

	x = SubGlyphFlagArgsAreXyValues | SubGlyphFlagScale
	if got, want := x.String(), "ArgsAreXyValues|Scale"; got != want {
		t.Errorf("SubGlyphFlag.String(SubGlyphFlagArgsAreXyValues | SubGlyphFlagScale) = %v, want %v", got, want)
	}

	x = SubGlyphFlagArgsAreWords | SubGlyphFlagArgsAreXyValues |
		SubGlyphFlagRoundXyToGrid | SubGlyphFlagScale | SubGlyphFlagXyScale |
		SubGlyphFlag2x2 | SubGlyphFlagUseMyMetrics
	if got, want := x.String(), "ArgsAreWords|ArgsAreXyValues|RoundXyToGrid|Scale|XyScale|2x2|UseMyMetrics"; got != want {
		t.Errorf("SubGlyphFlag.String(SubGlyphFlagArgsAreWords | SubGlyphFlagArgsAreXyValues | SubGlyphFlagRoundXyToGrid | SubGlyphFlagScale | SubGlyphFlagXyScale | SubGlyphFlag2x2 | SubGlyphFlagUseMyMetrics) = %v, want %v", got, want)
	}
}
