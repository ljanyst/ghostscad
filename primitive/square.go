// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Square struct {
	ParentImpl
	Dims   Vec2
	Center bool   "optional"
	prefix string "prefix"
}

func NewSquare(dims Vec2) *Square {
	return &Square{
		Dims:   dims,
		Center: true,
	}
}

func (o *Square) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString(
		fmt.Sprintf("square([%f, %f], center=%t);\n", o.Dims[0], o.Dims[1], o.Center),
	)
}
