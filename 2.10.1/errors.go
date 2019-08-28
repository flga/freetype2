package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// const char* getFreetypeError(FT_Error err)
// {
//     #undef __FTERRORS_H__
//     #define FT_ERRORDEF( e, v, s )  case e: return s;
//     #define FT_ERROR_START_LIST     switch (err) {
//     #define FT_ERROR_END_LIST       }
//     #include FT_ERRORS_H
//     return "Unknown error";
// }
import "C"
import (
	"errors"
	"sync"
)

var magicStringUnknownError = "Unknown error" // must be equal to the default case of C.getFreetypeError

func testErrCannotOpenResource() error {
	return getErr(C.FT_Err_Cannot_Open_Resource)
}

func testUnmappedErr() error {
	return getErr(C.int(981298371)) //error code that *hopefully* does not exist
}

// Generic errors
var (
	ErrUnknownError         = errors.New(magicStringUnknownError)
	ErrCannotOpenResource   = errors.New("cannot open resource")
	ErrUnknownFileFormat    = errors.New("unknown file format")
	ErrInvalidFileFormat    = errors.New("broken file")
	ErrInvalidVersion       = errors.New("invalid FreeType version")
	ErrLowerModuleVersion   = errors.New("module version is too low")
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrUnimplementedFeature = errors.New("unimplemented feature")
	ErrInvalidTable         = errors.New("broken table")
	ErrInvalidOffset        = errors.New("broken offset within table")
	ErrArrayTooLarge        = errors.New("array allocation size too large")
	ErrMissingModule        = errors.New("missing module")
	ErrMissingProperty      = errors.New("missing property")
)

// Glyph/character errors
var (
	ErrInvalidGlyphIndex    = errors.New("invalid glyph index")
	ErrInvalidCharacterCode = errors.New("invalid character code")
	ErrInvalidGlyphFormat   = errors.New("unsupported glyph image format")
	ErrCannotRenderGlyph    = errors.New("cannot render this glyph format")
	ErrInvalidOutline       = errors.New("invalid outline")
	ErrInvalidComposite     = errors.New("invalid composite glyph")
	ErrTooManyHints         = errors.New("too many hints")
	ErrInvalidPixelSize     = errors.New("invalid pixel size")
)

// Handle errors
var (
	ErrInvalidHandle        = errors.New("invalid object handle")
	ErrInvalidLibraryHandle = errors.New("invalid library handle")
	ErrInvalidDriverHandle  = errors.New("invalid module handle")
	ErrInvalidFaceHandle    = errors.New("invalid face handle")
	ErrInvalidSizeHandle    = errors.New("invalid size handle")
	ErrInvalidSlotHandle    = errors.New("invalid glyph slot handle")
	ErrInvalidCharMapHandle = errors.New("invalid charmap handle")
	ErrInvalidCacheHandle   = errors.New("invalid cache manager handle")
	ErrInvalidStreamHandle  = errors.New("invalid stream handle")
)

// Driver errors
var (
	ErrTooManyDrivers    = errors.New("too many modules")
	ErrTooManyExtensions = errors.New("too many extensions")
)

// Memory errors
var (
	ErrOutOfMemory    = errors.New("out of memory")
	ErrUnlistedObject = errors.New("unlisted object")
)

// Stream errors
var (
	ErrCannotOpenStream       = errors.New("cannot open stream")
	ErrInvalidStreamSeek      = errors.New("invalid stream seek")
	ErrInvalidStreamSkip      = errors.New("invalid stream skip")
	ErrInvalidStreamRead      = errors.New("invalid stream read")
	ErrInvalidStreamOperation = errors.New("invalid stream operation")
	ErrInvalidFrameOperation  = errors.New("invalid frame operation")
	ErrNestedFrameAccess      = errors.New("nested frame access")
	ErrInvalidFrameRead       = errors.New("invalid frame read")
)

// Raster errors
var (
	ErrRasterUninitialized  = errors.New("raster uninitialized")
	ErrRasterCorrupted      = errors.New("raster corrupted")
	ErrRasterOverflow       = errors.New("raster overflow")
	ErrRasterNegativeHeight = errors.New("negative height while rastering")
)

// Cache errors
var (
	ErrTooManyCaches = errors.New("too many registered caches")
)

// TrueType and SFNT errors
var (
	ErrInvalidOpcode          = errors.New("invalid opcode")
	ErrTooFewArguments        = errors.New("too few arguments")
	ErrStackOverflow          = errors.New("stack overflow")
	ErrCodeOverflow           = errors.New("code overflow")
	ErrBadArgument            = errors.New("bad argument")
	ErrDivideByZero           = errors.New("division by zero")
	ErrInvalidReference       = errors.New("invalid reference")
	ErrDebugOpCode            = errors.New("found debug opcode")
	ErrENDFInExecStream       = errors.New("found ENDF opcode in execution stream")
	ErrNestedDEFS             = errors.New("nested DEFS")
	ErrInvalidCodeRange       = errors.New("invalid code range")
	ErrExecutionTooLong       = errors.New("execution context too long")
	ErrTooManyFunctionDefs    = errors.New("too many function definitions")
	ErrTooManyInstructionDefs = errors.New("too many instruction definitions")
	ErrTableMissing           = errors.New("SFNT font table missing")
	ErrHorizHeaderMissing     = errors.New("horizontal header (hhea) table missing")
	ErrLocationsMissing       = errors.New("locations (loca) table missing")
	ErrNameTableMissing       = errors.New("name table missing")
	ErrCMapTableMissing       = errors.New("character map (cmap) table missing")
	ErrHmtxTableMissing       = errors.New("horizontal metrics (hmtx) table missing")
	ErrPostTableMissing       = errors.New("PostScript (post) table missing")
	ErrInvalidHorizMetrics    = errors.New("invalid horizontal metrics")
	ErrInvalidCharMapFormat   = errors.New("invalid character map (cmap) format")
	ErrInvalidPPem            = errors.New("invalid ppem value")
	ErrInvalidVertMetrics     = errors.New("invalid vertical metrics")
	ErrCouldNotFindContext    = errors.New("could not find context")
	ErrInvalidPostTableFormat = errors.New("invalid PostScript (post) table format")
	ErrInvalidPostTable       = errors.New("invalid PostScript (post) table")
	ErrDEFInGlyfBytecode      = errors.New("found FDEF or IDEF opcode in glyf bytecode")
	ErrMissingBitmap          = errors.New("missing bitmap in strike")
)

// CFF, CID, and Type 1 errors
var (
	ErrSyntaxError        = errors.New("opcode syntax error")
	ErrStackUnderflow     = errors.New("argument stack underflow")
	ErrIgnore             = errors.New("ignore")
	ErrNoUnicodeGlyphName = errors.New("no Unicode glyph name found")
	ErrGlyphTooBig        = errors.New("glyph too big for hinting")
)

// BDF errors
var (
	ErrMissingStartfontField       = errors.New("`STARTFONT' field missing")
	ErrMissingFontField            = errors.New("`FONT' field missing")
	ErrMissingSizeField            = errors.New("`SIZE' field missing")
	ErrMissingFontboundingboxField = errors.New("`FONTBOUNDINGBOX' field missing")
	ErrMissingCharsField           = errors.New("`CHARS' field missing")
	ErrMissingStartcharField       = errors.New("`STARTCHAR' field missing")
	ErrMissingEncodingField        = errors.New("`ENCODING' field missing")
	ErrMissingBbxField             = errors.New("`BBX' field missing")
	ErrBbxTooBig                   = errors.New("`BBX' too big")
	ErrCorruptedFontHeader         = errors.New("font header corrupted or missing fields")
	ErrCorruptedFontGlyphs         = errors.New("font glyphs corrupted or missing fields")
)

var errMap = map[C.int]error{
	C.FT_Err_Cannot_Open_Resource:          ErrCannotOpenResource,
	C.FT_Err_Unknown_File_Format:           ErrUnknownFileFormat,
	C.FT_Err_Invalid_File_Format:           ErrInvalidFileFormat,
	C.FT_Err_Invalid_Version:               ErrInvalidVersion,
	C.FT_Err_Lower_Module_Version:          ErrLowerModuleVersion,
	C.FT_Err_Invalid_Argument:              ErrInvalidArgument,
	C.FT_Err_Unimplemented_Feature:         ErrUnimplementedFeature,
	C.FT_Err_Invalid_Table:                 ErrInvalidTable,
	C.FT_Err_Invalid_Offset:                ErrInvalidOffset,
	C.FT_Err_Array_Too_Large:               ErrArrayTooLarge,
	C.FT_Err_Missing_Module:                ErrMissingModule,
	C.FT_Err_Missing_Property:              ErrMissingProperty,
	C.FT_Err_Invalid_Glyph_Index:           ErrInvalidGlyphIndex,
	C.FT_Err_Invalid_Character_Code:        ErrInvalidCharacterCode,
	C.FT_Err_Invalid_Glyph_Format:          ErrInvalidGlyphFormat,
	C.FT_Err_Cannot_Render_Glyph:           ErrCannotRenderGlyph,
	C.FT_Err_Invalid_Outline:               ErrInvalidOutline,
	C.FT_Err_Invalid_Composite:             ErrInvalidComposite,
	C.FT_Err_Too_Many_Hints:                ErrTooManyHints,
	C.FT_Err_Invalid_Pixel_Size:            ErrInvalidPixelSize,
	C.FT_Err_Invalid_Handle:                ErrInvalidHandle,
	C.FT_Err_Invalid_Library_Handle:        ErrInvalidLibraryHandle,
	C.FT_Err_Invalid_Driver_Handle:         ErrInvalidDriverHandle,
	C.FT_Err_Invalid_Face_Handle:           ErrInvalidFaceHandle,
	C.FT_Err_Invalid_Size_Handle:           ErrInvalidSizeHandle,
	C.FT_Err_Invalid_Slot_Handle:           ErrInvalidSlotHandle,
	C.FT_Err_Invalid_CharMap_Handle:        ErrInvalidCharMapHandle,
	C.FT_Err_Invalid_Cache_Handle:          ErrInvalidCacheHandle,
	C.FT_Err_Invalid_Stream_Handle:         ErrInvalidStreamHandle,
	C.FT_Err_Too_Many_Drivers:              ErrTooManyDrivers,
	C.FT_Err_Too_Many_Extensions:           ErrTooManyExtensions,
	C.FT_Err_Out_Of_Memory:                 ErrOutOfMemory,
	C.FT_Err_Unlisted_Object:               ErrUnlistedObject,
	C.FT_Err_Cannot_Open_Stream:            ErrCannotOpenStream,
	C.FT_Err_Invalid_Stream_Seek:           ErrInvalidStreamSeek,
	C.FT_Err_Invalid_Stream_Skip:           ErrInvalidStreamSkip,
	C.FT_Err_Invalid_Stream_Read:           ErrInvalidStreamRead,
	C.FT_Err_Invalid_Stream_Operation:      ErrInvalidStreamOperation,
	C.FT_Err_Invalid_Frame_Operation:       ErrInvalidFrameOperation,
	C.FT_Err_Nested_Frame_Access:           ErrNestedFrameAccess,
	C.FT_Err_Invalid_Frame_Read:            ErrInvalidFrameRead,
	C.FT_Err_Raster_Uninitialized:          ErrRasterUninitialized,
	C.FT_Err_Raster_Corrupted:              ErrRasterCorrupted,
	C.FT_Err_Raster_Overflow:               ErrRasterOverflow,
	C.FT_Err_Raster_Negative_Height:        ErrRasterNegativeHeight,
	C.FT_Err_Too_Many_Caches:               ErrTooManyCaches,
	C.FT_Err_Invalid_Opcode:                ErrInvalidOpcode,
	C.FT_Err_Too_Few_Arguments:             ErrTooFewArguments,
	C.FT_Err_Stack_Overflow:                ErrStackOverflow,
	C.FT_Err_Code_Overflow:                 ErrCodeOverflow,
	C.FT_Err_Bad_Argument:                  ErrBadArgument,
	C.FT_Err_Divide_By_Zero:                ErrDivideByZero,
	C.FT_Err_Invalid_Reference:             ErrInvalidReference,
	C.FT_Err_Debug_OpCode:                  ErrDebugOpCode,
	C.FT_Err_ENDF_In_Exec_Stream:           ErrENDFInExecStream,
	C.FT_Err_Nested_DEFS:                   ErrNestedDEFS,
	C.FT_Err_Invalid_CodeRange:             ErrInvalidCodeRange,
	C.FT_Err_Execution_Too_Long:            ErrExecutionTooLong,
	C.FT_Err_Too_Many_Function_Defs:        ErrTooManyFunctionDefs,
	C.FT_Err_Too_Many_Instruction_Defs:     ErrTooManyInstructionDefs,
	C.FT_Err_Table_Missing:                 ErrTableMissing,
	C.FT_Err_Horiz_Header_Missing:          ErrHorizHeaderMissing,
	C.FT_Err_Locations_Missing:             ErrLocationsMissing,
	C.FT_Err_Name_Table_Missing:            ErrNameTableMissing,
	C.FT_Err_CMap_Table_Missing:            ErrCMapTableMissing,
	C.FT_Err_Hmtx_Table_Missing:            ErrHmtxTableMissing,
	C.FT_Err_Post_Table_Missing:            ErrPostTableMissing,
	C.FT_Err_Invalid_Horiz_Metrics:         ErrInvalidHorizMetrics,
	C.FT_Err_Invalid_CharMap_Format:        ErrInvalidCharMapFormat,
	C.FT_Err_Invalid_PPem:                  ErrInvalidPPem,
	C.FT_Err_Invalid_Vert_Metrics:          ErrInvalidVertMetrics,
	C.FT_Err_Could_Not_Find_Context:        ErrCouldNotFindContext,
	C.FT_Err_Invalid_Post_Table_Format:     ErrInvalidPostTableFormat,
	C.FT_Err_Invalid_Post_Table:            ErrInvalidPostTable,
	C.FT_Err_DEF_In_Glyf_Bytecode:          ErrDEFInGlyfBytecode,
	C.FT_Err_Missing_Bitmap:                ErrMissingBitmap,
	C.FT_Err_Syntax_Error:                  ErrSyntaxError,
	C.FT_Err_Stack_Underflow:               ErrStackUnderflow,
	C.FT_Err_Ignore:                        ErrIgnore,
	C.FT_Err_No_Unicode_Glyph_Name:         ErrNoUnicodeGlyphName,
	C.FT_Err_Glyph_Too_Big:                 ErrGlyphTooBig,
	C.FT_Err_Missing_Startfont_Field:       ErrMissingStartfontField,
	C.FT_Err_Missing_Font_Field:            ErrMissingFontField,
	C.FT_Err_Missing_Size_Field:            ErrMissingSizeField,
	C.FT_Err_Missing_Fontboundingbox_Field: ErrMissingFontboundingboxField,
	C.FT_Err_Missing_Chars_Field:           ErrMissingCharsField,
	C.FT_Err_Missing_Startchar_Field:       ErrMissingStartcharField,
	C.FT_Err_Missing_Encoding_Field:        ErrMissingEncodingField,
	C.FT_Err_Missing_Bbx_Field:             ErrMissingBbxField,
	C.FT_Err_Bbx_Too_Big:                   ErrBbxTooBig,
	C.FT_Err_Corrupted_Font_Header:         ErrCorruptedFontHeader,
	C.FT_Err_Corrupted_Font_Glyphs:         ErrCorruptedFontGlyphs,
}

// getErr is a var so that we can override it in tests, if needed
var getErr = func(code C.int) error {
	if code == 0 {
		return nil
	}

	if err, ok := errMap[code]; ok {
		return err
	}

	str := C.GoString(C.getFreetypeError(code))
	if str == magicStringUnknownError {
		return ErrUnknownError
	}

	return errors.New(str)
}

var mockGetErrMu sync.Mutex

func mockGetErr(fn func(c int) error) (restore func()) {
	mockGetErrMu.Lock()
	orig := getErr
	getErr = func(c C.int) error {
		return fn(int(c))
	}
	return func() {
		getErr = orig
		mockGetErrMu.Unlock()
	}
}
