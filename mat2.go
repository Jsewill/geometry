/* */
package geometry

import (
	"math"
)

type Mat2x1 [2]Mat1x1

func (a Mat2x1) Add(b Mat2x1) Mat2x1 {
	return Mat2x1{
		Mat1x1{a[0][0] + b[0][0]},
		Mat1x1{a[1][0] + b[1][0]},
	}
}

func (m Mat2x1) AddScalar(f float64) Mat2x1 {
	return Mat2x1{
		Mat1x1{m[0][0] + f},
		Mat1x1{m[1][0] + f},
	}
}

func (m Mat2x1) Len() int {
	return 2
}

func (m Mat2x1) MultiplyScalar(f float64) Mat2x1 {
	return Mat2x1{
		Mat1x1{m[0][0] * f},
		Mat1x1{m[1][0] * f},
	}
}

func (a Mat2x1) MultiplyMat2x1(b Mat1x1) Mat2x1 {
	return Mat2x1{
		Mat1x1{a[0][0] * b[0]},
		Mat1x1{a[1][0] * b[0]},
	}
}

func (a Mat2x1) MultiplyMat2x2(b Mat1x2) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			a[0][0] * b[0],
			a[0][0] * b[1],
		},
		Mat1x2{
			a[1][0] * b[0],
			a[1][0] * b[1],
		},
	}
}

func (a Mat2x1) MultiplyMat2x3(b Mat1x3) Mat2x3 {
	return Mat2x3{
		Mat1x3{
			a[0][0] * b[0],
			a[0][0] * b[1],
			a[0][0] * b[2],
		},
		Mat1x3{
			a[1][0] * b[0],
			a[1][0] * b[1],
			a[1][0] * b[2],
		},
	}
}

func (a Mat2x1) MultiplyMat2x4(b Mat1x4) Mat2x4 {
	return Mat2x4{
		Mat1x4{
			a[0][0] * b[0],
			a[0][0] * b[1],
			a[0][0] * b[2],
			a[0][0] * b[3],
		},
		Mat1x4{
			a[1][0] * b[0],
			a[1][0] * b[1],
			a[1][0] * b[2],
			a[1][0] * b[3],
		},
	}
}

func (m Mat2x1) MultiplyVec1(v Vec1) Vec2 {
	return Vec2{
		m[0][0] * v[0],
		m[1][0] * v[0],
	}
}

func (a Mat2x1) Subtract(b Mat2x1) Mat2x1 {
	return Mat2x1{
		Mat1x1{a[0][0] - b[0][0]},
		Mat1x1{a[1][0] - b[1][0]},
	}
}

func (m Mat2x1) SubtractScalar(f float64) Mat2x1 {
	return Mat2x1{
		Mat1x1{m[0][0] - f},
		Mat1x1{m[1][0] - f},
	}
}

func (m Mat2x1) Transpose() Mat1x2 {
	return Mat1x2{m[0][0], m[1][0]}
}

type Mat2x2 [2]Mat1x2

func (a Mat2x2) Add(b Mat2x2) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			a[0][0] + b[0][0],
			a[0][1] + b[0][1],
		},
		Mat1x2{
			a[1][0] + b[1][0],
			a[1][1] + b[1][1],
		},
	}
}

func (m Mat2x2) AddScalar(f float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0] + f,
			m[0][1] + f,
		},
		Mat1x2{
			m[1][0] + f,
			m[1][1] + f,
		},
	}
}

func (m Mat2x2) Determinant() float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

func (m Mat2x2) Homogenize() Mat3x3 {
	return Mat3x3{
		Mat1x3{m[0][0], m[0][1], 0},
		Mat1x3{m[1][0], m[1][1], 0},
		Mat1x3{0, 0, 1},
	}
}

func (m Mat2x2) Inverse() Mat2x2 {
	//Determinant
	d := m[0][0]*m[1][1] - m[0][1]*m[1][0]

	//Invertible?
	if d == 0 {
		return Mat2x2{}
	}

	//Return the inverse matrix
	return Mat2x2{
		Mat1x2{
			d * m[1][1],
			d * -m[0][1],
		},
		Mat1x2{
			d * -m[1][0],
			d * m[0][0],
		},
	}
}

func (m Mat2x2) Len() int {
	return 4
}

func (m Mat2x2) LinearReflectX() Mat2x2 {
	return Mat2x2{
		Mat1x2{-m[0][0], m[0][1]},
		Mat1x2{-m[1][0], m[1][1]},
	}
}

func (m Mat2x2) LinearReflectXY() Mat2x2 {
	return Mat2x2{
		Mat1x2{-m[0][0], -m[0][1]},
		Mat1x2{-m[1][0], -m[1][1]},
	}
}

func (m Mat2x2) LinearReflectY() Mat2x2 {
	return Mat2x2{
		Mat1x2{m[0][0], -m[0][1]},
		Mat1x2{m[1][0], -m[1][1]},
	}
}

//Rotate in the XY Plane
func (m Mat2x2) LinearRotateXY(Θ float64) Mat2x2 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat2x2{
		Mat1x2{
			m[0][0]*c + m[0][1]*s,
			m[0][0]*-s + m[0][1]*c,
		},
		Mat1x2{
			m[1][0]*c + m[1][1]*s,
			m[1][0]*-s + m[1][1]*c,
		},
	}
}

func (m Mat2x2) LinearShearX(τ float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0],
			m[0][0]*τ + m[0][1],
		},
		Mat1x2{
			m[1][0],
			m[1][0]*τ + m[1][1],
		},
	}
}

func (m Mat2x2) LinearShearXY(τx, τy float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0] + m[0][1]*τy,
			m[0][0]*τx + m[0][1],
		},
		Mat1x2{
			m[1][0] + m[1][1]*τy,
			m[1][0]*τx + m[1][1],
		},
	}
}

func (m Mat2x2) LinearShearY(τ float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0] + m[0][1]*τ,
			m[0][1],
		},
		Mat1x2{
			m[1][0] + m[1][1]*τ,
			m[1][1],
		},
	}
}

func (m Mat2x2) LinearTranslateX(δ float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0],
			m[0][0]*δ + m[0][1],
		},
		Mat1x2{
			m[1][0],
			m[1][0]*δ + m[1][1],
		},
	}
}

func (m Mat2x2) LinearTranslateXY(δ Vec2) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0],
			m[0][0]*δ[0] + m[0][1]*δ[1],
		},
		Mat1x2{
			m[1][0],
			m[1][0]*δ[0] + m[1][1]*δ[1],
		},
	}
}

func (m Mat2x2) LinearTranslateY(δ float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0],
			m[0][1] * δ,
		},
		Mat1x2{
			m[1][0],
			m[1][1] * δ,
		},
	}
}

func (m Mat2x2) MultiplyScalar(f float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{m[0][0] * f, m[0][1] * f},
		Mat1x2{m[1][0] * f, m[1][1] * f},
	}
}

func (a Mat2x2) MultiplyMat2x1(b Mat2x1) Mat2x1 {
	return Mat2x1{
		Mat1x1{a[0][0]*b[0][0] + a[0][1]*b[1][0]},
		Mat1x1{a[1][0]*b[0][0] + a[1][1]*b[1][0]},
	}
}

func (a Mat2x2) MultiplyMat2x2(b Mat2x2) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			a[0][0]*b[0][0] + a[0][1]*b[1][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1],
		},
		Mat1x2{
			a[1][0]*b[0][0] + a[1][1]*b[1][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1],
		},
	}
}

func (a Mat2x2) MultiplyMat2x3(b Mat2x3) Mat2x3 {
	return Mat2x3{
		Mat1x3{
			a[0][0]*b[0][0] + a[0][1]*b[1][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1],
			a[0][0]*b[0][2] + a[0][1]*b[1][2],
		},
		Mat1x3{
			a[1][0]*b[0][0] + a[1][1]*b[1][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1],
			a[1][0]*b[0][2] + a[1][1]*b[1][2],
		},
	}
}

func (a Mat2x2) MultiplyMat2x4(b Mat2x4) Mat2x4 {
	return Mat2x4{
		Mat1x4{
			a[0][0]*b[0][0] + a[0][1]*b[1][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1],
			a[0][0]*b[0][2] + a[0][1]*b[1][2],
			a[0][0]*b[0][3] + a[0][1]*b[1][3],
		},
		Mat1x4{
			a[1][0]*b[0][0] + a[1][1]*b[1][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1],
			a[1][0]*b[0][2] + a[1][1]*b[1][2],
			a[1][0]*b[0][3] + a[1][1]*b[1][3],
		},
	}
}

func (m Mat2x2) MultiplyVec2(v Vec2) Vec2 {
	return Vec2{
		m[0][0]*v[0] + m[0][1]*v[1],
		m[1][0]*v[0] + m[1][1]*v[1],
	}
}

func (m Mat2x2) ReflectX() Mat2x2 {
	return Mat2x2{
		Mat1x2{-m[0][0], m[0][1]},
		Mat1x2{-m[1][0], m[1][1]},
	}
}

func (m Mat2x2) Scale(Δ float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0] * Δ,
			m[0][1],
		},
		Mat1x2{
			m[1][0] * Δ,
			m[1][1],
		},
	}
}

func (m Mat2x2) ScaleX(Δ float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0] * Δ,
			m[0][1],
		},
		Mat1x2{
			m[1][0] * Δ,
			m[1][1],
		},
	}
}

func (m Mat2x2) ShearY(τ float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0] + m[0][1]*τ,
			m[0][1],
		},
		Mat1x2{
			m[1][0] + m[1][1]*τ,
			m[1][1],
		},
	}
}

func (a Mat2x2) Subtract(b Mat2x2) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			a[0][0] - b[0][0],
			a[0][1] - b[0][1],
		},
		Mat1x2{
			a[1][0] - b[1][0],
			a[1][1] - b[1][1],
		},
	}
}

func (m Mat2x2) SubtractScalar(f float64) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0] - f,
			m[0][1] - f,
		},
		Mat1x2{
			m[1][0] - f,
			m[1][1] - f,
		},
	}
}

func (m Mat2x2) Trace() float64 {
	return m[0][0] + m[1][1]
}

func (m Mat2x2) TranslateX(δ Vec1) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			m[0][0],
			m[0][0]*δ[0] + m[0][1],
		},
		Mat1x2{
			m[1][0],
			m[1][0]*δ[0] + m[1][1],
		},
	}
}

func (m Mat2x2) Transpose() Mat2x2 {
	return Mat2x2{
		Mat1x2{m[0][0], m[1][0]},
		Mat1x2{m[0][1], m[1][1]},
	}
}

type Mat2x3 [2]Mat1x3

func (a Mat2x3) Add(b Mat2x3) Mat2x3 {
	return Mat2x3{
		Mat1x3{
			a[0][0] + b[0][0],
			a[0][1] + b[0][1],
			a[0][2] + b[0][2],
		},
		Mat1x3{
			a[1][0] + b[1][0],
			a[1][1] + b[1][1],
			a[1][2] + b[1][2],
		},
	}
}

func (m Mat2x3) AddScalar(f float64) Mat2x3 {
	return Mat2x3{
		Mat1x3{
			m[0][0] + f,
			m[0][1] + f,
			m[0][2] + f,
		},
		Mat1x3{
			m[1][0] + f,
			m[1][1] + f,
			m[1][2] + f,
		},
	}
}

func (m Mat2x3) Len() int {
	return 6
}

func (m Mat2x3) MultiplyScalar(f float64) Mat2x3 {
	return Mat2x3{
		Mat1x3{m[0][0] * f, m[0][1] * f, m[0][2] * f},
		Mat1x3{m[1][0] * f, m[1][1] * f, m[1][2] * f},
	}
}

//Might as well be a Vec2
func (a Mat2x3) MultiplyMat3x1(b Mat3x1) Mat2x1 {
	return Mat2x1{
		Mat1x1{a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0]},
		Mat1x1{a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0]},
	}
}

func (a Mat2x3) MultiplyMat3x2(b Mat3x2) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1],
		},
		Mat1x2{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1],
		},
	}
}

func (a Mat2x3) MultiplyMat3x3(b Mat3x3) Mat2x3 {
	return Mat2x3{
		Mat1x3{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1],
			a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2],
		},
		Mat1x3{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1],
			a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2],
		},
	}
}

func (a Mat2x3) MultiplyMat3x4(b Mat3x4) Mat2x4 {
	return Mat2x4{
		Mat1x4{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1],
			a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2],
			a[0][0]*b[0][3] + a[0][1]*b[1][3] + a[0][2]*b[2][3],
		},
		Mat1x4{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1],
			a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2],
			a[1][0]*b[0][3] + a[1][1]*b[1][3] + a[1][2]*b[2][3],
		},
	}
}

func (m Mat2x3) MultiplyVec3(v Vec3) Vec2 {
	return Vec2{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2],
	}
}

func (a Mat2x3) Subtract(b Mat2x3) Mat2x3 {
	return Mat2x3{
		Mat1x3{
			a[0][0] - b[0][0],
			a[0][1] - b[0][1],
			a[0][2] - b[0][2],
		},
		Mat1x3{
			a[1][0] - b[1][0],
			a[1][1] - b[1][1],
			a[1][2] - b[1][2],
		},
	}
}

func (m Mat2x3) SubtractScalar(f float64) Mat2x3 {
	return Mat2x3{
		Mat1x3{
			m[0][0] - f,
			m[0][1] - f,
			m[0][2] - f,
		},
		Mat1x3{
			m[1][0] - f,
			m[1][1] - f,
			m[1][2] - f,
		},
	}
}

func (m Mat2x3) Transpose() Mat3x2 {
	return Mat3x2{
		Mat1x2{m[0][0], m[1][0]},
		Mat1x2{m[0][1], m[1][1]},
		Mat1x2{m[0][2], m[1][2]},
	}
}

type Mat2x4 [2]Mat1x4

func (a Mat2x4) Add(b Mat2x4) Mat2x4 {
	return Mat2x4{
		Mat1x4{
			a[0][0] + b[0][0],
			a[0][1] + b[0][1],
			a[0][2] + b[0][2],
			a[0][3] + b[0][3],
		},
		Mat1x4{
			a[1][0] + b[1][0],
			a[1][1] + b[1][1],
			a[1][2] + b[1][2],
			a[1][3] + b[1][3],
		},
	}
}

func (m Mat2x4) AddScalar(f float64) Mat2x4 {
	return Mat2x4{
		Mat1x4{
			m[0][0] + f,
			m[0][1] + f,
			m[0][2] + f,
			m[0][3] + f,
		},
		Mat1x4{
			m[1][0] + f,
			m[1][1] + f,
			m[1][2] + f,
			m[1][3] + f,
		},
	}
}

func (m Mat2x4) Len() int {
	return 4
}

func (m Mat2x4) MultiplyScalar(f float64) Mat2x4 {
	return Mat2x4{
		Mat1x4{m[0][0] * f, m[0][1] * f, m[0][2] * f, m[0][3] * f},
		Mat1x4{m[1][0] * f, m[1][1] * f, m[1][2] * f, m[1][3] * f},
	}
}

func (a Mat2x4) MultiplyMat4x1(b Mat4x1) Mat2x1 {
	return Mat2x1{
		Mat1x1{a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0]},
		Mat1x1{a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0]},
	}
}

func (a Mat2x4) MultiplyMat4x2(b Mat4x2) Mat2x2 {
	return Mat2x2{
		Mat1x2{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1] + a[0][3]*b[3][1],
		},
		Mat1x2{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1] + a[1][3]*b[3][1],
		},
	}
}

func (a Mat2x4) MultiplyMat4x3(b Mat4x3) Mat2x3 {
	return Mat2x3{
		Mat1x3{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1] + a[0][3]*b[3][1],
			a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2] + a[0][3]*b[3][2],
		},
		Mat1x3{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1] + a[1][3]*b[3][1],
			a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2] + a[1][3]*b[3][2],
		},
	}
}

func (a Mat2x4) MultiplyMat4x4(b Mat4x4) Mat2x4 {
	return Mat2x4{
		Mat1x4{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1] + a[0][3]*b[3][1],
			a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2] + a[0][3]*b[3][2],
			a[0][0]*b[0][3] + a[0][1]*b[1][3] + a[0][2]*b[2][3] + a[0][3]*b[3][3],
		},
		Mat1x4{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1] + a[1][3]*b[3][1],
			a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2] + a[1][3]*b[3][2],
			a[1][0]*b[0][3] + a[1][1]*b[1][3] + a[1][2]*b[2][3] + a[1][3]*b[3][3],
		},
	}
}

func (m Mat2x4) MultiplyVec4(v Vec4) Vec2 {
	return Vec2{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2] + m[0][3]*v[3],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2] + m[1][3]*v[3],
	}
}

func (a Mat2x4) Subtract(b Mat2x4) Mat2x4 {
	return Mat2x4{
		Mat1x4{
			a[0][0] - b[0][0],
			a[0][1] - b[0][1],
			a[0][2] - b[0][2],
			a[0][3] - b[0][3],
		},
		Mat1x4{
			a[1][0] - b[1][0],
			a[1][1] - b[1][1],
			a[1][2] - b[1][2],
			a[1][3] - b[1][3],
		},
	}
}

func (m Mat2x4) SubtractScalar(f float64) Mat2x4 {
	return Mat2x4{
		Mat1x4{
			m[0][0] - f,
			m[0][1] - f,
			m[0][2] - f,
			m[0][3] - f,
		},
		Mat1x4{
			m[1][0] - f,
			m[1][1] - f,
			m[1][2] - f,
			m[1][3] - f,
		},
	}
}

func (m Mat2x4) Transpose() Mat4x2 {
	return Mat4x2{
		Mat1x2{m[0][0], m[1][0]},
		Mat1x2{m[0][1], m[1][1]},
		Mat1x2{m[0][2], m[1][2]},
		Mat1x2{m[0][3], m[1][3]},
	}
}

type Mat2xN []float64
