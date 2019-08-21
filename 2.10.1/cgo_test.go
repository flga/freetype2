package freetype2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sliceFromZeroTerminatedUint32(t *testing.T) {
	tests := []struct {
		n    int
		want []rune
	}{
		{0, nil},
		{1, []rune{1}},
		{2, []rune{1, 2}},
		{3, []rune{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := sliceFromZeroTerminatedUint32(testMakeList(tt.n)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceFromZeroTerminatedUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}
