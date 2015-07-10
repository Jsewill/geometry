/* */

package geometry

import (
	"log"
	"math"
)

type Bounds [3]*Plane //Bounding planes: inner, outer, outer respectively

type Plane struct {
	Vec3
	Distance float64
}

func IsPointInSphere(p, v Vec3, r float64) bool {
	return math.Pow(r, 2) >= v.Subtract(p).Square()
}

func IsPointInCylinder(p, v1, v2 Vec3, r float64) bool {
	ab, ax := v2.Subtract(v1), p.Subtract(v1)
	dot := ax.Dot(ab)

	/*Angle is obtuse (between 0.5∏r and ∏r), point is on the side of the start cap opposite the end cap.
	Therefore, it is outside the cylinder. Essentially, we're testing if the point is between two
	infinite planes (representing the end caps, perpendicular to the edge which the two vertices make up).

	Note, we are checking against vertex + radius in IsPointInSphere().*/
	if dot < 0 {
		return false
	}

	/*Divide by ||ab||², resulting in

	||ab||||ax||cosΘ
	----------------
	     ||ab||²

	Which simply means, if greater than 1, that the point is outside the cylinder, on the wrong side of
	the end cap. Remember those infinite planes? This tests the other cap. Projecting ax onto ab.*/
	if dot /= ab.Square(); dot > 1.0 {
		return false
	}

	//Get distance along the aforementioned projection vector by ab*dot, then making it relative to ax.
	dist := ax.Subtract(ab.Multiply(dot))

	//If the magnitude of the projected vertex exceeds the radius, we're outside the cylinder volume.
	if dist.Square() > math.Pow(r, 2) {
		return false
	}

	return true
}

func IsPointBetweenPlanes(p Vec3, planes Bounds) bool {
	parallel := planes[1].Cross(planes[2].Vec3)

	if (Vec3{0, 0, 0}) != parallel {
		log.Println("Planes are not parallel!")
	}

	return (planes[1].Dot(p)+planes[1].Distance)*(planes[2].Dot(p)+planes[2].Distance) < 0
}

func NewPlane(t Triangle) *Plane {
	p := new(Plane)

	p.Vec3 = t[1].Subtract(t[0].Vec3).Cross(t[2].Subtract(t[1].Vec3)).Normalize()
	p.Distance = -p.Dot(t[0].Vec3)

	return p
}

func (p *Plane) Bounds(radius float64) Bounds {
	tc := p.DistanceToCenter(radius)
	return [3]*Plane{p, &Plane{p.Vec3, p.Distance + tc}, &Plane{p.Vec3, p.Distance - tc}}
}

func (p *Plane) DistanceToCenter(r float64) float64 {
	//@TODO: Implement matrix stuff so this ugly code can go away. (Vec3 through Mat3x3)
	v := Vec3{r, 0, 0}
	b1 := math.Abs(p.Dot(v))
	b2 := math.Abs(p.Dot(Vec3{0, r, 0}))
	b3 := math.Abs(p.Dot(Vec3{0, 0, r}))

	//v.Length() would seem to simply be equal to r.
	cb := math.Max(b1, math.Max(b2, b3)) / (v.Length() * p.Length())
	return r * cb
}

func BoundsFromPolygon(poly Polygon, radius float64) (b []Bounds) {
	for _, t := range poly.Triangles() {
		plane := NewPlane(t)
		b = append(b, plane.Bounds(radius))
	}

	return b
}
