package freetype2

import (
	"testing"
)

func TestEncoding_String(t *testing.T) {
	tests := []struct {
		name string
		e    Encoding
		want string
	}{
		{name: "None", e: EncodingNone, want: "None"},
		{name: "MsSymbol", e: EncodingMsSymbol, want: "MsSymbol"},
		{name: "Unicode", e: EncodingUnicode, want: "Unicode"},
		{name: "SJIS", e: EncodingSJIS, want: "SJIS"},
		{name: "PRC", e: EncodingPRC, want: "PRC"},
		{name: "Big5", e: EncodingBig5, want: "Big5"},
		{name: "Wansung", e: EncodingWansung, want: "Wansung"},
		{name: "Johab", e: EncodingJohab, want: "Johab"},
		{name: "AdobeStandard", e: EncodingAdobeStandard, want: "AdobeStandard"},
		{name: "AdobeExpert", e: EncodingAdobeExpert, want: "AdobeExpert"},
		{name: "AdobeCustom", e: EncodingAdobeCustom, want: "AdobeCustom"},
		{name: "AdobeLatin1", e: EncodingAdobeLatin1, want: "AdobeLatin1"},
		{name: "AppleRoman", e: EncodingAppleRoman, want: "AppleRoman"},
		{name: "Unknown", e: 9192389, want: "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Encoding.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPixelMode_String(t *testing.T) {
	tests := []struct {
		name string
		p    PixelMode
		want string
	}{
		{name: "Mono", p: PixelModeMono, want: "Mono"},
		{name: "Gray", p: PixelModeGray, want: "Gray"},
		{name: "Gray2", p: PixelModeGray2, want: "Gray2"},
		{name: "Gray4", p: PixelModeGray4, want: "Gray4"},
		{name: "LCD", p: PixelModeLCD, want: "LCD"},
		{name: "LCDV", p: PixelModeLCDV, want: "LCDV"},
		{name: "BGRA", p: PixelModeBGRA, want: "BGRA"},
		{name: "Unknown", p: 912319823, want: "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("PixelMode.String() = %v, want %v", got, tt.want)
			}
		})
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

func TestGlyphFormat_String(t *testing.T) {
	tests := []struct {
		name string
		g    GlyphFormat
		want string
	}{
		{name: "Composite", g: GlyphFormatComposite, want: "Composite"},
		{name: "Bitmap", g: GlyphFormatBitmap, want: "Bitmap"},
		{name: "Outline", g: GlyphFormatOutline, want: "Outline"},
		{name: "Plotter", g: GlyphFormatPlotter, want: "Plotter"},
		{name: "Unknown", g: 912319823, want: "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.String(); got != tt.want {
				t.Errorf("GlyphFormat.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKerningMode_String(t *testing.T) {
	tests := []struct {
		name string
		x    KerningMode
		want string
	}{
		{name: "Default", x: KerningModeDefault, want: "Default"},
		{name: "Unfitted", x: KerningModeUnfitted, want: "Unfitted"},
		{name: "Unscaled", x: KerningModeUnscaled, want: "Unscaled"},
		{name: "Unknown", x: 901929, want: "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("KerningMode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRenderMode_String(t *testing.T) {
	tests := []struct {
		name string
		r    RenderMode
		want string
	}{
		{name: "Normal", r: RenderModeNormal, want: "Normal"},
		{name: "Light", r: RenderModeLight, want: "Light"},
		{name: "Mono", r: RenderModeMono, want: "Mono"},
		{name: "LCD", r: RenderModeLCD, want: "LCD"},
		{name: "LCDV", r: RenderModeLCDV, want: "LCDV"},
		{name: "Unknown", r: 90102, want: "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("RenderMode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
