package utils

import (
	"bufio"
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/setanarut/v"
	"golang.org/x/image/math/fixed"
)

// TangentAngle return tangent angle of two points
func TangentAngle(start, end v.Vec) float64 {
	return math.Atan2(end.Y-start.Y, end.X-start.X)
}

func Clip(number, min, max float64) float64 {
	if number < min {
		number = min
	} else if number > max {
		number = max
	}
	return number
}

// OppositeAngle returns opposite angle
func OppositeAngle(angle float64) float64 {
	return math.Mod((angle + math.Pi), (2 * math.Pi))
}

// PointOnCircle returns point at angle
func PointOnCircle(center v.Vec, radius float64, angle float64) v.Vec {
	x := center.X + (radius * math.Cos(angle))
	y := center.Y + (radius * math.Sin(angle))
	return v.Vec{x, y}
}

// Radians converts degrees to radians
func Radians(degree float64) float64 {
	return degree * (math.Pi / 180)
}

// Linspace returns evenly spaced numbers over a specified closed interval.
func Linspace(start, stop float64, num int) (res []float64) {
	if num <= 0 {
		return []float64{}
	}
	if num == 1 {
		return []float64{start}
	}
	step := (stop - start) / float64(num-1)
	res = make([]float64, num)
	res[0] = start
	for i := 1; i < num; i++ {
		res[i] = start + float64(i)*step
	}
	res[num-1] = stop
	return
}

// WritePNG writes PNG mage to disk.
func WritePNG(filePath string, img image.Image) {
	outFile, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, img)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func CloneRGBAImage(img *image.RGBA) image.Image {
	clone := image.NewRGBA(img.Rect)
	copy(clone.Pix, img.Pix)
	return clone
}

func ToFixed(v v.Vec) fixed.Point26_6 {
	return fixed.Point26_6{X: fixed.Int26_6(v.X * 64), Y: fixed.Int26_6(v.Y * 64)}
}

// RotateAbout rotates point about origin
func RotateAbout(point v.Vec, angle float64, origin v.Vec) v.Vec {
	b := v.Vec{}
	b.X = math.Cos(angle)*(point.X-origin.X) - math.Sin(angle)*(point.Y-origin.Y) + origin.X
	b.Y = math.Sin(angle)*(point.X-origin.X) + math.Cos(angle)*(point.Y-origin.Y) + origin.Y
	return b
}
