// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Translation struct {
	ParentImpl
	Offset Vec3
	Items  *List       "forward:Add"
	prefix *PrefixImpl "forward:Disable,ShowOnly,Highlight,Transparent"
}

func NewTranslation(offset Vec3, items ...Primitive) *Translation {
	ret := &Translation{
		Offset: offset,
		Items:  NewList(),
		prefix: &PrefixImpl{},
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *Translation) Render(w *bufio.Writer) {
	w.WriteString(o.prefix.String())
	w.WriteString(
		fmt.Sprintf("translate([%f, %f, %f]) ", o.Offset[0], o.Offset[1], o.Offset[2]),
	)
	o.Items.Render(w)
}
