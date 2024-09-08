package gog

import (
	"math"

	"github.com/setanarut/gog/v2/path"
	"github.com/setanarut/gog/v2/utils"
	"github.com/setanarut/gog/v2/vec"
)

// BBox returns a bounding box path with min and max points.
func BBox(min, max vec.Vec2) *path.Path {
	points := []vec.Vec2{min, {max.X, min.Y}, max, {min.X, max.Y}, min}
	return path.NewPath(points).CalculateLength()
}

// Line returns a line Path
func Line(start, end vec.Vec2) *path.Path {
	return path.NewPath([]vec.Vec2{start, end}).CalculateLength()
}

// CubicBezier returns a cubic-bezier Path.
func CubicBezier(x0, y0, x1, y1, x2, y2, x3, y3 float64, samples int) *path.Path {
	cb := cubicBezier{[4]vec.Vec2{{x0, y0}, {x1, y1}, {x2, y2}, {x3, y3}}}
	p := path.NewPath(cb.flatten(samples)).CalculateLength()
	return p
}

// Rect returns a rectangle-shaped path.Path.
func Rect(topLeft vec.Vec2, w, h float64) *path.Path {
	Sq := path.NewPath([]vec.Vec2{{}, {w, 0}, {w, h}, {0, h}})
	Sq.Close().Translate(topLeft.X, topLeft.Y).CalculateLength()
	return Sq
}

// Square returns Square-shaped path.Path with side.
func Square(topLeft vec.Vec2, side float64) *path.Path {
	return Rect(topLeft, side, side)
}

// Ellipse returns ellipse-shaped path.Path.
func Ellipse(origin vec.Vec2, xRadius, yRadius float64) *path.Path {
	samples := int(utils.Clip(xRadius, 20, 80))
	return EllipseSamples(origin, xRadius, yRadius, samples)
}

// Circle returns circle-shaped path.Path.
func Circle(origin vec.Vec2, radius float64) *path.Path {
	samples := int(utils.Clip(radius, 20, 80))
	return EllipseSamples(origin, radius, radius, samples)
}

// RegularPolygon returns regular polygon shaped path.Path.
func RegularPolygon(origin vec.Vec2, n int, radius float64) *path.Path {
	align_angle := (math.Pi / 2) - (math.Pi*2)/float64(n)/2
	return EllipseSamples(origin, radius, radius, n).Rotate(align_angle)
}

// Spiral returns spriral
//
// n: Loop count
//
// radius: Spiral radius
//
// angleStep: The amount of angle increase with each step.
//
// s := SpiralPoints(250, 250, 0.05)
func Spiral(n int, radius, angleStep float64) *path.Path {
	points := make([]vec.Vec2, n)

	for i := 0; i < n; i++ {
		angle := angleStep * float64(i)         // Açıyı her adımda arttırıyoruz
		r := radius * (float64(i) / float64(n)) // Yarıçap her adımda artıyor

		// x ve y koordinatlarını hesaplıyoruz
		x := r * math.Cos(angle)
		y := r * math.Sin(angle)

		points[i] = vec.Vec2{x, y}
	}
	return path.NewPath(points)
}

// Lemniscate generates the points for an infinity symbol (Lemniscate)
//
// n: Number of points
//
// w: Half width of the infinity symbol
func Lemniscate(n int, w float64) *path.Path {
	points := make([]vec.Vec2, n)
	step := (2 * math.Pi) / float64(n) // Calculate the step size for the angle range from -pi to pi
	for i := 0; i < n; i++ {
		t := -math.Pi + step*float64(i) // Adjust t to be between -pi and pi
		x := w * math.Cos(t) / (1 + math.Pow(math.Sin(t), 2))
		y := w * math.Cos(t) * math.Sin(t) / (1 + math.Pow(math.Sin(t), 2))
		points[i] = vec.Vec2{x, y}
	}
	return path.NewPath(points)
}

// EllipseSamples returns an ellipse-shaped path.Path.
func EllipseSamples(origin vec.Vec2, xRadius, yRadius float64, samples int) *path.Path {
	points := make([]vec.Vec2, 0)
	angleStep := 2 * math.Pi / float64(samples)
	for i := 0; i < samples; i++ {
		angle := float64(i) * angleStep
		pt := vec.Vec2{xRadius * math.Cos(angle), yRadius * math.Sin(angle)}
		points = append(points, pt.Add(origin))
	}

	return path.NewPath(points).Close().CalculateLength()
}
