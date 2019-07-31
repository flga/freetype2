package freetype2

import (
	"reflect"
	"testing"

	"github.com/flga/freetype2/2.10.1/truetype"
)

func Test_newCharMap(t *testing.T) {
	if got, want := newCharMap(nil), (CharMap{}); got != want {
		t.Errorf("newCharMap(nil) = %v, want %v", got, want)
	}

	tests := []struct {
		name string
		face func() (testface, error)
		want []CharMap
	}{
		{
			name: "Go Regular",
			face: goRegular,
			want: []CharMap{
				{Format: 4, Language: truetype.MacLangEnglish, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: truetype.AppleEncodingUnicode2_0, index: 0, valid: true},
				{Format: 6, Language: truetype.MacLangEnglish, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: truetype.MacEncodingRoman, index: 1, valid: true},
				{Format: 4, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: truetype.MicrosoftEncodingUnicodeCs, index: 2, valid: true},
			},
		},
		{
			name: "Bungee Layers Regular",
			face: bungeeLayersReg,
			want: []CharMap{
				{Format: 4, Language: truetype.MacLangEnglish, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: truetype.AppleEncodingUnicode2_0, index: 0, valid: true},
				{Format: 6, Language: truetype.MacLangEnglish, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: truetype.MacEncodingRoman, index: 1, valid: true},
				{Format: 4, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: truetype.MicrosoftEncodingUnicodeCs, index: 2, valid: true},
				{Format: -1, Language: 0, Encoding: EncodingAdobeStandard, PlatformID: truetype.PlatformAdobe, EncodingID: truetype.AdobeEncodingStandard, index: 3, valid: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			gotC := face.charmaps()
			if len(gotC) != len(tt.want) {
				t.Fatalf("got %d maps, want %d", len(gotC), len(tt.want))
			}

			got := make([]CharMap, len(gotC))
			for i, c := range gotC {
				got[i] = newCharMap(c)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharMap_Index(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		want, wantOk := 0, false
		var cmap CharMap
		got, gotOk := cmap.Index()
		if got != want || gotOk != wantOk {
			t.Errorf("CharMap.Index() = %v, %v, want %v, %v", got, gotOk, want, wantOk)
		}
	})

	t.Run("notoSansJpReg", func(t *testing.T) {
		face, err := notoSansJpReg()
		if err != nil {
			t.Fatalf("unable to open face: %v", err)
		}

		for i, c := range face.CharMaps() {
			want, wantOk := i, true
			got, gotOk := c.Index()
			if got != want || gotOk != wantOk {
				t.Errorf("CharMap.Index() = %v, %v, want %v, %v", got, gotOk, want, wantOk)
			}
		}
	})
}

func TestSizeRequestType_String(t *testing.T) {
	if got, want := SizeRequestTypeNominal.String(), "Nominal"; got != want {
		t.Errorf("SizeRequestTypeNominal.String() = %v, want %v", got, want)
	}
	if got, want := SizeRequestTypeRealDim.String(), "RealDim"; got != want {
		t.Errorf("SizeRequestTypeRealDim.String() = %v, want %v", got, want)
	}
	if got, want := SizeRequestTypeBBox.String(), "BBox"; got != want {
		t.Errorf("SizeRequestTypeBBox.String() = %v, want %v", got, want)
	}
	if got, want := SizeRequestTypeCell.String(), "Cell"; got != want {
		t.Errorf("SizeRequestTypeCell.String() = %v, want %v", got, want)
	}
	if got, want := SizeRequestTypeScales.String(), "Scales"; got != want {
		t.Errorf("SizeRequestTypeScales.String() = %v, want %v", got, want)
	}
	if got, want := SizeRequestType(8912387).String(), "Unknown"; got != want {
		t.Errorf("8912387.String() = %v, want %v", got, want)
	}
}

func TestOutlineFlag_String(t *testing.T) {
	var x OutlineFlag
	if got, want := x.String(), ""; got != want {
		t.Errorf("OutlineFlag.String(0) = %v, want %v", got, want)
	}

	x = OutlineIgnoreDropouts
	if got, want := x.String(), "IgnoreDropouts"; got != want {
		t.Errorf("OutlineFlag.String(OutlineIgnoreDropouts) = %v, want %v", got, want)
	}

	x = OutlineEvenOddFill | OutlineIncludeStubs
	if got, want := x.String(), "EvenOddFill|IncludeStubs"; got != want {
		t.Errorf("OutlineFlag.String(OutlineEvenOddFill | OutlineIncludeStubs) = %v, want %v", got, want)
	}

	x = OutlineOwner | OutlineIgnoreDropouts | OutlineHighPrecision
	if got, want := x.String(), "Owner|IgnoreDropouts|HighPrecision"; got != want {
		t.Errorf("OutlineFlag.String(OutlineOwner | OutlineIgnoreDropouts | OutlineHighPrecision) = %v, want %v", got, want)
	}

	x = OutlineOwner | OutlineEvenOddFill | OutlineReverseFill |
		OutlineIgnoreDropouts | OutlineSmartDropouts | OutlineIncludeStubs |
		OutlineHighPrecision | OutlineSinglePass
	if got, want := x.String(), "Owner|EvenOddFill|ReverseFill|IgnoreDropouts|SmartDropouts|IncludeStubs|HighPrecision|SinglePass"; got != want {
		t.Errorf("OutlineFlag.String(OutlineOwner | OutlineEvenOddFill | OutlineReverseFill | OutlineIgnoreDropouts | OutlineSmartDropouts | OutlineIncludeStubs | OutlineHighPrecision | OutlineSinglePass) = %v, want %v", got, want)
	}
}
