package freetype2

import (
	"testing"

	"github.com/flga/freetype2/fixed"
)

func TestFace_HasPSGlyphNames(t *testing.T) {
	tests := []struct {
		name string
		face func() (testface, error)
		want bool
	}{
		{
			name: "nilFace",
			face: nilFace,
			want: false,
		},
		{
			name: "goRegular",
			face: goRegular,
			want: false,
		},
		{
			name: "bungeeLayersReg",
			face: bungeeLayersReg,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if got := face.HasPSGlyphNames(); got != tt.want {
				t.Errorf("Face.HasPSGlyphNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFace_PSFontInfo(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		want    PSFontInfo
		wantErr error
	}{
		{
			name:    "nilFace",
			face:    nilFace,
			want:    PSFontInfo{},
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name:    "goRegular",
			face:    goRegular,
			want:    PSFontInfo{},
			wantErr: ErrInvalidArgument,
		},
		{
			name: "bungeeLayersReg",
			face: bungeeLayersReg,
			want: PSFontInfo{
				Version:            "1.0",
				Notice:             "Bungee is a trademark of The Font Bureau.",
				FullName:           "Bungee Layers Regular",
				FamilyName:         "Bungee Layers",
				Weight:             "Normal",
				ItalicAngle:        0,
				IsFixedPitch:       false,
				UnderlinePosition:  0,
				UnderlineThickness: 0,
			},
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

			got, err := face.PSFontInfo()
			if err != tt.wantErr {
				t.Errorf("Face.PSFontInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.PSFontInfo() = %v", diff)
			}
		})
	}
}

func TestFace_PSPrivate(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		want    PSPrivate
		wantErr error
	}{
		{
			name:    "nilFace",
			face:    nilFace,
			want:    PSPrivate{},
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name:    "goRegular",
			face:    goRegular,
			want:    PSPrivate{},
			wantErr: ErrInvalidArgument,
		},
		{
			name:    "bungeeLayersReg",
			face:    bungeeLayersReg,
			want:    PSPrivate{},
			wantErr: ErrInvalidArgument,
		},
		{
			name: "nimbusMono",
			face: nimbusMono,
			want: PSPrivate{
				UniqueID:            0,
				LenIV:               4,
				NumBlueValues:       8,
				NumOtherBlues:       0,
				NumFamilyBlues:      0,
				NumFamilyOtherBlues: 0,
				BlueValues:          [14]int16{-16, 0, 417, 433, 563, 575, 603, 616, 0, 0, 0, 0, 0, 0},
				OtherBlues:          [10]int16{},
				FamilyBlues:         [14]int16{},
				FamilyOtherBlues:    [10]int16{},
				BlueScale:           fixed.Int16_16(2596864),
				BlueShift:           7,
				BlueFuzz:            1,
				StandardWidth:       52,
				StandardHeight:      51,
				NumSnapWidths:       12,
				NumSnapHeights:      12,
				ForceBold:           false,
				RoundStemUp:         false,
				SnapWidths:          [13]int16{28, 36, 39, 43, 47, 52, 56, 59, 66, 114, 149, 179, 0},
				SnapHeights:         [13]int16{35, 39, 43, 47, 51, 54, 58, 67, 72, 113, 133, 152, 0},
				ExpansionFactor:     fixed.Int16_16(3932),
				LanguageGroup:       0,
				Password:            5839,
				MinFeature:          [2]int16{16, 16},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load font: %v", err)
			}
			defer face.Free()

			got, err := face.PSPrivate()
			if err != tt.wantErr {
				t.Errorf("Face.PSPrivate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.PSPrivate() = %v", diff)
			}
		})
	}
}
