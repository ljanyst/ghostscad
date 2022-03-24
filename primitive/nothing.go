// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
)

type Nothing struct {
	ParentImpl
	prefix string "prefix"
}

func NewNothing() *Nothing {
	return &Nothing{}
}

func (o *Nothing) Render(w *bufio.Writer) {
	w.WriteString("/* Nothing */\n")
}
