// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package shapes

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

func newLine3d(p1, p2 Vec3, thickness float64, fn uint16) Primitive {
	return NewHull(
		NewTranslation(p1, NewSphere(thickness/2).SetFn(fn)),
		NewTranslation(p2, NewSphere(thickness/2).SetFn(fn)),
	)
}

type Polyline3d struct {
	Primitive Primitive
	Points    []Vec3
	Thickness float64
	Fn        uint16 "optional"
}

func NewPolyline3d(points []Vec3, thickness float64) *Polyline3d {
	return &Polyline3d{
		Points:    points,
		Thickness: thickness,
		Fn:        24,
	}
}

func (o *Polyline3d) Build() Primitive {
	segs := NewList()
	for i := 1; i < len(o.Points); i++ {
		segs.Add(newLine3d(o.Points[i-1], o.Points[i], o.Thickness, o.Fn))
	}

	o.Primitive = segs
	return o.Primitive
}
