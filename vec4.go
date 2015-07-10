/* */
package geometry

import (
	"math"
)

type Vec4 [4]float64

func (v Vec4) A() float64 {
	return v[3]
}

func (v Vec4) Add(a Vec4) Vec4 {
	return Vec4{v[0] + a.X(), v[1] + a.Y(), v[2] + a.Z(), v[3] + v[3]}
}

func (v Vec4) B() float64 {
	return v[2]
}

func (v Vec4) ByRow(m Mat1x4) Mat4x4 {
	return Mat4x4{
		Mat1x4{v[0] * m[0], v[0] * m[1], v[0] * m[2], v[0] * m[3]},
		Mat1x4{v[1] * m[0], v[1] * m[1], v[1] * m[2], v[1] * m[3]},
		Mat1x4{v[2] * m[0], v[2] * m[1], v[2] * m[2], v[2] * m[3]},
		Mat1x4{v[3] * m[0], v[3] * m[1], v[3] * m[2], v[3] * m[3]},
	}
}

func (v Vec4) Component(c int) float64 {
	if c > v.Len()-1 {
		return 0
	}

	return v[c]
}

func (v Vec4) Components() []float64 {
	return v[:]
}

//May or may not actually be correct. While a 4D vector cross product does not have the properties that 0, 1, 3, and 7 dimension vectors do, I wanted to provide it.
func (v Vec4) Cross(c Vec4) Vec4 {
	return Vec4{v[2]*c.W() - v[3]*c.Z(), v[1]*c.Z() - v[2]*c.Y(), v[2]*c.X() - v[0]*c.Z(), v[0]*c.Y() - v[1]*c.X()}
}

func (v Vec4) Dot(d Vec4) float64 {
	return v[0]*d.X() + v[1]*d.Y() + v[2]*d.Z() + v[3]*d.W()
}

//@TODO Implement approximate equality comparison, as float calculations tend to differ.
func (v Vec4) Equal(c Vec4) bool {
	return v[0] == c.X() && v[1] == c.Y() && v[2] == c.Z() && v[3] == c.W()
}

func (v Vec4) G() float64 {
	return v[1]
}

func (v Vec4) Len() int {
	return 4
}

func (v Vec4) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec4) Multiply(m float64) Vec4 {
	return Vec4{v[0] * m, v[1] * m, v[2] * m, v[3] * m}
}

func (v Vec4) Normalize() Vec4 {
	d := 1.0 / v.Length()
	return v.Multiply(d)
}

func (v Vec4) P() float64 {
	return v[2]
}

func (v Vec4) Q() float64 {
	return v[3]
}

func (v Vec4) R() float64 {
	return v[0]
}

func (v Vec4) RG() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec4) RGB() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec4) RGBA() Vec4 {
	return Vec4{v[0], v[1], v[2], v[3]}
}

func (v Vec4) S() float64 {
	return v[0]
}

func (v Vec4) Square() float64 {
	return v.Dot(v)
}

func (v Vec4) ST() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec4) STP() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec4) STPQ() Vec4 {
	return Vec4{v[0], v[1], v[2], v[3]}
}

func (v Vec4) Subtract(s Vec4) Vec4 {
	return Vec4{v[0] - s.X(), v[1] - s.Y(), v[2] - s.Z(), v[3] - s.W()}
}

func (v Vec4) Sum() float64 {
	return v[0] + v[1] + v[2] + v[3]
}

func (v Vec4) T() float64 {
	return v[1]
}

func (v Vec4) U() float64 {
	return v[0]
}

func (v Vec4) UV() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec4) UVP() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec4) UVPQ() Vec4 {
	return Vec4{v[0], v[1], v[2], v[3]}
}

func (v Vec4) V() float64 {
	return v[1]
}

func (v Vec4) Vec1() Vec1 {
	return Vec1{v[0]}
}

func (v Vec4) Vec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec4) Vec3() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec4) VecN() VecN {
	return VecN(v[:])
}

func (v Vec4) W() float64 {
	return v[3]
}

func (v Vec4) X() float64 {
	return v[0]
}

func (v Vec4) XY() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec4) XYZ() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec4) XYZW() Vec4 {
	return Vec4{v[0], v[1], v[2], v[3]}
}

func (v Vec4) Y() float64 {
	return v[1]
}

func (v Vec4) Z() float64 {
	return v[2]
}
