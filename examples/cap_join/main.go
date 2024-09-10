package main

import (
	"image/color"

	"github.com/setanarut/gog/v2"
	"github.com/setanarut/gog/v2/vec"
)

func main() {
	c := gog.NewContext(450, 220).Clear(color.Gray{100})
	curvePath := gog.CubicBezier(13, 17, 44, 12, 67, 23, 29, 59, 100)
	curvePath.Scale(vec.Vec2{1.4, 1.4})
	strokeStyle := gog.DefaultStrokeStyle()
	for i := 0; i < 5; i++ {
		curvePath.SetPos(vec.Vec2{float64(i)*80 + 80, 50})
		strokeStyle.Cap = gog.CapMode(i)
		c.Fill(curvePath, color.RGBA{30, 144, 255, 255})
		c.Stroke(curvePath, strokeStyle.SetLineWidth(20).SetColor(color.Black))
		c.Stroke(curvePath, strokeStyle.SetLineWidth(1.5).SetColor(color.White))
	}
	curvePath.Close()
	for i := 0; i < 3; i++ {
		curvePath.SetPos(vec.Vec2{float64(i)*80 + 80, 150})
		strokeStyle.Join = gog.JoinMode(i)
		c.Fill(curvePath, color.RGBA{30, 144, 255, 255})
		c.Stroke(curvePath, strokeStyle.SetLineWidth(20).SetColor(color.Black))
		c.Stroke(curvePath, strokeStyle.SetLineWidth(1.5).SetColor(color.White))
	}
	c.SavePNG("cap_join.png")
}
