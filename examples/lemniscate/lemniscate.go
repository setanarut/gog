package main

import (
	"github.com/setanarut/gog"
)

func main() {
	g := gog.New(500, 500)
	inf := gog.Lemniscate(50, 240)
	inf.Translate(250, 250)
	inf.DebugDraw(g)
	// inf.Stroke(g)
	g.SavePNG("lemniscate.png")
}
