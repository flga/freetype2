package freetype2

import (
	"testing"
)

func TestGetErr(t *testing.T) {
	var want error

	want = ErrCannotOpenResource
	if got := testErrCannotOpenResource(); got != want {
		t.Errorf("want err: %v, got %v", want, got)
	}

	want = ErrUnknownError
	if got := testUnmappedErr(); got != want {
		t.Errorf("want err: %v, got %v", want, got)
	}
}
