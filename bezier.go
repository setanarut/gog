package gog

type CubicBezier struct {
	a, b, c, d Point
}

// GetPoint returns point at time t
func (cb *CubicBezier) GetPoint(t float64) Point {
	x := (1-t)*(1-t)*(1-t)*cb.a.X + 3*(1-t)*(1-t)*t*cb.b.X + 3*(1-t)*t*t*cb.c.X + t*t*t*cb.d.X
	y := (1-t)*(1-t)*(1-t)*cb.a.Y + 3*(1-t)*(1-t)*t*cb.b.Y + 3*(1-t)*t*t*cb.c.Y + t*t*t*cb.d.Y
	return Point{X: x, Y: y}
}
