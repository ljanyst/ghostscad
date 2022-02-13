// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type RotationExtrusion struct {
	Parent
	Angle     float64   "optional"
	Convexity uint16    "optional"
	Circular  *Circular "forward:SetFa,SetFs,SetFn"
	Items     *List     "forward:Add"
}

func NewRotationExtrusion(items ...Primitive) *RotationExtrusion {
	ret := &RotationExtrusion{
		Angle:     360,
		Convexity: 10,
		Circular:  &Circular{},
		Items:     NewList(),
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *RotationExtrusion) Render(w *bufio.Writer) {
	w.WriteString("rotate_extrude(")
	w.WriteString(fmt.Sprintf("angle=%f", o.Angle))
	w.WriteString(fmt.Sprintf(", convexity=%d", o.Convexity))
	w.WriteString(o.Circular.String())
	w.WriteString(")")
	o.Items.Render(w)
}
