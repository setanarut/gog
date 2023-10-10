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
	newPath.SetAnchor(newPath.GetCentroid())
	newPath.calculateLength()
	return newPath
}

// BBox returns a bounding box path with min and max points.
func BBox(min, max Point) *Path {
	points := []Point{min, {max.X, min.Y}, max, {min.X, max.Y}, min}
	return NewPath(points)
}

// Line returns a line Path.
func Line(start, end Point) *Path {
	return NewPath([]Point{start, end})
}

// Square returns a square Path with side s.
func Square(topLeft Point, s float64) *Path {
	Sq := NewPath([]Point{{}, {s, 0}, {s, s}, {0, s}})
	Sq.Close().Translate(topLeft.X, topLeft.Y)
	return Sq
}

// Ellipse returns an ellipse-shaped Path.
func Ellipse(origin Point, xRadius, yRadius float64) *Path {
	samples := int(clip(xRadius, 20, 80))
	return ellipseSamples(origin, xRadius, yRadius, samples)
}

// Circle returns an circle-shaped Path.
func Circle(origin Point, radius float64) *Path {
	samples := int(clip(radius, 20, 80))
	return ellipseSamples(origin, radius, radius, samples)
}

// RegularPolygon returns an circle-shaped Path.
func RegularPolygon(origin Point, n int, radius float64) *Path {
	align_angle := (math.Pi / 2) - (math.Pi*2)/float64(n)/2
	return ellipseSamples(origin, radius, radius, n).Rotate(align_angle)
}

// ellipseSamples returns an ellipseSamples-shaped Path.
func ellipseSamples(origin Point, xRadius, yRadius float64, samples int) *Path {
	points := make([]Point, 0)
	angleStep := 2 * math.Pi / float64(samples)
	for i := 0; i < samples; i++ {
		angle := float64(i) * angleStep
		pt := Point{xRadius * math.Cos(angle), yRadius * math.Sin(angle)}
		points = append(points, pt.Add(origin))
	}
	return NewPath(points).Close()
}
