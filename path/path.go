package path

import (
	"fmt"
	"math"
	"slices"

	"github.com/setanarut/gog/v2/utils"
	"github.com/setanarut/gog/v2/vec"
)

// Path object
type Path struct {
	// Anchor point
	Anchor vec.Vec2

	// points holds coordinates of Path
	points []vec.Vec2
	// length holds total length of Path
	length float64
}

// NewPath returns new Path from points
func NewPath(points []vec.Vec2) *Path {
	newPath := &Path{
		points: points,
	}

	if len(points) > 1 {
		// if newPath.IsClosed() == false {
		// 	newPath.Close()
		// }
		newPath.SetAnchor(newPath.Centroid())
	}
	newPath.CalculateLength()
	return newPath
}

// // DebugDraw draw for debug
// func (p *Path) DebugDraw(c *context) *Path {
// 	c.DebugDraw(p)
// 	return p
// }

// AppendPoints appends points to the end of the Path
func (p *Path) AppendPoints(points ...vec.Vec2) *Path {
	p.points = append(p.points, points...)
	p.CalculateLength()
	return p
}

// DeleteEnd removes end point of the Path if the number of points is more than two.
func (p *Path) DeleteEnd() *Path {
	if len(p.points) > 2 {
		p.points = p.points[:len(p.points)-1]
		p.CalculateLength()
	}
	return p
}

// DeleteAtIndex removes point from Path at index if the number of points is more than two.
func (p *Path) DeleteAtIndex(index int) *Path {
	if len(p.points) > 2 {
		p.points = slices.Delete(p.points, index, index+1)
		p.CalculateLength()
	}
	return p
}

// InsertAtIndex inserts point to the Path at index
func (p *Path) InsertAtIndex(pnt vec.Vec2, index int) *Path {
	p.points = slices.Insert(p.points, index, pnt)
	p.CalculateLength()
	return p
}

// PointAtIndex returns point at index
func (p *Path) PointAtIndex(index int) vec.Vec2 {
	pt := p.points[index]
	return pt
}

// InsertAtLength inserts point at length if coord is empty
func (p *Path) InsertAtLength(length float64) {
	length = utils.Clip(length, 0, p.length)
	traveledDist := 0.0
	for i := 0; i < len(p.points)-1; i++ {
		start, end := p.points[i], p.points[i+1]
		segmentLength := start.Distance(end)
		if traveledDist+segmentLength >= length {
			fracSeg := (length - traveledDist) / segmentLength
			pt := start.Add(end.Sub(start).Mult(vec.Vec2{fracSeg, fracSeg}))
			if pt.Distance(end) > 0.1 {
				p.InsertAtIndex(pt, i+1)
			}
			p.CalculateLength()
			return
		}
		traveledDist += segmentLength
	}
}

// InsertAtTime inserts point at time
func (p *Path) InsertAtTime(t float64) {
	t = utils.Clip(t, 0, 1)
	length := t * p.length
	p.InsertAtLength(length)

}

// RemoveDoubles removes double points
func (p *Path) RemoveDoubles() *Path {
	uniquePoints := make([]vec.Vec2, 0)
	seen := make(map[vec.Vec2]bool)
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

// Len returns number of points
func (p Path) Len() int {
	return len(p.points)
}

// ResetAnchor sets anchor point to centroid
func (p *Path) ResetAnchor() *Path {
	p.Anchor = p.Centroid()
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
		p.CalculateLength()
	}
	return p
}

// Returns Perpendicular line points at time t with length
func (p *Path) Perpendicular(t float64, length float64) (p1 vec.Vec2, p2 vec.Vec2) {
	pos, ang := p.PointAngleAtTime(t)
	ang += math.Pi * 0.5
	p1 = utils.PointOnCircle(pos, length/2, ang)
	p2 = utils.PointOnCircle(pos, length/2, utils.OppositeAngle(ang))
	return p1, p2
}

// Centroid Calculates and returns the path's centroid point
func (p *Path) Centroid() vec.Vec2 {
	total := float64(len(p.points))
	centroidPoint := vec.Vec2{0, 0}
	noDoublePath := clonePath(p).RemoveDoubles()
	for _, pt := range noDoublePath.points {
		centroidPoint = centroidPoint.Add(pt)
	}
	noDoublePath = nil
	return centroidPoint.Div(total)
}

// SetAnchor Sets Path's anchor point
func (p *Path) SetAnchor(pt vec.Vec2) *Path {
	p.Anchor = pt
	return p
}

// PointAngleAtLength Returns point and tangent angle at length
func (p *Path) PointAngleAtLength(length float64) (vec.Vec2, float64) {
	length = utils.Clip(length, 0, p.length)
	traveledDist := 0.0
	for i := 0; i < len(p.points)-1; i++ {
		start, end := p.points[i], p.points[i+1]
		segmentLength := start.Distance(end)
		if traveledDist+segmentLength >= length {
			fracSeg := (length - traveledDist) / segmentLength
			pt := start.Add(end.Sub(start).Mult(vec.Vec2{fracSeg, fracSeg}))
			return pt, utils.TangentAngle(start, end)
		}
		traveledDist += segmentLength
	}
	return vec.Vec2{}, 0.0
}

// PointAngleAtTime Returns point and tangent angle at time t
func (p *Path) PointAngleAtTime(t float64) (vec.Vec2, float64) {
	t = utils.Clip(t, 0, 1)
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

// Points return points
func (p *Path) Points() []vec.Vec2 {
	return p.points
}

// SetPoints sets points
func (p *Path) SetPoints(pts []vec.Vec2) {
	p.points = pts
}

// Bounds returns bounds min/max
func (p *Path) Bounds() (vec.Vec2, vec.Vec2) {
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
func (p *Path) Start() vec.Vec2 {
	if len(p.points) > 0 {
		return p.points[0]
	}
	return vec.Vec2{0, 0}
}

// End returns end point of Path
func (p *Path) End() vec.Vec2 {
	if len(p.points) > 0 {
		return p.points[len(p.points)-1]
	}
	return vec.Vec2{0, 0}
}

// SetPos Aligns the Path with the anchor point to the desired point.
// In other words, it sets the position.
func (p *Path) SetPos(position vec.Vec2) *Path {
	p.Translate(position.X-p.Anchor.X, position.Y-p.Anchor.Y)
	return p
}

// Translate translates the Path
func (p *Path) Translate(x, y float64) *Path {
	q := vec.Vec2{x, y}
	for i := 0; i < len(p.points); i++ {
		p.points[i] = p.points[i].Add(q)
	}
	p.Anchor = p.Anchor.Add(q)
	return p
}

// Rotate rotates the Path about Path.Anchor point
func (p *Path) Rotate(angle float64) *Path {
	for i := 0; i < p.Len(); i++ {
		p.points[i] = vec.RotateAbout(p.points[i], angle, p.Anchor)
	}
	return p
}

// Rotated returns new rotated Path about Path.Anchor point
func (p *Path) Rotated(angle float64) *Path {
	return clonePath(p).Rotate(angle)
}

// Scale scales the Path at the Anchor point.
func (p *Path) Scale(factor vec.Vec2) *Path {
	for i := 0; i < len(p.points); i++ {
		p.points[i] = factor.Mult(p.points[i].Sub(p.Anchor)).Add(p.Anchor)
	}
	p.CalculateLength()
	return p
}

// CalculateLength calculates total length of path
func (p *Path) CalculateLength() *Path {
	p.length = 0.0
	for i := 0; i < len(p.points)-1; i++ {
		p.length += p.points[i].Distance(p.points[i+1])
	}
	return p
}

// Length returns total length of Path
func (p Path) Length() float64 {
	return p.length
}

// Reverse reverses Path.
// The starting point becomes the end and the end becomes the beginning.
func (p *Path) Reverse() *Path {
	slices.Reverse(p.points)
	return p
}

// DeepCopyPath returns copy of path
func clonePath(p *Path) *Path {
	newPath := new(Path)
	newCoords := make([]vec.Vec2, len(p.Points()))
	copy(newCoords, p.Points())
	newPath.SetPoints(newCoords)
	newPath.Anchor = p.Anchor
	newPath.length = p.length
	return newPath
}