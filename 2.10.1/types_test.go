package freetype2

import (
	"reflect"
	"testing"

	"github.com/flga/freetype2/2.10.1/truetype"
)

func Test_newBitmap(t *testing.T) {
	got := newBitmap(emptyBitmap)
	want := Bitmap{}
	if diff := diff(got, want); diff != nil {
		t.Errorf("newBitmap(empty) %v", diff)
	}

	face, err := goRegular()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	if err := face.LoadGlyph(0x24, LoadRender|LoadColor); err != nil {
		t.Fatalf("unable laod glyph: %v", err)
	}

	got = newBitmap(face.ptr.glyph.bitmap)
	want = Bitmap{
		Rows:      0xb,
		Width:     0xa,
		Pitch:     10,
		Buffer:    goRegularGlyphBuf(0x24, 0),
		NumGrays:  0x100,
		PixelMode: PixelModeGray,
	}
	if diff := diff(got, want); diff != nil {
		t.Errorf("newBitmap() %v", diff)
	}
}

func Test_newCharMap(t *testing.T) {
	if got, want := newCharMap(nil), (CharMap{}); got != want {
		t.Errorf("newCharMap(nil) = %v, want %v", got, want)
	}

	face, err := goRegular()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	got := newCharMap(face.charmaps()[1])
	want := CharMap{
		Format:     6,
		Language:   truetype.MacLangEnglish,
		Encoding:   EncodingAppleRoman,
		PlatformID: truetype.PlatformMacintosh,
		EncodingID: truetype.MacEncodingRoman,
		index:      1,
		valid:      true,
	}
	if diff := diff(got, want); diff != nil {
		t.Errorf("newCharMap() %v", diff)
	}
}

func TestCharMap_Index(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		want, wantOk := 0, false
		got, gotOk := (CharMap{}).Index()
		if got != want || gotOk != wantOk {
			t.Errorf("CharMap.Index() = %v, %v, want %v, %v", got, gotOk, want, wantOk)
		}
	})

	t.Run("notoSansJpReg", func(t *testing.T) {
		face, err := notoSansJpReg()
		if err != nil {
			t.Fatalf("unable to open face: %v", err)
		}

		for i, c := range face.CharMaps() {
			want, wantOk := i, true
			got, gotOk := c.Index()
			if got != want || gotOk != wantOk {
				t.Errorf("CharMap.Index() = %v, %v, want %v, %v", got, gotOk, want, wantOk)
			}
		}
	})
}

func Test_newSize(t *testing.T) {
	if got, want := newSize(nil), (Size{}); got != want {
		t.Errorf("newSize(nil) = %v, want %v", got, want)
	}
	face, err := goRegular()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(20<<6, 20<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	got := newSize(face.ptr.size)
	want := Size{
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
	}
	if diff := diff(got, want); diff != nil {
		t.Errorf("newSize() %v", diff)
	}
}

func Test_newGlyphSlot(t *testing.T) {
	got := newGlyphSlot(nil)
	want := GlyphSlot{}
	if diff := diff(got, want); diff != nil {
		t.Errorf("newGlyphSlot(nil) = %v", diff)
	}

	face, err := notoSansJpReg() // noto has both horizontal and vertical modes
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

	got = newGlyphSlot(face.ptr.glyph)
	want = GlyphSlot{
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
		Advance:           Vector26_6{X: 576, Y: 0},
		Format:            GlyphFormatBitmap,
		Bitmap:            Bitmap{}, // separate test
		BitmapLeft:        0,
		BitmapTop:         11,
		Outline:           Outline{}, // separate test
		NumSubglyphs:      0,
		LsbDelta:          0,
		RsbDelta:          0,
	}

	// don't care about the actual values, separate tests, just make sure
	// they're not 0 vals
	if reflect.DeepEqual(got.Bitmap, Bitmap{}) {
		t.Errorf("newGlyphSlot() want non-zero %T", Bitmap{})
	}
	if reflect.DeepEqual(got.Outline, Outline{}) {
		t.Errorf("newGlyphSlot() want non-zero %T", Outline{})
	}
	// set to 0 for comparison
	got.Bitmap = Bitmap{}
	got.Outline = Outline{}

	if diff := diff(got, want); diff != nil {
		t.Errorf("newGlyphSlot() %v", diff)
	}
}

func Test_newOutline(t *testing.T) {
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

	got := newOutline(face.ptr.glyph.outline)
	want := Outline{
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
