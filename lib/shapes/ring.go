// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package shapes

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

type Ring struct {
	Primitive  Primitive
	Radius1    float64
	Radius2    float64
	StartAngle float64 "optional"
	RingAngle  float64 "optional"

	StartAnchor *Anchor
	EndAnchor   *Anchor
}

func NewRing(r1, r2 float64) *Ring {
	return &Ring{
		Radius1:    r1,
		Radius2:    r2,
		StartAngle: 0,
		RingAngle:  360,
	}
}

func (o *Ring) Build() Primitive {
	o.StartAnchor = NewAnchor()
	o.EndAnchor = NewAnchor()
	o.Primitive =
		// Rotate to the start angle
		NewRotation(
			Vec3{0, 0, o.StartAngle},
			// Place the start anchor at beginning of the ring
			NewTranslation(
				Vec3{o.Radius1, 0, 0},
				o.StartAnchor),

			// Produce the ring
			NewRotationExtrusion(
				NewTranslation(
					Vec3{o.Radius1, 0, 0},
					NewCircle(o.Radius2))).SetAngle(o.RingAngle),

			// Place the end anchor at the end of the ring
			NewRotation(
				Vec3{0, 0, o.RingAngle},
				NewTranslation(
					Vec3{o.Radius1, 0, 0},
					o.EndAnchor)))

	return o.Primitive
}
