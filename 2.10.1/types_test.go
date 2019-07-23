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

	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %s", err)
	}
	defer l.Free()

	goRegular, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goRegular.Free()

	bungeeLayersReg, err := l.NewFaceFromPath(testdata("bungee", "BungeeLayers-Regular.otf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer bungeeLayersReg.Free()

	tests := []struct {
		name string
		font *Face
		want []CharMap
	}{
		{
			name: "Go Regular",
			font: goRegular,
			want: []CharMap{
				{Format: 4, Language: truetype.MacLangEnglish, Encoding: EncodingUnicode, PlatformID: truetype.PlatformAppleUnicode, EncodingID: truetype.AppleEncodingUnicode2_0, index: 0, valid: true},
				{Format: 6, Language: truetype.MacLangEnglish, Encoding: EncodingAppleRoman, PlatformID: truetype.PlatformMacintosh, EncodingID: truetype.MacEncodingRoman, index: 1, valid: true},
				{Format: 4, Language: 0, Encoding: EncodingUnicode, PlatformID: truetype.PlatformMicrosoft, EncodingID: truetype.MicrosoftEncodingUnicodeCs, index: 2, valid: true},
			},
		},
		{
			name: "Bungee Layers Regular",
			font: bungeeLayersReg,
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
			gotC := tt.font.charmaps()
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
	l, err := NewLibrary()
	if err != nil {
		t.Fatalf("unable to create lib: %s", err)
	}
	defer l.Free()

	goRegular, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer goRegular.Free()

	bungeeLayersReg, err := l.NewFaceFromPath(testdata("bungee", "BungeeLayers-Regular.otf"), 0, 0)
	if err != nil {
		t.Fatalf("unable to open font: %s", err)
	}
	defer bungeeLayersReg.Free()

	charmaps := bungeeLayersReg.CharMaps()

	tests := []struct {
		name    string
		cmap    CharMap
		wantIdx int
		wantOk  bool
	}{
		{
			name:    "invalid",
			cmap:    CharMap{},
			wantIdx: 0,
			wantOk:  false,
		},
		{
			name:    "bungee 0",
			cmap:    charmaps[0],
			wantIdx: 0,
			wantOk:  true,
		},
		{
			name:    "bungee 1",
			cmap:    charmaps[1],
			wantIdx: 1,
			wantOk:  true,
		},
		{
			name:    "bungee 2",
			cmap:    charmaps[2],
			wantIdx: 2,
			wantOk:  true,
		},
		{
			name:    "bungee 3",
			cmap:    charmaps[3],
			wantIdx: 3,
			wantOk:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdx, gotOk := tt.cmap.Index()
			if gotIdx != tt.wantIdx {
				t.Errorf("CharMap.Index() gotIdx = %v, want %v", gotIdx, tt.wantIdx)
			}
			if gotOk != tt.wantOk {
				t.Errorf("CharMap.Index() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
