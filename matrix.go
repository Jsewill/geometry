/* */
package geometry

import (
	"math"
)

type Matrix interface {
	Assessor
}

func Homogenous1x2() Mat1x2 {
	return Mat1x2{0, 1}
}

func Homogenous1x3() Mat1x3 {
	return Mat1x3{0, 0, 1}
}

func Homogenous1x4() Mat1x4 {
	return Mat1x4{0, 0, 0, 1}
}

func Homogenous2x1() Mat2x1 {
	return Mat2x1{
		Mat1x1{0},
		Mat1x1{1},
	}
}

func Homogenous3x1() Mat3x1 {
	return Mat3x1{
		Mat1x1{0},
		Mat1x1{0},
		Mat1x1{1},
	}
}

func Homogenous4x1() Mat4x1 {
	return Mat4x1{
		Mat1x1{0},
		Mat1x1{0},
		Mat1x1{0},
		Mat1x1{1},
	}
}

func Identity1() Mat1x1 {
	return Mat1x1{1}
}

func Identity2() Mat2x2 {
	return Mat2x2{
		Mat1x2{1, 0},
		Mat1x2{0, 1},
	}
}

func Identity3() Mat3x3 {
	return Mat3x3{
		Mat1x3{1, 0, 0},
		Mat1x3{0, 1, 0},
		Mat1x3{0, 0, 1},
	}
}

func Identity4() Mat4x4 {
	return Mat4x4{
		Mat1x4{1, 0, 0, 0},
		Mat1x4{0, 1, 0, 0},
		Mat1x4{0, 0, 1, 0},
		Mat1x4{0, 0, 0, 1},
	}
}

func LinearReflect1() Mat1x1 {
	return Mat1x1{-1}
}

func LinearReflect2(Γ Vec2) Mat2x2 {
	return Mat2x2{
		Mat1x2{Γ[0], 0},
		Mat1x2{0, Γ[1]},
	}
}

func LinearReflect3(Γ Vec3) Mat3x3 {
	return Mat3x3{
		Mat1x3{Γ[0], 0, 0},
		Mat1x3{0, Γ[1], 0},
		Mat1x3{0, 0, Γ[2]},
	}
}

//Takes in a Vec3 comprised of 1s and/or -1s. If the value in the vector are anything else, scaling may take place
func LinearReflect4(Γ Vec4) Mat4x4 {
	return Mat4x4{
		Mat1x4{Γ[0], 0, 0, 0},
		Mat1x4{0, Γ[1], 0, 0},
		Mat1x4{0, 0, Γ[2], 0},
		Mat1x4{0, 0, 0, Γ[3]},
	}
}

//Takes one angle representing the subsequent rotation in the imaginary XY plane.
func LinearRotate2(Θ float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{math.Cos(Θ), -math.Sin(Θ)},
		Mat1x2{math.Sin(Θ), math.Cos(Θ)},
	}
}

//Takes up to three angles representing the subsequent rotation in each plane. Order is as follows: YZ, XZ, XY
func LinearRotate3(Θ Vec3) Mat3x3 {
	var planes int
	for i, θ := range Θ {
		if θ != 0 {
			planes |= 1 << uint(i)
		}
	}

	if planes == 0 {
		return Identity3()
	}

	var c, s float64
	var a, b Mat3x3

	switch {
	//X-axis rotation
	//Rotate in the YZ plane.
	case planes&1 > 0:
		c, s = math.Cos(Θ[0]), math.Sin(Θ[0])
		a = Mat3x3{
			Mat1x3{1, 0, 0},
			Mat1x3{0, c, -s},
			Mat1x3{0, s, c},
		}

		fallthrough
	//Y-axis rotation
	//Rotate in the XZ plane.
	case planes&2 > 0:
		c, s = math.Cos(Θ[1]), math.Sin(Θ[1])

		b = Mat3x3{
			Mat1x3{c, 0, s},
			Mat1x3{0, 1, 0},
			Mat1x3{-s, 0, c},
		}

		if planes&1 == 0 {
			a = b
		} else {
			a = a.MultiplyMat3x3(b)
		}

		fallthrough
	//Z-axis rotation
	//Rotate in the XY plane.
	case planes&4 > 0:
		c, s = math.Cos(Θ[2]), math.Sin(Θ[2])
		b = Mat3x3{
			Mat1x3{c, -s, 0},
			Mat1x3{s, c, 0},
			Mat1x3{0, 0, 1},
		}

		if planes&3 == 0 {
			a = b
		} else {
			a = a.MultiplyMat3x3(b)
		}
	}

	return a
}

//Z-axis rotation
//Rotate in the XY plane.
func LinearRotate3XY(Θ float64) Mat3x3 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat3x3{
		Mat1x3{c, -s, 0},
		Mat1x3{s, c, 0},
		Mat1x3{0, 0, 1},
	}
}

//Y-axis rotation
//Rotate in the XZ plane.
func LinearRotate3XZ(Θ float64) Mat3x3 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat3x3{
		Mat1x3{c, 0, s},
		Mat1x3{0, 1, 0},
		Mat1x3{-s, 0, c},
	}
}

//X-axis rotation
//Rotate in the YZ plane.
func LinearRotate3YZ(Θ float64) Mat3x3 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat3x3{
		Mat1x3{1, 0, 0},
		Mat1x3{0, c, -s},
		Mat1x3{0, s, c},
	}
}

/*
4D Rotation. Rotation happens in a plane. Just as 2D rotations happen in a plane perpendicular to the imaginary Z-axis (XY),
and 3D rotations happen in 3 planes (XY, XZ, YZ), so 4D has six rotations planes (XY, XZ, XW, YZ, YW, ZW). Interestingly,
the number of planes in an arbitrary dimension corresponds to a triangular number (0, 1, 3, 6, 10, 15, 21, 28, 36, etc).

Two Dimensions:
	XY

Three Dimensions:
	XY XZ
	YZ

Four Dimensions:
	XY XZ XW
	YZ YW
	ZW

See a pattern?
*/

//Takes up to six angles representing the subsequent rotation in each plane. Order is as follows: ZW, YW, YZ, XW, XZ, XY
func LinearRotate4(Θ [6]float64) Mat4x4 {
	var planes int
	for i, θ := range Θ {
		if θ != 0 {
			planes |= 1 << uint(i)
		}
	}

	if planes == 0 {
		return Identity4()
	}

	var c, s float64
	var a, b Mat4x4

	switch {
	//Rotate in the ZW plane.
	case planes&1 > 0:
		c, s = math.Cos(Θ[0]), math.Sin(Θ[0])
		a = Mat4x4{
			Mat1x4{1, 0, 0, 0},
			Mat1x4{0, 1, 0, 0},
			Mat1x4{0, 0, c, -s},
			Mat1x4{0, 0, s, c},
		}

		fallthrough
	//Rotate in the YW plane.
	case planes&2 > 0:
		c, s = math.Cos(Θ[1]), math.Sin(Θ[1])
		b = Mat4x4{
			Mat1x4{1, 0, 0, 0},
			Mat1x4{0, c, 0, -s},
			Mat1x4{0, 0, 1, 0},
			Mat1x4{0, s, 0, c},
		}

		if planes&1 == 0 {
			a = b
		} else {
			a = a.MultiplyMat4x4(b)
		}

		fallthrough
	//Rotate in the YZ plane.
	case planes&4 > 0:
		c, s = math.Cos(Θ[2]), math.Sin(Θ[2])
		b = Mat4x4{
			Mat1x4{1, 0, 0, 0},
			Mat1x4{0, c, s, 0},
			Mat1x4{0, -s, c, 0},
			Mat1x4{0, 0, 0, 1},
		}

		if planes&3 == 0 {
			a = b
		} else {
			a = a.MultiplyMat4x4(b)
		}

		fallthrough
	//Rotate in the XW plane.
	case planes&8 > 0:
		c, s = math.Cos(Θ[3]), math.Sin(Θ[3])
		b = Mat4x4{
			Mat1x4{c, 0, 0, s},
			Mat1x4{0, 1, 0, 0},
			Mat1x4{0, 0, 1, 0},
			Mat1x4{-s, 0, 0, c},
		}

		if planes&7 == 0 {
			a = b
		} else {
			a = a.MultiplyMat4x4(b)
		}

		fallthrough
	//Rotate in the XZ plane.
	case planes&16 > 0:
		c, s = math.Cos(Θ[4]), math.Sin(Θ[4])
		b = Mat4x4{
			Mat1x4{c, 0, -s, 0},
			Mat1x4{0, 1, 0, 0},
			Mat1x4{s, 0, c, 0},
			Mat1x4{0, 0, 0, 1},
		}

		if planes&15 == 0 {
			a = b
		} else {
			a = a.MultiplyMat4x4(b)
		}

		fallthrough
	//Rotate in the XY plane.
	case planes&32 > 0:
		c, s = math.Cos(Θ[5]), math.Sin(Θ[5])
		b = Mat4x4{
			Mat1x4{c, s, 0, 0},
			Mat1x4{-s, c, 0, 0},
			Mat1x4{0, 0, 1, 0},
			Mat1x4{0, 0, 0, 1},
		}

		if planes&31 == 0 {
			a = b
		} else {
			a = a.MultiplyMat4x4(b)
		}
	}

	return a
}

//Rotate in the XY plane.
func LinearRotate4XY(Θ float64) Mat4x4 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat4x4{
		Mat1x4{c, s, 0, 0},
		Mat1x4{-s, c, 0, 0},
		Mat1x4{0, 0, 1, 0},
		Mat1x4{0, 0, 0, 1},
	}
}

//Rotate in the XZ plane.
func LinearRotate4XZ(Θ float64) Mat4x4 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat4x4{
		Mat1x4{c, 0, -s, 0},
		Mat1x4{0, 1, 0, 0},
		Mat1x4{s, 0, c, 0},
		Mat1x4{0, 0, 0, 1},
	}
}

//Rotate in the XW plane.
func LinearRotate4XW(Θ float64) Mat4x4 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat4x4{
		Mat1x4{c, 0, 0, s},
		Mat1x4{0, 1, 0, 0},
		Mat1x4{0, 0, 1, 0},
		Mat1x4{-s, 0, 0, c},
	}
}

//Rotate in the YZ plane.
func LinearRotate4YZ(Θ float64) Mat4x4 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat4x4{
		Mat1x4{1, 0, 0, 0},
		Mat1x4{0, c, s, 0},
		Mat1x4{0, -s, c, 0},
		Mat1x4{0, 0, 0, 1},
	}
}

//Rotate in the YW plane.
func LinearRotate4YW(Θ float64) Mat4x4 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat4x4{
		Mat1x4{1, 0, 0, 0},
		Mat1x4{0, c, 0, -s},
		Mat1x4{0, 0, 1, 0},
		Mat1x4{0, s, 0, c},
	}
}

//Rotate in the ZW plane.
func LinearRotate4ZW(Θ float64) Mat4x4 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat4x4{
		Mat1x4{1, 0, 0, 0},
		Mat1x4{0, 1, 0, 0},
		Mat1x4{0, 0, c, -s},
		Mat1x4{0, 0, s, c},
	}
}

func LinearScale1(s Vec1) Mat1x1 {
	return Mat1x1{s[0]}
}

func LinearScale2(s Vec2) Mat2x2 {
	return Mat2x2{
		Mat1x2{s[0], 0},
		Mat1x2{0, s[1]},
	}
}

func LinearScale3(s Vec3) Mat3x3 {
	return Mat3x3{
		Mat1x3{s[0], 0, 0},
		Mat1x3{0, s[1], 0},
		Mat1x3{0, 0, s[2]},
	}
}

func LinearScale4(s Vec4) Mat4x4 {
	return Mat4x4{
		Mat1x4{s[0], 0, 0, 0},
		Mat1x4{0, s[1], 0, 0},
		Mat1x4{0, 0, s[2], 0},
		Mat1x4{0, 0, 0, s[3]},
	}
}

func LinearShear2(τ Vec2) Mat2x2 {
	return Mat2x2{
		Mat1x2{1, τ[0]},
		Mat1x2{τ[1], 1},
	}
}

func LinearShear3(τx, τy, τz Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{1, τy[0], τz[0]},
		Mat1x3{τx[0], 1, τz[1]},
		Mat1x3{τx[1], τy[1], 1},
	}
}

func LinearShear4(τx, τy, τz, τw Vec3) Mat4x4 {
	return Mat4x4{
		Mat1x4{1, τy[0], τz[0], τw[0]},
		Mat1x4{τx[0], 1, τz[1], τw[1]},
		Mat1x4{τx[1], τy[1], 1, τw[2]},
		Mat1x4{τx[2], τy[2], τz[2], 1},
	}
}

func LinearTranslate1(δ Vec1) Mat1x1 {
	return Mat1x1{δ[0]}
}

func LinearTranslate2(δ Vec2) Mat2x2 {
	return Mat2x2{
		Mat1x2{1, δ[0]},
		Mat1x2{0, δ[1]},
	}
}

func LinearTranslate3(δ Vec3) Mat3x3 {
	return Mat3x3{
		Mat1x3{1, 0, δ[0]},
		Mat1x3{0, 1, δ[1]},
		Mat1x3{0, 0, δ[2]},
	}
}

func LinearTranslate4(δ Vec4) Mat4x4 {
	return Mat4x4{
		Mat1x4{1, 0, 0, δ[0]},
		Mat1x4{0, 1, 0, δ[1]},
		Mat1x4{0, 0, 1, δ[2]},
		Mat1x4{0, 0, 0, δ[3]},
	}
}

func Reflect2() Mat2x2 {
	return Mat2x2{
		Mat1x2{-1, 0},
		Mat1x2{0, 1},
	}
}

func Reflect3(Γ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{Γ[0], 0, 0},
		Mat1x3{0, Γ[1], 0},
		Mat1x3{0, 0, 1},
	}
}

//Takes in a Vec3 comprised of 1s and/or -1s. If the value in the vector are anything else, scaling may take place
func Reflect4(Γ Vec3) Mat4x4 {
	return Mat4x4{
		Mat1x4{Γ[0], 0, 0, 0},
		Mat1x4{0, Γ[1], 0, 0},
		Mat1x4{0, 0, Γ[2], 0},
		Mat1x4{0, 0, 0, 1},
	}
}

func Rotate3(Θ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{math.Cos(Θ), -math.Sin(Θ), 0},
		Mat1x3{math.Sin(Θ), math.Cos(Θ), 0},
		Mat1x3{0, 0, 1},
	}
}

func Rotate4(Θ Vec3) Mat4x4 {
	var planes int
	for i, θ := range Θ {
		if θ != 0 {
			planes |= 1 << uint(i)
		}
	}

	if planes == 0 {
		return Identity4()
	}

	var c, s float64
	var a, b Mat4x4

	switch {
	//X-axis rotation
	//Rotate in the YZ plane.
	case planes&1 > 0:
		c, s = math.Cos(Θ[0]), math.Sin(Θ[0])
		a = Mat4x4{
			Mat1x4{1, 0, 0, 0},
			Mat1x4{0, c, -s, 0},
			Mat1x4{0, s, c, 0},
			Mat1x4{0, 0, 0, 1},
		}

		fallthrough
		//Y-axis rotation
		//Rotate in the XZ plane.
	case planes&2 > 0:
		c, s = math.Cos(Θ[1]), math.Sin(Θ[1])
		b = Mat4x4{
			Mat1x4{c, 0, s, 0},
			Mat1x4{0, 1, 0, 0},
			Mat1x4{-s, 0, c, 0},
			Mat1x4{0, 0, 0, 1},
		}

		if planes&1 == 0 {
			a = b
		} else {
			a = a.MultiplyMat4x4(b)
		}

		fallthrough
	//Z-axis rotation
	//Rotate in the XY plane.
	case planes&4 > 0:
		c, s = math.Cos(Θ[2]), math.Sin(Θ[2])

		b = Mat4x4{
			Mat1x4{c, -s, 0, 0},
			Mat1x4{s, c, 0, 0},
			Mat1x4{0, 0, 1, 0},
			Mat1x4{0, 0, 0, 1},
		}

		if planes&3 == 0 {
			a = b
		} else {
			a = a.MultiplyMat4x4(b)
		}
	}

	return a
}

func Rotate4X(Θ float64) Mat4x4 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat4x4{
		Mat1x4{1, 0, 0, 0},
		Mat1x4{0, c, -s, 0},
		Mat1x4{0, s, c, 0},
		Mat1x4{0, 0, 0, 1},
	}
}

func Rotate4Y(Θ float64) Mat4x4 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat4x4{
		Mat1x4{c, 0, s, 0},
		Mat1x4{0, 1, 0, 0},
		Mat1x4{-s, 0, c, 0},
		Mat1x4{0, 0, 0, 1},
	}
}

func Rotate4Z(Θ float64) Mat4x4 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat4x4{
		Mat1x4{c, -s, 0, 0},
		Mat1x4{s, c, 0, 0},
		Mat1x4{0, 0, 1, 0},
		Mat1x4{0, 0, 0, 1},
	}
}

func Scale2(Δ Vec1) Mat2x2 {
	return Mat2x2{
		Mat1x2{Δ[0], 0},
		Mat1x2{0, 1},
	}
}

func Scale3(Δ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{Δ[0], 0, 0},
		Mat1x3{0, Δ[1], 0},
		Mat1x3{0, 0, 1},
	}
}

func Scale4(Δ Vec3) Mat4x4 {
	return Mat4x4{
		Mat1x4{Δ[0], 0, 0, 0},
		Mat1x4{0, Δ[1], 0, 0},
		Mat1x4{0, 0, Δ[2], 0},
		Mat1x4{0, 0, 0, 1},
	}
}

func Shear3(τ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{1, τ[0], 0},
		Mat1x3{τ[1], 1, 0},
		Mat1x3{0, 0, 1},
	}
}

func Shear4(τx, τy, τz Vec2) Mat4x4 {
	return Mat4x4{
		Mat1x4{1, τy[0], τz[0], 0},
		Mat1x4{τx[0], 1, τz[1], 0},
		Mat1x4{τx[1], τy[1], 1, 0},
		Mat1x4{0, 0, 0, 1},
	}
}

func Translate2(δ Vec1) Mat2x2 {
	return Mat2x2{
		Mat1x2{1, δ[0]},
		Mat1x2{0, 1},
	}
}

func Translate3(δ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{1, 0, δ[0]},
		Mat1x3{0, 1, δ[1]},
		Mat1x3{0, 0, 1},
	}
}

func Translate4(δ Vec3) Mat4x4 {
	return Mat4x4{
		Mat1x4{1, 0, 0, δ[0]},
		Mat1x4{0, 1, 0, δ[1]},
		Mat1x4{0, 0, 1, δ[2]},
		Mat1x4{0, 0, 0, 1},
	}
}
