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
	return Vec4{v[0] + a[0], v[1] + a[1], v[2] + a[2], v[3] + a[3]}
}

func (v Vec4) AddScalar(a float64) Vec4 {
	return Vec4{v[0] + a, v[1] + a, v[2] + a, v[3] + a}
}

func (v Vec4) B() float64 {
	return v[2]
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

//May or may not actually be correct. While a 4D vector cross product does not have the properties that 0, 1, 3, and 7 dimension vectors do, I wanted to define it.
func (v Vec4) Cross(c Vec4) Vec4 {
	return Vec4{v[2]*c[3] - v[3]*c[2], v[1]*c[2] - v[2]*c[1], v[2]*c[0] - v[0]*c[2], v[0]*c[1] - v[1]*c[0]}
}

func (v Vec4) Divide(d Vec4) Vec4 {
	r := Vec4{}
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
	if d[3] == 0 || v[3] == 0 {
		r[3] = 0
	} else {
		r[3] = v[3] / d[3]
	}
	return r
}

func (v Vec4) DivideScalar(d float64) Vec4 {
	r := Vec4{}
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
	if v[3] == 0 {
		r[3] = 0
	} else {
		r[3] = v[3] / d
	}
	return r
}

func (v Vec4) Dot(d Vec4) float64 {
	return v[0]*d[0] + v[1]*d[1] + v[2]*d[2] + v[3]*d[3]
}

//@TODO Implement approximate equality comparison, as float calculations tend to differ.
func (v Vec4) Equal(c Vec4) bool {
	return v[0] == c[0] && v[1] == c[1] && v[2] == c[2] && v[3] == c[3]
}

func (v Vec4) Expand() VecN {
	return VecN(append(v[:], 0))
}

func (v Vec4) G() float64 {
	return v[1]
}

func (v Vec4) Homogenize() VecN {
	return VecN(append(v[:], 1))
}

func (v Vec4) Len() int {
	return 4
}

func (v Vec4) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec4) MultiplyMat1x1(m Mat1x1) Mat4x1 {
	return Mat4x1{
		Mat1x1{v[0] * m[0]},
		Mat1x1{v[1] * m[0]},
		Mat1x1{v[2] * m[0]},
		Mat1x1{v[3] * m[0]},
	}
}

func (v Vec4) MultiplyMat1x2(m Mat1x2) Mat4x2 {
	return Mat4x2{
		Mat1x2{v[0] * m[0], v[0] * m[1]},
		Mat1x2{v[1] * m[0], v[1] * m[1]},
		Mat1x2{v[2] * m[0], v[2] * m[1]},
		Mat1x2{v[3] * m[0], v[3] * m[1]},
	}
}

func (v Vec4) MultiplyMat1x3(m Mat1x3) Mat4x3 {
	return Mat4x3{
		Mat1x3{v[0] * m[0], v[0] * m[1], v[0] * m[2]},
		Mat1x3{v[1] * m[0], v[1] * m[1], v[1] * m[2]},
		Mat1x3{v[2] * m[0], v[2] * m[1], v[2] * m[2]},
		Mat1x3{v[3] * m[0], v[3] * m[1], v[3] * m[2]},
	}
}

func (v Vec4) MultiplyMat1x4(m Mat1x4) Mat4x4 {
	return Mat4x4{
		Mat1x4{v[0] * m[0], v[0] * m[1], v[0] * m[2], v[0] * m[3]},
		Mat1x4{v[1] * m[0], v[1] * m[1], v[1] * m[2], v[1] * m[3]},
		Mat1x4{v[2] * m[0], v[2] * m[1], v[2] * m[2], v[2] * m[3]},
		Mat1x4{v[3] * m[0], v[3] * m[1], v[3] * m[2], v[3] * m[3]},
	}
}

func (v Vec4) MultiplyScalar(m float64) Vec4 {
	return Vec4{v[0] * m, v[1] * m, v[2] * m, v[3] * m}
}

//Component-wise vector multiplication
func (a Vec4) MultiplyVec4(b Vec4) Vec4 {
	return Vec4{a[0] * b[0], a[1] * b[1], a[2] * b[2], a[3] * b[3]}
}

func (v Vec4) Normalize() Vec4 {
	d := 1.0 / v.Length()
	return v.MultiplyScalar(d)
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

func (v Vec4) Reflect(n Vec4) Vec4 {
	return n.MultiplyScalar(-2 * v.Dot(n)).Subtract(v)
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
	return Vec4{v[0] - s[0], v[1] - s[1], v[2] - s[2], v[3] - s[3]}
}

func (v Vec4) SubtractScalar(s float64) Vec4 {
	return Vec4{v[0] - s, v[1] - s, v[2] - s, v[3] - s}
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

func (v Vec4) Vec4() Vec4 {
	return Vec4{v[0], v[1], v[2], v[3]}
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
