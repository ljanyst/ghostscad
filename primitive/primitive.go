// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

// Implementation of OpenSCAD primitives. Everything here mostly follows the
// naming and defaults as defined here: https://openscad.org/cheatsheet/
// Additionally, the following new primitives have been added: Anchor, Custom,
// List, Nothing. These are documented here.
package primitive

//go:generate go run github.com/ljanyst/ghostscad/util/tagged_method_generator

import (
	"bufio"
)

type Primitive interface {
	SetParent(Primitive)
	Parent() Primitive
	Disable() Primitive
	Highlight() Primitive
	ShowOnly() Primitive
	Transparent() Primitive
	Prefix() string
	Render(w *bufio.Writer)
}

type ParentImpl struct {
	parent Primitive
}

func (o *ParentImpl) SetParent(p Primitive) {
	o.parent = p
}

func (o *ParentImpl) Parent() Primitive {
	return o.parent
}
