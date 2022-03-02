package main

import (
	"math"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
	"github.com/ljanyst/ghostscad/sys"
)

func newColor() Primitive {
	lst := NewList()
	for i := 0.0; i < 36; i++ {
		for j := 0.0; j < 36; j++ {
			r := byte(255 * (0.5 + math.Sin(DegToRad(10*i))/2))
			g := byte(255 * (0.5 + math.Sin(DegToRad(10*j))/2))
			b := byte(255 * (0.5 + math.Sin(DegToRad(10*(i+j)))/2))
			h := 11 + 10*math.Cos(DegToRad(10*i))*math.Sin(DegToRad(10*j))
			lst.Add(
				NewColorV(r, g, b, 255,
					NewTranslation(Vec3{i, j, 0},
						NewCube(Vec3{1, 1, h}).SetCenter(false),
					)))
		}
	}
	return NewTranslation(Vec3{-18, -18, 0}, lst)
}

func newOffset() Primitive {
	base := NewDifference(
		NewOffset(NewSquare(Vec2{20, 20})).SetR(10),
		NewOffset(NewSquare(Vec2{20, 20})).SetR(8),
	)

	obj := NewLinearExtrusion(20).SetTwist(90).SetSlices(250).Add(base)
	return NewTranslation(Vec3{60, 0, 10}, obj)
}

func newMinkowski() Primitive {
	obj := NewMinkowski(
		NewCube(Vec3{30, 30, 20}).SetCenter(false),
		NewCylinder(1, 2),
	)
	return NewTranslation(Vec3{-75, -15, 0}, obj)
}

func main() {
	sys.SetFn(360)
	sys.RenderOne(NewList(
		newColor(),
		newOffset(),
		newMinkowski(),
	))
}
