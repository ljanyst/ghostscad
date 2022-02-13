// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type LinearExtrusion struct {
	ParentImpl
	Height    float64
	Center    bool        "optional"
	Convexity uint16      "optional"
	Twist     uint16      "optional"
	Slices    uint16      "optional"
	Scale     float64     "optional"
	Fn        uint16      "optional"
	Items     *List       "forward:Add"
	prefix    *PrefixImpl "forward:Disable,ShowOnly,Highlight,Transparent"
}

func NewLinearExtrusion(height float64, items ...Primitive) *LinearExtrusion {
	ret := &LinearExtrusion{
		Height:    height,
		Center:    true,
		Convexity: 10,
		Slices:    20,
		Scale:     1.0,
		Fn:        16,
		Items:     NewList(),
		prefix:    &PrefixImpl{},
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *LinearExtrusion) Render(w *bufio.Writer) {
	w.WriteString(o.prefix.String())
	w.WriteString("linear_extrude(")
	w.WriteString(fmt.Sprintf("height=%f", o.Height))
	w.WriteString(fmt.Sprintf(", center=%t", o.Center))
	w.WriteString(fmt.Sprintf(", convexity=%d", o.Convexity))
	w.WriteString(fmt.Sprintf(", twist=%d", o.Twist))
	w.WriteString(fmt.Sprintf(", slices=%d", o.Slices))
	w.WriteString(fmt.Sprintf(", scale=%f", o.Scale))
	w.WriteString(fmt.Sprintf(", $fn=%d", o.Fn))
	w.WriteString(")")
	o.Items.Render(w)
}
