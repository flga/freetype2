// Package fixed implements fixed point types
package fixed

import (
	"fmt"

	"golang.org/x/image/math/fixed"
)

// Int26_6 is a signed 26.6 fixed-point type used for vectorial pixel coordinates.
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_f26dot6
type Int26_6 = fixed.Int26_6

// Int16_16 is a signed 16.16 fixed-point number.
//
// The integer part ranges from -32768 to 32767, inclusive. The fractional part has 16 bits of precision.
// For example, the number one-and-a-quarter is Int16_16(1<<6 + 1<<14).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_fixed
type Int16_16 int32

func (x Int16_16) String() string {
	const shift, mask = 16, 1<<16 - 1
	if x >= 0 {
		return fmt.Sprintf("%d:%d", int32(x>>shift), int32(x&mask))
	}
	x = -x
	if x >= 0 {
		return fmt.Sprintf("-%d:%d", int32(x>>shift), int32(x&mask))
	}
	return "-32768:0" // The minimum value is -(1<<15).
}

// Floor returns the greatest integer value less than or equal to x.
//
// Its return type is int, not Int16_16.
func (x Int16_16) Floor() int { return int((x) >> 16) }

// Round returns the nearest integer value to x. Ties are rounded up.
//
// Its return type is int, not Int16_16.
func (x Int16_16) Round() int { return int((x + 1<<15) >> 16) }

// Ceil returns the least integer value greater than or equal to x.
//
// Its return type is int, not Int16_16.
func (x Int16_16) Ceil() int { return int((x + 1<<16 - 1) >> 16) }

// Mul returns x*y in 16.16 fixed-point arithmetic.
func (x Int16_16) Mul(y Int16_16) Int16_16 {
	return Int16_16((int64(x)*int64(y) + 1<<15) >> 16)
}

// F32 converts the underlying value to float32.
func (x Int16_16) F32() float32 {
	return float32(x) / float32(1<<16)
}

// F64 converts the underlying value to float64.
func (x Int16_16) F64() float64 {
	return float64(x) / float64(1<<16)
}

// Int2_14 is a signed 2.14 fixed-point number.
//
// The integer part ranges from -2 to 1, inclusive. The fractional part has 14 bits of precision.
// For example, the number one-and-a-quarter is Int2_14(1<<2 + 1<<12).
//
// See https://www.freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_f2dot14
type Int2_14 int16

func (x Int2_14) String() string {
	const shift, mask = 14, 1<<14 - 1
	if x >= 0 {
		return fmt.Sprintf("%d:%d", int16(x>>shift), int16(x&mask))
	}
	x = -x
	if x >= 0 {
		return fmt.Sprintf("-%d:%d", int16(x>>shift), int16(x&mask))
	}
	return "-2:0" // The minimum value is -(1<<1).
}

// Floor returns the greatest integer value less than or equal to x.
//
// Its return type is int, not Int2_14.
func (x Int2_14) Floor() int { return int((x) >> 14) }

// Round returns the nearest integer value to x. Ties are rounded up.
//
// Its return type is int, not Int2_14.
func (x Int2_14) Round() int { return int((x + 1<<13) >> 14) }

// Ceil returns the least integer value greater than or equal to x.
//
// Its return type is int, not Int2_14.
func (x Int2_14) Ceil() int { return int((x + 1<<14 - 1) >> 14) }

// Mul returns x*y in 2.14 fixed-point arithmetic.
func (x Int2_14) Mul(y Int2_14) Int2_14 {
	return Int2_14((int32(x)*int32(y) + 1<<13) >> 14)
}

// F32 converts the underlying value to float32.
func (x Int2_14) F32() float32 {
	return float32(x) / float32(1<<14)
}

// F64 converts the underlying value to float64.
func (x Int2_14) F64() float64 {
	return float64(x) / float64(1<<14)
}
