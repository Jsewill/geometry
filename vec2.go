/* */
package geometry

import (
	"math"
)

type Vec2 [2]float64

func (v Vec2) Add(a Vec2) Vec2 {
	return Vec2{v[0] + a.X(), v[1] + a.Y()}
}

func (v Vec2) ByRow(m Mat1x2) Mat2x2 {
	return Mat2x2{
		Mat1x2{v[0] * m[0], v[0] * m[1]},
		Mat1x2{v[1] * m[0], v[1] * m[1]},
	}
}

func (v Vec2) Component(c int) float64 {
	if c > v.Len()-1 {
		return 0
	}

	return v[c]
}

func (v Vec2) Components() []float64 {
	return v[:]
}

func (v Vec2) Cross(c Vec2) Vec2 {
	return Vec2{v[1]*c.X() - v[0]*c.Y(), v[0]*c.Y() - v[1]*c.X()}
}

func (v Vec2) Dot(d Vec2) float64 {
	return v[0]*d.X() + v[1]*d.Y()
}

//@TODO Implement approximate equality comparison, as float calculations tend to differ.
func (v Vec2) Equal(c Vec2) bool {
	return v[0] == c.X() && v[1] == c.Y()
}

func (v Vec2) G() float64 {
	return v[1]
}

//Returns the size of the vector
func (v Vec2) Len() int {
	return 2
}

//Returns the vector length; e.g., Hypotenuse
func (v Vec2) Length() float64 {
	return math.Hypot(v[0], v[1])
}

func (v Vec2) Multiply(m float64) Vec2 {
	return Vec2{v[0] * m, v[1] * m}
}

func (v Vec2) Normalize() Vec2 {
	d := 1.0 / v.Length()
	return v.Multiply(d)
}

func (v Vec2) R() float64 {
	return v[0]
}

func (v Vec2) RG() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec2) S() float64 {
	return v[0]
}

//Returns the Dot product of itself.
func (v Vec2) Square() float64 {
	return v.Dot(v)
}

//Returns a vector squared component-wise. May not have any use.
func (v Vec2) Squared() Vec2 {
	return Vec2{v[0] * v[0], v[1] * v[1]}
}

func (v Vec2) ST() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec2) Subtract(s Vec2) Vec2 {
	return Vec2{v[0] - s.X(), v[1] - s.Y()}
}

func (v Vec2) Sum() float64 {
	return v[0] + v[1]
}

func (v Vec2) T() float64 {
	return v[1]
}

func (v Vec2) U() float64 {
	return v[0]
}

func (v Vec2) UV() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec2) V() float64 {
	return v[1]
}

func (v Vec2) Vec1() Vec1 {
	return Vec1{v[0]}
}

func (v Vec2) VecN() VecN {
	return VecN(v[:])
}

func (v Vec2) X() float64 {
	return v[0]
}

func (v Vec2) XY() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec2) Y() float64 {
	return v[1]
}
