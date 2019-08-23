package freetype2

import (
	"testing"
)

func TestFace_BDFCharsetID(t *testing.T) {
	tests := []struct {
		name         string
		face         func() (testface, error)
		wantEncoding string
		wantRegistry string
		wantErr      error
	}{
		{name: "nilFace", face: nilFace, wantEncoding: "", wantRegistry: "", wantErr: ErrInvalidFaceHandle},
		{name: "gohuBdf", face: gohuBdf, wantEncoding: "1", wantRegistry: "ISO8859", wantErr: nil},
		{name: "gohuPcf", face: gohuPcf, wantEncoding: "1", wantRegistry: "ISO8859", wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			gotEncoding, gotRegistry, err := face.BDFCharsetID()
			if err != tt.wantErr {
				t.Errorf("Face.BDFCharsetID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEncoding != tt.wantEncoding {
				t.Errorf("Face.BDFCharsetID() gotEncoding = %v, want %v", gotEncoding, tt.wantEncoding)
			}
			if gotRegistry != tt.wantRegistry {
				t.Errorf("Face.BDFCharsetID() gotRegistry = %v, want %v", gotRegistry, tt.wantRegistry)
			}
		})
	}
}

func TestFace_BDFProperty(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		prop    string
		want    BDFProperty
		wantErr error
	}{
		{name: "nilFace", face: nilFace, prop: "", want: BDFProperty{}, wantErr: ErrInvalidFaceHandle},

		{name: "gohuBdf-thisDoesNotExist", face: gohuBdf, prop: "thisDoesNotExist", want: BDFProperty{}, wantErr: ErrInvalidArgument},
		{name: "gohuPcf-thisDoesNotExist", face: gohuPcf, prop: "thisDoesNotExist", want: BDFProperty{}, wantErr: ErrInvalidArgument},

		{name: "gohuBdf-FOUNDRY", face: gohuBdf, prop: "FOUNDRY", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "Gohu"}, wantErr: nil},
		{name: "gohuBdf-FAMILY_NAME", face: gohuBdf, prop: "FAMILY_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "GohuFont"}, wantErr: nil},
		{name: "gohuBdf-WEIGHT_NAME", face: gohuBdf, prop: "WEIGHT_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "Medium"}, wantErr: nil},
		{name: "gohuBdf-SLANT", face: gohuBdf, prop: "SLANT", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "R"}, wantErr: nil},
		{name: "gohuBdf-SETWIDTH_NAME", face: gohuBdf, prop: "SETWIDTH_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "Normal"}, wantErr: nil},
		{name: "gohuBdf-ADD_STYLE_NAME", face: gohuBdf, prop: "ADD_STYLE_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: ""}, wantErr: nil},
		{name: "gohuBdf-PIXEL_SIZE", face: gohuBdf, prop: "PIXEL_SIZE", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 11}, wantErr: nil},
		{name: "gohuBdf-POINT_SIZE", face: gohuBdf, prop: "POINT_SIZE", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 80}, wantErr: nil},
		{name: "gohuBdf-RESOLUTION_X", face: gohuBdf, prop: "RESOLUTION_X", want: BDFProperty{Type: BDFPropertyTypeCardinal, Cardinal: 100}, wantErr: nil},
		{name: "gohuBdf-RESOLUTION_Y", face: gohuBdf, prop: "RESOLUTION_Y", want: BDFProperty{Type: BDFPropertyTypeCardinal, Cardinal: 100}, wantErr: nil},
		{name: "gohuBdf-SPACING", face: gohuBdf, prop: "SPACING", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "C"}, wantErr: nil},
		{name: "gohuBdf-AVERAGE_WIDTH", face: gohuBdf, prop: "AVERAGE_WIDTH", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 60}, wantErr: nil},
		{name: "gohuBdf-CHARSET_REGISTRY", face: gohuBdf, prop: "CHARSET_REGISTRY", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "ISO8859"}, wantErr: nil},
		{name: "gohuBdf-CHARSET_ENCODING", face: gohuBdf, prop: "CHARSET_ENCODING", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "1"}, wantErr: nil},
		{name: "gohuBdf-FONTNAME_REGISTRY", face: gohuBdf, prop: "FONTNAME_REGISTRY", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: ""}, wantErr: nil},
		{name: "gohuBdf-FONT_NAME", face: gohuBdf, prop: "FONT_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "GohuFont"}, wantErr: nil},
		{name: "gohuBdf-FACE_NAME", face: gohuBdf, prop: "FACE_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "GohuFont"}, wantErr: nil},
		{name: "gohuBdf-FONT_VERSION", face: gohuBdf, prop: "FONT_VERSION", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "003.000"}, wantErr: nil},
		{name: "gohuBdf-FONT_ASCENT", face: gohuBdf, prop: "FONT_ASCENT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 9}, wantErr: nil},
		{name: "gohuBdf-FONT_DESCENT", face: gohuBdf, prop: "FONT_DESCENT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 2}, wantErr: nil},
		{name: "gohuBdf-UNDERLINE_POSITION", face: gohuBdf, prop: "UNDERLINE_POSITION", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: -1}, wantErr: nil},
		{name: "gohuBdf-UNDERLINE_THICKNESS", face: gohuBdf, prop: "UNDERLINE_THICKNESS", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 1}, wantErr: nil},
		{name: "gohuBdf-X_HEIGHT", face: gohuBdf, prop: "X_HEIGHT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 4}, wantErr: nil},
		{name: "gohuBdf-CAP_HEIGHT", face: gohuBdf, prop: "CAP_HEIGHT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 7}, wantErr: nil},
		{name: "gohuBdf-RAW_ASCENT", face: gohuBdf, prop: "RAW_ASCENT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 818}, wantErr: nil},
		{name: "gohuBdf-RAW_DESCENT", face: gohuBdf, prop: "RAW_DESCENT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 182}, wantErr: nil},
		{name: "gohuBdf-NORM_SPACE", face: gohuBdf, prop: "NORM_SPACE", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 6}, wantErr: nil},
		{name: "gohuBdf-FIGURE_WIDTH", face: gohuBdf, prop: "FIGURE_WIDTH", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 6}, wantErr: nil},
		{name: "gohuBdf-AVG_LOWERCASE_WIDTH", face: gohuBdf, prop: "AVG_LOWERCASE_WIDTH", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 60}, wantErr: nil},
		{name: "gohuBdf-AVG_UPPERCASE_WIDTH", face: gohuBdf, prop: "AVG_UPPERCASE_WIDTH", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "60"}, wantErr: nil},

		{name: "gohuPcf-FOUNDRY", face: gohuPcf, prop: "FOUNDRY", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "Gohu"}, wantErr: nil},
		{name: "gohuPcf-FAMILY_NAME", face: gohuPcf, prop: "FAMILY_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "GohuFont"}, wantErr: nil},
		{name: "gohuPcf-WEIGHT_NAME", face: gohuPcf, prop: "WEIGHT_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "Medium"}, wantErr: nil},
		{name: "gohuPcf-SLANT", face: gohuPcf, prop: "SLANT", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "R"}, wantErr: nil},
		{name: "gohuPcf-SETWIDTH_NAME", face: gohuPcf, prop: "SETWIDTH_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "Normal"}, wantErr: nil},
		{name: "gohuPcf-ADD_STYLE_NAME", face: gohuPcf, prop: "ADD_STYLE_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: ""}, wantErr: nil},
		{name: "gohuPcf-PIXEL_SIZE", face: gohuPcf, prop: "PIXEL_SIZE", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 11}, wantErr: nil},
		{name: "gohuPcf-POINT_SIZE", face: gohuPcf, prop: "POINT_SIZE", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 80}, wantErr: nil},
		{name: "gohuPcf-RESOLUTION_X", face: gohuPcf, prop: "RESOLUTION_X", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 100}, wantErr: nil},
		{name: "gohuPcf-RESOLUTION_Y", face: gohuPcf, prop: "RESOLUTION_Y", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 100}, wantErr: nil},
		{name: "gohuPcf-SPACING", face: gohuPcf, prop: "SPACING", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "C"}, wantErr: nil},
		{name: "gohuPcf-AVERAGE_WIDTH", face: gohuPcf, prop: "AVERAGE_WIDTH", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 60}, wantErr: nil},
		{name: "gohuPcf-CHARSET_REGISTRY", face: gohuPcf, prop: "CHARSET_REGISTRY", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "ISO8859"}, wantErr: nil},
		{name: "gohuPcf-CHARSET_ENCODING", face: gohuPcf, prop: "CHARSET_ENCODING", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "1"}, wantErr: nil},
		{name: "gohuPcf-FONTNAME_REGISTRY", face: gohuPcf, prop: "FONTNAME_REGISTRY", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: ""}, wantErr: nil},
		{name: "gohuPcf-FONT_NAME", face: gohuPcf, prop: "FONT_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "GohuFont"}, wantErr: nil},
		{name: "gohuPcf-FACE_NAME", face: gohuPcf, prop: "FACE_NAME", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "GohuFont"}, wantErr: nil},
		{name: "gohuPcf-FONT_VERSION", face: gohuPcf, prop: "FONT_VERSION", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "003.000"}, wantErr: nil},
		{name: "gohuPcf-FONT_ASCENT", face: gohuPcf, prop: "FONT_ASCENT", want: BDFProperty{}, wantErr: ErrInvalidArgument},
		{name: "gohuPcf-FONT_DESCENT", face: gohuPcf, prop: "FONT_DESCENT", want: BDFProperty{}, wantErr: ErrInvalidArgument},
		{name: "gohuPcf-UNDERLINE_POSITION", face: gohuPcf, prop: "UNDERLINE_POSITION", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: -1}, wantErr: nil},
		{name: "gohuPcf-UNDERLINE_THICKNESS", face: gohuPcf, prop: "UNDERLINE_THICKNESS", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 1}, wantErr: nil},
		{name: "gohuPcf-X_HEIGHT", face: gohuPcf, prop: "X_HEIGHT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 4}, wantErr: nil},
		{name: "gohuPcf-CAP_HEIGHT", face: gohuPcf, prop: "CAP_HEIGHT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 7}, wantErr: nil},
		{name: "gohuPcf-RAW_ASCENT", face: gohuPcf, prop: "RAW_ASCENT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 818}, wantErr: nil},
		{name: "gohuPcf-RAW_DESCENT", face: gohuPcf, prop: "RAW_DESCENT", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 182}, wantErr: nil},
		{name: "gohuPcf-NORM_SPACE", face: gohuPcf, prop: "NORM_SPACE", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 6}, wantErr: nil},
		{name: "gohuPcf-FIGURE_WIDTH", face: gohuPcf, prop: "FIGURE_WIDTH", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 6}, wantErr: nil},
		{name: "gohuPcf-AVG_LOWERCASE_WIDTH", face: gohuPcf, prop: "AVG_LOWERCASE_WIDTH", want: BDFProperty{Type: BDFPropertyTypeInteger, Integer: 60}, wantErr: nil},
		{name: "gohuPcf-AVG_UPPERCASE_WIDTH", face: gohuPcf, prop: "AVG_UPPERCASE_WIDTH", want: BDFProperty{Type: BDFPropertyTypeAtom, Atom: "60"}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			got, err := face.BDFProperty(tt.prop)
			if err != tt.wantErr {
				t.Errorf("Face.BDFProperty() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.BDFProperty() = %v", diff)
			}
		})
	}
}
