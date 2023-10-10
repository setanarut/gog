package gog

import (
	"bufio"
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/cia-rana/goapng"
)

func clip(number, min, max float64) float64 {
	if number < min {
		number = min
	} else if number > max {
		number = max
	}
	return number
}

// OppositeAngle returns opposite angle
func oppositeAngle(angle float64) float64 {
	return math.Mod((angle + math.Pi), (2 * math.Pi))
}

// pointOnCircle returns point at angle
func pointOnCircle(center Point, radius float64, angle float64) Point {
	x := center.X + (radius * math.Cos(angle))
	y := center.Y + (radius * math.Sin(angle))
	return Point{x, y}
}

// // deepCopyPath returns copy of this Path
// func deepCopyPath(p *Path) *Path {
// 	newPath := new(Path)
// 	newCoords := make([]Point, len(p.points))
// 	copy(newCoords, p.points)
// 	newPath.points = newCoords
// 	newPath.Style = p.Style
// 	newPath.Anchor = p.Anchor
// 	newPath.length = p.length
// 	return newPath
// }

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

// writePNG writes PNG mage to disk.
func writePNG(filePath string, img image.Image) {
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

func writeAnimatedPNG(filePath string, images []image.Image, delay uint16) {
	totalFrames := len(images)
	delays := make([]uint16, totalFrames)
	for i := range delays {
		delays[i] = delay
	}
	animPng := goapng.APNG{}
	animPng.Images = images
	animPng.Delays = delays
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	goapng.EncodeAll(file, &animPng)
}

func cloneRGBAImage(img *image.RGBA) image.Image {
	clone := image.NewRGBA(img.Rect)
	copy(clone.Pix, img.Pix)
	return clone
}
