// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package shapes

import (
	. "github.com/ljanyst/ghostscad/primitive"
)

type Arc struct {
	Primitive  Primitive
	Radius     float64
	StartAngle float64
	EndAngle   float64
	Width      float64 "optional"
	Fn         uint16  "optional"
}

func NewArc(radius, startAngle, endAngle float64) *Arc {
	return &Arc{
		Radius:     radius,
		StartAngle: startAngle,
		EndAngle:   endAngle,
		Width:      1,
		Fn:         24,
	}
}

func (o *Arc) Build() Primitive {
	o.Primitive = NewDifference(
		NewSector(o.Radius+0.5*o.Width, o.StartAngle, o.EndAngle).SetFn(o.Fn).Build(),
		NewSector(o.Radius-0.5*o.Width, o.StartAngle-0.1, o.EndAngle+0.1).SetFn(o.Fn).Build(),
	)
	return o.Primitive
}
