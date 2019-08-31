package freetype2

import (
	"reflect"
	"testing"

	"github.com/flga/freetype2/2.10.1/truetype"
)

func Test_newBitmap(t *testing.T) {
	got := &Bitmap{ptr: nil}
	want := &Bitmap{}
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

	got = &Bitmap{ptr: &face.ptr.glyph.bitmap}
	got.reload()
	want = &Bitmap{
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
	if got, want := newSize(nil), (*Size)(nil); got != want {
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
	want := &Size{
		SizeMetrics: SizeMetrics{
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
	// got := newGlyphSlot(nil)
	// var want *GlyphSlot
	// if diff := diff(got, want); diff != nil {
	// 	t.Errorf("newGlyphSlot(nil) = %v", diff)
	// }

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

	got := &GlyphSlot{ptr: face.ptr.glyph}
	got.reload()
	want := &GlyphSlot{
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
		Bitmap:            nil, // separate test
		BitmapLeft:        0,
		BitmapTop:         11,
		Outline:           nil, // separate test
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
	got.Bitmap = nil
	got.Outline = nil

	if diff := diff(got, want); diff != nil {
		t.Errorf("newGlyphSlot() %v", diff)
	}
}

func TestGlyphSlot_SubGlyphInfo(t *testing.T) {
	var nilSlot *GlyphSlot
	if got, err := nilSlot.SubGlyphInfo(0); got != (SubGlyphInfo{}) || err != ErrInvalidArgument {
		t.Errorf("GlyphSlot.SubGlyphInfo() got %v, %v, want %v, %v", got, err, SubGlyphInfo{}, ErrInvalidArgument)
	}

	face, err := bungeeColorWin()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable set char size: %v", err)
	}

	if err := face.LoadChar(0x3a, LoadNoRecurse); err != nil {
		t.Fatalf("unable load char: %v", err)
	}

	want0 := SubGlyphInfo{
		Index:     0x18,
		Flags:     0x226,
		Arg1:      0,
		Arg2:      0,
		Transform: Matrix{Xx: 65536, Xy: 0, Yx: 0, Yy: 65536},
	}

	want1 := SubGlyphInfo{
		Index:     0x18,
		Flags:     0x107,
		Arg1:      0,
		Arg2:      355,
		Transform: Matrix{Xx: 65536, Xy: 0, Yx: 0, Yy: 65536},
	}

	got0, err := face.GlyphSlot().SubGlyphInfo(0)
	if diff := diff(got0, want0); diff != nil || err != nil {
		t.Errorf("GlyphSlot.SubGlyphInfo(0) = %v, %v", diff, err)
	}
	got1, err := face.GlyphSlot().SubGlyphInfo(1)
	if diff := diff(got1, want1); diff != nil || err != nil {
		t.Errorf("GlyphSlot.SubGlyphInfo(1) = %v, %v", diff, err)
	}
	if got, err := face.GlyphSlot().SubGlyphInfo(2); got != (SubGlyphInfo{}) || err != ErrInvalidArgument {
		t.Errorf("GlyphSlot.SubGlyphInfo(2) got %v, %v, want %v, %v", got, err, SubGlyphInfo{}, ErrInvalidArgument)
	}

	// not a composite
	if err := face.LoadChar(0x3a, LoadDefault); err != nil {
		t.Fatalf("unable load char: %v", err)
	}
	if got, err := face.GlyphSlot().SubGlyphInfo(0); got != (SubGlyphInfo{}) || err != ErrInvalidArgument {
		t.Errorf("GlyphSlot.SubGlyphInfo() got %v, %v, want %v, %v", got, err, SubGlyphInfo{}, ErrInvalidArgument)
	}
}

func TestGlyphSlot_RenderGlyph(t *testing.T) {
	var slot *GlyphSlot
	if err := slot.RenderGlyph(0); err != ErrInvalidArgument {
		t.Errorf("GlyphSlot.RenderGlyph() error = %v, want %v", err, ErrInvalidArgument)
	}

	face, err := goRegular()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable set char size: %v", err)
	}

	if err := face.LoadChar('A', LoadDefault); err != nil {
		t.Fatalf("unable load char: %v", err)
	}

	slot = face.GlyphSlot()
	want := &GlyphSlot{
		GlyphIndex: 0x24,
		Metrics: GlyphMetrics{
			Width:        640,
			Height:       704,
			HoriBearingX: 0,
			HoriBearingY: 704,
			HoriAdvance:  576,
			VertBearingX: -320,
			VertBearingY: 64,
			VertAdvance:  896,
		},
		LinearHoriAdvance: 611968,
		LinearVertAdvance: 884352,
		Advance:           Vector26_6{X: 576, Y: 0},
		Format:            GlyphFormatOutline,
		Bitmap: &Bitmap{
			Rows:      11,
			Width:     10,
			Pitch:     10,
			Buffer:    nil,
			NumGrays:  256,
			PixelMode: PixelModeGray,
		},
		BitmapLeft: 0,
		BitmapTop:  11,
		Outline: &Outline{
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
		},
		NumSubglyphs: 0,
		LsbDelta:     0,
		RsbDelta:     0,
	}

	if diff := diff(slot, want); diff != nil {
		t.Error(diff)
	}

	want = &GlyphSlot{
		GlyphIndex: 0x24,
		Metrics: GlyphMetrics{
			Width:        640,
			Height:       704,
			HoriBearingX: 0,
			HoriBearingY: 704,
			HoriAdvance:  576,
			VertBearingX: -320,
			VertBearingY: 64,
			VertAdvance:  896,
		},
		LinearHoriAdvance: 611968,
		LinearVertAdvance: 884352,
		Advance:           Vector26_6{X: 576, Y: 0},
		Format:            GlyphFormatBitmap,
		Bitmap: &Bitmap{
			Rows:      11,
			Width:     10,
			Pitch:     10,
			Buffer:    testGlyphSlotRenderGlyphdata(),
			NumGrays:  256,
			PixelMode: PixelModeGray,
		},
		BitmapLeft: 0,
		BitmapTop:  11,
		Outline: &Outline{
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
		},
		NumSubglyphs: 0,
		LsbDelta:     0,
		RsbDelta:     0,
	}
	if err := slot.RenderGlyph(RenderModeNormal); err != nil {
		t.Errorf("GlyphSlot.RenderGlyph() error = %v", err)
	}
	if diff := diff(slot, want); diff != nil {
		t.Error(diff)
	}

}

func TestSize_Free(t *testing.T) {
	var nilSize *Size
	if err := nilSize.Free(); err != nil {
		t.Errorf("Size.Free() error = %v", err)
	}

	face, err := goRegular()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}

	s, err := face.NewSize()
	if err != nil {
		t.Fatalf("unable to create size: %v", err)
	}
	if err := s.Free(); err != nil {
		t.Fatalf("unable to free: %v", err)
	}
	if s.ptr != nil {
		t.Fatalf("ptr should be nil")
	}

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	s = face.Size()
	if err := face.Free(); err != nil {
		t.Fatalf("unable to free: %v", err)
	}
	if s.ptr != nil {
		t.Fatalf("ptr should be nil")
	}
}
