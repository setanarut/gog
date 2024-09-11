package main

import (
	"fmt"

	"github.com/setanarut/gog/v2"
	"github.com/setanarut/gog/v2/shapes"
)

func main() {
	// ctx := gog.NewContext(200, 100)
	// rect := shapes.Rect(vec.Vec2{}, 30, 10)
	// lemn := shapes.Lemniscate(100, 100)
	// lemnTotalLength := lemn.Length()
	// for _, length := range utils.Linspace(0, lemnTotalLength, 120) {
	// 	ctx.Clear(color.Black)
	// 	ctx.Stroke(lemn, gog.DefaultStrokeStyle())
	// 	pos, ang := lemn.PointAngleAtLength(length)
	// 	ctx.Fill(rect.SetPos(pos).Rotated(ang), color.White)
	// 	ctx.AppendAnimationFrame()
	// }
	// ctx.SaveAPNG("point_angle.png", 2)
	drawim()
}
func drawim() {
	ctx := gog.NewContext(400, 200)
	lemn := shapes.Lemniscate(50, 173).SetPos(ctx.Center)
	fmt.Println("a", lemn.Anchor)
	fmt.Println(lemn.Centroid())
	ctx.DebugDraw(lemn)
	ctx.SavePNG("point_angle.png")
}
