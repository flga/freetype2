package freetype2

import "testing"

func TestFace_Gasp(t *testing.T) {
	tests := []struct {
		name         string
		face         func() (testface, error)
		verticalPPem int
		want         GaspFlag
		wantOk       bool
	}{
		{name: "nilFace", face: nilFace, verticalPPem: 0, want: 0, wantOk: false},
		{name: "bungeeColorWin", face: bungeeColorWin, verticalPPem: 64, want: 0, wantOk: false},
		{name: "goRegular", face: goRegular, verticalPPem: 64, want: GaspFlagDoGridfit | GaspFlagDoGray | GaspFlagSymmetricGridfit | GaspFlagSymmetricSmoothing, wantOk: true},
		{name: "twemojiMozilla", face: twemojiMozilla, verticalPPem: 64, want: GaspFlagDoGray, wantOk: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if got, gotOk := face.GaspFlags(tt.verticalPPem); got != tt.want || gotOk != tt.wantOk {
				t.Errorf("Face.GaspFlags() = %v, %v want %v, %v", got, gotOk, tt.want, tt.wantOk)
			}
		})
	}
}
