package main

import (
	"math"

	"github.com/ljanyst/ghostscad/lib/shapes"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/go-gl/mathgl/mgl64"
)

func main() {
	sys.Initialize()
	f := func(x, y float64) float64 {
		return (math.Pow(y, 2) / 4) - (math.Pow(x, 2) / 4)
	}
	sys.RenderOne(shapes.NewGraph(f, Vec2{-3, 3}, Vec2{-3, 3}, Vec2{0.25, 0.25}).Build())
}
