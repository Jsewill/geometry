/* */
package geometry

import (
	"math"
)

type Vec3 [3]float64

func (v Vec3) Add(a Vec3) Vec3 {
	return Vec3{v[0] + a[0], v[1] + a[1], v[2] + a[2]}
}

func (v Vec3) AddScalar(a float64) Vec3 {
	return Vec3{v[0] + a, v[1] + a, v[2] + a}
}

func (v Vec3) B() float64 {
	return v[2]
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
	return Vec3{v[1]*c[2] - v[2]*c[1], v[2]*c[0] - v[0]*c[2], v[0]*c[1] - v[1]*c[0]}
}

func (v Vec3) Divide(d Vec3) Vec3 {
	r := Vec3{}
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
	if d[2] == 0 || v[2] == 0 {
		r[2] = 0
	} else {
		r[2] = v[2] / d[2]
	}
	return r
}

func (v Vec3) DivideScalar(d float64) Vec3 {
	r := Vec3{}
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
	if v[2] == 0 {
		r[2] = 0
	} else {
		r[2] = v[2] / d
	}
	return r
}

func (v Vec3) Dot(d Vec3) float64 {
	return v[0]*d[0] + v[1]*d[1] + v[2]*d[2]
}

//@TODO Implement approximate equality comparison, as float calculations tend to differ.
func (v Vec3) Equal(c Vec3) bool {
	return v[0] == c[0] && v[1] == c[1] && v[2] == c[2]
}

func (v Vec3) Expand() Vec4 {
	return Vec4{v[0], v[1], v[2], 0}
}

func (v Vec3) G() float64 {
	return v[1]
}

func (v Vec3) Homogenize() Vec4 {
	return Vec4{v[0], v[1], v[2], 1}
}

func (v Vec3) Len() int {
	return 3
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec3) MultiplyMat1x1(m Mat1x1) Mat3x1 {
	return Mat3x1{
		Mat1x1{v[0] * m[0]},
		Mat1x1{v[1] * m[0]},
		Mat1x1{v[2] * m[0]},
	}
}

func (v Vec3) MultiplyMat1x2(m Mat1x2) Mat3x2 {
	return Mat3x2{
		Mat1x2{v[0] * m[0], v[0] * m[1]},
		Mat1x2{v[1] * m[0], v[1] * m[1]},
		Mat1x2{v[2] * m[0], v[2] * m[1]},
	}
}

func (v Vec3) MultiplyMat1x3(m Mat1x3) Mat3x3 {
	return Mat3x3{
		Mat1x3{v[0] * m[0], v[0] * m[1], v[0] * m[2]},
		Mat1x3{v[1] * m[0], v[1] * m[1], v[1] * m[2]},
		Mat1x3{v[2] * m[0], v[2] * m[1], v[2] * m[2]},
	}
}

func (v Vec3) MultiplyMat1x4(m Mat1x4) Mat3x4 {
	return Mat3x4{
		Mat1x4{v[0] * m[0], v[0] * m[1], v[0] * m[2], v[0] * m[3]},
		Mat1x4{v[1] * m[0], v[1] * m[1], v[1] * m[2], v[1] * m[3]},
		Mat1x4{v[2] * m[0], v[2] * m[1], v[2] * m[2], v[2] * m[3]},
	}
}

func (v Vec3) MultiplyScalar(m float64) Vec3 {
	return Vec3{v[0] * m, v[1] * m, v[2] * m}
}

//Component-wise vector multiplication
func (a Vec3) MultiplyVec3(b Vec3) Vec3 {
	return Vec3{a[0] * b[0], a[1] * b[1], a[2] * b[2]}
}

func (v Vec3) Normalize() Vec3 {
	d := 1.0 / v.Length()
	return v.MultiplyScalar(d)
}

func (v Vec3) P() float64 {
	return v[2]
}

func (v Vec3) R() float64 {
	return v[0]
}

func (v Vec3) Reflect(n Vec3) Vec3 {
	return n.MultiplyScalar(-2 * v.Dot(n)).Subtract(v)
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
	return Vec3{v[0] - s[0], v[1] - s[1], v[2] - s[2]}
}

func (v Vec3) SubtractScalar(s float64) Vec3 {
	return Vec3{v[0] - s, v[1] - s, v[2] - s}
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

func (v Vec3) Vec3() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v Vec3) Vec4(w float64) Vec4 {
	return Vec4{v[0], v[1], v[2], w}
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
