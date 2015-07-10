/* */
package geometry

import (
	"math"
)

type Vector interface {
	Assessor
	Component(int) float64
	Components() []float64
	VecN() VecN
}

type IDVector interface {
	IDVectorOperator
	IDVectorAccessor
}

type IDVectorAccessor interface {
	Vector
	X() float64
}

//@TODO More Vec1 related operations
type IDVectorOperator interface {
	Add(Vec1) Vec1
	Dot(Vec1) float64
	Equal(Vec1) bool
	Multiply(float64) Vec1
	Subtract(Vec1) Vec1
	Square() float64
}

type IIDVector interface {
	IIDVectorAccessor
	IIDVectorOperator
}

type IIDVectorAccessor interface {
	IDVectorAccessor
	Vec1() Vec1
	XY() Vec2
	Y() float64
}

type IIDVectorOperator interface {
	Add(Vec2) Vec2
	Cross(Vec2) Vec2
	Dot(Vec2) float64
	Equal(Vec2) bool
	Length() float64
	Multiply(float64) Vec2
	Normalize() Vec2
	Square() float64
	Subtract(Vec2) Vec2
}

type IIIDVector interface {
	IIIDVectorAccessor
	IIIDVectorOperator
}

type IIIDVectorAccessor interface {
	IIDVectorAccessor
	Vec2() Vec2
	XYZ() Vec3
	Z() float64
}

type IIIDVectorOperator interface {
	Add(Vec3) Vec3
	Cross(Vec3) Vec3
	Dot(Vec3) float64
	Equal(Vec3) bool
	Length() float64
	Multiply(float64) Vec3
	Normalize() Vec3
	Square() float64
	Subtract(Vec3) Vec3
}

type IVDVector interface {
	IVDVectorAccessor
	IVDVectorOperator
}

type IVDVectorAccessor interface {
	IIIDVectorAccessor
	Vec3() Vec3
	W() float64
	XYZW() Vec4
}

type IVDVectorOperator interface {
	Add(Vec4) Vec4
	Cross(Vec4) Vec4
	Dot(Vec4) float64
	Equal(Vec4) bool
	Length() float64
	Multiply(float64) Vec4
	Normalize() Vec4
	Square() float64
	Subtract(Vec4) Vec4
}

type NDVector interface {
	NDVectorAccessor
	NDVectorOperator
}

type NDVectorAccessor interface {
	IVDVectorAccessor
}

type NDVectorOperator interface {
	Add(Vector) VecN
	Cross(Vector) VecN
	CrossLength(Vector) float64
	Dot(Vector) float64
	Equal(Vector) bool
	Length() float64
	Multiply(float64) VecN
	Normalize() VecN
	Square() float64
	Subtract(Vector) VecN
}

type IDColorAccessor interface {
	R() float64
}

type IIDColorAccessor interface {
	IDColorAccessor
	G() float64
	RG() Vec2
}

type IIIDColorAccessor interface {
	IIDColorAccessor
	B() float64
	RGB() Vec3
}

type IVDColorAccessor interface {
	IIIDColorAccessor
	A() float64
	RGBA() Vec4
}

//@TODO Remove this redundant interface.
type ColorVector interface {
	Vector
}

type IDColorVector interface {
	IDColorAccessor
	IDVector
}

type IIDColorVector interface {
	IIDColorAccessor
	IIDVector
}

type IIIDColorVector interface {
	IIIDColorAccessor
	IIIDVector
}

type IVDColorVector interface {
	IVDColorAccessor
	IVDVector
}

type NDColorVector interface {
	IVDColorVector
}

type IDTextureMapAccessor interface {
	S() float64
}

type IIDTextureMapAccessor interface {
	IDTextureMapAccessor
	T() float64
	ST() Vec2
}

type IIIDTextureMapAccessor interface {
	IIDTextureMapAccessor
	P() float64
	STP() Vec3
}

type IVDTextureMapAccessor interface {
	IIIDTextureMapAccessor
	Q() float64
	STPQ() Vec4
}

//@TODO Remove this redundant interface.
type TextureMapVector interface {
	Vector
}

type IDTextureMapVector interface {
	IDTextureMapAccessor
	IDVector
}

type IIDTextureMapVector interface {
	IIDTextureMapAccessor
	IIDVector
}

type IIIDTextureMapVector interface {
	IIIDTextureMapAccessor
	IIIDVector
}

type IVDTextureMapVector interface {
	IVDTextureMapAccessor
	IVDVector
}

type NDTextureMapVector interface {
	IVDTextureMapVector
}

func AddVec1(a, b IDVectorAccessor) IDVector {
	return Vec1{a.X() + b.X()}
}

func AddVec2(a, b IIDVectorAccessor) IIDVector {
	return Vec2{a.X() + b.X(), a.Y() + b.Y()}
}

func AddVec3(a, b IIIDVectorAccessor) IIIDVector {
	return Vec3{a.X() + b.X(), a.Y() + b.Y(), a.Z() + b.Z()}
}

func AddVec4(a, b IVDVectorAccessor) IVDVector {
	return Vec4{a.X() + b.X(), a.Y() + b.Y(), a.Z() + b.Z(), a.W() + b.W()}
}

func AddVecN(a, b Vector) NDVector {
	if a.Len() < 1 {
		return VecN(a.Components())
	}

	o := VecN{}

	for i, c := range b.Components() {
		o = append(o, a.Component(i)+c)
	}

	return o
}

func CrossVec1(a, b Vector) IDVector {
	return Vec1{0.0}
}

func CrossVec3(a, b IIIDVectorAccessor) IIIDVector {
	return Vec3{a.Y()*b.Z() - a.Z()*b.Y(), a.Z()*b.X() - a.X()*b.Z(), a.X()*b.Y() - a.Y()*b.X()}
}

func DotVec1(a, b IDVectorAccessor) float64 {
	return a.X() * b.X()
}

func DotVec2(a, b IIDVectorAccessor) float64 {
	return a.X()*b.X() + a.Y()*b.Y()
}

func DotVec3(a, b IIIDVectorAccessor) float64 {
	return a.X()*b.X() + a.Y()*b.Y() + a.Z()*b.Z()
}

func DotVec4(a, b IVDVectorAccessor) float64 {
	return a.X()*b.X() + a.Y()*b.Y() + a.Z()*b.Z() + a.W()*b.W()
}

func DotVecN(a, b Vector) float64 {
	if b.Len() < 1 {
		return SumVecN(a)
	}

	r := 0.0

	for i, c := range a.Components() {
		r += b.Component(i) * c
	}

	return r
}

func EqualVec1(a, b IDVectorAccessor) bool {
	return a.X() == b.X()
}

func EqualVec2(a, b IIDVectorAccessor) bool {
	return a.X() == b.X() && a.Y() == b.Y()
}

func EqualVec3(a, b IIIDVectorAccessor) bool {
	return a.X() == b.X() && a.Y() == b.Y() && a.Z() == b.Z()
}

func EqualVec4(a, b IVDVectorAccessor) bool {
	return a.X() == b.X() && a.Y() == b.Y() && a.Z() == b.Z() && a.W() == b.W()
}

func EqualVecN(a, b Vector) bool {
	if b.Len() != a.Len() {
		return false
	}

	for i, c := range a.Components() {
		if c != b.Component(i) {
			return false
		}
	}

	return true
}

func LengthVec1(a IDVectorAccessor) float64 {
	return math.Sqrt(DotVec1(a, a))
}

func LengthVec2(a IIDVectorAccessor) float64 {
	return math.Sqrt(DotVec2(a, a))
}

func LengthVec3(a IIIDVectorAccessor) float64 {
	return math.Sqrt(DotVec3(a, a))
}

func LengthVec4(a IVDVectorAccessor) float64 {
	return math.Sqrt(DotVec4(a, a))
}

func LengthVecN(a Vector) float64 {
	return math.Sqrt(DotVecN(a, a))
}

func MultiplyVec1(a IDVectorAccessor, m float64) IDVector {
	return Vec1{a.X() * m}
}

func MultiplyVec2(a IIDVectorAccessor, m float64) IIDVector {
	return Vec2{a.X() * m, a.Y() * m}
}

func MultiplyVec3(a IIIDVectorAccessor, m float64) IIIDVector {
	return Vec3{a.X() * m, a.Y() * m, a.Z() * m}
}

func MultiplyVec4(a IVDVectorAccessor, m float64) IVDVector {
	return Vec4{a.X() * m, a.Y() * m, a.Z() * m, a.W() * m}
}

func MultiplyVecN(a Vector, m float64) NDVector {
	p := VecN{}

	for _, n := range a.Components() {
		p = append(p, n*m)
	}

	return p
}

func NormalizeVec1(a IDVectorAccessor) IDVector {
	return Vec1{a.X()}
}

func NormalizeVec2(a IIDVectorAccessor) IIDVector {
	d := 1.0 / LengthVec1(a)
	return MultiplyVec2(a, d)
}

func NormalizeVec3(a IIIDVectorAccessor) IIIDVector {
	d := 1.0 / LengthVec1(a)
	return MultiplyVec3(a, d)
}

func NormalizeVec4(a IVDVectorAccessor) IVDVector {
	d := 1.0 / LengthVec1(a)
	return MultiplyVec4(a, d)
}

func NormalizeVecN(a Vector) NDVector {
	if a.Len() < 2 {
		return VecN(a.Components())
	}

	d := 1.0 / LengthVecN(a)
	return MultiplyVecN(a, d)
}

func SubtractVec1(a, b IDVectorAccessor) IDVector {
	return Vec1{a.X() - b.X()}
}

func SubtractVec2(a, b IIDVectorAccessor) IIDVector {
	return Vec2{a.X() - b.X(), a.Y() - b.Y()}
}

func SubtractVec3(a, b IIIDVectorAccessor) IIIDVector {
	return Vec3{a.X() - b.X(), a.Y() - b.Y(), a.Z() - b.Z()}
}

func SubtractVec4(a, b IVDVectorAccessor) IVDVector {
	return Vec4{a.X() - b.X(), a.Y() - b.Y(), a.Z() - b.Z(), a.W() - b.W()}
}

func SubtractVecN(a, b Vector) NDVector {
	if b.Len() < 1 {
		return VecN(a.Components())
	}

	d := VecN{}

	for i, n := range a.Components() {
		d = append(d, n-b.Component(i))
	}

	return d
}

func SumVec1(a IDVectorAccessor) float64 {
	return a.X()
}

func SumVec2(a IIDVectorAccessor) float64 {
	return a.X() + a.Y()
}

func SumVec3(a IIIDVectorAccessor) float64 {
	return a.X() + a.Y() + a.Z()
}

func SumVec4(a IVDVectorAccessor) float64 {
	return a.X() + a.Y() + a.Z() + a.W()
}

func SumVecN(a Vector) float64 {
	s := 0.0

	for _, n := range a.Components() {
		s += n
	}

	return s
}
