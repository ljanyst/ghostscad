package main

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
	"github.com/ljanyst/ghostscad/sys"
)

func buildShapes() Primitive {
	return NewTranslation(
		Vec3{-40, 0, 0},
		NewSphere(5),
		NewTranslation(
			Vec3{20, 0, 0},
			NewCube(Vec3{10, 10, 10}),
		),
		NewTranslation(
			Vec3{40, 0, 0},
			NewImport("die.stl"),
		),
		NewTranslation(
			Vec3{60, 0, 0},
			NewCylinder(10, 3).SetRBottom(5),
		),
	)
}

func main() {
	sys.RenderAndExit(buildShapes())
}
