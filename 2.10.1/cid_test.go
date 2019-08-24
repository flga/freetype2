package freetype2

import (
	"testing"
)

func TestFace_CIDRegistryOrderingSupplement(t *testing.T) {
	tests := []struct {
		name           string
		face           func() (testface, error)
		wantRegistry   string
		wantOrdering   string
		wantSupplement int
		wantErr        error
	}{
		{name: "nilFace", face: nilFace, wantErr: ErrInvalidArgument},
		{name: "bungeeColorWin", face: bungeeColorWin, wantErr: ErrInvalidArgument},
		{name: "bungeeColorMac", face: bungeeColorMac, wantErr: ErrInvalidArgument},
		{name: "goRegular", face: goRegular, wantErr: ErrInvalidArgument},
		{name: "notoSansJpReg", face: notoSansJpReg, wantRegistry: "Adobe", wantOrdering: "Identity", wantSupplement: 0, wantErr: nil},
		{name: "bungeeLayersReg", face: bungeeLayersReg, wantErr: ErrInvalidArgument},
		{name: "nimbusMono", face: nimbusMono, wantErr: ErrInvalidArgument},
		{name: "gohuBdf", face: gohuBdf, wantErr: ErrInvalidArgument},
		{name: "gohuPcf", face: gohuPcf, wantErr: ErrInvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			gotRegistry, gotOrdering, gotSupplement, err := face.CIDRegistryOrderingSupplement()
			if err != tt.wantErr {
				t.Errorf("Face.CIDRegistryOrderingSupplement() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotRegistry != tt.wantRegistry {
				t.Errorf("Face.CIDRegistryOrderingSupplement() gotRegistry = %v, want %v", gotRegistry, tt.wantRegistry)
			}
			if gotOrdering != tt.wantOrdering {
				t.Errorf("Face.CIDRegistryOrderingSupplement() gotOrdering = %v, want %v", gotOrdering, tt.wantOrdering)
			}
			if gotSupplement != tt.wantSupplement {
				t.Errorf("Face.CIDRegistryOrderingSupplement() gotSupplement = %v, want %v", gotSupplement, tt.wantSupplement)
			}
		})
	}
}

func TestFace_IsInternallyCIDKeyed(t *testing.T) {
	tests := []struct {
		name string
		face func() (testface, error)
		want bool
	}{
		{name: "nilFace", face: nilFace, want: false},
		{name: "bungeeColorWin", face: bungeeColorWin, want: false},
		{name: "bungeeColorMac", face: bungeeColorMac, want: false},
		{name: "goRegular", face: goRegular, want: false},
		{name: "notoSansJpReg", face: notoSansJpReg, want: true},
		{name: "bungeeLayersReg", face: bungeeLayersReg, want: false},
		{name: "nimbusMono", face: nimbusMono, want: false},
		{name: "gohuBdf", face: gohuBdf, want: false},
		{name: "gohuPcf", face: gohuPcf, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if got := face.IsInternallyCIDKeyed(); got != tt.want {
				t.Errorf("Face.IsInternallyCIDKeyed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFace_CIDFromGlyphIndex(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		idx     GlyphIndex
		want    uint
		wantErr error
	}{
		{name: "nilFace", face: nilFace, idx: 0, want: 0, wantErr: ErrInvalidArgument},
		{name: "bungeeColorWin", face: bungeeColorWin, idx: 0, want: 0, wantErr: ErrInvalidArgument},
		{name: "bungeeColorMac", face: bungeeColorMac, idx: 0, want: 0, wantErr: ErrInvalidArgument},
		{name: "goRegular", face: goRegular, idx: 0, want: 0, wantErr: ErrInvalidArgument},
		{name: "notoSansJpReg-0", face: notoSansJpReg, idx: 0, want: 0, wantErr: nil},
		{name: "notoSansJpReg-1200", face: notoSansJpReg, idx: 1200, want: 1456, wantErr: nil},
		{name: "bungeeLayersReg", face: bungeeLayersReg, idx: 0, want: 0, wantErr: ErrInvalidArgument},
		{name: "nimbusMono", face: nimbusMono, idx: 0, want: 0, wantErr: ErrInvalidArgument},
		{name: "gohuBdf", face: gohuBdf, idx: 0, want: 0, wantErr: ErrInvalidArgument},
		{name: "gohuPcf", face: gohuPcf, idx: 0, want: 0, wantErr: ErrInvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			got, err := face.CIDFromGlyphIndex(tt.idx)
			if err != tt.wantErr {
				t.Errorf("Face.CIDFromGlyphIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Face.CIDFromGlyphIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
