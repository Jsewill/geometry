/* */

package geometry

type PolyType int

const (
	VertexPoly PolyType = iota + 1
	EdgePoly
	TrianglePoly
	QuadrilateralPoly
)

type Assessor interface {
	Len() int
}

//Zero or more sides.
type Geometry interface {
	Assessor
	Vertex(int) *Vertex
	Vertices() []*Vertex
}

//Many sides.
type Polygon interface {
	Edge(i int) Edge
	Edges() []Edge
	Geometry
	Ngon() Ngon
	Triangle(int) Triangle
	Triangles() []Triangle
}

/* Four or more sides. Many names were considered, including some strange ones:
Tessperigon
Periquadrilateral
Quadriperigon
Peritessigon
*/
type Quadriplurilateral interface {
	Polygon
	Quadrilateral(int) Quadrilateral
	Quadrilaterals() []Quadrilateral
}

// Slice of Slices getters' backbone
func ExtractVertexSlices(p Polygon, components PolyType) [][]*Vertex {
	qty := int(components - 1)
	verts := p.Len()
	limit := verts - qty

	if limit < 1 {
		return nil
	}

	set := make([][]*Vertex, verts)
	vertices := p.Vertices()

	for i := range set {
		wrap := make([]*Vertex, qty+1)

		for j := 0; j < qty+1; j++ {
			wrap[j] = vertices[(i+j)%verts]
		}

		set[i] = wrap[:]
	}

	return set
}
