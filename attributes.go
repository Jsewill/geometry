/* */

package geometry

type AttributeType int

const (
	PositionAttribute AttributeType = iota
	ColorAttribute
	NormalAttribute
	TextureMapAttribute
)

type Attribute struct {
	AttributeType
	VecN
}

func (a *Attribute) Set(t AttributeType, v Vector) {
	a.AttributeType = t
	a.VecN = v.Components()
}

func (a *Attribute) Type() AttributeType {
	return a.AttributeType
}
