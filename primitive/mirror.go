// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Mirror struct {
	ParentImpl
	Mirror Vec3
	Items  *List  "forward:Add"
	prefix string "prefix"
}

func NewMirror(mirror Vec3, items ...Primitive) *Mirror {
	ret := &Mirror{
		Mirror: mirror,
		Items:  NewList(),
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *Mirror) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString(
		fmt.Sprintf("mirror([%f, %f, %f]) ", o.Mirror[0], o.Mirror[1], o.Mirror[2]),
	)
	o.Items.Render(w)
}
