package main

import (
	"image/color"

	"github.com/setanarut/gog"
	"golang.org/x/image/colornames"
)

func main() {
	c := gog.New(450, 220).Clear(color.Gray{100})
	curve := gog.CubicBezier(13, 17, 44, 12, 67, 23, 29, 59, 100).Scale(gog.P(1.4, 1.4))
	for i := 0; i < 5; i++ {
		curve.SetPos(gog.P(float64(i)*80+80, 50))
		curve.Style.Cap = gog.CapMode(i)
		curve.SetFill(colornames.Dodgerblue).Fill(c)
		curve.SetStroke(colornames.Black).SetLineWidth(20).Stroke(c)
		curve.SetStroke(colornames.White).SetLineWidth(1.5).Stroke(c)
	}
	curve.Close()
	for i := 0; i < 3; i++ {
		curve.SetPos(gog.P(float64(i)*80+80, 150))
		curve.Style.Join = gog.JoinMode(i)
		curve.SetFill(colornames.Dodgerblue).Fill(c)
		curve.SetStroke(colornames.Black).SetLineWidth(20).Stroke(c)
		curve.SetStroke(colornames.White).SetLineWidth(1.5).Stroke(c)
	}
	c.SavePNG("cap_join.png")
}
