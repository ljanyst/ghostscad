package main

import (
	"math"

	"github.com/ljanyst/ghostscad/lib/shapes"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/go-gl/mathgl/mgl64"
)

func main() {
	f := func(x, y float64) float64 {
		return (math.Pow(y, 2) / math.Pow(2, 2)) - (math.Pow(x, 2) / math.Pow(2, 2))
	}
	sys.RenderOne(shapes.NewGraph(f, Vec2{-3, 3}, Vec2{-3, 3}, Vec2{0.25, 0.25}).Build())
}
