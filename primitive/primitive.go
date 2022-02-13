// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

//go:generate go run github.com/ljanyst/ghostscad/util/tagged_method_generator

import (
	"bufio"
)

type Primitive interface {
	SetParent(Primitive)
	Parent() Primitive
	Render(w *bufio.Writer)
}

type ParentImpl struct {
	parent Primitive
}

type PrefixImpl struct {
	prefix string
}

func (o *ParentImpl) SetParent(p Primitive) {
	o.parent = p
}

func (o *ParentImpl) Parent() Primitive {
	return o.parent
}

func (o *PrefixImpl) Disable() {
	o.prefix = "*"
}

func (o *PrefixImpl) ShowOnly() {
	o.prefix = "!"
}

func (o *PrefixImpl) Highlight() {
	o.prefix = "#"
}

func (o *PrefixImpl) Transparent() {
	o.prefix = "%"
}

func (o *PrefixImpl) String() string {
	return o.prefix
}
