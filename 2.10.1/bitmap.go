package freetype2

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include FT_BITMAP_H
import (
	"C"
)
import (
	"image/color"

	"github.com/flga/freetype2/fixed"
)

// NewBitmap creates a new empty bitmap.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bitmap_handling.html#ft_bitmap_init
func (l *Library) NewBitmap() (*Bitmap, error) {
	if l == nil || l.ptr == nil {
		return nil, ErrInvalidLibraryHandle
	}

	ptr := (*C.FT_Bitmap)(C.calloc(1, C.sizeof_FT_Bitmap))
	C.FT_Bitmap_Init(ptr)

	ret := &Bitmap{
		ptr:         ptr,
		l:           l,
		userCreated: true,
	}
	l.dealloc = append(l.dealloc, func() { ret.Free() })

	return ret, nil
}

// Free destroys a bitmap created with NewBitmap.
//
// The library argument is taken to have access to FreeType's memory handling functions.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bitmap_handling.html#ft_bitmap_done
func (b *Bitmap) Free() error {
	if b == nil || b.ptr == nil {
		return nil
	}

	if !b.userCreated {
		return nil
	}

	if err := getErr(C.FT_Bitmap_Done(b.l.ptr, b.ptr)); err != nil {
		return err
	}

	*b = Bitmap{}
	return nil
}

// CopyTo copies a bitmap into another one.
//
// b.buffer and target.buffer must neither be equal nor overlap.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bitmap_handling.html#ft_bitmap_copy
func (b *Bitmap) CopyTo(l *Library, target *Bitmap) error {
	if b == nil || b.ptr == nil {
		return ErrInvalidArgument
	}
	if l == nil || l.ptr == nil {
		return ErrInvalidLibraryHandle
	}
	if target == nil || target.ptr == nil {
		return ErrInvalidArgument
	}

	if err := getErr(C.FT_Bitmap_Copy(l.ptr, b.ptr, target.ptr)); err != nil {
		return err
	}

	target.reload()
	return nil
}

// Embolden emboldens a bitmap.
//
// The new bitmap will be about xStrength pixels wider and yStrength pixels higher.
// The left and bottom borders are kept unchanged.
//
// The current implementation restricts xStrength to be less than or equal to 8
// if bitmap is of pixel_mode PixelModeMono.
//
// If you want to embolden the bitmap owned by a GlyphSlot, you should call
// GlyphSlot.OwnBitmap() on the slot first.
//
// Bitmaps in PixelModeGray2 and PixelModeGray4 format are converted to
// PixelModeGray format (i.e., 8bpp).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bitmap_handling.html#ft_bitmap_embolden
func (b *Bitmap) Embolden(l *Library, xStrength, yStrength fixed.Int26_6) error {
	if b == nil || b.ptr == nil {
		return ErrInvalidArgument
	}

	if l == nil || l.ptr == nil {
		return ErrInvalidLibraryHandle
	}

	if err := getErr(C.FT_Bitmap_Embolden(l.ptr, b.ptr, C.FT_Pos(xStrength), C.FT_Pos(yStrength))); err != nil {
		return err
	}

	b.reload()
	return nil
}

// Convert converts b into a new bitmap object with depth 1bpp, 2bpp, 4bpp, 8bpp
// or 32bpp to a bitmap object with depth 8bpp, making the number of used bytes
// per line (a.k.a. the ‘pitch’) a multiple of alignment.
//
// Use Free remove the new bitmap object.
//
// The library argument is taken to have access to FreeType's memory handling
// functions.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bitmap_handling.html#ft_bitmap_convert
func (b *Bitmap) Convert(l *Library, alignment int) (*Bitmap, error) {
	if b == nil || b.ptr == nil {
		return nil, ErrInvalidArgument
	}

	if l == nil || l.ptr == nil {
		return nil, ErrInvalidLibraryHandle
	}

	target, err := l.NewBitmap()
	if err != nil {
		return nil, err
	}
	if err := getErr(C.FT_Bitmap_Convert(l.ptr, b.ptr, target.ptr, C.FT_Int(alignment))); err != nil {
		return nil, err
	}

	target.reload()
	return target, nil
}

// Blend blends b onto another bitmap, using a given color.
//
// It returns the offset vector to the upper left corner of the target bitmap.
// It should represent an integer offset; the function will set the lowest six
// bits to zero to enforce that.
//
// b can have any PixelMode.
//
// srcOffset is the offset vector to the upper left corner of the source bitmap.
// It should represent an integer offset; the function will set the lowest six
// bits to zero to enforce that.
//
// target should be either initialized as empty with a call to NewBitmap, or it
// should be of type PixelModeBGRA.
//
// This function doesn't perform clipping.
//
// The bitmap in target gets allocated or reallocated as needed; the targetOffset
// vector is updated accordingly.
//
// In case of allocation or reallocation, the bitmap's pitch is set to 4 * width.
// Both source and target must have the same bitmap flow (as indicated by the
// sign of the pitch field).
//
// b.buffer and target.buffer must neither be equal nor overlap.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bitmap_handling.html#ft_bitmap_blend
func (b *Bitmap) Blend(l *Library, srcOffset Vector26_6, target *Bitmap, color color.RGBA) (targetOffset Vector26_6, err error) {
	if b == nil || b.ptr == nil {
		return Vector26_6{}, ErrInvalidArgument
	}
	if l == nil || l.ptr == nil {
		return Vector26_6{}, ErrInvalidLibraryHandle
	}
	if target == nil || target.ptr == nil {
		return Vector26_6{}, ErrInvalidArgument
	}

	var ctargetOffset C.FT_Vector

	err = getErr(C.FT_Bitmap_Blend(
		l.ptr,
		b.ptr,
		C.FT_Vector{
			x: C.FT_Pos(srcOffset.X),
			y: C.FT_Pos(srcOffset.Y),
		},
		target.ptr,
		&ctargetOffset,
		C.FT_Color{
			blue:  C.FT_Byte(color.B),
			green: C.FT_Byte(color.G),
			red:   C.FT_Byte(color.R),
			alpha: C.FT_Byte(color.A),
		},
	))
	if err != nil {
		return Vector26_6{}, err
	}

	target.reload()
	return Vector26_6{
		X: fixed.Int26_6(ctargetOffset.x),
		Y: fixed.Int26_6(ctargetOffset.y),
	}, nil
}

// OwnBitmap makes sure that a glyph slot owns its bitmap.
//
// This function is to be used in combination with Bitmap.Embolden.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-bitmap_handling.html#ft_glyphslot_own_bitmap
func (s *GlyphSlot) OwnBitmap() error {
	if s == nil || s.ptr == nil {
		return nil
	}

	return getErr(C.FT_GlyphSlot_Own_Bitmap(s.ptr))
}
