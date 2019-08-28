package freetype2

import (
	"testing"

	"github.com/flga/freetype2/fixed"
)

func TestFace_MultiMaster(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		want    MultiMaster
		wantErr error
	}{
		{
			name:    "nilFace",
			face:    nilFace,
			want:    MultiMaster{},
			wantErr: ErrInvalidFaceHandle,
		},
		{
			name:    "goRegular",
			face:    goRegular,
			want:    MultiMaster{},
			wantErr: ErrInvalidArgument,
		},
		{face: faceFromPath("variable/blinker/Blinker_variable.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/Decovar/DecovarAlpha-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/FiraCode/FiraCode-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/movement/MovementV.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/league-spartan/LeagueSpartanVariable.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/Hepta-Slab/HeptaSlab-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/nunito/NunitoVFBeta.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/KayakVF/KayakVF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/Amstelvar/Amstelvar-Roman-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/Amstelvar/Amstelvar-Roman-VF-APPS.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/Amstelvar/Hidden-Axis-Amstel.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/RibbonVF/RibbonVF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/gnomon/gnomon-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/PT Root UI VF/PT Root UI_VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/soulcraft/Soulcraft.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/source-code-pro/SourceCodeVariable-Roman.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/barlow/BarlowGX.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/CrimsonPro/CrimsonPro-Roman-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/waba-border/WabaBorderGX.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/Graduate-Variable-Font/GRADUATE.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/zycon/Zycon.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/work-sans/WorkSans-Roman-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/width-and-vertical-width-vf/WidthAndVWidthVF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/width-and-vertical-width-vf/WidthAndVWidthVF.otf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/markazitext/MarkaziText-VF.ttf"), wantErr: ErrInvalidArgument},
		{
			face: faceFromPath("variable/impossible/ImposMM.pfb"),
			want: MultiMaster{
				NumAxis:    3,
				NumDesigns: 8,
				Axis: []MMAxis{
					{Name: "Width", Max: 1000, Min: 0},
					{Name: "OpticalSize", Max: 1000, Min: 0},
					{Name: "Serif", Max: 1000, Min: 0},
				},
			},
			wantErr: nil,
		},
		{face: faceFromPath("variable/Libre-Franklin/LibreFranklinGX-Romans-v4015.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/tiny/TINY5x3GX.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/IBM-Plex-Sans-Variable/IBMPlexSansVar-Roman.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/IBM-Plex-Sans-Variable/IBMPlexSansVar-Italic.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/cabin/Cabin_V.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/varfonts-ofl/ZinzinVF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/leaguemono/LeagueMonoVariable.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/changa-vf/Changa-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/Lora-Cyrillic/Lora-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/mutatorSans/MutatorSans.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/sudo-font/SudoVariable.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/Gingham/Gingham.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/secuela-variable/Secuela-Regular-v_1_787-TTF-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoHK.otf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoJP.otf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoKR.otf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoTW.otf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoCN.otf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoMO.otf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/titillium-web-vf/TitilliumWeb-Roman-VF.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/VotoSerifGX-OFL/VotoSerifGX.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/iA Writer Mono/iAWriterMonoV.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/selawik/Selawik-variable.ttf"), wantErr: ErrInvalidArgument},
		{face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), wantErr: ErrInvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %s", err)
			}
			defer face.Free()

			got, err := face.MultiMaster()
			if err != tt.wantErr {
				t.Errorf("Face.MultiMaster() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.MultiMaster() = %v", diff)
			}
		})
	}
}

func TestFace_MMVar(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		want    *MMVar
		wantErr error
	}{
		{name: "nilFaces", face: nilFace, want: nil, wantErr: ErrInvalidFaceHandle},
		{name: "goRegular", face: goRegular, want: nil, wantErr: ErrInvalidArgument},
		{
			name: "Blinker_variable.ttf",
			face: faceFromPath("variable/blinker/Blinker_variable.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 8,
				Axis: []VarAxis{
					{Name: "Weight", Min: 1310720, Def: 5373952, Max: 14417920, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{1310720}, Strid: 257, Psid: 258},
					{Coords: []fixed.Int16_16{2359296}, Strid: 259, Psid: 260},
					{Coords: []fixed.Int16_16{4063232}, Strid: 261, Psid: 262},
					{Coords: []fixed.Int16_16{5373952}, Strid: 263, Psid: 264},
					{Coords: []fixed.Int16_16{7602176}, Strid: 265, Psid: 266},
					{Coords: []fixed.Int16_16{9830400}, Strid: 267, Psid: 268},
					{Coords: []fixed.Int16_16{11796480}, Strid: 269, Psid: 270},
					{Coords: []fixed.Int16_16{14417920}, Strid: 271, Psid: 272},
				},
			},
		},
		{
			name: "DecovarAlpha-VF.ttf",
			face: faceFromPath("variable/Decovar/DecovarAlpha-VF.ttf"),
			want: &MMVar{
				NumAxis:        15,
				NumDesigns:     0,
				NumNamedstyles: 17,
				Axis: []VarAxis{
					{Name: "BLDA", Min: 0, Def: 0, Max: 65536000, Tag: 0x424c4441, Strid: 256, Flags: 0x00000000},
					{Name: "TRMD", Min: 0, Def: 0, Max: 65536000, Tag: 0x54524d44, Strid: 257, Flags: 0x00000000},
					{Name: "TRMC", Min: 0, Def: 0, Max: 65536000, Tag: 0x54524d43, Strid: 258, Flags: 0x00000000},
					{Name: "SKLD", Min: 0, Def: 0, Max: 65536000, Tag: 0x534b4c44, Strid: 259, Flags: 0x00000000},
					{Name: "TRML", Min: 0, Def: 0, Max: 65536000, Tag: 0x54524d4c, Strid: 260, Flags: 0x00000000},
					{Name: "SKLA", Min: 0, Def: 0, Max: 65536000, Tag: 0x534b4c41, Strid: 261, Flags: 0x00000000},
					{Name: "TRMF", Min: 0, Def: 0, Max: 65536000, Tag: 0x54524d46, Strid: 262, Flags: 0x00000000},
					{Name: "TRMK", Min: 0, Def: 0, Max: 65536000, Tag: 0x54524d4b, Strid: 263, Flags: 0x00000000},
					{Name: "BLDB", Min: 0, Def: 0, Max: 65536000, Tag: 0x424c4442, Strid: 264, Flags: 0x00000000},
					{Name: "WMX2", Min: 0, Def: 0, Max: 65536000, Tag: 0x574d5832, Strid: 265, Flags: 0x00000000},
					{Name: "TRMB", Min: 0, Def: 0, Max: 65536000, Tag: 0x54524d42, Strid: 266, Flags: 0x00000000},
					{Name: "TRMA", Min: 0, Def: 0, Max: 65536000, Tag: 0x54524d41, Strid: 267, Flags: 0x00000000},
					{Name: "SKLB", Min: 0, Def: 0, Max: 65536000, Tag: 0x534b4c42, Strid: 268, Flags: 0x00000000},
					{Name: "TRMG", Min: 0, Def: 0, Max: 65536000, Tag: 0x54524d47, Strid: 269, Flags: 0x00000000},
					{Name: "TRME", Min: 0, Def: 0, Max: 65536000, Tag: 0x54524d45, Strid: 270, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Strid: 271, Psid: 65535},
					{Coords: []fixed.Int16_16{65536000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Strid: 272, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 0, 0, 0, 65536000, 0, 0, 0, 0, 0, 0}, Strid: 273, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 65536000, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Strid: 274, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 0, 0, 65536000, 0, 0, 0, 0, 0, 0, 0}, Strid: 275, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 32768000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Strid: 276, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65536000, 0, 0, 0}, Strid: 277, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65536000, 0, 0, 0, 0}, Strid: 278, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 65536000, 0, 0, 0, 0, 65536000, 0, 0, 0, 0}, Strid: 279, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 65536000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Strid: 280, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 65536000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Strid: 281, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65536000}, Strid: 282, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 32768000, 32768000, 0, 0, 0, 0, 0, 0, 0, 0}, Strid: 283, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65536000, 0}, Strid: 284, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 0, 0, 0, 0, 65536000, 0, 0, 0, 0, 0}, Strid: 285, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 65536000, 0, 0, 0, 65536000, 65536000, 0, 0, 0, 0}, Strid: 286, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 0, 49152000, 0, 16384000, 65536000, 16384000, 16384000, 65536000, 49152000, 32768000, 32768000, 65536000, 49152000, 32768000}, Strid: 287, Psid: 65535},
				},
			},
		},
		{
			name: "FiraCode-VF.ttf",
			face: faceFromPath("variable/FiraCode/FiraCode-VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 5,
				Axis: []VarAxis{
					{Name: "Weight", Min: 19660800, Def: 19660800, Max: 45875200, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{19660800}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{29491200}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 261, Psid: 65535},
				},
			},
		},
		{
			name: "MovementV.ttf",
			face: faceFromPath("variable/movement/MovementV.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 6,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6553600, Def: 6553600, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "SPAC", Min: 6553600, Def: 6553600, Max: 7864320, Tag: 0x53504143, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{6553600, 6553600}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 7864320}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{58982400, 6553600}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{58982400, 7864320}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{16384000, 6553600}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{16384000, 7864320}, Strid: 263, Psid: 65535},
				},
			},
		},
		{
			name: "LeagueSpartanVariable.ttf",
			face: faceFromPath("variable/league-spartan/LeagueSpartanVariable.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 8,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{18153472}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{24903680}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{58982400}, Strid: 264, Psid: 65535},
				},
			},
		},
		{
			name: "HeptaSlab-VF.ttf",
			face: faceFromPath("variable/Hepta-Slab/HeptaSlab-VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 10,
				Axis: []VarAxis{
					{Name: "Weight", Min: 65536, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{65536}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800}, Strid: 265, Psid: 65535},
					{Coords: []fixed.Int16_16{58982400}, Strid: 266, Psid: 65535},
				},
			},
		},
		{
			name: "NunitoVFBeta.ttf",
			face: faceFromPath("variable/nunito/NunitoVFBeta.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 9,
				Axis: []VarAxis{
					{Name: "Weight", Min: 3538944, Def: 3538944, Max: 10223616, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{2752512}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{3997696}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{5308416}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{6619136}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{8192000}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{9895936}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{11665408}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{13631488}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{3538944}, Strid: 17, Psid: 6},
				},
			},
		},
		{
			name: "KayakVF.ttf",
			face: faceFromPath("variable/KayakVF/KayakVF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 5,
				Axis: []VarAxis{
					{Name: "Weight", Min: 0, Def: 0, Max: 65536000, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0}, Strid: 257, Psid: 258},
					{Coords: []fixed.Int16_16{16384000}, Strid: 259, Psid: 260},
					{Coords: []fixed.Int16_16{32768000}, Strid: 261, Psid: 262},
					{Coords: []fixed.Int16_16{49152000}, Strid: 263, Psid: 264},
					{Coords: []fixed.Int16_16{65536000}, Strid: 265, Psid: 266},
				},
			},
		},
		{
			name: "Amstelvar-Roman-VF.ttf",
			face: faceFromPath("variable/Amstelvar/Amstelvar-Roman-VF.ttf"),
			want: &MMVar{
				NumAxis:        10,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6553600, Def: 26214400, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "OpticalSize", Min: 524288, Def: 917504, Max: 9437184, Tag: VarAxisTagOpsz, Strid: 257, Flags: 0x00000000},
					{Name: "Width", Min: 3276800, Def: 6553600, Max: 8192000, Tag: VarAxisTagWdth, Strid: 258, Flags: 0x00000000},
					{Name: "XTRA", Min: 29097984, Def: 64225280, Max: 72351744, Tag: 0x58545241, Strid: 259, Flags: 0x00000000},
					{Name: "XOPQ", Min: 2359296, Def: 11534336, Max: 34471936, Tag: 0x584f5051, Strid: 260, Flags: 0x00000000},
					{Name: "YOPQ", Min: 1900544, Def: 8126464, Max: 14680064, Tag: 0x594f5051, Strid: 261, Flags: 0x00000000},
					{Name: "YTAS", Min: 43712512, Def: 50266112, Max: 56819712, Tag: 0x59544153, Strid: 262, Flags: 0x00000000},
					{Name: "YTDE", Min: 9175040, Def: 15728640, Max: 22282240, Tag: 0x59544445, Strid: 263, Flags: 0x00000000},
					{Name: "YTUC", Min: 36044800, Def: 49152000, Max: 55705600, Tag: 0x59545543, Strid: 264, Flags: 0x00000000},
					{Name: "YTLC", Min: 29163520, Def: 32768000, Max: 39321600, Tag: 0x59544c43, Strid: 265, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{26214400, 917504, 6553600, 64225280, 11534336, 8126464, 50266112, 15728640, 49152000, 32768000}, Strid: 17, Psid: 6},
				},
			},
		},
		{
			name: "Amstelvar-Roman-VF-APPS.ttf",
			face: faceFromPath("variable/Amstelvar/Amstelvar-Roman-VF-APPS.ttf"),
			want: &MMVar{
				NumAxis:        10,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6553600, Def: 26214400, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "OpticalSize", Min: 65536, Def: 917504, Max: 65536000, Tag: VarAxisTagOpsz, Strid: 257, Flags: 0x00000000},
					{Name: "Width", Min: 3276800, Def: 6553600, Max: 8192000, Tag: VarAxisTagWdth, Strid: 258, Flags: 0x00000000},
					{Name: "XTRA", Min: 29097984, Def: 64225280, Max: 72351744, Tag: 0x58545241, Strid: 259, Flags: 0x00000000},
					{Name: "XOPQ", Min: 2359296, Def: 11534336, Max: 34471936, Tag: 0x584f5051, Strid: 260, Flags: 0x00000000},
					{Name: "YOPQ", Min: 1900544, Def: 8126464, Max: 14680064, Tag: 0x594f5051, Strid: 261, Flags: 0x00000000},
					{Name: "YTAS", Min: 43712512, Def: 50266112, Max: 56819712, Tag: 0x59544153, Strid: 262, Flags: 0x00000000},
					{Name: "YTDE", Min: 9175040, Def: 15728640, Max: 22282240, Tag: 0x59544445, Strid: 263, Flags: 0x00000000},
					{Name: "YTUC", Min: 36044800, Def: 49152000, Max: 55705600, Tag: 0x59545543, Strid: 264, Flags: 0x00000000},
					{Name: "YTLC", Min: 29163520, Def: 32768000, Max: 39321600, Tag: 0x59544c43, Strid: 265, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{26214400, 917504, 6553600, 64225280, 11534336, 8126464, 50266112, 15728640, 49152000, 32768000}, Strid: 17, Psid: 6},
				},
			},
		},
		{
			name: "Hidden-Axis-Amstel.ttf",
			face: faceFromPath("variable/Amstelvar/Hidden-Axis-Amstel.ttf"),
			want: &MMVar{
				NumAxis:        10,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6553600, Def: 26214400, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000001},
					{Name: "OpticalSize", Min: 524288, Def: 917504, Max: 9437184, Tag: VarAxisTagOpsz, Strid: 257, Flags: 0x00000000},
					{Name: "Width", Min: 3276800, Def: 6553600, Max: 8192000, Tag: VarAxisTagWdth, Strid: 258, Flags: 0x00000000},
					{Name: "XTRA", Min: 29097984, Def: 64225280, Max: 72351744, Tag: 0x58545241, Strid: 259, Flags: 0x00000000},
					{Name: "XOPQ", Min: 2359296, Def: 11534336, Max: 34471936, Tag: 0x584f5051, Strid: 260, Flags: 0x00000000},
					{Name: "YOPQ", Min: 1900544, Def: 8126464, Max: 14680064, Tag: 0x594f5051, Strid: 261, Flags: 0x00000000},
					{Name: "YTAS", Min: 43712512, Def: 50266112, Max: 56819712, Tag: 0x59544153, Strid: 262, Flags: 0x00000000},
					{Name: "YTDE", Min: 9175040, Def: 15728640, Max: 22282240, Tag: 0x59544445, Strid: 263, Flags: 0x00000000},
					{Name: "YTUC", Min: 36044800, Def: 49152000, Max: 55705600, Tag: 0x59545543, Strid: 264, Flags: 0x00000000},
					{Name: "YTLC", Min: 29163520, Def: 32768000, Max: 39321600, Tag: 0x59544c43, Strid: 265, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{26214400, 917504, 6553600, 64225280, 11534336, 8126464, 50266112, 15728640, 49152000, 32768000}, Strid: 17, Psid: 6},
				},
			},
		},
		{
			name: "RibbonVF.ttf",
			face: faceFromPath("variable/RibbonVF/RibbonVF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 2,
				Axis: []VarAxis{
					{Name: "Width", Min: 0, Def: 0, Max: 65536000, Tag: VarAxisTagWdth, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0}, Strid: 257, Psid: 258},
					{Coords: []fixed.Int16_16{65536000}, Strid: 259, Psid: 260},
				},
			},
		},
		{
			name: "gnomon-VF.ttf",
			face: faceFromPath("variable/gnomon/gnomon-VF.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "TOTD", Min: 393216, Def: 589824, Max: 1179648, Tag: 0x544f5444, Strid: 256, Flags: 0x00000000},
					{Name: "DIST", Min: 0, Def: 65536, Max: 196608, Tag: 0x44495354, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{589824, 65536}, Strid: 17, Psid: 6},
				},
			},
		},
		{
			name: "PT Root UI_VF.ttf",
			face: faceFromPath("variable/PT Root UI VF/PT Root UI_VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 4,
				Axis: []VarAxis{
					{Name: "Weight", Min: 19660800, Def: 26214400, Max: 45875200, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{19660800}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 260, Psid: 65535},
				},
			},
		},
		{
			name: "Soulcraft.ttf",
			face: faceFromPath("variable/soulcraft/Soulcraft.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Width", Min: 0, Def: 0, Max: 6553600, Tag: VarAxisTagWdth, Strid: 256, Flags: 0x00000000},
					{Name: "Slant", Min: 0, Def: 0, Max: 6553600, Tag: VarAxisTagSlnt, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0, 0}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "SourceCodeVariable-Roman.ttf",
			face: faceFromPath("variable/source-code-pro/SourceCodeVariable-Roman.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 7,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 26214400, Max: 58982400, Tag: VarAxisTagWght, Strid: 279, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200}, Strid: 280, Psid: 281},
					{Coords: []fixed.Int16_16{19660800}, Strid: 282, Psid: 283},
					{Coords: []fixed.Int16_16{26214400}, Strid: 284, Psid: 285},
					{Coords: []fixed.Int16_16{32768000}, Strid: 286, Psid: 287},
					{Coords: []fixed.Int16_16{39321600}, Strid: 288, Psid: 289},
					{Coords: []fixed.Int16_16{45875200}, Strid: 290, Psid: 291},
					{Coords: []fixed.Int16_16{58982400}, Strid: 292, Psid: 293},
				},
			},
		},
		{
			name: "BarlowGX.ttf",
			face: faceFromPath("variable/barlow/BarlowGX.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 55,
				Axis: []VarAxis{
					{Name: "Weight", Min: 1441792, Def: 1441792, Max: 12320768, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 19660800, Def: 19660800, Max: 32768000, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{1966080, 19660800}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{2555904, 19660800}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{3473408, 19660800}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{4653056, 19660800}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{6291456, 19660800}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{7602176, 19660800}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{9240576, 19660800}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{10878976, 19660800}, Strid: 265, Psid: 65535},
					{Coords: []fixed.Int16_16{12320768, 19660800}, Strid: 266, Psid: 65535},
					{Coords: []fixed.Int16_16{1966080, 26214400}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{2555904, 26214400}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{3473408, 26214400}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{4653056, 26214400}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{6291456, 26214400}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{7602176, 26214400}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{9240576, 26214400}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{10878976, 26214400}, Strid: 265, Psid: 65535},
					{Coords: []fixed.Int16_16{12320768, 26214400}, Strid: 266, Psid: 65535},
					{Coords: []fixed.Int16_16{1966080, 32768000}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{2555904, 32768000}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{3473408, 32768000}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{4653056, 32768000}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{6291456, 32768000}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{7602176, 32768000}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{9240576, 32768000}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{10878976, 32768000}, Strid: 265, Psid: 65535},
					{Coords: []fixed.Int16_16{12320768, 32768000}, Strid: 266, Psid: 65535},
					{Coords: []fixed.Int16_16{1966080, 19660800}, Strid: 267, Psid: 65535},
					{Coords: []fixed.Int16_16{2555904, 19660800}, Strid: 268, Psid: 65535},
					{Coords: []fixed.Int16_16{3473408, 19660800}, Strid: 269, Psid: 65535},
					{Coords: []fixed.Int16_16{4653056, 19660800}, Strid: 270, Psid: 65535},
					{Coords: []fixed.Int16_16{6291456, 19660800}, Strid: 271, Psid: 65535},
					{Coords: []fixed.Int16_16{7602176, 19660800}, Strid: 272, Psid: 65535},
					{Coords: []fixed.Int16_16{9240576, 19660800}, Strid: 273, Psid: 65535},
					{Coords: []fixed.Int16_16{10878976, 19660800}, Strid: 274, Psid: 65535},
					{Coords: []fixed.Int16_16{12320768, 19660800}, Strid: 275, Psid: 65535},
					{Coords: []fixed.Int16_16{1966080, 26214400}, Strid: 267, Psid: 65535},
					{Coords: []fixed.Int16_16{2555904, 26214400}, Strid: 268, Psid: 65535},
					{Coords: []fixed.Int16_16{3473408, 26214400}, Strid: 269, Psid: 65535},
					{Coords: []fixed.Int16_16{4653056, 26214400}, Strid: 270, Psid: 65535},
					{Coords: []fixed.Int16_16{6291456, 26214400}, Strid: 271, Psid: 65535},
					{Coords: []fixed.Int16_16{7602176, 26214400}, Strid: 272, Psid: 65535},
					{Coords: []fixed.Int16_16{9240576, 26214400}, Strid: 273, Psid: 65535},
					{Coords: []fixed.Int16_16{10878976, 26214400}, Strid: 274, Psid: 65535},
					{Coords: []fixed.Int16_16{12320768, 26214400}, Strid: 275, Psid: 65535},
					{Coords: []fixed.Int16_16{1966080, 32768000}, Strid: 267, Psid: 65535},
					{Coords: []fixed.Int16_16{2555904, 32768000}, Strid: 268, Psid: 65535},
					{Coords: []fixed.Int16_16{3473408, 32768000}, Strid: 269, Psid: 65535},
					{Coords: []fixed.Int16_16{4653056, 32768000}, Strid: 270, Psid: 65535},
					{Coords: []fixed.Int16_16{6291456, 32768000}, Strid: 271, Psid: 65535},
					{Coords: []fixed.Int16_16{7602176, 32768000}, Strid: 272, Psid: 65535},
					{Coords: []fixed.Int16_16{9240576, 32768000}, Strid: 273, Psid: 65535},
					{Coords: []fixed.Int16_16{10878976, 32768000}, Strid: 274, Psid: 65535},
					{Coords: []fixed.Int16_16{12320768, 32768000}, Strid: 275, Psid: 65535},
					{Coords: []fixed.Int16_16{1441792, 19660800}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "CrimsonPro-Roman-VF.ttf",
			face: faceFromPath("variable/CrimsonPro/CrimsonPro-Roman-VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 8,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{58982400}, Strid: 264, Psid: 65535},
				},
			},
		},
		{
			name: "WabaBorderGX.ttf",
			face: faceFromPath("variable/waba-border/WabaBorderGX.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 4,
				Axis: []VarAxis{
					{Name: "Weight", Min: 0, Def: 0, Max: 6553600, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 0, Def: 0, Max: 6553600, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0, 0}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 0}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 6553600}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 6553600}, Strid: 261, Psid: 65535},
				},
			},
		},
		{
			name: "AdobeVFPrototype.otf",
			face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 9,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 25516065, Max: 58982400, Tag: VarAxisTagWght, Strid: 259, Flags: 0x00000000},
					{Name: "CNTR", Min: 0, Def: 0, Max: 6553600, Tag: 0x434e5452, Strid: 260, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200, 0}, Strid: 261, Psid: 262},
					{Coords: []fixed.Int16_16{19660800, 0}, Strid: 263, Psid: 264},
					{Coords: []fixed.Int16_16{26214400, 0}, Strid: 265, Psid: 266},
					{Coords: []fixed.Int16_16{39321600, 0}, Strid: 267, Psid: 268},
					{Coords: []fixed.Int16_16{45875200, 0}, Strid: 269, Psid: 270},
					{Coords: []fixed.Int16_16{58982400, 0}, Strid: 271, Psid: 272},
					{Coords: []fixed.Int16_16{58982400, 3276800}, Strid: 273, Psid: 274},
					{Coords: []fixed.Int16_16{58982400, 6553600}, Strid: 275, Psid: 276},
					{Coords: []fixed.Int16_16{25516065, 0}, Strid: 17, Psid: 6},
				},
			},
		},
		{
			name: "AdobeVFPrototype.ttf",
			face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 9,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 25516065, Max: 58982400, Tag: VarAxisTagWght, Strid: 259, Flags: 0x00000000},
					{Name: "CNTR", Min: 0, Def: 0, Max: 6553600, Tag: 0x434e5452, Strid: 260, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200, 0}, Strid: 261, Psid: 262},
					{Coords: []fixed.Int16_16{19660800, 0}, Strid: 263, Psid: 264},
					{Coords: []fixed.Int16_16{26214400, 0}, Strid: 265, Psid: 266},
					{Coords: []fixed.Int16_16{39321600, 0}, Strid: 267, Psid: 268},
					{Coords: []fixed.Int16_16{45875200, 0}, Strid: 269, Psid: 270},
					{Coords: []fixed.Int16_16{58982400, 0}, Strid: 271, Psid: 272},
					{Coords: []fixed.Int16_16{58982400, 3276800}, Strid: 273, Psid: 274},
					{Coords: []fixed.Int16_16{58982400, 6553600}, Strid: 275, Psid: 276},
					{Coords: []fixed.Int16_16{25516065, 0}, Strid: 17, Psid: 6},
				},
			},
		},
		{
			name: "GRADUATE.ttf",
			face: faceFromPath("variable/Graduate-Variable-Font/GRADUATE.ttf"),
			want: &MMVar{
				NumAxis:        12,
				NumDesigns:     0,
				NumNamedstyles: 27,
				Axis: []VarAxis{
					{Name: "XOPQ", Min: 2621440, Def: 2621440, Max: 13107200, Tag: 0x584f5051, Strid: 256, Flags: 0x00000000},
					{Name: "XTRA", Min: 6553600, Def: 26214400, Max: 52428800, Tag: 0x58545241, Strid: 257, Flags: 0x00000000},
					{Name: "OPSZ", Min: 524288, Def: 1048576, Max: 1048576, Tag: 0x4f50535a, Strid: 258, Flags: 0x00000000},
					{Name: "GRAD", Min: 0, Def: 0, Max: 1310720, Tag: 0x47524144, Strid: 259, Flags: 0x00000000},
					{Name: "YTRA", Min: 49152000, Def: 49152000, Max: 55705600, Tag: 0x59545241, Strid: 260, Flags: 0x00000000},
					{Name: "CNTR", Min: 0, Def: 0, Max: 6553600, Tag: 0x434e5452, Strid: 261, Flags: 0x00000000},
					{Name: "YOPQ", Min: 6553600, Def: 6553600, Max: 52428800, Tag: 0x594f5051, Strid: 262, Flags: 0x00000000},
					{Name: "SERF", Min: 0, Def: 0, Max: 1966080, Tag: 0x53455246, Strid: 263, Flags: 0x00000000},
					{Name: "YTAS", Min: 0, Def: 0, Max: 3276800, Tag: 0x59544153, Strid: 264, Flags: 0x00000000},
					{Name: "YTLC", Min: 42598400, Def: 42598400, Max: 49152000, Tag: 0x59544c43, Strid: 265, Flags: 0x00000000},
					{Name: "YTDE", Min: 0, Def: 0, Max: 3276800, Tag: 0x59544445, Strid: 266, Flags: 0x00000000},
					{Name: "SELE", Min: -1310720, Def: 0, Max: 0, Tag: 0x53454c45, Strid: 267, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{2621440, 6553600, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 268, Psid: 65535},
					{Coords: []fixed.Int16_16{2621440, 26214400, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 269, Psid: 65535},
					{Coords: []fixed.Int16_16{2621440, 52428800, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 270, Psid: 65535},
					{Coords: []fixed.Int16_16{3801088, 6553600, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 271, Psid: 65535},
					{Coords: []fixed.Int16_16{3801088, 26214400, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 272, Psid: 65535},
					{Coords: []fixed.Int16_16{3801088, 52428800, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 273, Psid: 65535},
					{Coords: []fixed.Int16_16{4259840, 6553600, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 274, Psid: 65535},
					{Coords: []fixed.Int16_16{4259840, 26214400, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 275, Psid: 65535},
					{Coords: []fixed.Int16_16{4259840, 52428800, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 276, Psid: 65535},
					{Coords: []fixed.Int16_16{5177344, 6553600, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 277, Psid: 65535},
					{Coords: []fixed.Int16_16{5439488, 26214400, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 278, Psid: 65535},
					{Coords: []fixed.Int16_16{5701632, 52428800, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 279, Psid: 65535},
					{Coords: []fixed.Int16_16{6881280, 6553600, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 280, Psid: 65535},
					{Coords: []fixed.Int16_16{6881280, 26214400, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 281, Psid: 65535},
					{Coords: []fixed.Int16_16{6881280, 52428800, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 282, Psid: 65535},
					{Coords: []fixed.Int16_16{8454144, 6553600, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 283, Psid: 65535},
					{Coords: []fixed.Int16_16{8454144, 26214400, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 284, Psid: 65535},
					{Coords: []fixed.Int16_16{8454144, 52428800, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 285, Psid: 65535},
					{Coords: []fixed.Int16_16{10027008, 6553600, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 286, Psid: 65535},
					{Coords: []fixed.Int16_16{10027008, 26214400, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 287, Psid: 65535},
					{Coords: []fixed.Int16_16{10027008, 52428800, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 288, Psid: 65535},
					{Coords: []fixed.Int16_16{11665408, 6553600, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 289, Psid: 65535},
					{Coords: []fixed.Int16_16{11665408, 26214400, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 290, Psid: 65535},
					{Coords: []fixed.Int16_16{11665408, 52428800, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 291, Psid: 65535},
					{Coords: []fixed.Int16_16{12451840, 6553600, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 292, Psid: 65535},
					{Coords: []fixed.Int16_16{12779520, 26214400, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 293, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200, 52428800, 1048576, 0, 49152000, 0, 6553600, 0, 0, 42598400, 0, 0}, Strid: 294, Psid: 65535},
				},
			},
		},
		{
			name: "Zycon.ttf",
			face: faceFromPath("variable/zycon/Zycon.ttf"),
			want: &MMVar{
				NumAxis:        6,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "T1  ", Min: 0, Def: 0, Max: 65536, Tag: 0x54312020, Strid: 256, Flags: 0x00000000},
					{Name: "T2  ", Min: 0, Def: 0, Max: 65536, Tag: 0x54322020, Strid: 257, Flags: 0x00000000},
					{Name: "T3  ", Min: 0, Def: 0, Max: 65536, Tag: 0x54332020, Strid: 258, Flags: 0x00000000},
					{Name: "T4  ", Min: 0, Def: 0, Max: 65536, Tag: 0x54342020, Strid: 259, Flags: 0x00000000},
					{Name: "M1  ", Min: -65536, Def: 0, Max: 65536, Tag: 0x4d312020, Strid: 260, Flags: 0x00000000},
					{Name: "M2  ", Min: -65536, Def: 0, Max: 65536, Tag: 0x4d322020, Strid: 261, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0, 0, 0, 0, 0, 0}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "WorkSans-Roman-VF.ttf",
			face: faceFromPath("variable/work-sans/WorkSans-Roman-VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 9,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6553600, Def: 26214400, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{6553600}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{58982400}, Strid: 265, Psid: 65535},
				},
			},
		},
		{
			name: "WidthAndVWidthVF.ttf",
			face: faceFromPath("variable/width-and-vertical-width-vf/WidthAndVWidthVF.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 12,
				Axis: []VarAxis{
					{Name: "Width", Min: 65536, Def: 65536000, Max: 65536000, Tag: VarAxisTagWdth, Strid: 256, Flags: 0x00000000},
					{Name: "VWID", Min: 65536, Def: 65536000, Max: 65536000, Tag: 0x56574944, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{65536, 32768000}, Strid: 258, Psid: 259},
					{Coords: []fixed.Int16_16{32768000, 65536}, Strid: 260, Psid: 261},
					{Coords: []fixed.Int16_16{6553600, 32768000}, Strid: 262, Psid: 263},
					{Coords: []fixed.Int16_16{32768000, 6553600}, Strid: 264, Psid: 265},
					{Coords: []fixed.Int16_16{16384000, 32768000}, Strid: 266, Psid: 267},
					{Coords: []fixed.Int16_16{32768000, 16384000}, Strid: 268, Psid: 269},
					{Coords: []fixed.Int16_16{32768000, 32768000}, Strid: 270, Psid: 271},
					{Coords: []fixed.Int16_16{49152000, 32768000}, Strid: 272, Psid: 273},
					{Coords: []fixed.Int16_16{32768000, 49152000}, Strid: 274, Psid: 275},
					{Coords: []fixed.Int16_16{65536000, 32768000}, Strid: 276, Psid: 277},
					{Coords: []fixed.Int16_16{32768000, 65536000}, Strid: 278, Psid: 279},
					{Coords: []fixed.Int16_16{65536000, 65536000}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "WidthAndVWidthVF.otf",
			face: faceFromPath("variable/width-and-vertical-width-vf/WidthAndVWidthVF.otf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 12,
				Axis: []VarAxis{
					{Name: "Width", Min: 65536, Def: 65536000, Max: 65536000, Tag: VarAxisTagWdth, Strid: 256, Flags: 0x00000000},
					{Name: "VWID", Min: 65536, Def: 65536000, Max: 65536000, Tag: 0x56574944, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{65536, 32768000}, Strid: 258, Psid: 259},
					{Coords: []fixed.Int16_16{32768000, 65536}, Strid: 260, Psid: 261},
					{Coords: []fixed.Int16_16{6553600, 32768000}, Strid: 262, Psid: 263},
					{Coords: []fixed.Int16_16{32768000, 6553600}, Strid: 264, Psid: 265},
					{Coords: []fixed.Int16_16{16384000, 32768000}, Strid: 266, Psid: 267},
					{Coords: []fixed.Int16_16{32768000, 16384000}, Strid: 268, Psid: 269},
					{Coords: []fixed.Int16_16{32768000, 32768000}, Strid: 270, Psid: 271},
					{Coords: []fixed.Int16_16{49152000, 32768000}, Strid: 272, Psid: 273},
					{Coords: []fixed.Int16_16{32768000, 49152000}, Strid: 274, Psid: 275},
					{Coords: []fixed.Int16_16{65536000, 32768000}, Strid: 276, Psid: 277},
					{Coords: []fixed.Int16_16{32768000, 65536000}, Strid: 278, Psid: 279},
					{Coords: []fixed.Int16_16{65536000, 65536000}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "MarkaziText-VF.ttf",
			face: faceFromPath("variable/markazitext/MarkaziText-VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 4,
				Axis: []VarAxis{
					{Name: "Weight", Min: 26214400, Def: 26214400, Max: 45875200, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{26214400}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{32221867}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{39867733}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 260, Psid: 65535},
				},
			},
		},
		{
			name: "ImposMM.pfb",
			face: faceFromPath("variable/impossible/ImposMM.pfb"),
			want: &MMVar{
				NumAxis:        3,
				NumDesigns:     8,
				NumNamedstyles: 0,
				Axis: []VarAxis{
					{Name: "Width", Min: 0, Def: 65471000, Max: 65536000, Tag: VarAxisTagWdth, Strid: 0, Flags: 0x00000000},
					{Name: "OpticalSize", Min: 0, Def: 65471000, Max: 65536000, Tag: VarAxisTagOpsz, Strid: 0, Flags: 0x00000000},
					{Name: "Serif", Min: 0, Def: 65471000, Max: 65536000, Tag: 0xffffffff, Strid: 0, Flags: 0x00000000},
				},
				Namedstyle: nil,
			},
		},
		{
			name: "LibreFranklinGX-Romans-v4015.ttf",
			face: faceFromPath("variable/Libre-Franklin/LibreFranklinGX-Romans-v4015.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 9,
				Axis: []VarAxis{
					{Name: "Weight", Min: 2621440, Def: 2621440, Max: 13107200, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{2621440}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{4325376}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{5505024}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{6946816}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{10092544}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{11665408}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200}, Strid: 265, Psid: 65535},
				},
			},
		},
		{
			name: "TINY5x3GX.ttf",
			face: faceFromPath("variable/tiny/TINY5x3GX.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 16,
				Axis: []VarAxis{
					{Name: "Weight", Min: 0, Def: 0, Max: 19660800, Tag: VarAxisTagWght, Strid: 263, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{1310720}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{2621440}, Strid: 265, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160}, Strid: 266, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880}, Strid: 267, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600}, Strid: 268, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320}, Strid: 269, Psid: 65535},
					{Coords: []fixed.Int16_16{9175040}, Strid: 270, Psid: 65535},
					{Coords: []fixed.Int16_16{10485760}, Strid: 271, Psid: 65535},
					{Coords: []fixed.Int16_16{11796480}, Strid: 272, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200}, Strid: 273, Psid: 65535},
					{Coords: []fixed.Int16_16{14417920}, Strid: 274, Psid: 65535},
					{Coords: []fixed.Int16_16{15728640}, Strid: 275, Psid: 65535},
					{Coords: []fixed.Int16_16{17039360}, Strid: 276, Psid: 65535},
					{Coords: []fixed.Int16_16{18350080}, Strid: 277, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800}, Strid: 278, Psid: 65535},
					{Coords: []fixed.Int16_16{0}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "IBMPlexSansVar-Roman.ttf",
			face: faceFromPath("variable/IBM-Plex-Sans-Variable/IBMPlexSansVar-Roman.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 16,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6553600, Def: 26214400, Max: 45875200, Tag: VarAxisTagWght, Strid: 261, Flags: 0x00000000},
					{Name: "Width", Min: 5570560, Def: 6553600, Max: 6553600, Tag: VarAxisTagWdth, Strid: 262, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{6553600, 6553600}, Strid: 263, Psid: 264},
					{Coords: []fixed.Int16_16{13107200, 6553600}, Strid: 265, Psid: 266},
					{Coords: []fixed.Int16_16{19660800, 6553600}, Strid: 267, Psid: 268},
					{Coords: []fixed.Int16_16{26214400, 6553600}, Strid: 269, Psid: 270},
					{Coords: []fixed.Int16_16{29491200, 6553600}, Strid: 271, Psid: 272},
					{Coords: []fixed.Int16_16{32768000, 6553600}, Strid: 273, Psid: 274},
					{Coords: []fixed.Int16_16{39321600, 6553600}, Strid: 275, Psid: 276},
					{Coords: []fixed.Int16_16{45875200, 6553600}, Strid: 277, Psid: 278},
					{Coords: []fixed.Int16_16{6553600, 5570560}, Strid: 279, Psid: 280},
					{Coords: []fixed.Int16_16{13107200, 5570560}, Strid: 281, Psid: 282},
					{Coords: []fixed.Int16_16{19660800, 5570560}, Strid: 283, Psid: 284},
					{Coords: []fixed.Int16_16{26214400, 5570560}, Strid: 285, Psid: 286},
					{Coords: []fixed.Int16_16{29491200, 5570560}, Strid: 287, Psid: 288},
					{Coords: []fixed.Int16_16{32768000, 5570560}, Strid: 289, Psid: 290},
					{Coords: []fixed.Int16_16{39321600, 5570560}, Strid: 291, Psid: 292},
					{Coords: []fixed.Int16_16{45875200, 5570560}, Strid: 293, Psid: 294},
				},
			},
		},
		{
			name: "IBMPlexSansVar-Italic.ttf",
			face: faceFromPath("variable/IBM-Plex-Sans-Variable/IBMPlexSansVar-Italic.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 16,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6553600, Def: 26214400, Max: 45875200, Tag: VarAxisTagWght, Strid: 261, Flags: 0x00000000},
					{Name: "Width", Min: 5570560, Def: 6553600, Max: 6553600, Tag: VarAxisTagWdth, Strid: 262, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{6553600, 6553600}, Strid: 263, Psid: 264},
					{Coords: []fixed.Int16_16{13107200, 6553600}, Strid: 265, Psid: 266},
					{Coords: []fixed.Int16_16{19660800, 6553600}, Strid: 267, Psid: 268},
					{Coords: []fixed.Int16_16{26214400, 6553600}, Strid: 269, Psid: 270},
					{Coords: []fixed.Int16_16{29491200, 6553600}, Strid: 271, Psid: 272},
					{Coords: []fixed.Int16_16{32768000, 6553600}, Strid: 273, Psid: 274},
					{Coords: []fixed.Int16_16{39321600, 6553600}, Strid: 275, Psid: 276},
					{Coords: []fixed.Int16_16{45875200, 6553600}, Strid: 277, Psid: 278},
					{Coords: []fixed.Int16_16{6553600, 5570560}, Strid: 279, Psid: 280},
					{Coords: []fixed.Int16_16{13107200, 5570560}, Strid: 281, Psid: 282},
					{Coords: []fixed.Int16_16{19660800, 5570560}, Strid: 283, Psid: 284},
					{Coords: []fixed.Int16_16{26214400, 5570560}, Strid: 285, Psid: 286},
					{Coords: []fixed.Int16_16{29491200, 5570560}, Strid: 287, Psid: 288},
					{Coords: []fixed.Int16_16{32768000, 5570560}, Strid: 289, Psid: 290},
					{Coords: []fixed.Int16_16{39321600, 5570560}, Strid: 291, Psid: 292},
					{Coords: []fixed.Int16_16{45875200, 5570560}, Strid: 293, Psid: 294},
				},
			},
		},
		{
			name: "Cabin_V.ttf",
			face: faceFromPath("variable/cabin/Cabin_V.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 8,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6225920, Def: 6225920, Max: 8388608, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 0, Def: 0, Max: 6553600, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{6225920, 0}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{7143424, 0}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{7602176, 0}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{8388608, 0}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{6160384, 6553600}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{7143424, 6553600}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{7602176, 6553600}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{8388608, 6553600}, Strid: 261, Psid: 65535},
				},
			},
		},
		{
			name: "ZinzinVF.ttf",
			face: faceFromPath("variable/varfonts-ofl/ZinzinVF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "SWSH", Min: 0, Def: 0, Max: 65536000, Tag: 0x53575348, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "LeagueMonoVariable.ttf",
			face: faceFromPath("variable/leaguemono/LeagueMonoVariable.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 41,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6553600, Def: 6553600, Max: 52428800, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 3276800, Def: 6553600, Max: 13107200, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{6553600, 3276800}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200, 3276800}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800, 3276800}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400, 3276800}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000, 3276800}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600, 3276800}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200, 3276800}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800, 3276800}, Strid: 265, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 5242880}, Strid: 266, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200, 5242880}, Strid: 267, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800, 5242880}, Strid: 268, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400, 5242880}, Strid: 269, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000, 5242880}, Strid: 270, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600, 5242880}, Strid: 271, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200, 5242880}, Strid: 272, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800, 5242880}, Strid: 273, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 7536640}, Strid: 274, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200, 7536640}, Strid: 275, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800, 7536640}, Strid: 276, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400, 7536640}, Strid: 277, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000, 7536640}, Strid: 278, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600, 7536640}, Strid: 279, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200, 7536640}, Strid: 280, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800, 7536640}, Strid: 281, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 9830400}, Strid: 282, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200, 9830400}, Strid: 283, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800, 9830400}, Strid: 284, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400, 9830400}, Strid: 285, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000, 9830400}, Strid: 286, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600, 9830400}, Strid: 287, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200, 9830400}, Strid: 288, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800, 9830400}, Strid: 289, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 13107200}, Strid: 290, Psid: 65535},
					{Coords: []fixed.Int16_16{13107200, 13107200}, Strid: 291, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800, 13107200}, Strid: 292, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400, 13107200}, Strid: 293, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000, 13107200}, Strid: 294, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600, 13107200}, Strid: 295, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200, 13107200}, Strid: 296, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800, 13107200}, Strid: 297, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 6553600}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "Changa-VF.ttf",
			face: faceFromPath("variable/changa-vf/Changa-VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 6553600, Def: 6553600, Max: 13107200, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{6553600}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "Lora-VF.ttf",
			face: faceFromPath("variable/Lora-Cyrillic/Lora-VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 3,
				Axis: []VarAxis{
					{Name: "Weight", Min: 26214400, Def: 26214400, Max: 45875200, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{26214400}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 259, Psid: 65535},
				},
			},
		},
		{
			name: "MutatorSans.ttf",
			face: faceFromPath("variable/mutatorSans/MutatorSans.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 6,
				Axis: []VarAxis{
					{Name: "Width", Min: 0, Def: 0, Max: 65536000, Tag: VarAxisTagWdth, Strid: 256, Flags: 0x00000000},
					{Name: "Weight", Min: 0, Def: 0, Max: 65536000, Tag: VarAxisTagWght, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0, 0}, Strid: 258, Psid: 259},
					{Coords: []fixed.Int16_16{0, 65536000}, Strid: 260, Psid: 261},
					{Coords: []fixed.Int16_16{65536000, 0}, Strid: 262, Psid: 263},
					{Coords: []fixed.Int16_16{65536000, 65536000}, Strid: 264, Psid: 265},
					{Coords: []fixed.Int16_16{21430272, 32768000}, Strid: 266, Psid: 65535},
					{Coords: []fixed.Int16_16{21430272, 32768000}, Strid: 267, Psid: 65535},
				},
			},
		},
		{
			name: "SudoVariable.ttf",
			face: faceFromPath("variable/sudo-font/SudoVariable.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 11,
				Axis: []VarAxis{
					{Name: "ital", Min: 0, Def: 0, Max: 65536, Tag: VarAxisTagItal, Strid: 256, Flags: 0x00000000},
					{Name: "Weight", Min: 13107200, Def: 26214400, Max: 45875200, Tag: VarAxisTagWght, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0, 13107200}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 20275200}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 25395200}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 35635200}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 45875200}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{65536, 13107200}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{65536, 20275200}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{65536, 25395200}, Strid: 265, Psid: 65535},
					{Coords: []fixed.Int16_16{65536, 35635200}, Strid: 266, Psid: 65535},
					{Coords: []fixed.Int16_16{65536, 45875200}, Strid: 267, Psid: 65535},
					{Coords: []fixed.Int16_16{0, 26214400}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "Gingham.ttf",
			face: faceFromPath("variable/Gingham/Gingham.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 9,
				Axis: []VarAxis{
					{Name: "Weight", Min: 19660800, Def: 19660800, Max: 45875200, Tag: VarAxisTagWght, Strid: 257, Flags: 0x00000000},
					{Name: "Width", Min: 65536, Def: 65536, Max: 9830400, Tag: VarAxisTagWdth, Strid: 258, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{19660800, 65536}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400, 65536}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200, 65536}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800, 6553600}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400, 6553600}, Strid: 256, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200, 6553600}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800, 9830400}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400, 9830400}, Strid: 265, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200, 9830400}, Strid: 266, Psid: 65535},
				},
			},
		},
		{
			name: "Secuela-Regular-v_1_787-TTF-VF.ttf",
			face: faceFromPath("variable/secuela-variable/Secuela-Regular-v_1_787-TTF-VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 5,
				Axis: []VarAxis{
					{Name: "Weight", Min: 19660800, Def: 19660800, Max: 52428800, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{19660800}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{32768000}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{52428800}, Strid: 261, Psid: 65535},
				},
			},
		},
		{
			name: "SourceHanSansVFProtoHK.otf",
			face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoHK.otf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 4915200, Def: 6553600, Max: 6553600, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200, 6553600}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "SourceHanSansVFProtoJP.otf",
			face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoJP.otf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 4915200, Def: 6553600, Max: 6553600, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200, 6553600}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "SourceHanSansVFProtoKR.otf",
			face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoKR.otf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 4915200, Def: 6553600, Max: 6553600, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200, 6553600}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "SourceHanSansVFProtoTW.otf",
			face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoTW.otf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 4915200, Def: 6553600, Max: 6553600, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200, 6553600}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "SourceHanSansVFProtoCN.otf",
			face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoCN.otf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 4915200, Def: 6553600, Max: 6553600, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200, 6553600}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "SourceHanSansVFProtoMO.otf",
			face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoMO.otf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 1,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
					{Name: "Width", Min: 4915200, Def: 6553600, Max: 6553600, Tag: VarAxisTagWdth, Strid: 257, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200, 6553600}, Strid: 2, Psid: 6},
				},
			},
		},
		{
			name: "TitilliumWeb-Roman-VF.ttf",
			face: faceFromPath("variable/titillium-web-vf/TitilliumWeb-Roman-VF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 6,
				Axis: []VarAxis{
					{Name: "Weight", Min: 13107200, Def: 13107200, Max: 58982400, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{13107200}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{19660800}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{58982400}, Strid: 262, Psid: 65535},
				},
			},
		},
		{
			name: "VotoSerifGX.ttf",
			face: faceFromPath("variable/VotoSerifGX-OFL/VotoSerifGX.ttf"),
			want: &MMVar{
				NumAxis:        3,
				NumDesigns:     0,
				NumNamedstyles: 567,
				Axis: []VarAxis{
					{Name: "Width", Min: 3276800, Def: 8519680, Max: 8519680, Tag: VarAxisTagWdth, Strid: 256, Flags: 0x00000000},
					{Name: "Weight", Min: 1835008, Def: 1835008, Max: 12713984, Tag: VarAxisTagWght, Strid: 257, Flags: 0x00000000},
					{Name: "OpticalSize", Min: 786432, Def: 786432, Max: 4718592, Tag: VarAxisTagOpsz, Strid: 258, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{3276800, 1835008, 786432}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 2424832, 786432}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 3342336, 786432}, Strid: 261, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 4587520, 786432}, Strid: 262, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 6029312, 786432}, Strid: 263, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 7733248, 786432}, Strid: 264, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 9437184, 786432}, Strid: 265, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 11206656, 786432}, Strid: 266, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 12713984, 786432}, Strid: 267, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 1835008, 786432}, Strid: 268, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 2424832, 786432}, Strid: 269, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 3342336, 786432}, Strid: 270, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 4587520, 786432}, Strid: 271, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 6029312, 786432}, Strid: 272, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 7733248, 786432}, Strid: 273, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 9437184, 786432}, Strid: 274, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 11206656, 786432}, Strid: 275, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 12713984, 786432}, Strid: 276, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 1835008, 786432}, Strid: 277, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 2424832, 786432}, Strid: 278, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 3342336, 786432}, Strid: 279, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 4587520, 786432}, Strid: 280, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 6029312, 786432}, Strid: 281, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 7733248, 786432}, Strid: 282, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 9437184, 786432}, Strid: 283, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 11206656, 786432}, Strid: 284, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 12713984, 786432}, Strid: 285, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 1835008, 786432}, Strid: 286, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 2424832, 786432}, Strid: 287, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 3342336, 786432}, Strid: 288, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 4587520, 786432}, Strid: 289, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 6029312, 786432}, Strid: 290, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 7733248, 786432}, Strid: 291, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 9437184, 786432}, Strid: 292, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 11206656, 786432}, Strid: 293, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 12713984, 786432}, Strid: 294, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 1835008, 786432}, Strid: 295, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 2424832, 786432}, Strid: 296, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 3342336, 786432}, Strid: 297, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 4587520, 786432}, Strid: 298, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 6029312, 786432}, Strid: 299, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 7733248, 786432}, Strid: 300, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 9437184, 786432}, Strid: 301, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 11206656, 786432}, Strid: 302, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 12713984, 786432}, Strid: 303, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 1835008, 786432}, Strid: 304, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 2424832, 786432}, Strid: 305, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 3342336, 786432}, Strid: 306, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 4587520, 786432}, Strid: 307, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 6029312, 786432}, Strid: 308, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 7733248, 786432}, Strid: 309, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 9437184, 786432}, Strid: 310, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 11206656, 786432}, Strid: 311, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 12713984, 786432}, Strid: 312, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 1835008, 786432}, Strid: 313, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 2424832, 786432}, Strid: 314, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 3342336, 786432}, Strid: 315, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 4587520, 786432}, Strid: 316, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 6029312, 786432}, Strid: 317, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 7733248, 786432}, Strid: 318, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 9437184, 786432}, Strid: 319, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 11206656, 786432}, Strid: 320, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 12713984, 786432}, Strid: 321, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 1835008, 786432}, Strid: 322, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 2424832, 786432}, Strid: 323, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 3342336, 786432}, Strid: 324, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 4587520, 786432}, Strid: 325, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 6029312, 786432}, Strid: 326, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 7733248, 786432}, Strid: 327, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 9437184, 786432}, Strid: 328, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 11206656, 786432}, Strid: 329, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 12713984, 786432}, Strid: 330, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 1835008, 786432}, Strid: 331, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 2424832, 786432}, Strid: 332, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 3342336, 786432}, Strid: 333, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 4587520, 786432}, Strid: 334, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 6029312, 786432}, Strid: 335, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 7733248, 786432}, Strid: 336, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 9437184, 786432}, Strid: 337, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 11206656, 786432}, Strid: 338, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 12713984, 786432}, Strid: 339, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 1835008, 1179648}, Strid: 340, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 2424832, 1179648}, Strid: 341, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 3342336, 1179648}, Strid: 342, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 4587520, 1179648}, Strid: 343, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 6029312, 1179648}, Strid: 344, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 7733248, 1179648}, Strid: 345, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 9437184, 1179648}, Strid: 346, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 11206656, 1179648}, Strid: 347, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 12713984, 1179648}, Strid: 348, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 1835008, 1179648}, Strid: 349, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 2424832, 1179648}, Strid: 350, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 3342336, 1179648}, Strid: 351, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 4587520, 1179648}, Strid: 352, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 6029312, 1179648}, Strid: 353, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 7733248, 1179648}, Strid: 354, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 9437184, 1179648}, Strid: 355, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 11206656, 1179648}, Strid: 356, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 12713984, 1179648}, Strid: 357, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 1835008, 1179648}, Strid: 358, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 2424832, 1179648}, Strid: 359, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 3342336, 1179648}, Strid: 360, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 4587520, 1179648}, Strid: 361, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 6029312, 1179648}, Strid: 362, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 7733248, 1179648}, Strid: 363, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 9437184, 1179648}, Strid: 364, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 11206656, 1179648}, Strid: 365, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 12713984, 1179648}, Strid: 366, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 1835008, 1179648}, Strid: 367, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 2424832, 1179648}, Strid: 368, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 3342336, 1179648}, Strid: 369, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 4587520, 1179648}, Strid: 370, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 6029312, 1179648}, Strid: 371, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 7733248, 1179648}, Strid: 372, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 9437184, 1179648}, Strid: 373, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 11206656, 1179648}, Strid: 374, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 12713984, 1179648}, Strid: 375, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 1835008, 1179648}, Strid: 376, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 2424832, 1179648}, Strid: 377, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 3342336, 1179648}, Strid: 378, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 4587520, 1179648}, Strid: 379, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 6029312, 1179648}, Strid: 380, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 7733248, 1179648}, Strid: 381, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 9437184, 1179648}, Strid: 382, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 11206656, 1179648}, Strid: 383, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 12713984, 1179648}, Strid: 384, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 1835008, 1179648}, Strid: 385, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 2424832, 1179648}, Strid: 386, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 3342336, 1179648}, Strid: 387, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 4587520, 1179648}, Strid: 388, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 6029312, 1179648}, Strid: 389, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 7733248, 1179648}, Strid: 390, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 9437184, 1179648}, Strid: 391, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 11206656, 1179648}, Strid: 392, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 12713984, 1179648}, Strid: 393, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 1835008, 1179648}, Strid: 394, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 2424832, 1179648}, Strid: 395, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 3342336, 1179648}, Strid: 396, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 4587520, 1179648}, Strid: 397, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 6029312, 1179648}, Strid: 398, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 7733248, 1179648}, Strid: 399, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 9437184, 1179648}, Strid: 400, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 11206656, 1179648}, Strid: 401, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 12713984, 1179648}, Strid: 402, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 1835008, 1179648}, Strid: 403, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 2424832, 1179648}, Strid: 404, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 3342336, 1179648}, Strid: 405, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 4587520, 1179648}, Strid: 406, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 6029312, 1179648}, Strid: 407, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 7733248, 1179648}, Strid: 408, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 9437184, 1179648}, Strid: 409, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 11206656, 1179648}, Strid: 410, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 12713984, 1179648}, Strid: 411, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 1835008, 1179648}, Strid: 412, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 2424832, 1179648}, Strid: 413, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 3342336, 1179648}, Strid: 414, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 4587520, 1179648}, Strid: 415, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 6029312, 1179648}, Strid: 416, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 7733248, 1179648}, Strid: 417, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 9437184, 1179648}, Strid: 418, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 11206656, 1179648}, Strid: 419, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 12713984, 1179648}, Strid: 420, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 1835008, 1572864}, Strid: 421, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 2424832, 1572864}, Strid: 422, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 3342336, 1572864}, Strid: 423, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 4587520, 1572864}, Strid: 424, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 6029312, 1572864}, Strid: 425, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 7733248, 1572864}, Strid: 426, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 9437184, 1572864}, Strid: 427, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 11206656, 1572864}, Strid: 428, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 12713984, 1572864}, Strid: 429, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 1835008, 1572864}, Strid: 430, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 2424832, 1572864}, Strid: 431, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 3342336, 1572864}, Strid: 432, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 4587520, 1572864}, Strid: 433, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 6029312, 1572864}, Strid: 434, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 7733248, 1572864}, Strid: 435, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 9437184, 1572864}, Strid: 436, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 11206656, 1572864}, Strid: 437, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 12713984, 1572864}, Strid: 438, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 1835008, 1572864}, Strid: 439, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 2424832, 1572864}, Strid: 440, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 3342336, 1572864}, Strid: 441, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 4587520, 1572864}, Strid: 442, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 6029312, 1572864}, Strid: 443, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 7733248, 1572864}, Strid: 444, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 9437184, 1572864}, Strid: 445, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 11206656, 1572864}, Strid: 446, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 12713984, 1572864}, Strid: 447, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 1835008, 1572864}, Strid: 448, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 2424832, 1572864}, Strid: 449, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 3342336, 1572864}, Strid: 450, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 4587520, 1572864}, Strid: 451, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 6029312, 1572864}, Strid: 452, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 7733248, 1572864}, Strid: 453, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 9437184, 1572864}, Strid: 454, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 11206656, 1572864}, Strid: 455, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 12713984, 1572864}, Strid: 456, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 1835008, 1572864}, Strid: 457, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 2424832, 1572864}, Strid: 458, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 3342336, 1572864}, Strid: 459, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 4587520, 1572864}, Strid: 460, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 6029312, 1572864}, Strid: 461, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 7733248, 1572864}, Strid: 462, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 9437184, 1572864}, Strid: 463, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 11206656, 1572864}, Strid: 464, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 12713984, 1572864}, Strid: 465, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 1835008, 1572864}, Strid: 466, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 2424832, 1572864}, Strid: 467, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 3342336, 1572864}, Strid: 468, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 4587520, 1572864}, Strid: 469, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 6029312, 1572864}, Strid: 470, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 7733248, 1572864}, Strid: 471, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 9437184, 1572864}, Strid: 472, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 11206656, 1572864}, Strid: 473, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 12713984, 1572864}, Strid: 474, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 1835008, 1572864}, Strid: 475, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 2424832, 1572864}, Strid: 476, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 3342336, 1572864}, Strid: 477, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 4587520, 1572864}, Strid: 478, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 6029312, 1572864}, Strid: 479, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 7733248, 1572864}, Strid: 480, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 9437184, 1572864}, Strid: 481, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 11206656, 1572864}, Strid: 482, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 12713984, 1572864}, Strid: 483, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 1835008, 1572864}, Strid: 484, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 2424832, 1572864}, Strid: 485, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 3342336, 1572864}, Strid: 486, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 4587520, 1572864}, Strid: 487, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 6029312, 1572864}, Strid: 488, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 7733248, 1572864}, Strid: 489, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 9437184, 1572864}, Strid: 490, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 11206656, 1572864}, Strid: 491, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 12713984, 1572864}, Strid: 492, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 1835008, 1572864}, Strid: 493, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 2424832, 1572864}, Strid: 494, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 3342336, 1572864}, Strid: 495, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 4587520, 1572864}, Strid: 496, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 6029312, 1572864}, Strid: 497, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 7733248, 1572864}, Strid: 498, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 9437184, 1572864}, Strid: 499, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 11206656, 1572864}, Strid: 500, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 12713984, 1572864}, Strid: 501, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 1835008, 2359296}, Strid: 502, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 2424832, 2359296}, Strid: 503, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 3342336, 2359296}, Strid: 504, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 4587520, 2359296}, Strid: 505, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 6029312, 2359296}, Strid: 506, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 7733248, 2359296}, Strid: 507, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 9437184, 2359296}, Strid: 508, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 11206656, 2359296}, Strid: 509, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 12713984, 2359296}, Strid: 510, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 1835008, 2359296}, Strid: 511, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 2424832, 2359296}, Strid: 512, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 3342336, 2359296}, Strid: 513, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 4587520, 2359296}, Strid: 514, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 6029312, 2359296}, Strid: 515, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 7733248, 2359296}, Strid: 516, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 9437184, 2359296}, Strid: 517, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 11206656, 2359296}, Strid: 518, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 12713984, 2359296}, Strid: 519, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 1835008, 2359296}, Strid: 520, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 2424832, 2359296}, Strid: 521, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 3342336, 2359296}, Strid: 522, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 4587520, 2359296}, Strid: 523, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 6029312, 2359296}, Strid: 524, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 7733248, 2359296}, Strid: 525, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 9437184, 2359296}, Strid: 526, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 11206656, 2359296}, Strid: 527, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 12713984, 2359296}, Strid: 528, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 1835008, 2359296}, Strid: 529, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 2424832, 2359296}, Strid: 530, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 3342336, 2359296}, Strid: 531, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 4587520, 2359296}, Strid: 532, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 6029312, 2359296}, Strid: 533, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 7733248, 2359296}, Strid: 534, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 9437184, 2359296}, Strid: 535, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 11206656, 2359296}, Strid: 536, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 12713984, 2359296}, Strid: 537, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 1835008, 2359296}, Strid: 538, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 2424832, 2359296}, Strid: 539, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 3342336, 2359296}, Strid: 540, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 4587520, 2359296}, Strid: 541, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 6029312, 2359296}, Strid: 542, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 7733248, 2359296}, Strid: 543, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 9437184, 2359296}, Strid: 544, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 11206656, 2359296}, Strid: 545, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 12713984, 2359296}, Strid: 546, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 1835008, 2359296}, Strid: 547, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 2424832, 2359296}, Strid: 548, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 3342336, 2359296}, Strid: 549, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 4587520, 2359296}, Strid: 550, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 6029312, 2359296}, Strid: 551, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 7733248, 2359296}, Strid: 552, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 9437184, 2359296}, Strid: 553, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 11206656, 2359296}, Strid: 554, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 12713984, 2359296}, Strid: 555, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 1835008, 2359296}, Strid: 556, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 2424832, 2359296}, Strid: 557, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 3342336, 2359296}, Strid: 558, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 4587520, 2359296}, Strid: 559, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 6029312, 2359296}, Strid: 560, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 7733248, 2359296}, Strid: 561, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 9437184, 2359296}, Strid: 562, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 11206656, 2359296}, Strid: 563, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 12713984, 2359296}, Strid: 564, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 1835008, 2359296}, Strid: 565, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 2424832, 2359296}, Strid: 566, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 3342336, 2359296}, Strid: 567, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 4587520, 2359296}, Strid: 568, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 6029312, 2359296}, Strid: 569, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 7733248, 2359296}, Strid: 570, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 9437184, 2359296}, Strid: 571, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 11206656, 2359296}, Strid: 572, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 12713984, 2359296}, Strid: 573, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 1835008, 2359296}, Strid: 574, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 2424832, 2359296}, Strid: 575, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 3342336, 2359296}, Strid: 576, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 4587520, 2359296}, Strid: 577, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 6029312, 2359296}, Strid: 578, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 7733248, 2359296}, Strid: 579, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 9437184, 2359296}, Strid: 580, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 11206656, 2359296}, Strid: 581, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 12713984, 2359296}, Strid: 582, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 1835008, 3145728}, Strid: 583, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 2424832, 3145728}, Strid: 584, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 3342336, 3145728}, Strid: 585, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 4587520, 3145728}, Strid: 586, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 6029312, 3145728}, Strid: 587, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 7733248, 3145728}, Strid: 588, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 9437184, 3145728}, Strid: 589, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 11206656, 3145728}, Strid: 590, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 12713984, 3145728}, Strid: 591, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 1835008, 3145728}, Strid: 592, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 2424832, 3145728}, Strid: 593, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 3342336, 3145728}, Strid: 594, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 4587520, 3145728}, Strid: 595, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 6029312, 3145728}, Strid: 596, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 7733248, 3145728}, Strid: 597, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 9437184, 3145728}, Strid: 598, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 11206656, 3145728}, Strid: 599, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 12713984, 3145728}, Strid: 600, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 1835008, 3145728}, Strid: 601, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 2424832, 3145728}, Strid: 602, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 3342336, 3145728}, Strid: 603, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 4587520, 3145728}, Strid: 604, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 6029312, 3145728}, Strid: 605, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 7733248, 3145728}, Strid: 606, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 9437184, 3145728}, Strid: 607, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 11206656, 3145728}, Strid: 608, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 12713984, 3145728}, Strid: 609, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 1835008, 3145728}, Strid: 610, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 2424832, 3145728}, Strid: 611, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 3342336, 3145728}, Strid: 612, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 4587520, 3145728}, Strid: 613, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 6029312, 3145728}, Strid: 614, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 7733248, 3145728}, Strid: 615, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 9437184, 3145728}, Strid: 616, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 11206656, 3145728}, Strid: 617, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 12713984, 3145728}, Strid: 618, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 1835008, 3145728}, Strid: 619, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 2424832, 3145728}, Strid: 620, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 3342336, 3145728}, Strid: 621, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 4587520, 3145728}, Strid: 622, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 6029312, 3145728}, Strid: 623, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 7733248, 3145728}, Strid: 624, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 9437184, 3145728}, Strid: 625, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 11206656, 3145728}, Strid: 626, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 12713984, 3145728}, Strid: 627, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 1835008, 3145728}, Strid: 628, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 2424832, 3145728}, Strid: 629, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 3342336, 3145728}, Strid: 630, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 4587520, 3145728}, Strid: 631, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 6029312, 3145728}, Strid: 632, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 7733248, 3145728}, Strid: 633, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 9437184, 3145728}, Strid: 634, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 11206656, 3145728}, Strid: 635, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 12713984, 3145728}, Strid: 636, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 1835008, 3145728}, Strid: 637, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 2424832, 3145728}, Strid: 638, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 3342336, 3145728}, Strid: 639, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 4587520, 3145728}, Strid: 640, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 6029312, 3145728}, Strid: 641, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 7733248, 3145728}, Strid: 642, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 9437184, 3145728}, Strid: 643, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 11206656, 3145728}, Strid: 644, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 12713984, 3145728}, Strid: 645, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 1835008, 3145728}, Strid: 646, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 2424832, 3145728}, Strid: 647, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 3342336, 3145728}, Strid: 648, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 4587520, 3145728}, Strid: 649, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 6029312, 3145728}, Strid: 650, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 7733248, 3145728}, Strid: 651, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 9437184, 3145728}, Strid: 652, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 11206656, 3145728}, Strid: 653, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 12713984, 3145728}, Strid: 654, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 1835008, 3145728}, Strid: 655, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 2424832, 3145728}, Strid: 656, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 3342336, 3145728}, Strid: 657, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 4587520, 3145728}, Strid: 658, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 6029312, 3145728}, Strid: 659, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 7733248, 3145728}, Strid: 660, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 9437184, 3145728}, Strid: 661, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 11206656, 3145728}, Strid: 662, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 12713984, 3145728}, Strid: 663, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 1835008, 3932160}, Strid: 664, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 2424832, 3932160}, Strid: 665, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 3342336, 3932160}, Strid: 666, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 4587520, 3932160}, Strid: 667, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 6029312, 3932160}, Strid: 668, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 7733248, 3932160}, Strid: 669, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 9437184, 3932160}, Strid: 670, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 11206656, 3932160}, Strid: 671, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 12713984, 3932160}, Strid: 672, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 1835008, 3932160}, Strid: 673, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 2424832, 3932160}, Strid: 674, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 3342336, 3932160}, Strid: 675, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 4587520, 3932160}, Strid: 676, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 6029312, 3932160}, Strid: 677, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 7733248, 3932160}, Strid: 678, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 9437184, 3932160}, Strid: 679, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 11206656, 3932160}, Strid: 680, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 12713984, 3932160}, Strid: 681, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 1835008, 3932160}, Strid: 682, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 2424832, 3932160}, Strid: 683, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 3342336, 3932160}, Strid: 684, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 4587520, 3932160}, Strid: 685, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 6029312, 3932160}, Strid: 686, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 7733248, 3932160}, Strid: 687, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 9437184, 3932160}, Strid: 688, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 11206656, 3932160}, Strid: 689, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 12713984, 3932160}, Strid: 690, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 1835008, 3932160}, Strid: 691, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 2424832, 3932160}, Strid: 692, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 3342336, 3932160}, Strid: 693, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 4587520, 3932160}, Strid: 694, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 6029312, 3932160}, Strid: 695, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 7733248, 3932160}, Strid: 696, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 9437184, 3932160}, Strid: 697, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 11206656, 3932160}, Strid: 698, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 12713984, 3932160}, Strid: 699, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 1835008, 3932160}, Strid: 700, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 2424832, 3932160}, Strid: 701, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 3342336, 3932160}, Strid: 702, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 4587520, 3932160}, Strid: 703, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 6029312, 3932160}, Strid: 704, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 7733248, 3932160}, Strid: 705, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 9437184, 3932160}, Strid: 706, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 11206656, 3932160}, Strid: 707, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 12713984, 3932160}, Strid: 708, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 1835008, 3932160}, Strid: 709, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 2424832, 3932160}, Strid: 710, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 3342336, 3932160}, Strid: 711, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 4587520, 3932160}, Strid: 712, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 6029312, 3932160}, Strid: 713, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 7733248, 3932160}, Strid: 714, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 9437184, 3932160}, Strid: 715, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 11206656, 3932160}, Strid: 716, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 12713984, 3932160}, Strid: 717, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 1835008, 3932160}, Strid: 718, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 2424832, 3932160}, Strid: 719, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 3342336, 3932160}, Strid: 720, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 4587520, 3932160}, Strid: 721, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 6029312, 3932160}, Strid: 722, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 7733248, 3932160}, Strid: 723, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 9437184, 3932160}, Strid: 724, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 11206656, 3932160}, Strid: 725, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 12713984, 3932160}, Strid: 726, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 1835008, 3932160}, Strid: 727, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 2424832, 3932160}, Strid: 728, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 3342336, 3932160}, Strid: 729, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 4587520, 3932160}, Strid: 730, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 6029312, 3932160}, Strid: 731, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 7733248, 3932160}, Strid: 732, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 9437184, 3932160}, Strid: 733, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 11206656, 3932160}, Strid: 734, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 12713984, 3932160}, Strid: 735, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 1835008, 3932160}, Strid: 736, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 2424832, 3932160}, Strid: 737, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 3342336, 3932160}, Strid: 738, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 4587520, 3932160}, Strid: 739, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 6029312, 3932160}, Strid: 740, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 7733248, 3932160}, Strid: 741, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 9437184, 3932160}, Strid: 742, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 11206656, 3932160}, Strid: 743, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 12713984, 3932160}, Strid: 744, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 1835008, 4718592}, Strid: 745, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 2424832, 4718592}, Strid: 746, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 3342336, 4718592}, Strid: 747, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 4587520, 4718592}, Strid: 748, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 6029312, 4718592}, Strid: 749, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 7733248, 4718592}, Strid: 750, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 9437184, 4718592}, Strid: 751, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 11206656, 4718592}, Strid: 752, Psid: 65535},
					{Coords: []fixed.Int16_16{3276800, 12713984, 4718592}, Strid: 753, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 1835008, 4718592}, Strid: 754, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 2424832, 4718592}, Strid: 755, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 3342336, 4718592}, Strid: 756, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 4587520, 4718592}, Strid: 757, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 6029312, 4718592}, Strid: 758, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 7733248, 4718592}, Strid: 759, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 9437184, 4718592}, Strid: 760, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 11206656, 4718592}, Strid: 761, Psid: 65535},
					{Coords: []fixed.Int16_16{3932160, 12713984, 4718592}, Strid: 762, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 1835008, 4718592}, Strid: 763, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 2424832, 4718592}, Strid: 764, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 3342336, 4718592}, Strid: 765, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 4587520, 4718592}, Strid: 766, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 6029312, 4718592}, Strid: 767, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 7733248, 4718592}, Strid: 768, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 9437184, 4718592}, Strid: 769, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 11206656, 4718592}, Strid: 770, Psid: 65535},
					{Coords: []fixed.Int16_16{4587520, 12713984, 4718592}, Strid: 771, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 1835008, 4718592}, Strid: 772, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 2424832, 4718592}, Strid: 773, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 3342336, 4718592}, Strid: 774, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 4587520, 4718592}, Strid: 775, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 6029312, 4718592}, Strid: 776, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 7733248, 4718592}, Strid: 777, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 9437184, 4718592}, Strid: 778, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 11206656, 4718592}, Strid: 779, Psid: 65535},
					{Coords: []fixed.Int16_16{5242880, 12713984, 4718592}, Strid: 780, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 1835008, 4718592}, Strid: 781, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 2424832, 4718592}, Strid: 782, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 3342336, 4718592}, Strid: 783, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 4587520, 4718592}, Strid: 784, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 6029312, 4718592}, Strid: 785, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 7733248, 4718592}, Strid: 786, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 9437184, 4718592}, Strid: 787, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 11206656, 4718592}, Strid: 788, Psid: 65535},
					{Coords: []fixed.Int16_16{5898240, 12713984, 4718592}, Strid: 789, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 1835008, 4718592}, Strid: 790, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 2424832, 4718592}, Strid: 791, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 3342336, 4718592}, Strid: 792, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 4587520, 4718592}, Strid: 793, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 6029312, 4718592}, Strid: 794, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 7733248, 4718592}, Strid: 795, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 9437184, 4718592}, Strid: 796, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 11206656, 4718592}, Strid: 797, Psid: 65535},
					{Coords: []fixed.Int16_16{6553600, 12713984, 4718592}, Strid: 798, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 1835008, 4718592}, Strid: 799, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 2424832, 4718592}, Strid: 800, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 3342336, 4718592}, Strid: 801, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 4587520, 4718592}, Strid: 802, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 6029312, 4718592}, Strid: 803, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 7733248, 4718592}, Strid: 804, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 9437184, 4718592}, Strid: 805, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 11206656, 4718592}, Strid: 806, Psid: 65535},
					{Coords: []fixed.Int16_16{7208960, 12713984, 4718592}, Strid: 807, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 1835008, 4718592}, Strid: 808, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 2424832, 4718592}, Strid: 809, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 3342336, 4718592}, Strid: 810, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 4587520, 4718592}, Strid: 811, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 6029312, 4718592}, Strid: 812, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 7733248, 4718592}, Strid: 813, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 9437184, 4718592}, Strid: 814, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 11206656, 4718592}, Strid: 815, Psid: 65535},
					{Coords: []fixed.Int16_16{7864320, 12713984, 4718592}, Strid: 816, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 1835008, 4718592}, Strid: 817, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 2424832, 4718592}, Strid: 818, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 3342336, 4718592}, Strid: 819, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 4587520, 4718592}, Strid: 820, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 6029312, 4718592}, Strid: 821, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 7733248, 4718592}, Strid: 822, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 9437184, 4718592}, Strid: 823, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 11206656, 4718592}, Strid: 824, Psid: 65535},
					{Coords: []fixed.Int16_16{8519680, 12713984, 4718592}, Strid: 825, Psid: 65535},
				},
			},
		},
		{
			name: "iAWriterMonoV.ttf",
			face: faceFromPath("variable/iA Writer Mono/iAWriterMonoV.ttf"),
			want: &MMVar{
				NumAxis:        2,
				NumDesigns:     0,
				NumNamedstyles: 4,
				Axis: []VarAxis{
					{Name: "Weight", Min: 26214400, Def: 26214400, Max: 45875200, Tag: VarAxisTagWght, Strid: 261, Flags: 0x00000000},
					{Name: "SPCG", Min: 0, Def: 0, Max: 9830400, Tag: 0x53504347, Strid: 262, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{26214400, 0}, Strid: 263, Psid: 500},
					{Coords: []fixed.Int16_16{29491200, 0}, Strid: 264, Psid: 501},
					{Coords: []fixed.Int16_16{42598400, 0}, Strid: 265, Psid: 502},
					{Coords: []fixed.Int16_16{45875200, 0}, Strid: 266, Psid: 503},
				},
			},
		},
		{
			name: "Selawik-variable.ttf",
			face: faceFromPath("variable/selawik/Selawik-variable.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 5,
				Axis: []VarAxis{
					{Name: "Weight", Min: 19660800, Def: 26214400, Max: 45875200, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{19660800}, Strid: 257, Psid: 65535},
					{Coords: []fixed.Int16_16{22937600}, Strid: 258, Psid: 65535},
					{Coords: []fixed.Int16_16{26214400}, Strid: 259, Psid: 65535},
					{Coords: []fixed.Int16_16{39321600}, Strid: 260, Psid: 65535},
					{Coords: []fixed.Int16_16{45875200}, Strid: 261, Psid: 65535},
				},
			},
		},
		{
			name: "BPdotsSquareVF.ttf",
			face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"),
			want: &MMVar{
				NumAxis:        1,
				NumDesigns:     0,
				NumNamedstyles: 3,
				Axis: []VarAxis{
					{Name: "Weight", Min: 0, Def: 0, Max: 65536000, Tag: VarAxisTagWght, Strid: 256, Flags: 0x00000000},
				},
				Namedstyle: []VarNamedStyle{
					{Coords: []fixed.Int16_16{0}, Strid: 257, Psid: 258},
					{Coords: []fixed.Int16_16{32768000}, Strid: 259, Psid: 260},
					{Coords: []fixed.Int16_16{65536000}, Strid: 261, Psid: 262},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %s", err)
			}
			defer face.Free()

			var freed bool
			defer mockDoneMMVar(func() {
				freed = true
			})()

			got, err := face.MMVar()
			if err != tt.wantErr {
				t.Errorf("Face.MMVar() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil && !freed {
				t.Errorf("Face.MMVar() doneMMVar should have been called")
			}

			if diff := diff(got, tt.want); diff != nil {
				t.Errorf("Face.MMVar() = %v", diff)
			}
		})
	}

	t.Run("freed lib", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}

		if err := l.Free(); err != nil {
			t.Fatalf("unable to free lib: %v", err)
		}

		if _, err := face.MMVar(); err != ErrInvalidFaceHandle {
			t.Errorf("Face.MMVar() error = %v, wantErr %v", err, ErrInvalidFaceHandle)
		}
	})

	t.Run("invalid lib", func(t *testing.T) {
		l, err := NewLibrary()
		if err != nil {
			t.Fatalf("unable to create lib: %v", err)
		}
		face, err := l.NewFaceFromPath(testdata("go", "Go-Regular.ttf"), 0, 0)
		if err != nil {
			t.Fatalf("unable to load face: %v", err)
		}

		ptr := l.ptr
		l.ptr = nil
		defer func() {
			l.ptr = ptr
			l.Free()
		}()

		if _, err := face.MMVar(); err != ErrInvalidFaceHandle {
			t.Errorf("Face.MMVar() error = %v, wantErr %v", err, ErrInvalidFaceHandle)
		}
	})
}

func TestFace_GetSetVarDesignCoords(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		coords  []fixed.Int16_16
		wantErr error
	}{
		{name: "nilFace", face: nilFace, wantErr: ErrInvalidFaceHandle},
		{name: "goRegular", face: goRegular, wantErr: ErrInvalidArgument},

		{name: "Blinker_variable.ttf", face: faceFromPath("variable/blinker/Blinker_variable.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "DecovarAlpha-VF.ttf", face: faceFromPath("variable/Decovar/DecovarAlpha-VF.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000, 0x30000, 0x40000, 0x50000, 0x60000, 0x70000, 0x80000, 0x90000, 0xa0000, 0xb0000, 0xc0000, 0xd0000, 0xe0000, 0xf0000}},
		{name: "FiraCode-VF.ttf", face: faceFromPath("variable/FiraCode/FiraCode-VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "MovementV.ttf", face: faceFromPath("variable/movement/MovementV.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "LeagueSpartanVariable.ttf", face: faceFromPath("variable/league-spartan/LeagueSpartanVariable.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "HeptaSlab-VF.ttf", face: faceFromPath("variable/Hepta-Slab/HeptaSlab-VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "NunitoVFBeta.ttf", face: faceFromPath("variable/nunito/NunitoVFBeta.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "KayakVF.ttf", face: faceFromPath("variable/KayakVF/KayakVF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "Amstelvar-Roman-VF.ttf", face: faceFromPath("variable/Amstelvar/Amstelvar-Roman-VF.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000, 0x30000, 0x40000, 0x50000, 0x60000, 0x70000, 0x80000, 0x90000, 0xa0000}},
		{name: "Amstelvar-Roman-VF-APPS.ttf", face: faceFromPath("variable/Amstelvar/Amstelvar-Roman-VF-APPS.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000, 0x30000, 0x40000, 0x50000, 0x60000, 0x70000, 0x80000, 0x90000, 0xa0000}},
		{name: "Hidden-Axis-Amstel.ttf", face: faceFromPath("variable/Amstelvar/Hidden-Axis-Amstel.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000, 0x30000, 0x40000, 0x50000, 0x60000, 0x70000, 0x80000, 0x90000, 0xa0000}},
		{name: "RibbonVF.ttf", face: faceFromPath("variable/RibbonVF/RibbonVF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "gnomon-VF.ttf", face: faceFromPath("variable/gnomon/gnomon-VF.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "PT Root UI_VF.ttf", face: faceFromPath("variable/PT Root UI VF/PT Root UI_VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "Soulcraft.ttf", face: faceFromPath("variable/soulcraft/Soulcraft.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "SourceCodeVariable-Roman.ttf", face: faceFromPath("variable/source-code-pro/SourceCodeVariable-Roman.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "BarlowGX.ttf", face: faceFromPath("variable/barlow/BarlowGX.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "CrimsonPro-Roman-VF.ttf", face: faceFromPath("variable/CrimsonPro/CrimsonPro-Roman-VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "WabaBorderGX.ttf", face: faceFromPath("variable/waba-border/WabaBorderGX.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "AdobeVFPrototype.ttf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "GRADUATE.ttf", face: faceFromPath("variable/Graduate-Variable-Font/GRADUATE.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000, 0x30000, 0x40000, 0x50000, 0x60000, 0x70000, 0x80000, 0x90000, 0xa0000, 0xb0000, 0xc0000}},
		{name: "Zycon.ttf", face: faceFromPath("variable/zycon/Zycon.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000, 0x30000, 0x40000, 0x50000, 0x60000}},
		{name: "WorkSans-Roman-VF.ttf", face: faceFromPath("variable/work-sans/WorkSans-Roman-VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "WidthAndVWidthVF.ttf", face: faceFromPath("variable/width-and-vertical-width-vf/WidthAndVWidthVF.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "WidthAndVWidthVF.otf", face: faceFromPath("variable/width-and-vertical-width-vf/WidthAndVWidthVF.otf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "MarkaziText-VF.ttf", face: faceFromPath("variable/markazitext/MarkaziText-VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "ImposMM.pfb", face: faceFromPath("variable/impossible/ImposMM.pfb"), coords: []fixed.Int16_16{0x101d0, 0x1ffb8, 0x2fda0}},
		{name: "LibreFranklinGX-Romans-v4015.ttf", face: faceFromPath("variable/Libre-Franklin/LibreFranklinGX-Romans-v4015.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "TINY5x3GX.ttf", face: faceFromPath("variable/tiny/TINY5x3GX.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "IBMPlexSansVar-Roman.ttf", face: faceFromPath("variable/IBM-Plex-Sans-Variable/IBMPlexSansVar-Roman.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "IBMPlexSansVar-Italic.ttf", face: faceFromPath("variable/IBM-Plex-Sans-Variable/IBMPlexSansVar-Italic.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "Cabin_V.ttf", face: faceFromPath("variable/cabin/Cabin_V.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "ZinzinVF.ttf", face: faceFromPath("variable/varfonts-ofl/ZinzinVF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "LeagueMonoVariable.ttf", face: faceFromPath("variable/leaguemono/LeagueMonoVariable.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "Changa-VF.ttf", face: faceFromPath("variable/changa-vf/Changa-VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "Lora-VF.ttf", face: faceFromPath("variable/Lora-Cyrillic/Lora-VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "MutatorSans.ttf", face: faceFromPath("variable/mutatorSans/MutatorSans.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "SudoVariable.ttf", face: faceFromPath("variable/sudo-font/SudoVariable.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "Gingham.ttf", face: faceFromPath("variable/Gingham/Gingham.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "Secuela-Regular-v_1_787-TTF-VF.ttf", face: faceFromPath("variable/secuela-variable/Secuela-Regular-v_1_787-TTF-VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "SourceHanSansVFProtoHK.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoHK.otf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "SourceHanSansVFProtoJP.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoJP.otf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "SourceHanSansVFProtoKR.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoKR.otf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "SourceHanSansVFProtoTW.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoTW.otf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "SourceHanSansVFProtoCN.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoCN.otf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "SourceHanSansVFProtoMO.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoMO.otf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "TitilliumWeb-Roman-VF.ttf", face: faceFromPath("variable/titillium-web-vf/TitilliumWeb-Roman-VF.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "VotoSerifGX.ttf", face: faceFromPath("variable/VotoSerifGX-OFL/VotoSerifGX.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000, 0x30000}},
		{name: "iAWriterMonoV.ttf", face: faceFromPath("variable/iA Writer Mono/iAWriterMonoV.ttf"), coords: []fixed.Int16_16{0x10000, 0x20000}},
		{name: "Selawik-variable.ttf", face: faceFromPath("variable/selawik/Selawik-variable.ttf"), coords: []fixed.Int16_16{0x10000}},
		{name: "BPdotsSquareVF.ttf", face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), coords: []fixed.Int16_16{0x10000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %s", err)
			}
			defer face.Free()

			var setFreed bool
			restore := mockFree(func() {
				setFreed = true
			})

			if err := face.SetVarDesignCoords(tt.coords); err != tt.wantErr {
				t.Errorf("Face.SetVarDesignCoords() error = %v, wantErr %v", err, tt.wantErr)
			}
			restore()
			if len(tt.coords) > 0 && !setFreed {
				t.Errorf("Face.SetMMDesignCoords() free should have been called")
			}

			var getFreed bool
			restore = mockFree(func() {
				getFreed = true
			})

			got, err := face.VarDesignCoords()
			restore()
			if err != tt.wantErr {
				t.Errorf("Face.GetVarDesignCoords() error = %v", err)
			}

			if len(tt.coords) > 0 && !getFreed {
				t.Errorf("Face.GetVarDesignCoords() free should have been called")
			}

			if diff := diff(got, tt.coords); diff != nil {
				t.Errorf("Face.GetVarDesignCoords() %v", diff)
			}

			if got := face.Flags(); len(tt.coords) > 0 && got&FaceFlagVariation == 0 {
				t.Errorf("Face.GetSetVarDesignCoords() Face should have %s, got %s", FaceFlagVariation, got)
			}
		})
	}
}

func TestFace_SetMMDesignCoords(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		coords  []int
		wantErr error
	}{
		{name: "nilFace", face: nilFace, wantErr: ErrInvalidFaceHandle},
		{name: "goRegular", face: goRegular, wantErr: ErrInvalidArgument},

		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), wantErr: ErrInvalidArgument},
		{name: "BPdotsSquareVF.ttf", face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), wantErr: ErrInvalidArgument},

		{name: "ImposMM.pfb", face: faceFromPath("variable/impossible/ImposMM.pfb"), coords: []int{600, 699, 800}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %s", err)
			}
			defer face.Free()

			var freed bool
			defer mockFree(func() {
				freed = true
			})()

			if err := face.SetMMDesignCoords(tt.coords); err != tt.wantErr {
				t.Errorf("Face.SetMMDesignCoords() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(tt.coords) > 0 && !freed {
				t.Errorf("Face.SetMMDesignCoords() free should have been called")
			}

			if tt.wantErr != nil {
				return
			}

			got, err := face.VarDesignCoords()
			if err != nil {
				t.Errorf("Face.SetMMDesignCoords() error = %v", err)
			}

			var gotInts []int
			for _, v := range got {
				gotInts = append(gotInts, int(v>>16))
			}
			if diff := diff(gotInts, tt.coords); diff != nil {
				t.Errorf("Face.SetMMDesignCoords() %v", diff)
			}

			if got := face.Flags(); got&FaceFlagVariation == 0 {
				t.Errorf("Face.GetSetMMDesignCoords() Face should have %s, got %s", FaceFlagVariation, got)
			}
		})
	}
}

func TestFace_GetSetMMBlendCoords(t *testing.T) {
	type setter func(f *Face, coords []fixed.Int16_16) error
	type getter func(f *Face) ([]fixed.Int16_16, error)

	type test struct {
		name    string
		face    func() (testface, error)
		coords  []fixed.Int16_16
		wantErr error
	}

	tests := []test{
		{name: "nilFace", face: nilFace, wantErr: ErrInvalidFaceHandle},
		{name: "goRegular", face: goRegular, wantErr: ErrInvalidArgument},

		{name: "Blinker_variable.ttf", face: faceFromPath("variable/blinker/Blinker_variable.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "DecovarAlpha-VF.ttf", face: faceFromPath("variable/Decovar/DecovarAlpha-VF.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002, 0x00004, 0x00008, 0x00010, 0x00020, 0x00040, 0x00080, 0x00100, 0x00200, 0x00400, 0x00800, 0x01000, 0x02000, 0x04000}},
		{name: "FiraCode-VF.ttf", face: faceFromPath("variable/FiraCode/FiraCode-VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "MovementV.ttf", face: faceFromPath("variable/movement/MovementV.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "LeagueSpartanVariable.ttf", face: faceFromPath("variable/league-spartan/LeagueSpartanVariable.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "HeptaSlab-VF.ttf", face: faceFromPath("variable/Hepta-Slab/HeptaSlab-VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "NunitoVFBeta.ttf", face: faceFromPath("variable/nunito/NunitoVFBeta.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "KayakVF.ttf", face: faceFromPath("variable/KayakVF/KayakVF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "Amstelvar-Roman-VF.ttf", face: faceFromPath("variable/Amstelvar/Amstelvar-Roman-VF.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002, 0x00004, 0x00008, 0x00010, 0x00020, 0x00040, 0x00080, 0x00100, 0x00200}},
		{name: "Amstelvar-Roman-VF-APPS.ttf", face: faceFromPath("variable/Amstelvar/Amstelvar-Roman-VF-APPS.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002, 0x00004, 0x00008, 0x00010, 0x00020, 0x00040, 0x00080, 0x00100, 0x00200}},
		{name: "Hidden-Axis-Amstel.ttf", face: faceFromPath("variable/Amstelvar/Hidden-Axis-Amstel.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002, 0x00004, 0x00008, 0x00010, 0x00020, 0x00040, 0x00080, 0x00100, 0x00200}},
		{name: "RibbonVF.ttf", face: faceFromPath("variable/RibbonVF/RibbonVF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "gnomon-VF.ttf", face: faceFromPath("variable/gnomon/gnomon-VF.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "PT Root UI_VF.ttf", face: faceFromPath("variable/PT Root UI VF/PT Root UI_VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "Soulcraft.ttf", face: faceFromPath("variable/soulcraft/Soulcraft.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "SourceCodeVariable-Roman.ttf", face: faceFromPath("variable/source-code-pro/SourceCodeVariable-Roman.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "BarlowGX.ttf", face: faceFromPath("variable/barlow/BarlowGX.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "CrimsonPro-Roman-VF.ttf", face: faceFromPath("variable/CrimsonPro/CrimsonPro-Roman-VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "WabaBorderGX.ttf", face: faceFromPath("variable/waba-border/WabaBorderGX.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "AdobeVFPrototype.ttf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "GRADUATE.ttf", face: faceFromPath("variable/Graduate-Variable-Font/GRADUATE.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002, 0x00004, 0x00008, 0x00010, 0x00020, 0x00040, 0x00080, 0x00100, 0x00200, 0x00400, 0x00800}},
		{name: "Zycon.ttf", face: faceFromPath("variable/zycon/Zycon.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002, 0x00004, 0x00008, 0x00010, 0x00020}},
		{name: "WorkSans-Roman-VF.ttf", face: faceFromPath("variable/work-sans/WorkSans-Roman-VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "WidthAndVWidthVF.ttf", face: faceFromPath("variable/width-and-vertical-width-vf/WidthAndVWidthVF.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "WidthAndVWidthVF.otf", face: faceFromPath("variable/width-and-vertical-width-vf/WidthAndVWidthVF.otf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "MarkaziText-VF.ttf", face: faceFromPath("variable/markazitext/MarkaziText-VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "ImposMM.pfb", face: faceFromPath("variable/impossible/ImposMM.pfb"), coords: []fixed.Int16_16{0x00001, 0x00002, 0x00004}},
		{name: "LibreFranklinGX-Romans-v4015.ttf", face: faceFromPath("variable/Libre-Franklin/LibreFranklinGX-Romans-v4015.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "TINY5x3GX.ttf", face: faceFromPath("variable/tiny/TINY5x3GX.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "IBMPlexSansVar-Roman.ttf", face: faceFromPath("variable/IBM-Plex-Sans-Variable/IBMPlexSansVar-Roman.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "IBMPlexSansVar-Italic.ttf", face: faceFromPath("variable/IBM-Plex-Sans-Variable/IBMPlexSansVar-Italic.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "Cabin_V.ttf", face: faceFromPath("variable/cabin/Cabin_V.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "ZinzinVF.ttf", face: faceFromPath("variable/varfonts-ofl/ZinzinVF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "LeagueMonoVariable.ttf", face: faceFromPath("variable/leaguemono/LeagueMonoVariable.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "Changa-VF.ttf", face: faceFromPath("variable/changa-vf/Changa-VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "Lora-VF.ttf", face: faceFromPath("variable/Lora-Cyrillic/Lora-VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "MutatorSans.ttf", face: faceFromPath("variable/mutatorSans/MutatorSans.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "SudoVariable.ttf", face: faceFromPath("variable/sudo-font/SudoVariable.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "Gingham.ttf", face: faceFromPath("variable/Gingham/Gingham.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "Secuela-Regular-v_1_787-TTF-VF.ttf", face: faceFromPath("variable/secuela-variable/Secuela-Regular-v_1_787-TTF-VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "SourceHanSansVFProtoHK.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoHK.otf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "SourceHanSansVFProtoJP.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoJP.otf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "SourceHanSansVFProtoKR.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoKR.otf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "SourceHanSansVFProtoTW.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoTW.otf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "SourceHanSansVFProtoCN.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoCN.otf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "SourceHanSansVFProtoMO.otf", face: faceFromPath("variable/variable-font-collection-test/SourceHanSansVFProtoMO.otf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "TitilliumWeb-Roman-VF.ttf", face: faceFromPath("variable/titillium-web-vf/TitilliumWeb-Roman-VF.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "VotoSerifGX.ttf", face: faceFromPath("variable/VotoSerifGX-OFL/VotoSerifGX.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002, 0x00004}},
		{name: "iAWriterMonoV.ttf", face: faceFromPath("variable/iA Writer Mono/iAWriterMonoV.ttf"), coords: []fixed.Int16_16{0x00001, 0x00002}},
		{name: "Selawik-variable.ttf", face: faceFromPath("variable/selawik/Selawik-variable.ttf"), coords: []fixed.Int16_16{0x00001}},
		{name: "BPdotsSquareVF.ttf", face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), coords: []fixed.Int16_16{0x00001}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run := func(t *testing.T, tt test, getter getter, setter setter) {
				face, err := tt.face()
				if err != nil {
					t.Fatalf("unable to load face: %s", err)
				}
				defer face.Free()

				var setFreed bool
				restore := mockFree(func() {
					setFreed = true
				})

				if err := setter(face.Face, tt.coords); err != tt.wantErr {
					t.Errorf("Face.SetMMBlendCoords() error = %v, wantErr %v", err, tt.wantErr)
				}
				restore()
				if len(tt.coords) > 0 && !setFreed {
					t.Errorf("Face.SetMMBlendCoords() free should have been called")
				}

				var getFreed bool
				restore = mockFree(func() {
					getFreed = true
				})

				got, err := getter(face.Face)
				restore()
				if err != tt.wantErr {
					t.Errorf("Face.MMBlendCoords() error = %v", err)
				}
				if len(tt.coords) > 0 && !getFreed {
					t.Errorf("Face.MMBlendCoords() free should have been called")
				}

				if diff := diff(got, tt.coords); diff != nil {
					t.Errorf("Face.MMBlendCoords() %v", diff)
				}

				if got := face.Flags(); len(tt.coords) > 0 && got&FaceFlagVariation == 0 {
					t.Errorf("Face.GetSetMMBlendCoords() Face should have %s, got %s", FaceFlagVariation, got)
				}
			}

			run(t, tt, (*Face).MMBlendCoords, (*Face).SetMMBlendCoords)
			run(t, tt, (*Face).VarBlendCoords, (*Face).SetVarBlendCoords)
		})
	}
}

func TestFace_GetSetMMWeightVector(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		vec     []fixed.Int16_16
		wantErr error
	}{
		{name: "nilFace", face: nilFace, wantErr: ErrInvalidFaceHandle},
		{name: "goRegular", face: goRegular, wantErr: ErrInvalidArgument},

		// this panics on cffdriver.c:878 mm->set_mm_weightvector is NULL
		// {name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), wantErr: ErrInvalidArgument},
		{name: "BPdotsSquareVF.ttf", face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), wantErr: ErrInvalidArgument},

		{name: "ImposMM.pfb", face: faceFromPath("variable/impossible/ImposMM.pfb"), vec: []fixed.Int16_16{0x2000, 0x2000, 0x2000, 0x2000, 0x2000, 0x2000, 0x2000, 0x2000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %s", err)
			}
			defer face.Free()

			var setFreed bool
			restore := mockFree(func() {
				setFreed = true
			})

			if err := face.SetMMWeightVector(tt.vec); err != tt.wantErr {
				t.Errorf("Face.SetMMWeightVector() error = %v, wantErr %v", err, tt.wantErr)
			}
			restore()
			if len(tt.vec) > 0 && !setFreed {
				t.Errorf("Face.SetMMWeightVector() free should have been called")
			}

			var getFreed bool
			restore = mockFree(func() {
				getFreed = true
			})

			got, err := face.MMWeightVector()
			restore()
			if err != tt.wantErr {
				t.Errorf("Face.MMWeightVector() error = %v", err)
			}

			if len(tt.vec) > 0 && !getFreed {
				t.Errorf("Face.MMWeightVector() free should have been called")
			}

			if diff := diff(got, tt.vec); diff != nil {
				t.Errorf("Face.GetSetMMWeightVector() %v", diff)
			}
		})
	}
}

func TestFace_SetNamedInstance(t *testing.T) {
	tests := []struct {
		name    string
		face    func() (testface, error)
		idx     int
		wantErr error
	}{
		{name: "nilFace", face: nilFace, wantErr: ErrInvalidFaceHandle},
		{name: "goRegular", face: goRegular, wantErr: ErrInvalidArgument},

		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 0, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 1, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 2, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 3, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 4, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 5, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 6, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 7, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 8, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 9, wantErr: nil},
		{name: "AdobeVFPrototype.otf", face: faceFromPath("variable/adobe-variable-font-prototype/AdobeVFPrototype.otf"), idx: 10, wantErr: ErrInvalidArgument},
		{name: "BPdotsSquareVF.ttf", face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), idx: 0, wantErr: nil},
		{name: "BPdotsSquareVF.ttf", face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), idx: 1, wantErr: nil},
		{name: "BPdotsSquareVF.ttf", face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), idx: 2, wantErr: nil},
		{name: "BPdotsSquareVF.ttf", face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), idx: 3, wantErr: nil},
		{name: "BPdotsSquareVF.ttf", face: faceFromPath("variable/BPdotsSquareVF/BPdotsSquareVF.ttf"), idx: 4, wantErr: ErrInvalidArgument},

		{name: "ImposMM.pfb", face: faceFromPath("variable/impossible/ImposMM.pfb"), idx: 0, wantErr: nil},
		{name: "ImposMM.pfb", face: faceFromPath("variable/impossible/ImposMM.pfb"), idx: -1, wantErr: nil},
		{name: "ImposMM.pfb", face: faceFromPath("variable/impossible/ImposMM.pfb"), idx: 999, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			face, err := tt.face()
			if err != nil {
				t.Fatalf("unable to load face: %s", err)
			}
			defer face.Free()

			if err := face.SetNamedInstance(tt.idx); err != tt.wantErr {
				t.Errorf("Face.SetNamedInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
