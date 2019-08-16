package freetype2

import (
	"fmt"
	"testing"
)

func TestFace_ForEachColorGlyphLayer(t *testing.T) {
	tests := []struct {
		name       string
		face       func() (testface, error)
		baseGlyph  GlyphIndex
		wantLayers []ColorLayer
	}{
		{
			name:       "nilFace",
			face:       nilFace,
			wantLayers: nil,
		},
		{
			name:       "goRegular",
			face:       goRegular,
			baseGlyph:  0x24, //A
			wantLayers: nil,
		},
		{
			name:       "chromacheckColr",
			face:       chromacheckColr,
			baseGlyph:  0x1,
			wantLayers: []ColorLayer{{0x1, 0}},
		},
		{
			name:      "bungeeColorWin",
			face:      bungeeColorWin,
			baseGlyph: 0x2b, //A
			wantLayers: []ColorLayer{
				{0x124, 0},
				{0x125, 1},
			},
		},
		{
			name:       "bungeeColorMac",
			face:       bungeeColorMac,
			baseGlyph:  0x2b, //A
			wantLayers: nil,
		},
		{
			name:      "twemojiMozilla",
			face:      twemojiMozilla,
			baseGlyph: 0x3ae, //smiling cat with heart-eyes
			wantLayers: []ColorLayer{
				{0x1d4b, 194},
				{0x1d4c, 723},
				{0x1d4d, 665},
				{0x1d4e, 194},
				{0x1d4f, 724},
				{0x1d54, 58},
				{0x1d51, 476},
				{0x1d55, 2},
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

			var invocations int
			var got []ColorLayer
			cb := func(layerGlyphIndex, layerColorIndex int) (done bool) {
				invocations++
				got = append(got, ColorLayer{
					GlyphIndex: layerGlyphIndex,
					ColorIndex: layerColorIndex,
				})
				return false
			}

			if got := face.ForEachColorGlyphLayer(tt.baseGlyph, cb); got != len(tt.wantLayers) {
				t.Errorf("Face.ForEachColorGlyphLayer() = %v, want %v", got, len(tt.wantLayers))
			}
			if invocations != len(tt.wantLayers) {
				t.Errorf("callback should have been invoked %d times, was %d", len(tt.wantLayers), invocations)
			}
			if diff := diff(got, tt.wantLayers); diff != nil {
				t.Errorf("values are not equal: %v", diff)
			}
		})
	}

	t.Run("nill callback", func(t *testing.T) {
		face, err := twemojiMozilla()
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}
		defer face.Free()

		want := 8
		if got := face.ForEachColorGlyphLayer(0x3ae, nil); got != want {
			t.Errorf("Face.ForEachColorGlyphLayer() = %v, want %v", got, want)
		}
	})
}

func TestFace_ForEachColorGlyphLayer_cancellation(t *testing.T) {
	face, err := twemojiMozilla()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()

	baseGlyph := GlyphIndex(0x3ae) //smiling cat with heart-eyes
	layers := []ColorLayer{
		{0x1d4b, 194},
		{0x1d4c, 723},
		{0x1d4d, 665},
		{0x1d4e, 194},
		{0x1d4f, 724},
		{0x1d54, 58},
		{0x1d51, 476},
		{0x1d55, 2},
	}

	tests := []struct {
		stopAt, wantInvocations int
		wantLayers              []ColorLayer
	}{
		// a simple for would be enough, but this is clearer
		{stopAt: 1, wantInvocations: 1, wantLayers: layers[:1]},
		{stopAt: 2, wantInvocations: 2, wantLayers: layers[:2]},
		{stopAt: 3, wantInvocations: 3, wantLayers: layers[:3]},
		{stopAt: 4, wantInvocations: 4, wantLayers: layers[:4]},
		{stopAt: 5, wantInvocations: 5, wantLayers: layers[:5]},
		{stopAt: 6, wantInvocations: 6, wantLayers: layers[:6]},
		{stopAt: 7, wantInvocations: 7, wantLayers: layers[:7]},
		{stopAt: 8, wantInvocations: 8, wantLayers: layers[:8]},
		{stopAt: 9, wantInvocations: 8, wantLayers: layers},
		{stopAt: 10, wantInvocations: 8, wantLayers: layers},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.stopAt), func(t *testing.T) {
			var invocations int
			var got []ColorLayer
			cb := func(layerGlyphIndex, layerColorIndex int) (done bool) {
				invocations++
				got = append(got, ColorLayer{
					GlyphIndex: layerGlyphIndex,
					ColorIndex: layerColorIndex,
				})
				return invocations == tt.stopAt
			}

			if got := face.ForEachColorGlyphLayer(baseGlyph, cb); got != len(layers) {
				t.Errorf("Face.ForEachColorGlyphLayer() = %v, want %v", got, len(layers))
			}
			if invocations != tt.wantInvocations {
				t.Errorf("callback should have been invoked %d times, was %d", tt.wantInvocations, invocations)
			}
			if diff := diff(got, tt.wantLayers); diff != nil {
				t.Errorf("values are not equal: %v", diff)
			}
		})
	}
}

func TestFace_GetColorGlyphLayers(t *testing.T) {
	tests := []struct {
		name       string
		face       func() (testface, error)
		baseGlyph  GlyphIndex
		wantLayers []ColorLayer
	}{
		{
			name:       "nilFace",
			face:       nilFace,
			wantLayers: nil,
		},
		{
			name:       "goRegular",
			face:       goRegular,
			baseGlyph:  0x24, //A
			wantLayers: nil,
		},
		{
			name:       "chromacheckColr",
			face:       chromacheckColr,
			baseGlyph:  0x1,
			wantLayers: []ColorLayer{{0x1, 0}},
		},
		{
			name:      "bungeeColorWin",
			face:      bungeeColorWin,
			baseGlyph: 0x2b, //A
			wantLayers: []ColorLayer{
				{0x124, 0},
				{0x125, 1},
			},
		},
		{
			name:       "bungeeColorMac",
			face:       bungeeColorMac,
			baseGlyph:  0x2b, //A
			wantLayers: nil,
		},
		{
			name:      "twemojiMozilla",
			face:      twemojiMozilla,
			baseGlyph: 0x3ae, //smiling cat with heart-eyes
			wantLayers: []ColorLayer{
				{0x1d4b, 194},
				{0x1d4c, 723},
				{0x1d4d, 665},
				{0x1d4e, 194},
				{0x1d4f, 724},
				{0x1d54, 58},
				{0x1d51, 476},
				{0x1d55, 2},
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

			got := face.GetColorGlyphLayers(tt.baseGlyph)
			if diff := diff(got, tt.wantLayers); diff != nil {
				t.Errorf("Face.GetColorGlyphLayer() %v", diff)
			}
		})
	}
}

func benchmarkFaceForEachColorGlyphLayer(b *testing.B, faceFn func() (testface, error), idx GlyphIndex) {
	b.StopTimer()

	cb := func(_, _ int) bool { return false }
	face, err := faceFn()
	if err != nil {
		b.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()
	//warmup
	face.ForEachColorGlyphLayer(idx, cb)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		face.ForEachColorGlyphLayer(idx, cb)
	}
}
func BenchmarkFace_ForEachColorGlyphLayer_8(b *testing.B) {
	benchmarkFaceForEachColorGlyphLayer(b, twemojiMozilla, 0x3ae)
}

func BenchmarkFace_ForEachColorGlyphLayer_2(b *testing.B) {
	benchmarkFaceForEachColorGlyphLayer(b, bungeeColorWin, 0x2b)
}

func BenchmarkFace_ForEachColorGlyphLayer_1(b *testing.B) {
	benchmarkFaceForEachColorGlyphLayer(b, chromacheckColr, 0x1)
}

func benchmarkFaceGetColorGlyphLayers(b *testing.B, faceFn func() (testface, error), idx GlyphIndex) {
	b.StopTimer()

	face, err := faceFn()
	if err != nil {
		b.Fatalf("unable to load face: %v", err)
	}
	defer face.Free()
	//warmup
	face.GetColorGlyphLayers(idx)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		face.GetColorGlyphLayers(idx)
	}
}

func BenchmarkFace_GetColorGlyphLayers_8(b *testing.B) {
	benchmarkFaceGetColorGlyphLayers(b, twemojiMozilla, 0x3ae)
}

func BenchmarkFace_GetColorGlyphLayers_2(b *testing.B) {
	benchmarkFaceGetColorGlyphLayers(b, bungeeColorWin, 0x2b)
}

func BenchmarkFace_GetColorGlyphLayers_1(b *testing.B) {
	benchmarkFaceGetColorGlyphLayers(b, chromacheckColr, 0x1)
}

// BenchmarkFace_ForEachColorGlyphLayer_8-4   	 1000000	      1139 ns/op	     168 B/op	      12 allocs/op
// BenchmarkFace_ForEachColorGlyphLayer_2-4   	 3000000	       423 ns/op	      72 B/op	       6 allocs/op
// BenchmarkFace_ForEachColorGlyphLayer_1-4   	 5000000	       289 ns/op	      56 B/op	       5 allocs/op
// BenchmarkFace_GetColorGlyphLayers_8-4      	 5000000	       313 ns/op	     156 B/op	       4 allocs/op
// BenchmarkFace_GetColorGlyphLayers_2-4      	 5000000	       238 ns/op	      60 B/op	       4 allocs/op
// BenchmarkFace_GetColorGlyphLayers_1-4      	10000000	       220 ns/op	      44 B/op	       4 allocs/op
