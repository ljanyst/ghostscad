// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
)

type ListOp struct {
	ParentImpl
	Items  *List "forward:Add"
	name   string
	prefix string "prefix"
}

func newListOp(name string, items ...Primitive) *ListOp {
	ret := &ListOp{
		Items: NewList(),
		name:  name,
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func NewIntersection(items ...Primitive) *ListOp {
	return newListOp("intersection", items...)
}

func NewUnion(items ...Primitive) *ListOp {
	return newListOp("union", items...)
}

func NewDifference(items ...Primitive) *ListOp {
	return newListOp("difference", items...)
}

func (o *ListOp) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString(o.name)
	w.WriteString("()")
	o.Items.Render(w)
}
