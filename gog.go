// Package gog is a Go Drawing Library
package gog

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/srwiley/rasterx"
	"github.com/srwiley/scanFT"
	"golang.org/x/image/colornames"
	"golang.org/x/image/math/fixed"
)

var Pi float64 = math.Pi

// New returns a new drawing context.
func New(width, height int) *context {
	c := new(context)
	c.surface = image.NewRGBA(image.Rect(0, 0, width, height))
	c.painter = scanFT.NewRGBAPainter(c.surface)
	c.scannerFreeType = scanFT.NewScannerFT(width, height, c.painter)
	c.stroker = rasterx.NewStroker(width, height, c.scannerFreeType)
	c.filler = &c.stroker.Filler
	c.Center = Point{float64(width) / 2, float64(height) / 2}
	c.Clear(Black)
	return c
}

type context struct {
	surface         *image.RGBA
	painter         *scanFT.RGBAPainter
	scannerFreeType *scanFT.ScannerFT
	filler          *rasterx.Filler
	stroker         *rasterx.Stroker
	AnimationFrames []image.Image
	// Center point of Canvas
	Center Point
}

// Fill draws path with fill
func (canv *context) Fill(p *Path) {
	canv.filler.Start(p.Start().Fixed())
	for _, pt := range p.points {
		canv.filler.Line(pt.Fixed())
	}
	canv.filler.SetColor(p.Style.Fill)
	canv.filler.Stop(p.IsClosed())
	canv.filler.Draw()
	canv.filler.Clear()
}

// Stroke draw paths with stroke
func (canv *context) Stroke(p *Path) {
	var capFunction rasterx.CapFunc
	var joinStyle rasterx.JoinMode

	switch p.Style.Cap {
	case ButtCap:
		capFunction = rasterx.ButtCap
	case SquareCap:
		capFunction = rasterx.SquareCap
	case RoundCap:
		capFunction = rasterx.RoundCap
	case CubicCap:
		capFunction = rasterx.CubicCap
	case QuadraticCap:
		capFunction = rasterx.QuadraticCap
	}

	switch p.Style.Join {
	case MiterJoin:
		joinStyle = rasterx.Miter
	case RoundJoin:
		joinStyle = rasterx.Round
	case BevelJoin:
		joinStyle = rasterx.Bevel

	}

	canv.stroker.SetStroke(
		fixed.Int26_6(p.Style.LineWidth*64), // line width
		fixed.Int26_6(3*64),                 // miter limit
		capFunction,                         // cap L
		capFunction,                         // cap T
		rasterx.RoundGap,                    // gap
		joinStyle)                           // join mode

	canv.stroker.Start(p.Start().Fixed())
	for i := 1; i < len(p.points); i++ {
		canv.stroker.Line(p.points[i].Fixed())
	}

	canv.stroker.SetColor(p.Style.Stroke)
	canv.stroker.Stop(p.IsClosed())
	canv.stroker.Draw()
	canv.stroker.Clear()
}

// Clear clears canvas
func (canv *context) Clear(c color.Color) *context {
	// m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(canv.surface, canv.surface.Bounds(),
		&image.Uniform{c}, image.Point{}, draw.Src)
	return canv
}

// AppendAnimationFrame appends current canvas to animation frames.
func (canv *context) AppendAnimationFrame() {
	canv.AnimationFrames = append(canv.AnimationFrames, cloneRGBAImage(canv.surface))
}

// SavePNG saves current canvas as static image
func (canv *context) SavePNG(filePath string) {
	writePNG(filePath, canv.surface)
}

// SaveAPNG Saves APNG animation addes with AppendAnimationFrame().
//
// The successive delay times, one per frame, in 100ths of a second. (2 for 50 FPS, 4 for 25 FPS)
func (canv *context) SaveAPNG(filePath string, delay int) {
	if len(canv.AnimationFrames) == 0 {
		panic("There is no frame in the image sequence, add at least one frame with AppendAnimationFrame().")
	}
	writeAnimatedPNG(filePath, canv.AnimationFrames, uint16(delay))
}

// DebugDraw draws Path attributes for debug
func (c *context) DebugDraw(path *Path) {

	p := path.Clone()
	// BBOX
	BBox(p.Bounds()).SetStroke(colornames.Magenta).Stroke(c)
	// END
	dot := Circle(p.Start(), 2)
	dot.SetFill(colornames.Yellow)
	dot.SetPos(p.End())
	c.Fill(dot)
	// START
	dot.SetFill(colornames.Yellow)
	dot.SetPos(p.Start())
	c.Fill(dot)
	// SECOND POINT
	dot.SetPos(p.points[1]).SetFill(colornames.Orangered)
	c.Fill(dot)
	// POINTS
	dot.SetFill(colornames.White)
	for i := 2; i < p.Len()-1; i++ {
		dot.SetPos(p.points[i])
		c.Fill(dot)
	}
	// STROKE PATH
	st := p.Style.Stroke
	p.SetLineWidth(1)
	p.SetStroke(Gray)
	p.Stroke(c)
	p.SetStroke(st)

	// centroid
	dot.SetLineWidth(2)
	dot.SetStroke(colornames.Cyan)
	dot.SetPos(p.Centroid()).Scale(Point{2, 2})
	c.Stroke(dot)

	// BBox Center
	dot.SetStroke(colornames.Orange)
	a, b := p.Bounds()
	dot.SetPos(a.Lerp(b, 0.5))
	c.Stroke(dot)

}
