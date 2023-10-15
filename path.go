package gog

import (
	"fmt"
	"image/color"
	"math"
	"slices"
)

// Path object
type Path struct {
	// points holds coordinates of Path
	points []Point
	// Style holds the fill color, line color, thickness and DrawMode options.
	Style Style
	// Anchor point
	Anchor Point
	// length holds total length of Path
	length float64
}

// DebugDraw draw for debug
func (p *Path) DebugDraw(c *context) *Path {
	c.DebugDraw(p)
	return p
}

// AppendPoints appends points to the end of the Path
func (p *Path) AppendPoints(points ...Point) *Path {
	p.points = append(p.points, points...)
	p.calculateLength()
	return p
}

// DeleteEnd removes end point of the Path if the number of points is more than two.
func (p *Path) DeleteEnd() *Path {
	if len(p.points) > 2 {
		p.points = p.points[:len(p.points)-1]
		p.calculateLength()
	}
	return p
}

// DeleteAtIndex removes point from Path at index if the number of points is more than two.
func (p *Path) DeleteAtIndex(index int) *Path {
	if len(p.points) > 2 {
		p.points = slices.Delete(p.points, index, index+1)
		p.calculateLength()
	}
	return p
}

// InsertAtIndex inserts point to the Path at index
func (p *Path) InsertAtIndex(pnt Point, index int) *Path {
	p.points = slices.Insert(p.points, index, pnt)
	p.calculateLength()
	return p
}

// PointAtIndex returns point at index
func (p *Path) PointAtIndex(index int) Point {
	pt := p.points[index]
	return pt
}

// InsertAtLength inserts point at length if coord is empty
func (p *Path) InsertAtLength(length float64) {
	length = clip(length, 0, p.length)
	traveledDist := 0.0
	for i := 0; i < len(p.points)-1; i++ {
		start, end := p.points[i], p.points[i+1]
		segmentLength := start.Distance(end)
		if traveledDist+segmentLength >= length {
			fracSeg := (length - traveledDist) / segmentLength
			pt := start.Add(end.Sub(start).Mul(Point{fracSeg, fracSeg}))
			if pt.Distance(end) > 0.1 {
				p.InsertAtIndex(pt, i+1)
			}
			p.calculateLength()
			return
		}
		traveledDist += segmentLength
	}
}

// InsertAtTime inserts point at time
func (p *Path) InsertAtTime(t float64) {
	t = clip(t, 0, 1)
	length := t * p.length
	p.InsertAtLength(length)

}

// RemoveDoubles removes double points
func (p *Path) RemoveDoubles() *Path {
	uniquePoints := make([]Point, 0)
	seen := make(map[Point]bool)
	closed := p.IsClosed()
	p.Open()
	for _, p := range p.points {
		if !seen[p] {
			seen[p] = true
			uniquePoints = append(uniquePoints, p)
		}
	}
	p.points = uniquePoints
	if closed {
		p.Close()
	}
	return p
}

// Clone returns copy of path
func (p *Path) Clone() *Path {
	return deepCopyPath(p)
}

// Len returns number of points
func (p Path) Len() int {
	return len(p.points)
}

// ResetAnchor sets anchor point to centroid
func (p *Path) ResetAnchor() *Path {
	p.Anchor = p.Centroid()
	return p
}

// Fill fills the path and draws it on the canvas.
func (p *Path) Fill(c *context) *Path {
	c.Fill(p)
	return p
}

// Stroke Draw the path as a stroke
func (p *Path) Stroke(c *context) *Path {
	c.Stroke(p)
	return p
}

// FillStroke fills then strokes path
func (p *Path) FillStroke(c *context) *Path {
	c.Fill(p)
	c.Stroke(p)
	return p
}

// StrokeFill stroke then fill path
func (p *Path) StrokeFill(c *context) *Path {
	c.Stroke(p)
	c.Fill(p)
	return p
}

// calculateLenght calculates total length of path
func (p *Path) calculateLength() *Path {
	p.length = 0.0
	for i := 0; i < len(p.points)-1; i++ {
		p.length += p.points[i].Distance(p.points[i+1])
	}
	return p
}

// Open opens Path
func (p *Path) Open() *Path {
	if p.IsClosed() {
		p.points = p.points[:len(p.points)-1]
		return p
	}
	return p
}

// Close closes Path
func (p *Path) Close() *Path {
	if !p.IsClosed() {
		p.points = append(p.points, p.points[0])
		p.calculateLength()
	}
	return p
}

// Returns Perpendicular line points at time t with length
func (p *Path) Perpendicular(t float64, length float64) (p1 Point, p2 Point) {
	pos, ang := p.PointAngleAtTime(t)
	ang += math.Pi * 0.5
	p1 = pointOnCircle(pos, length/2, ang)
	p2 = pointOnCircle(pos, length/2, oppositeAngle(ang))
	return p1, p2
}

// Centroid Calculates and returns the path's centroid point
func (p *Path) Centroid() Point {
	total := float64(len(p.points))
	centroidPoint := Point{0, 0}
	noDoublePath := p.Clone().RemoveDoubles()
	for _, pt := range noDoublePath.points {
		centroidPoint = centroidPoint.Add(pt)
	}
	noDoublePath = nil
	return centroidPoint.Div(total)
}

// SetAnchor Sets Path's anchor point
func (p *Path) SetAnchor(pt Point) *Path {
	p.Anchor = pt
	return p
}

// PointAngleAtLength Returns point and tangent angle at length
func (p *Path) PointAngleAtLength(length float64) (Point, float64) {
	length = clip(length, 0, p.length)
	traveledDist := 0.0
	for i := 0; i < len(p.points)-1; i++ {
		start, end := p.points[i], p.points[i+1]
		segmentLength := start.Distance(end)
		if traveledDist+segmentLength >= length {
			fracSeg := (length - traveledDist) / segmentLength
			pt := start.Add(end.Sub(start).Mul(Point{fracSeg, fracSeg}))
			return pt, TangentAngle(start, end)
		}
		traveledDist += segmentLength
	}
	return Point{}, 0.0
}

// PointAngleAtTime Returns point and tangent angle at time t
func (p *Path) PointAngleAtTime(t float64) (Point, float64) {
	t = clip(t, 0, 1)
	length := t * p.length
	return p.PointAngleAtLength(length)
}

// IsClosed returns true if Path closed
func (p *Path) IsClosed() bool {
	if p.Start().Distance(p.End()) < 0.01 {
		return true
	} else {
		return false
	}
}

// PrintPoints prints path points to standard output.
func (p *Path) PrintPoints() {
	fmt.Println(p.points)
}

// GetPoints return points
func (p *Path) GetPoints() []Point {
	return p.points
}

// Bounds returns bounds min/max
func (p *Path) Bounds() (Point, Point) {
	min := p.Start()
	max := p.Start()
	for i := 0; i < p.Len(); i++ {
		if p.points[i].X < min.X {
			min.X = p.points[i].X
		}
		if p.points[i].Y < min.Y {
			min.Y = p.points[i].Y
		}
		if p.points[i].X > max.X {
			max.X = p.points[i].X
		}
		if p.points[i].Y > max.Y {
			max.Y = p.points[i].Y
		}
	}
	return min, max
}

// Start returns start point of Path
func (p *Path) Start() Point {
	if len(p.points) > 0 {
		return p.points[0]
	}
	return Point{0, 0}
}

// End returns end point of Path
func (p *Path) End() Point {
	if len(p.points) > 0 {
		return p.points[len(p.points)-1]
	}
	return Point{0, 0}
}

// SetStyle sets Style of Path.
func (p *Path) SetStyle(s Style) *Path {
	p.Style = s
	return p
}

// SetFill sets fill color of Path.
func (p *Path) SetFill(c color.Color) *Path {
	p.Style.Fill = c
	return p
}

// SetStroke sets stroke color of Path.
func (p *Path) SetStroke(c color.Color) *Path {
	p.Style.Stroke = c
	return p
}

// SetLineWidth sets line thickness of Path.
func (p *Path) SetLineWidth(w float64) *Path {
	p.Style.LineWidth = w
	return p
}

// SetPos Aligns the Path with the anchor point to the desired point.
// In other words, it sets the position.
func (p *Path) SetPos(position Point) *Path {
	p.Translate(position.X-p.Anchor.X, position.Y-p.Anchor.Y)
	return p
}

// Translate translates the Path
func (p *Path) Translate(x, y float64) *Path {
	q := Point{x, y}
	for i := 0; i < len(p.points); i++ {
		p.points[i] = p.points[i].Add(q)
	}
	p.Anchor = p.Anchor.Add(q)
	return p
}

// Rotate rotates the Path about
func (p *Path) Rotate(angle float64) *Path {
	for i := 0; i < p.Len(); i++ {
		p.points[i] = p.points[i].Rotate(angle, p.Anchor)
	}
	return p
}

// Scale scales the Path at the Anchor point.
func (p *Path) Scale(factor Point) *Path {
	for i := 0; i < len(p.points); i++ {
		p.points[i] = factor.Mul(p.points[i].Sub(p.Anchor)).Add(p.Anchor)
	}
	p.calculateLength()
	return p
}

// Length returns total length of Path
func (p Path) Length() float64 {
	p.calculateLength()
	return p.length
}

// Reverse reverses Path.
// The starting point becomes the end and the end becomes the beginning.
func (p *Path) Reverse() *Path {
	slices.Reverse(p.points)
	return p
}
