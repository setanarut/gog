package main

import (
	"image/color"

	"github.com/setanarut/gog/v2"
	"github.com/setanarut/gog/v2/shapes"
	"github.com/setanarut/gog/v2/utils"
	"github.com/setanarut/v"
)

func main() {
	ctx := gog.NewContext(250, 100)
	rect := shapes.Rect(v.Vec{}, 30, 10)
	lemn := shapes.Lemniscate(100, 100).SetPos(ctx.Center)
	lemnTotalLength := lemn.Length()
	for _, length := range utils.Linspace(0, lemnTotalLength, 120) {
		ctx.Clear(color.Black)
		ctx.Stroke(lemn, gog.DefaultStrokeStyle())
		pos, ang := lemn.PointAngleAtLength(length)
		ctx.Fill(rect.SetPos(pos).Rotated(ang), color.White)
		ctx.AppendAnimationFrame()
	}
	ctx.SaveAPNG("path_follow.png", 2)
}
