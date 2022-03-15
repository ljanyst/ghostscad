package main

import (
	"github.com/ljanyst/ghostscad/lib/shapes"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/go-gl/mathgl/mgl64"
)

func main() {
	points := []Vec3{{1, 2, 3}, {4, -5, -6}, {-1, -3, -5}, {0, 0, 0}}
	sys.RenderOne(shapes.NewPolyline3d(points, 1).Build())
}
