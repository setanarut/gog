package main

import (
	"fmt"

	"github.com/setanarut/gog/v2/vec"
)

func main() {
	// a, b := point.P(10, 3), point.P(102, 334)
	av, bv := vec.Vec2{10, 3}, vec.Vec2{102, 334}

	// d := a.Distance(b)
	dv := av.Distance(bv)
	// fmt.Println(d)
	fmt.Println(dv)

}
