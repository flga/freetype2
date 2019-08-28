package freetype2

import (
	"reflect"
	"testing"

	"github.com/flga/freetype2/fixed"
)

func TestVector_Transform(t *testing.T) {
	tests := []struct {
		name string
		v    Vector16_16
		m    Matrix
		want Vector16_16
	}{
		{
			name: "fixed",
			v: Vector16_16{
				X: 2 << 16,
				Y: 1 << 16,
			},
			m: Matrix{
				Xx: 2 << 16, Xy: 1 << 16,
				Yx: -1 << 16, Yy: 2 << 16,
			},
			want: Vector16_16{
				X: 5 << 16,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Transform(tt.m); got != tt.want {
				t.Errorf("Vector.Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Mul(t *testing.T) {
	tests := []struct {
		name string
		a    Matrix
		b    Matrix
		want Matrix
	}{
		{
			name: "fixed",
			a: Matrix{
				Xx: 2 << 16, Xy: -2 << 16,
				Yx: 5 << 16, Yy: 3 << 16,
			},
			b: Matrix{
				Xx: -1 << 16, Xy: 4 << 16,
				Yx: 7 << 16, Yy: -6 << 16,
			},
			want: Matrix{
				Xx: -16 << 16, Xy: 20 << 16,
				Yx: 16 << 16, Yy: 2 << 16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Mul(tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Invert(t *testing.T) {
	asFixed := func(f float64) fixed.Int16_16 {
		return fixed.Int16_16(f * float64(0x10000))
	}
	tests := []struct {
		name    string
		a       Matrix
		want    Matrix
		wantErr error
	}{
		{
			name: "fixed",
			a: Matrix{
				Xx: 5 << 16, Xy: 4 << 16,
				Yx: 2 << 16, Yy: 2 << 16,
			},
			want: Matrix{
				Xx: asFixed(1), Xy: asFixed(-2),
				Yx: asFixed(-1), Yy: asFixed(2.5),
			},
		},
		{
			name: "not invertible",
			a: Matrix{
				Xx: 3 << 16, Xy: 4 << 16,
				Yx: 6 << 16, Yy: 8 << 16,
			},
			want:    Matrix{},
			wantErr: ErrInvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Invert()
			if err != tt.wantErr {
				t.Errorf("Matrix.Invert() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Matrix.Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Sin(t *testing.T) {
	tests := []struct {
		name string
		a    Angle
		want fixed.Int16_16
	}{
		{name: "PI", a: AnglePI, want: 0x0},
		{name: "2PI", a: Angle2PI, want: 0x0},
		{name: "PI2", a: AnglePI2, want: 0x10000},
		{name: "PI4", a: AnglePI4, want: 0xb505},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Sin(); got != tt.want {
				t.Errorf("Angle.Sin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Cos(t *testing.T) {
	tests := []struct {
		name string
		a    Angle
		want fixed.Int16_16
	}{
		{name: "PI", a: AnglePI, want: -0x10000},
		{name: "2PI", a: Angle2PI, want: 0x10000},
		{name: "PI2", a: AnglePI2, want: 0x0},
		{name: "PI4", a: AnglePI4, want: 0xb505},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Cos(); got != tt.want {
				t.Errorf("Angle.Cos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Tan(t *testing.T) {
	tests := []struct {
		name string
		a    Angle
		want fixed.Int16_16
	}{
		{name: "PI", a: AnglePI, want: 0x0},
		{name: "2PI", a: Angle2PI, want: 0x0},
		// {name: "PI2", a: AnglePI2, want: -0xc378000}, fails on 32 bit
		{name: "PI4", a: AnglePI4, want: 0x10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Tan(); got != tt.want {
				t.Errorf("Angle.Tan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtan2(t *testing.T) {
	tests := []struct {
		name string
		x    fixed.Int16_16
		y    fixed.Int16_16
		want Angle
	}{
		{x: 2, y: 1, want: 0x1a90b0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Atan2(tt.x, tt.y); got != tt.want {
				t.Errorf("Atan2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Diff(t *testing.T) {
	tests := []struct {
		name string
		a    Angle
		b    Angle
		want Angle
	}{
		{a: AnglePI, b: Angle2PI, want: AnglePI},
		{a: AnglePI2, b: AnglePI4, want: -AnglePI4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Diff(tt.b); got != tt.want {
				t.Errorf("Angle.Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_UnitVector(t *testing.T) {
	tests := []struct {
		name string
		a    Angle
		want Vector
	}{
		{name: "PI", a: AnglePI, want: Vector{X: -0x10000, Y: 0x0}},
		{name: "2PI", a: Angle2PI, want: Vector{X: 0x10000, Y: 0x0}},
		{name: "PI2", a: AnglePI2, want: Vector{X: 0x0, Y: 0x10000}},
		{name: "PI4", a: AnglePI4, want: Vector{X: 0xb505, Y: 0xb505}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.UnitVector(); got != tt.want {
				t.Errorf("Angle.UnitVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Rotate(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		a    Angle
		want Vector
	}{
		{name: "PI", v: Vector{X: -0x10000, Y: 0x0}, a: AnglePI, want: Vector{X: 0x10000, Y: 0x0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Rotate(tt.a); got != tt.want {
				t.Errorf("Vector.Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Length(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want Pos
	}{
		{v: Vector{X: 1, Y: 1}, want: 1},
		{v: Vector{X: 1 << 16, Y: 1 << 16}, want: 0x16a0a},
		{v: Vector{X: 6 << 16, Y: 3 << 16}, want: 0x6b54d},
		{v: Vector{X: 6 << 6, Y: 3 << 6}, want: 0x1ad},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Length(); got != tt.want {
				t.Errorf("Vector.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Polarize(t *testing.T) {
	tests := []struct {
		name       string
		v          Vector
		wantLength Pos
		wantAngle  Angle
	}{
		{v: Vector{X: 1, Y: 1}, wantLength: 0x1, wantAngle: 0x2d0000},
		{v: Vector{X: 1 << 16, Y: 1 << 16}, wantLength: 0x16a09, wantAngle: 0x2d0000},
		{v: Vector{X: 6 << 16, Y: 3 << 16}, wantLength: 0x6b54c, wantAngle: 0x1a90b0},
		{v: Vector{X: 6 << 6, Y: 3 << 6}, wantLength: 0x1ad, wantAngle: 0x1a90b0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLength, gotAngle := tt.v.Polarize()
			if gotLength != tt.wantLength {
				t.Errorf("Vector.Polarize() gotLength = %v, want %v", gotLength, tt.wantLength)
			}
			if gotAngle != tt.wantAngle {
				t.Errorf("Vector.Polarize() gotAngle = %v, want %v", gotAngle, tt.wantAngle)
			}
		})
	}
}

func TestVectorFromPolar(t *testing.T) {
	tests := []struct {
		name   string
		length fixed.Int16_16
		angle  Angle
		want   Vector
	}{
		{length: 0x00001, angle: 0x2d0000, want: Vector{X: 1, Y: 1}},
		{length: 0x16a09, angle: 0x2d0000, want: Vector{X: 1<<16 - 1, Y: 1<<16 - 1}},
		{length: 0x6b54c, angle: 0x1a90b0, want: Vector{X: 6<<16 - 1, Y: 3<<16 + 1}},
		{length: 0x001ad, angle: 0x1a90b0, want: Vector{X: 6 << 6, Y: 3 << 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VectorFromPolar(tt.length, tt.angle); got != tt.want {
				t.Errorf("VectorFromPolar() = %v, want %v", got, tt.want)
			}
		})
	}
}
