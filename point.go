package gog

import (
	"math"

	"golang.org/x/image/math/fixed"
)

// A Point is an X, Y coordinate pair. The axes increase right and down.
type Point struct {
	X, Y float64
}

// Add returns the vector p+q.
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Sub returns the vector p-q.
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// Mul returns xy * Point.
func (p Point) Mul(factor Point) Point {
	return Point{p.X * factor.X, p.Y * factor.Y}
}

// Div returns the vector p/k.
func (p Point) Div(k float64) Point {
	return Point{p.X / k, p.Y / k}
}

func (a Point) Fixed() fixed.Point26_6 {
	return fixed.Point26_6{X: fixed.Int26_6(a.X * 64), Y: fixed.Int26_6(a.Y * 64)}
}

// Distance to other point
func (a Point) Distance(other Point) float64 {
	return math.Hypot(a.X-other.X, a.Y-other.Y)
}

// Lerp Linear interpolates points
func (a Point) Lerp(other Point, t float64) Point {
	// a + (b-a) * t
	return a.Add(other.Sub(a).Mul(Point{t, t}))
}

// Rotate rotates point about origin o
func (a Point) Rotate(angle float64, o Point) Point {
	b := Point{}
	b.X = math.Cos(angle)*(a.X-o.X) - math.Sin(angle)*(a.Y-o.Y) + o.X
	b.Y = math.Sin(angle)*(a.X-o.X) + math.Cos(angle)*(a.Y-o.Y) + o.Y
	return b
}
