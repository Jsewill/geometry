/* Mat1xN implements row vectors. Whereas the VecN types are akin to column vectors */
package geometry

import (
	_ "text/tabwriter"
)

type Mat1x1 [1]float64

func (a Mat1x1) Add(b Mat1x1) Mat1x1 {
	return Mat1x1{a[0] + b[0]}
}

func (m Mat1x1) AddScalar(f float64) Mat1x1 {
	return Mat1x1{
		m[0] + f,
	}
}

func (m Mat1x1) Concatenate(f float64) Mat1x2 {
	return Mat1x2{m[0], f}
}

func (m Mat1x1) Homogenize() Mat2x2 {
	return Mat2x2{
		Mat1x2{m[0], 0},
		Mat1x2{0, 1},
	}
}

func (m Mat1x1) Len() int {
	return 1
}

func (m Mat1x1) Inverse() Mat1x1 {
	return Mat1x1{m[0]}
}

func (m Mat1x1) MultiplyScalar(f float64) Mat1x1 {
	return Mat1x1{m[0] * f}
}

func (a Mat1x1) MultiplyMat1x1(b Mat1x1) Mat1x1 {
	return Mat1x1{a[0] * b[0]}
}

func (a Mat1x1) MultiplyMat1x2(b Mat1x2) Mat1x2 {
	return Mat1x2{
		a[0] * b[0],
		a[0] * b[1],
	}
}

func (a Mat1x1) MultiplyMat1x3(b Mat1x3) Mat1x3 {
	return Mat1x3{
		a[0] * b[0],
		a[0] * b[1],
		a[0] * b[2],
	}
}

func (a Mat1x1) MultiplyMat1x4(b Mat1x4) Mat1x4 {
	return Mat1x4{
		a[0] * b[0],
		a[0] * b[1],
		a[0] * b[2],
		a[0] * b[3],
	}
}

func (m Mat1x1) MultiplyVec1(v Vec1) float64 {
	return m[0] * v[0]
}

func (m Mat1x1) String() string {
	return ""
}

func (a Mat1x1) Subtract(b Mat1x1) Mat1x1 {
	return Mat1x1{a[0] - b[0]}
}

func (m Mat1x1) SubtractScalar(f float64) Mat1x1 {
	return Mat1x1{
		m[0] - f,
	}
}

func (m Mat1x1) Trace() float64 {
	return m[0]
}

func (m Mat1x1) Transpose() Mat1x1 {
	return Mat1x1{m[0]}
}

type Mat1x2 [2]float64

func (a Mat1x2) Add(b Mat1x2) Mat1x2 {
	return Mat1x2{
		a[0] + b[0],
		a[1] + b[1],
	}
}

func (m Mat1x2) AddScalar(f float64) Mat1x2 {
	return Mat1x2{
		m[0] + f,
		m[1] + f,
	}
}

func (m Mat1x2) Concatenate(f float64) Mat1x3 {
	return Mat1x3{m[0], m[1], f}
}

func (m Mat1x2) Homogenize() Mat1x3 {
	return Mat1x3{m[0], m[1], 0}
}

func (m Mat1x2) Len() int {
	return 2
}

func (m Mat1x2) Mat1x1() Mat1x1 {
	return Mat1x1{m[0]}
}

func (m Mat1x2) MultiplyScalar(f float64) Mat1x2 {
	return Mat1x2{m[0] * f, m[1] * f}
}

func (a Mat1x2) MultiplyMat2x1(b Mat2x1) Mat1x1 {
	return Mat1x1{a[0]*b[0][0] + a[1]*b[1][0]}
}

func (a Mat1x2) MultiplyMat2x2(b Mat2x2) Mat1x2 {
	return Mat1x2{
		a[0]*b[0][0] + a[1]*b[1][0],
		a[0]*b[0][1] + a[1]*b[1][1],
	}
}

func (a Mat1x2) MultiplyMat2x3(b Mat2x3) Mat1x3 {
	return Mat1x3{
		a[0]*b[0][0] + a[1]*b[1][0],
		a[0]*b[0][1] + a[1]*b[1][1],
		a[0]*b[0][2] + a[1]*b[1][2],
	}
}

func (a Mat1x2) MultiplyMat2x4(b Mat2x4) Mat1x4 {
	return Mat1x4{
		a[0]*b[0][0] + a[1]*b[1][0],
		a[0]*b[0][1] + a[1]*b[1][1],
		a[0]*b[0][2] + a[1]*b[1][2],
		a[0]*b[0][3] + a[1]*b[1][3],
	}
}

func (m Mat1x2) MultiplyVec2(v Vec2) float64 {
	return m[0]*v[0] + m[1]*v[1]
}

func (a Mat1x2) Subtract(b Mat1x2) Mat1x2 {
	return Mat1x2{
		a[0] - b[0],
		a[1] - b[1],
	}
}

func (m Mat1x2) SubtractScalar(f float64) Mat1x2 {
	return Mat1x2{
		m[0] - f,
		m[1] - f,
	}
}

func (m Mat1x2) Transpose() Mat2x1 {
	return Mat2x1{
		Mat1x1{m[0]},
		Mat1x1{m[1]},
	}
}

type Mat1x3 [3]float64

func (a Mat1x3) Add(b Mat1x3) Mat1x3 {
	return Mat1x3{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

func (m Mat1x3) AddScalar(f float64) Mat1x3 {
	return Mat1x3{
		m[0] + f,
		m[1] + f,
		m[2] + f,
	}
}

func (m Mat1x3) Concatenate(f float64) Mat1x4 {
	return Mat1x4{m[0], m[1], m[2], f}
}

func (m Mat1x3) Homogenize() Mat1x4 {
	return Mat1x4{m[0], m[1], m[2], 0}
}

func (m Mat1x3) Len() int {
	return 3
}

func (m Mat1x3) Mat1x1() Mat1x1 {
	return Mat1x1{m[0]}
}

func (m Mat1x3) Mat1x2() Mat1x2 {
	return Mat1x2{m[0], m[1]}
}

func (m Mat1x3) MultiplyScalar(f float64) Mat1x3 {
	return Mat1x3{m[0] * f, m[1] * f, m[2] * f}
}

func (a Mat1x3) MultiplyMat3x1(b Mat3x1) Mat1x1 {
	return Mat1x1{a[0]*b[0][0] + a[1]*b[1][0] + a[2]*b[2][0]}
}

func (a Mat1x3) MultiplyMat3x2(b Mat3x2) Mat1x2 {
	return Mat1x2{
		a[0]*b[0][0] + a[1]*b[1][0] + a[2]*b[2][0],
		a[0]*b[0][1] + a[1]*b[1][1] + a[2]*b[2][1],
	}
}

func (a Mat1x3) MultiplyMat3x3(b Mat3x3) Mat1x3 {
	return Mat1x3{
		a[0]*b[0][0] + a[1]*b[1][0] + a[2]*b[2][0],
		a[0]*b[0][1] + a[1]*b[1][1] + a[2]*b[2][1],
		a[0]*b[0][2] + a[1]*b[1][2] + a[2]*b[2][2],
	}
}

func (a Mat1x3) MultiplyMat3x4(b Mat3x4) Mat1x4 {
	return Mat1x4{
		a[0]*b[0][0] + a[1]*b[1][0] + a[2]*b[2][0],
		a[0]*b[0][1] + a[1]*b[1][1] + a[2]*b[2][1],
		a[0]*b[0][2] + a[1]*b[1][2] + a[2]*b[2][2],
		a[0]*b[0][3] + a[1]*b[1][3] + a[2]*b[2][3],
	}
}

func (m Mat1x3) MultiplyVec3(v Vec3) float64 {
	return m[0]*v[0] + m[1]*v[1] + m[2]*v[2]
}

func (a Mat1x3) Subtract(b Mat1x3) Mat1x3 {
	return Mat1x3{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
}

func (m Mat1x3) SubtractScalar(f float64) Mat1x3 {
	return Mat1x3{
		m[0] - f,
		m[1] - f,
		m[2] - f,
	}
}

func (m Mat1x3) Transpose() Mat3x1 {
	return Mat3x1{
		Mat1x1{m[0]},
		Mat1x1{m[1]},
		Mat1x1{m[2]},
	}
}

type Mat1x4 [4]float64

func (a Mat1x4) Add(b Mat1x4) Mat1x4 {
	return Mat1x4{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
		a[3] + b[3],
	}
}

func (m Mat1x4) AddScalar(f float64) Mat1x4 {
	return Mat1x4{
		m[0] + f,
		m[1] + f,
		m[2] + f,
		m[3] + f,
	}
}

func (m Mat1x4) Concatenate(f float64) Mat1xN {
	return Mat1xN{m[0], m[1], m[2], m[3], f}
}

func (m Mat1x4) Homogenize() Mat1xN {
	return Mat1xN{m[0], m[1], m[2], m[3], 0}
}

func (m Mat1x4) Len() int {
	return 4
}

func (m Mat1x4) Mat1x1() Mat1x1 {
	return Mat1x1{m[0]}
}

func (m Mat1x4) Mat1x2() Mat1x2 {
	return Mat1x2{m[0], m[1]}
}

func (m Mat1x4) Mat1x3() Mat1x3 {
	return Mat1x3{m[0], m[1], m[2]}
}

func (m Mat1x4) MultiplyScalar(f float64) Mat1x4 {
	return Mat1x4{m[0] * f, m[1] * f, m[2] * f, m[3] * f}
}

func (a Mat1x4) MultiplyMat4x1(b Mat4x1) Mat1x1 {
	return Mat1x1{a[0]*b[0][0] + a[1]*b[1][0] + a[2]*b[2][0] + a[3]*b[3][0]}
}

func (a Mat1x4) MultiplyMat4x2(b Mat4x2) Mat1x2 {
	return Mat1x2{
		a[0]*b[0][0] + a[1]*b[1][0] + a[2]*b[2][0] + a[3]*b[3][0],
		a[0]*b[0][1] + a[1]*b[1][1] + a[2]*b[2][1] + a[3]*b[3][1],
	}
}

func (a Mat1x4) MultiplyMat4x3(b Mat4x3) Mat1x3 {
	return Mat1x3{
		a[0]*b[0][0] + a[1]*b[1][0] + a[2]*b[2][0] + a[3]*b[3][0],
		a[0]*b[0][1] + a[1]*b[1][1] + a[2]*b[2][1] + a[3]*b[3][1],
		a[0]*b[0][2] + a[1]*b[1][2] + a[2]*b[2][2] + a[3]*b[3][2],
	}
}

func (a Mat1x4) MultiplyMat4x4(b Mat4x4) Mat1x4 {
	return Mat1x4{
		a[0]*b[0][0] + a[1]*b[1][0] + a[2]*b[2][0] + a[3]*b[3][0],
		a[0]*b[0][1] + a[1]*b[1][1] + a[2]*b[2][1] + a[3]*b[3][1],
		a[0]*b[0][2] + a[1]*b[1][2] + a[2]*b[2][2] + a[3]*b[3][2],
		a[0]*b[0][3] + a[1]*b[1][3] + a[2]*b[2][3] + a[3]*b[3][3],
	}
}

func (m Mat1x4) MultiplyVec4(v Vec4) float64 {
	return m[0]*v[0] + m[1]*v[1] + m[2]*v[2] + m[3]*v[3]
}

func (a Mat1x4) Subtract(b Mat1x4) Mat1x4 {
	return Mat1x4{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
		a[3] - b[3],
	}
}

func (m Mat1x4) SubtractScalar(f float64) Mat1x4 {
	return Mat1x4{
		m[0] - f,
		m[1] - f,
		m[2] - f,
		m[3] - f,
	}
}

func (m Mat1x4) Transpose() Mat4x1 {
	return Mat4x1{
		Mat1x1{m[0]},
		Mat1x1{m[1]},
		Mat1x1{m[2]},
		Mat1x1{m[3]},
	}
}

type Mat1xN []float64

func (m Mat1xN) Concatenate(f float64) Mat1xN {
	return append(m, f)
}
