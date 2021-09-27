// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Translation struct {
	Offset Vec3
	Items  *List
}

func NewTranslation(offset Vec3, items ...Primitive) *Translation {
	return &Translation{
		Offset: offset,
		Items:  NewList(items...),
	}
}

func (o *Translation) Add(items ...Primitive) *Translation {
	o.Items.Add(items...)
	return o
}

func (o *Translation) Render(w *bufio.Writer) {
	w.WriteString(
		fmt.Sprintf("translate([%f, %f, %f]) ", o.Offset[0], o.Offset[1], o.Offset[2]),
	)
	o.Items.Render(w)
}
