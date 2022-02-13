// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

//go:generate go run github.com/ljanyst/ghostscad/util/tagged_method_generator

import (
	"bufio"
)

type Primitive interface {
	SetParent(Primitive)
	GetParent() Primitive
	Render(w *bufio.Writer)
}

type Parent struct {
	parent Primitive
}

func (c *Parent) SetParent(p Primitive) {
	c.parent = p
}

func (c *Parent) GetParent() Primitive {
	return c.parent
}
