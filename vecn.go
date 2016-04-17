/* */
package geometry

import (
	"math"
)

type VecN []float64

func (v VecN) A() float64 {
	return v.Component(3)
}

//Adds the two vectors together, possibly returning a vector containing more components than the receiver.
func (v VecN) Add(a Vector) VecN {
	if a.Len() < 1 {
		return VecN(v[:])
	}

	o := VecN{}

	for i, c := range a.Components() {
		o = append(o, v.Component(i)+c)
	}

	return o
}

func (v VecN) AddScalar(a float64) VecN {
	if a == 0 {
		return VecN(v[:])
	}

	o := VecN{}

	for i, c := range v.Components() {
		o = append(o, v.Component(i)+c)
	}

	return o
}

func (v VecN) B() float64 {
	return v.Component(2)
}

func (v VecN) Component(c int) float64 {
	if c > v.Len()-1 {
		return 0
	}

	return v[c]
}

func (v VecN) Components() []float64 {
	return v[:]
}

//Returns v Cross c, limited to the minimum size (number of components) shared by the vectors. 0 and 1 dimensional vector cross products are always 0
func (v VecN) Cross(c Vector) VecN {
	vlen, clen, d := v.Len(), c.Len(), 0

	if clen > vlen {
		d = vlen
	} else {
		d = clen
	}

	switch {
	case d >= 7:
		//@TODO Implement 7D Cross
	case d >= 3:
		return VecN{v.Component(1)*c.Component(2) - v.Component(2)*c.Component(1), v.Component(2)*c.Component(0) - v.Component(0)*c.Component(2), v.Component(0)*c.Component(1) - v.Component(1)*c.Component(0)}
	}

	return VecN{0.0}
}

//Returns the length of V Cross C. Can be used to determine orthogonality.
func (v VecN) CrossLength(c Vector) float64 {
	if v.Len() < 1 && c.Len() < 1 {
		return 0.0
	}

	cd := 0.0

	//Dot for Vector @TODO: Remove if Vector gets a Dot or argument gets changed to a better type
	for _, co := range c.Components() {
		cd += co * co
	}

	vd := v.Dot(v)

	return math.Sqrt(vd * cd * (1 - (v.Dot(c) / (vd * cd))))
}

func (v VecN) Divide(d Vector) VecN {
	if v.Len() < 1 || d.Len() < 1 {
		return VecN(v[:])
	}
	p := VecN{}

	for i, n := range v {
		if n == 0 || d.Component(i) == 0 {
			p = append(p, 0)
		} else {
			p = append(p, n/d.Component(i))
		}
	}

	return p
}

func (v VecN) DivideScalar(d float64) VecN {
	if d == 0 {
		return make(VecN, v.Len())
	}
	p := VecN{}

	for _, n := range v {
		if n == 0 {
			p = append(p, 0)
		} else {
			p = append(p, n/d)
		}
	}

	return p
}

func (v VecN) Dot(d Vector) float64 {
	if d.Len() < 1 {
		return v.Sum()
	}

	r := 0.0

	for i, c := range v {
		r += d.Component(i) * c
	}

	return r
}

func (v VecN) Equal(e Vector) bool {
	if e.Len() != v.Len() {
		return false
	}

	for i, c := range v {
		if c != e.Component(i) {
			return false
		}
	}

	return true
}

func (v VecN) Expand() VecN {
	return VecN(append(v[:], 0))
}

func (v VecN) G() float64 {
	return v.Component(1)
}

func (v VecN) Homogenize() VecN {
	return VecN(append(v[:], 1))
}

func (v VecN) Len() int {
	return len(v)
}

func (v VecN) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v VecN) MultiplyScalar(m float64) VecN {
	p := VecN{}

	for _, n := range v {
		p = append(p, n*m)
	}

	return p
}

func (v VecN) Normalize() VecN {
	d := 1.0 / v.Length()
	return v.MultiplyScalar(d)
}

func (v VecN) P() float64 {
	return v.Component(2)
}

func (v VecN) Q() float64 {
	return v.Component(3)
}

func (v VecN) R() float64 {
	return v.Component(0)
}

func (v VecN) RG() Vec2 {
	return Vec2{v.Component(0), v.Component(1)}
}

func (v VecN) RGB() Vec3 {
	return Vec3{v.Component(0), v.Component(1), v.Component(2)}
}

func (v VecN) RGBA() Vec4 {
	return Vec4{v.Component(0), v.Component(1), v.Component(2), v.Component(3)}
}

func (v VecN) S() float64 {
	return v.Component(0)
}

//Returns the Dot product of itself. For VecN this is faster than v.Dot(v)
func (v VecN) Square() float64 {
	s := 0.0

	for _, n := range v {
		s += n * n
	}

	return s
}

func (v VecN) ST() Vec2 {
	return Vec2{v.Component(0), v.Component(1)}
}

func (v VecN) STP() Vec3 {
	return Vec3{v.Component(0), v.Component(1), v.Component(2)}
}

func (v VecN) STPQ() Vec4 {
	return Vec4{v.Component(0), v.Component(1), v.Component(2), v.Component(3)}
}

func (v VecN) Subtract(s Vector) VecN {
	if s.Len() < 1 {
		return VecN(v[:])
	}

	d := VecN{}

	for i, n := range v {
		d = append(d, n-s.Component(i))
	}

	return d
}

func (v VecN) SubtractScalar(s float64) VecN {
	if s == 0 {
		return VecN(v[:])
	}

	d := VecN{}

	for _, n := range v {
		d = append(d, n-s)
	}

	return d
}

func (v VecN) Sum() float64 {
	s := 0.0

	for _, n := range v {
		s += n
	}

	return s
}

func (v VecN) T() float64 {
	return v.Component(1)
}

func (v VecN) U() float64 {
	return v.Component(0)
}

func (v VecN) UV() Vec2 {
	return Vec2{v.Component(0), v.Component(1)}
}

func (v VecN) UVP() Vec3 {
	return Vec3{v.Component(0), v.Component(1), v.Component(2)}
}

func (v VecN) UVPQ() Vec4 {
	return Vec4{v.Component(0), v.Component(1), v.Component(2), v.Component(3)}
}

func (v VecN) V() float64 {
	return v.Component(1)
}

func (v VecN) Vec1() Vec1 {
	return Vec1{v.Component(0)}
}

func (v VecN) Vec2() Vec2 {
	return Vec2{v.Component(0), v.Component(1)}
}

func (v VecN) Vec3() Vec3 {
	return Vec3{v.Component(0), v.Component(1), v.Component(2)}
}

func (v VecN) Vec4() Vec4 {
	return Vec4{v.Component(0), v.Component(1), v.Component(2), v.Component(3)}
}

func (v VecN) VecN() VecN {
	return VecN(v[:])
}

func (v VecN) W() float64 {
	return v.Component(3)
}

func (v VecN) X() float64 {
	return v.Component(0)
}

func (v VecN) XY() Vec2 {
	return Vec2{v.Component(0), v.Component(1)}
}

func (v VecN) XYZ() Vec3 {
	return Vec3{v.Component(0), v.Component(1), v.Component(2)}
}

func (v VecN) XYZW() Vec4 {
	return Vec4{v.Component(0), v.Component(1), v.Component(2), v.Component(3)}
}

func (v VecN) Y() float64 {
	return v.Component(1)
}

func (v VecN) Z() float64 {
	return v.Component(2)
}
