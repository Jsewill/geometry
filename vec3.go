/* */
package geometry

import (
	"math"
)

type Vec3 [3]float64

func (v Vec3) Add(a Vec3) Vec3 {
	return Vec3{v[0] + a.X(), v[1] + a.Y(), v[2] + a.Z()}
}

func (v Vec3) B() float64 {
	return v[2]
}

func (v Vec3) ByRow(m Mat1x3) Mat3x3 {
	return Mat3x3{
		Mat1x3{v[0] * m[0], v[0] * m[1], v[0] * m[2]},
		Mat1x3{v[1] * m[0], v[1] * m[1], v[1] * m[2]},
		Mat1x3{v[2] * m[0], v[2] * m[1], v[2] * m[2]},
	}
}

func (v Vec3) Component(c int) float64 {
	if c > v.Len()-1 {
		return 0
	}

	return v[c]
}

func (v Vec3) Components() []float64 {
	return v[:]
}

func (v Vec3) Cross(c Vec3) Vec3 {
	return Vec3{v[1]*c.Z() - v[2]*c.Y(), v[2]*c.X() - v[0]*c.Z(), v[0]*c.Y() - v[1]*c.X()}
}

func (v Vec3) Dot(d Vec3) float64 {
	return v[0]*d.X() + v[1]*d.Y() + v[2]*d.Z()
}

//@TODO Implement approximate equality comparison, as float calculations tend to differ.
func (v Vec3) Equal(c Vec3) bool {
	return v[0] == c.X() && v[1] == c.Y() && v[2] == c.Z()
}

func (v Vec3) G() float64 {
	return v[1]
}

func (v Vec3) Len() int {
	return 3
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec3) Multiply(m float64) Vec3 {
	return Vec3{v[0] * m, v[1] * m, v[2] * m}
}

func (v Vec3) Normalize() Vec3 {
	d := 1.0 / v.Length()
	return v.Multiply(d)
}

func (v Vec3) P() float64 {
	return v[2]
}

func (v Vec3) R() float64 {
	return v[0]
}

func (v Vec3) RG() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec3) RGB() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec3) S() float64 {
	return v[0]
}

func (v Vec3) Square() float64 {
	return v.Dot(v)
}

func (v Vec3) ST() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec3) STP() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec3) Subtract(s Vec3) Vec3 {
	return Vec3{v[0] - s.X(), v[1] - s.Y(), v[2] - s.Z()}
}

func (v Vec3) Sum() float64 {
	return v[0] + v[1] + v[2]
}

func (v Vec3) T() float64 {
	return v[1]
}

func (v Vec3) U() float64 {
	return v[0]
}

func (v Vec3) UV() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec3) UVP() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec3) V() float64 {
	return v[1]
}

func (v Vec3) Vec1() Vec1 {
	return Vec1{v[0]}
}

func (v Vec3) Vec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec3) VecN() VecN {
	return VecN(v[:])
}

func (v Vec3) X() float64 {
	return v[0]
}

func (v Vec3) XY() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec3) XYZ() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec3) Y() float64 {
	return v[1]
}

func (v Vec3) Z() float64 {
	return v[2]
}
