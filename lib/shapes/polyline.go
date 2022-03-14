// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package shapes

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/lib/degmath"
	. "github.com/ljanyst/ghostscad/primitive"
)

func newLine(p1, p2 Vec2, width float64, round bool) Primitive {
	angle := 90 - Atan((p2[1]-p1[1])/(p2[0]-p1[0]))
	offsetX := 0.5 * width * Cos(angle)
	offsetY := 0.5 * width * Sin(angle)

	offset1 := Vec2{-offsetX, offsetY}
	offset2 := Vec2{offsetX, -offsetY}
	points := []Vec2{
		p1.Add(offset1),
		p2.Add(offset1),
		p2.Add(offset2),
		p1.Add(offset2),
	}

	ret := NewList(NewPolygon(points))

	if round {
		ret.Add(
			NewTranslation(Vec3{p1[0], p1[1], 0}, NewCircle(width/2).SetFn(48)),
			NewTranslation(Vec3{p2[0], p2[1], 0}, NewCircle(width/2).SetFn(48)),
		)
	}

	return ret
}

type Polyline struct {
	Primitive Primitive
	Points    []Vec2
	Width     float64
	Round     bool "optional"
}

func NewPolyline(points []Vec2, width float64) *Polyline {
	return &Polyline{
		Points: points,
		Width:  width,
	}
}

func (o *Polyline) Build() Primitive {
	segs := NewList()
	for i := 1; i < len(o.Points); i++ {
		segs.Add(newLine(o.Points[i-1], o.Points[i], o.Width, o.Round))
	}

	o.Primitive = segs
	return o.Primitive
}
