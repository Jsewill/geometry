/* */
package geometry

type Vec1 [1]float64

func (v Vec1) Add(a Vec1) Vec1 {
	return Vec1{v[0] + a[0]}
}

func (v Vec1) AddScalar(a float64) Vec1 {
	return Vec1{v[0] + a}
}

func (v Vec1) Component(c int) float64 {
	return v[0]
}

func (v Vec1) Components() []float64 {
	return v[:]
}

func (v Vec1) Divide(d Vec1) Vec1 {
	if d[0] == 0 {
		return Vec1{}
	}
	return Vec1{v[0] / d[0]}
}

func (v Vec1) DivideScalar(d float64) Vec1 {
	if d == 0 {
		return Vec1{}
	}
	return Vec1{v[0] / d}
}

func (v Vec1) Dot(d Vec1) float64 {
	return v[0] * d[0]
}

//@TODO Implement approximate equality comparison, as float calculations tend to differ.
func (v Vec1) Equal(c Vec1) bool {
	return v[0] == c[0]
}

func (v Vec1) Expand() Vec2 {
	return Vec2{v[0], 0}
}

func (v Vec1) Homogenize() Vec2 {
	return Vec2{v[0], 1}
}

func (v Vec1) Len() int {
	return 1
}

func (v Vec1) MultiplyMat1x1(m Mat1x1) Mat1x1 {
	return Mat1x1{v[0] * m[0]}
}

func (v Vec1) MultiplyMat1x2(m Mat1x2) Mat1x2 {
	return Mat1x2{v[0] * m[0], v[0] * m[1]}
}

func (v Vec1) MultiplyMat1x3(m Mat1x3) Mat1x3 {
	return Mat1x3{v[0] * m[0], v[0] * m[1], v[0] * m[2]}
}

func (v Vec1) MultiplyMat1x4(m Mat1x4) Mat1x4 {
	return Mat1x4{v[0] * m[0], v[0] * m[1], v[0] * m[2], v[0] * m[3]}
}

func (v Vec1) MultiplyScalar(m float64) Vec1 {
	return Vec1{v[0] * m}
}

//Component-wise vector multiplication
func (a Vec1) MultiplyVec1(b Vec1) Vec1 {
	return Vec1{a[0] * b[0]}
}

//Normalizing a 1D vector does nothing, so does this.
func (v Vec1) Normalize() Vec1 {
	return Vec1{v[0]}
}

func (v Vec1) R() float64 {
	return v[0]
}

func (v Vec1) Reflect(n Vec1) Vec1 {
	return n.MultiplyScalar(-2 * v.Dot(n)).Subtract(v)
}

func (v Vec1) S() float64 {
	return v[0]
}

func (v Vec1) Subtract(s Vec1) Vec1 {
	return Vec1{v[0] - s[0]}
}

func (v Vec1) SubtractScalar(s float64) Vec1 {
	return Vec1{v[0] - s}
}

//Returns the first component.
func (v Vec1) Sum() float64 {
	return v[0]
}

func (v Vec1) Square() float64 {
	return v[0] * v[0]
}

func (v Vec1) Squared() Vec1 {
	return Vec1{v[0] * v[0]}
}

func (v Vec1) Vec1() Vec1 {
	return Vec1{v[0]}
}

func (v Vec1) Vec2(y float64) Vec2 {
	return Vec2{v[0], y}
}

func (v Vec1) Vec3(y, z float64) Vec3 {
	return Vec3{v[0], y, z}
}

func (v Vec1) Vec4(y, z, w float64) Vec4 {
	return Vec4{v[0], y, z, w}
}

func (v Vec1) VecN() VecN {
	return VecN(v[:])
}

func (v Vec1) X() float64 {
	return v[0]
}
