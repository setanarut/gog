package gog_test

import (
	"fmt"
	"image/color"

	"github.com/setanarut/gog"
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

// Get point and tangent angle at time t
func ExamplePath_PointAngleAtTime() {
	line := gog.NewPath([]gog.Point{{0, 0}, {10, 10}})
	point, angle := line.PointAngleAtTime(0.5)
	fmt.Println(point, angle)
	// Output:
	// {5 5} 0.7853981633974483
}

// Get point and tangent angle at length
func ExamplePath_PointAngleAtLength() {
	line := gog.NewPath([]gog.Point{{0, 0}, {10, 10}})
	point, angle := line.PointAngleAtLength(line.Length() / 2)
	fmt.Println(point, angle)
	// Output:
	// {5 5} 0.7853981633974483
}

func ExamplePath_InsertAtLength() {
	line := gog.NewPath([]gog.Point{{0, 0}, {0, 10}, {0, 20}})
	line.InsertAtLength(10.5)
	line.PrintPoints()
	// Output:
	// [{0 0} {0 10} {0 10.5} {0 20}]
}

// Insert point to path points at index
func ExamplePath_InsertAtIndex() {
	line := gog.NewPath([]gog.Point{{0, 0}, {10, 10}})
	line.InsertAtIndex(gog.Point{66, 66}, 1)
	line.PrintPoints()
	// Output:
	// [{0 0} {66 66} {10 10}]
}

func ExamplePath_SetAnchor() {
	line := gog.NewPath([]gog.Point{{0, 0}, {10, 10}})
	fmt.Println(line.Anchor) // Centroid of Path
	line.SetAnchor(gog.Point{3, 3})
	fmt.Println(line.Anchor)
	line.ResetAnchor()
	fmt.Println(line.Anchor)
	fmt.Println(line.Centroid() == line.Anchor)
	// Output:
	// {5 5}
	// {3 3}
	// {5 5}
	// true
}
