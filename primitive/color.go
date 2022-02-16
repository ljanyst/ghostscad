// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Color struct {
	ParentImpl
	Color  string
	Items  *List  "forward:Add"
	prefix string "prefix"
}

func NewColor(color string, items ...Primitive) *Color {
	ret := &Color{
		Color: color,
		Items: NewList(),
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func NewColorV(r, g, b, a byte, items ...Primitive) *Color {
	return NewColor(fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a), items...)
}

func (o *Color) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString(fmt.Sprintf("color(%q)", o.Color))
	o.Items.Render(w)
}
