/**/

package geometry

type Quadrilateral [4]*Vertex

func (q Quadrilateral) Edge(i int) Edge {
	return Edge{q[i%4], q[(i+1)%4]}
}

func (q Quadrilateral) Edges() (e []Edge) {
	edge := Edge{}
	for _, item := range ExtractVertexSlices(q, EdgePoly) {
		copy(edge[:], item[:])
		e = append(e, edge)
	}

	return
}

func (q Quadrilateral) Len() int {
	return len(q)
}

func (q Quadrilateral) Ngon() (n Ngon) {
	return Ngon(q[:])
}

func (q Quadrilateral) Quadrilateral(i int) Quadrilateral {
	return Quadrilateral{q[i%4], q[(i+1)%4], q[(i+2)%4], q[(i+3)%4]}
}

func (q Quadrilateral) Quadrilaterals() (qa []Quadrilateral) {
	quad := Quadrilateral{}
	for _, item := range ExtractVertexSlices(q, QuadrilateralPoly) {
		copy(quad[:], item[:])
		qa = append(qa, quad)
	}

	return
}

func (q Quadrilateral) Triangle(i int) Triangle {
	return Triangle{q[i%4], q[(i+1)%4], q[(i+2)%4]}
}

func (q Quadrilateral) Triangles() (t []Triangle) {
	tri := Triangle{}
	for _, item := range ExtractVertexSlices(q, TrianglePoly) {
		copy(tri[:], item[:])
		t = append(t, tri)
	}

	return
}

func (q Quadrilateral) Vertex(i int) *Vertex {
	i %= 4
	return &Vertex{q[i].Attributes, q[i].Color, q[i].Normal, q[i].TextureMap, q[i].Vec3}
}

func (q Quadrilateral) Vertices() []*Vertex {
	return q[:]
}
