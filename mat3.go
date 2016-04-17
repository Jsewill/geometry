/* */
package geometry

import (
	"math"
)

type Mat3x1 [3]Mat1x1

func (a Mat3x1) Add(b Mat3x1) Mat3x1 {
	return Mat3x1{
		Mat1x1{a[0][0] + b[0][0]},
		Mat1x1{a[1][0] + b[1][0]},
		Mat1x1{a[2][0] + b[2][0]},
	}
}

func (m Mat3x1) AddScalar(f float64) Mat3x1 {
	return Mat3x1{
		Mat1x1{m[0][0] + f},
		Mat1x1{m[1][0] + f},
		Mat1x1{m[2][0] + f},
	}
}

func (m Mat3x1) Len() int {
	return 3
}

func (m Mat3x1) MultiplyScalar(f float64) Mat3x1 {
	return Mat3x1{
		Mat1x1{m[0][0] * f},
		Mat1x1{m[1][0] * f},
		Mat1x1{m[2][0] * f},
	}
}

func (a Mat3x1) MultiplyMat1x1(b Mat1x1) Mat3x1 {
	return Mat3x1{
		Mat1x1{a[0][0] * b[0]},
		Mat1x1{a[1][0] * b[0]},
		Mat1x1{a[2][0] * b[0]},
	}
}

func (a Mat3x1) MultiplyMat1x2(b Mat1x2) Mat3x2 {
	return Mat3x2{
		Mat1x2{
			a[0][0] * b[0],
			a[0][0] * b[1],
		},
		Mat1x2{
			a[1][0] * b[0],
			a[1][0] * b[1],
		},
		Mat1x2{
			a[2][0] * b[0],
			a[2][0] * b[1],
		},
	}
}

func (a Mat3x1) MultiplyMat1x3(b Mat1x3) Mat3x3 {
	return Mat3x3{
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
		Mat1x3{
			a[2][0] * b[0],
			a[2][0] * b[1],
			a[2][0] * b[2],
		},
	}
}

func (a Mat3x1) MultiplyMat1x4(b Mat1x4) Mat3x4 {
	return Mat3x4{
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
		Mat1x4{
			a[2][0] * b[0],
			a[2][0] * b[1],
			a[2][0] * b[2],
			a[2][0] * b[3],
		},
	}
}

func (m Mat3x1) MultiplyVec1(v Vec1) Vec3 {
	return Vec3{
		m[0][0] * v[0],
		m[1][0] * v[0],
		m[2][0] * v[0],
	}
}

func (a Mat3x1) Subtract(b Mat3x1) Mat3x1 {
	return Mat3x1{
		Mat1x1{a[0][0] - b[0][0]},
		Mat1x1{a[1][0] - b[1][0]},
		Mat1x1{a[2][0] - b[2][0]},
	}
}

func (m Mat3x1) SubtractScalar(f float64) Mat3x1 {
	return Mat3x1{
		Mat1x1{m[0][0] - f},
		Mat1x1{m[1][0] - f},
		Mat1x1{m[2][0] - f},
	}
}

func (m Mat3x1) Transpose() Mat1x3 {
	return Mat1x3{m[0][0], m[1][0], m[2][0]}
}

type Mat3x2 [3]Mat1x2

func (a Mat3x2) Add(b Mat3x2) Mat3x2 {
	return Mat3x2{
		Mat1x2{
			a[0][0] + b[0][0],
			a[0][1] + b[0][1],
		},
		Mat1x2{
			a[1][0] + b[1][0],
			a[1][1] + b[1][1],
		},
		Mat1x2{
			a[2][0] + b[2][0],
			a[2][1] + b[2][1],
		},
	}
}

func (m Mat3x2) AddScalar(f float64) Mat3x2 {
	return Mat3x2{
		Mat1x2{
			m[0][0] + f,
			m[0][1] + f,
		},
		Mat1x2{
			m[1][0] + f,
			m[1][1] + f,
		},
		Mat1x2{
			m[2][0] + f,
			m[2][1] + f,
		},
	}
}

func (m Mat3x2) Len() int {
	return 6
}

func (m Mat3x2) MultiplyScalar(f float64) Mat3x2 {
	return Mat3x2{
		Mat1x2{m[0][0] * f, m[0][1] * f},
		Mat1x2{m[1][0] * f, m[1][1] * f},
		Mat1x2{m[2][0] * f, m[2][1] * f},
	}
}

func (a Mat3x2) MultiplyMat2x1(b Mat2x1) Mat3x1 {
	return Mat3x1{
		Mat1x1{a[0][0]*b[0][0] + a[0][1]*b[1][0]},
		Mat1x1{a[1][0]*b[0][0] + a[1][1]*b[1][0]},
		Mat1x1{a[2][0]*b[0][0] + a[2][1]*b[1][0]},
	}
}

func (a Mat3x2) MultiplyMat2x2(b Mat2x2) Mat3x2 {
	return Mat3x2{
		Mat1x2{
			a[0][0]*b[0][0] + a[0][1]*b[1][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1],
		},
		Mat1x2{
			a[1][0]*b[0][0] + a[1][1]*b[1][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1],
		},
		Mat1x2{
			a[2][0]*b[0][0] + a[2][1]*b[1][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1],
		},
	}
}

func (a Mat3x2) MultiplyMat2x3(b Mat2x3) Mat3x3 {
	return Mat3x3{
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
		Mat1x3{
			a[2][0]*b[0][0] + a[2][1]*b[1][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1],
			a[2][0]*b[0][2] + a[2][1]*b[1][2],
		},
	}
}

func (a Mat3x2) MultiplyMat2x4(b Mat2x4) Mat3x4 {
	return Mat3x4{
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
		Mat1x4{
			a[2][0]*b[0][0] + a[2][1]*b[1][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1],
			a[2][0]*b[0][2] + a[2][1]*b[1][2],
			a[2][0]*b[0][3] + a[2][1]*b[1][3],
		},
	}
}

func (m Mat3x2) MultiplyVec2(v Vec2) Vec3 {
	return Vec3{
		m[0][0]*v[0] + m[0][1]*v[1],
		m[1][0]*v[0] + m[1][1]*v[1],
		m[2][0]*v[0] + m[2][1]*v[1],
	}
}

func (a Mat3x2) Subtract(b Mat3x2) Mat3x2 {
	return Mat3x2{
		Mat1x2{
			a[0][0] - b[0][0],
			a[0][1] - b[0][1],
		},
		Mat1x2{
			a[1][0] - b[1][0],
			a[1][1] - b[1][1],
		},
		Mat1x2{
			a[2][0] - b[2][0],
			a[2][1] - b[2][1],
		},
	}
}

func (m Mat3x2) SubtractScalar(f float64) Mat3x2 {
	return Mat3x2{
		Mat1x2{
			m[0][0] - f,
			m[0][1] - f,
		},
		Mat1x2{
			m[1][0] - f,
			m[1][1] - f,
		},
		Mat1x2{
			m[2][0] - f,
			m[2][1] - f,
		},
	}
}

func (m Mat3x2) Transpose() Mat2x3 {
	return Mat2x3{
		Mat1x3{m[0][0], m[1][0], m[2][0]},
		Mat1x3{m[0][1], m[1][1], m[2][1]},
	}
}

type Mat3x3 [3]Mat1x3

func (a Mat3x3) Add(b Mat3x3) Mat3x3 {
	return Mat3x3{
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
		Mat1x3{
			a[2][0] + b[2][0],
			a[2][1] + b[2][1],
			a[2][2] + b[2][2],
		},
	}
}

func (m Mat3x3) AddScalar(f float64) Mat3x3 {
	return Mat3x3{
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
		Mat1x3{
			m[2][0] + f,
			m[2][1] + f,
			m[2][2] + f,
		},
	}
}

func (m Mat3x3) Convolute(k Mat3x3) float64 {
	return m[0][0]*k[2][2] + m[0][1]*k[2][1] + m[0][2]*k[2][0] + m[1][0]*k[1][2] + m[1][1]*k[1][1] + m[1][2]*k[1][0] + m[2][0]*k[0][2] + m[2][1]*k[0][1] + m[2][2]*k[0][0]
}

func (m Mat3x3) Determinant() float64 {
	return m[0][0]*m[1][1]*m[2][2] + m[0][1]*m[1][2]*m[2][0] + m[0][2]*m[1][0]*m[2][1] - m[0][2]*m[1][1]*m[2][0] - m[0][1]*m[1][0]*m[2][2] - m[0][0]*m[1][2]*m[2][0]
}

func (m Mat3x3) Homogenize() Mat4x4 {
	return Mat4x4{
		Mat1x4{m[0][0], m[0][1], m[0][2], 0},
		Mat1x4{m[1][0], m[1][1], m[1][2], 0},
		Mat1x4{m[2][0], m[2][1], m[2][2], 0},
		Mat1x4{0, 0, 0, 1},
	}
}

func (m Mat3x3) Inverse() Mat3x3 {
	//Determinant
	d := m[0][0]*m[1][1]*m[2][2] + m[0][1]*m[1][2]*m[2][0] + m[0][2]*m[1][0]*m[2][1] - m[0][2]*m[1][1]*m[2][0] - m[0][1]*m[1][0]*m[2][2] - m[0][0]*m[1][2]*m[2][0]

	//Invertible?
	if d == 0 {
		return Mat3x3{}
	}

	d = 1.0 / d

	//Return the transpose of the matrix of the determinants of each sub matrix
	return Mat3x3{
		Mat1x3{
			d * (m[1][1]*m[2][2] - m[1][2]*m[2][1]),
			d * (m[0][2]*m[2][1] - m[0][1]*m[2][2]),
			d * (m[0][1]*m[1][2] - m[0][2]*m[1][1]),
		},
		Mat1x3{
			d * (m[1][2]*m[2][0] - m[1][0]*m[2][2]),
			d * (m[0][0]*m[2][2] - m[0][2]*m[2][0]),
			d * (m[0][2]*m[1][0] - m[0][0]*m[1][2]),
		},
		Mat1x3{
			d * (m[1][0]*m[2][1] - m[1][1]*m[2][0]),
			d * (m[0][1]*m[2][0] - m[0][0]*m[2][1]),
			d * (m[0][0]*m[1][1] - m[0][1]*m[1][0]),
		},
	}
}

func (m Mat3x3) Len() int {
	return 9
}

func (m Mat3x3) LinearReflectX() Mat3x3 {
	return Mat3x3{
		Mat1x3{-m[0][0], m[0][1], m[0][2]},
		Mat1x3{-m[1][0], m[1][1], m[1][2]},
		Mat1x3{-m[2][0], m[2][1], m[2][2]},
	}
}

func (m Mat3x3) LinearReflectXY() Mat3x3 {
	return Mat3x3{
		Mat1x3{-m[0][0], -m[0][1], m[0][2]},
		Mat1x3{-m[1][0], -m[1][1], m[1][2]},
		Mat1x3{-m[2][0], -m[2][1], m[2][2]},
	}
}

func (m Mat3x3) LinearReflectXYZ() Mat3x3 {
	return Mat3x3{
		Mat1x3{-m[0][0], -m[0][1], -m[0][2]},
		Mat1x3{-m[1][0], -m[1][1], -m[1][2]},
		Mat1x3{-m[2][0], -m[2][1], -m[2][2]},
	}
}

func (m Mat3x3) LinearReflectXZ() Mat3x3 {
	return Mat3x3{
		Mat1x3{-m[0][0], m[0][1], -m[0][2]},
		Mat1x3{-m[1][0], m[1][1], -m[1][2]},
		Mat1x3{-m[2][0], m[2][1], -m[2][2]},
	}
}

func (m Mat3x3) LinearReflectY() Mat3x3 {
	return Mat3x3{
		Mat1x3{m[0][0], -m[0][1], m[0][2]},
		Mat1x3{m[1][0], -m[1][1], m[1][2]},
		Mat1x3{m[2][0], -m[2][1], m[2][2]},
	}
}

func (m Mat3x3) LinearReflectYZ() Mat3x3 {
	return Mat3x3{
		Mat1x3{m[0][0], -m[0][1], -m[0][2]},
		Mat1x3{m[1][0], -m[1][1], -m[1][2]},
		Mat1x3{m[2][0], -m[2][1], -m[2][2]},
	}
}

func (m Mat3x3) LinearReflectZ() Mat3x3 {
	return Mat3x3{
		Mat1x3{m[0][0], m[0][1], -m[0][2]},
		Mat1x3{m[1][0], m[1][1], -m[1][2]},
		Mat1x3{m[2][0], m[2][1], -m[2][2]},
	}
}

//Takes up to three angles representing the subsequent rotation in each plane. Order is as follows: YZ, XZ, XY
func (m Mat3x3) LinearRotate(Θ Vec3) Mat3x3 {
	var planes int
	for i, θ := range Θ {
		if θ != 0 {
			planes |= 1 << uint(i)
		}
	}

	if planes == 0 {
		return Mat3x3{
			Mat1x3{m[0][0], m[0][1], m[0][2]},
			Mat1x3{m[1][0], m[1][1], m[1][2]},
			Mat1x3{m[2][0], m[2][1], m[2][2]},
		}
	}

	var c, s float64
	var a, b Mat3x3

	switch {
	//X-axis rotation
	//Rotate in the YZ Plane
	case planes&1 > 0:
		c, s = math.Cos(Θ[0]), math.Sin(Θ[0])
		b = Mat3x3{
			Mat1x3{
				m[0][0],
				m[0][1]*c + m[0][2]*s,
				m[0][1]*-s + m[0][2]*c,
			},
			Mat1x3{
				m[1][0],
				m[1][1]*c + m[1][2]*s,
				m[1][1]*-s + m[1][2]*c,
			},
			Mat1x3{
				m[2][0],
				m[2][1]*c + m[2][2]*s,
				m[2][1]*-s + m[2][2]*c,
			},
		}

		fallthrough
	//Y-axis rotation
	//Rotate in the XZ Plane
	case planes&2 > 0:
		c, s = math.Cos(Θ[1]), math.Sin(Θ[1])
		b = Mat3x3{
			Mat1x3{
				m[0][0]*c + m[0][2]*-s,
				m[0][1],
				m[0][0]*s + m[0][2]*c,
			},
			Mat1x3{
				m[1][0]*c + m[1][2]*-s,
				m[1][1],
				m[1][0]*s + m[1][2]*c,
			},
			Mat1x3{
				m[2][0]*c + m[2][2]*-s,
				m[2][1],
				m[2][0]*s + m[2][2]*c,
			},
		}

		if planes&1 == 0 {
			a = b
		} else {
			a = a.MultiplyMat3x3(b)
		}

		fallthrough
	//Z-axis rotation
	//Rotate in the XY Plane
	case planes&4 > 0:
		c, s = math.Cos(Θ[2]), math.Sin(Θ[2])
		b = Mat3x3{
			Mat1x3{
				m[0][0]*c + m[0][1]*s,
				m[0][0]*-s + m[0][1]*c,
				m[0][2],
			},
			Mat1x3{
				m[1][0]*c + m[1][1]*s,
				m[1][0]*-s + m[1][1]*c,
				m[1][2],
			},
			Mat1x3{
				m[2][0]*c + m[2][1]*s,
				m[2][0]*-s + m[2][1]*c,
				m[2][2],
			},
		}

		if planes&3 == 0 {
			a = b
		} else {
			a = a.MultiplyMat3x3(b)
		}
	}

	return a
}

//Rotate in the YZ Plane
func (m Mat3x3) LinearRotateX(Θ float64) Mat3x3 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1]*c + m[0][2]*s,
			m[0][1]*-s + m[0][2]*c,
		},
		Mat1x3{
			m[1][0],
			m[1][1]*c + m[1][2]*s,
			m[1][1]*-s + m[1][2]*c,
		},
		Mat1x3{
			m[2][0],
			m[2][1]*c + m[2][2]*s,
			m[2][1]*-s + m[2][2]*c,
		},
	}
}

func (m Mat3x3) LinearRotateXY(Θx, Θy float64) Mat3x3 {
	xc, xs, yc, ys := math.Cos(Θx), math.Sin(Θx), math.Cos(Θy), math.Sin(Θy)
	x := Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1]*xc + m[0][2]*xs,
			m[0][1]*-xs + m[0][2]*xc,
		},
		Mat1x3{
			m[1][0],
			m[1][1]*xc + m[1][2]*xs,
			m[1][1]*-xs + m[1][2]*xc,
		},
		Mat1x3{
			m[2][0],
			m[2][1]*xc + m[2][2]*xs,
			m[2][1]*-xs + m[2][2]*xc,
		},
	}

	return Mat3x3{
		Mat1x3{
			x[0][0]*yc + x[0][2]*-ys,
			x[0][1],
			x[0][0]*ys + x[0][2]*yc,
		},
		Mat1x3{
			x[1][0]*yc + x[1][2]*-ys,
			x[1][1],
			x[1][0]*ys + x[1][2]*yc,
		},
		Mat1x3{
			x[2][0]*yc + x[2][2]*-ys,
			x[2][1],
			x[2][0]*ys + x[2][2]*yc,
		},
	}
}

func (m Mat3x3) LinearRotateXZ(Θx, Θz float64) Mat3x3 {
	xc, xs, zc, zs := math.Cos(Θx), math.Sin(Θx), math.Cos(Θz), math.Sin(Θz)
	x := Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1]*xc + m[0][2]*xs,
			m[0][1]*-xs + m[0][2]*xc,
		},
		Mat1x3{
			m[1][0],
			m[1][1]*xc + m[1][2]*xs,
			m[1][1]*-xs + m[1][2]*xc,
		},
		Mat1x3{
			m[2][0],
			m[2][1]*xc + m[2][2]*xs,
			m[2][1]*-xs + m[2][2]*xc,
		},
	}

	return Mat3x3{
		Mat1x3{
			x[0][0]*zc + x[0][1]*zs,
			x[0][0]*-zs + x[0][1]*zc,
			x[0][2],
		},
		Mat1x3{
			x[1][0]*zc + x[1][1]*zs,
			x[1][0]*-zs + x[1][1]*zc,
			x[1][2],
		},
		Mat1x3{
			x[2][0]*zc + x[2][1]*zs,
			x[2][0]*-zs + x[2][1]*zc,
			x[2][2],
		},
	}
}

func (m Mat3x3) LinearRotateXYZ(Θx, Θy, Θz float64) Mat3x3 {
	xc, xs, yc, ys, zc, zs := math.Cos(Θx), math.Sin(Θx), math.Cos(Θy), math.Sin(Θy), math.Cos(Θz), math.Sin(Θz)
	x := Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1]*xc + m[0][2]*xs,
			m[0][1]*-xs + m[0][2]*xc,
		},
		Mat1x3{
			m[1][0],
			m[1][1]*xc + m[1][2]*xs,
			m[1][1]*-xs + m[1][2]*xc,
		},
		Mat1x3{
			m[2][0],
			m[2][1]*xc + m[2][2]*xs,
			m[2][1]*-xs + m[2][2]*xc,
		},
	}

	xy := Mat3x3{
		Mat1x3{
			x[0][0]*yc + x[0][2]*-ys,
			x[0][1],
			x[0][0]*ys + x[0][2]*yc,
		},
		Mat1x3{
			x[1][0]*yc + x[1][2]*-ys,
			x[1][1],
			x[1][0]*ys + x[1][2]*yc,
		},
		Mat1x3{
			x[2][0]*yc + x[2][2]*-ys,
			x[2][1],
			x[2][0]*ys + x[2][2]*yc,
		},
	}

	return Mat3x3{
		Mat1x3{
			xy[0][0]*zc + xy[0][1]*zs,
			xy[0][0]*-zs + xy[0][1]*zc,
			xy[0][2],
		},
		Mat1x3{
			xy[1][0]*zc + xy[1][1]*zs,
			xy[1][0]*-zs + xy[1][1]*zc,
			xy[1][2],
		},
		Mat1x3{
			xy[2][0]*zc + xy[2][1]*zs,
			xy[2][0]*-zs + xy[2][1]*zc,
			xy[2][2],
		},
	}
}

//Rotate in the XZ Plane
func (m Mat3x3) LinearRotateY(Θ float64) Mat3x3 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat3x3{
		Mat1x3{
			m[0][0]*c + m[0][2]*-s,
			m[0][1],
			m[0][0]*s + m[0][2]*c,
		},
		Mat1x3{
			m[1][0]*c + m[1][2]*-s,
			m[1][1],
			m[1][0]*s + m[1][2]*c,
		},
		Mat1x3{
			m[2][0]*c + m[2][2]*-s,
			m[2][1],
			m[2][0]*s + m[2][2]*c,
		},
	}
}

func (m Mat3x3) LinearRotateYZ(Θy, Θz float64) Mat3x3 {
	yc, ys, zc, zs := math.Cos(Θy), math.Sin(Θy), math.Cos(Θz), math.Sin(Θz)
	y := Mat3x3{
		Mat1x3{
			m[0][0]*yc + m[0][2]*-ys,
			m[0][1],
			m[0][0]*ys + m[0][2]*yc,
		},
		Mat1x3{
			m[1][0]*yc + m[1][2]*-ys,
			m[1][1],
			m[1][0]*ys + m[1][2]*yc,
		},
		Mat1x3{
			m[2][0]*yc + m[2][2]*-ys,
			m[2][1],
			m[2][0]*ys + m[2][2]*yc,
		},
	}

	return Mat3x3{
		Mat1x3{
			y[0][0]*zc + y[0][1]*zs,
			y[0][0]*-zs + y[0][1]*zc,
			y[0][2],
		},
		Mat1x3{
			y[1][0]*zc + y[1][1]*zs,
			y[1][0]*-zs + y[1][1]*zc,
			y[1][2],
		},
		Mat1x3{
			y[2][0]*zc + y[2][1]*zs,
			y[2][0]*-zs + y[2][1]*zc,
			y[2][2],
		},
	}
}

//Rotate in the XY Plane
func (m Mat3x3) LinearRotateZ(Θ float64) Mat3x3 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat3x3{
		Mat1x3{
			m[0][0]*c + m[0][1]*s,
			m[0][0]*-s + m[0][1]*c,
			m[0][2],
		},
		Mat1x3{
			m[1][0]*c + m[1][1]*s,
			m[1][0]*-s + m[1][1]*c,
			m[1][2],
		},
		Mat1x3{
			m[2][0]*c + m[2][1]*s,
			m[2][0]*-s + m[2][1]*c,
			m[2][2],
		},
	}
}

func (m Mat3x3) LinearShearX(τ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][0]*τ[0] + m[0][1],
			m[0][0]*τ[1] + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][0]*τ[0] + m[1][1],
			m[1][0]*τ[1] + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][0]*τ[0] + m[2][1],
			m[2][0]*τ[1] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearShearXY(τx, τy Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] + m[0][1]*τy[0],
			m[0][0]*τx[0] + m[0][1],
			m[0][0]*τx[1] + m[0][1]*τy[1] + m[0][2],
		},
		Mat1x3{
			m[1][0] + m[1][1]*τy[0],
			m[1][0]*τx[0] + m[1][1],
			m[1][0]*τx[1] + m[1][1]*τy[1] + m[1][2],
		},
		Mat1x3{
			m[2][0] + m[2][1]*τy[0],
			m[2][0]*τx[0] + m[2][1],
			m[2][0]*τx[1] + m[2][1]*τy[1] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearShearXYZ(τx, τy, τz Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] + m[0][1]*τy[0] + m[0][2]*τz[0],
			m[0][0]*τx[0] + m[0][1] + m[0][2]*τz[1],
			m[0][0]*τx[1] + m[0][1]*τy[1] + m[0][2],
		},
		Mat1x3{
			m[1][0] + m[1][1]*τy[0] + m[1][2]*τz[0],
			m[1][0]*τx[0] + m[1][1] + m[1][2]*τz[1],
			m[1][0]*τx[1] + m[1][1]*τy[1] + m[1][2],
		},
		Mat1x3{
			m[2][0] + m[2][1]*τy[0] + m[2][2]*τz[0],
			m[2][0]*τx[0] + m[2][1] + m[2][2]*τz[1],
			m[2][0]*τx[1] + m[2][1]*τy[1] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearShearXZ(τx, τz Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] + m[0][2]*τz[0],
			m[0][0]*τx[0] + m[0][1] + m[0][2]*τz[1],
			m[0][0]*τx[1] + m[0][2],
		},
		Mat1x3{
			m[1][0] + m[1][2]*τz[0],
			m[1][0]*τx[0] + m[1][1] + m[1][2]*τz[1],
			m[1][0]*τx[1] + m[1][2],
		},
		Mat1x3{
			m[2][0] + m[2][2]*τz[0],
			m[2][0]*τx[0] + m[2][1] + m[2][2]*τz[1],
			m[2][0]*τx[1] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearShearY(τ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] + m[0][1]*τ[0],
			m[0][1],
			m[0][1]*τ[1] + m[0][2],
		},
		Mat1x3{
			m[1][0] + m[1][1]*τ[0],
			m[1][1],
			m[1][1]*τ[1] + m[1][2],
		},
		Mat1x3{
			m[2][0] + m[2][1]*τ[0],
			m[2][1],
			m[2][1]*τ[1] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearShearYZ(τy, τz Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] + m[0][1]*τy[0] + m[0][2]*τz[0],
			m[0][1] + m[0][2]*τz[1],
			m[0][1]*τy[1] + m[0][2],
		},
		Mat1x3{
			m[1][0] + m[1][1]*τy[0] + m[1][2]*τz[0],
			m[1][1] + m[1][2]*τz[1],
			m[1][1]*τy[1] + m[1][2],
		},
		Mat1x3{
			m[2][0] + m[2][1]*τy[0] + m[2][2]*τz[0],
			m[2][1] + m[2][2]*τz[1],
			m[2][1]*τy[1] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearShearZ(τ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] + m[0][2]*τ[0],
			m[0][1] + m[0][2]*τ[1],
			m[0][2],
		},
		Mat1x3{
			m[1][0] + m[1][2]*τ[0],
			m[1][1] + m[1][2]*τ[1],
			m[1][2],
		},
		Mat1x3{
			m[2][0] + m[2][2]*τ[0],
			m[2][1] + m[2][2]*τ[1],
			m[2][2],
		},
	}
}

func (m Mat3x3) LinearTranslate(δ Vec3) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][0]*δ[0] + m[0][1]*δ[1] + m[0][2]*δ[2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][0]*δ[0] + m[1][1]*δ[1] + m[1][2]*δ[2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][0]*δ[0] + m[2][1]*δ[1] + m[2][2]*δ[2],
		},
	}
}

func (m Mat3x3) LinearTranslateX(δ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][0]*δ + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][0]*δ + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][0]*δ + m[2][2],
		},
	}
}

func (m Mat3x3) LinearTranslateXY(δ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][0]*δ[0] + m[0][1]*δ[1] + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][0]*δ[0] + m[1][1]*δ[1] + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][0]*δ[0] + m[2][1]*δ[1] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearTranslateXYZ(δ Vec3) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][0]*δ[0] + m[0][1]*δ[1] + m[0][2]*δ[2] + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][0]*δ[0] + m[1][1]*δ[1] + m[1][2]*δ[2] + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][0]*δ[0] + m[2][1]*δ[1] + m[2][2]*δ[2] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearTranslateXZ(δ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][0]*δ[0] + m[0][2]*δ[1] + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][0]*δ[0] + m[1][2]*δ[1] + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][0]*δ[0] + m[2][2]*δ[1] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearTranslateY(δ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][1]*δ + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][1]*δ + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][1]*δ + m[2][2],
		},
	}
}

func (m Mat3x3) LinearTranslateYZ(δ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][1]*δ[0] + m[0][2]*δ[1] + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][1]*δ[0] + m[1][2]*δ[1] + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][1]*δ[0] + m[2][2]*δ[1] + m[2][2],
		},
	}
}

func (m Mat3x3) LinearTranslateZ(δ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][2]*δ + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][2]*δ + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][2]*δ + m[2][2],
		},
	}
}

func (m Mat3x3) Mat3x4() Mat3x4 {
	return Mat3x4{
		Mat1x4{m[0][0], m[0][1], m[0][2], 0},
		Mat1x4{m[1][0], m[1][1], m[1][2], 0},
		Mat1x4{m[2][0], m[2][1], m[2][2], 0},
	}
}

func (m Mat3x3) MultiplyScalar(f float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{m[0][0] * f, m[0][1] * f, m[0][2] * f},
		Mat1x3{m[1][0] * f, m[1][1] * f, m[1][2] * f},
		Mat1x3{m[2][0] * f, m[2][1] * f, m[2][2] * f},
	}
}

func (a Mat3x3) MultiplyMat3x1(b Mat3x1) Mat3x1 {
	return Mat3x1{
		Mat1x1{a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0]},
		Mat1x1{a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0]},
		Mat1x1{a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0]},
	}
}

func (a Mat3x3) MultiplyMat3x2(b Mat3x2) Mat3x2 {
	return Mat3x2{
		Mat1x2{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1],
		},
		Mat1x2{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1],
		},
		Mat1x2{
			a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1],
		},
	}
}

func (a Mat3x3) MultiplyMat3x3(b Mat3x3) Mat3x3 {
	return Mat3x3{
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
		Mat1x3{
			a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1],
			a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2],
		},
	}
}

func (a Mat3x3) MultiplyMat3x4(b Mat3x4) Mat3x4 {
	return Mat3x4{
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
		Mat1x4{
			a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1],
			a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2],
			a[2][0]*b[0][3] + a[2][1]*b[1][3] + a[2][2]*b[2][3],
		},
	}
}

func (m Mat3x3) MultiplyVec3(v Vec3) Vec3 {
	return Vec3{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2],
		m[2][0]*v[0] + m[2][1]*v[1] + m[2][2]*v[2],
	}
}

func (m Mat3x3) ReflectX() Mat3x3 {
	return Mat3x3{
		Mat1x3{-m[0][0], m[0][1], m[0][2]},
		Mat1x3{-m[1][0], m[1][1], m[1][2]},
		Mat1x3{-m[2][0], m[2][1], m[2][2]},
	}
}

func (m Mat3x3) ReflectXY() Mat3x3 {
	return Mat3x3{
		Mat1x3{-m[0][0], -m[0][1], m[0][2]},
		Mat1x3{-m[1][0], -m[1][1], m[1][2]},
		Mat1x3{-m[2][0], -m[2][1], m[2][2]},
	}
}

func (m Mat3x3) ReflectY() Mat3x3 {
	return Mat3x3{
		Mat1x3{m[0][0], -m[0][1], m[0][2]},
		Mat1x3{m[1][0], -m[1][1], m[1][2]},
		Mat1x3{m[2][0], -m[2][1], m[2][2]},
	}
}

func (m Mat3x3) Rotate(Θ float64) Mat3x3 {
	c, s := math.Cos(Θ), math.Sin(Θ)
	return Mat3x3{
		Mat1x3{
			m[0][0]*c + m[0][1]*s,
			m[0][0]*-s + m[0][1]*c,
			m[0][2],
		},
		Mat1x3{
			m[1][0]*c + m[1][1]*s,
			m[1][0]*-s + m[1][1]*c,
			m[1][2],
		},
		Mat1x3{
			m[2][0]*c + m[2][1]*s,
			m[2][0]*-s + m[2][1]*c,
			m[2][2],
		},
	}
}

func (m Mat3x3) ScaleX(Δ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] * Δ,
			m[0][1],
			m[0][2],
		},
		Mat1x3{
			m[1][0] * Δ,
			m[1][1],
			m[1][2],
		},
		Mat1x3{
			m[2][0] * Δ,
			m[2][1],
			m[2][2],
		},
	}
}

func (m Mat3x3) ScaleXY(Δ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] * Δ[0],
			m[0][1] * Δ[1],
			m[0][2],
		},
		Mat1x3{
			m[1][0] * Δ[0],
			m[1][1] * Δ[1],
			m[1][2],
		},
		Mat1x3{
			m[2][0] * Δ[0],
			m[2][1] * Δ[1],
			m[2][2],
		},
	}
}

func (m Mat3x3) ScaleY(Δ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1] * Δ,
			m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1] * Δ,
			m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1] * Δ,
			m[2][2],
		},
	}
}

func (m Mat3x3) ShearX(τ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][0]*τ + m[0][1],
			m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][0]*τ + m[1][1],
			m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][0]*τ + m[2][1],
			m[2][2],
		},
	}
}

func (m Mat3x3) ShearXY(τx, τy float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] + m[0][1]*τy,
			m[0][0]*τx + m[0][1],
			m[0][2],
		},
		Mat1x3{
			m[1][0] + m[1][1]*τy,
			m[1][0]*τx + m[1][1],
			m[1][2],
		},
		Mat1x3{
			m[2][0] + m[2][1]*τy,
			m[2][0]*τx + m[2][1],
			m[2][2],
		},
	}
}

func (m Mat3x3) ShearY(τ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0] + m[0][1]*τ,
			m[0][1],
			m[0][2],
		},
		Mat1x3{
			m[1][0] + m[1][1]*τ,
			m[1][1],
			m[1][2],
		},
		Mat1x3{
			m[2][0] + m[2][1]*τ,
			m[2][1],
			m[2][2],
		},
	}
}

func (a Mat3x3) Subtract(b Mat3x3) Mat3x3 {
	return Mat3x3{
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
		Mat1x3{
			a[2][0] - b[2][0],
			a[2][1] - b[2][1],
			a[2][2] - b[2][2],
		},
	}
}

func (m Mat3x3) SubtractScalar(f float64) Mat3x3 {
	return Mat3x3{
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
		Mat1x3{
			m[2][0] - f,
			m[2][1] - f,
			m[2][2] - f,
		},
	}
}

func (m Mat3x3) Trace() float64 {
	return m[0][0] + m[1][1] + m[2][2]
}

func (m Mat3x3) TranslateX(δ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][0]*δ + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][0]*δ + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][0]*δ + m[2][2],
		},
	}
}

func (m Mat3x3) TranslateXY(δ Vec2) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][0]*δ[0] + m[0][1]*δ[1] + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][0]*δ[0] + m[1][1]*δ[1] + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][0]*δ[0] + m[2][1]*δ[1] + m[2][2],
		},
	}
}

func (m Mat3x3) TranslateY(δ float64) Mat3x3 {
	return Mat3x3{
		Mat1x3{
			m[0][0],
			m[0][1],
			m[0][1]*δ + m[0][2],
		},
		Mat1x3{
			m[1][0],
			m[1][1],
			m[1][1]*δ + m[1][2],
		},
		Mat1x3{
			m[2][0],
			m[2][1],
			m[2][1]*δ + m[2][2],
		},
	}
}

func (m Mat3x3) Transpose() Mat3x3 {
	return Mat3x3{
		Mat1x3{m[0][0], m[1][0], m[2][0]},
		Mat1x3{m[0][1], m[1][1], m[2][1]},
		Mat1x3{m[0][2], m[1][2], m[2][2]},
	}
}

type Mat3x4 [3]Mat1x4

func (a Mat3x4) Add(b Mat3x4) Mat3x4 {
	return Mat3x4{
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
		Mat1x4{
			a[2][0] + b[2][0],
			a[2][1] + b[2][1],
			a[2][2] + b[2][2],
			a[2][3] + b[2][3],
		},
	}
}

func (m Mat3x4) AddScalar(f float64) Mat3x4 {
	return Mat3x4{
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
		Mat1x4{
			m[2][0] + f,
			m[2][1] + f,
			m[2][2] + f,
			m[2][3] + f,
		},
	}
}

/*
func (m Mat3x4) Inverse() Mat4x3 {
	if(Identity3() == m.MultiplyMat4x3(m.Transpose())) {

	}
}
*/
func (m Mat3x4) Len() int {
	return 12
}

func (m Mat3x4) Mat3x3() Mat3x3 {
	return Mat3x3{
		Mat1x3{m[0][0], m[0][1], m[0][2]},
		Mat1x3{m[1][0], m[1][1], m[1][2]},
		Mat1x3{m[2][0], m[2][1], m[2][2]},
	}
}

func (m Mat3x4) MultiplyScalar(f float64) Mat3x4 {
	return Mat3x4{
		Mat1x4{m[0][0] * f, m[0][1] * f, m[0][2] * f, m[0][3] * f},
		Mat1x4{m[1][0] * f, m[1][1] * f, m[1][2] * f, m[1][3] * f},
		Mat1x4{m[2][0] * f, m[2][1] * f, m[2][2] * f, m[2][3] * f},
	}
}

func (a Mat3x4) MultiplyMat4x1(b Mat4x1) Mat3x1 {
	return Mat3x1{
		Mat1x1{a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0]},
		Mat1x1{a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0]},
		Mat1x1{a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0] + a[2][3]*b[3][0]},
	}
}

func (a Mat3x4) MultiplyMat4x2(b Mat4x2) Mat3x2 {
	return Mat3x2{
		Mat1x2{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1] + a[0][3]*b[3][1],
		},
		Mat1x2{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1] + a[1][3]*b[3][1],
		},
		Mat1x2{
			a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0] + a[2][3]*b[3][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1] + a[2][3]*b[3][1],
		},
	}
}

func (a Mat3x4) MultiplyMat4x3(b Mat4x3) Mat3x3 {
	return Mat3x3{
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
		Mat1x3{
			a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0] + a[2][3]*b[3][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1] + a[2][3]*b[3][1],
			a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2] + a[2][3]*b[3][2],
		},
	}
}

func (a Mat3x4) MultiplyMat4x4(b Mat4x4) Mat3x4 {
	return Mat3x4{
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
		Mat1x4{
			a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0] + a[2][3]*b[3][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1] + a[2][3]*b[3][1],
			a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2] + a[2][3]*b[3][2],
			a[2][0]*b[0][3] + a[2][1]*b[1][3] + a[2][2]*b[2][3] + a[2][3]*b[3][3],
		},
	}
}

func (m Mat3x4) MultiplyVec4(v Vec4) Vec3 {
	return Vec3{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2] + m[0][3]*v[3],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2] + m[1][3]*v[3],
		m[2][0]*v[0] + m[2][1]*v[1] + m[2][2]*v[2] + m[2][3]*v[3],
	}
}

func (a Mat3x4) Subtract(b Mat3x4) Mat3x4 {
	return Mat3x4{
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
		Mat1x4{
			a[2][0] - b[2][0],
			a[2][1] - b[2][1],
			a[2][2] - b[2][2],
			a[2][3] - b[2][3],
		},
	}
}

func (m Mat3x4) SubtractScalar(f float64) Mat3x4 {
	return Mat3x4{
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
		Mat1x4{
			m[2][0] - f,
			m[2][1] - f,
			m[2][2] - f,
			m[2][3] - f,
		},
	}
}

func (m Mat3x4) Transpose() Mat4x3 {
	return Mat4x3{
		Mat1x3{m[0][0], m[1][0], m[2][0]},
		Mat1x3{m[0][1], m[1][1], m[2][1]},
		Mat1x3{m[0][2], m[1][2], m[2][2]},
		Mat1x3{m[0][3], m[1][3], m[2][3]},
	}
}

type Mat3xN []float64
