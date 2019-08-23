package freetype2

import (
	"testing"
)

func TestFace_SfntNameCount(t *testing.T) {
	tests := []struct {
		name string
		face func() (testface, error)
		want int
	}{
		{name: "nilFace", face: nilFace, want: 0},
		{name: "bungeeColorWin", face: bungeeColorWin, want: 42},
		{name: "bungeeColorMac", face: bungeeColorMac, want: 42},
		{name: "goRegular", face: goRegular, want: 25},
		{name: "notoSansJpReg", face: notoSansJpReg, want: 16},
		{name: "nimbusMono", face: nimbusMono, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			if got := face.SfntNameCount(); got != tt.want {
				t.Errorf("Face.SfntNameCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFace_SfntName(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		idx     int
		want    SfntName
		wantErr error
	}{
		{
			name:    "nilFace",
			face:    nilFace,
			idx:     0,
			want:    SfntName{},
			wantErr: ErrInvalidArgument,
		},

		{name: "out of bounds", face: goRegular, idx: -1, want: SfntName{}, wantErr: ErrInvalidArgument},
		{name: "out of bounds", face: goRegular, idx: 25, want: SfntName{}, wantErr: ErrInvalidArgument},

		{name: "bungeeColorWin-0", face: bungeeColorWin, idx: 0, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 0, Name: "Copyright 2008 The Bungee Project Authors (david@djr.com)"}, wantErr: nil},
		{name: "bungeeColorWin-1", face: bungeeColorWin, idx: 1, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 1, Name: "Bungee Color Regular"}, wantErr: nil},
		{name: "bungeeColorWin-2", face: bungeeColorWin, idx: 2, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 2, Name: "Regular"}, wantErr: nil},
		{name: "bungeeColorWin-3", face: bungeeColorWin, idx: 3, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 3, Name: "1.000;djr ;BungeeColor-Regular"}, wantErr: nil},
		{name: "bungeeColorWin-4", face: bungeeColorWin, idx: 4, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 4, Name: "Bungee Color Regular"}, wantErr: nil},
		{name: "bungeeColorWin-5", face: bungeeColorWin, idx: 5, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 5, Name: "Version 1.000;PS 1.0;hotconv 1.0.72;makeotf.lib2.5.5900"}, wantErr: nil},
		{name: "bungeeColorWin-6", face: bungeeColorWin, idx: 6, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 6, Name: "BungeeColor-Regular"}, wantErr: nil},
		{name: "bungeeColorWin-7", face: bungeeColorWin, idx: 7, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 7, Name: "Bungee is a trademark of The Font Bureau."}, wantErr: nil},
		{name: "bungeeColorWin-8", face: bungeeColorWin, idx: 8, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 8, Name: "David Jonathan Ross"}, wantErr: nil},
		{name: "bungeeColorWin-9", face: bungeeColorWin, idx: 9, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 9, Name: "David Jonathan Ross"}, wantErr: nil},
		{name: "bungeeColorWin-10", face: bungeeColorWin, idx: 10, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 11, Name: "http://www.djr.com"}, wantErr: nil},
		{name: "bungeeColorWin-11", face: bungeeColorWin, idx: 11, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 12, Name: "http://www.djr.com"}, wantErr: nil},
		{name: "bungeeColorWin-12", face: bungeeColorWin, idx: 12, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 13, Name: "This Font Software is licensed under the SIL Open Font License, Version 1.1. This license is available with a FAQ at: http://scripts.sil.org/OFL"}, wantErr: nil},
		{name: "bungeeColorWin-13", face: bungeeColorWin, idx: 13, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 14, Name: "http://scripts.sil.org/OFL"}, wantErr: nil},
		{name: "bungeeColorWin-14", face: bungeeColorWin, idx: 14, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 16, Name: "Bungee Color"}, wantErr: nil},
		{name: "bungeeColorWin-15", face: bungeeColorWin, idx: 15, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 256, Name: "Round forms"}, wantErr: nil},
		{name: "bungeeColorWin-16", face: bungeeColorWin, idx: 16, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 257, Name: "Round E"}, wantErr: nil},
		{name: "bungeeColorWin-17", face: bungeeColorWin, idx: 17, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 258, Name: "Sans-serif I"}, wantErr: nil},
		{name: "bungeeColorWin-18", face: bungeeColorWin, idx: 18, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 259, Name: "Sans-serif L"}, wantErr: nil},
		{name: "bungeeColorWin-19", face: bungeeColorWin, idx: 19, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 260, Name: "Alternate ampersand"}, wantErr: nil},
		{name: "bungeeColorWin-20", face: bungeeColorWin, idx: 20, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 261, Name: "Small quotes"}, wantErr: nil},
		{name: "bungeeColorWin-21", face: bungeeColorWin, idx: 21, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 0, Name: "Copyright 2008 The Bungee Project Authors (david@djr.com)"}, wantErr: nil},
		{name: "bungeeColorWin-22", face: bungeeColorWin, idx: 22, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 1, Name: "Bungee Color Regular"}, wantErr: nil},
		{name: "bungeeColorWin-23", face: bungeeColorWin, idx: 23, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 2, Name: "Regular"}, wantErr: nil},
		{name: "bungeeColorWin-24", face: bungeeColorWin, idx: 24, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 3, Name: "1.000;djr ;BungeeColor-Regular"}, wantErr: nil},
		{name: "bungeeColorWin-25", face: bungeeColorWin, idx: 25, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 4, Name: "Bungee Color Regular Regular"}, wantErr: nil},
		{name: "bungeeColorWin-26", face: bungeeColorWin, idx: 26, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 5, Name: "Version 1.000;PS 1.0;hotconv 1.0.72;makeotf.lib2.5.5900"}, wantErr: nil},
		{name: "bungeeColorWin-27", face: bungeeColorWin, idx: 27, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 6, Name: "BungeeColor-Regular"}, wantErr: nil},
		{name: "bungeeColorWin-28", face: bungeeColorWin, idx: 28, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 7, Name: "Bungee is a trademark of The Font Bureau."}, wantErr: nil},
		{name: "bungeeColorWin-29", face: bungeeColorWin, idx: 29, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 8, Name: "David Jonathan Ross"}, wantErr: nil},
		{name: "bungeeColorWin-30", face: bungeeColorWin, idx: 30, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 9, Name: "David Jonathan Ross"}, wantErr: nil},
		{name: "bungeeColorWin-31", face: bungeeColorWin, idx: 31, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 11, Name: "http://www.djr.com"}, wantErr: nil},
		{name: "bungeeColorWin-32", face: bungeeColorWin, idx: 32, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 12, Name: "http://www.djr.com"}, wantErr: nil},
		{name: "bungeeColorWin-33", face: bungeeColorWin, idx: 33, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 13, Name: "This Font Software is licensed under the SIL Open Font License, Version 1.1. This license is available with a FAQ at: http://scripts.sil.org/OFL"}, wantErr: nil},
		{name: "bungeeColorWin-34", face: bungeeColorWin, idx: 34, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 14, Name: "http://scripts.sil.org/OFL"}, wantErr: nil},
		{name: "bungeeColorWin-35", face: bungeeColorWin, idx: 35, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 16, Name: "Bungee Color"}, wantErr: nil},
		{name: "bungeeColorWin-36", face: bungeeColorWin, idx: 36, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 256, Name: "Round forms"}, wantErr: nil},
		{name: "bungeeColorWin-37", face: bungeeColorWin, idx: 37, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 257, Name: "Round E"}, wantErr: nil},
		{name: "bungeeColorWin-38", face: bungeeColorWin, idx: 38, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 258, Name: "Sans-serif I"}, wantErr: nil},
		{name: "bungeeColorWin-39", face: bungeeColorWin, idx: 39, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 259, Name: "Sans-serif L"}, wantErr: nil},
		{name: "bungeeColorWin-40", face: bungeeColorWin, idx: 40, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 260, Name: "Alternate ampersand"}, wantErr: nil},
		{name: "bungeeColorWin-41", face: bungeeColorWin, idx: 41, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 261, Name: "Small quotes"}, wantErr: nil},

		{name: "bungeeColorMac-0", face: bungeeColorMac, idx: 0, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 0, Name: "Copyright 2008 The Bungee Project Authors (david@djr.com)"}, wantErr: nil},
		{name: "bungeeColorMac-1", face: bungeeColorMac, idx: 1, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 1, Name: "Bungee Color Regular"}, wantErr: nil},
		{name: "bungeeColorMac-2", face: bungeeColorMac, idx: 2, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 2, Name: "Regular"}, wantErr: nil},
		{name: "bungeeColorMac-3", face: bungeeColorMac, idx: 3, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 3, Name: "1.000;djr ;BungeeColor-Regular"}, wantErr: nil},
		{name: "bungeeColorMac-4", face: bungeeColorMac, idx: 4, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 4, Name: "Bungee Color Regular"}, wantErr: nil},
		{name: "bungeeColorMac-5", face: bungeeColorMac, idx: 5, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 5, Name: "Version 1.000;PS 1.0;hotconv 1.0.72;makeotf.lib2.5.5900"}, wantErr: nil},
		{name: "bungeeColorMac-6", face: bungeeColorMac, idx: 6, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 6, Name: "BungeeColor-Regular"}, wantErr: nil},
		{name: "bungeeColorMac-7", face: bungeeColorMac, idx: 7, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 7, Name: "Bungee is a trademark of The Font Bureau."}, wantErr: nil},
		{name: "bungeeColorMac-8", face: bungeeColorMac, idx: 8, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 8, Name: "David Jonathan Ross"}, wantErr: nil},
		{name: "bungeeColorMac-9", face: bungeeColorMac, idx: 9, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 9, Name: "David Jonathan Ross"}, wantErr: nil},
		{name: "bungeeColorMac-10", face: bungeeColorMac, idx: 10, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 11, Name: "http://www.djr.com"}, wantErr: nil},
		{name: "bungeeColorMac-11", face: bungeeColorMac, idx: 11, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 12, Name: "http://www.djr.com"}, wantErr: nil},
		{name: "bungeeColorMac-12", face: bungeeColorMac, idx: 12, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 13, Name: "This Font Software is licensed under the SIL Open Font License, Version 1.1. This license is available with a FAQ at: http://scripts.sil.org/OFL"}, wantErr: nil},
		{name: "bungeeColorMac-13", face: bungeeColorMac, idx: 13, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 14, Name: "http://scripts.sil.org/OFL"}, wantErr: nil},
		{name: "bungeeColorMac-14", face: bungeeColorMac, idx: 14, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 16, Name: "Bungee Color"}, wantErr: nil},
		{name: "bungeeColorMac-15", face: bungeeColorMac, idx: 15, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 256, Name: "Round forms"}, wantErr: nil},
		{name: "bungeeColorMac-16", face: bungeeColorMac, idx: 16, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 257, Name: "Round E"}, wantErr: nil},
		{name: "bungeeColorMac-17", face: bungeeColorMac, idx: 17, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 258, Name: "Sans-serif I"}, wantErr: nil},
		{name: "bungeeColorMac-18", face: bungeeColorMac, idx: 18, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 259, Name: "Sans-serif L"}, wantErr: nil},
		{name: "bungeeColorMac-19", face: bungeeColorMac, idx: 19, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 260, Name: "Alternate ampersand"}, wantErr: nil},
		{name: "bungeeColorMac-20", face: bungeeColorMac, idx: 20, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 261, Name: "Small quotes"}, wantErr: nil},
		{name: "bungeeColorMac-21", face: bungeeColorMac, idx: 21, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 0, Name: "Copyright 2008 The Bungee Project Authors (david@djr.com)"}, wantErr: nil},
		{name: "bungeeColorMac-22", face: bungeeColorMac, idx: 22, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 1, Name: "Bungee Color Regular"}, wantErr: nil},
		{name: "bungeeColorMac-23", face: bungeeColorMac, idx: 23, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 2, Name: "Regular"}, wantErr: nil},
		{name: "bungeeColorMac-24", face: bungeeColorMac, idx: 24, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 3, Name: "1.000;djr ;BungeeColor-Regular"}, wantErr: nil},
		{name: "bungeeColorMac-25", face: bungeeColorMac, idx: 25, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 4, Name: "Bungee Color Regular Regular"}, wantErr: nil},
		{name: "bungeeColorMac-26", face: bungeeColorMac, idx: 26, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 5, Name: "Version 1.000;PS 1.0;hotconv 1.0.72;makeotf.lib2.5.5900"}, wantErr: nil},
		{name: "bungeeColorMac-27", face: bungeeColorMac, idx: 27, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 6, Name: "BungeeColor-Regular"}, wantErr: nil},
		{name: "bungeeColorMac-28", face: bungeeColorMac, idx: 28, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 7, Name: "Bungee is a trademark of The Font Bureau."}, wantErr: nil},
		{name: "bungeeColorMac-29", face: bungeeColorMac, idx: 29, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 8, Name: "David Jonathan Ross"}, wantErr: nil},
		{name: "bungeeColorMac-30", face: bungeeColorMac, idx: 30, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 9, Name: "David Jonathan Ross"}, wantErr: nil},
		{name: "bungeeColorMac-31", face: bungeeColorMac, idx: 31, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 11, Name: "http://www.djr.com"}, wantErr: nil},
		{name: "bungeeColorMac-32", face: bungeeColorMac, idx: 32, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 12, Name: "http://www.djr.com"}, wantErr: nil},
		{name: "bungeeColorMac-33", face: bungeeColorMac, idx: 33, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 13, Name: "This Font Software is licensed under the SIL Open Font License, Version 1.1. This license is available with a FAQ at: http://scripts.sil.org/OFL"}, wantErr: nil},
		{name: "bungeeColorMac-34", face: bungeeColorMac, idx: 34, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 14, Name: "http://scripts.sil.org/OFL"}, wantErr: nil},
		{name: "bungeeColorMac-35", face: bungeeColorMac, idx: 35, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 16, Name: "Bungee Color"}, wantErr: nil},
		{name: "bungeeColorMac-36", face: bungeeColorMac, idx: 36, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 256, Name: "Round forms"}, wantErr: nil},
		{name: "bungeeColorMac-37", face: bungeeColorMac, idx: 37, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 257, Name: "Round E"}, wantErr: nil},
		{name: "bungeeColorMac-38", face: bungeeColorMac, idx: 38, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 258, Name: "Sans-serif I"}, wantErr: nil},
		{name: "bungeeColorMac-39", face: bungeeColorMac, idx: 39, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 259, Name: "Sans-serif L"}, wantErr: nil},
		{name: "bungeeColorMac-40", face: bungeeColorMac, idx: 40, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 260, Name: "Alternate ampersand"}, wantErr: nil},
		{name: "bungeeColorMac-41", face: bungeeColorMac, idx: 41, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 261, Name: "Small quotes"}, wantErr: nil},

		{name: "goRegular-0", face: goRegular, idx: 0, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 0, Name: "Copyright (c) 2016 by Bigelow & Holmes Inc.. All rights reserved."}, wantErr: nil},
		{name: "goRegular-1", face: goRegular, idx: 1, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 1, Name: "Go"}, wantErr: nil},
		{name: "goRegular-2", face: goRegular, idx: 2, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 2, Name: "Regular"}, wantErr: nil},
		{name: "goRegular-3", face: goRegular, idx: 3, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 3, Name: "Bigelow&HolmesInc.: Go Regular: 2016"}, wantErr: nil},
		{name: "goRegular-4", face: goRegular, idx: 4, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 4, Name: "Go Regular"}, wantErr: nil},
		{name: "goRegular-5", face: goRegular, idx: 5, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 5, Name: "Version 2.008; ttfautohint (v1.6)"}, wantErr: nil},
		{name: "goRegular-6", face: goRegular, idx: 6, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 6, Name: "GoRegular"}, wantErr: nil},
		{name: "goRegular-7", face: goRegular, idx: 7, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 8, Name: "Bigelow & Holmes Inc."}, wantErr: nil},
		{name: "goRegular-8", face: goRegular, idx: 8, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 9, Name: "Kris Holmes and Charles Bigelow"}, wantErr: nil},
		{name: "goRegular-9", face: goRegular, idx: 9, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 10, Name: "Go is a humanistic sans-serif font for the Go language. Its x-height, stem weight, and distinctive forms of zero, capital O, lowercase l, figure one, and capital I follow the DIN 1450 font legibility standard. Go's WGL character set includes Unicode Latin, Greek and Cyrillic alphabets plus symbols and graphical elements."}, wantErr: nil},
		{name: "goRegular-10", face: goRegular, idx: 10, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 12, Name: "lucidafonts.com"}, wantErr: nil},
		{
			name: "goRegular-11",
			face: goRegular,
			idx:  11, want: SfntName{
				PlatformID: 1,
				EncodingID: 0,
				LanguageID: 0,
				NameID:     13,
				Name: `Copyright (c) 2016 Bigelow & Holmes Inc.. All rights reserved.

Distribution of this font is governed by the following license. If you do not agree to this license, including the disclaimer, do not distribute or modify this font.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

   * Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

   * Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

   * Neither the name of Google Inc. nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

DISCLAIMER: THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.`,
			},
			wantErr: nil,
		},
		{name: "goRegular-12", face: goRegular, idx: 12, want: SfntName{PlatformID: 1, EncodingID: 0, LanguageID: 0, NameID: 18, Name: "Go Regular"}, wantErr: nil},
		{name: "goRegular-13", face: goRegular, idx: 13, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 0, Name: "Copyright (c) 2016 by Bigelow & Holmes Inc.. All rights reserved."}, wantErr: nil},
		{name: "goRegular-14", face: goRegular, idx: 14, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 1, Name: "Go"}, wantErr: nil},
		{name: "goRegular-15", face: goRegular, idx: 15, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 2, Name: "Regular"}, wantErr: nil},
		{name: "goRegular-16", face: goRegular, idx: 16, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 3, Name: "Bigelow&HolmesInc.: Go Regular: 2016"}, wantErr: nil},
		{name: "goRegular-17", face: goRegular, idx: 17, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 4, Name: "Go Regular"}, wantErr: nil},
		{name: "goRegular-18", face: goRegular, idx: 18, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 5, Name: "Version 2.008; ttfautohint (v1.6)"}, wantErr: nil},
		{name: "goRegular-19", face: goRegular, idx: 19, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 6, Name: "GoRegular"}, wantErr: nil},
		{name: "goRegular-20", face: goRegular, idx: 20, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 8, Name: "Bigelow & Holmes Inc."}, wantErr: nil},
		{name: "goRegular-21", face: goRegular, idx: 21, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 9, Name: "Kris Holmes and Charles Bigelow"}, wantErr: nil},
		{name: "goRegular-22", face: goRegular, idx: 22, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 10, Name: "Go is a humanistic sans-serif font for the Go language. Its x-height, stem weight, and distinctive forms of zero, capital O, lowercase l, figure one, and capital I follow the DIN 1450 font legibility standard. Go's WGL character set includes Unicode Latin, Greek and Cyrillic alphabets plus symbols and graphical elements."}, wantErr: nil},
		{name: "goRegular-23", face: goRegular, idx: 23, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 12, Name: "lucidafonts.com"}, wantErr: nil},
		{
			name: "goRegular-24",
			face: goRegular,
			idx:  24,
			want: SfntName{
				PlatformID: 3,
				EncodingID: 1,
				LanguageID: 1033,
				NameID:     13,
				Name: `Copyright (c) 2016 Bigelow & Holmes Inc.. All rights reserved.

Distribution of this font is governed by the following license. If you do not agree to this license, including the disclaimer, do not distribute or modify this font.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

   * Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

   * Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

   * Neither the name of Google Inc. nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

DISCLAIMER: THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.`,
			},
			wantErr: nil,
		},

		{name: "notoSansJpReg-0", face: notoSansJpReg, idx: 0, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 0, Name: "Copyright © 2014, 2015 Adobe Systems Incorporated (http://www.adobe.com/)."}, wantErr: nil},
		{name: "notoSansJpReg-1", face: notoSansJpReg, idx: 1, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 1, Name: "Noto Sans JP Regular"}, wantErr: nil},
		{name: "notoSansJpReg-2", face: notoSansJpReg, idx: 2, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 2, Name: "Regular"}, wantErr: nil},
		{name: "notoSansJpReg-3", face: notoSansJpReg, idx: 3, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 3, Name: "1.004;GOOG;NotoSansJP-Regular;ADOBE"}, wantErr: nil},
		{name: "notoSansJpReg-4", face: notoSansJpReg, idx: 4, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 4, Name: "Noto Sans JP Regular"}, wantErr: nil},
		{name: "notoSansJpReg-5", face: notoSansJpReg, idx: 5, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 5, Name: "Version 1.004;PS 1.004;hotconv 1.0.82;makeotf.lib2.5.63406"}, wantErr: nil},
		{name: "notoSansJpReg-6", face: notoSansJpReg, idx: 6, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 6, Name: "NotoSansJP-Regular"}, wantErr: nil},
		{name: "notoSansJpReg-7", face: notoSansJpReg, idx: 7, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 7, Name: "Noto is a trademark of Google Inc."}, wantErr: nil},
		{name: "notoSansJpReg-8", face: notoSansJpReg, idx: 8, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 8, Name: "Adobe Systems Incorporated"}, wantErr: nil},
		{name: "notoSansJpReg-9", face: notoSansJpReg, idx: 9, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 9, Name: "Ryoko NISHIZUKA 西塚涼子 (kana & ideographs); Paul D. Hunt (Latin, Greek & Cyrillic); Wenlong ZHANG 张文龙 (bopomofo); Sandoll Communication 산돌커뮤니케이션, Soo-young JANG 장수영 & Joo-yeon KANG 강주연 (hangul elements, letters & syllables)"}, wantErr: nil},
		{name: "notoSansJpReg-10", face: notoSansJpReg, idx: 10, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 10, Name: "Dr. Ken Lunde (project architect, glyph set definition & overall production); Masataka HATTORI 服部正貴 (production & ideograph elements)"}, wantErr: nil},
		{name: "notoSansJpReg-11", face: notoSansJpReg, idx: 11, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 11, Name: "http://www.google.com/get/noto/"}, wantErr: nil},
		{name: "notoSansJpReg-12", face: notoSansJpReg, idx: 12, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 12, Name: "http://www.adobe.com/type/"}, wantErr: nil},
		{name: "notoSansJpReg-13", face: notoSansJpReg, idx: 13, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 13, Name: "This Font Software is licensed under the SIL Open Font License, Version 1.1. This Font Software is distributed on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the SIL Open Font License for the specific language, permissions and limitations governing your use of this Font Software."}, wantErr: nil},
		{name: "notoSansJpReg-14", face: notoSansJpReg, idx: 14, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 14, Name: "http://scripts.sil.org/OFL"}, wantErr: nil},
		{name: "notoSansJpReg-15", face: notoSansJpReg, idx: 15, want: SfntName{PlatformID: 3, EncodingID: 1, LanguageID: 1033, NameID: 16, Name: "Noto Sans JP"}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %v", err)
			}
			defer face.Free()

			got, err := face.SfntName(tt.idx)
			if err != tt.wantErr {
				t.Errorf("Face.SfntName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.SfntName() = %v", diff)
			}
		})
	}
}

func TestFace_SfntLangTag(t *testing.T) {
	t.Skip("need a font with a format 1 name table")
	// tests := []struct {
	// 	name    string
	// 	face    func() (testface, error)
	// 	id      truetype.LanguageID
	// 	want    SfntLangTag
	// 	wantErr error
	// }{
	// 	{name: "nilFace", face: nilFace, id: 0, want: "", wantErr: ErrInvalidArgument},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		face, err := tt.face()
	// 		if err != nil {
	// 			t.Fatalf("unable to load face: %v", err)
	// 		}
	// 		defer face.Free()

	// 		got, err := face.SfntLangTag(tt.id)
	// 		if err != tt.wantErr {
	// 			t.Errorf("Face.SfntLangTag() error = %v, wantErr %v", err, tt.wantErr)
	// 		}
	// 		if got != tt.want {
	// 			t.Errorf("Face.SfntLangTag() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
