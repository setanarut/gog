# gog

[gog](https://pkg.go.dev/github.com/setanarut/gog#section-documentation) is a Go Object-oriented Graphic drawing library for generative art, PNG or ANPG animations

Instead of immediate drawing, each shape is a `Path{}` struct made up of points. The Path has `Fill()` `Stroke()` `FillStroke()` `StrokeFill()` functions for drawing on the canvas. There is also a `DrawDebug()` function for Debug purposes that draws all the properties of the Path.

All transformations are made with reference to the Path.Anchor point.

- `Path.Translate()`
- `Path.SetPos()`
- `Path.Rotate()`
- `Path.Rotated()`
- `Path.Scale()`

```Go
package main

import (
	"github.com/setanarut/gog"
	"image/color"
)

func main() {
	c := gog.New(250, 250)
	curve := gog.CubicBezier(100, 95, 50, 300, 190, 88, 140, 200)
	curve.SetPos(c.Center).Scale(gog.P(1.3, 1.3))
	for i := 0; i < 150; i++ {
		c.Clear(color.Gray{30})
		curve.Rotate((gog.Pi * 2) / 150)
		c.DebugDraw(curve)
		c.AppendAnimationFrame()
	}
	c.SaveAPNG("curve_anim.png", 3)
}
```

![curve](./examples/curve_anim/curve_anim.png)

## Motion path example

```go
rect := gog.Rect(gog.Point{}, 30, 10)
poly := gog.Circle(c.Center, 50).Scale(gog.P(1, 0.5))
for _, t := range gog.Linspace(0, 1, 150) {
	c.Clear(gog.Black)
	poly.Stroke(c)
	p, a := poly.PointAngleAtTime(t)
	rect.SetPos(p).Rotated(a).SetFill(colornames.Yellow).Fill(c)
	c.AppendAnimationFrame()
}
```

[Full code](./examples/point_angle/point_angle.go) is available in example folder

![curve](./examples/point_angle/point_angle.png)
