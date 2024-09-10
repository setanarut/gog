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
	"golang.org/x/image/math/fixed"
)

var yellow = color.RGBA{255, 255, 0, 255}
var orangered = color.RGBA{255, 69, 0, 255}

var debugPathStrokeStyle = &StrokeStyle{
	Color:     color.White,
	LineWidth: 1,
}
var debugCentroidStrokeStyle = &StrokeStyle{
	Color:     color.RGBA{0, 255, 255, 255},
	LineWidth: 2,
}
var debugBBoxCenterStrokeStyle = &StrokeStyle{
	Color:     color.RGBA{255, 128, 0, 255},
	LineWidth: 2,
}

type Context struct {
	AnimationFrames []image.Image
	// Center point of Canvas
	Center vec.Vec2

	surface         *image.RGBA
	painter         *scanFT.RGBAPainter
	scannerFreeType *scanFT.ScannerFT
	filler          *rasterx.Filler
	stroker         *rasterx.Stroker
}

// NewContext returns a new drawing context.
func NewContext(width, height int) *Context {
	ctx := new(Context)
	ctx.surface = image.NewRGBA(image.Rect(0, 0, width, height))
	ctx.painter = scanFT.NewRGBAPainter(ctx.surface)
	ctx.scannerFreeType = scanFT.NewScannerFT(width, height, ctx.painter)
	ctx.stroker = rasterx.NewStroker(width, height, ctx.scannerFreeType)
	ctx.filler = &ctx.stroker.Filler
	ctx.Center = vec.Vec2{float64(width) / 2, float64(height) / 2}
	ctx.Clear(color.Black)
	return ctx
}

// Fill draws path with fillColor
func (ctx *Context) Fill(p *path.Path, fillColor color.Color) {
	ctx.filler.Start(vec.ToFixed(p.Start()))
	for _, pt := range p.Points() {
		ctx.filler.Line(vec.ToFixed(pt))
	}
	ctx.filler.SetColor(fillColor)
	ctx.filler.Stop(p.IsClosed())
	ctx.filler.Draw()
	ctx.filler.Clear()
}

// Stroke draw paths with StrokeStyle
func (ctx *Context) Stroke(p *path.Path, strokeStyle *StrokeStyle) {
	var capFunction rasterx.CapFunc
	var joinStyle rasterx.JoinMode

	switch strokeStyle.Cap {
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

	switch strokeStyle.Join {
	case MiterJoin:
		joinStyle = rasterx.Miter
	case RoundJoin:
		joinStyle = rasterx.Round
	case BevelJoin:
		joinStyle = rasterx.Bevel

	}

	ctx.stroker.SetStroke(
		fixed.Int26_6(strokeStyle.LineWidth*64), // line width
		fixed.Int26_6(3*64),                     // miter limit
		capFunction,                             // cap L
		capFunction,                             // cap T
		rasterx.RoundGap,                        // gap
		joinStyle)                               // join mode

	ctx.stroker.Start(vec.ToFixed(p.Start()))
	for i := 1; i < len(p.Points()); i++ {
		ctx.stroker.Line(vec.ToFixed(p.Points()[i]))
	}

	ctx.stroker.SetColor(strokeStyle.Color)
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
	circle := Circle(pth.Start(), 2)
	ctx.Fill(circle, yellow)

	// Draw end point
	ctx.Fill(circle.SetPos(pth.End()), yellow)

	// Draw second point
	ctx.Fill(circle.SetPos(pth.Points()[1]), orangered)

	// Draw all points
	for i := 2; i < pth.Len()-1; i++ {
		circle.SetPos(pth.Points()[i])
		ctx.Fill(circle, color.White)
	}
	// Draw path stroke
	ctx.Stroke(pth, debugPathStrokeStyle)

	// Draw Centroid
	circle.SetPos(pth.Centroid()).Scale(vec.Vec2{2, 2})
	ctx.Stroke(circle, debugCentroidStrokeStyle)

	// Draw BBox center
	a, b := pth.Bounds()
	circle.SetPos(a.Lerp(b, 0.5))
	ctx.Stroke(circle, debugBBoxCenterStrokeStyle)

}
