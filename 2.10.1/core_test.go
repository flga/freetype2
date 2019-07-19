package freetype2

import (
	"os"
	"path/filepath"
	"testing"
)

func testdata(file string) string {
	return filepath.Join("..", "testdata", file)
}

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
		t.Log("testing " + tc.filename)
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to init lib: %s", err)
		}
		defer l.Free()

		f, err := constructor(l, testdata(tc.filename))
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

	if _, err := NewFace(nil, nil, 0); err != errInvalidLib {
		t.Errorf("NewFace(nil), want err: %v, got %v", errInvalidLib, err)
	}
	if _, err := NewFaceFromPath(nil, "", 0); err != errInvalidLib {
		t.Errorf("NewFaceFromPath(nil), want err: %v, got %v", errInvalidLib, err)
	}

	t.Run("from path", func(t *testing.T) {
		for _, tc := range tests {
			test(t, tc, func(l *Library, path string) (*Face, error) {
				return NewFaceFromPath(l, path, 0)
			})
		}
	})

	t.Run("from reader", func(t *testing.T) {
		for _, tc := range tests {
			test(t, tc, func(l *Library, path string) (*Face, error) {
				r, err := os.Open(path)
				if err != nil {
					return nil, err
				}
				defer r.Close()
				return NewFace(l, r, 0)
			})
		}
	})
}

func TestFaceFree(t *testing.T) {
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to init lib: %s", err)
	}
	defer l.Free()

	var called bool
	sentinel := func() { called = true }

	f, err := NewFaceFromPath(l, testdata("Go-Regular.ttf"), 0)
	if err != nil {
		t.Fatalf("unable to open face: %s", err)
	}
	f.dealloc = append(f.dealloc, sentinel)

	if err := f.Free(); err != nil {
		t.Fatalf("unable to free face: %s", err)
	}
	if f.ptr != nil {
		t.Fatalf("Free should set ptr to nil")
	}
	if called != true {
		t.Fatalf("Free should call every function in dealoc")
	}
	if err := f.Free(); err != nil {
		t.Fatalf("Free on an already freed face should be a noop, got: %s", err)
	}
}
