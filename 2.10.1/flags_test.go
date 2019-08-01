package freetype2

import "testing"

func TestSubGlyphFlag_String(t *testing.T) {
	tests := []struct {
		name string
		x    SubGlyphFlag
		want string
	}{
		{name: "0", x: 0, want: ""},
		{name: "one", x: SubGlyphFlagXyScale, want: "XyScale"},
		{name: "two", x: SubGlyphFlagArgsAreXyValues | SubGlyphFlagScale, want: "ArgsAreXyValues|Scale"},
		{
			name: "all",
			x: SubGlyphFlagArgsAreWords | SubGlyphFlagArgsAreXyValues | SubGlyphFlagRoundXyToGrid |
				SubGlyphFlagScale | SubGlyphFlagXyScale | SubGlyphFlag2x2 | SubGlyphFlagUseMyMetrics,
			want: "ArgsAreWords|ArgsAreXyValues|RoundXyToGrid|Scale|XyScale|2x2|UseMyMetrics",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("SubGlyphFlag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSTypeFlag_String(t *testing.T) {
	tests := []struct {
		name string
		x    FSTypeFlag
		want string
	}{
		{name: "0", x: 0, want: "InstallableEmbedding"},
		{name: "one", x: FsTypeFlagRestrictedLicenseEmbedding, want: "RestrictedLicenseEmbedding"},
		{
			name: "two",
			x:    FsTypeFlagRestrictedLicenseEmbedding | FsTypeFlagPreviewAndPrintEmbedding,
			want: "RestrictedLicenseEmbedding|PreviewAndPrintEmbedding",
		},
		{
			name: "all",
			x: FsTypeFlagInstallableEmbedding | FsTypeFlagRestrictedLicenseEmbedding | FsTypeFlagPreviewAndPrintEmbedding |
				FsTypeFlagEditableEmbedding | FsTypeFlagNoSubsetting | FsTypeFlagBitmapEmbeddingOnly,
			want: "RestrictedLicenseEmbedding|PreviewAndPrintEmbedding|EditableEmbedding|NoSubsetting|BitmapEmbeddingOnly",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("FSTypeFlag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadFlag_String(t *testing.T) {
	tests := []struct {
		name string
		x    LoadFlag
		want string
	}{
		{name: "0", x: 0, want: "Default"},
		{name: "unknown", x: 1 << 30, want: "TargetNormal"},
		{name: "one", x: LoadColor, want: "Color|TargetNormal"},
		{name: "two", x: LoadMonochrome | LoadNoAutohint, want: "Monochrome|NoAutohint|TargetNormal"},
		{
			name: "all",
			x: LoadDefault | LoadNoScale | LoadNoHinting | LoadRender | LoadNoBitmap | LoadVerticalLayout |
				LoadForceAutohint | LoadPedantic | LoadNoRecurse | LoadIgnoreTransform | LoadMonochrome |
				LoadLinearDesign | LoadNoAutohint | LoadColor | LoadComputeMetrics | LoadBitmapMetricsOnly | LoadTargetNormal,
			want: "NoScale|NoHinting|Render|NoBitmap|VerticalLayout|ForceAutohint|Pedantic|NoRecurse|IgnoreTransform|" +
				"Monochrome|LinearDesign|NoAutohint|Color|ComputeMetrics|BitmapMetricsOnly|TargetNormal",
		},
		{
			name: "all TargetLight",
			x: LoadDefault | LoadNoScale | LoadNoHinting | LoadRender | LoadNoBitmap | LoadVerticalLayout |
				LoadForceAutohint | LoadPedantic | LoadNoRecurse | LoadIgnoreTransform | LoadMonochrome |
				LoadLinearDesign | LoadNoAutohint | LoadColor | LoadComputeMetrics | LoadBitmapMetricsOnly | LoadTargetLight,
			want: "NoScale|NoHinting|Render|NoBitmap|VerticalLayout|ForceAutohint|Pedantic|NoRecurse|IgnoreTransform|" +
				"Monochrome|LinearDesign|NoAutohint|Color|ComputeMetrics|BitmapMetricsOnly|TargetLight",
		},
		{
			name: "all TargetMono",
			x: LoadDefault | LoadNoScale | LoadNoHinting | LoadRender | LoadNoBitmap | LoadVerticalLayout |
				LoadForceAutohint | LoadPedantic | LoadNoRecurse | LoadIgnoreTransform | LoadMonochrome |
				LoadLinearDesign | LoadNoAutohint | LoadColor | LoadComputeMetrics | LoadBitmapMetricsOnly | LoadTargetMono,
			want: "NoScale|NoHinting|Render|NoBitmap|VerticalLayout|ForceAutohint|Pedantic|NoRecurse|IgnoreTransform|" +
				"Monochrome|LinearDesign|NoAutohint|Color|ComputeMetrics|BitmapMetricsOnly|TargetMono",
		},
		{
			name: "all TargetLCD",
			x: LoadDefault | LoadNoScale | LoadNoHinting | LoadRender | LoadNoBitmap | LoadVerticalLayout |
				LoadForceAutohint | LoadPedantic | LoadNoRecurse | LoadIgnoreTransform | LoadMonochrome |
				LoadLinearDesign | LoadNoAutohint | LoadColor | LoadComputeMetrics | LoadBitmapMetricsOnly | LoadTargetLCD,
			want: "NoScale|NoHinting|Render|NoBitmap|VerticalLayout|ForceAutohint|Pedantic|NoRecurse|IgnoreTransform|" +
				"Monochrome|LinearDesign|NoAutohint|Color|ComputeMetrics|BitmapMetricsOnly|TargetLCD",
		},
		{
			name: "all TargetLCDV",
			x: LoadDefault | LoadNoScale | LoadNoHinting | LoadRender | LoadNoBitmap | LoadVerticalLayout |
				LoadForceAutohint | LoadPedantic | LoadNoRecurse | LoadIgnoreTransform | LoadMonochrome |
				LoadLinearDesign | LoadNoAutohint | LoadColor | LoadComputeMetrics | LoadBitmapMetricsOnly | LoadTargetLCDV,
			want: "NoScale|NoHinting|Render|NoBitmap|VerticalLayout|ForceAutohint|Pedantic|NoRecurse|IgnoreTransform|" +
				"Monochrome|LinearDesign|NoAutohint|Color|ComputeMetrics|BitmapMetricsOnly|TargetLCDV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("LoadFlag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFaceFlag_String(t *testing.T) {
	tests := []struct {
		name string
		x    FaceFlag
		want string
	}{
		{name: "0", x: 0, want: ""},
		{name: "one", x: FaceFlagColor, want: "Color"},
		{name: "two", x: FaceFlagKerning | FaceFlagCidKeyed, want: "Kerning|CidKeyed"},
		{name: "all", x: FaceFlagScalable | FaceFlagFixedSizes | FaceFlagFixedWidth | FaceFlagSfnt | FaceFlagHorizontal |
			FaceFlagVertical | FaceFlagKerning | FaceFlagMultipleMasters | FaceFlagGlyphNames | FaceFlagHinter |
			FaceFlagCidKeyed | FaceFlagTricky | FaceFlagColor | FaceFlagVariation,
			want: "Scalable|FixedSizes|FixedWidth|Sfnt|Horizontal|Vertical|Kerning|MultipleMasters|GlyphNames|Hinter|" +
				"CidKeyed|Tricky|Color|Variation",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("FaceFlag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStyleFlag_String(t *testing.T) {
	tests := []struct {
		name string
		x    StyleFlag
		want string
	}{
		{name: "0", x: 0, want: ""},
		{name: "one", x: StyleFlagItalic, want: "Italic"},
		{name: "all", x: StyleFlagItalic | StyleFlagBold, want: "Italic|Bold"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("StyleFlag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutlineFlag_String(t *testing.T) {
	tests := []struct {
		name string
		x    OutlineFlag
		want string
	}{
		{name: "0", x: 0, want: ""},
		{name: "one", x: OutlineIgnoreDropouts, want: "IgnoreDropouts"},
		{name: "two", x: OutlineEvenOddFill | OutlineIncludeStubs, want: "EvenOddFill|IncludeStubs"},
		{
			name: "all",
			x: OutlineOwner | OutlineEvenOddFill | OutlineReverseFill | OutlineIgnoreDropouts |
				OutlineSmartDropouts | OutlineIncludeStubs | OutlineHighPrecision | OutlineSinglePass,
			want: "Owner|EvenOddFill|ReverseFill|IgnoreDropouts|SmartDropouts|IncludeStubs|HighPrecision|SinglePass",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("OutlineFlag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
