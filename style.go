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

var debugStyle *Style = &Style{
	Fill:        color.RGBA{255, 255, 0, 255}, //Yellow
	StrokeColor: color.RGBA{255, 0, 255, 255}, //Magenta
	LineWidth:   1.5,
	Cap:         ButtCap,
	Join:        MiterJoin,
}

// Style of path
type Style struct {
	Fill        color.Color
	StrokeColor color.Color
	LineWidth   float64
	// Line cap style constant
	//
	// 0=ButtCap 1=SquareCap 2=RoundCap 3=CubicCap 4=QuadraticCap
	Cap CapMode
	// Line join style
	//
	// 0=MiterJoin 1=RoundJoin 2=BevelJoin
	Join JoinMode
}

// NewStyle shorthand for create Style{}
func NewStyle(fillColor, strokeColor color.Color, lineWidth float64, cap CapMode, join JoinMode) *Style {
	return &Style{
		Fill:        fillColor,
		StrokeColor: strokeColor,
		LineWidth:   lineWidth,
		Cap:         cap,
		Join:        join,
	}
}

// DefaultStyle returns default style
func DefaultStyle() *Style {
	return &Style{
		Fill:        color.Gray{128},
		StrokeColor: color.White,
		LineWidth:   1.5,
		Cap:         ButtCap,
		Join:        MiterJoin,
	}
}
