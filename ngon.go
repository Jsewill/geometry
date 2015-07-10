/**/

package geometry

type Ngon []*Vertex

func (n Ngon) Edge(i int) Edge {
	l := len(n)
	return Edge{n[i%l], n[(i+1)%l]}
}

func (n Ngon) Edges() (e []Edge) {
	edge := Edge{}
	for _, item := range ExtractVertexSlices(n, EdgePoly) {
		copy(edge[:], item[:])
		e = append(e, edge)
	}

	return
}

func (n Ngon) Len() int {
	return len(n)
}

func (n Ngon) Ngon() Ngon {
	return Ngon(n[:])
}

func (n Ngon) Quadrilateral(i int) Quadrilateral {
	l := len(n)
	return Quadrilateral{n[i%l], n[(i+1)%l], n[(i+2)%l], n[(i+3)%l]}
}

func (n Ngon) Quadrilaterals() (q []Quadrilateral) {
	quad := Quadrilateral{}
	for _, item := range ExtractVertexSlices(n, QuadrilateralPoly) {
		copy(quad[:], item[:])
		q = append(q, quad)
	}
	return
}

func (n Ngon) Triangle(i int) Triangle {
	l := len(n)
	return Triangle{n[i%l], n[(i+1)%l], n[(i+2)%l]}
}

func (n Ngon) Triangles() (t []Triangle) {
	tri := Triangle{}
	for _, item := range ExtractVertexSlices(n, TrianglePoly) {
		copy(tri[:], item[:])
		t = append(t, tri)
	}

	return
}

func (n Ngon) Vertex(i int) *Vertex {
	return &Vertex{n[i].Attributes, n[i].Color, n[i].Normal, n[i].TextureMap, n[i].Vec3}
}

func (n Ngon) Vertices() []*Vertex {
	return n[:]
}
