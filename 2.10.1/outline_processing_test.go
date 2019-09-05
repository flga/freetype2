package freetype2

import (
	"fmt"
	"math"
	"testing"

	"github.com/flga/freetype2/fixed"
)

func Test_Outline_reload(t *testing.T) {
	face, err := notoSansJpReg()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	if err := face.LoadGlyph(0x22, LoadRender|LoadColor); err != nil {
		t.Fatalf("unable laod glyph: %v", err)
	}

	got := &Outline{ptr: &face.ptr.glyph.outline}
	got.reload()
	want := &Outline{
		Points: []Vector{
			{X: 0x000000ab, Y: 0x00000102}, {X: 0x000000cb, Y: 0x00000178}, {X: 0x000000e2, Y: 0x000001cf},
			{X: 0x000000f8, Y: 0x00000222}, {X: 0x0000010c, Y: 0x0000027c}, {X: 0x00000110, Y: 0x0000027c},
			{X: 0x00000126, Y: 0x00000223}, {X: 0x0000013a, Y: 0x000001cf}, {X: 0x00000152, Y: 0x00000178},
			{X: 0x00000172, Y: 0x00000102}, {X: 0x000001c5, Y: 0x00000000}, {X: 0x0000021d, Y: 0x00000000},
			{X: 0x0000013e, Y: 0x000002c0}, {X: 0x000000e1, Y: 0x000002c0}, {X: 0x00000002, Y: 0x00000000},
			{X: 0x00000056, Y: 0x00000000}, {X: 0x00000096, Y: 0x000000c0}, {X: 0x00000186, Y: 0x000000c0},
		},
		Tags: []byte{
			0x01, 0x01, 0x02, 0x02, 0x01, 0x01, 0x02, 0x02, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		},
		Contours: []int16{0x0009, 0x0011},
		Flags:    0x00000104,
	}

	if diff := diff(got, want); diff != nil {
		t.Errorf("newOutline() %v", diff)
	}
}

func TestOutlineFlag_String(t *testing.T) {
	tests := []struct {
		name string
		x    OutlineFlag
		want string
	}{
		{name: "0", x: 0, want: ""},
		{name: "one", x: OutlineIgnoreDropouts, want: "IgnoreDropouts"},
		{name: "two", x: OutlineEvenOddFill | OutlineIncludeStubs, want: "EvenOddFill|IncludeStubs"},
		{
			name: "all",
			x: OutlineOwner | OutlineEvenOddFill | OutlineReverseFill | OutlineIgnoreDropouts |
				OutlineSmartDropouts | OutlineIncludeStubs | OutlineHighPrecision | OutlineSinglePass,
			want: "Owner|EvenOddFill|ReverseFill|IgnoreDropouts|SmartDropouts|IncludeStubs|HighPrecision|SinglePass",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("OutlineFlag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutline_Copy(t *testing.T) {
	t.Run("nil source", func(t *testing.T) {
		var o Outline
		if err := o.CopyTo(nil); err != ErrInvalidOutline {
			t.Errorf("Outline.CopyTo() error = %v, want %v", err, ErrInvalidOutline)
		}
	})

	t.Run("nil target", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}
		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		if err := face.GlyphSlot().Outline.CopyTo(nil); err != ErrInvalidOutline {
			t.Errorf("Outline.CopyTo() error = %v, want %v", err, ErrInvalidOutline)
		}
	})

	t.Run("different size", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}
		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		OutlineGlyphA, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}

		if err := face.LoadChar('B', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot = face.GlyphSlot()
		OutlineGlyphB, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}

		outlineA, ok := OutlineGlyphA.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline")
		}

		outlineB, ok := OutlineGlyphB.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline")
		}

		if err := outlineA.Outline.CopyTo(outlineB.Outline); err != ErrInvalidArgument {
			t.Errorf("Outline.CopyTo() error = %v, want %v", err, ErrInvalidArgument)
		}
	})

	t.Run("ok", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}
		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		glyph, err := face.GlyphSlot().Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		glyphCopy1, err := glyph.Copy()
		if err != nil {
			t.Fatalf("unable to copy glyph: %v", err)
		}
		glyphCopy2, err := glyph.Copy()
		if err != nil {
			t.Fatalf("unable to copy glyph: %v", err)
		}

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline")
		}
		outlineGlyphCopy1, ok := glyphCopy1.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline")
		}
		outlineGlyphCopy2, ok := glyphCopy2.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline")
		}

		outline := outlineGlyph.Outline
		outlineCopy1 := outlineGlyphCopy1.Outline
		outlineCopy2 := outlineGlyphCopy2.Outline

		wantOriginal := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		wantReversed := &Outline{
			Points: []Vector{
				{94, 0},
				{161, 182},
				{421, 182},
				{488, 0},
				{587, 0},
				{345, 704},
				{254, 704},
				{8, 0},
				{292, 565},
				{396, 255},
				{187, 255},
			},
			Tags:     []byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x95, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131 | OutlineReverseFill,
		}

		// baseline
		if diff := diff(outline, wantOriginal); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}
		if diff := diff(outlineCopy1, wantOriginal); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}
		if diff := diff(outlineCopy2, wantOriginal); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}

		// reverse
		outlineCopy1.Reverse()
		if diff := diff(outline, wantOriginal); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}
		if diff := diff(outlineCopy1, wantReversed); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}
		if diff := diff(outlineCopy2, wantOriginal); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}

		// copy 1
		if err := outlineCopy1.CopyTo(outlineCopy2); err != nil {
			t.Errorf("Outline.CopyTo() error = %v", err)
		}
		if diff := diff(outline, wantOriginal); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}
		if diff := diff(outlineCopy1, wantReversed); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}
		if diff := diff(outlineCopy2, wantReversed); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}

		// copy 2
		if err := outline.CopyTo(outlineCopy1); err != nil {
			t.Errorf("Outline.CopyTo() error = %v", err)
		}
		if diff := diff(outline, wantOriginal); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}
		if diff := diff(outlineCopy1, wantOriginal); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}
		if diff := diff(outlineCopy2, wantReversed); diff != nil {
			t.Errorf("Outline.CopyTo() %v", diff)
		}
	})
}

func TestOutline_Translate(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var o Outline
		o.Translate(0, 0)
	})

	t.Run("glyph slot - glyph", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Translate() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{72, 64},
				{318, 768},
				{409, 768},
				{651, 64},
				{552, 64},
				{485, 246},
				{225, 246},
				{158, 64},
				{251, 319},
				{460, 319},
				{356, 629},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		outline.Translate(64, 64)
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Translate() = %v", diff)
		}
	})

	t.Run("glyph slot - direct", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		outline := slot.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Translate() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{72, 64},
				{318, 768},
				{409, 768},
				{651, 64},
				{552, 64},
				{485, 246},
				{225, 246},
				{158, 64},
				{251, 319},
				{460, 319},
				{356, 629},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		outline.Translate(64, 64)
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Translate() = %v", diff)
		}
	})

	t.Run("no overwrite", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Translate() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{72, 64},
				{318, 768},
				{409, 768},
				{651, 64},
				{552, 64},
				{485, 246},
				{225, 246},
				{158, 64},
				{251, 319},
				{460, 319},
				{356, 629},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		outline.Translate(64, 64)
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Translate() = %v", diff)
		}

		wantInSlot := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(slot.Outline, wantInSlot); diff != nil {
			t.Fatalf("Outline.Translate() = %v", diff)
		}
	})
}

func TestOutline_Transform(t *testing.T) {
	deg := float64(90)
	angle := deg / 360.0 * math.Pi * 2.0
	matrix := Matrix{
		Xx: (fixed.Int16_16)(math.Cos(angle) * 0x10000),
		Xy: (fixed.Int16_16)(-math.Sin(angle) * 0x10000),
		Yx: (fixed.Int16_16)(math.Sin(angle) * 0x10000),
		Yy: (fixed.Int16_16)(math.Cos(angle) * 0x10000),
	}

	t.Run("nil", func(t *testing.T) {
		var o Outline
		o.Transform(matrix)
	})

	t.Run("glyph slot - glyph", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Transform() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{0, 8},
				{-704, 254},
				{-704, 345},
				{0, 587},
				{0, 488},
				{-182, 421},
				{-182, 161},
				{0, 94},
				{-255, 187},
				{-255, 396},
				{-565, 292},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		outline.Transform(matrix)
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Transform() = %v", diff)
		}
	})

	t.Run("glyph slot - direct", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		outline := slot.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Transform() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{0, 8},
				{-704, 254},
				{-704, 345},
				{0, 587},
				{0, 488},
				{-182, 421},
				{-182, 161},
				{0, 94},
				{-255, 187},
				{-255, 396},
				{-565, 292},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		outline.Transform(matrix)
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Transform() = %v", diff)
		}
	})

	t.Run("no overwrite", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Transform() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{0, 8},
				{-704, 254},
				{-704, 345},
				{0, 587},
				{0, 488},
				{-182, 421},
				{-182, 161},
				{0, 94},
				{-255, 187},
				{-255, 396},
				{-565, 292},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		outline.Transform(matrix)
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Transform() = %v", diff)
		}

		wantInSlot := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(slot.Outline, wantInSlot); diff != nil {
			t.Fatalf("Outline.Transform() = %v", diff)
		}
	})
}

func TestOutline_Embolden(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var o Outline
		if err := o.Embolden(10 << 6); err != ErrInvalidOutline {
			t.Errorf("Outline.Embolden() error = %v, want %v", err, ErrInvalidOutline)
		}
	})

	t.Run("glyph slot - glyph", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Embolden() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{-123, 0},
				{347, 1344},
				{893, 1344},
				{1355, 0},
				{585, 0},
				{547, 224},
				{675, 224},
				{637, 0},
				{716, 725},
				{507, 725},
				{611, 540},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if err := outline.Embolden(10 << 6); err != nil {
			t.Fatalf("Outline.Embolden() error = %v", err)
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Embolden() = %v", diff)
		}
	})

	t.Run("glyph slot - direct", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		outline := slot.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Embolden() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{-123, 0},
				{347, 1344},
				{893, 1344},
				{1355, 0},
				{585, 0},
				{547, 224},
				{675, 224},
				{637, 0},
				{716, 725},
				{507, 725},
				{611, 540},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if err := outline.Embolden(10 << 6); err != nil {
			t.Fatalf("Outline.Embolden() error = %v", err)
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Embolden() = %v", diff)
		}
	})

	t.Run("no overwrite", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Embolden() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{-123, 0},
				{347, 1344},
				{893, 1344},
				{1355, 0},
				{585, 0},
				{547, 224},
				{675, 224},
				{637, 0},
				{716, 725},
				{507, 725},
				{611, 540},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if err := outline.Embolden(10 << 6); err != nil {
			t.Fatalf("Outline.Embolden() error = %v", err)
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Embolden() = %v", diff)
		}

		wantInSlot := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(slot.Outline, wantInSlot); diff != nil {
			t.Fatalf("Outline.Embolden() = %v", diff)
		}
	})
}

func TestOutline_EmboldenXY(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var o Outline
		if err := o.EmboldenXY(10<<6, 10<<6); err != ErrInvalidOutline {
			t.Errorf("Outline.EmboldenXY() error = %v, want %v", err, ErrInvalidOutline)
		}
	})

	t.Run("glyph slot - glyph", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.EmboldenXY() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{-123, 0},
				{347, 1344},
				{893, 1344},
				{1355, 0},
				{585, 0},
				{547, 224},
				{675, 224},
				{637, 0},
				{716, 725},
				{507, 725},
				{611, 540},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if err := outline.EmboldenXY(10<<6, 10<<6); err != nil {
			t.Fatalf("Outline.EmboldenXY() error = %v", err)
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.EmboldenXY() = %v", diff)
		}
	})

	t.Run("glyph slot - direct", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		outline := slot.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.EmboldenXY() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{-123, 0},
				{347, 1344},
				{893, 1344},
				{1355, 0},
				{585, 0},
				{547, 224},
				{675, 224},
				{637, 0},
				{716, 725},
				{507, 725},
				{611, 540},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if err := outline.EmboldenXY(10<<6, 10<<6); err != nil {
			t.Fatalf("Outline.EmboldenXY() error = %v", err)
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.EmboldenXY() = %v", diff)
		}
	})

	t.Run("no overwrite", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.EmboldenXY() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{-123, 0},
				{347, 1344},
				{893, 1344},
				{1355, 0},
				{585, 0},
				{547, 224},
				{675, 224},
				{637, 0},
				{716, 725},
				{507, 725},
				{611, 540},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if err := outline.EmboldenXY(10<<6, 10<<6); err != nil {
			t.Fatalf("Outline.EmboldenXY() error = %v", err)
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.EmboldenXY() = %v", diff)
		}

		wantInSlot := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(slot.Outline, wantInSlot); diff != nil {
			t.Fatalf("Outline.EmboldenXY() = %v", diff)
		}
	})
}

func TestOutline_Reverse(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var o Outline
		o.Reverse()
	})

	t.Run("glyph slot - glyph", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Reverse() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{94, 0},
				{161, 182},
				{421, 182},
				{488, 0},
				{587, 0},
				{345, 704},
				{254, 704},
				{8, 0},
				{292, 565},
				{396, 255},
				{187, 255},
			},
			Tags:     []byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x95, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131 | OutlineReverseFill,
		}
		outline.Reverse()
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Reverse() = %v", diff)
		}
	})

	t.Run("glyph slot - direct", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		outline := slot.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Reverse() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{94, 0},
				{161, 182},
				{421, 182},
				{488, 0},
				{587, 0},
				{345, 704},
				{254, 704},
				{8, 0},
				{292, 565},
				{396, 255},
				{187, 255},
			},
			Tags:     []byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x95, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130 | OutlineReverseFill,
		}
		outline.Reverse()
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Reverse() = %v", diff)
		}
	})

	t.Run("no overwrite", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}

		if err := face.LoadChar('A', LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}

		slot := face.GlyphSlot()
		glyph, err := slot.Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		defer glyph.Free()

		outlineGlyph, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("glyph is not an outline: %v", err)
		}

		outline := outlineGlyph.Outline
		want := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131,
		}
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Reverse() = %v", diff)
		}

		want = &Outline{
			Points: []Vector{
				{94, 0},
				{161, 182},
				{421, 182},
				{488, 0},
				{587, 0},
				{345, 704},
				{254, 704},
				{8, 0},
				{292, 565},
				{396, 255},
				{187, 255},
			},
			Tags:     []byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x95, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x131 | OutlineReverseFill,
		}
		outline.Reverse()
		if diff := diff(outline, want); diff != nil {
			t.Fatalf("Outline.Reverse() = %v", diff)
		}

		wantInSlot := &Outline{
			Points: []Vector{
				{8, 0},
				{254, 704},
				{345, 704},
				{587, 0},
				{488, 0},
				{421, 182},
				{161, 182},
				{94, 0},
				{187, 255},
				{396, 255},
				{292, 565},
			},
			Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
			Contours: []int16{7, 10},
			Flags:    0x130,
		}
		if diff := diff(slot.Outline, wantInSlot); diff != nil {
			t.Fatalf("Outline.Reverse() = %v", diff)
		}
	})
}

func TestOutline_Check(t *testing.T) {
	var o Outline
	if valid := o.Check(); valid != false {
		t.Errorf("Outline.Check() = %v, want %v", valid, false)
	}

	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %v", err)
	}
	defer l.Free()

	face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	if err := face.LoadChar('A', LoadDefault); err != nil {
		t.Fatalf("unable to load char: %v", err)
	}

	slot := face.GlyphSlot()

	if valid := slot.Outline.Check(); valid != true {
		t.Errorf("Outline.Check() = %v, want %v", valid, true)
	}

	glyph, err := slot.Glyph()
	if err != nil {
		t.Fatalf("unable to get glyph: %v", err)
	}
	defer glyph.Free()

	outlineGlyph, ok := glyph.(*OutlineGlyph)
	if !ok {
		t.Fatalf("glyph is not an outline: %v", err)
	}

	if valid := outlineGlyph.Outline.Check(); valid != true {
		t.Errorf("Outline.Check() = %v, want %v", valid, true)
	}

}

func TestOutline_CBox(t *testing.T) {
	var o Outline
	var want BBox
	if got := o.CBox(); got != want {
		t.Errorf("Outline.CBox() = %v, want %v", got, want)
	}

	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %v", err)
	}
	defer l.Free()

	face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	if err := face.LoadChar('A', LoadDefault); err != nil {
		t.Fatalf("unable to load char: %v", err)
	}

	slot := face.GlyphSlot()
	glyph, err := slot.Glyph()
	if err != nil {
		t.Fatalf("unable to get glyph: %v", err)
	}
	defer glyph.Free()

	outlineGlyph, ok := glyph.(*OutlineGlyph)
	if !ok {
		t.Fatalf("glyph is not an outline: %v", err)
	}

	want = BBox{
		XMin: 8,
		XMax: 587,
		YMin: 0,
		YMax: 704,
	}

	if got := slot.Outline.CBox(); got != want {
		t.Errorf("Outline.CBox() = %v, want %v", got, want)
	}

	if got := outlineGlyph.Outline.CBox(); got != want {
		t.Errorf("Outline.CBox() = %v, want %v", got, want)
	}
}

func TestOutline_BBox(t *testing.T) {
	var o Outline
	var want BBox
	if got := o.BBox(); got != want {
		t.Errorf("Outline.BBox() = %v, want %v", got, want)
	}

	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %v", err)
	}
	defer l.Free()

	face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	if err := face.LoadChar('A', LoadDefault); err != nil {
		t.Fatalf("unable to load char: %v", err)
	}

	slot := face.GlyphSlot()
	glyph, err := slot.Glyph()
	if err != nil {
		t.Fatalf("unable to get glyph: %v", err)
	}
	defer glyph.Free()

	outlineGlyph, ok := glyph.(*OutlineGlyph)
	if !ok {
		t.Fatalf("glyph is not an outline: %v", err)
	}

	want = BBox{
		XMin: 8,
		XMax: 587,
		YMin: 0,
		YMax: 704,
	}

	if got := slot.Outline.BBox(); got != want {
		t.Errorf("Outline.BBox() = %v, want %v", got, want)
	}

	if got := outlineGlyph.Outline.BBox(); got != want {
		t.Errorf("Outline.BBox() = %v, want %v", got, want)
	}
}

func TestOutline_Orientation(t *testing.T) {
	var o Outline
	if got := o.Orientation(); got != OrientationNone {
		t.Errorf("Outline.Orientation() = %v, want %v", got, OrientationNone)
	}

	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %v", err)
	}
	defer l.Free()

	face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	if err := face.LoadChar('A', LoadDefault); err != nil {
		t.Fatalf("unable to load char: %v", err)
	}

	slot := face.GlyphSlot()
	glyph, err := slot.Glyph()
	if err != nil {
		t.Fatalf("unable to get glyph: %v", err)
	}
	defer glyph.Free()

	outlineGlyph, ok := glyph.(*OutlineGlyph)
	if !ok {
		t.Fatalf("glyph is not an outline: %v", err)
	}

	{
		if got := slot.Outline.Orientation(); got != OrientationFillRight {
			t.Errorf("Outline.Orientation() = %v, want %v", got, OrientationFillRight)
		}
		if got := outlineGlyph.Outline.Orientation(); got != OrientationFillRight {
			t.Errorf("Outline.Orientation() = %v, want %v", got, OrientationFillRight)
		}
	}

	{
		slot.Outline.Reverse()
		if got := slot.Outline.Orientation(); got != OrientationFillLeft {
			t.Errorf("Outline.Orientation() = %v, want %v", got, OrientationFillLeft)
		}
		if got := outlineGlyph.Outline.Orientation(); got != OrientationFillRight {
			t.Errorf("Outline.Orientation() = %v, want %v", got, OrientationFillRight)
		}
		slot.Outline.Reverse()
	}

	{
		outlineGlyph.Outline.Reverse()
		if got := slot.Outline.Orientation(); got != OrientationFillRight {
			t.Errorf("Outline.Orientation() = %v, want %v", got, OrientationFillRight)
		}
		if got := outlineGlyph.Outline.Orientation(); got != OrientationFillLeft {
			t.Errorf("Outline.Orientation() = %v, want %v", got, OrientationFillLeft)
		}
		outlineGlyph.Outline.Reverse()
	}

	{
		outlineGlyph.Outline.Reverse()
		slot.Outline.Reverse()
		if got := slot.Outline.Orientation(); got != OrientationFillLeft {
			t.Errorf("Outline.Orientation() = %v, want %v", got, OrientationFillLeft)
		}
		if got := outlineGlyph.Outline.Orientation(); got != OrientationFillLeft {
			t.Errorf("Outline.Orientation() = %v, want %v", got, OrientationFillLeft)
		}
	}
}

func TestLibrary_NewOutline(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %v", err)
	}
	defer l.Free()

	tests := []struct {
		name     string
		l        *Library
		points   int
		contours int
		want     *Outline
		wantErr  error
	}{
		{name: "nil lib", l: nil, wantErr: ErrInvalidLibraryHandle},

		{name: "negative", l: l, points: -1, wantErr: ErrInvalidArgument},
		{name: "negative", l: l, contours: -1, wantErr: ErrInvalidArgument},
		{name: "negative", l: l, points: -1, contours: -1, wantErr: ErrInvalidArgument},

		{name: "oob", l: l, points: 1 << 30, wantErr: ErrArrayTooLarge},

		{name: "too many contours", l: l, points: 10, contours: 100, wantErr: ErrInvalidArgument},

		{
			name:     "ok",
			l:        l,
			points:   10,
			contours: 10,
			want: &Outline{
				userCreated: true,
				Points:      make([]Vector, 10),
				Tags:        make([]byte, 10),
				Contours:    make([]int16, 10),
				Flags:       OutlineOwner,
			},
		},
		{
			name:     "ok",
			l:        l,
			points:   10,
			contours: 0,
			want: &Outline{
				userCreated: true,
				Points:      make([]Vector, 10),
				Tags:        make([]byte, 10),
				Contours:    nil,
				Flags:       OutlineOwner,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "ok" {
				fmt.Print()
			}
			got, err := tt.l.NewOutline(tt.points, tt.contours)
			defer got.Free()
			if err != tt.wantErr {
				t.Errorf("Library.NewOutline() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Library.NewOutline() = %v", diff)
			}
		})
	}
}

type dummyDecomposer int

func (d dummyDecomposer) MoveTo(Vector) error                         { return nil }
func (d dummyDecomposer) LineTo(Vector) error                         { return nil }
func (d dummyDecomposer) ConicTo(control, to Vector) error            { return nil }
func (d dummyDecomposer) CubicTo(control1, control2, to Vector) error { return nil }

func TestDecomposerTable(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		table := &decomposerTable{
			table: make(map[uintptr]OutlineDecomposer),
		}

		value := dummyDecomposer(42)
		_, handle, err := table.acquire(value, 0, 0)
		if err != nil {
			t.Fatalf("err %v", err)
		}

		if handle != 0 {
			t.Fatalf("handle should be 0")
		}

		if got := table.valueOf(handle); got != value {
			t.Fatalf("got %v, want %v", got, value)
		}

		table.release(handle)
		if got := table.valueOf(handle); got != nil {
			t.Fatalf("got %v, want %v", got, nil)
		}
	})

	t.Run("overlap", func(t *testing.T) {
		table := &decomposerTable{
			table: make(map[uintptr]OutlineDecomposer),
		}

		value1 := dummyDecomposer(42)
		value2 := dummyDecomposer(420)
		_, handle1, err := table.acquire(value1, 0, 0)
		if err != nil {
			t.Fatalf("err %v", err)
		}

		table.idx = 0 //simulate overflow
		_, handle2, err := table.acquire(value2, 0, 0)
		if err != nil {
			t.Fatalf("err %v", err)
		}

		if got := table.valueOf(handle1); got != value1 {
			t.Fatalf("got %v, want %v", got, value1)
		}
		if got := table.valueOf(handle2); got != value2 {
			t.Fatalf("got %v, want %v", got, value2)
		}

		table.release(handle1)
		if got := table.valueOf(handle1); got != nil {
			t.Fatalf("got %v, want %v", got, nil)
		}
		if got := table.valueOf(handle2); got != value2 {
			t.Fatalf("got %v, want %v", got, value2)
		}

		table.release(handle2)
		if got := table.valueOf(handle1); got != nil {
			t.Fatalf("got %v, want %v", got, nil)
		}
		if got := table.valueOf(handle2); got != nil {
			t.Fatalf("got %v, want %v", got, nil)
		}
	})
}

type stackRendererFrame struct {
	y     int
	spans []Span
}

type stackRenderer struct {
	frames []stackRendererFrame
}

func (r *stackRenderer) GraySpans(y int, spans []Span) {
	r.frames = append(r.frames, stackRendererFrame{y, spans})
}

func TestOutlineRender(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		defer l.Free()

		newOutline, err := l.NewOutline(10, 10)
		if err != nil {
			t.Fatalf("unable to create outline: %v", err)
		}
		defer newOutline.Free()

		var nilOutline *Outline

		if err := nilOutline.Render(nil, RasterParams{}); err != ErrInvalidOutline {
			t.Errorf("Outline.Render() error = %v, want %v", err, ErrInvalidOutline)
		}

		if err := newOutline.Render(nil, RasterParams{}); err != ErrInvalidLibraryHandle {
			t.Errorf("Outline.Render() error = %v, want %v", err, ErrInvalidLibraryHandle)
		}
	})

	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %v", err)
	}
	defer l.Free()

	face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	if err := face.LoadChar('A', LoadDefault); err != nil {
		t.Fatalf("unable to load char: %v", err)
	}

	want := []stackRendererFrame{
		{
			y: 0,
			spans: []Span{
				{X: 0, Len: 1, Coverage: 0xb3},
				{X: 1, Len: 1, Coverage: 0xa7},
				{X: 7, Len: 1, Coverage: 0x8f},
				{X: 8, Len: 1, Coverage: 0xf4},
				{X: 9, Len: 1, Coverage: 0x0a},
			},
		},
		{
			y: 1,
			spans: []Span{
				{X: 0, Len: 1, Coverage: 0x5a},
				{X: 1, Len: 1, Coverage: 0xf6},
				{X: 2, Len: 1, Coverage: 0x0e},
				{X: 6, Len: 1, Coverage: 0x04},
				{X: 7, Len: 1, Coverage: 0xe9},
				{X: 8, Len: 1, Coverage: 0xa7},
			},
		},
		{
			y: 2,
			spans: []Span{
				{X: 0, Len: 1, Coverage: 0x0b},
				{X: 1, Len: 1, Coverage: 0xf5},
				{X: 2, Len: 1, Coverage: 0x75},
				{X: 3, Len: 3, Coverage: 0x28},
				{X: 6, Len: 1, Coverage: 0x61},
				{X: 7, Len: 1, Coverage: 0xff},
				{X: 8, Len: 1, Coverage: 0x4f},
			},
		},
		{
			y: 3,
			spans: []Span{
				{X: 1, Len: 1, Coverage: 0xa7},
				{X: 2, Len: 1, Coverage: 0xff},
				{X: 3, Len: 3, Coverage: 0xfc},
				{X: 6, Len: 1, Coverage: 0xff},
				{X: 7, Len: 1, Coverage: 0xf0},
				{X: 8, Len: 1, Coverage: 0x07},
			},
		},
		{
			y: 4,
			spans: []Span{
				{X: 1, Len: 1, Coverage: 0x4e},
				{X: 2, Len: 1, Coverage: 0xfd},
				{X: 3, Len: 1, Coverage: 0x1a},
				{X: 5, Len: 1, Coverage: 0x09},
				{X: 6, Len: 1, Coverage: 0xf3},
				{X: 7, Len: 1, Coverage: 0x9f},
			},
		},
		{
			y: 5,
			spans: []Span{
				{X: 1, Len: 1, Coverage: 0x05},
				{X: 2, Len: 1, Coverage: 0xee},
				{X: 3, Len: 1, Coverage: 0x6f},
				{X: 5, Len: 1, Coverage: 0x53},
				{X: 6, Len: 1, Coverage: 0xff},
				{X: 7, Len: 1, Coverage: 0x47},
			},
		},
		{
			y: 6,
			spans: []Span{
				{X: 2, Len: 1, Coverage: 0x9b},
				{X: 3, Len: 1, Coverage: 0xc5},
				{X: 5, Len: 1, Coverage: 0xa8},
				{X: 6, Len: 1, Coverage: 0xeb},
				{X: 7, Len: 1, Coverage: 0x04},
			},
		},
		{
			y: 7,
			spans: []Span{
				{X: 2, Len: 1, Coverage: 0x41},
				{X: 3, Len: 1, Coverage: 0xfe},
				{X: 4, Len: 1, Coverage: 0x27},
				{X: 5, Len: 1, Coverage: 0xf4},
				{X: 6, Len: 1, Coverage: 0x97},
			},
		},
		{
			y: 8,
			spans: []Span{
				{X: 2, Len: 1, Coverage: 0x02},
				{X: 3, Len: 1, Coverage: 0xe5},
				{X: 4, Len: 1, Coverage: 0xc4},
				{X: 5, Len: 1, Coverage: 0xff},
				{X: 6, Len: 1, Coverage: 0x3f},
			},
		},
		{
			y: 9,
			spans: []Span{
				{X: 3, Len: 1, Coverage: 0x8e},
				{X: 4, Len: 1, Coverage: 0xff},
				{X: 5, Len: 1, Coverage: 0xe5},
				{X: 6, Len: 1, Coverage: 0x02},
			},
		},
		{
			y: 10,
			spans: []Span{
				{X: 3, Len: 1, Coverage: 0x35},
				{X: 4, Len: 1, Coverage: 0xff},
				{X: 5, Len: 1, Coverage: 0x8f},
			},
		},
	}

	renderer := &stackRenderer{}
	params := RasterParams{
		Flags:     RasterFlagAA | RasterFlagDirect,
		GraySpans: renderer.GraySpans,
	}

	if err := face.GlyphSlot().Outline.Render(l, params); err != nil {
		t.Errorf("Outline.Render() error = %v", err)
	}

	if diff := diff(renderer.frames, want); diff != nil {
		t.Errorf("Outline.Render() %v", diff)
	}
}

type move struct{ to Vector }
type line move
type conic struct{ control, to Vector }
type cubic struct{ control1, control2, to Vector }

type stackDecomposer struct {
	stack []interface{}
}

func (d *stackDecomposer) MoveTo(to Vector) error {
	d.stack = append(d.stack, move{to})
	return nil
}
func (d *stackDecomposer) LineTo(to Vector) error {
	d.stack = append(d.stack, line{to})
	return nil
}
func (d *stackDecomposer) ConicTo(control, to Vector) error {
	d.stack = append(d.stack, conic{control, to})
	return nil
}
func (d *stackDecomposer) CubicTo(control1, control2, to Vector) error {
	d.stack = append(d.stack, cubic{control1, control2, to})
	return nil
}

func TestOutlineDecompose(t *testing.T) {
	getGlyph := func(face testface, r rune) *OutlineGlyph {
		if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
			t.Fatalf("unable to set char size: %v", err)
		}
		if err := face.LoadChar(r, LoadDefault); err != nil {
			t.Fatalf("unable to load char: %v", err)
		}
		glyph, err := face.GlyphSlot().Glyph()
		if err != nil {
			t.Fatalf("unable to get glyph: %v", err)
		}
		ret, ok := glyph.(*OutlineGlyph)
		if !ok {
			t.Fatalf("not an outline")
		}
		return ret
	}

	goReg, err := goRegular()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer goReg.Free()

	srcSans, err := faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoMO.otf")()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer srcSans.Free()

	goA := getGlyph(goReg, 'A')
	defer goA.Free()

	goB := getGlyph(goReg, 'B')
	defer goB.Free()

	srcA := getGlyph(srcSans, 'A')
	defer srcA.Free()

	srcB := getGlyph(srcSans, 'B')
	defer srcB.Free()

	tests := []struct {
		name    string
		outline *Outline
		want    []interface{}
		wantErr error
	}{
		{name: "nilOutline", outline: nil, wantErr: ErrInvalidOutline},
		{
			name:    "goA",
			outline: goA.Outline,
			want: []interface{}{
				move{to: Vector{8, 0}},
				line{to: Vector{254, 704}},
				line{to: Vector{345, 704}},
				line{to: Vector{587, 0}},
				line{to: Vector{488, 0}},
				line{to: Vector{421, 182}},
				line{to: Vector{161, 182}},
				line{to: Vector{94, 0}},
				line{to: Vector{8, 0}},
				move{to: Vector{187, 255}},
				line{to: Vector{396, 255}},
				line{to: Vector{292, 565}},
				line{to: Vector{187, 255}},
			},
			wantErr: nil,
		},
		{
			name:    "goB",
			outline: goB.Outline,
			want: []interface{}{
				move{to: Vector{72, 0}},
				line{to: Vector{72, 704}},
				line{to: Vector{280, 704}},
				conic{control: Vector{407, 704}, to: Vector{460, 666}},
				conic{control: Vector{513, 629}, to: Vector{513, 537}},
				conic{control: Vector{513, 402}, to: Vector{369, 350}},
				conic{control: Vector{539, 302}, to: Vector{539, 168}},
				conic{control: Vector{539, 104}, to: Vector{503, 59}},
				conic{control: Vector{475, 24}, to: Vector{434, 12}},
				conic{control: Vector{394, 0}, to: Vector{302, 0}},
				line{to: Vector{72, 0}},
				move{to: Vector{164, 68}},
				line{to: Vector{232, 68}},
				conic{control: Vector{361, 68}, to: Vector{401, 87}},
				conic{control: Vector{441, 107}, to: Vector{441, 171}},
				conic{control: Vector{441, 239}, to: Vector{389, 275}},
				conic{control: Vector{337, 312}, to: Vector{239, 312}},
				line{to: Vector{164, 312}},
				line{to: Vector{164, 68}},
				move{to: Vector{164, 380}},
				line{to: Vector{242, 380}},
				conic{control: Vector{418, 380}, to: Vector{418, 526}},
				conic{control: Vector{418, 590}, to: Vector{383, 613}},
				conic{control: Vector{348, 636}, to: Vector{249, 636}},
				line{to: Vector{164, 636}},
				line{to: Vector{164, 380}},
			},
			wantErr: nil,
		},
		{
			name:    "srcA",
			outline: srcA.Outline,
			want: []interface{}{
				move{to: Vector{84, 0}},
				line{to: Vector{116, 0}},
				line{to: Vector{116, 482}},
				cubic{control1: Vector{116, 523}, control2: Vector{111, 570}, to: Vector{109, 611}},
				line{to: Vector{112, 611}},
				line{to: Vector{141, 502}},
				line{to: Vector{227, 177}},
				line{to: Vector{256, 177}},
				line{to: Vector{342, 502}},
				line{to: Vector{367, 611}},
				line{to: Vector{371, 611}},
				cubic{control1: Vector{369, 570}, control2: Vector{365, 523}, to: Vector{365, 482}},
				line{to: Vector{365, 0}},
				line{to: Vector{399, 0}},
				line{to: Vector{399, 650}},
				line{to: Vector{353, 650}},
				line{to: Vector{271, 337}},
				line{to: Vector{243, 224}},
				line{to: Vector{240, 224}},
				line{to: Vector{212, 337}},
				line{to: Vector{130, 650}},
				line{to: Vector{84, 650}},
				line{to: Vector{84, 0}},
				move{to: Vector{654, -11}},
				cubic{control1: Vector{761, -11}, control2: Vector{837, 103}, to: Vector{837, 328}},
				cubic{control1: Vector{837, 550}, control2: Vector{761, 661}, to: Vector{654, 661}},
				cubic{control1: Vector{547, 661}, control2: Vector{471, 550}, to: Vector{471, 328}},
				cubic{control1: Vector{471, 103}, control2: Vector{547, -11}, to: Vector{654, -11}},
				move{to: Vector{654, 19}},
				cubic{control1: Vector{565, 19}, control2: Vector{507, 123}, to: Vector{507, 328}},
				cubic{control1: Vector{507, 532}, control2: Vector{565, 631}, to: Vector{654, 631}},
				cubic{control1: Vector{743, 631}, control2: Vector{801, 532}, to: Vector{801, 328}},
				cubic{control1: Vector{801, 123}, control2: Vector{743, 19}, to: Vector{654, 19}},
				move{to: Vector{0, -108}},
				line{to: Vector{896, -108}},
				line{to: Vector{896, 788}},
				line{to: Vector{0, 788}},
				line{to: Vector{0, -108}},
				move{to: Vector{17, -90}},
				line{to: Vector{17, 770}},
				line{to: Vector{878, 770}},
				line{to: Vector{878, -90}},
				line{to: Vector{17, -90}},
			},
			wantErr: nil,
		},
		{
			name:    "srcB",
			outline: srcB.Outline,
			want: []interface{}{
				move{to: Vector{84, 0}},
				line{to: Vector{116, 0}},
				line{to: Vector{116, 482}},
				cubic{control1: Vector{116, 523}, control2: Vector{111, 570}, to: Vector{109, 611}},
				line{to: Vector{112, 611}},
				line{to: Vector{141, 502}},
				line{to: Vector{227, 177}},
				line{to: Vector{256, 177}},
				line{to: Vector{342, 502}},
				line{to: Vector{367, 611}},
				line{to: Vector{371, 611}},
				cubic{control1: Vector{369, 570}, control2: Vector{365, 523}, to: Vector{365, 482}},
				line{to: Vector{365, 0}},
				line{to: Vector{399, 0}},
				line{to: Vector{399, 650}},
				line{to: Vector{353, 650}},
				line{to: Vector{271, 337}},
				line{to: Vector{243, 224}},
				line{to: Vector{240, 224}},
				line{to: Vector{212, 337}},
				line{to: Vector{130, 650}},
				line{to: Vector{84, 650}},
				line{to: Vector{84, 0}},
				move{to: Vector{654, -11}},
				cubic{control1: Vector{761, -11}, control2: Vector{837, 103}, to: Vector{837, 328}},
				cubic{control1: Vector{837, 550}, control2: Vector{761, 661}, to: Vector{654, 661}},
				cubic{control1: Vector{547, 661}, control2: Vector{471, 550}, to: Vector{471, 328}},
				cubic{control1: Vector{471, 103}, control2: Vector{547, -11}, to: Vector{654, -11}},
				move{to: Vector{654, 19}},
				cubic{control1: Vector{565, 19}, control2: Vector{507, 123}, to: Vector{507, 328}},
				cubic{control1: Vector{507, 532}, control2: Vector{565, 631}, to: Vector{654, 631}},
				cubic{control1: Vector{743, 631}, control2: Vector{801, 532}, to: Vector{801, 328}},
				cubic{control1: Vector{801, 123}, control2: Vector{743, 19}, to: Vector{654, 19}},
				move{to: Vector{0, -108}},
				line{to: Vector{896, -108}},
				line{to: Vector{896, 788}},
				line{to: Vector{0, 788}},
				line{to: Vector{0, -108}},
				move{to: Vector{17, -90}},
				line{to: Vector{17, 770}},
				line{to: Vector{878, 770}},
				line{to: Vector{878, -90}},
				line{to: Vector{17, -90}},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &stackDecomposer{}
			if err := tt.outline.Decompose(d, 0, 0); err != tt.wantErr {
				t.Errorf("Outline.Decompose() error = %v, want %v", err, tt.wantErr)
			}
			if diff := diff(d.stack, tt.want); diff != nil {
				t.Errorf("Outline.Decompose() %v", diff)
			}
		})
	}
}

func Test_Outline_Free(t *testing.T) {
	var nilOutline *Outline
	if err := nilOutline.Free(); err != nil {
		t.Errorf("Outline.Free() error = %v", err)
	}

	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to init lib: %v", err)
	}
	defer l.Free()

	face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open face: %v", err)
	}
	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}
	if err := face.LoadChar('A', LoadDefault); err != nil {
		t.Fatalf("unable to load char: %v", err)
	}

	glyph, err := face.GlyphSlot().Glyph()
	if err != nil {
		t.Fatalf("unable to get glyph: %v", err)
	}
	defer glyph.Free()

	wantNewOutline := &Outline{
		userCreated: true,
		Points:      make([]Vector, 10),
		Tags:        make([]byte, 10),
		Contours:    make([]int16, 10),
		Flags:       OutlineOwner,
	}
	wantSlotOutline := &Outline{
		Points: []Vector{
			{8, 0},
			{254, 704},
			{345, 704},
			{587, 0},
			{488, 0},
			{421, 182},
			{161, 182},
			{94, 0},
			{187, 255},
			{396, 255},
			{292, 565},
		},
		Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
		Contours: []int16{7, 10},
		Flags:    0x130,
	}
	wantSlotOutlineCopy := &Outline{
		Points: []Vector{
			{8, 0},
			{254, 704},
			{345, 704},
			{587, 0},
			{488, 0},
			{421, 182},
			{161, 182},
			{94, 0},
			{187, 255},
			{396, 255},
			{292, 565},
		},
		Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
		Contours: []int16{7, 10},
		Flags:    0x131,
	}

	slotOutline := face.GlyphSlot().Outline
	slotOutlineCopy := (glyph.(*OutlineGlyph)).Outline
	newOutline, err := l.NewOutline(10, 10)
	if err != nil {
		t.Fatalf("unable to create outline: %v", err)
	}

	if diff := diff(slotOutline, wantSlotOutline); diff != nil {
		t.Errorf("%v", diff)
	}
	if diff := diff(slotOutlineCopy, wantSlotOutlineCopy); diff != nil {
		t.Errorf("%v", diff)
	}
	if diff := diff(newOutline, wantNewOutline); diff != nil {
		t.Errorf("%v", diff)
	}

	if err := slotOutline.Free(); err != nil {
		t.Errorf("Outline.Free() error = %v", err)
	}
	if err := slotOutlineCopy.Free(); err != nil {
		t.Errorf("Outline.Free() error = %v", err)
	}
	if err := newOutline.Free(); err != nil {
		t.Errorf("Outline.Free() error = %v", err)
	}

	if diff := diff(slotOutline, wantSlotOutline); diff != nil {
		t.Errorf("%v", diff)
	}
	if diff := diff(slotOutlineCopy, wantSlotOutlineCopy); diff != nil {
		t.Errorf("%v", diff)
	}
	if diff := diff(newOutline, &Outline{}); diff != nil {
		t.Errorf("%v", diff)
	}

	glyph.Free()
	if diff := diff(slotOutlineCopy, &Outline{}); diff != nil {
		t.Errorf("%v", diff)
	}
}
