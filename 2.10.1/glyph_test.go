package freetype2

import (
	"fmt"
	"math"
	"testing"

	"github.com/flga/freetype2/fixed"
)

func TestLibrary_NewGlyph(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %v", err)
	}
	defer l.Free()
	tests := []struct {
		name    string
		l       *Library
		format  GlyphFormat
		want    Glyph
		wantErr error
	}{
		{
			name:    "nilLib",
			l:       nil,
			wantErr: ErrInvalidArgument,
		},
		{
			name:    "emptyLib",
			l:       &Library{},
			wantErr: ErrInvalidArgument,
		},
		{
			name:    "Composite",
			l:       l,
			format:  GlyphFormatComposite,
			wantErr: ErrInvalidGlyphFormat,
		},
		{
			name:   "Bitmap",
			l:      l,
			format: GlyphFormatBitmap,
			want: &BitmapGlyph{
				format:  GlyphFormatBitmap,
				advance: Vector16_16{},
				Left:    0,
				Top:     0,
				Bitmap:  Bitmap{},
			},
		},
		{
			name:   "Outline",
			l:      l,
			format: GlyphFormatOutline,
			want: &OutlineGlyph{
				format:  GlyphFormatOutline,
				advance: Vector16_16{},
				Outline: Outline{},
			},
		},
		{
			name:    "Plotter",
			l:       l,
			format:  GlyphFormatPlotter,
			wantErr: ErrInvalidGlyphFormat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.NewGlyph(tt.format)
			if err != tt.wantErr {
				t.Fatalf("Library.NewGlyph() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != nil {
				defer got.Free()
			}

			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Library.NewGlyph() = %v", diff)
			}
		})
	}
}

func TestGlyph_Free(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %v", err)
	}
	defer l.Free()

	g, err := l.NewGlyph(GlyphFormatOutline)
	if err != nil {
		t.Fatalf("unable to create glyph: %v", err)
	}

	g.Free()
	if g.getptr() != nil {
		t.Fatalf("g.ptr is not nil")
	}
}

func TestGlyphSlot_Glyph(t *testing.T) {
	slot := func(facefn func() (testface, error), char rune, flags LoadFlag, size fixed.Int26_6, dpi uint) func() (*GlyphSlot, error) {
		return func() (*GlyphSlot, error) {
			face, err := facefn()
			if err != nil {
				return nil, err
			}

			if face.Flags()&FaceFlagFixedSizes > 0 {
				if err := face.SelectSize(0); err != nil {
					face.Free()
					return nil, err
				}
			} else {
				if err := face.SetCharSize(size, size, dpi, dpi); err != nil {
					face.Free()
					return nil, err
				}
			}

			if err := face.LoadChar(char, flags); err != nil {
				face.Free()
				return nil, err
			}

			return face.GlyphSlot(), nil
		}
	}

	nilSlot := func() (*GlyphSlot, error) {
		return nil, nil
	}
	emptySlot := func() (*GlyphSlot, error) {
		return &GlyphSlot{}, nil
	}

	tests := []struct {
		name    string
		slot    func() (*GlyphSlot, error)
		want    Glyph
		wantErr error
	}{
		{
			name:    "nilSlot",
			slot:    nilSlot,
			want:    nil,
			wantErr: ErrInvalidSlotHandle,
		},
		{
			name:    "emptySlot",
			slot:    emptySlot,
			want:    nil,
			wantErr: ErrInvalidSlotHandle,
		},
		{
			name:    "composite",
			slot:    slot(bungeeColorWin, 0x3a, LoadNoRecurse, 14<<6, 72),
			wantErr: ErrInvalidGlyphFormat,
		},
		{
			name: "outline",
			slot: slot(goRegular, 'A', LoadDefault, 14<<6, 72),
			want: &OutlineGlyph{
				format:  GlyphFormatOutline,
				advance: Vector16_16{X: 589824, Y: 0},
				Outline: Outline{
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
				},
			},
			wantErr: nil,
		},
		{
			name: "bitmap",
			slot: slot(bungeeColorMac, 'A', LoadDefault, 14<<6, 72),
			want: &BitmapGlyph{
				format:  GlyphFormatBitmap,
				advance: Vector16_16{X: 917504, Y: 0},
				Left:    0,
				Top:     15,
				Bitmap: Bitmap{
					Rows:  15,
					Width: 13,
					Pitch: 13,
					Buffer: []byte{
						0x00, 0x00, 0x15, 0xbf, 0xc7, 0xc7, 0xc7, 0xc7, 0xc7, 0xbf, 0x15, 0x00, 0x00,
						0x00, 0x00, 0x80, 0xdf, 0xdd, 0xda, 0xdb, 0xda, 0xde, 0xdf, 0x80, 0x00, 0x00,
						0x00, 0x05, 0xd1, 0xde, 0xd4, 0xd9, 0xdb, 0xd9, 0xd4, 0xde, 0xd1, 0x05, 0x00,
						0x00, 0x3e, 0xdf, 0xd8, 0xda, 0xde, 0xd4, 0xdf, 0xd8, 0xd9, 0xdf, 0x3e, 0x00,
						0x00, 0xa6, 0xde, 0xd5, 0xde, 0x98, 0x07, 0xb8, 0xde, 0xd5, 0xdf, 0xa6, 0x00,
						0x13, 0xde, 0xdb, 0xd8, 0xdf, 0x3c, 0x00, 0x64, 0xdf, 0xd7, 0xdc, 0xde, 0x11,
						0x5e, 0xdf, 0xd7, 0xdc, 0xd4, 0x03, 0x00, 0x11, 0xdd, 0xdc, 0xd7, 0xdf, 0x5e,
						0xa8, 0xde, 0xd6, 0xde, 0xc7, 0x45, 0x45, 0x50, 0xd5, 0xde, 0xd6, 0xde, 0xa8,
						0xc6, 0xdd, 0xd6, 0xde, 0xde, 0xde, 0xdf, 0xde, 0xdd, 0xde, 0xd6, 0xdd, 0xc6,
						0xc8, 0xde, 0xd0, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd0, 0xde, 0xc8,
						0xc7, 0xdd, 0xd6, 0xde, 0xde, 0xdf, 0xdf, 0xdf, 0xde, 0xde, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x8f, 0x26, 0x26, 0x26, 0xaa, 0xdf, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x60, 0x00, 0x00, 0x00, 0x8a, 0xdf, 0xd6, 0xdd, 0xc7,
						0xce, 0xde, 0xde, 0xdf, 0x6c, 0x00, 0x00, 0x00, 0x96, 0xdf, 0xde, 0xde, 0xce,
						0xa1, 0xc8, 0xc7, 0xc9, 0x48, 0x00, 0x00, 0x00, 0x69, 0xc9, 0xc7, 0xc8, 0xa1,
					},
					NumGrays:  256,
					PixelMode: PixelModeGray,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot, err := tt.slot()
			if err != nil {
				t.Fatalf("unable to load slot: %v", err)
			}

			got, gotErr := slot.Glyph()
			if gotErr != tt.wantErr {
				t.Errorf("GlyphSlot.Glyph() error = %v, want %v", gotErr, tt.wantErr)
			}

			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("GlyphSlot.Glyph() = %v", diff)
			}
		})
	}
}

func TestGlyph_Copy(t *testing.T) {
	tests := []struct {
		name    string
		glyph   func() (testglyph, error)
		want    Glyph
		wantErr error
	}{
		{
			name:    "nilGlyph",
			glyph:   newNilGlyph,
			want:    nil,
			wantErr: ErrInvalidArgument,
		},
		{
			name:    "emptyGlyph",
			glyph:   newZeroGlyph(),
			want:    nil,
			wantErr: ErrInvalidArgument,
		},
		{
			name:  "empty Bitmap",
			glyph: newEmptyGlyph(GlyphFormatBitmap),
			want: &BitmapGlyph{
				format:  GlyphFormatBitmap,
				advance: Vector16_16{X: 0, Y: 0},
				Left:    0,
				Top:     0,
				Bitmap:  Bitmap{},
			},
			wantErr: nil,
		},
		{
			name:  "real Bitmap",
			glyph: newTestGlyph(bungeeColorMac, 'A', LoadDefault, 14<<6, 72),
			want: &BitmapGlyph{
				format:  GlyphFormatBitmap,
				advance: Vector16_16{X: 917504, Y: 0},
				Left:    0,
				Top:     15,
				Bitmap: Bitmap{
					Rows:  15,
					Width: 13,
					Pitch: 13,
					Buffer: []byte{
						0x00, 0x00, 0x15, 0xbf, 0xc7, 0xc7, 0xc7, 0xc7, 0xc7, 0xbf, 0x15, 0x00, 0x00,
						0x00, 0x00, 0x80, 0xdf, 0xdd, 0xda, 0xdb, 0xda, 0xde, 0xdf, 0x80, 0x00, 0x00,
						0x00, 0x05, 0xd1, 0xde, 0xd4, 0xd9, 0xdb, 0xd9, 0xd4, 0xde, 0xd1, 0x05, 0x00,
						0x00, 0x3e, 0xdf, 0xd8, 0xda, 0xde, 0xd4, 0xdf, 0xd8, 0xd9, 0xdf, 0x3e, 0x00,
						0x00, 0xa6, 0xde, 0xd5, 0xde, 0x98, 0x07, 0xb8, 0xde, 0xd5, 0xdf, 0xa6, 0x00,
						0x13, 0xde, 0xdb, 0xd8, 0xdf, 0x3c, 0x00, 0x64, 0xdf, 0xd7, 0xdc, 0xde, 0x11,
						0x5e, 0xdf, 0xd7, 0xdc, 0xd4, 0x03, 0x00, 0x11, 0xdd, 0xdc, 0xd7, 0xdf, 0x5e,
						0xa8, 0xde, 0xd6, 0xde, 0xc7, 0x45, 0x45, 0x50, 0xd5, 0xde, 0xd6, 0xde, 0xa8,
						0xc6, 0xdd, 0xd6, 0xde, 0xde, 0xde, 0xdf, 0xde, 0xdd, 0xde, 0xd6, 0xdd, 0xc6,
						0xc8, 0xde, 0xd0, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd0, 0xde, 0xc8,
						0xc7, 0xdd, 0xd6, 0xde, 0xde, 0xdf, 0xdf, 0xdf, 0xde, 0xde, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x8f, 0x26, 0x26, 0x26, 0xaa, 0xdf, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x60, 0x00, 0x00, 0x00, 0x8a, 0xdf, 0xd6, 0xdd, 0xc7,
						0xce, 0xde, 0xde, 0xdf, 0x6c, 0x00, 0x00, 0x00, 0x96, 0xdf, 0xde, 0xde, 0xce,
						0xa1, 0xc8, 0xc7, 0xc9, 0x48, 0x00, 0x00, 0x00, 0x69, 0xc9, 0xc7, 0xc8, 0xa1,
					},
					NumGrays:  256,
					PixelMode: PixelModeGray,
				},
			},
			wantErr: nil,
		},
		{
			name:  "empty Outline",
			glyph: newEmptyGlyph(GlyphFormatOutline),
			want: &OutlineGlyph{
				format:  GlyphFormatOutline,
				advance: Vector16_16{X: 0, Y: 0},
				Outline: Outline{
					Flags: OutlineOwner,
				},
			},
			wantErr: nil,
		},
		{
			name:  "real Outline",
			glyph: newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72),
			want: &OutlineGlyph{
				format:  GlyphFormatOutline,
				advance: Vector16_16{X: 589824, Y: 0},
				Outline: Outline{
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
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			glyph, err := tt.glyph()
			if err != nil {
				t.Fatalf("unable to get glyph: %v", err)
			}
			defer glyph.Free()
			var got Glyph
			if glyph.Glyph != nil {
				var err error
				got, err = glyph.Copy()
				if err != tt.wantErr {
					t.Errorf("Glyph.Copy() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Glyph.Copy() = %v", diff)
			}
		})
	}
}

func TestGlyph_Transform(t *testing.T) {
	deg := float64(90)
	angle := deg / 360.0 * math.Pi * 2.0
	matrix := Matrix{
		Xx: (fixed.Int16_16)(math.Cos(angle) * 0x10000),
		Xy: (fixed.Int16_16)(-math.Sin(angle) * 0x10000),
		Yx: (fixed.Int16_16)(math.Sin(angle) * 0x10000),
		Yy: (fixed.Int16_16)(math.Cos(angle) * 0x10000),
	}
	delta := Vector26_6{X: 64, Y: 64}

	tests := []struct {
		name    string
		glyph   func() (testglyph, error)
		matrix  Matrix
		delta   Vector26_6
		want    Glyph
		wantErr error
	}{
		{
			name:    "nilGlyph",
			glyph:   newNilGlyph,
			want:    nil,
			wantErr: ErrInvalidArgument,
		},
		{
			name:    "emptyGlyph",
			glyph:   newZeroGlyph(GlyphFormatBitmap),
			want:    &BitmapGlyph{},
			wantErr: ErrInvalidArgument,
		},
		{
			name:   "bitmap",
			glyph:  newTestGlyph(bungeeColorMac, 'A', LoadDefault, 14<<6, 72),
			matrix: matrix,
			delta:  Vector26_6{},
			want: &BitmapGlyph{
				format:  GlyphFormatBitmap,
				advance: Vector16_16{X: 917504, Y: 0},
				Left:    0,
				Top:     15,
				Bitmap: Bitmap{
					Rows:  15,
					Width: 13,
					Pitch: 13,
					Buffer: []byte{
						0x00, 0x00, 0x15, 0xbf, 0xc7, 0xc7, 0xc7, 0xc7, 0xc7, 0xbf, 0x15, 0x00, 0x00,
						0x00, 0x00, 0x80, 0xdf, 0xdd, 0xda, 0xdb, 0xda, 0xde, 0xdf, 0x80, 0x00, 0x00,
						0x00, 0x05, 0xd1, 0xde, 0xd4, 0xd9, 0xdb, 0xd9, 0xd4, 0xde, 0xd1, 0x05, 0x00,
						0x00, 0x3e, 0xdf, 0xd8, 0xda, 0xde, 0xd4, 0xdf, 0xd8, 0xd9, 0xdf, 0x3e, 0x00,
						0x00, 0xa6, 0xde, 0xd5, 0xde, 0x98, 0x07, 0xb8, 0xde, 0xd5, 0xdf, 0xa6, 0x00,
						0x13, 0xde, 0xdb, 0xd8, 0xdf, 0x3c, 0x00, 0x64, 0xdf, 0xd7, 0xdc, 0xde, 0x11,
						0x5e, 0xdf, 0xd7, 0xdc, 0xd4, 0x03, 0x00, 0x11, 0xdd, 0xdc, 0xd7, 0xdf, 0x5e,
						0xa8, 0xde, 0xd6, 0xde, 0xc7, 0x45, 0x45, 0x50, 0xd5, 0xde, 0xd6, 0xde, 0xa8,
						0xc6, 0xdd, 0xd6, 0xde, 0xde, 0xde, 0xdf, 0xde, 0xdd, 0xde, 0xd6, 0xdd, 0xc6,
						0xc8, 0xde, 0xd0, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd0, 0xde, 0xc8,
						0xc7, 0xdd, 0xd6, 0xde, 0xde, 0xdf, 0xdf, 0xdf, 0xde, 0xde, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x8f, 0x26, 0x26, 0x26, 0xaa, 0xdf, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x60, 0x00, 0x00, 0x00, 0x8a, 0xdf, 0xd6, 0xdd, 0xc7,
						0xce, 0xde, 0xde, 0xdf, 0x6c, 0x00, 0x00, 0x00, 0x96, 0xdf, 0xde, 0xde, 0xce,
						0xa1, 0xc8, 0xc7, 0xc9, 0x48, 0x00, 0x00, 0x00, 0x69, 0xc9, 0xc7, 0xc8, 0xa1,
					},
					NumGrays:  256,
					PixelMode: 2,
				},
			},
			wantErr: ErrInvalidGlyphFormat,
		},
		{
			name:   "outline none",
			glyph:  newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72),
			matrix: Matrix{},
			delta:  Vector26_6{},
			want: &OutlineGlyph{
				format:  GlyphFormatOutline,
				advance: Vector16_16{X: 589824, Y: 0},
				Outline: Outline{
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
				},
			},
			wantErr: nil,
		},
		{
			name:   "outline matrix",
			glyph:  newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72),
			matrix: matrix,
			delta:  Vector26_6{},
			want: &OutlineGlyph{
				format:  GlyphFormatOutline,
				advance: Vector16_16{X: 0, Y: 589824},
				Outline: Outline{
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
				},
			},
			wantErr: nil,
		},
		{
			name:   "outline delta",
			glyph:  newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72),
			matrix: Matrix{},
			delta:  delta,
			want: &OutlineGlyph{
				format:  GlyphFormatOutline,
				advance: Vector16_16{X: 589824, Y: 0},
				Outline: Outline{
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
				},
			},
			wantErr: nil,
		},
		{
			name:   "outline matrix & delta",
			glyph:  newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72),
			matrix: matrix,
			delta:  delta,
			want: &OutlineGlyph{
				format:  GlyphFormatOutline,
				advance: Vector16_16{X: 0, Y: 589824},
				Outline: Outline{
					Points: []Vector{
						{64, 72},
						{-640, 318},
						{-640, 409},
						{64, 651},
						{64, 552},
						{-118, 485},
						{-118, 225},
						{64, 158},
						{-191, 251},
						{-191, 460},
						{-501, 356},
					},
					Tags:     []byte{0x95, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11},
					Contours: []int16{7, 10},
					Flags:    0x131,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			glyph, err := tt.glyph()
			if err != nil {
				t.Fatalf("unable to load gltph: %v", err)
			}
			defer glyph.Free()
			if glyph.Glyph != nil {
				if err := glyph.Glyph.Transform(tt.matrix, tt.delta); err != tt.wantErr {
					t.Errorf("Glyph.Transform() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if diff := diff(glyph.Glyph, tt.want); diff != nil {
				t.Errorf("Glyph.Transform() = %v", diff)
			}
		})
	}
}

func TestGlyph_CBox(t *testing.T) {
	outline := newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72)
	bitmap := newTestGlyph(bungeeColorMac, 'A', LoadDefault, 14<<6, 72)

	tests := []struct {
		name  string
		glyph func() (testglyph, error)
		mode  BBoxMode
		want  BBox
	}{
		{name: "emptyGlyph", glyph: newZeroGlyph(), want: BBox{}},
		{name: "outline unscaled", glyph: outline, mode: GlyphBBoxUnscaled, want: BBox{XMin: 8, YMin: 0, XMax: 587, YMax: 704}},
		{name: "outline subpixels", glyph: outline, mode: GlyphBBoxSubpixels, want: BBox{XMin: 8, YMin: 0, XMax: 587, YMax: 704}},
		{name: "outline gridfit", glyph: outline, mode: GlyphBBoxGridfit, want: BBox{XMin: 0, YMin: 0, XMax: 640, YMax: 704}},
		{name: "outline truncate", glyph: outline, mode: GlyphBBoxTruncate, want: BBox{XMin: 0, YMin: 0, XMax: 9, YMax: 11}},
		{name: "outline pixels", glyph: outline, mode: GlyphBBoxPixels, want: BBox{XMin: 0, YMin: 0, XMax: 10, YMax: 11}},
		{name: "bitmap unscaled", glyph: bitmap, mode: GlyphBBoxUnscaled, want: BBox{XMin: 0, YMin: 0, XMax: 832, YMax: 960}},
		{name: "bitmap subpixels", glyph: bitmap, mode: GlyphBBoxSubpixels, want: BBox{XMin: 0, YMin: 0, XMax: 832, YMax: 960}},
		{name: "bitmap gridfit", glyph: bitmap, mode: GlyphBBoxGridfit, want: BBox{XMin: 0, YMin: 0, XMax: 832, YMax: 960}},
		{name: "bitmap truncate", glyph: bitmap, mode: GlyphBBoxTruncate, want: BBox{XMin: 0, YMin: 0, XMax: 13, YMax: 15}},
		{name: "bitmap pixels", glyph: bitmap, mode: GlyphBBoxPixels, want: BBox{XMin: 0, YMin: 0, XMax: 13, YMax: 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			glyph, err := tt.glyph()
			if err != nil {
				t.Fatalf("unable to load glyph: %v", err)
			}
			defer glyph.Free()

			if got := glyph.CBox(tt.mode); got != tt.want {
				t.Errorf("Glyph.CBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGlyph_ToBitmap(t *testing.T) {
	outline := newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72)
	bitmap := newTestGlyph(bungeeColorMac, 'A', LoadDefault, 14<<6, 72)

	type args struct {
		mode    RenderMode
		origin  Vector26_6
		destroy bool
	}
	tests := []struct {
		name    string
		glyph   func() (testglyph, error)
		args    args
		want    Glyph
		wantErr error
	}{
		{
			name:    "emptyGlyph",
			glyph:   newZeroGlyph(),
			want:    (*BitmapGlyph)(nil),
			wantErr: ErrInvalidArgument,
		},
		{
			name:  "bitmap",
			glyph: bitmap,
			args: args{
				mode:    RenderModeNormal,
				origin:  Vector26_6{},
				destroy: false,
			},
			want: &BitmapGlyph{
				format:  GlyphFormatBitmap,
				advance: Vector16_16{X: 917504, Y: 0},
				Left:    0,
				Top:     15,
				Bitmap: Bitmap{
					Rows:  15,
					Width: 13,
					Pitch: 13,
					Buffer: []byte{
						0x00, 0x00, 0x15, 0xbf, 0xc7, 0xc7, 0xc7, 0xc7, 0xc7, 0xbf, 0x15, 0x00, 0x00,
						0x00, 0x00, 0x80, 0xdf, 0xdd, 0xda, 0xdb, 0xda, 0xde, 0xdf, 0x80, 0x00, 0x00,
						0x00, 0x05, 0xd1, 0xde, 0xd4, 0xd9, 0xdb, 0xd9, 0xd4, 0xde, 0xd1, 0x05, 0x00,
						0x00, 0x3e, 0xdf, 0xd8, 0xda, 0xde, 0xd4, 0xdf, 0xd8, 0xd9, 0xdf, 0x3e, 0x00,
						0x00, 0xa6, 0xde, 0xd5, 0xde, 0x98, 0x07, 0xb8, 0xde, 0xd5, 0xdf, 0xa6, 0x00,
						0x13, 0xde, 0xdb, 0xd8, 0xdf, 0x3c, 0x00, 0x64, 0xdf, 0xd7, 0xdc, 0xde, 0x11,
						0x5e, 0xdf, 0xd7, 0xdc, 0xd4, 0x03, 0x00, 0x11, 0xdd, 0xdc, 0xd7, 0xdf, 0x5e,
						0xa8, 0xde, 0xd6, 0xde, 0xc7, 0x45, 0x45, 0x50, 0xd5, 0xde, 0xd6, 0xde, 0xa8,
						0xc6, 0xdd, 0xd6, 0xde, 0xde, 0xde, 0xdf, 0xde, 0xdd, 0xde, 0xd6, 0xdd, 0xc6,
						0xc8, 0xde, 0xd0, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd0, 0xde, 0xc8,
						0xc7, 0xdd, 0xd6, 0xde, 0xde, 0xdf, 0xdf, 0xdf, 0xde, 0xde, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x8f, 0x26, 0x26, 0x26, 0xaa, 0xdf, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x60, 0x00, 0x00, 0x00, 0x8a, 0xdf, 0xd6, 0xdd, 0xc7,
						0xce, 0xde, 0xde, 0xdf, 0x6c, 0x00, 0x00, 0x00, 0x96, 0xdf, 0xde, 0xde, 0xce,
						0xa1, 0xc8, 0xc7, 0xc9, 0x48, 0x00, 0x00, 0x00, 0x69, 0xc9, 0xc7, 0xc8, 0xa1,
					},
					NumGrays:  256,
					PixelMode: 2,
				},
			},
			wantErr: nil,
		},
		{
			name:  "bitmap origin",
			glyph: bitmap,
			args: args{
				mode:    RenderModeNormal,
				origin:  Vector26_6{X: 64, Y: 64},
				destroy: false,
			},
			want: &BitmapGlyph{
				format:  GlyphFormatBitmap,
				advance: Vector16_16{X: 917504, Y: 0},
				Left:    0,
				Top:     15,
				Bitmap: Bitmap{
					Rows:  15,
					Width: 13,
					Pitch: 13,
					Buffer: []byte{
						0x00, 0x00, 0x15, 0xbf, 0xc7, 0xc7, 0xc7, 0xc7, 0xc7, 0xbf, 0x15, 0x00, 0x00,
						0x00, 0x00, 0x80, 0xdf, 0xdd, 0xda, 0xdb, 0xda, 0xde, 0xdf, 0x80, 0x00, 0x00,
						0x00, 0x05, 0xd1, 0xde, 0xd4, 0xd9, 0xdb, 0xd9, 0xd4, 0xde, 0xd1, 0x05, 0x00,
						0x00, 0x3e, 0xdf, 0xd8, 0xda, 0xde, 0xd4, 0xdf, 0xd8, 0xd9, 0xdf, 0x3e, 0x00,
						0x00, 0xa6, 0xde, 0xd5, 0xde, 0x98, 0x07, 0xb8, 0xde, 0xd5, 0xdf, 0xa6, 0x00,
						0x13, 0xde, 0xdb, 0xd8, 0xdf, 0x3c, 0x00, 0x64, 0xdf, 0xd7, 0xdc, 0xde, 0x11,
						0x5e, 0xdf, 0xd7, 0xdc, 0xd4, 0x03, 0x00, 0x11, 0xdd, 0xdc, 0xd7, 0xdf, 0x5e,
						0xa8, 0xde, 0xd6, 0xde, 0xc7, 0x45, 0x45, 0x50, 0xd5, 0xde, 0xd6, 0xde, 0xa8,
						0xc6, 0xdd, 0xd6, 0xde, 0xde, 0xde, 0xdf, 0xde, 0xdd, 0xde, 0xd6, 0xdd, 0xc6,
						0xc8, 0xde, 0xd0, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd0, 0xde, 0xc8,
						0xc7, 0xdd, 0xd6, 0xde, 0xde, 0xdf, 0xdf, 0xdf, 0xde, 0xde, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x8f, 0x26, 0x26, 0x26, 0xaa, 0xdf, 0xd6, 0xdd, 0xc7,
						0xc7, 0xdd, 0xd6, 0xdf, 0x60, 0x00, 0x00, 0x00, 0x8a, 0xdf, 0xd6, 0xdd, 0xc7,
						0xce, 0xde, 0xde, 0xdf, 0x6c, 0x00, 0x00, 0x00, 0x96, 0xdf, 0xde, 0xde, 0xce,
						0xa1, 0xc8, 0xc7, 0xc9, 0x48, 0x00, 0x00, 0x00, 0x69, 0xc9, 0xc7, 0xc8, 0xa1,
					},
					NumGrays:  256,
					PixelMode: 2,
				},
			},
			wantErr: nil,
		},
		{
			name:  "outline",
			glyph: outline,
			args: args{
				mode:    RenderModeNormal,
				origin:  Vector26_6{},
				destroy: false,
			},
			want: &BitmapGlyph{
				format:  GlyphFormatBitmap,
				advance: Vector16_16{X: 589824, Y: 0},
				Left:    0,
				Top:     11,
				Bitmap: Bitmap{
					Rows:      11,
					Width:     10,
					Pitch:     10,
					Buffer:    goRegularToBitmap(),
					NumGrays:  256,
					PixelMode: 2,
				},
			},
			wantErr: nil,
		},
		{
			name:  "outline origin",
			glyph: outline,
			args: args{
				mode:    RenderModeNormal,
				origin:  Vector26_6{X: 64, Y: 64},
				destroy: false,
			},
			want: &BitmapGlyph{
				format:  GlyphFormatBitmap,
				advance: Vector16_16{X: 589824, Y: 0},
				Left:    1,
				Top:     12,
				Bitmap: Bitmap{
					Rows:      11,
					Width:     10,
					Pitch:     10,
					Buffer:    goRegularToBitmap(),
					NumGrays:  256,
					PixelMode: 2,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			glyph, err := tt.glyph()
			if err != nil {
				t.Fatalf("unable to load glyph: %v", err)
			}
			defer glyph.Free()

			got, err := glyph.ToBitmap(tt.args.mode, tt.args.origin, tt.args.destroy)
			if err != tt.wantErr {
				t.Errorf("Glyph.ToBitmap() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Glyph.ToBitmap() = %v", diff)
			}
		})
	}

	t.Run("destroy", func(t *testing.T) {
		wantOriginal := &OutlineGlyph{
			format:  GlyphFormatOutline,
			advance: Vector16_16{X: 589824, Y: 0},
			Outline: Outline{
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
			},
		}

		wantBitmap := &BitmapGlyph{
			format:  GlyphFormatBitmap,
			advance: Vector16_16{X: 589824, Y: 0},
			Left:    0,
			Top:     11,
			Bitmap: Bitmap{
				Rows:      11,
				Width:     10,
				Pitch:     10,
				Buffer:    goRegularToBitmap(),
				NumGrays:  256,
				PixelMode: 2,
			},
		}

		wantTransformed := &OutlineGlyph{
			format:  GlyphFormatOutline,
			advance: Vector16_16{X: 0, Y: 589824},
			Outline: Outline{
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
			},
		}

		basetest, err := newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72)()
		if err != nil {
			t.Fatalf("unable to load glyph: %v", err)
		}
		defer basetest.Free()
		base := basetest.Glyph

		copy, err := base.Copy()
		if err != nil {
			t.Fatalf("unable to copy: %v", err)
		}

		// test base
		if diff := diff(base, wantOriginal); diff != nil {
			t.Errorf("Glyph.ToBitmap() base = %v", diff)
		}
		if diff := diff(copy, wantOriginal); diff != nil {
			t.Errorf("Glyph.ToBitmap() copy = %v", diff)
		}

		// convert without destroying
		var gotBitmap Glyph
		{
			if gotBitmap, err = base.ToBitmap(RenderModeNormal, Vector26_6{}, false); err != nil {
				t.Fatalf("unable to convert to bitmap: %v", err)
			}
			if diff := diff(gotBitmap, wantBitmap); diff != nil {
				t.Errorf("Glyph.ToBitmap() converted = %v", diff)
			}

			// original should be intact
			if diff := diff(base, wantOriginal); diff != nil {
				t.Errorf("Glyph.ToBitmap() original = %v", diff)
			}
		}

		// apply transform to original
		{
			deg := float64(90)
			angle := deg / 360.0 * math.Pi * 2.0
			matrix := Matrix{
				Xx: (fixed.Int16_16)(math.Cos(angle) * 0x10000),
				Xy: (fixed.Int16_16)(-math.Sin(angle) * 0x10000),
				Yx: (fixed.Int16_16)(math.Sin(angle) * 0x10000),
				Yy: (fixed.Int16_16)(math.Cos(angle) * 0x10000),
			}
			if err := base.Transform(matrix, Vector26_6{}); err != nil {
				t.Fatalf("unable to apply transform: %v", err)
			}

			// original should be transformed
			if diff := diff(base, wantTransformed); diff != nil {
				t.Errorf("Glyph.ToBitmap() transformed = %v", diff)
			}
			// bitmap should be intact
			if diff := diff(gotBitmap, wantBitmap); diff != nil {
				t.Errorf("Glyph.ToBitmap() converted = %v", diff)
			}
		}

		// convert copy destroying
		{
			var gotBitmap Glyph
			if gotBitmap, err = copy.ToBitmap(RenderModeNormal, Vector26_6{}, true); err != nil {
				t.Fatalf("unable to convert to bitmap: %v", err)
			}
			if diff := diff(gotBitmap, wantBitmap); diff != nil {
				t.Errorf("Glyph.ToBitmap() converted = %v", diff)
			}

			// copy should be nil
			if copy.getptr() != nil {
				t.Errorf("Glyph.ToBitmap() copy.ptr is not nil")
			}
		}
	})
}

func TestGlyph_Format(t *testing.T) {
	tests := []struct {
		name  string
		glyph func() (testglyph, error)
		want  GlyphFormat
	}{
		{
			name:  "emptyOutline",
			glyph: newZeroGlyph(GlyphFormatOutline),
			want:  0,
		},
		{
			name:  "emptyBitmap",
			glyph: newZeroGlyph(GlyphFormatBitmap),
			want:  0,
		},
		{
			name:  "outline",
			glyph: newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72),
			want:  GlyphFormatOutline,
		},
		{
			name:  "bitmap",
			glyph: newTestGlyph(bungeeColorMac, 'A', LoadDefault, 14<<6, 72),
			want:  GlyphFormatBitmap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			glyph, err := tt.glyph()
			if err != nil {
				t.Fatalf("unable to load glyph: %v", err)
			}
			defer glyph.Free()

			if got := glyph.Format(); got != tt.want {
				t.Errorf("Glyph.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGlyph_Advance(t *testing.T) {
	tests := []struct {
		name  string
		glyph func() (testglyph, error)
		want  Vector16_16
	}{
		{
			name:  "emptyOutline",
			glyph: newZeroGlyph(GlyphFormatOutline),
			want:  Vector16_16{},
		},
		{
			name:  "emptyBitmap",
			glyph: newZeroGlyph(GlyphFormatBitmap),
			want:  Vector16_16{},
		},
		{
			name:  "outline",
			glyph: newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72),
			want:  Vector16_16{X: 589824, Y: 0},
		},
		{
			name:  "bitmap",
			glyph: newTestGlyph(bungeeColorMac, 'A', LoadDefault, 14<<6, 72),
			want:  Vector16_16{X: 917504, Y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			glyph, err := tt.glyph()
			if err != nil {
				t.Fatalf("unable to load glyph: %v", err)
			}
			defer glyph.Free()

			if got := glyph.Advance(); got != tt.want {
				t.Errorf("Glyph.Advance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGlyph_reload(t *testing.T) {
	testOutline, err := newTestGlyph(goRegular, 'A', LoadDefault, 14<<6, 72)()
	if err != nil {
		t.Fatalf("unable to load glyph: %v", err)
	}
	defer testOutline.Free()

	testBitmap, err := newTestGlyph(bungeeColorMac, 'A', LoadDefault, 14<<6, 72)()
	if err != nil {
		t.Fatalf("unable to load glyph: %v", err)
	}
	defer testBitmap.Free()

	outline, ok := testOutline.Glyph.(*OutlineGlyph)
	if !ok {
		t.Fatal("unable to convert glyph")
	}
	bitmap, ok := testBitmap.Glyph.(*BitmapGlyph)
	if !ok {
		t.Fatal("unable to convert glyph")
	}

	wantOutline := &OutlineGlyph{
		format:  GlyphFormatOutline,
		advance: Vector16_16{X: 589824, Y: 0},
		Outline: Outline{
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
		},
	}

	wantBitmap := &BitmapGlyph{
		format:  GlyphFormatBitmap,
		advance: Vector16_16{X: 917504, Y: 0},
		Left:    0,
		Top:     15,
		Bitmap: Bitmap{
			Rows:  15,
			Width: 13,
			Pitch: 13,
			Buffer: []byte{
				0x00, 0x00, 0x15, 0xbf, 0xc7, 0xc7, 0xc7, 0xc7, 0xc7, 0xbf, 0x15, 0x00, 0x00,
				0x00, 0x00, 0x80, 0xdf, 0xdd, 0xda, 0xdb, 0xda, 0xde, 0xdf, 0x80, 0x00, 0x00,
				0x00, 0x05, 0xd1, 0xde, 0xd4, 0xd9, 0xdb, 0xd9, 0xd4, 0xde, 0xd1, 0x05, 0x00,
				0x00, 0x3e, 0xdf, 0xd8, 0xda, 0xde, 0xd4, 0xdf, 0xd8, 0xd9, 0xdf, 0x3e, 0x00,
				0x00, 0xa6, 0xde, 0xd5, 0xde, 0x98, 0x07, 0xb8, 0xde, 0xd5, 0xdf, 0xa6, 0x00,
				0x13, 0xde, 0xdb, 0xd8, 0xdf, 0x3c, 0x00, 0x64, 0xdf, 0xd7, 0xdc, 0xde, 0x11,
				0x5e, 0xdf, 0xd7, 0xdc, 0xd4, 0x03, 0x00, 0x11, 0xdd, 0xdc, 0xd7, 0xdf, 0x5e,
				0xa8, 0xde, 0xd6, 0xde, 0xc7, 0x45, 0x45, 0x50, 0xd5, 0xde, 0xd6, 0xde, 0xa8,
				0xc6, 0xdd, 0xd6, 0xde, 0xde, 0xde, 0xdf, 0xde, 0xdd, 0xde, 0xd6, 0xdd, 0xc6,
				0xc8, 0xde, 0xd0, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0xd0, 0xde, 0xc8,
				0xc7, 0xdd, 0xd6, 0xde, 0xde, 0xdf, 0xdf, 0xdf, 0xde, 0xde, 0xd6, 0xdd, 0xc7,
				0xc7, 0xdd, 0xd6, 0xdf, 0x8f, 0x26, 0x26, 0x26, 0xaa, 0xdf, 0xd6, 0xdd, 0xc7,
				0xc7, 0xdd, 0xd6, 0xdf, 0x60, 0x00, 0x00, 0x00, 0x8a, 0xdf, 0xd6, 0xdd, 0xc7,
				0xce, 0xde, 0xde, 0xdf, 0x6c, 0x00, 0x00, 0x00, 0x96, 0xdf, 0xde, 0xde, 0xce,
				0xa1, 0xc8, 0xc7, 0xc9, 0x48, 0x00, 0x00, 0x00, 0x69, 0xc9, 0xc7, 0xc8, 0xa1,
			},
			NumGrays:  256,
			PixelMode: 2,
		},
	}

	if diff := diff(outline, wantOutline); diff != nil {
		t.Fatalf("wrong outline: %v", diff)
	}

	if diff := diff(bitmap, wantBitmap); diff != nil {
		t.Fatalf("wrong bitmap: %v", diff)
	}

	outlinePtr := outline.ptr
	bitmapPtr := bitmap.ptr

	outline.ptr = nil
	bitmap.ptr = nil
	outline.reload()
	bitmap.reload()

	if diff := diff(outline, &OutlineGlyph{}); diff != nil {
		t.Fatalf("wrong outline: %v", diff)
	}
	if diff := diff(bitmap, &BitmapGlyph{}); diff != nil {
		t.Fatalf("wrong outline: %v", diff)
	}

	var outlinePanic, bitmapPanic bool

	fn := func(g Glyph, panicked *bool) {
		defer func() {
			if r := recover(); r != nil {
				*panicked = true
			}
		}()
		g.reload()
	}

	outline.ptr = bitmapPtr
	bitmap.ptr = outlinePtr

	fn(outline, &outlinePanic)
	fn(bitmap, &bitmapPanic)

	if !outlinePanic {
		t.Fatal("outline should have panicked")
	}
	if !bitmapPanic {
		t.Fatal("bitmap should have panicked")
	}

	outline.ptr = outlinePtr // there's a defered Free
	bitmap.ptr = bitmapPtr   // there's a defered Free
}

func ExampleGlyph_Format_outline() {
	lib, err := NewLibrary()
	if err != nil {
		panic(err)
	}
	defer lib.Free()

	face, err := lib.NewFaceFromPath("../testdata/go/Go-Regular.ttf", 0, 0)
	if err != nil {
		panic(err)
	}
	defer face.Free()

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		panic(err)
	}

	if err := face.LoadChar('A', LoadDefault); err != nil {
		panic(err)
	}

	glyph, err := face.GlyphSlot().Glyph()
	if err != nil {
		panic(err)
	}

	format := glyph.Format()
	_, outlineOk := glyph.(*OutlineGlyph)
	_, bitmapOk := glyph.(*BitmapGlyph)
	fmt.Printf("format: %v, outline: %v, bitmap: %v\n", format, outlineOk, bitmapOk)

	// Output: format: Outline, outline: true, bitmap: false
}

func ExampleGlyph_Format_bitmap() {
	lib, err := NewLibrary()
	if err != nil {
		panic(err)
	}
	defer lib.Free()

	face, err := lib.NewFaceFromPath("../testdata/bungee/BungeeColor-Regular_sbix_MacOS.ttf", 0, 0)
	if err != nil {
		panic(err)
	}
	defer face.Free()

	if err := face.SelectSize(0); err != nil {
		panic(err)
	}

	if err := face.LoadChar('A', LoadDefault); err != nil {
		panic(err)
	}

	glyph, err := face.GlyphSlot().Glyph()
	if err != nil {
		panic(err)
	}

	format := glyph.Format()
	_, outlineOk := glyph.(*OutlineGlyph)
	_, bitmapOk := glyph.(*BitmapGlyph)
	fmt.Printf("format: %v, outline: %v, bitmap: %v\n", format, outlineOk, bitmapOk)

	// Output: format: Bitmap, outline: false, bitmap: true
}
