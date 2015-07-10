/**/

package geometry

type Edge [2]*Vertex

func (e Edge) Edge(i int) Edge {
	return e
}

func (e Edge) Edges() []Edge {
	return []Edge{e}
}

func (e Edge) Len() int {
	return len(e)
}

func (e Edge) Vertex(i int) *Vertex {
	i %= 2
	return &Vertex{e[i].Attributes, e[i].Color, e[i].Normal, e[i].TextureMap, e[i].Vec3}
}

func (e Edge) Vertices() []*Vertex {
	return e[:]
}
