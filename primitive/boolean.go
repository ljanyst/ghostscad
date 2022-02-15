// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
)

type Boolean struct {
	ParentImpl
	Items  *List "forward:Add"
	name   string
	prefix string "prefix"
}

func newBoolean(name string, items ...Primitive) *Boolean {
	ret := &Boolean{
		Items: NewList(),
		name:  name,
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func NewIntersection(items ...Primitive) *Boolean {
	return newBoolean("intersection", items...)
}

func NewUnion(items ...Primitive) *Boolean {
	return newBoolean("union", items...)
}

func NewDifference(items ...Primitive) *Boolean {
	return newBoolean("difference", items...)
}

func (o *Boolean) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString(o.name)
	w.WriteString("()")
	o.Items.Render(w)
}
