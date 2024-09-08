package gog

import (
	"github.com/setanarut/gog/v2/utils"
	"github.com/setanarut/gog/v2/vec"
)

type cubicBezier struct {
	pts [4]vec.Vec2
}

// getPoint returns point at time t
func (cb *cubicBezier) getPoint(t float64) vec.Vec2 {
	x := (1-t)*(1-t)*(1-t)*cb.pts[0].X + 3*(1-t)*(1-t)*t*cb.pts[1].X + 3*(1-t)*t*t*cb.pts[2].X + t*t*t*cb.pts[3].X
	y := (1-t)*(1-t)*(1-t)*cb.pts[0].Y + 3*(1-t)*(1-t)*t*cb.pts[1].Y + 3*(1-t)*t*t*cb.pts[2].Y + t*t*t*cb.pts[3].Y
	return vec.Vec2{X: x, Y: y}
}

// flatten returns flattan Bezier
func (cb *cubicBezier) flatten(samples int) []vec.Vec2 {

	coords := make([]vec.Vec2, 0)
	for _, t := range utils.Linspace(0, 1, samples) {
		coords = append(coords, cb.getPoint(t))
	}
	return coords
}
