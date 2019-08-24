package freetype2

import (
	"testing"
)

func TestFace_PFRMetrics(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		want    PFRMetrics
		wantErr error
	}{
		{name: "nilFace", face: nilFace, want: PFRMetrics{}, wantErr: ErrInvalidFaceHandle},
		{name: "bungeeColorWin", face: bungeeColorWin, want: PFRMetrics{OutlineResolution: 1000, MetricsResolution: 1000, MetricsXScale: 0, MetricsYScale: 0}, wantErr: ErrUnknownFileFormat},
		{name: "bungeeColorMac", face: bungeeColorMac, want: PFRMetrics{OutlineResolution: 0, MetricsResolution: 0, MetricsXScale: 0, MetricsYScale: 0}, wantErr: ErrUnknownFileFormat},
		{name: "goRegular", face: goRegular, want: PFRMetrics{OutlineResolution: 2048, MetricsResolution: 2048, MetricsXScale: 0, MetricsYScale: 0}, wantErr: ErrUnknownFileFormat},
		{name: "notoSansJpReg", face: notoSansJpReg, want: PFRMetrics{OutlineResolution: 1000, MetricsResolution: 1000, MetricsXScale: 0, MetricsYScale: 0}, wantErr: ErrUnknownFileFormat},
		{name: "bungeeLayersReg", face: bungeeLayersReg, want: PFRMetrics{OutlineResolution: 1000, MetricsResolution: 1000, MetricsXScale: 0, MetricsYScale: 0}, wantErr: ErrUnknownFileFormat},
		{name: "nimbusMono", face: nimbusMono, want: PFRMetrics{OutlineResolution: 1000, MetricsResolution: 1000, MetricsXScale: 0, MetricsYScale: 0}, wantErr: ErrUnknownFileFormat},
		{name: "gohuBdf", face: gohuBdf, want: PFRMetrics{OutlineResolution: 0, MetricsResolution: 0, MetricsXScale: 0, MetricsYScale: 0}, wantErr: ErrUnknownFileFormat},
		{name: "gohuPcf", face: gohuPcf, want: PFRMetrics{OutlineResolution: 0, MetricsResolution: 0, MetricsXScale: 0, MetricsYScale: 0}, wantErr: ErrUnknownFileFormat},
		{name: "amelia", face: amelia, want: PFRMetrics{OutlineResolution: 2048, MetricsResolution: 2048, MetricsXScale: 0, MetricsYScale: 0}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			got, err := face.PFRMetrics()
			if err != tt.wantErr {
				t.Errorf("Face.PFRMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.PFRMetrics() = %v", diff)
			}
		})
	}
}

func TestFace_PFRKerning(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		left    GlyphIndex
		right   GlyphIndex
		want    Vector
		wantErr error
	}{
		{name: "nilFace", face: nilFace, left: 0, right: 0, want: Vector{}, wantErr: ErrInvalidFaceHandle},
		{name: "bungeeColorWin", face: bungeeColorWin, left: 43, right: 65, want: Vector{X: 0, Y: 0}, wantErr: nil},
		{name: "bungeeColorMac", face: bungeeColorMac, left: 43, right: 65, want: Vector{X: 0, Y: 0}, wantErr: nil},
		{name: "goRegular", face: goRegular, left: 36, right: 58, want: Vector{X: 0, Y: 0}, wantErr: nil},
		{name: "notoSansJpReg", face: notoSansJpReg, left: 34, right: 56, want: Vector{X: 0, Y: 0}, wantErr: nil},
		{name: "arimoRegular", face: arimoRegular, left: 36, right: 58, want: Vector{X: -76, Y: 0}, wantErr: nil},
		{name: "bungeeLayersReg", face: bungeeLayersReg, left: 2, right: 99, want: Vector{X: 0, Y: 0}, wantErr: nil},
		{name: "nimbusMono", face: nimbusMono, left: 854, right: 22, want: Vector{X: 0, Y: 0}, wantErr: nil},
		{name: "gohuBdf", face: gohuBdf, left: 34, right: 56, want: Vector{X: 0, Y: 0}, wantErr: nil},
		{name: "gohuPcf", face: gohuPcf, left: 34, right: 56, want: Vector{X: 0, Y: 0}, wantErr: nil},
		{name: "amelia", face: amelia, left: 36, right: 58, want: Vector{X: 0, Y: 0}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			got, err := face.PFRKerning(tt.left, tt.right)
			if err != tt.wantErr {
				t.Errorf("Face.PFRKerning() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Face.PFRKerning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFace_PFRAdvance(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		idx     GlyphIndex
		want    Pos
		wantErr error
	}{
		{name: "nilFace", face: nilFace, idx: 0, want: 0, wantErr: ErrInvalidFaceHandle},
		{name: "bungeeColorWin", face: bungeeColorWin, idx: 43, want: 0, wantErr: ErrInvalidArgument},
		{name: "bungeeColorMac", face: bungeeColorMac, idx: 43, want: 0, wantErr: ErrInvalidArgument},
		{name: "goRegular", face: goRegular, idx: 36, want: 0, wantErr: ErrInvalidArgument},
		{name: "notoSansJpReg", face: notoSansJpReg, idx: 34, want: 0, wantErr: ErrInvalidArgument},
		{name: "bungeeLayersReg", face: bungeeLayersReg, idx: 2, want: 0, wantErr: ErrInvalidArgument},
		{name: "nimbusMono", face: nimbusMono, idx: 854, want: 0, wantErr: ErrInvalidArgument},
		{name: "gohuBdf", face: gohuBdf, idx: 34, want: 0, wantErr: ErrInvalidArgument},
		{name: "gohuPcf", face: gohuPcf, idx: 34, want: 0, wantErr: ErrInvalidArgument},
		{name: "amelia", face: amelia, idx: 36, want: 1053, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			got, err := face.PFRAdvance(tt.idx)
			if err != tt.wantErr {
				t.Errorf("Face.PFRAdvance() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Face.PFRAdvance() = %v, want %v", got, tt.want)
			}
		})
	}
}
