/* */
package geometry

import (
	"math"
)

const (
	RadiansPerDegree float64 = math.Pi / 180
	DegreesPerRadian float64 = 180 / math.Pi
)

// Convert Degrees to Radians
func Radians(d float64) float64 {
	return d * RadiansPerDegree
}

// Convert Radians to Degrees
func Degrees(r float64) float64 {
	return r * DegreesPerRadian
}
