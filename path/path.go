package path

import (
	"fmt"
	"math"
	"slices"

	"github.com/setanarut/gog/v2/utils"
	"github.com/setanarut/v"
)

// Path object
type Path struct {
	// Anchor point
	Anchor v.Vec
	// Points holds coordinates of Path
	Points []v.Vec
}

// NewPath returns new Path from points
func NewPath(points []v.Vec) *Path {
	newPath := &Path{
		Points: points,
	}
	if len(points) > 1 {
		newPath.SetAnchorToCentroid()
	}
	return newPath
}

// AppendPoints appends points to the end of the Path
//
// This is an operation that changes the path length.
// If the length is necessary, length should be taken with Length().
func (p *Path) AppendPoints(points ...v.Vec) *Path {
	p.Points = append(p.Points, points...)
	return p
}

// DeleteEnd removes end point of the Path if the number of points is more than two.
//
// This is an operation that changes the path length.
// If the length is necessary,length should be taken with Length().
func (p *Path) DeleteEnd() *Path {
	if len(p.Points) > 2 {
		p.Points = p.Points[:len(p.Points)-1]
	}
	return p
}

// DeleteAtIndex removes point from Path at index if the number of points is more than two.
//
// This is an operation that changes the path length.
// If the length is necessary, length should be taken with Length().
func (p *Path) DeleteAtIndex(index int) *Path {
	if len(p.Points) > 2 {
		p.Points = slices.Delete(p.Points, index, index+1)
	}
	return p
}

// InsertAtIndex inserts point to the Path at index
//
// This is an operation that changes the path length.
// If the length is necessary, length should be taken with Length().
func (p *Path) InsertAtIndex(pnt v.Vec, index int) *Path {
	p.Points = slices.Insert(p.Points, index, pnt)
	return p
}

// PointAtIndex returns point at index
func (p *Path) PointAtIndex(index int) v.Vec {
	pt := p.Points[index]
	return pt
}

// InsertAtLength inserts point at length if coord is empty
//
// This is an operation that changes the path length.
// If the length is necessary, length should be taken with Length().
func (p *Path) InsertAtLength(length float64) {
	traveledDist := 0.0
	for i := range len(p.Points) - 1 {
		start, end := p.Points[i], p.Points[i+1]
		segmentLength := start.Dist(end)
		if traveledDist+segmentLength >= length {
			fracSeg := (length - traveledDist) / segmentLength
			pt := start.Add(end.Sub(start).Mul(v.Vec{fracSeg, fracSeg}))
			if pt.Dist(end) > 0.1 {
				p.InsertAtIndex(pt, i+1)
			}
			return
		}
		traveledDist += segmentLength
	}
}

// RemoveDoubles removes double points
func (p *Path) RemoveDoubles() *Path {
	uniquePoints := make([]v.Vec, 0)
	seen := make(map[v.Vec]bool)
	closed := p.IsClosed()
	p.Open()
	for _, p := range p.Points {
		if !seen[p] {
			seen[p] = true
			uniquePoints = append(uniquePoints, p)
		}
	}
	p.Points = uniquePoints
	if closed {
		p.Close()
	}
	return p
}

// Len returns number of points
func (p Path) Len() int {
	return len(p.Points)
}

// ResetAnchor sets anchor point to centroid
func (p *Path) ResetAnchor() *Path {
	p.Anchor = p.Centroid()
	return p
}

// Open opens Path
func (p *Path) Open() *Path {
	if p.IsClosed() {
		p.Points = p.Points[:len(p.Points)-1]
		return p
	}
	return p
}

// Close closes Path
//
// This is an operation that changes the path length.
// If the length is necessary, length should be taken with Length().
func (p *Path) Close() *Path {
	if !p.IsClosed() {
		p.Points = append(p.Points, p.Points[0])
	}
	return p
}

// Returns Perpendicular line points at length
func (p *Path) Perpendicular(length float64, lineLength float64) (start v.Vec, end v.Vec) {
	pos, ang := p.PointAngleAtLength(length)
	ang += math.Pi * 0.5
	start = utils.PointOnCircle(pos, length/2, ang)
	end = utils.PointOnCircle(pos, length/2, utils.OppositeAngle(ang))
	return start, end
}

// Centroid calculates and returns the path's centroid point.
// Costly operation. Don't use unless necessary.
func (p *Path) Centroid() v.Vec {
	total := float64(len(p.Points))
	centroidPoint := v.Vec{0, 0}
	clonePath := p.Clone()
	if clonePath.IsClosed() {
		clonePath.Open()
	}
	clonePath.RemoveDoubles()
	for _, pt := range clonePath.Points {
		centroidPoint = centroidPoint.Add(pt)
	}
	clonePath = nil
	return centroidPoint.DivS(total)
}

// SetAnchor Sets Path's anchor point
func (p *Path) SetAnchor(pt v.Vec) *Path {
	p.Anchor = pt
	return p
}

// SetAnchorToCentroid Sets Path's anchor point to centroid
func (p *Path) SetAnchorToCentroid() *Path {
	return p.SetAnchor(p.Centroid())
}

// PointAngleAtLength Returns point and tangent angle at length
func (p *Path) PointAngleAtLength(length float64) (v.Vec, float64) {
	traveledDist := 0.0
	for i := range len(p.Points) - 1 {
		start, end := p.Points[i], p.Points[i+1]
		segmentLength := start.Dist(end)
		if traveledDist+segmentLength >= length {
			fracSeg := (length - traveledDist) / segmentLength
			pt := start.Add(end.Sub(start).Mul(v.Vec{fracSeg, fracSeg}))
			return pt, utils.TangentAngle(start, end)
		}
		traveledDist += segmentLength
	}
	return v.Vec{}, 0.0
}

// IsClosed returns true if Path closed
func (p *Path) IsClosed() bool {
	if p.Start().Dist(p.End()) < 0.1 {
		return true
	} else {
		return false
	}
}

// PrintPoints prints path points to standard output.
func (p *Path) PrintPoints() {
	fmt.Println(p.Points)
}

// SetPoints sets points
func (p *Path) SetPoints(pts []v.Vec) {
	p.Points = pts
}

// Bounds returns bounds min/max
func (p *Path) Bounds() (v.Vec, v.Vec) {
	min := p.Start()
	max := p.Start()
	for i := 0; i < p.Len(); i++ {
		if p.Points[i].X < min.X {
			min.X = p.Points[i].X
		}
		if p.Points[i].Y < min.Y {
			min.Y = p.Points[i].Y
		}
		if p.Points[i].X > max.X {
			max.X = p.Points[i].X
		}
		if p.Points[i].Y > max.Y {
			max.Y = p.Points[i].Y
		}
	}
	return min, max
}

// Start returns start point of Path
func (p *Path) Start() v.Vec {
	if len(p.Points) > 0 {
		return p.Points[0]
	}
	return v.Vec{0, 0}
}

// End returns end point of Path
func (p *Path) End() v.Vec {
	if len(p.Points) > 0 {
		return p.Points[len(p.Points)-1]
	}
	return v.Vec{0, 0}
}

// SetPos Aligns the Path with the anchor point to the desired point.
// In other words, it sets the position.
func (p *Path) SetPos(position v.Vec) *Path {
	p.Translate(position.X-p.Anchor.X, position.Y-p.Anchor.Y)
	return p
}

// Translate translates the Path
func (p *Path) Translate(x, y float64) *Path {
	q := v.Vec{x, y}
	for i := 0; i < len(p.Points); i++ {
		p.Points[i] = p.Points[i].Add(q)
	}
	p.Anchor = p.Anchor.Add(q)
	return p
}

// Rotate rotates the Path about Path.Anchor point
func (p *Path) Rotate(angle float64) *Path {
	for i := 0; i < p.Len(); i++ {
		p.Points[i] = utils.RotateAbout(p.Points[i], angle, p.Anchor)
	}
	return p
}

// Rotated returns new rotated Path about Path.Anchor point
func (p *Path) Rotated(angle float64) *Path {
	return p.Clone().Rotate(angle)
}

// Scale scales the Path at the Anchor point.
// CalculateLength() after scaling for
func (p *Path) Scale(factor v.Vec) *Path {
	for i := 0; i < len(p.Points); i++ {
		p.Points[i] = factor.Mul(p.Points[i].Sub(p.Anchor)).Add(p.Anchor)
	}
	return p
}

// Length calculates and returns total length of path
// Costly operation. Don't use unless necessary.
func (p *Path) Length() float64 {
	length := 0.0
	for i := 0; i < len(p.Points)-1; i++ {
		length += p.Points[i].Dist(p.Points[i+1])
	}
	return length
}

// Reverse reverses Path.
// The starting point becomes the end and the end becomes the beginning.
func (p *Path) Reverse() *Path {
	slices.Reverse(p.Points)
	return p
}

// DeepCopyPath returns copy of path
func (p *Path) Clone() *Path {
	newPath := new(Path)
	newCoords := make([]v.Vec, len(p.Points))
	copy(newCoords, p.Points)
	newPath.SetPoints(newCoords)
	newPath.Anchor = p.Anchor
	return newPath
}
