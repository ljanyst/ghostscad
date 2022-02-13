// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
)

type List struct {
	ParentImpl
	Items  []Primitive
	prefix *PrefixImpl "forward:Disable,ShowOnly,Highlight,Transparent"
}

func NewList() *List {
	return &List{
		prefix: &PrefixImpl{},
	}
}

func (o *List) Add(items ...Primitive) *List {
	for _, item := range items {
		item.SetParent(o)
	}
	o.Items = append(o.Items, items...)
	return o
}

func (o *List) Render(w *bufio.Writer) {
	w.WriteString(o.prefix.String())
	w.WriteString("{\n")
	for _, item := range o.Items {
		item.Render(w)
	}
	w.WriteString("}\n")
}
