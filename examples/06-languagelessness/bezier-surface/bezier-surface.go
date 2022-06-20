package main

import (
	"github.com/ljanyst/ghostscad/lib/shapes"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/go-gl/mathgl/mgl64"
)

func bezierSurface(ctrlPoints [][]Vec3) func(float64, float64) Vec3 {
	// We could to some memoization here, but the purpose of this code is to demonstrate
	// a usecase, so we value clarity and don't care much abount recomputing tha same values
	// over and over again
	return func(tx, ty float64) Vec3 {
		ptsY := []Vec3{}
		for _, pts := range ctrlPoints {
			ptsY = append(ptsY, BezierCurve3D(tx, pts))
		}
		return BezierCurve3D(ty, ptsY)
	}
}

func main() {
	sys.Initialize()
	ctrlPoints := [][]Vec3{
		{{0, 0, 20}, {60, 0, -35}, {90, 0, 60}, {200, 0, 5}},
		{{0, 50, 30}, {100, 60, -25}, {120, 50, 120}, {200, 50, 5}},
		{{0, 100, 0}, {60, 120, 35}, {90, 100, 60}, {200, 100, 45}},
		{{0, 150, 0}, {60, 150, -35}, {90, 180, 60}, {200, 150, 45}},
	}
	f := bezierSurface(ctrlPoints)
	sys.RenderOne(shapes.NewGraphT(f, 0.02).Build())
}
