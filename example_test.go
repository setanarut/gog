package gog_test

import (
	"fmt"
	"github/setanarut/gog"
	"image/color"
)

// 150-frame rotating cubic bezier APNG animation
func Example() {
	ctx := gog.New(250, 250)
	curve := gog.CubicBezier(100, 95, 50, 300, 190, 88, 140, 200)
	curve.SetPos(ctx.Center)
	for i := 0; i < 150; i++ {
		ctx.Clear(color.Gray{30})
		curve.Rotate((gog.Pi * 2) / 150)
		ctx.DebugDraw(curve)
		ctx.AppendAnimationFrame()
	}
	// ctx.SaveAPNG("anim.png", 3)
	fmt.Println(len(ctx.AnimationFrames))
	// Output:
	// 150
}

// Creates new line and prints start and end point
func ExampleLine() {
	line := gog.Line(gog.Point{0, 0}, gog.Point{25, 80})
	fmt.Println(line.Start(), line.End())
	// Output:
	// {0 0} {25 80}
}

// Creates new line and prints point and angle at time
func ExampleLine_pointAngleAt() {
	line := gog.Line(gog.Point{0, 0}, gog.Point{25, 80})
	p, a := line.PointAngleAt(0.5)
	fmt.Println(p, a)
	// Output:
	// {12.5 40} 1.2679114584199251
}
