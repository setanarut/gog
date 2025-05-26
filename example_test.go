package gog_test

import (
	"fmt"
	"image/color"
	"math"

	"github.com/setanarut/gog/v2"
	"github.com/setanarut/gog/v2/path"
	"github.com/setanarut/gog/v2/shapes"
	"github.com/setanarut/v"
)

// 150-frame rotating cubic bezier APNG animation
func Example() {
	ctx := gog.NewContext(250, 250)
	curve := shapes.CubicBezier(100, 95, 50, 300, 190, 88, 140, 200, 50)
	curve.SetPos(ctx.Center)
	for range 150 {
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
	line := shapes.Line(v.Vec{X: 0, Y: 0}, v.Vec{X: 25, Y: 80})
	fmt.Println(line.Start(), line.End())
	// Output:
	// (0.0, 0.0) (25.0, 80.0)
}

// Get point and tangent angle at length
func ExamplePath_PointAngleAtLength() {
	line := path.NewPath([]v.Vec{{X: 0, Y: 0}, {X: 10, Y: 10}})
	point, angle := line.PointAngleAtLength(line.Length() / 2)
	fmt.Println(point, angle)
	// Output:
	// (5.0, 5.0) 0.7853981633974483
}

func ExamplePath_InsertAtLength() {
	line := path.NewPath([]v.Vec{{0, 0}, {0, 10}, {0, 20}})
	line.InsertAtLength(10.5)
	line.PrintPoints()
	// Output:
	// [(0.0, 0.0) (0.0, 10.0) (0.0, 10.5) (0.0, 20.0)]
}

// Insert point to path points at index
func ExamplePath_InsertAtIndex() {
	line := path.NewPath([]v.Vec{{0, 0}, {10, 10}})
	line.InsertAtIndex(v.Vec{66, 66}, 1)
	line.PrintPoints()
	// Output:
	// [(0.0, 0.0) (66.0, 66.0) (10.0, 10.0)]
}

func ExamplePath_SetAnchor() {
	line := path.NewPath([]v.Vec{{X: 0, Y: 0}, {X: 10, Y: 10}})
	fmt.Println(line.Anchor) // Centroid of Path
	line.SetAnchor(v.Vec{X: 3, Y: 3})
	fmt.Println(line.Anchor)
	line.ResetAnchor()
	fmt.Println(line.Anchor)
	fmt.Println(line.Centroid() == line.Anchor)
	// Output:
	// (5.0, 5.0)
	// (3.0, 3.0)
	// (5.0, 5.0)
	// true
}

func ExamplePath_RemoveDoubles() {
	path := path.NewPath([]v.Vec{{0, 0}, {77, 77}, {77, 77}, {0, 0}, {0, 0}})
	path.RemoveDoubles()
	fmt.Println(path.Points)
	// Output:
	// [(0.0, 0.0) (77.0, 77.0) (0.0, 0.0)]
}
