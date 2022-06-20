package main

import (
	"github.com/ljanyst/ghostscad/lib/shapes"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/go-gl/mathgl/mgl64"
)

func main() {
	sys.Initialize()
	points := []Vec3{{0, 0, 0}, {40, 60, 0}, {-50, 90, 0}, {0, 200, 0}}
	sys.RenderOne(shapes.NewBezier(points, 2).Build())
}
