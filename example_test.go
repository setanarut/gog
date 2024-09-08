package gog_test

import (
	"fmt"
	"image/color"

	"github.com/setanarut/gog/v2"
	"github.com/setanarut/gog/v2/path"
	"github.com/setanarut/gog/v2/vec"
)

// 150-frame rotating cubic bezier APNG animation
func Example() {
	ctx := gog.New(250, 250)
	curve := gog.CubicBezier(100, 95, 50, 300, 190, 88, 140, 200, 50)
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
	line := gog.Line(vec.Vec2{X: 0, Y: 0}, vec.Vec2{X: 25, Y: 80})
	fmt.Println(line.Start(), line.End())
	// Output:
	// Vec2{X: 0.000000, Y: 0.000000} Vec2{X: 25.000000, Y: 80.000000}
}

// Get point and tangent angle at time t
func ExamplePath_PointAngleAtTime() {
	line := path.NewPath([]vec.Vec2{{X: 0, Y: 0}, {X: 10, Y: 10}})
	point, angle := line.PointAngleAtTime(0.5)
	fmt.Println(point, angle)
	// Output:
	// Vec2{X: 5.000000, Y: 5.000000} 0.7853981633974483
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
func ExampleStyle() {
	myStyle := path.NewStyle(color.RGBA{255, 0, 0, 255},
		color.Gray{128}, 10, path.RoundCap, path.RoundJoin)
	myStyle2 := path.Style{
		Fill:      color.RGBA{255, 255, 0, 255},
		Stroke:    color.RGBA{255, 0, 255, 255},
		LineWidth: 7,
		Cap:       path.CubicCap,
		Join:      path.BevelJoin,
	}
	square := gog.Square(vec.Vec2{10, 10}, 50).SetStyle(myStyle)
	square2 := gog.Square(vec.Vec2{10, 10}, 50)
	square2.Style = myStyle2
	square2.SetFill(color.RGBA{0, 255, 255, 255})
	fmt.Printf("%+v\n%+v", square.Style, square2.Style)
	// Output:
	// {Fill:{R:255 G:0 B:0 A:255} Stroke:{Y:128} LineWidth:10 Cap:2 Join:1}
	// {Fill:{R:0 G:255 B:255 A:255} Stroke:{R:255 G:0 B:255 A:255} LineWidth:7 Cap:3 Join:2}
}
