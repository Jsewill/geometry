/* */
package geometry

import (
	"math"
)

type Vec2 [2]float64

func (v Vec2) Add(a Vec2) Vec2 {
	return Vec2{v[0] + a[0], v[1] + a[1]}
}

func (v Vec2) AddScalar(a float64) Vec2 {
	return Vec2{v[0] + a, v[1] + a}
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
	return Vec2{v[1]*c[0] - v[0]*c[1], v[0]*c[1] - v[1]*c[0]}
}

func (v Vec2) Divide(d Vec2) Vec2 {
	r := Vec2{}
	if d[0] == 0 || v[0] == 0 {
		r[0] = 0
	} else {
		r[0] = v[0] / d[0]
	}
	if d[1] == 0 || v[1] == 0 {
		r[1] = 0
	} else {
		r[1] = v[1] / d[1]
	}
	return r
}

func (v Vec2) DivideScalar(d float64) Vec2 {
	r := Vec2{}
	if d == 0 {
		return r
	}
	if v[0] == 0 {
		r[0] = 0
	} else {
		r[0] = v[0] / d
	}
	if v[1] == 0 {
		r[1] = 0
	} else {
		r[1] = v[1] / d
	}
	return r
}

func (v Vec2) Dot(d Vec2) float64 {
	return v[0]*d[0] + v[1]*d[0]
}

//@TODO Implement approximate equality comparison, as float calculations tend to differ.
func (v Vec2) Equal(c Vec2) bool {
	return v[0] == c[0] && v[1] == c[1]
}

func (v Vec2) Expand() Vec3 {
	return Vec3{v[0], v[1], 0}
}

func (v Vec2) G() float64 {
	return v[1]
}

func (v Vec2) Homogenize() Vec3 {
	return Vec3{v[0], v[1], 1}
}

//Returns the size of the vector
func (v Vec2) Len() int {
	return 2
}

//Returns the vector length; e.g., Hypotenuse
func (v Vec2) Length() float64 {
	return math.Hypot(v[0], v[1])
}

func (v Vec2) MultiplyMat1x1(m Mat1x1) Mat2x1 {
	return Mat2x1{
		Mat1x1{v[0] * m[0]},
		Mat1x1{v[1] * m[0]},
	}
}

func (v Vec2) MultiplyMat1x2(m Mat1x2) Mat2x2 {
	return Mat2x2{
		Mat1x2{v[0] * m[0], v[0] * m[1]},
		Mat1x2{v[1] * m[0], v[1] * m[1]},
	}
}

func (v Vec2) MultiplyMat1x3(m Mat1x3) Mat2x3 {
	return Mat2x3{
		Mat1x3{v[0] * m[0], v[0] * m[1], v[0] * m[2]},
		Mat1x3{v[1] * m[0], v[1] * m[1], v[1] * m[2]},
	}
}

func (v Vec2) MultiplyMat1x4(m Mat1x4) Mat2x4 {
	return Mat2x4{
		Mat1x4{v[0] * m[0], v[0] * m[1], v[0] * m[2], v[0] * m[3]},
		Mat1x4{v[1] * m[0], v[1] * m[1], v[1] * m[2], v[1] * m[3]},
	}
}

func (v Vec2) MultiplyScalar(m float64) Vec2 {
	return Vec2{v[0] * m, v[1] * m}
}

//Component-wise vector multiplication
func (a Vec2) MultiplyVec2(b Vec2) Vec2 {
	return Vec2{a[0] * b[0], a[1] * b[1]}
}

func (v Vec2) Normalize() Vec2 {
	d := 1.0 / v.Length()
	return v.MultiplyScalar(d)
}

func (v Vec2) R() float64 {
	return v[0]
}

func (v Vec2) Reflect(n Vec2) Vec2 {
	return n.MultiplyScalar(-2 * v.Dot(n)).Subtract(v)
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
	return Vec2{v[0] - s[0], v[1] - s[1]}
}

func (v Vec2) SubtractScalar(s float64) Vec2 {
	return Vec2{v[0] - s, v[1] - s}
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

func (v Vec2) Vec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec2) Vec3(z float64) Vec3 {
	return Vec3{v[0], v[1], z}
}

func (v Vec2) Vec4(z, w float64) Vec4 {
	return Vec4{v[0], v[1], z, w}
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
