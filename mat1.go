/* Mat1xN implements row vectors. Whereas the VecN types are akin to column vectors */
package geometry

type Mat1x1 [1]float64

func (m Mat1x1) Len() int {
	return 1
}

func (m Mat1x1) Multiply(v Vec1) float64 {
	return m[0] * v[0]
}

type Mat1x2 [2]float64

func (m Mat1x2) Len() int {
	return 2
}

func (m Mat1x2) Multiply(v Vec2) float64 {
	return m[0]*v[0] + m[1]*v[1]
}

func (m Mat1x2) Transpose() Mat2x1 {
	return Mat2x1{
		Mat1x1{m[0]},
		Mat1x1{m[1]},
	}
}

type Mat1x3 [3]float64

func (m Mat1x3) Len() int {
	return 3
}

func (m Mat1x3) Multiply(v Vec3) float64 {
	return m[0]*v[0] + m[1]*v[1] + m[2]*v[2]
}

func (m Mat1x3) Transpose() Mat3x1 {
	return Mat3x1{
		Mat1x1{m[0]},
		Mat1x1{m[1]},
		Mat1x1{m[2]},
	}
}

type Mat1x4 [4]float64

func (m Mat1x4) Len() int {
	return 4
}

func (m Mat1x4) Multiply(v Vec4) float64 {
	return m[0]*v[0] + m[1]*v[1] + m[2]*v[2] + m[3]*v[3]
}

func (m Mat1x4) Transpose() Mat4x1 {
	return Mat4x1{
		Mat1x1{m[0]},
		Mat1x1{m[1]},
		Mat1x1{m[2]},
		Mat1x1{m[3]},
	}
}
