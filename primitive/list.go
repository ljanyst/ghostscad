// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
)

type List struct {
	Parent
	Items []Primitive
}

func NewList() *List {
	return &List{}
}

func (o *List) Add(items ...Primitive) *List {
	for _, item := range items {
		item.SetParent(o)
	}
	o.Items = append(o.Items, items...)
	return o
}

func (o *List) Render(w *bufio.Writer) {
	w.WriteString("{\n")
	for _, item := range o.Items {
		item.Render(w)
	}
	w.WriteString("}\n")
}
