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

func ExamplePath_RemoveDoubles() {
	path := gog.NewPath([]gog.Point{{0, 0}, {77, 77}, {77, 77}, {0, 0}, {0, 0}})
	path.RemoveDoubles()
	fmt.Println(path.GetPoints())
	// Output:
	// [{0 0} {77 77} {0 0}]
}
func ExampleStyle() {
	myStyle := gog.NewStyle(color.RGBA{255, 0, 0, 255},
		color.Gray{128}, 10, gog.RoundCap, gog.RoundJoin)
	myStyle2 := gog.Style{
		Fill:      color.RGBA{255, 255, 0, 255},
		Stroke:    color.RGBA{255, 0, 255, 255},
		LineWidth: 7,
		Cap:       gog.CubicCap,
		Join:      gog.BevelJoin,
	}
	square := gog.Square(gog.Point{10, 10}, 50).SetStyle(myStyle)
	square2 := gog.Square(gog.Point{10, 10}, 50)
	square2.Style = myStyle2
	square2.SetFill(color.RGBA{0, 255, 255, 255})
	fmt.Printf("%+v\n%+v", square.Style, square2.Style)
	// Output:
	// {Fill:{R:255 G:0 B:0 A:255} Stroke:{Y:128} LineWidth:10 Cap:2 Join:1}
	// {Fill:{R:0 G:255 B:255 A:255} Stroke:{R:255 G:0 B:255 A:255} LineWidth:7 Cap:3 Join:2}
}
