package gog

import (
	"math"
)

// NewPath returns new Path from points
func NewPath(points []Point) *Path {
	newPath := &Path{
		points: points,
		Style:  DefaultStyle(),
	}

	if len(points) > 1 {
		// if newPath.IsClosed() == false {
		// 	newPath.Close()
		// }
		newPath.SetAnchor(newPath.Centroid())
	}
	newPath.calculateLength()
	return newPath
}

// BBox returns a bounding box path with min and max points.
func BBox(min, max Point) *Path {
	points := []Point{min, {max.X, min.Y}, max, {min.X, max.Y}, min}
	return NewPath(points).calculateLength()
}

// Line returns a line Path.
func Line(start, end Point) *Path {
	return NewPath([]Point{start, end}).calculateLength()
}

// CubicBezier returns a cubic-bezier Path.
func CubicBezier(x0, y0, x1, y1, x2, y2, x3, y3 float64, samples int) *Path {
	cb := cubicBezier{[4]Point{{x0, y0}, {x1, y1}, {x2, y2}, {x3, y3}}}
	p := NewPath(cb.flatten(samples)).calculateLength()
	return p
}

// Rect returns a rectangle-shaped Path.
func Rect(topLeft Point, w, h float64) *Path {
	Sq := NewPath([]Point{{}, {w, 0}, {w, h}, {0, h}})
	Sq.Close().Translate(topLeft.X, topLeft.Y).calculateLength()
	return Sq
}

// Square returns Square-shaped Path with side.
func Square(topLeft Point, side float64) *Path {
	return Rect(topLeft, side, side)
}

// Ellipse returns ellipse-shaped Path.
func Ellipse(origin Point, xRadius, yRadius float64) *Path {
	samples := int(clip(xRadius, 20, 80))
	return EllipseSamples(origin, xRadius, yRadius, samples)
}

// Circle returns circle-shaped Path.
func Circle(origin Point, radius float64) *Path {
	samples := int(clip(radius, 20, 80))
	return EllipseSamples(origin, radius, radius, samples)
}

// RegularPolygon returns regular polygon shaped Path.
func RegularPolygon(origin Point, n int, radius float64) *Path {
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
func Spiral(n int, radius, angleStep float64) *Path {
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		angle := angleStep * float64(i)         // Açıyı her adımda arttırıyoruz
		r := radius * (float64(i) / float64(n)) // Yarıçap her adımda artıyor

		// x ve y koordinatlarını hesaplıyoruz
		x := r * math.Cos(angle)
		y := r * math.Sin(angle)

		points[i] = Point{x, y}
	}
	return NewPath(points)
}

// Lemniscate generates the points for an infinity symbol (Lemniscate)
//
// n: Number of points
//
// w: Half width of the infinity symbol
func Lemniscate(n int, w float64) *Path {
	points := make([]Point, n)
	step := (2 * math.Pi) / float64(n) // Calculate the step size for the angle range from -pi to pi
	for i := 0; i < n; i++ {
		t := -math.Pi + step*float64(i) // Adjust t to be between -pi and pi
		x := w * math.Cos(t) / (1 + math.Pow(math.Sin(t), 2))
		y := w * math.Cos(t) * math.Sin(t) / (1 + math.Pow(math.Sin(t), 2))
		points[i] = Point{x, y}
	}
	return NewPath(points)
}

// EllipseSamples returns an ellipse-shaped Path.
func EllipseSamples(origin Point, xRadius, yRadius float64, samples int) *Path {
	points := make([]Point, 0)
	angleStep := 2 * math.Pi / float64(samples)
	for i := 0; i < samples; i++ {
		angle := float64(i) * angleStep
		pt := Point{xRadius * math.Cos(angle), yRadius * math.Sin(angle)}
		points = append(points, pt.Add(origin))
	}

	return NewPath(points).Close().calculateLength()
}
