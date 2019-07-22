package freetype2

import (
	"testing"
)

func TestEncoding_String(t *testing.T) {
	if got, want := EncodingNone.String(), "None"; got != want {
		t.Errorf("EncodingNone.String() = %s, want %s", got, want)
	}
	if got, want := EncodingMsSymbol.String(), "MsSymbol"; got != want {
		t.Errorf("EncodingMsSymbol.String() = %s, want %s", got, want)
	}
	if got, want := EncodingUnicode.String(), "Unicode"; got != want {
		t.Errorf("EncodingUnicode.String() = %s, want %s", got, want)
	}
	if got, want := EncodingSJIS.String(), "SJIS"; got != want {
		t.Errorf("EncodingSJIS.String() = %s, want %s", got, want)
	}
	if got, want := EncodingPRC.String(), "PRC"; got != want {
		t.Errorf("EncodingPRC.String() = %s, want %s", got, want)
	}
	if got, want := EncodingBig5.String(), "Big5"; got != want {
		t.Errorf("EncodingBig5.String() = %s, want %s", got, want)
	}
	if got, want := EncodingWansung.String(), "Wansung"; got != want {
		t.Errorf("EncodingWansung.String() = %s, want %s", got, want)
	}
	if got, want := EncodingJohab.String(), "Johab"; got != want {
		t.Errorf("EncodingJohab.String() = %s, want %s", got, want)
	}
	if got, want := EncodingAdobeStandard.String(), "AdobeStandard"; got != want {
		t.Errorf("EncodingAdobeStandard.String() = %s, want %s", got, want)
	}
	if got, want := EncodingAdobeExpert.String(), "AdobeExpert"; got != want {
		t.Errorf("EncodingAdobeExpert.String() = %s, want %s", got, want)
	}
	if got, want := EncodingAdobeCustom.String(), "AdobeCustom"; got != want {
		t.Errorf("EncodingAdobeCustom.String() = %s, want %s", got, want)
	}
	if got, want := EncodingAdobeLatin1.String(), "AdobeLatin1"; got != want {
		t.Errorf("EncodingAdobeLatin1.String() = %s, want %s", got, want)
	}
	if got, want := EncodingAppleRoman.String(), "AppleRoman"; got != want {
		t.Errorf("EncodingAppleRoman.String() = %s, want %s", got, want)
	}
	if got, want := Encoding(912319823).String(), "Unknown"; got != want {
		t.Errorf("912319823.String() = %s, want %s", got, want)
	}
}

func TestPixelMode_BitsPerPixel(t *testing.T) {
	tests := []struct {
		p    PixelMode
		want int
	}{
		{PixelModeMono, 1},
		{PixelModeGray, 8},
		{PixelModeGray2, 2},
		{PixelModeGray4, 4},
		{PixelModeLCD, 8},
		{PixelModeLCDV, 8},
		{PixelModeBGRA, 32},
		{912319823, 0},
	}
	for _, tt := range tests {
		t.Run(PixelModeBGRA.String(), func(t *testing.T) {
			if got := tt.p.BitsPerPixel(); got != tt.want {
				t.Errorf("PixelMode.BitsPerPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPixelMode_String(t *testing.T) {
	if got, want := PixelModeMono.String(), "Mono"; got != want {
		t.Errorf("PixelModeMono.String() = %s, want %s", got, want)
	}
	if got, want := PixelModeGray.String(), "Gray"; got != want {
		t.Errorf("PixelModeGray.String() = %s, want %s", got, want)
	}
	if got, want := PixelModeGray2.String(), "Gray2"; got != want {
		t.Errorf("PixelModeGray2.String() = %s, want %s", got, want)
	}
	if got, want := PixelModeGray4.String(), "Gray4"; got != want {
		t.Errorf("PixelModeGray4.String() = %s, want %s", got, want)
	}
	if got, want := PixelModeLCD.String(), "LCD"; got != want {
		t.Errorf("PixelModeLCD.String() = %s, want %s", got, want)
	}
	if got, want := PixelModeLCDV.String(), "LCDV"; got != want {
		t.Errorf("PixelModeLCDV.String() = %s, want %s", got, want)
	}
	if got, want := PixelModeBGRA.String(), "BGRA"; got != want {
		t.Errorf("PixelModeBGRA.String() = %s, want %s", got, want)
	}
	if got, want := PixelMode(912319823).String(), "Unknown"; got != want {
		t.Errorf("912319823.String() = %s, want %s", got, want)
	}
}

func TestGlyphFormat_String(t *testing.T) {
	if got, want := GlyphFormatComposite.String(), "Composite"; got != want {
		t.Errorf("GlyphFormatComposite.String() = %s, want %s", got, want)
	}
	if got, want := GlyphFormatBitmap.String(), "Bitmap"; got != want {
		t.Errorf("GlyphFormatBitmap.String() = %s, want %s", got, want)
	}
	if got, want := GlyphFormatOutline.String(), "Outline"; got != want {
		t.Errorf("GlyphFormatOutline.String() = %s, want %s", got, want)
	}
	if got, want := GlyphFormatPlotter.String(), "Plotter"; got != want {
		t.Errorf("GlyphFormatPlotter.String() = %s, want %s", got, want)
	}
	if got, want := GlyphFormat(912319823).String(), "Unknown"; got != want {
		t.Errorf("912319823.String() = %s, want %s", got, want)
	}
}
