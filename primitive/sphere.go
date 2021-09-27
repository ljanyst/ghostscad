// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Sphere struct {
	Radius float64
	Circular
	name string
}

func NewSphere(radius float64) *Sphere {
	return &Sphere{
		Radius: radius,
		name:   "sphere",
	}
}

func NewCircle(radius float64) *Sphere {
	return &Sphere{
		Radius: radius,
		name:   "circle",
	}
}

func (o *Sphere) SetFa(val float64) *Sphere {
	o.Circular.SetFa(val)
	return o
}

func (o *Sphere) SetFs(val float64) *Sphere {
	o.Circular.SetFs(val)
	return o
}

func (o *Sphere) SetFn(val uint16) *Sphere {
	o.Circular.SetFn(val)
	return o
}

func (o *Sphere) Render(w *bufio.Writer) {

	w.WriteString(fmt.Sprintf("%s(r=%f%s);\n",
		o.name,
		o.Radius,
		o.Circular.String(),
	))
}
