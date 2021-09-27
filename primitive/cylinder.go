// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Cylinder struct {
	H        float64
	RTop     float64
	RBottom  float64   "optional"
	Center   bool      "optional"
	Circular *Circular "forward:SetFa,SetFs,SetFn"
}

func NewCylinder(h, r float64) *Cylinder {
	return &Cylinder{
		H:        h,
		RTop:     r,
		RBottom:  r,
		Center:   true,
		Circular: &Circular{},
	}
}

func (o *Cylinder) Render(w *bufio.Writer) {
	w.WriteString(
		fmt.Sprintf(
			"cylinder(h=%f, r1=%f, r2=%f, center=%t%s);\n",
			o.H, o.RBottom, o.RTop, o.Center, o.Circular.String(),
		),
	)
}
