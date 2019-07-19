package freetype2

import (
	"os"
	"testing"
)

func TestNewLibrary(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to init lib: %s", err)
	}
	if l == nil || l.ptr == nil {
		t.Fatalf("got a nil library")
	}
	if err := l.Free(); err != nil {
		t.Fatalf("unable to free library: %s", err)
	}
	if l.ptr != nil {
		t.Fatalf("Free should set ptr to nil")
	}
	if err := l.Free(); err != nil {
		t.Fatalf("Free on an already freed library should be a noop, got: %s", err)
	}
}

func TestVersion(t *testing.T) {
	t.Run("Library.Version()", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to init lib: %s", err)
		}
		defer l.Free()

		want := Version{Major: 2, Minor: 10, Patch: 1}
		if got := l.Version(); got != want {
			t.Fatalf("unexpected version, want: %v, got %v", want, got)
		}

		var nilLib *Library
		want = Version{}
		if got := nilLib.Version(); got != want {
			t.Fatalf("Version should return 0 value when lib is nil")
		}
	})

	t.Run("Version.String()", func(t *testing.T) {
		v := Version{Major: 0, Minor: 1, Patch: 2}
		want := "0.1.2"
		if got := v.String(); got != want {
			t.Fatalf("want: %v, got %v", want, got)
		}
	})
}

func TestLibraryFree(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to init lib: %s", err)
	}

	face1, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0)
	if err != nil {
		t.Fatalf("unable to create face: %s", err)
	}
	face2, err := l.NewFaceFromPath(testdata("go", "Go-Bold.ttf"), 0)
	if err != nil {
		t.Fatalf("unable to create face: %s", err)
	}

	if err := l.Free(); err != nil {
		t.Fatalf("unable to free lib: %s", err)
	}

	if face1.ptr != nil {
		t.Fatalf("lib was freed but face1.ptr was not set to nil")
	}
	if face2.ptr != nil {
		t.Fatalf("lib was freed but face2.ptr was not set to nil")
	}

	if want, got := 0, len(l.faces); got != want {
		t.Fatalf("lib retained references to faces, wanted len(l.faces) to be %d, got %d", want, got)
	}
}

func TestNewFace(t *testing.T) {
	type testCase struct{ filename, family, style string }
	tests := []testCase{
		{filename: "Go-Bold-Italic.ttf", family: "Go", style: "Bold Italic"},
		{filename: "Go-Bold.ttf", family: "Go", style: "Bold"},
		{filename: "Go-Italic.ttf", family: "Go", style: "Italic"},
		{filename: "Go-Medium-Italic.ttf", family: "Go Medium", style: "Italic"},
		{filename: "Go-Medium.ttf", family: "Go Medium", style: "Regular"},
		{filename: "Go-Mono-Bold-Italic.ttf", family: "Go Mono", style: "Bold Italic"},
		{filename: "Go-Mono-Bold.ttf", family: "Go Mono", style: "Bold"},
		{filename: "Go-Mono-Italic.ttf", family: "Go Mono", style: "Italic"},
		{filename: "Go-Mono.ttf", family: "Go Mono", style: "Regular"},
		{filename: "Go-Regular.ttf", family: "Go", style: "Regular"},
		{filename: "Go-Smallcaps-Italic.ttf", family: "Go Smallcaps", style: "Italic"},
		{filename: "Go-Smallcaps.ttf", family: "Go Smallcaps", style: "Regular"},
	}

	test := func(t *testing.T, tc testCase, constructor func(l *Library, path string) (*Face, error)) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to init lib: %s", err)
		}
		defer l.Free()

		f, err := constructor(l, testdata("go", tc.filename))
		if err != nil {
			t.Fatalf("unable to open face: %s", err)
		}
		defer f.Free()

		if f == nil || f.ptr == nil {
			t.Fatalf("got nil face")
		}

		if got := f.FamilyName(); got != tc.family {
			t.Errorf("want family %q, got %q", tc.family, got)
		}
		if got := f.StyleName(); got != tc.style {
			t.Errorf("want style %q, got %q", tc.style, got)
		}
	}

	t.Run("NewFace(nilLib)", func(t *testing.T) {
		var l *Library
		want := errInvalidLib
		if _, err := l.NewFace(nil, 0); err != want {
			t.Errorf("want err: %v, got %v", want, err)
		}
	})

	t.Run("NewFaceFromPath(nilLib)", func(t *testing.T) {
		var l *Library
		want := errInvalidLib
		if _, err := l.NewFaceFromPath("", 0); err != want {
			t.Errorf("want err: %v, got %v", want, err)
		}
	})

	t.Run("NewFace(noData)", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to init lib: %s", err)
		}
		defer l.Free()

		want := ErrUnknownFileFormat
		if _, err := l.NewFace(zeroReader{}, 0); err != want {
			t.Errorf("want err: %v, got %v", want, err)
		}
	})

	t.Run("NewFaceFromPath(fileNotExists)", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to init lib: %s", err)
		}
		defer l.Free()
		want := ErrCannotOpenResource
		if _, err := l.NewFaceFromPath("idontexist.ttf", 0); err != want {
			t.Errorf("want err: %v, got %v", want, err)
		}
	})

	t.Run("NewFaceFromPath(fileEmpty)", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to init lib: %s", err)
		}
		defer l.Free()
		want := ErrUnknownFileFormat
		if _, err := l.NewFaceFromPath(testdata("emptyfile"), 0); err != want {
			t.Errorf("want err: %v, got %v", want, err)
		}
	})

	t.Run("NewFaceFromPath()", func(t *testing.T) {
		for _, tc := range tests {
			test(t, tc, func(l *Library, path string) (*Face, error) {
				return l.NewFaceFromPath(path, 0)
			})
		}
	})

	t.Run("NewFace()", func(t *testing.T) {
		for _, tc := range tests {
			test(t, tc, func(l *Library, path string) (*Face, error) {
				r, err := os.Open(path)
				if err != nil {
					return nil, err
				}
				defer r.Close()
				return l.NewFace(r, 0)
			})
		}
	})
}
