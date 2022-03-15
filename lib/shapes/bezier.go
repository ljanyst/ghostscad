// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package shapes

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

type Bezier struct {
	Primitive Primitive
	Points    []Vec3
	Thickness float64
	Fn        uint16 "optional"
}

func NewBezier(points []Vec3, thickness float64) *Bezier {
	return &Bezier{
		Points:    points,
		Thickness: thickness,
		Fn:        48,
	}
}

func (o *Bezier) Build() Primitive {
	step := 1 / float64(o.Fn)
	points := []Vec3{}
	for t := 0.; t < 1; t += step {
		points = append(points, BezierCurve3D(t, o.Points))
	}
	points = append(points, BezierCurve3D(1, o.Points))
	o.Primitive = NewPolyline3d(points, o.Thickness).Build()
	return o.Primitive
}
