package freetype2

import (
	"testing"
)

func TestFace_NewSize(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		want    *Size
		wantErr error
	}{
		{
			name:    "nilFace",
			face:    nilFace,
			want:    nil,
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name:    "goRegular",
			face:    goRegular,
			want:    &Size{},
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

			got, err := face.NewSize()
			if err != tt.wantErr {
				t.Errorf("Face.NewSize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.NewSize() = %v", diff)
			}
		})
	}

	t.Run("free", func(t *testing.T) {
		face, err := goRegular()
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}

		got, err := face.NewSize()
		if err != nil {
			t.Fatalf("Face.NewSize() error = %v", err)
		}

		if got.ptr == nil {
			t.Fatalf("Face.NewSize() ptr is nil")
		}

		face.Free()
		if got.ptr != nil {
			t.Fatalf("Face.NewSize() ptr is not nil")
		}
	})
}

func TestFace_ActivateSize(t *testing.T) {
	var nilFace *Face
	if err := nilFace.ActivateSize(nil); err != ErrInvalidFaceHandle {
		t.Errorf("Face.ActivateSize() error = %v, want %v", err, ErrInvalidFaceHandle)
	}

	face, err := goRegular()
	if err != nil {
		t.Fatalf("unable to load face: %v", err)
	}

	if err := face.ActivateSize(nil); err != ErrInvalidSizeHandle {
		t.Errorf("Face.ActivateSize() error = %v, want %v", err, ErrInvalidSizeHandle)
	}

	if err := face.SetCharSize(14<<6, 14<<6, 72, 72); err != nil {
		t.Fatalf("unable to set char size: %v", err)
	}

	want := &Size{
		SizeMetrics: SizeMetrics{
			XPpem:      14,
			YPpem:      14,
			XScale:     28672,
			YScale:     28672,
			Ascender:   896,
			Descender:  -192,
			Height:     1024,
			MaxAdvance: 960,
		},
	}

	if diff := diff(face.Size(), want); diff != nil {
		t.Fatalf("size is different: %v", diff)
	}

	s, err := face.NewSize()
	if err != nil {
		t.Fatalf("unable to create size: %v", err)
	}

	if err := face.ActivateSize(s); err != nil {
		t.Errorf("Face.ActivateSize() error = %v", err)
	}

	if diff := diff(face.Size(), &Size{}); diff != nil {
		t.Fatalf("size is different: %v", diff)
	}
}
