/* */
package geometry

type Vec1 [1]float64

func (v Vec1) Add(a Vec1) Vec1 {
	return Vec1{v[0] + a.X()}
}

func (v Vec1) Component(c int) float64 {
	return v[0]
}

func (v Vec1) Components() []float64 {
	return v[:]
}

func (v Vec1) Dot(d Vec1) float64 {
	return v[0] * d.X()
}

//@TODO Implement approximate equality comparison, as float calculations tend to differ.
func (v Vec1) Equal(c Vec1) bool {
	return v[0] == c.X()
}

func (v Vec1) Len() int {
	return 1
}

func (v Vec1) Multiply(m float64) Vec1 {
	return Vec1{v[0] * m}
}

//Normalizing a 1D vector does nothing, so does this.
func (v Vec1) Normalize() Vec1 {
	return Vec1{v[0]}
}

func (v Vec1) R() float64 {
	return v[0]
}

func (v Vec1) S() float64 {
	return v[0]
}

func (v Vec1) Subtract(s Vec1) Vec1 {
	return Vec1{v[0] - s.X()}
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

func (v Vec1) VecN() VecN {
	return VecN(v[:])
}

func (v Vec1) X() float64 {
	return v[0]
}
