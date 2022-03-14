// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package shapes

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

type SmootedCube struct {
	Primitive Primitive
	Dims      Vec3
	Radius    float64
	Center    bool "optional"
}

func NewSmoothedCube(dims Vec3, r float64) *SmootedCube {
	return &SmootedCube{
		Dims:   dims,
		Radius: r,
		Center: true,
	}
}

func (o *SmootedCube) Build() Primitive {
	tr := Vec3{
		o.Radius,
		o.Radius,
		o.Radius,
	}
	if o.Center {
		tr = tr.Add(o.Dims.Mul(-0.5))
	}

	cDims := Vec3{
		o.Dims[0] - 2*o.Radius,
		o.Dims[1] - 2*o.Radius,
		o.Dims[2] - 2*o.Radius,
	}
	o.Primitive = NewTranslation(
		tr,
		NewMinkowski(
			NewCube(cDims).SetCenter(false),
			NewSphere(o.Radius)))
	return o.Primitive
}
