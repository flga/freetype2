package freetype2

import (
	"testing"
)

func TestFace_WinFntHeader(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		want    WinFntHeader
		wantErr error
	}{
		{name: "nilFace", face: nilFace, want: WinFntHeader{}, wantErr: ErrInvalidFaceHandle},
		{name: "goRegular", face: goRegular, want: WinFntHeader{}, wantErr: ErrInvalidArgument},
		{
			name: "bitout",
			face: bitout,
			want: WinFntHeader{
				Version:  512,
				FileSize: 5792,
				Copyright: [60]byte{
					0x31, 0x39, 0x39, 0x39, 0x20, 0x4d, 0x61, 0x67, 0x6e, 0x75,
					0x73, 0x20, 0x48, 0xf6, 0x67, 0x62, 0x65, 0x72, 0x67, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
				},
				FileType:             0,
				NominalPointSize:     11,
				VerticalResolution:   96,
				HorizontalResolution: 96,
				Ascent:               14,
				InternalLeading:      2,
				ExternalLeading:      0,
				Italic:               0,
				Underline:            0,
				StrikeOut:            0,
				Weight:               400,
				Charset:              0,
				PixelWidth:           0,
				PixelHeight:          17,
				PitchAndFamily:       33,
				AvgWidth:             7,
				MaxWidth:             15,
				FirstChar:            30,
				LastChar:             255,
				DefaultChar:          31,
				BreakChar:            32,
				BytesPerRow:          280,
				DeviceOffset:         0,
				FaceNameOffset:       5786,
				BitsPointer:          0,
				BitsOffset:           1026,
				Reserved:             0,
				Flags:                0,
				ASpace:               0,
				BSpace:               0,
				CSpace:               0,
				ColorTableOffset:     0,
				Reserved1:            [4]uint64{2819521476297798, 3664165449761896, 0, 0},
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

			got, err := face.WinFntHeader()
			if err != tt.wantErr {
				t.Errorf("Face.WinFntHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.WinFntHeader() = %v", diff)
			}
		})
	}
}
