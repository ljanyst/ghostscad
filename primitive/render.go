// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Render struct {
	ParentImpl
	Convexity uint16
	Items     *List  "forward:Add"
	prefix    string "prefix"
}

func NewRender(convexity uint16, items ...Primitive) *Render {
	ret := &Render{
		Convexity: convexity,
		Items:     NewList(),
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *Render) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString(
		fmt.Sprintf("render(convexity = %d)", o.Convexity),
	)
	o.Items.Render(w)
}
