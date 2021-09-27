// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Scale struct {
	Scale Vec3
	Items *List "forward:Add"
}

func NewScale(scale Vec3, items ...Primitive) *Scale {
	return &Scale{
		Scale: scale,
		Items: NewList(items...),
	}
}

func (o *Scale) Render(w *bufio.Writer) {
	w.WriteString(
		fmt.Sprintf("scale([%f, %f, %f]) ", o.Scale[0], o.Scale[1], o.Scale[2]),
	)
	o.Items.Render(w)
}
