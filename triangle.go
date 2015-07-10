/**/

package geometry

type Triangle [3]*Vertex

func (t Triangle) Edge(i int) Edge {
	return Edge{t[i%3], t[(i+1)%3]}
}

func (t Triangle) Edges() (e []Edge) {
	edge := Edge{}
	for _, item := range ExtractVertexSlices(t, EdgePoly) {
		copy(edge[:], item[:])
		e = append(e, edge)
	}

	return
}

func (t Triangle) Len() int {
	return len(t)
}

func (t Triangle) Ngon() (n Ngon) {
	return Ngon(t[:])
}

func (t Triangle) Quadrilateral(i int) Quadrilateral {
	return Quadrilateral{t[i%3], t[(i+1)%3], t[(i+2)%3], t[(i+3)%3]}
}

func (t Triangle) Triangle(i int) Triangle {
	return Triangle{t[i%3], t[(i+1)%3], t[(i+2)%3]}
}

func (t Triangle) Triangles() (ta []Triangle) {
	tri := Triangle{}
	for _, item := range ExtractVertexSlices(t, TrianglePoly) {
		copy(tri[:], item[:])
		ta = append(ta, tri)
	}

	return
}

func (t Triangle) Vertex(i int) *Vertex {
	i %= len(t)
	return &Vertex{t[i].Attributes, t[i].Color, t[i].Normal, t[i].TextureMap, t[i].Vec3}
}

func (t Triangle) Vertices() []*Vertex {
	return t[:]
}
