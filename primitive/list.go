// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
)

type List struct {
	Items []Primitive
}

func NewList(items ...Primitive) *List {
	return &List{
		Items: items,
	}
}

func (o *List) Add(items ...Primitive) *List {
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
