/**/

package geometry

type Vertex struct {
	Attributes []Attribute
	Color      Vec4
	Normal     Vec3
	TextureMap Vec2
	Vec3       //Position
}

func (v *Vertex) Len() int {
	return 1
}

func (v *Vertex) Vertex(i int) *Vertex {
	return v
}

func (v *Vertex) Vertices() []Vertex {
	return []Vertex{*v}
}
