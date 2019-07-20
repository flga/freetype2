package freetype2

// #include <stdlib.h>
// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_TRUETYPE_TABLES_H
// #include FT_TRUETYPE_IDS_H
import "C"

// PlatformID is an enum for the PlatformID in CharMap and SfntName structs.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_platform_xxx
type PlatformID int

const (
	// PlatformAppleUnicode is used by Apple to indicate a Unicode character map and/or name entry. See AppleEncodingIDs
	// for corresponding values. Note that name entries in this format are coded as big-endian UCS-2 character codes only.
	PlatformAppleUnicode PlatformID = C.TT_PLATFORM_APPLE_UNICODE
	// PlatformMacintosh is used by Apple to indicate a MacOS-specific charmap and/or name entry. See MacEncodingIDs for
	// corresponding values. Note that most TrueType fonts contain an Apple roman charmap to be usable on MacOS systems
	// (even if they contain a Microsoft charmap as well).
	PlatformMacintosh PlatformID = C.TT_PLATFORM_MACINTOSH
	// PlatformMicrosoft is used by Microsoft to indicate Windows-specific charmaps. See MicrosoftEncodingIDs for
	// corresponding values. Note that most fonts contain a Unicode charmap using (PlatformMicrosoft, EncodingMSIDUnicodeCS).
	PlatformMicrosoft PlatformID = C.TT_PLATFORM_MICROSOFT
	// PlatformCustom is used to indicate application-specific charmaps.
	PlatformCustom PlatformID = C.TT_PLATFORM_CUSTOM
	//PlatformAdobe isn't part of any font format specification, but is used by FreeType to report Adobe-specific
	// charmaps in a CharMap struct. See AdobeEncodingIDs.
	PlatformAdobe PlatformID = C.TT_PLATFORM_ADOBE
)

// EncodingID is an enum for EncodingID in CharMap and SfntName structs.
type EncodingID int

// AppleEncodingIDs are the valid encoding ids for PlatformAppleUnicode
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_apple_id_xxx
var AppleEncodingIDs = struct {
	// Unicode version 1.0.
	Default EncodingID
	// Unicode 1.1; specifies Hangul characters starting at U+34xx.
	Unicode1_1 EncodingID
	// Unicode 2.0 and beyond (UTF-16 BMP only).
	Unicode2_0 EncodingID
	// Unicode 3.1 and beyond, using UTF-32.
	Unicode32 EncodingID
	// From Adobe, not Apple. Not a normal cmap. Specifies variations on a real cmap.
	VariantSelector EncodingID
	// Used for fallback fonts that provide complete Unicode coverage with a type 13 cmap.
	FullUnicode EncodingID
}{
	Default:         EncodingID(C.TT_APPLE_ID_DEFAULT),
	Unicode1_1:      EncodingID(C.TT_APPLE_ID_UNICODE_1_1),
	Unicode2_0:      EncodingID(C.TT_APPLE_ID_UNICODE_2_0),
	Unicode32:       EncodingID(C.TT_APPLE_ID_UNICODE_32),
	VariantSelector: EncodingID(C.TT_APPLE_ID_VARIANT_SELECTOR),
	FullUnicode:     EncodingID(C.TT_APPLE_ID_FULL_UNICODE),
}

// MacEncodingIDs are the valid encoding ids for PlatformMacintosh
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_mac_id_xxx
var MacEncodingIDs = struct {
	Roman              EncodingID
	Japanese           EncodingID
	TraditionalChinese EncodingID
	Korean             EncodingID
	Arabic             EncodingID
	Hebrew             EncodingID
	Greek              EncodingID
	Russian            EncodingID
	Rsymbol            EncodingID
	Devanagari         EncodingID
	Gurmukhi           EncodingID
	Gujarati           EncodingID
	Oriya              EncodingID
	Bengali            EncodingID
	Tamil              EncodingID
	Telugu             EncodingID
	Kannada            EncodingID
	Malayalam          EncodingID
	Sinhalese          EncodingID
	Burmese            EncodingID
	Khmer              EncodingID
	Thai               EncodingID
	Laotian            EncodingID
	Georgian           EncodingID
	Armenian           EncodingID
	Maldivian          EncodingID
	SimplifiedChinese  EncodingID
	Tibetan            EncodingID
	Mongolian          EncodingID
	Geez               EncodingID
	Slavic             EncodingID
	Vietnamese         EncodingID
	Sindhi             EncodingID
	Uninterp           EncodingID
}{
	Roman:              EncodingID(C.TT_MAC_ID_ROMAN),
	Japanese:           EncodingID(C.TT_MAC_ID_JAPANESE),
	TraditionalChinese: EncodingID(C.TT_MAC_ID_TRADITIONAL_CHINESE),
	Korean:             EncodingID(C.TT_MAC_ID_KOREAN),
	Arabic:             EncodingID(C.TT_MAC_ID_ARABIC),
	Hebrew:             EncodingID(C.TT_MAC_ID_HEBREW),
	Greek:              EncodingID(C.TT_MAC_ID_GREEK),
	Russian:            EncodingID(C.TT_MAC_ID_RUSSIAN),
	Rsymbol:            EncodingID(C.TT_MAC_ID_RSYMBOL),
	Devanagari:         EncodingID(C.TT_MAC_ID_DEVANAGARI),
	Gurmukhi:           EncodingID(C.TT_MAC_ID_GURMUKHI),
	Gujarati:           EncodingID(C.TT_MAC_ID_GUJARATI),
	Oriya:              EncodingID(C.TT_MAC_ID_ORIYA),
	Bengali:            EncodingID(C.TT_MAC_ID_BENGALI),
	Tamil:              EncodingID(C.TT_MAC_ID_TAMIL),
	Telugu:             EncodingID(C.TT_MAC_ID_TELUGU),
	Kannada:            EncodingID(C.TT_MAC_ID_KANNADA),
	Malayalam:          EncodingID(C.TT_MAC_ID_MALAYALAM),
	Sinhalese:          EncodingID(C.TT_MAC_ID_SINHALESE),
	Burmese:            EncodingID(C.TT_MAC_ID_BURMESE),
	Khmer:              EncodingID(C.TT_MAC_ID_KHMER),
	Thai:               EncodingID(C.TT_MAC_ID_THAI),
	Laotian:            EncodingID(C.TT_MAC_ID_LAOTIAN),
	Georgian:           EncodingID(C.TT_MAC_ID_GEORGIAN),
	Armenian:           EncodingID(C.TT_MAC_ID_ARMENIAN),
	Maldivian:          EncodingID(C.TT_MAC_ID_MALDIVIAN),
	SimplifiedChinese:  EncodingID(C.TT_MAC_ID_SIMPLIFIED_CHINESE),
	Tibetan:            EncodingID(C.TT_MAC_ID_TIBETAN),
	Mongolian:          EncodingID(C.TT_MAC_ID_MONGOLIAN),
	Geez:               EncodingID(C.TT_MAC_ID_GEEZ),
	Slavic:             EncodingID(C.TT_MAC_ID_SLAVIC),
	Vietnamese:         EncodingID(C.TT_MAC_ID_VIETNAMESE),
	Sindhi:             EncodingID(C.TT_MAC_ID_SINDHI),
	Uninterp:           EncodingID(C.TT_MAC_ID_UNINTERP),
}

// MicrosoftEncodingIDs are the valid encoding ids for PlatformMicrosoft
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_ms_id_xxx
var MicrosoftEncodingIDs = struct {
	// Microsoft symbol encoding. See EncodingMsSymbol
	SymbolCs EncodingID
	// Microsoft WGL4 charmap, matching Unicode. See EncodingUnicode.
	UnicodeCs EncodingID
	// Shift JIS Japanese encoding. See EncodingSjis.
	Sjis EncodingID
	// Chinese encodings as used in the People's Republic of China (PRC). This means the encodings GB 2312 and its
	// supersets GBK and GB 18030. See  EncodingPrc.
	Prc EncodingID
	// Traditional Chinese as used in Taiwan and Hong Kong. See EncodingBig5.
	Big5 EncodingID
	// Korean Extended Wansung encoding. See EncodingWansung.
	Wansung EncodingID
	// Korean Johab encoding. See EncodingJohab.
	Johab EncodingID
	// UCS-4 or UTF-32 charmaps. This has been added to the OpenType specification version 1.4 (mid-2001).
	UCS4 EncodingID
}{
	SymbolCs:  EncodingID(C.TT_MS_ID_SYMBOL_CS),
	UnicodeCs: EncodingID(C.TT_MS_ID_UNICODE_CS),
	Sjis:      EncodingID(C.TT_MS_ID_SJIS),
	Prc:       EncodingID(C.TT_MS_ID_PRC),
	Big5:      EncodingID(C.TT_MS_ID_BIG_5),
	Wansung:   EncodingID(C.TT_MS_ID_WANSUNG),
	Johab:     EncodingID(C.TT_MS_ID_JOHAB),
	UCS4:      EncodingID(C.TT_MS_ID_UCS_4),
}

// AdobeEncodingIDs are the valid encoding ids for PlatformAdobe
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_adobe_id_xxx
var AdobeEncodingIDs = struct {
	Standard EncodingID
	Expert   EncodingID
	Custom   EncodingID
	Latin1   EncodingID
}{
	Standard: EncodingID(C.TT_ADOBE_ID_STANDARD),
	Expert:   EncodingID(C.TT_ADOBE_ID_EXPERT),
	Custom:   EncodingID(C.TT_ADOBE_ID_CUSTOM),
	Latin1:   EncodingID(C.TT_ADOBE_ID_LATIN_1),
}

// LanguageID is an enum of possible values of the language identifier field in the name records of the
// SFNT ‘name’ table.
type LanguageID int

// MacLangIDs are the valid language ids for PlatformMacintosh
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_mac_langid_xxx
var MacLangIDs = struct {
	English                   LanguageID
	French                    LanguageID
	German                    LanguageID
	Italian                   LanguageID
	Dutch                     LanguageID
	Swedish                   LanguageID
	Spanish                   LanguageID
	Danish                    LanguageID
	Portuguese                LanguageID
	Norwegian                 LanguageID
	Hebrew                    LanguageID
	Japanese                  LanguageID
	Arabic                    LanguageID
	Finnish                   LanguageID
	Greek                     LanguageID
	Icelandic                 LanguageID
	Maltese                   LanguageID
	Turkish                   LanguageID
	Croatian                  LanguageID
	ChineseTraditional        LanguageID
	Urdu                      LanguageID
	Hindi                     LanguageID
	Thai                      LanguageID
	Korean                    LanguageID
	Lithuanian                LanguageID
	Polish                    LanguageID
	Hungarian                 LanguageID
	Estonian                  LanguageID
	Lettish                   LanguageID
	Saamisk                   LanguageID
	Faeroese                  LanguageID
	Farsi                     LanguageID
	Russian                   LanguageID
	ChineseSimplified         LanguageID
	Flemish                   LanguageID
	Irish                     LanguageID
	Albanian                  LanguageID
	Romanian                  LanguageID
	Czech                     LanguageID
	Slovak                    LanguageID
	Slovenian                 LanguageID
	Yiddish                   LanguageID
	Serbian                   LanguageID
	Macedonian                LanguageID
	Bulgarian                 LanguageID
	Ukrainian                 LanguageID
	Byelorussian              LanguageID
	Uzbek                     LanguageID
	Kazakh                    LanguageID
	Azerbaijani               LanguageID
	AzerbaijaniCyrillicScript LanguageID
	AzerbaijaniArabicScript   LanguageID
	Armenian                  LanguageID
	Georgian                  LanguageID
	Moldavian                 LanguageID
	Kirghiz                   LanguageID
	Tajiki                    LanguageID
	Turkmen                   LanguageID
	Mongolian                 LanguageID
	MongolianMongolianScript  LanguageID
	MongolianCyrillicScript   LanguageID
	Pashto                    LanguageID
	Kurdish                   LanguageID
	Kashmiri                  LanguageID
	Sindhi                    LanguageID
	Tibetan                   LanguageID
	Nepali                    LanguageID
	Sanskrit                  LanguageID
	Marathi                   LanguageID
	Bengali                   LanguageID
	Assamese                  LanguageID
	Gujarati                  LanguageID
	Punjabi                   LanguageID
	Oriya                     LanguageID
	Malayalam                 LanguageID
	Kannada                   LanguageID
	Tamil                     LanguageID
	Telugu                    LanguageID
	Sinhalese                 LanguageID
	Burmese                   LanguageID
	Khmer                     LanguageID
	Lao                       LanguageID
	Vietnamese                LanguageID
	Indonesian                LanguageID
	Tagalog                   LanguageID
	MalayRomanScript          LanguageID
	MalayArabicScript         LanguageID
	Amharic                   LanguageID
	Tigrinya                  LanguageID
	Galla                     LanguageID
	Somali                    LanguageID
	Swahili                   LanguageID
	Ruanda                    LanguageID
	Rundi                     LanguageID
	Chewa                     LanguageID
	Malagasy                  LanguageID
	Esperanto                 LanguageID
	Welsh                     LanguageID
	Basque                    LanguageID
	Catalan                   LanguageID
	Latin                     LanguageID
	Quechua                   LanguageID
	Guarani                   LanguageID
	Aymara                    LanguageID
	Tatar                     LanguageID
	Uighur                    LanguageID
	Dzongkha                  LanguageID
	Javanese                  LanguageID
	Sundanese                 LanguageID
	Galician                  LanguageID
	Afrikaans                 LanguageID
	Breton                    LanguageID
	Inuktitut                 LanguageID
	ScottishGaelic            LanguageID
	ManxGaelic                LanguageID
	IrishGaelic               LanguageID
	Tongan                    LanguageID
	GreekPolytonic            LanguageID
	Greelandic                LanguageID
	AzerbaijaniRomanScript    LanguageID
}{
	English:                   LanguageID(C.TT_MAC_LANGID_ENGLISH),
	French:                    LanguageID(C.TT_MAC_LANGID_FRENCH),
	German:                    LanguageID(C.TT_MAC_LANGID_GERMAN),
	Italian:                   LanguageID(C.TT_MAC_LANGID_ITALIAN),
	Dutch:                     LanguageID(C.TT_MAC_LANGID_DUTCH),
	Swedish:                   LanguageID(C.TT_MAC_LANGID_SWEDISH),
	Spanish:                   LanguageID(C.TT_MAC_LANGID_SPANISH),
	Danish:                    LanguageID(C.TT_MAC_LANGID_DANISH),
	Portuguese:                LanguageID(C.TT_MAC_LANGID_PORTUGUESE),
	Norwegian:                 LanguageID(C.TT_MAC_LANGID_NORWEGIAN),
	Hebrew:                    LanguageID(C.TT_MAC_LANGID_HEBREW),
	Japanese:                  LanguageID(C.TT_MAC_LANGID_JAPANESE),
	Arabic:                    LanguageID(C.TT_MAC_LANGID_ARABIC),
	Finnish:                   LanguageID(C.TT_MAC_LANGID_FINNISH),
	Greek:                     LanguageID(C.TT_MAC_LANGID_GREEK),
	Icelandic:                 LanguageID(C.TT_MAC_LANGID_ICELANDIC),
	Maltese:                   LanguageID(C.TT_MAC_LANGID_MALTESE),
	Turkish:                   LanguageID(C.TT_MAC_LANGID_TURKISH),
	Croatian:                  LanguageID(C.TT_MAC_LANGID_CROATIAN),
	ChineseTraditional:        LanguageID(C.TT_MAC_LANGID_CHINESE_TRADITIONAL),
	Urdu:                      LanguageID(C.TT_MAC_LANGID_URDU),
	Hindi:                     LanguageID(C.TT_MAC_LANGID_HINDI),
	Thai:                      LanguageID(C.TT_MAC_LANGID_THAI),
	Korean:                    LanguageID(C.TT_MAC_LANGID_KOREAN),
	Lithuanian:                LanguageID(C.TT_MAC_LANGID_LITHUANIAN),
	Polish:                    LanguageID(C.TT_MAC_LANGID_POLISH),
	Hungarian:                 LanguageID(C.TT_MAC_LANGID_HUNGARIAN),
	Estonian:                  LanguageID(C.TT_MAC_LANGID_ESTONIAN),
	Lettish:                   LanguageID(C.TT_MAC_LANGID_LETTISH),
	Saamisk:                   LanguageID(C.TT_MAC_LANGID_SAAMISK),
	Faeroese:                  LanguageID(C.TT_MAC_LANGID_FAEROESE),
	Farsi:                     LanguageID(C.TT_MAC_LANGID_FARSI),
	Russian:                   LanguageID(C.TT_MAC_LANGID_RUSSIAN),
	ChineseSimplified:         LanguageID(C.TT_MAC_LANGID_CHINESE_SIMPLIFIED),
	Flemish:                   LanguageID(C.TT_MAC_LANGID_FLEMISH),
	Irish:                     LanguageID(C.TT_MAC_LANGID_IRISH),
	Albanian:                  LanguageID(C.TT_MAC_LANGID_ALBANIAN),
	Romanian:                  LanguageID(C.TT_MAC_LANGID_ROMANIAN),
	Czech:                     LanguageID(C.TT_MAC_LANGID_CZECH),
	Slovak:                    LanguageID(C.TT_MAC_LANGID_SLOVAK),
	Slovenian:                 LanguageID(C.TT_MAC_LANGID_SLOVENIAN),
	Yiddish:                   LanguageID(C.TT_MAC_LANGID_YIDDISH),
	Serbian:                   LanguageID(C.TT_MAC_LANGID_SERBIAN),
	Macedonian:                LanguageID(C.TT_MAC_LANGID_MACEDONIAN),
	Bulgarian:                 LanguageID(C.TT_MAC_LANGID_BULGARIAN),
	Ukrainian:                 LanguageID(C.TT_MAC_LANGID_UKRAINIAN),
	Byelorussian:              LanguageID(C.TT_MAC_LANGID_BYELORUSSIAN),
	Uzbek:                     LanguageID(C.TT_MAC_LANGID_UZBEK),
	Kazakh:                    LanguageID(C.TT_MAC_LANGID_KAZAKH),
	Azerbaijani:               LanguageID(C.TT_MAC_LANGID_AZERBAIJANI),
	AzerbaijaniCyrillicScript: LanguageID(C.TT_MAC_LANGID_AZERBAIJANI_CYRILLIC_SCRIPT),
	AzerbaijaniArabicScript:   LanguageID(C.TT_MAC_LANGID_AZERBAIJANI_ARABIC_SCRIPT),
	Armenian:                  LanguageID(C.TT_MAC_LANGID_ARMENIAN),
	Georgian:                  LanguageID(C.TT_MAC_LANGID_GEORGIAN),
	Moldavian:                 LanguageID(C.TT_MAC_LANGID_MOLDAVIAN),
	Kirghiz:                   LanguageID(C.TT_MAC_LANGID_KIRGHIZ),
	Tajiki:                    LanguageID(C.TT_MAC_LANGID_TAJIKI),
	Turkmen:                   LanguageID(C.TT_MAC_LANGID_TURKMEN),
	Mongolian:                 LanguageID(C.TT_MAC_LANGID_MONGOLIAN),
	MongolianMongolianScript:  LanguageID(C.TT_MAC_LANGID_MONGOLIAN_MONGOLIAN_SCRIPT),
	MongolianCyrillicScript:   LanguageID(C.TT_MAC_LANGID_MONGOLIAN_CYRILLIC_SCRIPT),
	Pashto:                    LanguageID(C.TT_MAC_LANGID_PASHTO),
	Kurdish:                   LanguageID(C.TT_MAC_LANGID_KURDISH),
	Kashmiri:                  LanguageID(C.TT_MAC_LANGID_KASHMIRI),
	Sindhi:                    LanguageID(C.TT_MAC_LANGID_SINDHI),
	Tibetan:                   LanguageID(C.TT_MAC_LANGID_TIBETAN),
	Nepali:                    LanguageID(C.TT_MAC_LANGID_NEPALI),
	Sanskrit:                  LanguageID(C.TT_MAC_LANGID_SANSKRIT),
	Marathi:                   LanguageID(C.TT_MAC_LANGID_MARATHI),
	Bengali:                   LanguageID(C.TT_MAC_LANGID_BENGALI),
	Assamese:                  LanguageID(C.TT_MAC_LANGID_ASSAMESE),
	Gujarati:                  LanguageID(C.TT_MAC_LANGID_GUJARATI),
	Punjabi:                   LanguageID(C.TT_MAC_LANGID_PUNJABI),
	Oriya:                     LanguageID(C.TT_MAC_LANGID_ORIYA),
	Malayalam:                 LanguageID(C.TT_MAC_LANGID_MALAYALAM),
	Kannada:                   LanguageID(C.TT_MAC_LANGID_KANNADA),
	Tamil:                     LanguageID(C.TT_MAC_LANGID_TAMIL),
	Telugu:                    LanguageID(C.TT_MAC_LANGID_TELUGU),
	Sinhalese:                 LanguageID(C.TT_MAC_LANGID_SINHALESE),
	Burmese:                   LanguageID(C.TT_MAC_LANGID_BURMESE),
	Khmer:                     LanguageID(C.TT_MAC_LANGID_KHMER),
	Lao:                       LanguageID(C.TT_MAC_LANGID_LAO),
	Vietnamese:                LanguageID(C.TT_MAC_LANGID_VIETNAMESE),
	Indonesian:                LanguageID(C.TT_MAC_LANGID_INDONESIAN),
	Tagalog:                   LanguageID(C.TT_MAC_LANGID_TAGALOG),
	MalayRomanScript:          LanguageID(C.TT_MAC_LANGID_MALAY_ROMAN_SCRIPT),
	MalayArabicScript:         LanguageID(C.TT_MAC_LANGID_MALAY_ARABIC_SCRIPT),
	Amharic:                   LanguageID(C.TT_MAC_LANGID_AMHARIC),
	Tigrinya:                  LanguageID(C.TT_MAC_LANGID_TIGRINYA),
	Galla:                     LanguageID(C.TT_MAC_LANGID_GALLA),
	Somali:                    LanguageID(C.TT_MAC_LANGID_SOMALI),
	Swahili:                   LanguageID(C.TT_MAC_LANGID_SWAHILI),
	Ruanda:                    LanguageID(C.TT_MAC_LANGID_RUANDA),
	Rundi:                     LanguageID(C.TT_MAC_LANGID_RUNDI),
	Chewa:                     LanguageID(C.TT_MAC_LANGID_CHEWA),
	Malagasy:                  LanguageID(C.TT_MAC_LANGID_MALAGASY),
	Esperanto:                 LanguageID(C.TT_MAC_LANGID_ESPERANTO),
	Welsh:                     LanguageID(C.TT_MAC_LANGID_WELSH),
	Basque:                    LanguageID(C.TT_MAC_LANGID_BASQUE),
	Catalan:                   LanguageID(C.TT_MAC_LANGID_CATALAN),
	Latin:                     LanguageID(C.TT_MAC_LANGID_LATIN),
	Quechua:                   LanguageID(C.TT_MAC_LANGID_QUECHUA),
	Guarani:                   LanguageID(C.TT_MAC_LANGID_GUARANI),
	Aymara:                    LanguageID(C.TT_MAC_LANGID_AYMARA),
	Tatar:                     LanguageID(C.TT_MAC_LANGID_TATAR),
	Uighur:                    LanguageID(C.TT_MAC_LANGID_UIGHUR),
	Dzongkha:                  LanguageID(C.TT_MAC_LANGID_DZONGKHA),
	Javanese:                  LanguageID(C.TT_MAC_LANGID_JAVANESE),
	Sundanese:                 LanguageID(C.TT_MAC_LANGID_SUNDANESE),
	Galician:                  LanguageID(C.TT_MAC_LANGID_GALICIAN),
	Afrikaans:                 LanguageID(C.TT_MAC_LANGID_AFRIKAANS),
	Breton:                    LanguageID(C.TT_MAC_LANGID_BRETON),
	Inuktitut:                 LanguageID(C.TT_MAC_LANGID_INUKTITUT),
	ScottishGaelic:            LanguageID(C.TT_MAC_LANGID_SCOTTISH_GAELIC),
	ManxGaelic:                LanguageID(C.TT_MAC_LANGID_MANX_GAELIC),
	IrishGaelic:               LanguageID(C.TT_MAC_LANGID_IRISH_GAELIC),
	Tongan:                    LanguageID(C.TT_MAC_LANGID_TONGAN),
	GreekPolytonic:            LanguageID(C.TT_MAC_LANGID_GREEK_POLYTONIC),
	Greelandic:                LanguageID(C.TT_MAC_LANGID_GREELANDIC),
	AzerbaijaniRomanScript:    LanguageID(C.TT_MAC_LANGID_AZERBAIJANI_ROMAN_SCRIPT),
}

// MicrosoftLangIDs are the valid language ids for PlatformMicrosoft
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_ms_langid_xxx
var MicrosoftLangIDs = struct {
	ArabicSaudiArabia           LanguageID
	ArabicIraq                  LanguageID
	ArabicEgypt                 LanguageID
	ArabicLibya                 LanguageID
	ArabicAlgeria               LanguageID
	ArabicMorocco               LanguageID
	ArabicTunisia               LanguageID
	ArabicOman                  LanguageID
	ArabicYemen                 LanguageID
	ArabicSyria                 LanguageID
	ArabicJordan                LanguageID
	ArabicLebanon               LanguageID
	ArabicKuwait                LanguageID
	ArabicUae                   LanguageID
	ArabicBahrain               LanguageID
	ArabicQatar                 LanguageID
	BulgarianBulgaria           LanguageID
	CatalanCatalan              LanguageID
	ChineseTaiwan               LanguageID
	ChinesePrc                  LanguageID
	ChineseHongKong             LanguageID
	ChineseSingapore            LanguageID
	ChineseMacao                LanguageID
	CzechCzechRepublic          LanguageID
	DanishDenmark               LanguageID
	GermanGermany               LanguageID
	GermanSwitzerland           LanguageID
	GermanAustria               LanguageID
	GermanLuxembourg            LanguageID
	GermanLiechtenstein         LanguageID
	GreekGreece                 LanguageID
	EnglishUnitedStates         LanguageID
	EnglishUnitedKingdom        LanguageID
	EnglishAustralia            LanguageID
	EnglishCanada               LanguageID
	EnglishNewZealand           LanguageID
	EnglishIreland              LanguageID
	EnglishSouthAfrica          LanguageID
	EnglishJamaica              LanguageID
	EnglishCaribbean            LanguageID
	EnglishBelize               LanguageID
	EnglishTrinidad             LanguageID
	EnglishZimbabwe             LanguageID
	EnglishPhilippines          LanguageID
	EnglishIndia                LanguageID
	EnglishMalaysia             LanguageID
	EnglishSingapore            LanguageID
	SpanishSpainTraditionalSort LanguageID
	SpanishMexico               LanguageID
	SpanishSpainModernSort      LanguageID
	SpanishGuatemala            LanguageID
	SpanishCostaRica            LanguageID
	SpanishPanama               LanguageID
	SpanishDominicanRepublic    LanguageID
	SpanishVenezuela            LanguageID
	SpanishColombia             LanguageID
	SpanishPeru                 LanguageID
	SpanishArgentina            LanguageID
	SpanishEcuador              LanguageID
	SpanishChile                LanguageID
	SpanishUruguay              LanguageID
	SpanishParaguay             LanguageID
	SpanishBolivia              LanguageID
	SpanishElSalvador           LanguageID
	SpanishHonduras             LanguageID
	SpanishNicaragua            LanguageID
	SpanishPuertoRico           LanguageID
	SpanishUnitedStates         LanguageID
	FinnishFinland              LanguageID
	FrenchFrance                LanguageID
	FrenchBelgium               LanguageID
	FrenchCanada                LanguageID
	FrenchSwitzerland           LanguageID
	FrenchLuxembourg            LanguageID
	FrenchMonaco                LanguageID
	HebrewIsrael                LanguageID
	HungarianHungary            LanguageID
	IcelandicIceland            LanguageID
	ItalianItaly                LanguageID
	ItalianSwitzerland          LanguageID
	JapaneseJapan               LanguageID
	KoreanKorea                 LanguageID
	DutchNetherlands            LanguageID
	DutchBelgium                LanguageID
	NorwegianNorwayBokmal       LanguageID
	NorwegianNorwayNynorsk      LanguageID
	PolishPoland                LanguageID
	PortugueseBrazil            LanguageID
	PortuguesePortugal          LanguageID
	RomanshSwitzerland          LanguageID
	RomanianRomania             LanguageID
	RussianRussia               LanguageID
	CroatianCroatia             LanguageID
	SerbianSerbiaLatin          LanguageID
	SerbianSerbiaCyrillic       LanguageID
	CroatianBosniaHerzegovina   LanguageID
	BosnianBosniaHerzegovina    LanguageID
	SerbianBosniaHerzLatin      LanguageID
	SerbianBosniaHerzCyrillic   LanguageID
	BosnianBosniaHerzCyrillic   LanguageID
	SlovakSlovakia              LanguageID
	AlbanianAlbania             LanguageID
	SwedishSweden               LanguageID
	SwedishFinland              LanguageID
	ThaiThailand                LanguageID
	TurkishTurkey               LanguageID
	UrduPakistan                LanguageID
	IndonesianIndonesia         LanguageID
	UkrainianUkraine            LanguageID
	BelarusianBelarus           LanguageID
	SlovenianSlovenia           LanguageID
	EstonianEstonia             LanguageID
	LatvianLatvia               LanguageID
	LithuanianLithuania         LanguageID
	TajikTajikistan             LanguageID
	VietnameseVietNam           LanguageID
	ArmenianArmenia             LanguageID
	AzeriAzerbaijanLatin        LanguageID
	AzeriAzerbaijanCyrillic     LanguageID
	BasqueBasque                LanguageID
	UpperSorbianGermany         LanguageID
	LowerSorbianGermany         LanguageID
	MacedonianMacedonia         LanguageID
	SetswanaSouthAfrica         LanguageID
	IsixhosaSouthAfrica         LanguageID
	IsizuluSouthAfrica          LanguageID
	AfrikaansSouthAfrica        LanguageID
	GeorgianGeorgia             LanguageID
	FaeroeseFaeroeIslands       LanguageID
	HindiIndia                  LanguageID
	MalteseMalta                LanguageID
	SamiNorthernNorway          LanguageID
	SamiNorthernSweden          LanguageID
	SamiNorthernFinland         LanguageID
	SamiLuleNorway              LanguageID
	SamiLuleSweden              LanguageID
	SamiSouthernNorway          LanguageID
	SamiSouthernSweden          LanguageID
	SamiSkoltFinland            LanguageID
	SamiInariFinland            LanguageID
	IrishIreland                LanguageID
	MalayMalaysia               LanguageID
	MalayBruneiDarussalam       LanguageID
	KazakhKazakhstan            LanguageID
	KyrgyzKyrgyzstan            LanguageID
	KiswahiliKenya              LanguageID
	TurkmenTurkmenistan         LanguageID
	UzbekUzbekistanLatin        LanguageID
	UzbekUzbekistanCyrillic     LanguageID
	TatarRussia                 LanguageID
	BengaliIndia                LanguageID
	BengaliBangladesh           LanguageID
	PunjabiIndia                LanguageID
	GujaratiIndia               LanguageID
	OdiaIndia                   LanguageID
	TamilIndia                  LanguageID
	TeluguIndia                 LanguageID
	KannadaIndia                LanguageID
	MalayalamIndia              LanguageID
	AssameseIndia               LanguageID
	MarathiIndia                LanguageID
	SanskritIndia               LanguageID
	MongolianMongolia           LanguageID
	MongolianPrc                LanguageID
	TibetanPrc                  LanguageID
	WelshUnitedKingdom          LanguageID
	KhmerCambodia               LanguageID
	LaoLaos                     LanguageID
	GalicianGalician            LanguageID
	KonkaniIndia                LanguageID
	SyriacSyria                 LanguageID
	SinhalaSriLanka             LanguageID
	InuktitutCanada             LanguageID
	InuktitutCanadaLatin        LanguageID
	AmharicEthiopia             LanguageID
	TamazightAlgeria            LanguageID
	NepaliNepal                 LanguageID
	FrisianNetherlands          LanguageID
	PashtoAfghanistan           LanguageID
	FilipinoPhilippines         LanguageID
	DhivehiMaldives             LanguageID
	HausaNigeria                LanguageID
	YorubaNigeria               LanguageID
	QuechuaBolivia              LanguageID
	QuechuaEcuador              LanguageID
	QuechuaPeru                 LanguageID
	SesothoSaLeboaSouthAfrica   LanguageID
	BashkirRussia               LanguageID
	LuxembourgishLuxembourg     LanguageID
	GreenlandicGreenland        LanguageID
	IgboNigeria                 LanguageID
	YiPrc                       LanguageID
	MapudungunChile             LanguageID
	MohawkMohawk                LanguageID
	BretonFrance                LanguageID
	UighurPrc                   LanguageID
	MaoriNewZealand             LanguageID
	OccitanFrance               LanguageID
	CorsicanFrance              LanguageID
	AlsatianFrance              LanguageID
	YakutRussia                 LanguageID
	KicheGuatemala              LanguageID
	KinyarwandaRwanda           LanguageID
	WolofSenegal                LanguageID
	DariAfghanistan             LanguageID
}{
	ArabicSaudiArabia:           LanguageID(C.TT_MS_LANGID_ARABIC_SAUDI_ARABIA),
	ArabicIraq:                  LanguageID(C.TT_MS_LANGID_ARABIC_IRAQ),
	ArabicEgypt:                 LanguageID(C.TT_MS_LANGID_ARABIC_EGYPT),
	ArabicLibya:                 LanguageID(C.TT_MS_LANGID_ARABIC_LIBYA),
	ArabicAlgeria:               LanguageID(C.TT_MS_LANGID_ARABIC_ALGERIA),
	ArabicMorocco:               LanguageID(C.TT_MS_LANGID_ARABIC_MOROCCO),
	ArabicTunisia:               LanguageID(C.TT_MS_LANGID_ARABIC_TUNISIA),
	ArabicOman:                  LanguageID(C.TT_MS_LANGID_ARABIC_OMAN),
	ArabicYemen:                 LanguageID(C.TT_MS_LANGID_ARABIC_YEMEN),
	ArabicSyria:                 LanguageID(C.TT_MS_LANGID_ARABIC_SYRIA),
	ArabicJordan:                LanguageID(C.TT_MS_LANGID_ARABIC_JORDAN),
	ArabicLebanon:               LanguageID(C.TT_MS_LANGID_ARABIC_LEBANON),
	ArabicKuwait:                LanguageID(C.TT_MS_LANGID_ARABIC_KUWAIT),
	ArabicUae:                   LanguageID(C.TT_MS_LANGID_ARABIC_UAE),
	ArabicBahrain:               LanguageID(C.TT_MS_LANGID_ARABIC_BAHRAIN),
	ArabicQatar:                 LanguageID(C.TT_MS_LANGID_ARABIC_QATAR),
	BulgarianBulgaria:           LanguageID(C.TT_MS_LANGID_BULGARIAN_BULGARIA),
	CatalanCatalan:              LanguageID(C.TT_MS_LANGID_CATALAN_CATALAN),
	ChineseTaiwan:               LanguageID(C.TT_MS_LANGID_CHINESE_TAIWAN),
	ChinesePrc:                  LanguageID(C.TT_MS_LANGID_CHINESE_PRC),
	ChineseHongKong:             LanguageID(C.TT_MS_LANGID_CHINESE_HONG_KONG),
	ChineseSingapore:            LanguageID(C.TT_MS_LANGID_CHINESE_SINGAPORE),
	ChineseMacao:                LanguageID(C.TT_MS_LANGID_CHINESE_MACAO),
	CzechCzechRepublic:          LanguageID(C.TT_MS_LANGID_CZECH_CZECH_REPUBLIC),
	DanishDenmark:               LanguageID(C.TT_MS_LANGID_DANISH_DENMARK),
	GermanGermany:               LanguageID(C.TT_MS_LANGID_GERMAN_GERMANY),
	GermanSwitzerland:           LanguageID(C.TT_MS_LANGID_GERMAN_SWITZERLAND),
	GermanAustria:               LanguageID(C.TT_MS_LANGID_GERMAN_AUSTRIA),
	GermanLuxembourg:            LanguageID(C.TT_MS_LANGID_GERMAN_LUXEMBOURG),
	GermanLiechtenstein:         LanguageID(C.TT_MS_LANGID_GERMAN_LIECHTENSTEIN),
	GreekGreece:                 LanguageID(C.TT_MS_LANGID_GREEK_GREECE),
	EnglishUnitedStates:         LanguageID(C.TT_MS_LANGID_ENGLISH_UNITED_STATES),
	EnglishUnitedKingdom:        LanguageID(C.TT_MS_LANGID_ENGLISH_UNITED_KINGDOM),
	EnglishAustralia:            LanguageID(C.TT_MS_LANGID_ENGLISH_AUSTRALIA),
	EnglishCanada:               LanguageID(C.TT_MS_LANGID_ENGLISH_CANADA),
	EnglishNewZealand:           LanguageID(C.TT_MS_LANGID_ENGLISH_NEW_ZEALAND),
	EnglishIreland:              LanguageID(C.TT_MS_LANGID_ENGLISH_IRELAND),
	EnglishSouthAfrica:          LanguageID(C.TT_MS_LANGID_ENGLISH_SOUTH_AFRICA),
	EnglishJamaica:              LanguageID(C.TT_MS_LANGID_ENGLISH_JAMAICA),
	EnglishCaribbean:            LanguageID(C.TT_MS_LANGID_ENGLISH_CARIBBEAN),
	EnglishBelize:               LanguageID(C.TT_MS_LANGID_ENGLISH_BELIZE),
	EnglishTrinidad:             LanguageID(C.TT_MS_LANGID_ENGLISH_TRINIDAD),
	EnglishZimbabwe:             LanguageID(C.TT_MS_LANGID_ENGLISH_ZIMBABWE),
	EnglishPhilippines:          LanguageID(C.TT_MS_LANGID_ENGLISH_PHILIPPINES),
	EnglishIndia:                LanguageID(C.TT_MS_LANGID_ENGLISH_INDIA),
	EnglishMalaysia:             LanguageID(C.TT_MS_LANGID_ENGLISH_MALAYSIA),
	EnglishSingapore:            LanguageID(C.TT_MS_LANGID_ENGLISH_SINGAPORE),
	SpanishSpainTraditionalSort: LanguageID(C.TT_MS_LANGID_SPANISH_SPAIN_TRADITIONAL_SORT),
	SpanishMexico:               LanguageID(C.TT_MS_LANGID_SPANISH_MEXICO),
	SpanishSpainModernSort:      LanguageID(C.TT_MS_LANGID_SPANISH_SPAIN_MODERN_SORT),
	SpanishGuatemala:            LanguageID(C.TT_MS_LANGID_SPANISH_GUATEMALA),
	SpanishCostaRica:            LanguageID(C.TT_MS_LANGID_SPANISH_COSTA_RICA),
	SpanishPanama:               LanguageID(C.TT_MS_LANGID_SPANISH_PANAMA),
	SpanishDominicanRepublic:    LanguageID(C.TT_MS_LANGID_SPANISH_DOMINICAN_REPUBLIC),
	SpanishVenezuela:            LanguageID(C.TT_MS_LANGID_SPANISH_VENEZUELA),
	SpanishColombia:             LanguageID(C.TT_MS_LANGID_SPANISH_COLOMBIA),
	SpanishPeru:                 LanguageID(C.TT_MS_LANGID_SPANISH_PERU),
	SpanishArgentina:            LanguageID(C.TT_MS_LANGID_SPANISH_ARGENTINA),
	SpanishEcuador:              LanguageID(C.TT_MS_LANGID_SPANISH_ECUADOR),
	SpanishChile:                LanguageID(C.TT_MS_LANGID_SPANISH_CHILE),
	SpanishUruguay:              LanguageID(C.TT_MS_LANGID_SPANISH_URUGUAY),
	SpanishParaguay:             LanguageID(C.TT_MS_LANGID_SPANISH_PARAGUAY),
	SpanishBolivia:              LanguageID(C.TT_MS_LANGID_SPANISH_BOLIVIA),
	SpanishElSalvador:           LanguageID(C.TT_MS_LANGID_SPANISH_EL_SALVADOR),
	SpanishHonduras:             LanguageID(C.TT_MS_LANGID_SPANISH_HONDURAS),
	SpanishNicaragua:            LanguageID(C.TT_MS_LANGID_SPANISH_NICARAGUA),
	SpanishPuertoRico:           LanguageID(C.TT_MS_LANGID_SPANISH_PUERTO_RICO),
	SpanishUnitedStates:         LanguageID(C.TT_MS_LANGID_SPANISH_UNITED_STATES),
	FinnishFinland:              LanguageID(C.TT_MS_LANGID_FINNISH_FINLAND),
	FrenchFrance:                LanguageID(C.TT_MS_LANGID_FRENCH_FRANCE),
	FrenchBelgium:               LanguageID(C.TT_MS_LANGID_FRENCH_BELGIUM),
	FrenchCanada:                LanguageID(C.TT_MS_LANGID_FRENCH_CANADA),
	FrenchSwitzerland:           LanguageID(C.TT_MS_LANGID_FRENCH_SWITZERLAND),
	FrenchLuxembourg:            LanguageID(C.TT_MS_LANGID_FRENCH_LUXEMBOURG),
	FrenchMonaco:                LanguageID(C.TT_MS_LANGID_FRENCH_MONACO),
	HebrewIsrael:                LanguageID(C.TT_MS_LANGID_HEBREW_ISRAEL),
	HungarianHungary:            LanguageID(C.TT_MS_LANGID_HUNGARIAN_HUNGARY),
	IcelandicIceland:            LanguageID(C.TT_MS_LANGID_ICELANDIC_ICELAND),
	ItalianItaly:                LanguageID(C.TT_MS_LANGID_ITALIAN_ITALY),
	ItalianSwitzerland:          LanguageID(C.TT_MS_LANGID_ITALIAN_SWITZERLAND),
	JapaneseJapan:               LanguageID(C.TT_MS_LANGID_JAPANESE_JAPAN),
	KoreanKorea:                 LanguageID(C.TT_MS_LANGID_KOREAN_KOREA),
	DutchNetherlands:            LanguageID(C.TT_MS_LANGID_DUTCH_NETHERLANDS),
	DutchBelgium:                LanguageID(C.TT_MS_LANGID_DUTCH_BELGIUM),
	NorwegianNorwayBokmal:       LanguageID(C.TT_MS_LANGID_NORWEGIAN_NORWAY_BOKMAL),
	NorwegianNorwayNynorsk:      LanguageID(C.TT_MS_LANGID_NORWEGIAN_NORWAY_NYNORSK),
	PolishPoland:                LanguageID(C.TT_MS_LANGID_POLISH_POLAND),
	PortugueseBrazil:            LanguageID(C.TT_MS_LANGID_PORTUGUESE_BRAZIL),
	PortuguesePortugal:          LanguageID(C.TT_MS_LANGID_PORTUGUESE_PORTUGAL),
	RomanshSwitzerland:          LanguageID(C.TT_MS_LANGID_ROMANSH_SWITZERLAND),
	RomanianRomania:             LanguageID(C.TT_MS_LANGID_ROMANIAN_ROMANIA),
	RussianRussia:               LanguageID(C.TT_MS_LANGID_RUSSIAN_RUSSIA),
	CroatianCroatia:             LanguageID(C.TT_MS_LANGID_CROATIAN_CROATIA),
	SerbianSerbiaLatin:          LanguageID(C.TT_MS_LANGID_SERBIAN_SERBIA_LATIN),
	SerbianSerbiaCyrillic:       LanguageID(C.TT_MS_LANGID_SERBIAN_SERBIA_CYRILLIC),
	CroatianBosniaHerzegovina:   LanguageID(C.TT_MS_LANGID_CROATIAN_BOSNIA_HERZEGOVINA),
	BosnianBosniaHerzegovina:    LanguageID(C.TT_MS_LANGID_BOSNIAN_BOSNIA_HERZEGOVINA),
	SerbianBosniaHerzLatin:      LanguageID(C.TT_MS_LANGID_SERBIAN_BOSNIA_HERZ_LATIN),
	SerbianBosniaHerzCyrillic:   LanguageID(C.TT_MS_LANGID_SERBIAN_BOSNIA_HERZ_CYRILLIC),
	BosnianBosniaHerzCyrillic:   LanguageID(C.TT_MS_LANGID_BOSNIAN_BOSNIA_HERZ_CYRILLIC),
	SlovakSlovakia:              LanguageID(C.TT_MS_LANGID_SLOVAK_SLOVAKIA),
	AlbanianAlbania:             LanguageID(C.TT_MS_LANGID_ALBANIAN_ALBANIA),
	SwedishSweden:               LanguageID(C.TT_MS_LANGID_SWEDISH_SWEDEN),
	SwedishFinland:              LanguageID(C.TT_MS_LANGID_SWEDISH_FINLAND),
	ThaiThailand:                LanguageID(C.TT_MS_LANGID_THAI_THAILAND),
	TurkishTurkey:               LanguageID(C.TT_MS_LANGID_TURKISH_TURKEY),
	UrduPakistan:                LanguageID(C.TT_MS_LANGID_URDU_PAKISTAN),
	IndonesianIndonesia:         LanguageID(C.TT_MS_LANGID_INDONESIAN_INDONESIA),
	UkrainianUkraine:            LanguageID(C.TT_MS_LANGID_UKRAINIAN_UKRAINE),
	BelarusianBelarus:           LanguageID(C.TT_MS_LANGID_BELARUSIAN_BELARUS),
	SlovenianSlovenia:           LanguageID(C.TT_MS_LANGID_SLOVENIAN_SLOVENIA),
	EstonianEstonia:             LanguageID(C.TT_MS_LANGID_ESTONIAN_ESTONIA),
	LatvianLatvia:               LanguageID(C.TT_MS_LANGID_LATVIAN_LATVIA),
	LithuanianLithuania:         LanguageID(C.TT_MS_LANGID_LITHUANIAN_LITHUANIA),
	TajikTajikistan:             LanguageID(C.TT_MS_LANGID_TAJIK_TAJIKISTAN),
	VietnameseVietNam:           LanguageID(C.TT_MS_LANGID_VIETNAMESE_VIET_NAM),
	ArmenianArmenia:             LanguageID(C.TT_MS_LANGID_ARMENIAN_ARMENIA),
	AzeriAzerbaijanLatin:        LanguageID(C.TT_MS_LANGID_AZERI_AZERBAIJAN_LATIN),
	AzeriAzerbaijanCyrillic:     LanguageID(C.TT_MS_LANGID_AZERI_AZERBAIJAN_CYRILLIC),
	BasqueBasque:                LanguageID(C.TT_MS_LANGID_BASQUE_BASQUE),
	UpperSorbianGermany:         LanguageID(C.TT_MS_LANGID_UPPER_SORBIAN_GERMANY),
	LowerSorbianGermany:         LanguageID(C.TT_MS_LANGID_LOWER_SORBIAN_GERMANY),
	MacedonianMacedonia:         LanguageID(C.TT_MS_LANGID_MACEDONIAN_MACEDONIA),
	SetswanaSouthAfrica:         LanguageID(C.TT_MS_LANGID_SETSWANA_SOUTH_AFRICA),
	IsixhosaSouthAfrica:         LanguageID(C.TT_MS_LANGID_ISIXHOSA_SOUTH_AFRICA),
	IsizuluSouthAfrica:          LanguageID(C.TT_MS_LANGID_ISIZULU_SOUTH_AFRICA),
	AfrikaansSouthAfrica:        LanguageID(C.TT_MS_LANGID_AFRIKAANS_SOUTH_AFRICA),
	GeorgianGeorgia:             LanguageID(C.TT_MS_LANGID_GEORGIAN_GEORGIA),
	FaeroeseFaeroeIslands:       LanguageID(C.TT_MS_LANGID_FAEROESE_FAEROE_ISLANDS),
	HindiIndia:                  LanguageID(C.TT_MS_LANGID_HINDI_INDIA),
	MalteseMalta:                LanguageID(C.TT_MS_LANGID_MALTESE_MALTA),
	SamiNorthernNorway:          LanguageID(C.TT_MS_LANGID_SAMI_NORTHERN_NORWAY),
	SamiNorthernSweden:          LanguageID(C.TT_MS_LANGID_SAMI_NORTHERN_SWEDEN),
	SamiNorthernFinland:         LanguageID(C.TT_MS_LANGID_SAMI_NORTHERN_FINLAND),
	SamiLuleNorway:              LanguageID(C.TT_MS_LANGID_SAMI_LULE_NORWAY),
	SamiLuleSweden:              LanguageID(C.TT_MS_LANGID_SAMI_LULE_SWEDEN),
	SamiSouthernNorway:          LanguageID(C.TT_MS_LANGID_SAMI_SOUTHERN_NORWAY),
	SamiSouthernSweden:          LanguageID(C.TT_MS_LANGID_SAMI_SOUTHERN_SWEDEN),
	SamiSkoltFinland:            LanguageID(C.TT_MS_LANGID_SAMI_SKOLT_FINLAND),
	SamiInariFinland:            LanguageID(C.TT_MS_LANGID_SAMI_INARI_FINLAND),
	IrishIreland:                LanguageID(C.TT_MS_LANGID_IRISH_IRELAND),
	MalayMalaysia:               LanguageID(C.TT_MS_LANGID_MALAY_MALAYSIA),
	MalayBruneiDarussalam:       LanguageID(C.TT_MS_LANGID_MALAY_BRUNEI_DARUSSALAM),
	KazakhKazakhstan:            LanguageID(C.TT_MS_LANGID_KAZAKH_KAZAKHSTAN),
	KyrgyzKyrgyzstan:            LanguageID(C.TT_MS_LANGID_KYRGYZ_KYRGYZSTAN),
	KiswahiliKenya:              LanguageID(C.TT_MS_LANGID_KISWAHILI_KENYA),
	TurkmenTurkmenistan:         LanguageID(C.TT_MS_LANGID_TURKMEN_TURKMENISTAN),
	UzbekUzbekistanLatin:        LanguageID(C.TT_MS_LANGID_UZBEK_UZBEKISTAN_LATIN),
	UzbekUzbekistanCyrillic:     LanguageID(C.TT_MS_LANGID_UZBEK_UZBEKISTAN_CYRILLIC),
	TatarRussia:                 LanguageID(C.TT_MS_LANGID_TATAR_RUSSIA),
	BengaliIndia:                LanguageID(C.TT_MS_LANGID_BENGALI_INDIA),
	BengaliBangladesh:           LanguageID(C.TT_MS_LANGID_BENGALI_BANGLADESH),
	PunjabiIndia:                LanguageID(C.TT_MS_LANGID_PUNJABI_INDIA),
	GujaratiIndia:               LanguageID(C.TT_MS_LANGID_GUJARATI_INDIA),
	OdiaIndia:                   LanguageID(C.TT_MS_LANGID_ODIA_INDIA),
	TamilIndia:                  LanguageID(C.TT_MS_LANGID_TAMIL_INDIA),
	TeluguIndia:                 LanguageID(C.TT_MS_LANGID_TELUGU_INDIA),
	KannadaIndia:                LanguageID(C.TT_MS_LANGID_KANNADA_INDIA),
	MalayalamIndia:              LanguageID(C.TT_MS_LANGID_MALAYALAM_INDIA),
	AssameseIndia:               LanguageID(C.TT_MS_LANGID_ASSAMESE_INDIA),
	MarathiIndia:                LanguageID(C.TT_MS_LANGID_MARATHI_INDIA),
	SanskritIndia:               LanguageID(C.TT_MS_LANGID_SANSKRIT_INDIA),
	MongolianMongolia:           LanguageID(C.TT_MS_LANGID_MONGOLIAN_MONGOLIA),
	MongolianPrc:                LanguageID(C.TT_MS_LANGID_MONGOLIAN_PRC),
	TibetanPrc:                  LanguageID(C.TT_MS_LANGID_TIBETAN_PRC),
	WelshUnitedKingdom:          LanguageID(C.TT_MS_LANGID_WELSH_UNITED_KINGDOM),
	KhmerCambodia:               LanguageID(C.TT_MS_LANGID_KHMER_CAMBODIA),
	LaoLaos:                     LanguageID(C.TT_MS_LANGID_LAO_LAOS),
	GalicianGalician:            LanguageID(C.TT_MS_LANGID_GALICIAN_GALICIAN),
	KonkaniIndia:                LanguageID(C.TT_MS_LANGID_KONKANI_INDIA),
	SyriacSyria:                 LanguageID(C.TT_MS_LANGID_SYRIAC_SYRIA),
	SinhalaSriLanka:             LanguageID(C.TT_MS_LANGID_SINHALA_SRI_LANKA),
	InuktitutCanada:             LanguageID(C.TT_MS_LANGID_INUKTITUT_CANADA),
	InuktitutCanadaLatin:        LanguageID(C.TT_MS_LANGID_INUKTITUT_CANADA_LATIN),
	AmharicEthiopia:             LanguageID(C.TT_MS_LANGID_AMHARIC_ETHIOPIA),
	TamazightAlgeria:            LanguageID(C.TT_MS_LANGID_TAMAZIGHT_ALGERIA),
	NepaliNepal:                 LanguageID(C.TT_MS_LANGID_NEPALI_NEPAL),
	FrisianNetherlands:          LanguageID(C.TT_MS_LANGID_FRISIAN_NETHERLANDS),
	PashtoAfghanistan:           LanguageID(C.TT_MS_LANGID_PASHTO_AFGHANISTAN),
	FilipinoPhilippines:         LanguageID(C.TT_MS_LANGID_FILIPINO_PHILIPPINES),
	DhivehiMaldives:             LanguageID(C.TT_MS_LANGID_DHIVEHI_MALDIVES),
	HausaNigeria:                LanguageID(C.TT_MS_LANGID_HAUSA_NIGERIA),
	YorubaNigeria:               LanguageID(C.TT_MS_LANGID_YORUBA_NIGERIA),
	QuechuaBolivia:              LanguageID(C.TT_MS_LANGID_QUECHUA_BOLIVIA),
	QuechuaEcuador:              LanguageID(C.TT_MS_LANGID_QUECHUA_ECUADOR),
	QuechuaPeru:                 LanguageID(C.TT_MS_LANGID_QUECHUA_PERU),
	SesothoSaLeboaSouthAfrica:   LanguageID(C.TT_MS_LANGID_SESOTHO_SA_LEBOA_SOUTH_AFRICA),
	BashkirRussia:               LanguageID(C.TT_MS_LANGID_BASHKIR_RUSSIA),
	LuxembourgishLuxembourg:     LanguageID(C.TT_MS_LANGID_LUXEMBOURGISH_LUXEMBOURG),
	GreenlandicGreenland:        LanguageID(C.TT_MS_LANGID_GREENLANDIC_GREENLAND),
	IgboNigeria:                 LanguageID(C.TT_MS_LANGID_IGBO_NIGERIA),
	YiPrc:                       LanguageID(C.TT_MS_LANGID_YI_PRC),
	MapudungunChile:             LanguageID(C.TT_MS_LANGID_MAPUDUNGUN_CHILE),
	MohawkMohawk:                LanguageID(C.TT_MS_LANGID_MOHAWK_MOHAWK),
	BretonFrance:                LanguageID(C.TT_MS_LANGID_BRETON_FRANCE),
	UighurPrc:                   LanguageID(C.TT_MS_LANGID_UIGHUR_PRC),
	MaoriNewZealand:             LanguageID(C.TT_MS_LANGID_MAORI_NEW_ZEALAND),
	OccitanFrance:               LanguageID(C.TT_MS_LANGID_OCCITAN_FRANCE),
	CorsicanFrance:              LanguageID(C.TT_MS_LANGID_CORSICAN_FRANCE),
	AlsatianFrance:              LanguageID(C.TT_MS_LANGID_ALSATIAN_FRANCE),
	YakutRussia:                 LanguageID(C.TT_MS_LANGID_YAKUT_RUSSIA),
	KicheGuatemala:              LanguageID(C.TT_MS_LANGID_KICHE_GUATEMALA),
	KinyarwandaRwanda:           LanguageID(C.TT_MS_LANGID_KINYARWANDA_RWANDA),
	WolofSenegal:                LanguageID(C.TT_MS_LANGID_WOLOF_SENEGAL),
	DariAfghanistan:             LanguageID(C.TT_MS_LANGID_DARI_AFGHANISTAN),
}

// NameID is the ‘name’ identifier field in the name records of an SFNT ‘name’ table.
// NameID values are platform independent.
type NameID int

// NameIDs contains all the possible values for NameID.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_name_id_xxx
var NameIDs = struct {
	Copyright            NameID
	FontFamily           NameID
	FontSubfamily        NameID
	UniqueID             NameID
	FullName             NameID
	VersionString        NameID
	PsName               NameID
	Trademark            NameID
	Manufacturer         NameID
	Designer             NameID
	Description          NameID
	VendorURL            NameID
	DesignerURL          NameID
	License              NameID
	LicenseURL           NameID
	TypographicFamily    NameID
	TypographicSubfamily NameID
	MacFullName          NameID
	SampleText           NameID
	CidFindfontName      NameID
	WwsFamily            NameID
	WwsSubfamily         NameID
	LightBackground      NameID
	DarkBackground       NameID
	VariationsPrefix     NameID
}{
	Copyright:            NameID(C.TT_NAME_ID_COPYRIGHT),
	FontFamily:           NameID(C.TT_NAME_ID_FONT_FAMILY),
	FontSubfamily:        NameID(C.TT_NAME_ID_FONT_SUBFAMILY),
	UniqueID:             NameID(C.TT_NAME_ID_UNIQUE_ID),
	FullName:             NameID(C.TT_NAME_ID_FULL_NAME),
	VersionString:        NameID(C.TT_NAME_ID_VERSION_STRING),
	PsName:               NameID(C.TT_NAME_ID_PS_NAME),
	Trademark:            NameID(C.TT_NAME_ID_TRADEMARK),
	Manufacturer:         NameID(C.TT_NAME_ID_MANUFACTURER),
	Designer:             NameID(C.TT_NAME_ID_DESIGNER),
	Description:          NameID(C.TT_NAME_ID_DESCRIPTION),
	VendorURL:            NameID(C.TT_NAME_ID_VENDOR_URL),
	DesignerURL:          NameID(C.TT_NAME_ID_DESIGNER_URL),
	License:              NameID(C.TT_NAME_ID_LICENSE),
	LicenseURL:           NameID(C.TT_NAME_ID_LICENSE_URL),
	TypographicFamily:    NameID(C.TT_NAME_ID_TYPOGRAPHIC_FAMILY),
	TypographicSubfamily: NameID(C.TT_NAME_ID_TYPOGRAPHIC_SUBFAMILY),
	MacFullName:          NameID(C.TT_NAME_ID_MAC_FULL_NAME),
	SampleText:           NameID(C.TT_NAME_ID_SAMPLE_TEXT),
	CidFindfontName:      NameID(C.TT_NAME_ID_CID_FINDFONT_NAME),
	WwsFamily:            NameID(C.TT_NAME_ID_WWS_FAMILY),
	WwsSubfamily:         NameID(C.TT_NAME_ID_WWS_SUBFAMILY),
	LightBackground:      NameID(C.TT_NAME_ID_LIGHT_BACKGROUND),
	DarkBackground:       NameID(C.TT_NAME_ID_DARK_BACKGROUND),
	VariationsPrefix:     NameID(C.TT_NAME_ID_VARIATIONS_PREFIX),
}

// CharMapLanguage reports the cmap language ID as specified in the OpenType standard.
// If the charmap at index i doesn't belong to an SFNT face, it returns 0 as the default value.
//
// For a format 14 cmap (to access Unicode IVS), the return value is 0xFFFFFFFF.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#ft_get_cmap_format
func (f *Face) CharMapLanguage(i int) LanguageID {
	if f == nil || f.ptr == nil {
		return -1
	}

	ccmap := f.getCCharMap(i)
	if ccmap == nil {
		return 0
	}

	return LanguageID(C.FT_Get_CMap_Language_ID(ccmap))
}

// CharMapFormat reports the format of an SFNT ‘cmap’ table.
// If the charmap at index i doesn't belong to an SFNT face, it returns -1.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#ft_get_cmap_format
func (f *Face) CharMapFormat(i int) int {
	if f == nil || f.ptr == nil {
		return -1
	}

	ccmap := f.getCCharMap(i)
	if ccmap == nil {
		return -1
	}

	return int(C.FT_Get_CMap_Format(ccmap))
}
