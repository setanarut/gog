package main

import (
	"image/color"

	"github.com/setanarut/gog/v2"
	"github.com/setanarut/gog/v2/shapes"
	"github.com/setanarut/v"
)

func main() {
	ctx := gog.NewContext(450, 220).Clear(color.Gray{100})
	curvePath := shapes.CubicBezier(13, 17, 44, 12, 67, 23, 29, 59, 100)
	curvePath.Scale(v.Vec{1.4, 1.4})
	strokeStyle := gog.DefaultStrokeStyle()
	for i := range 5 {
		curvePath.SetPos(v.Vec{float64(i)*80 + 80, 50})
		strokeStyle.Cap = gog.CapMode(i)
		ctx.Fill(curvePath, color.RGBA{30, 144, 255, 255})
		ctx.Stroke(curvePath, strokeStyle.SetLineWidth(20).SetColor(color.Black))
		ctx.Stroke(curvePath, strokeStyle.SetLineWidth(1.5).SetColor(color.White))
	}
	curvePath.Close()
	for i := range 3 {
		curvePath.SetPos(v.Vec{float64(i)*80 + 80, 150})
		strokeStyle.Join = gog.JoinMode(i)
		ctx.Fill(curvePath, color.RGBA{30, 144, 255, 255})
		ctx.Stroke(curvePath, strokeStyle.SetLineWidth(20).SetColor(color.Black))
		ctx.Stroke(curvePath, strokeStyle.SetLineWidth(1.5).SetColor(color.White))
	}
	ctx.SavePNG("cap_join.png")
}
