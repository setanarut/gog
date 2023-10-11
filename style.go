package gog

import (
	"image/color"
)

var White = color.White
var Black = color.Black
var Gray = color.Gray{128}
var Gray40 = color.Gray{40}

func NewStyle(fill, stroke color.Color, lineWidth float64) Style {
	return Style{
		Fill:      fill,
		Stroke:    stroke,
		LineWidth: lineWidth,
		Cap:       ButtCap,
		Join:      MiterJoin,
	}
}

// Style of path
type Style struct {
	Fill      color.Color
	Stroke    color.Color
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

func DefaultStyle() Style {
	return Style{
		Fill:      Gray,
		Stroke:    White,
		LineWidth: 1.5,
		Cap:       ButtCap,
		Join:      MiterJoin,
	}
}

type (
	JoinMode uint8
	CapMode  uint8
	DrawMode uint8
)

// Drawing style constant
const (
	Fill DrawMode = iota
	Stroke
	FillStroke
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
