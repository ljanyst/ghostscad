// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package shapes

import (
	"math"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/lib/degmath"
	. "github.com/ljanyst/ghostscad/primitive"
)

type Sector struct {
	Primitive  Primitive
	Radius     float64
	StartAngle float64
	EndAngle   float64
	Fn         uint16 "optional"
}

func NewSector(radius, startAngle, endAngle float64) *Sector {
	return &Sector{
		Radius:     radius,
		StartAngle: startAngle,
		EndAngle:   endAngle,
		Fn:         24,
	}
}

func (o *Sector) Build() Primitive {
	r := o.Radius / math.Cos(math.Pi/float64(o.Fn))
	step := float64(360) / float64(o.Fn)
	points := []Vec2{{0, 0}}

	for a := o.StartAngle; a >= o.EndAngle-360; a -= step {
		points = append(points, Vec2{r * Cos(a), r * Sin(a)})
	}
	points = append(points, Vec2{r * Cos(o.EndAngle), r * Sin(o.EndAngle)})
	o.Primitive = NewDifference(
		NewCircle(o.Radius).SetFn(o.Fn),
		NewPolygon(points),
	)
	return o.Primitive
}
