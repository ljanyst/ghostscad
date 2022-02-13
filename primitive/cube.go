// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Cube struct {
	ParentImpl
	Dims   Vec3
	Center bool   "optional"
	prefix string "prefix"
}

func NewCube(dims Vec3) *Cube {
	return &Cube{
		Dims:   dims,
		Center: true,
	}
}

func (o *Cube) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString(
		fmt.Sprintf("cube([%f, %f, %f], center=%t);\n", o.Dims[0], o.Dims[1], o.Dims[2], o.Center),
	)
}
