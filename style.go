package gog

import (
	"image/color"
)

// CapMode constants determines line cap style
const (
	ButtCap CapMode = iota
	SquareCap
	RoundCap
	CubicCap
	QuadraticCap
)

// JoinMode constants determine how stroke segments bridge the gap at a join
const (
	MiterJoin JoinMode = iota
	RoundJoin
	BevelJoin
)

type (
	JoinMode uint8
	CapMode  uint8
	DrawMode uint8
)

var debugStyle *StrokeStyle = &StrokeStyle{
	// FillColor:   color.RGBA{255, 255, 0, 255}, //Yellow
	Color:     color.RGBA{255, 0, 255, 255}, //Magenta
	LineWidth: 1.5,
	Cap:       ButtCap,
	Join:      MiterJoin,
}

// Style of path
type StrokeStyle struct {
	Color     color.Color
	LineWidth float64
	// Line cap style constant
	//
	// 0=ButtCap 1=SquareCap 2=RoundCap 3=CubicCap 4=QuadraticCap
	Cap CapMode
	// Line join style
	//
	// 0=MiterJoin 1=RoundJoin 2=BevelJoin
	Join JoinMode
}

func (s *StrokeStyle) SetColor(c color.Color) *StrokeStyle {
	s.Color = c
	return s
}
func (s *StrokeStyle) SetLineWidth(w float64) *StrokeStyle {
	s.LineWidth = w
	return s
}

// NewStyle shorthand for create Style{}
func NewStrokeStyle(fillColor, strokeColor color.Color, lineWidth float64, cap CapMode, join JoinMode) *StrokeStyle {
	return &StrokeStyle{
		Color:     strokeColor,
		LineWidth: lineWidth,
		Cap:       cap,
		Join:      join,
	}
}

// DefaultStyle returns default style
func DefaultStrokeStyle() *StrokeStyle {
	return &StrokeStyle{
		Color:     color.White,
		LineWidth: 1.5,
		Cap:       ButtCap,
		Join:      MiterJoin,
	}
}
