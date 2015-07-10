/* */
package geometry

type Mat3x1 [3]Mat1x1

func (m Mat3x1) Len() int {
	return 3
}

func (m Mat3x1) Multiply(v Vec1) Vec3 {
	return Vec3{
		m[0][0] * v[0],
		m[1][0] * v[0],
		m[2][0] * v[0],
	}
}

func (m Mat3x1) Transpose() Mat1x3 {
	return Mat1x3{m[0][0], m[1][0], m[2][0]}
}

type Mat3x2 [3]Mat1x2

func (m Mat3x2) Len() int {
	return 6
}

func (m Mat3x2) Multiply(v Vec2) Vec3 {
	return Vec3{
		m[0][0]*v[0] + m[0][1]*v[1],
		m[1][0]*v[0] + m[1][1]*v[1],
		m[2][0]*v[0] + m[2][1]*v[1],
	}
}

func (m Mat3x2) Transpose() Mat2x3 {
	return Mat2x3{
		Mat1x3{m[0][0], m[1][0], m[2][0]},
		Mat1x3{m[0][1], m[1][1], m[2][1]},
	}
}

type Mat3x3 [3]Mat1x3

func (m Mat3x3) Determinant() float64 {
	return m[0][0]*m[1][1]*m[2][2] + m[0][1]*m[1][2]*m[2][0] + m[0][2]*m[1][0]*m[2][1] - m[0][2]*m[1][1]*m[2][0] - m[0][1]*m[1][0]*m[2][2] - m[0][0]*m[1][2]*m[2][0]
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

func (m Mat3x3) Multiply(v Vec3) Vec3 {
	return Vec3{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2],
		m[2][0]*v[0] + m[2][1]*v[1] + m[2][2]*v[2],
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

func (m Mat3x4) Len() int {
	return 12
}

func (m Mat3x4) Multiply(v Vec4) Vec3 {
	return Vec3{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2] + m[0][3]*v[3],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2] + m[1][3]*v[3],
		m[2][0]*v[0] + m[2][1]*v[1] + m[2][2]*v[2] + m[2][3]*v[3],
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
