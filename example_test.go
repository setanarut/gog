package gog_test

import (
	"fmt"
	"image/color"
	"math"

	"github.com/setanarut/gog/v2"
	"github.com/setanarut/gog/v2/path"
	"github.com/setanarut/gog/v2/shapes"
	"github.com/setanarut/vec"
)

// 150-frame rotating cubic bezier APNG animation
func Example() {
	ctx := gog.NewContext(250, 250)
	curve := shapes.CubicBezier(100, 95, 50, 300, 190, 88, 140, 200, 50)
	curve.SetPos(ctx.Center)
	for i := 0; i < 150; i++ {
		ctx.Clear(color.Gray{30})
		curve.Rotate((math.Pi * 2) / 150)
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
	line := shapes.Line(vec.Vec2{X: 0, Y: 0}, vec.Vec2{X: 25, Y: 80})
	fmt.Println(line.Start(), line.End())
	// Output:
	// Vec2{X: 0.000000, Y: 0.000000} Vec2{X: 25.000000, Y: 80.000000}
}

// Get point and tangent angle at length
func ExamplePath_PointAngleAtLength() {
	line := path.NewPath([]vec.Vec2{{X: 0, Y: 0}, {X: 10, Y: 10}})
	point, angle := line.PointAngleAtLength(line.Length() / 2)
	fmt.Println(point, angle)
	// Output:
	// Vec2{X: 5.000000, Y: 5.000000} 0.7853981633974483
}

func ExamplePath_InsertAtLength() {
	line := path.NewPath([]vec.Vec2{{0, 0}, {0, 10}, {0, 20}})
	line.InsertAtLength(10.5)
	line.PrintPoints()
	// Output:
	// [Vec2{X: 0.000000, Y: 0.000000} Vec2{X: 0.000000, Y: 10.000000} Vec2{X: 0.000000, Y: 10.500000} Vec2{X: 0.000000, Y: 20.000000}]
}

// Insert point to path points at index
func ExamplePath_InsertAtIndex() {
	line := path.NewPath([]vec.Vec2{{0, 0}, {10, 10}})
	line.InsertAtIndex(vec.Vec2{66, 66}, 1)
	line.PrintPoints()
	// Output:
	// [Vec2{X: 0.000000, Y: 0.000000} Vec2{X: 66.000000, Y: 66.000000} Vec2{X: 10.000000, Y: 10.000000}]
}

func ExamplePath_SetAnchor() {
	line := path.NewPath([]vec.Vec2{{X: 0, Y: 0}, {X: 10, Y: 10}})
	fmt.Println(line.Anchor) // Centroid of Path
	line.SetAnchor(vec.Vec2{X: 3, Y: 3})
	fmt.Println(line.Anchor)
	line.ResetAnchor()
	fmt.Println(line.Anchor)
	fmt.Println(line.Centroid() == line.Anchor)
	// Output:
	// Vec2{X: 5.000000, Y: 5.000000}
	// Vec2{X: 3.000000, Y: 3.000000}
	// Vec2{X: 5.000000, Y: 5.000000}
	// true
}

func ExamplePath_RemoveDoubles() {
	path := path.NewPath([]vec.Vec2{{0, 0}, {77, 77}, {77, 77}, {0, 0}, {0, 0}})
	path.RemoveDoubles()
	fmt.Println(path.Points())
	// Output:
	// [Vec2{X: 0.000000, Y: 0.000000} Vec2{X: 77.000000, Y: 77.000000} Vec2{X: 0.000000, Y: 0.000000}]
}
