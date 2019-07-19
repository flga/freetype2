package freetype2

import (
	"testing"
)

func TestFaceFlags_String(t *testing.T) {
	var x FaceFlags
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
	var x StyleFlags

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

	f, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0)
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
