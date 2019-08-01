package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_TRUETYPE_TABLES_H
import (
	"C"
)

// SubGlyphFlag is a list of constants describing subglyphs.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_subglyph_flag_xxx
type SubGlyphFlag uint

// Please refer to the ‘glyf’ table description in the OpenType specification
// for the meaning of the various flags (which get synthesized for non-OpenType
// subglyphs).
//
// https://docs.microsoft.com/en-us/typography/opentype/spec/glyf#composite-glyph-description
const (
	SubGlyphFlagArgsAreWords    SubGlyphFlag = C.FT_SUBGLYPH_FLAG_ARGS_ARE_WORDS
	SubGlyphFlagArgsAreXyValues SubGlyphFlag = C.FT_SUBGLYPH_FLAG_ARGS_ARE_XY_VALUES
	SubGlyphFlagRoundXyToGrid   SubGlyphFlag = C.FT_SUBGLYPH_FLAG_ROUND_XY_TO_GRID
	SubGlyphFlagScale           SubGlyphFlag = C.FT_SUBGLYPH_FLAG_SCALE
	SubGlyphFlagXyScale         SubGlyphFlag = C.FT_SUBGLYPH_FLAG_XY_SCALE
	SubGlyphFlag2x2             SubGlyphFlag = C.FT_SUBGLYPH_FLAG_2X2
	SubGlyphFlagUseMyMetrics    SubGlyphFlag = C.FT_SUBGLYPH_FLAG_USE_MY_METRICS
)

func (x SubGlyphFlag) String() string {
	// the maximum concatenated len, at the time of writing, is 74.
	s := make([]byte, 0, 74)

	if x&SubGlyphFlagArgsAreWords == SubGlyphFlagArgsAreWords {
		s = append(s, []byte("ArgsAreWords|")...)
	}
	if x&SubGlyphFlagArgsAreXyValues == SubGlyphFlagArgsAreXyValues {
		s = append(s, []byte("ArgsAreXyValues|")...)
	}
	if x&SubGlyphFlagRoundXyToGrid == SubGlyphFlagRoundXyToGrid {
		s = append(s, []byte("RoundXyToGrid|")...)
	}
	if x&SubGlyphFlagScale == SubGlyphFlagScale {
		s = append(s, []byte("Scale|")...)
	}
	if x&SubGlyphFlagXyScale == SubGlyphFlagXyScale {
		s = append(s, []byte("XyScale|")...)
	}
	if x&SubGlyphFlag2x2 == SubGlyphFlag2x2 {
		s = append(s, []byte("2x2|")...)
	}
	if x&SubGlyphFlagUseMyMetrics == SubGlyphFlagUseMyMetrics {
		s = append(s, []byte("UseMyMetrics|")...)
	}

	if len(s) == 0 {
		return ""
	}

	return string(s[:len(s)-1]) // trim the leading |
}

// FSTypeFlag is a list of bit flags used in the fsType field of the OS/2 table
// in a TrueType or OpenType font and the FSType entry in a PostScript font.
// These bit flags are returned by Face,FSTypeFlags(); they inform client
// applications of embedding and subsetting restrictions associated with a font.
//
// See https://www.adobe.com/content/dam/Adobe/en/devnet/acrobat/pdfs/FontPolicies.pdf
// for more details.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_fstype_xxx
type FSTypeFlag uint

const (
	// FsTypeFlagInstallableEmbedding fonts with no fsType bit set may be
	// embedded and permanently installed on the remote system by an application.
	FsTypeFlagInstallableEmbedding FSTypeFlag = C.FT_FSTYPE_INSTALLABLE_EMBEDDING
	// FsTypeFlagRestrictedLicenseEmbedding fonts that have only this bit set
	// must not be modified, embedded or exchanged in any manner without first
	// obtaining permission of the font software copyright owner.
	FsTypeFlagRestrictedLicenseEmbedding FSTypeFlag = C.FT_FSTYPE_RESTRICTED_LICENSE_EMBEDDING
	// FsTypeFlagPreviewAndPrintEmbedding the font may be embedded and
	// temporarily loaded on the remote system. Documents containing Preview &
	// Print fonts must be opened ‘read-only’; no edits can be applied to the
	// document.
	FsTypeFlagPreviewAndPrintEmbedding FSTypeFlag = C.FT_FSTYPE_PREVIEW_AND_PRINT_EMBEDDING
	// FsTypeFlagEditableEmbedding the font may be embedded but must only be
	// installed temporarily on other systems. In contrast to Preview & Print
	// fonts, documents containing editable fonts may be opened for reading,
	// editing is permitted, and changes may be saved.
	FsTypeFlagEditableEmbedding FSTypeFlag = C.FT_FSTYPE_EDITABLE_EMBEDDING
	// FsTypeFlagNoSubsetting the font may not be subsetted prior to embedding.
	FsTypeFlagNoSubsetting FSTypeFlag = C.FT_FSTYPE_NO_SUBSETTING
	// FsTypeFlagBitmapEmbeddingOnly only bitmaps contained in the font may be
	// embedded; no outline data may be embedded. If there are no bitmaps
	// available in the font, then the font is unembeddable.
	FsTypeFlagBitmapEmbeddingOnly FSTypeFlag = C.FT_FSTYPE_BITMAP_EMBEDDING_ONLY
)

func (x FSTypeFlag) String() string {
	if x == FsTypeFlagInstallableEmbedding {
		return "InstallableEmbedding"
	}

	// the maximum concatenated len, at the time of writing, is 103.
	s := make([]byte, 0, 103)

	if x&FsTypeFlagRestrictedLicenseEmbedding == FsTypeFlagRestrictedLicenseEmbedding {
		s = append(s, []byte("RestrictedLicenseEmbedding|")...)
	}
	if x&FsTypeFlagPreviewAndPrintEmbedding == FsTypeFlagPreviewAndPrintEmbedding {
		s = append(s, []byte("PreviewAndPrintEmbedding|")...)
	}
	if x&FsTypeFlagEditableEmbedding == FsTypeFlagEditableEmbedding {
		s = append(s, []byte("EditableEmbedding|")...)
	}
	if x&FsTypeFlagNoSubsetting == FsTypeFlagNoSubsetting {
		s = append(s, []byte("NoSubsetting|")...)
	}
	if x&FsTypeFlagBitmapEmbeddingOnly == FsTypeFlagBitmapEmbeddingOnly {
		s = append(s, []byte("BitmapEmbeddingOnly|")...)
	}

	return string(s[:len(s)-1]) // trim the leading |
}

// LoadFlag is a list of bit field constants for LoadGlyph to indicate what
// kind of operations to perform during glyph loading.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_load_xxx
type LoadFlag int32

const (
	// LoadDefault is used as the default glyph load operation. In this case, the following happens:
	//
	// 1 - FreeType looks for a bitmap for the glyph corresponding to the face's current size. If one is found, the function
	// returns. The bitmap data can be accessed from the glyph slot (see note below).
	// 2 - If no embedded bitmap is searched for or found, FreeType looks for a scalable outline. If one is found, it is
	// loaded from the font file, scaled to device pixels, then ‘hinted’ to the pixel grid in order to optimize it. The
	// outline data can be accessed from the glyph slot (see note below).
	//
	// Note that by default the glyph loader doesn't render outlines into bitmaps. The following flags are used to modify
	// this default behaviour to more specific and useful cases.
	LoadDefault LoadFlag = C.FT_LOAD_DEFAULT
	// LoadNoScale don't scale the loaded outline glyph but keep it in font units.
	//
	// This flag implies LoadNoHinting and LoadNoBitmap, and unsets LoadRender.
	//
	// If the font is ‘tricky’ (see FaceFlagTricky for more), using LoadNoScale usually yields meaningless outlines
	// because the subglyphs must be scaled and positioned with hinting instructions. This can be solved by loading the
	// font without LoadNoScale and setting the character size to face.UnitsPerEM().
	LoadNoScale LoadFlag = C.FT_LOAD_NO_SCALE
	// LoadNoHinting disables hinting. This generally generates ‘blurrier’ bitmap glyphs when the glyph are rendered
	// in any of the anti-aliased modes. See also the note below.
	//
	// This flag is implied by LoadNoScale.
	LoadNoHinting LoadFlag = C.FT_LOAD_NO_HINTING
	// LoadRender call RenderGlyph after the glyph is loaded. By default, the glyph is rendered in RenderModeNormal mode.
	// This can be overridden by any LoadTarget or LoadMonochrome.
	//
	// This flag is unset by LoadNoScale.
	LoadRender LoadFlag = C.FT_LOAD_RENDER
	// LoadNoBitmap ignores bitmap strikes when loading. Bitmap-only fonts ignore this flag.
	//
	// LoadNoScale always sets this flag.
	LoadNoBitmap LoadFlag = C.FT_LOAD_NO_BITMAP
	// LoadVerticalLayout load the glyph for vertical text layout. In particular, the advance value in the GlyphSlot is
	// set to the VertAdvance value of the metrics field.
	//
	// If the face does not have FaceFlagVertical, you shouldn't use this flag currently. Reason is that in this case
	// vertical metrics get synthesized, and those values are not always consistent across various font formats.
	LoadVerticalLayout LoadFlag = C.FT_LOAD_VERTICAL_LAYOUT
	// LoadForceAutohint prefer the auto-hinter over the font's native hinter. See also the note below.
	LoadForceAutohint LoadFlag = C.FT_LOAD_FORCE_AUTOHINT
	// LoadPedantic makes the font driver perform pedantic verifications during glyph loading and hinting. This is
	// mostly used to detect broken glyphs in fonts. By default, FreeType tries to handle broken fonts also.
	//
	// In particular, errors from the TrueType bytecode engine are not passed to the application if this flag is not set;
	// this might result in partially hinted or distorted glyphs in case a glyph's bytecode is buggy.
	LoadPedantic LoadFlag = C.FT_LOAD_PEDANTIC
	// LoadNoRecurse don't load composite glyphs recursively. Instead, the font driver fills the NumSubglyph and
	// Subglyphs values of the glyph slot; it also sets glyph.Format to GlyphFormatComposite. The description of
	// subglyphs can then be accessed with GetSubGlyphInfo.
	//
	// Don't use this flag for retrieving metrics information since some font drivers only return rudimentary data.
	//
	// This flag implies LoadNoScale and LoadIgnoreTransform.
	LoadNoRecurse LoadFlag = C.FT_LOAD_NO_RECURSE
	// LoadIgnoreTransform ignore the transform matrix set by SetTransform.
	LoadIgnoreTransform LoadFlag = C.FT_LOAD_IGNORE_TRANSFORM
	// LoadMonochrome is used with LoadRender to indicate that you want to render an outline glyph to a 1-bit monochrome
	// bitmap glyph, with 8 pixels packed into each byte of the bitmap data.
	//
	// Note that this has no effect on the hinting algorithm used. You should rather use LoadTargetMono so that the
	// monochrome-optimized hinting algorithm is used.
	LoadMonochrome LoadFlag = C.FT_LOAD_MONOCHROME
	// LoadLinearDesign keep LinearHoriAdvance and LinearVertAdvance fields of GlyphSlot in font units. See GlyphSlot
	// for details.
	LoadLinearDesign LoadFlag = C.FT_LOAD_LINEAR_DESIGN
	// LoadNoAutohint disables the auto-hinter. See also the note below.
	LoadNoAutohint LoadFlag = C.FT_LOAD_NO_AUTOHINT
	// LoadColor loads colored glyphs. There are slight differences depending on the font format.
	//
	// [Since 2.5] Load embedded color bitmap images. The resulting color bitmaps, if available, will have the
	// PixelModeBGRA format, with pre-multiplied color channels. If the flag is not set and color bitmaps are found,
	// they are converted to 256-level gray bitmaps, using the PixelModeGray format.
	//
	// [Since 2.10, experimental] If the glyph index contains an entry in the face's ‘COLR’ table with a ‘CPAL’ palette
	// table (as defined in the OpenType specification), make RenderGlyph provide a default blending of the color glyph
	// layers associated with the glyph index, using the same bitmap format as embedded color bitmap images. This is
	// mainly for convenience; for full control of color layers use GetColorGlyphLayer and FreeType's color functions
	// like PaletteSelect instead of setting LoadColor for rendering so that the client application can handle blending
	// by itself.
	LoadColor LoadFlag = C.FT_LOAD_COLOR
	// LoadComputeMetrics [Since 2.6.1] Compute glyph metrics from the glyph data, without the use of bundled metrics
	// tables (for example, the ‘hdmx’ table in TrueType fonts). This flag is mainly used by font validating or font
	// editing applications, which need to ignore, verify, or edit those tables.
	//
	// Currently, this flag is only implemented for TrueType fonts.
	LoadComputeMetrics LoadFlag = C.FT_LOAD_COMPUTE_METRICS
	// LoadBitmapMetricsOnly [Since 2.7.1] request loading of the metrics and bitmap image information of a (possibly
	// embedded) bitmap glyph without allocating or copying the bitmap image data itself. No effect if the target glyph
	// is not a bitmap image.
	//
	// This flag unsets LoadRender.
	LoadBitmapMetricsOnly LoadFlag = C.FT_LOAD_BITMAP_METRICS_ONLY
	// LoadTargetNormal use the default hinting algorithm, optimized for
	// standard gray-level rendering. For monochrome output, use LoadTargetMono
	// instead.
	LoadTargetNormal LoadFlag = C.FT_LOAD_TARGET_NORMAL
	// LoadTargetLight is a lighter hinting algorithm for gray-level modes. Many
	// generated glyphs are fuzzier but better resemble their original shape.
	// This is achieved by snapping glyphs to the pixel grid only vertically
	// (Y-axis), as is done by FreeType's new CFF engine or Microsoft's
	// ClearType font renderer. This preserves inter-glyph spacing in horizontal
	// text. The snapping is done either by the native font driver, if the driver
	// itself and the font support it, or by the auto-hinter.
	//
	// Advance widths are rounded to integer values; however, using the LsbDelta
	// and RsbDelta fields of GlyphSlot, it is possible to get fractional
	// advance widths for subpixel positioning (which is recommended to use).
	//
	// If configuration option AF_CONFIG_OPTION_TT_SIZE_METRICS is active,
	// TrueType-like metrics are used to make this mode behave similarly as in
	// unpatched FreeType versions between 2.4.6 and 2.7.1 (inclusive).
	LoadTargetLight LoadFlag = C.FT_LOAD_TARGET_LIGHT
	// LoadTargetMono is a strong hinting algorithm that should only be used for
	// monochrome output. The result is probably unpleasant if the glyph is
	// rendered in non-monochrome modes.
	//
	// Note that for outline fonts only the TrueType font driver has proper
	// monochrome hinting support, provided the TTFs contain hints for B/W
	// rendering (which most fonts no longer provide). If these conditions are
	// not met it is very likely that you get ugly results at smaller sizes.
	LoadTargetMono LoadFlag = C.FT_LOAD_TARGET_MONO
	// LoadTargetLCD is a variant of LoadTargetLight optimized for horizontally
	// decimated LCD displays.
	LoadTargetLCD LoadFlag = C.FT_LOAD_TARGET_LCD
	// LoadTargetLCDV is a variant of LoadTargetNormal optimized for vertically
	// decimated LCD displays.
	LoadTargetLCDV LoadFlag = C.FT_LOAD_TARGET_LCD_V
)

func (x LoadFlag) String() string {
	// the maximum concatenated len, at the time of writing, is 237.
	s := make([]byte, 0, 237)

	if x == LoadDefault {
		return "Default"
	}

	if x&LoadNoScale == LoadNoScale {
		s = append(s, []byte("NoScale|")...)
	}
	if x&LoadNoHinting == LoadNoHinting {
		s = append(s, []byte("NoHinting|")...)
	}
	if x&LoadRender == LoadRender {
		s = append(s, []byte("Render|")...)
	}
	if x&LoadNoBitmap == LoadNoBitmap {
		s = append(s, []byte("NoBitmap|")...)
	}
	if x&LoadVerticalLayout == LoadVerticalLayout {
		s = append(s, []byte("VerticalLayout|")...)
	}
	if x&LoadForceAutohint == LoadForceAutohint {
		s = append(s, []byte("ForceAutohint|")...)
	}
	if x&LoadPedantic == LoadPedantic {
		s = append(s, []byte("Pedantic|")...)
	}
	if x&LoadNoRecurse == LoadNoRecurse {
		s = append(s, []byte("NoRecurse|")...)
	}
	if x&LoadIgnoreTransform == LoadIgnoreTransform {
		s = append(s, []byte("IgnoreTransform|")...)
	}
	if x&LoadMonochrome == LoadMonochrome {
		s = append(s, []byte("Monochrome|")...)
	}
	if x&LoadLinearDesign == LoadLinearDesign {
		s = append(s, []byte("LinearDesign|")...)
	}
	if x&LoadNoAutohint == LoadNoAutohint {
		s = append(s, []byte("NoAutohint|")...)
	}
	if x&LoadColor == LoadColor {
		s = append(s, []byte("Color|")...)
	}
	if x&LoadComputeMetrics == LoadComputeMetrics {
		s = append(s, []byte("ComputeMetrics|")...)
	}
	if x&LoadBitmapMetricsOnly == LoadBitmapMetricsOnly {
		s = append(s, []byte("BitmapMetricsOnly|")...)
	}

	switch {
	case x&LoadTargetLCDV == LoadTargetLCDV:
		s = append(s, []byte("TargetLCDV|")...)
	case x&LoadTargetLCD == LoadTargetLCD:
		s = append(s, []byte("TargetLCD|")...)
	case x&LoadTargetMono == LoadTargetMono:
		s = append(s, []byte("TargetMono|")...)
	case x&LoadTargetLight == LoadTargetLight:
		s = append(s, []byte("TargetLight|")...)
	default:
		s = append(s, []byte("TargetNormal|")...)
	}

	return string(s[:len(s)-1]) // trim the leading |
}

// FaceFlag is a list of bit flags of a given face.
// They inform client applications of properties of the corresponding face.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_face_flag_xxx
type FaceFlag int

const (
	// FaceFlagScalable the face contains outline glyphs. Note that a face can contain bitmap strikes also, i.e., a
	// face can have both this flag and FaceFlagFixedSizes set.
	FaceFlagScalable FaceFlag = C.FT_FACE_FLAG_SCALABLE

	// FaceFlagFixedSizes the face contains bitmap strikes. See also the Face.NumFixedSizes() and Face.AvailableSizes().
	FaceFlagFixedSizes FaceFlag = C.FT_FACE_FLAG_FIXED_SIZES

	// FaceFlagFixedWidth the face contains fixed-width characters (like Courier, Lucida, MonoType, etc.).
	FaceFlagFixedWidth FaceFlag = C.FT_FACE_FLAG_FIXED_WIDTH

	// FaceFlagSfnt the face uses the SFNT storage scheme. For now, this means TrueType and OpenType.
	FaceFlagSfnt FaceFlag = C.FT_FACE_FLAG_SFNT

	// FaceFlagHorizontal the face contains horizontal glyph metrics. This should be set for all common formats.
	FaceFlagHorizontal FaceFlag = C.FT_FACE_FLAG_HORIZONTAL

	// FaceFlagVertical the face contains vertical glyph metrics. This is only available in some formats, not all of them.
	FaceFlagVertical FaceFlag = C.FT_FACE_FLAG_VERTICAL

	// FaceFlagKerning the face contains kerning information. If set, the kerning distance can be retrieved using the
	// GetKerning() method.
	// Otherwise the function always returnS the vector (0,0). Note that FreeType doesn't handle kerning data from the
	// SFNT ‘GPOS’ table (as present in many OpenType fonts).
	FaceFlagKerning FaceFlag = C.FT_FACE_FLAG_KERNING

	// FaceFlagMultipleMasters the face contains multiple masters and is capable of interpolating between them.
	// Supported formats are Adobe MM, TrueType GX, and OpenType variation fonts.
	//
	// See https://www.freetype.org/freetype2/docs/reference/ft2-multiple_masters.html
	FaceFlagMultipleMasters FaceFlag = C.FT_FACE_FLAG_MULTIPLE_MASTERS

	// FaceFlagGlyphNames the face contains glyph names, which can be retrieved using GetGlyphName().
	// Note that some TrueType fonts contain broken glyph name tables. Use HasPSGlyphNames() when needed.
	FaceFlagGlyphNames FaceFlag = C.FT_FACE_FLAG_GLYPH_NAMES

	// FaceFlagHinter the font driver has a hinting machine of its own. For example, with TrueType fonts, it makes sense
	// to use data from the SFNT ‘gasp’ table only if the native TrueType hinting engine (with the bytecode interpreter)
	// is available and active.
	FaceFlagHinter FaceFlag = C.FT_FACE_FLAG_HINTER

	// FaceFlagCidKeyed The face is CID-keyed. In that case, the face is not accessed by glyph indices but by CID values.
	// For subsetted CID-keyed fonts this has the consequence that not all index values are a valid argument to
	// LoadGlyph(). Only the CID values for which corresponding glyphs in the subsetted font exist make LoadGlyph()
	// return successfully; in all other cases you get an ErrInvalidArgument error.
	//
	// Note that CID-keyed fonts that are in an SFNT wrapper (this is, all OpenType/CFF fonts) don't have this flag set
	// since the glyphs are accessed in the normal way (using contiguous indices); the ‘CID-ness’ isn't visible to the
	// application.
	FaceFlagCidKeyed FaceFlag = C.FT_FACE_FLAG_CID_KEYED

	// FaceFlagTricky the face is ‘tricky’, this is, it always needs the font format's native hinting engine to get a
	// reasonable result.
	// A typical example is the old Chinese font mingli.ttf (but not mingliu.ttc) that uses TrueType bytecode
	// instructions to move and scale all of its subglyphs.
	//
	// It is not possible to auto-hint such fonts using LoadForceAutohint; it will also ignore LoadNoHinting. You have
	// to set both LoadNoHinting and LoadNoAutohint to really disable hinting; however, you probably never want this
	// except for demonstration purposes.
	//
	// Currently, there are about a dozen TrueType fonts in the list of tricky fonts; they are hard-coded in file ttobjs.c.
	FaceFlagTricky FaceFlag = C.FT_FACE_FLAG_TRICKY

	// FaceFlagColor the face has color glyph tables. See LoadColor for more information.
	// [Since 2.5.1].
	FaceFlagColor FaceFlag = C.FT_FACE_FLAG_COLOR

	// FaceFlagVariation [Since 2.9] set if the current face (or named instance) has been altered with
	// SetMMDesignCoordinates, SetVarDesignCoordinates, or SetVarBlend_Coordinates.
	// This flag is unset by a call to SetNamedInstance.
	FaceFlagVariation FaceFlag = C.FT_FACE_FLAG_VARIATION
)

func (x FaceFlag) String() string {
	s := make([]byte, 0, 130) // len = sum of all the strings below.

	if x&FaceFlagScalable == FaceFlagScalable {
		s = append(s, []byte("Scalable|")...)
	}
	if x&FaceFlagFixedSizes == FaceFlagFixedSizes {
		s = append(s, []byte("FixedSizes|")...)
	}
	if x&FaceFlagFixedWidth == FaceFlagFixedWidth {
		s = append(s, []byte("FixedWidth|")...)
	}
	if x&FaceFlagSfnt == FaceFlagSfnt {
		s = append(s, []byte("Sfnt|")...)
	}
	if x&FaceFlagHorizontal == FaceFlagHorizontal {
		s = append(s, []byte("Horizontal|")...)
	}
	if x&FaceFlagVertical == FaceFlagVertical {
		s = append(s, []byte("Vertical|")...)
	}
	if x&FaceFlagKerning == FaceFlagKerning {
		s = append(s, []byte("Kerning|")...)
	}
	if x&FaceFlagMultipleMasters == FaceFlagMultipleMasters {
		s = append(s, []byte("MultipleMasters|")...)
	}
	if x&FaceFlagGlyphNames == FaceFlagGlyphNames {
		s = append(s, []byte("GlyphNames|")...)
	}
	if x&FaceFlagHinter == FaceFlagHinter {
		s = append(s, []byte("Hinter|")...)
	}
	if x&FaceFlagCidKeyed == FaceFlagCidKeyed {
		s = append(s, []byte("CidKeyed|")...)
	}
	if x&FaceFlagTricky == FaceFlagTricky {
		s = append(s, []byte("Tricky|")...)
	}
	if x&FaceFlagColor == FaceFlagColor {
		s = append(s, []byte("Color|")...)
	}
	if x&FaceFlagVariation == FaceFlagVariation {
		s = append(s, []byte("Variation|")...)
	}

	if len(s) == 0 {
		return ""
	}
	return string(s[:len(s)-1]) // trim the leading |
}

// StyleFlag is a list of bit flags to indicate the style of a given face.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-base_interface.html#ft_style_flag_xxx
type StyleFlag int

const (
	// StyleFlagItalic the face style is italic or oblique
	StyleFlagItalic StyleFlag = C.FT_STYLE_FLAG_ITALIC
	// StyleFlagBold the face is bold
	StyleFlagBold StyleFlag = C.FT_STYLE_FLAG_BOLD
)

func (x StyleFlag) String() string {
	s := make([]byte, 0, 12) // len = sum of all the strings below.

	if x&StyleFlagItalic == StyleFlagItalic {
		s = append(s, []byte("Italic|")...)
	}
	if x&StyleFlagBold == StyleFlagBold {
		s = append(s, []byte("Bold|")...)
	}

	if len(s) == 0 {
		return ""
	}
	return string(s[:len(s)-1]) // trim the leading |
}

// OutlineFlag is a list of bit-field constants used for the flags in an Outline flags field.
//
// NOTE:
// The flags OutlineIgnoreDropouts, OutlineSmartDropouts, and OutlineIncludeStubs
// are ignored by the smooth rasterizer.
//
// There exists a second mechanism to pass the drop-out mode to the B/W rasterizer;
// see the tags field in Outline.
//
// Please refer to the description of the ‘SCANTYPE’ instruction in the OpenType
// specification (in file ttinst1.doc) how simple drop-outs, smart drop-outs,
// and stubs are defined.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-outline_processing.html#ft_outline_xxx
type OutlineFlag uint

const (
	outlineNone OutlineFlag = C.FT_OUTLINE_NONE

	// OutlineOwner if set, this flag indicates that the outline's field arrays
	// (i.e., points, flags, and contours) are ‘owned’ by the outline object,
	// and should thus be freed when it is destroyed.
	OutlineOwner OutlineFlag = C.FT_OUTLINE_OWNER
	// OutlineEvenOddFill by default, outlines are filled using the non-zero
	// winding rule. If set to 1, the outline will be filled using the even-odd
	// fill rule (only works with the smooth rasterizer).
	OutlineEvenOddFill OutlineFlag = C.FT_OUTLINE_EVEN_ODD_FILL
	// OutlineReverseFill by default, outside contours of an outline are
	// oriented in clock-wise direction, as defined in the TrueType
	// specification. This flag is set if the outline uses the opposite
	// direction (typically for Type 1 fonts). This flag is ignored by the scan
	// converter.
	OutlineReverseFill OutlineFlag = C.FT_OUTLINE_REVERSE_FILL
	// OutlineIgnoreDropouts by default, the scan converter will try to detect
	// drop-outs in an outline and correct the glyph bitmap to ensure consistent
	// shape continuity. If set, this flag hints the scan-line converter to
	// ignore such cases.
	OutlineIgnoreDropouts OutlineFlag = C.FT_OUTLINE_IGNORE_DROPOUTS
	// OutlineSmartDropouts select smart dropout control. If unset, use simple
	// dropout control. Ignored if OutlineIgnoreDropouts is set.
	OutlineSmartDropouts OutlineFlag = C.FT_OUTLINE_SMART_DROPOUTS
	// OutlineIncludeStubs if set, turn pixels on for ‘stubs’, otherwise exclude
	// them. Ignored if OutlineIgnoreDropouts is set.
	OutlineIncludeStubs OutlineFlag = C.FT_OUTLINE_INCLUDE_STUBS
	// OutlineHighPrecision indicates that the scan-line converter should try to
	// convert this outline to bitmaps with the highest possible quality.
	// It is typically set for small character sizes. Note that this is only a
	// hint that might be completely ignored by a given scan-converter.
	OutlineHighPrecision OutlineFlag = C.FT_OUTLINE_HIGH_PRECISION
	// OutlineSinglePass is set to force a given scan-converter to only use a
	// single pass over the outline to render a bitmap glyph image. Normally,
	// it is set for very large character sizes. It is only a hint that might be
	// completely ignored by a given scan-converter.
	OutlineSinglePass OutlineFlag = C.FT_OUTLINE_SINGLE_PASS
)

func (x OutlineFlag) String() string {
	// the maximum concatenated len, at the time of writing, is 97.
	s := make([]byte, 0, 97)

	if x&OutlineOwner == OutlineOwner {
		s = append(s, []byte("Owner|")...)
	}
	if x&OutlineEvenOddFill == OutlineEvenOddFill {
		s = append(s, []byte("EvenOddFill|")...)
	}
	if x&OutlineReverseFill == OutlineReverseFill {
		s = append(s, []byte("ReverseFill|")...)
	}
	if x&OutlineIgnoreDropouts == OutlineIgnoreDropouts {
		s = append(s, []byte("IgnoreDropouts|")...)
	}
	if x&OutlineSmartDropouts == OutlineSmartDropouts {
		s = append(s, []byte("SmartDropouts|")...)
	}
	if x&OutlineIncludeStubs == OutlineIncludeStubs {
		s = append(s, []byte("IncludeStubs|")...)
	}
	if x&OutlineHighPrecision == OutlineHighPrecision {
		s = append(s, []byte("HighPrecision|")...)
	}
	if x&OutlineSinglePass == OutlineSinglePass {
		s = append(s, []byte("SinglePass|")...)
	}

	if len(s) == 0 {
		return ""
	}

	return string(s[:len(s)-1]) // trim the leading |
}
