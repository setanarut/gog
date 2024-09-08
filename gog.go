// Package gog is a Go Drawing Library
package gog

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/setanarut/gog/v2/path"
	"github.com/setanarut/gog/v2/utils"
	"github.com/setanarut/gog/v2/vec"
	"github.com/srwiley/rasterx"
	"github.com/srwiley/scanFT"
	"golang.org/x/image/colornames"
	"golang.org/x/image/math/fixed"
)

var Pi float64 = math.Pi

// New returns a new drawing context.
func New(width, height int) *Context {
	c := new(Context)
	c.surface = image.NewRGBA(image.Rect(0, 0, width, height))
	c.painter = scanFT.NewRGBAPainter(c.surface)
	c.scannerFreeType = scanFT.NewScannerFT(width, height, c.painter)
	c.stroker = rasterx.NewStroker(width, height, c.scannerFreeType)
	c.filler = &c.stroker.Filler
	c.Center = vec.Vec2{float64(width) / 2, float64(height) / 2}
	c.Clear(colornames.Black)
	return c
}

type Context struct {
	surface         *image.RGBA
	painter         *scanFT.RGBAPainter
	scannerFreeType *scanFT.ScannerFT
	filler          *rasterx.Filler
	stroker         *rasterx.Stroker
	AnimationFrames []image.Image
	// Center point of Canvas
	Center vec.Vec2
}

// Fill draws path with fill
func (canv *Context) Fill(p *path.Path) {
	canv.filler.Start(vec.ToFixed(p.Start()))
	for _, pt := range p.Points() {
		canv.filler.Line(vec.ToFixed(pt))
	}
	canv.filler.SetColor(p.Style.Fill)
	canv.filler.Stop(p.IsClosed())
	canv.filler.Draw()
	canv.filler.Clear()
}

// Stroke draw paths with stroke
func (canv *Context) Stroke(p *path.Path) {
	var capFunction rasterx.CapFunc
	var joinStyle rasterx.JoinMode

	switch p.Style.Cap {
	case path.ButtCap:
		capFunction = rasterx.ButtCap
	case path.SquareCap:
		capFunction = rasterx.SquareCap
	case path.RoundCap:
		capFunction = rasterx.RoundCap
	case path.CubicCap:
		capFunction = rasterx.CubicCap
	case path.QuadraticCap:
		capFunction = rasterx.QuadraticCap
	}

	switch p.Style.Join {
	case path.MiterJoin:
		joinStyle = rasterx.Miter
	case path.RoundJoin:
		joinStyle = rasterx.Round
	case path.BevelJoin:
		joinStyle = rasterx.Bevel

	}

	canv.stroker.SetStroke(
		fixed.Int26_6(p.Style.LineWidth*64), // line width
		fixed.Int26_6(3*64),                 // miter limit
		capFunction,                         // cap L
		capFunction,                         // cap T
		rasterx.RoundGap,                    // gap
		joinStyle)                           // join mode

	canv.stroker.Start(vec.ToFixed(p.Start()))
	for i := 1; i < len(p.Points()); i++ {
		canv.stroker.Line(vec.ToFixed(p.Points()[i]))
	}

	canv.stroker.SetColor(p.Style.Stroke)
	canv.stroker.Stop(p.IsClosed())
	canv.stroker.Draw()
	canv.stroker.Clear()
}

// Clear clears canvas
func (canv *Context) Clear(c color.Color) *Context {
	// m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(canv.surface, canv.surface.Bounds(),
		&image.Uniform{c}, image.Point{}, draw.Src)
	return canv
}

// AppendAnimationFrame appends current canvas to animation frames.
func (canv *Context) AppendAnimationFrame() {
	canv.AnimationFrames = append(canv.AnimationFrames, utils.CloneRGBAImage(canv.surface))
}

// SavePNG saves current canvas as static image
func (canv *Context) SavePNG(filePath string) {
	utils.WritePNG(filePath, canv.surface)
}

// Surface returns canvas surface image
func (canv *Context) Surface() *image.RGBA {
	return canv.surface
}

// SaveAPNG Saves APNG animation addes with AppendAnimationFrame().
//
// The successive delay times, one per frame, in 100ths of a second. (2 for 50 FPS, 4 for 25 FPS)
func (canv *Context) SaveAPNG(filePath string, delay int) {
	if len(canv.AnimationFrames) == 0 {
		panic("There is no frame in the image sequence, add at least one frame with AppendAnimationFrame().")
	}
	utils.WriteAnimatedPNG(filePath, canv.AnimationFrames, uint16(delay))
}

// DebugDraw draws Path attributes for debug
func (c *Context) DebugDraw(pt *path.Path) {

	p := pt.Clone()
	// BBOX
	c.Stroke(BBox(p.Bounds()).SetStrokeColor(colornames.Magenta))
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
	dot.SetPos(p.Points()[1]).SetFill(colornames.Orangered)
	c.Fill(dot)
	// POINTS
	dot.SetFill(colornames.White)
	for i := 2; i < p.Len()-1; i++ {
		dot.SetPos(p.Points()[i])
		c.Fill(dot)
	}
	// STROKE PATH
	st := p.Style.Stroke
	p.SetLineWidth(1)
	p.SetStrokeColor(pt.Style.Stroke)
	c.Stroke(p)
	p.SetStrokeColor(st)

	// centroid
	dot.SetLineWidth(2)
	dot.SetStrokeColor(colornames.Cyan)
	dot.SetPos(p.Centroid()).Scale(vec.Vec2{2, 2})
	c.Stroke(dot)

	// BBox Center
	dot.SetStrokeColor(colornames.Orange)
	a, b := p.Bounds()
	dot.SetPos(a.Lerp(b, 0.5))
	c.Stroke(dot)

}
