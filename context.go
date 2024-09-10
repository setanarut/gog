// Package gog is a Go Drawing Library
package gog

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/setanarut/gog/v2/path"
	"github.com/setanarut/gog/v2/utils"
	"github.com/setanarut/gog/v2/vec"
	"github.com/srwiley/rasterx"
	"github.com/srwiley/scanFT"
	"golang.org/x/image/colornames"
	"golang.org/x/image/math/fixed"
)

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

// NewContext returns a new drawing context.
func NewContext(width, height int) *Context {
	c := new(Context)
	c.surface = image.NewRGBA(image.Rect(0, 0, width, height))
	c.painter = scanFT.NewRGBAPainter(c.surface)
	c.scannerFreeType = scanFT.NewScannerFT(width, height, c.painter)
	c.stroker = rasterx.NewStroker(width, height, c.scannerFreeType)
	c.filler = &c.stroker.Filler
	c.Center = vec.Vec2{float64(width) / 2, float64(height) / 2}
	c.Clear(color.Black)
	return c
}

// Fill draws path with fill
func (c *Context) Fill(p *path.Path, s *Style) {
	c.filler.Start(vec.ToFixed(p.Start()))
	for _, pt := range p.Points() {
		c.filler.Line(vec.ToFixed(pt))
	}
	c.filler.SetColor(s.Fill)
	c.filler.Stop(p.IsClosed())
	c.filler.Draw()
	c.filler.Clear()
}

// Stroke draw paths with stroke
func (ctx *Context) Stroke(p *path.Path, style *Style) {
	var capFunction rasterx.CapFunc
	var joinStyle rasterx.JoinMode

	switch style.Cap {
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

	switch style.Join {
	case MiterJoin:
		joinStyle = rasterx.Miter
	case RoundJoin:
		joinStyle = rasterx.Round
	case BevelJoin:
		joinStyle = rasterx.Bevel

	}

	ctx.stroker.SetStroke(
		fixed.Int26_6(style.LineWidth*64), // line width
		fixed.Int26_6(3*64),               // miter limit
		capFunction,                       // cap L
		capFunction,                       // cap T
		rasterx.RoundGap,                  // gap
		joinStyle)                         // join mode

	ctx.stroker.Start(vec.ToFixed(p.Start()))
	for i := 1; i < len(p.Points()); i++ {
		ctx.stroker.Line(vec.ToFixed(p.Points()[i]))
	}

	ctx.stroker.SetColor(style.StrokeColor)
	ctx.stroker.Stop(p.IsClosed())
	ctx.stroker.Draw()
	ctx.stroker.Clear()
}

// Clear clears canvas
func (ctx *Context) Clear(c color.Color) *Context {
	// m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(ctx.surface, ctx.surface.Bounds(),
		&image.Uniform{c}, image.Point{}, draw.Src)
	return ctx
}

// AppendAnimationFrame appends current canvas to animation frames.
func (ctx *Context) AppendAnimationFrame() {
	ctx.AnimationFrames = append(ctx.AnimationFrames, utils.CloneRGBAImage(ctx.surface))
}

// ClearAnimationFrames clears context.AnimationFrames
func (ctx *Context) ClearAnimationFrames() {
	ctx.AnimationFrames = nil
}

// SavePNG saves current canvas as static image
func (ctx *Context) SavePNG(filePath string) {
	utils.WritePNG(filePath, ctx.surface)
}

// Surface returns canvas surface image
func (ctx *Context) Surface() *image.RGBA {
	return ctx.surface
}

// SaveAPNG Saves APNG animation addes with AppendAnimationFrame().
//
// The successive delay times, one per frame, in 100ths of a second. (2 for 50 FPS, 4 for 25 FPS)
func (ctx *Context) SaveAPNG(filePath string, delay int) {
	if len(ctx.AnimationFrames) == 0 {
		panic("There is no frame in the image sequence, add at least one frame with AppendAnimationFrame().")
	}
	utils.WriteAnimatedPNG(filePath, ctx.AnimationFrames, uint16(delay))
}

// DebugDraw draws Path attributes for debug
func (ctx *Context) DebugDraw(pth *path.Path) {

	// Draw Bounding box
	ctx.Stroke(BBox(pth.Bounds()), debugStyle)

	// Draw start point
	dot.SetFillColor(colornames.Yellow)
	dot.SetPos(pth.Start())
	ctx.Fill(dot)

	// Draw end point
	dot := Circle(pth.Start(), 2)
	dot.SetFillColor(colornames.Yellow)
	dot.SetPos(pth.End())
	ctx.Fill(dot)

	// Draw Second point
	dot.SetPos(pth.Points()[1]).SetFillColor(colornames.Orangered)
	ctx.Fill(dot)

	// Draw Points
	dot.SetFillColor(colornames.White)
	for i := 2; i < pth.Len()-1; i++ {
		dot.SetPos(pth.Points()[i])
		ctx.Fill(dot)
	}
	// Draw Stroke path
	st := Style.StrokeColor
	pth.SetLineWidth(1)
	pth.SetStrokeColor(p.Style.Stroke)
	ctx.Stroke(pth)
	pth.SetStrokeColor(st)

	// Draw Centroid
	dot.SetLineWidth(2)
	dot.SetStrokeColor(colornames.Cyan)
	dot.SetPos(pth.Centroid()).Scale(vec.Vec2{2, 2})
	ctx.Stroke(dot)

	// Draw BBox Center
	dot.SetStrokeColor(colornames.Orange)
	a, b := pth.Bounds()
	dot.SetPos(a.Lerp(b, 0.5))
	ctx.Stroke(dot)

}
