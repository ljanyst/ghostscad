// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Cylinder struct {
	ParentImpl
	H        float64
	RTop     float64
	RBottom  float64     "optional"
	Center   bool        "optional"
	Circular *Circular   "forward:SetFa,SetFs,SetFn"
	prefix   *PrefixImpl "forward:Disable,ShowOnly,Highlight,Transparent"
}

func NewCylinder(h, r float64) *Cylinder {
	return &Cylinder{
		H:        h,
		RTop:     r,
		RBottom:  r,
		Center:   true,
		Circular: &Circular{},
		prefix:   &PrefixImpl{},
	}
}

func (o *Cylinder) Render(w *bufio.Writer) {
	w.WriteString(o.prefix.String())
	w.WriteString(
		fmt.Sprintf(
			"cylinder(h=%f, r1=%f, r2=%f, center=%t%s);\n",
			o.H, o.RBottom, o.RTop, o.Center, o.Circular.String(),
		),
	)
}
