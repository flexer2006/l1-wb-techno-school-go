package twentyfour

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func Distance(a, b Point) float64 {
	return math.Hypot(a.x-b.x, a.y-b.y)
}
