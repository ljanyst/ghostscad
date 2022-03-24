// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
)

// Anchor enables easy tracking of transforms and enables the alignment
// of reference frames. This is useful when you need to create or use complex
// objects that need to be put together in complex ways.
type Anchor struct {
	ParentImpl
	prefix string "prefix"
}

func NewAnchor() *Anchor {
	return &Anchor{}
}

func (o *Anchor) Transform() *Transform {
	ret := &Transform{
		Items: NewList(),
	}
	ret.Items.SetParent(ret)

	var scan func(Primitive)
	scan = func(p Primitive) {
		if p == nil {
			return
		}
		scan(p.Parent())

		tr, ok := p.(*Transform)
		if !ok {
			return
		}
		ret.transforms = append(ret.transforms, tr.transforms...)
	}
	scan(o.Parent())

	return ret
}

func (o *Anchor) Render(w *bufio.Writer) {
	w.WriteString("/* Anchor */\n")
}
