// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Offset struct {
	ParentImpl
	R       float64 "optional"
	Delta   float64 "optional"
	Chamfer bool    "optional"
	Items   *List   "forward:Add"
	prefix  string  "prefix"
}

func NewOffset(items ...Primitive) *Offset {
	ret := &Offset{
		Items: NewList(),
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *Offset) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString("offset(")
	if o.R != 0 {
		w.WriteString(fmt.Sprintf("r = %f", o.R))
	} else if o.Delta != 0 {
		w.WriteString(fmt.Sprintf("delta = %f", o.Delta))
		if o.Chamfer {
			w.WriteString(fmt.Sprintf(", chamfer=%t", o.Chamfer))
		}
	}
	w.WriteString(")")
	o.Items.Render(w)
}
