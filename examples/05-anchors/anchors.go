package main

import (
	"github.com/ljanyst/ghostscad/lib/shapes"
	"github.com/ljanyst/ghostscad/lib/utils"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

func main() {
	sys.SetFn(120)
	ring1 := shapes.NewRing(20, 5).SetStartAngle(30).SetRingAngle(90)
	ring1.Build()

	ring2 := shapes.NewRing(20, 5).SetStartAngle(220).SetRingAngle(45)
	ring2.Build()

	a1 := NewAnchor()
	sphere21 := NewTranslation(
		Vec3{32, 41, 53},
		NewRotation(
			Vec3{32, 443, 12},
			a1,
			NewSphere(5)))

	a2 := NewAnchor()
	sphere22 := NewTranslation(
		Vec3{-43, 1, -322},
		NewRotation(
			Vec3{-443, 2, -74},
			a2,
			NewSphere(5)))

	sys.RenderOne(NewList(
		utils.AlignOrigin(ring1.StartAnchor).Add(
			NewSphere(5)),
		ring1.Primitive,
		utils.AlignOrigin(ring1.EndAnchor).Add(
			NewSphere(5)),

		utils.Align(ring2.StartAnchor, a1).Add(
			sphere21),
		ring2.Primitive,
		utils.Align(ring2.EndAnchor, a2).Add(
			sphere22),
	))
}
