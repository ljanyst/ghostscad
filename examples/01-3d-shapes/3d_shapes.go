package main

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
	"github.com/ljanyst/ghostscad/sys"
)

func newEmptyTriangle() Primitive {
	points := []Vec3{
		Vec3{0, -10, 60},
		Vec3{0, 10, 60},
		Vec3{0, 10, 0},
		Vec3{0, -10, 0},
		Vec3{60, -10, 60},
		Vec3{60, 10, 60},
		Vec3{10, -10, 50},
		Vec3{10, 10, 50},
		Vec3{10, 10, 30},
		Vec3{10, -10, 30},
		Vec3{30, -10, 50},
		Vec3{30, 10, 50},
	}
	faces := []Vec3{
		Vec3{0, 3, 2},
		Vec3{0, 2, 1},
		Vec3{4, 0, 5},
		Vec3{5, 0, 1},
		Vec3{5, 2, 4},
		Vec3{4, 2, 3},
		Vec3{6, 8, 9},
		Vec3{6, 7, 8},
		Vec3{6, 10, 11},
		Vec3{6, 11, 7},
		Vec3{10, 8, 11},
		Vec3{10, 9, 8},
		Vec3{3, 0, 9},
		Vec3{9, 0, 6},
		Vec3{10, 6, 0},
		Vec3{0, 4, 10},
		Vec3{3, 9, 10},
		Vec3{3, 10, 4},
		Vec3{1, 7, 11},
		Vec3{1, 11, 5},
		Vec3{1, 8, 7},
		Vec3{2, 8, 1},
		Vec3{8, 2, 11},
		Vec3{5, 11, 2},
	}
	return NewTranslation(
		Vec3{-5, 0, -5},
		NewScale(
			Vec3{1. / 6, 1. / 6, 1. / 6},
			NewPolyhedron(points, faces),
		),
	)
}

func newWave() Primitive {
	sv := 0.2173
	return NewScale(
		Vec3{sv, sv, sv},
		NewSurface("wave.dat").SetConvexity(5))
}

func buildShapes() Primitive {
	return NewTranslation(
		Vec3{-50, 0, 0},
		NewSphere(5).SetFs(1).SetFn(0),
		NewTranslation(
			Vec3{20, 0, 0},
			NewCube(Vec3{10, 10, 10}),
		),
		NewTranslation(
			Vec3{40, 0, 0},
			NewImport("die.stl").Highlight(),
		),
		NewTranslation(
			Vec3{60, 0, 0},
			NewCylinder(10, 3).SetRBottom(5).Transparent(),
		),
		NewTranslation(
			Vec3{80, 0, 0},
			newEmptyTriangle(),
		),
		NewTranslation(
			Vec3{100, 0, 0},
			newWave(),
		),
	)
}

func main() {
	sys.SetFn(360)
	sys.RenderAndExit(buildShapes())
}
