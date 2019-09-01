package freetype2

import "testing"

func TestFace_Advance(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		idx     GlyphIndex
		flags   LoadFlag
		want    Pos
		wantErr error
	}{
		{name: "nilFace", face: nilFace, wantErr: ErrInvalidFaceHandle},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadRender | LoadColor, want: 589824, wantErr: nil},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadDefault, want: 589824, wantErr: nil},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadNoScale, want: 1366, wantErr: nil},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadVerticalLayout, want: 917504, wantErr: nil},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadLinearDesign, want: 589824, wantErr: nil},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadComputeMetrics, want: 589824, wantErr: nil},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadRender | LoadColor | AdvanceFlagFastOnly, want: 0, wantErr: ErrUnimplementedFeature},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadDefault | AdvanceFlagFastOnly, want: 0, wantErr: ErrUnimplementedFeature},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadNoScale | AdvanceFlagFastOnly, want: 1366, wantErr: nil},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadVerticalLayout | AdvanceFlagFastOnly, want: 0, wantErr: ErrUnimplementedFeature},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadLinearDesign | AdvanceFlagFastOnly, want: 0, wantErr: ErrUnimplementedFeature},
		{name: "goRegular", face: goRegular, idx: 0x24, flags: LoadComputeMetrics | AdvanceFlagFastOnly, want: 0, wantErr: ErrUnimplementedFeature},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if face.Face != nil {
				if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					t.Fatalf("unable to set char size: %v", err)
				}
			}

			got, err := face.Advance(tt.idx, tt.flags)
			if err != tt.wantErr {
				t.Errorf("Face.Advance() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Face.Advance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFace_Advances(t *testing.T) {
	tests := []struct {
		name     string
		face     func() (testface, error)
		start    GlyphIndex
		count    int
		flags    LoadFlag
		want     []Pos
		wantErr  error
		wantFree bool
	}{
		{name: "nilFace", face: nilFace, wantErr: ErrInvalidFaceHandle},

		{name: "goRegular", face: goRegular, start: 0x24, count: 0, flags: LoadDefault, want: nil, wantErr: nil},

		{name: "goRegular", face: goRegular, start: 0x24, count: 1, flags: LoadDefault, want: []Pos{589824}, wantErr: nil, wantFree: true},

		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadRender | LoadColor, want: []Pos{589824, 589824, 655360, 655360, 589824}, wantErr: nil, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadDefault, want: []Pos{589824, 589824, 655360, 655360, 589824}, wantErr: nil, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadNoScale, want: []Pos{1366, 1366, 1479, 1479, 1366}, wantErr: nil, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadVerticalLayout, want: []Pos{917504, 917504, 917504, 917504, 917504}, wantErr: nil, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadLinearDesign, want: []Pos{589824, 589824, 655360, 655360, 589824}, wantErr: nil, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadComputeMetrics, want: []Pos{589824, 589824, 655360, 655360, 589824}, wantErr: nil, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadRender | LoadColor | AdvanceFlagFastOnly, want: nil, wantErr: ErrUnimplementedFeature, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadDefault | AdvanceFlagFastOnly, want: nil, wantErr: ErrUnimplementedFeature, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadNoScale | AdvanceFlagFastOnly, want: []Pos{1366, 1366, 1479, 1479, 1366}, wantErr: nil, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadVerticalLayout | AdvanceFlagFastOnly, want: nil, wantErr: ErrUnimplementedFeature, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadLinearDesign | AdvanceFlagFastOnly, want: nil, wantErr: ErrUnimplementedFeature, wantFree: true},
		{name: "goRegular", face: goRegular, start: 0x24, count: 5, flags: LoadComputeMetrics | AdvanceFlagFastOnly, want: nil, wantErr: ErrUnimplementedFeature, wantFree: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if face.Face != nil {
				if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
					t.Fatalf("unable to set char size: %v", err)
				}
			}

			var freed bool
			defer mockFree(func() { freed = true })()

			got, err := face.Advances(tt.start, tt.count, tt.flags)
			if err != tt.wantErr {
				t.Errorf("Face.Advances() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.Advances() = %v", diff)
			}
			if freed != tt.wantFree {
				t.Errorf("Face.Advances() free was not called")
			}
		})
	}
}
