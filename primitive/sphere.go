// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Sphere struct {
	Radius   float64
	Circular *Circular "forward:SetFa,SetFs,SetFn"
	name     string
}

func NewSphere(radius float64) *Sphere {
	return &Sphere{
		Radius:   radius,
		Circular: &Circular{},
		name:     "sphere",
	}
}

func NewCircle(radius float64) *Sphere {
	return &Sphere{
		Radius:   radius,
		Circular: &Circular{},
		name:     "circle",
	}
}

func (o *Sphere) Render(w *bufio.Writer) {
	w.WriteString(fmt.Sprintf("%s(r=%f%s);\n",
		o.name,
		o.Radius,
		o.Circular.String(),
	))
}
