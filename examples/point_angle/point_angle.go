package main

import (
	"github.com/setanarut/gog"
	"golang.org/x/image/colornames"
)

func main() {
	c := gog.New(200, 100)
	rect := gog.Rect(gog.Point{}, 30, 10)
	ellipse := gog.Circle(c.Center, 50).Scale(gog.P(1, 0.5))
	for _, t := range gog.Linspace(0, 1, 150) {
		c.Clear(gog.Black)
		ellipse.Stroke(c)
		p, a := ellipse.PointAngleAtTime(t)
		rect.SetPos(p).Rotated(a).SetFill(colornames.Yellow).Fill(c)
		c.AppendAnimationFrame()
	}
	c.SaveAPNG("point_angle.png", 2)
}
