package main

import (
	"image/color"
	"math"

	"github.com/setanarut/gog/v2"
	"github.com/setanarut/gog/v2/shapes"
	"github.com/setanarut/v"
)

func main() {
	ctx := gog.NewContext(250, 250)
	bezierPath := shapes.CubicBezier(100, 95, 50, 300, 190, 88, 140, 200, 50)
	bezierPath.SetPos(ctx.Center).Scale(v.Vec{1.3, 1.3})
	for i := 0; i < 150; i++ {
		ctx.Clear(color.Gray{30})
		bezierPath.Rotate((math.Pi * 2) / 150)
		ctx.DebugDraw(bezierPath)
		ctx.AppendAnimationFrame()
	}
	ctx.SaveAPNG("bezier_anim.png", 3)
}
