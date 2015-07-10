/* */
package geometry

type Mat2x1 [2]Mat1x1

func (m Mat2x1) Len() int {
	return 2
}

func (m Mat2x1) Multiply(v Vec1) Vec2 {
	return Vec2{
		m[0][0] * v[0],
		m[1][0] * v[0],
	}
}

func (m Mat2x1) Transpose() Mat1x2 {
	return Mat1x2{m[0][0], m[1][0]}
}

type Mat2x2 [2]Mat1x2

func (m Mat2x2) Determinant() float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
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

func (m Mat2x2) Multiply(v Vec2) Vec2 {
	return Vec2{
		m[0][0]*v[0] + m[0][1]*v[1],
		m[1][0]*v[0] + m[1][1]*v[1],
	}
}

func (m Mat2x2) Transpose() Mat2x2 {
	return Mat2x2{
		Mat1x2{m[0][0], m[1][0]},
		Mat1x2{m[0][1], m[1][1]},
	}
}

type Mat2x3 [2]Mat1x3

func (m Mat2x3) Len() int {
	return 6
}

func (m Mat2x3) Multiply(v Vec3) Vec2 {
	return Vec2{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2],
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

func (m Mat2x4) Len() int {
	return 4
}

func (m Mat2x4) Multiply(v Vec4) Vec2 {
	return Vec2{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2] + m[0][3]*v[3],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2] + m[1][3]*v[3],
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
