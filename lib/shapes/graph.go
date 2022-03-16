// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package shapes

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

type Graph struct {
	Primitive  Primitive
	Func       func(float64, float64) float64
	RangeX     Vec2
	RangeY     Vec2
	Resolution Vec2
	Thickness  float64 "optional"
	Convexity  int     "optional"
}

func NewGraph(f func(float64, float64) float64, rangeX, rangeY, resolution Vec2) *Graph {
	return &Graph{
		Func:       f,
		RangeX:     rangeX,
		RangeY:     rangeY,
		Resolution: resolution,
		Thickness:  1,
	}
}

func (o *Graph) Build() Primitive {
	pointsUp := []Vec3{}
	pointsDown := []Vec3{}
	faces := []Vec3{}

	itersX := int((o.RangeX[1]-o.RangeX[0])/o.Resolution[0]) + 1
	itersY := int((o.RangeY[1]-o.RangeY[0])/o.Resolution[1]) + 1

	lenArr := itersX * itersY
	indUp := func(i, j int) float64 {
		return float64(i*itersY + j)
	}
	indDown := func(i, j int) float64 {
		return indUp(i, j) + float64(lenArr)
	}

	f := func(x, y float64) Vec3 {
		val := o.Func(x, y)
		return Vec3{x, y, val}
	}
	offset := Vec3{0, 0, 0.5 * o.Thickness}

	for i := 0; i < itersX; i++ {
		for j := 0; j < itersY; j++ {
			x := o.RangeX[0] + float64(i)*o.Resolution[0]
			y := o.RangeY[0] + float64(j)*o.Resolution[1]
			pt := f(x, y)
			pointsUp = append(pointsUp, pt.Add(offset))
			pointsDown = append(pointsDown, pt.Sub(offset))

			if i > 0 && j > 0 {
				faces = append(faces,
					// upward facing faces clockwise
					Vec3{indUp(i-1, j), indUp(i, j), indUp(i, j-1)},
					Vec3{indUp(i, j-1), indUp(i-1, j-1), indUp(i-1, j)},
					// downward facing faces counter clockwise
					Vec3{indDown(i, j-1), indDown(i, j), indDown(i-1, j)},
					Vec3{indDown(i-1, j), indDown(i-1, j-1), indDown(i, j-1)},
				)
			}
		}
	}

	tmp := itersX - 1
	for j := 1; j < itersY; j++ {
		faces = append(faces,
			// left side
			Vec3{indUp(0, j), indUp(0, j-1), indDown(0, j)},
			Vec3{indDown(0, j), indUp(0, j-1), indDown(0, j-1)},
			// right side
			Vec3{indUp(tmp, j-1), indUp(tmp, j), indDown(tmp, j)},
			Vec3{indDown(tmp, j), indDown(tmp, j-1), indUp(tmp, j-1)},
		)
	}

	tmp = itersY - 1
	for i := 1; i < itersX; i++ {
		faces = append(faces,
			// front side
			Vec3{indUp(i-1, 0), indUp(i, 0), indDown(i-1, 0)},
			Vec3{indDown(i-1, 0), indUp(i, 0), indDown(i, 0)},
			// back side
			Vec3{indUp(i-1, tmp), indUp(i, tmp), indDown(i, tmp)},
			Vec3{indDown(i, tmp), indDown(i-1, tmp), indUp(i-1, tmp)},
		)
	}

	points := append(pointsUp, pointsDown...)
	ph := NewPolyhedron(points, faces)

	if o.Convexity != 0 {
		ph.SetConvexity(o.Convexity)
	}

	o.Primitive = ph
	return o.Primitive
}
