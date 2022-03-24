// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
)

// List groups operations. It renders curly brackets in the SCAD file.
type List struct {
	ParentImpl
	Items  []Primitive
	prefix string "prefix"
}

func NewList(items ...Primitive) *List {
	ret := &List{}
	ret.Add(items...)
	return ret
}

func (o *List) Add(items ...Primitive) *List {
	for _, item := range items {
		item.SetParent(o)
	}
	o.Items = append(o.Items, items...)
	return o
}

func (o *List) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString("{\n")
	for _, item := range o.Items {
		item.Render(w)
	}
	w.WriteString("}\n")
}
