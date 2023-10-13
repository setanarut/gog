# gog

gog is a Go Object-oriented Graphic drawing library for generative art, GIF/ANPG animations

Instead of instant drawing mode, each shape is a `Path` object made up of points. The path has `Fill()` `Stroke()` `FillStroke()` `StrokeFill()` functions for drawing on the canvas. Transformations (`Path.Rotate()`, `Path.Scale()`, `Path.Translate()`) are performed by reference to the `Path.Anchor` point.


```Go
package main

import (
	"github/setanarut/gog"
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
