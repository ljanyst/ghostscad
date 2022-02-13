// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Surface struct {
	Parent
	File      string
	Center    bool "optional"
	Invert    bool "optional"
	Convexity int  "optional"
}

func NewSurface(file string) *Surface {
	return &Surface{
		File:      file,
		Center:    true,
		Invert:    false,
		Convexity: 0,
	}
}

func (o *Surface) Render(w *bufio.Writer) {
	w.WriteString("surface(file=\"")
	w.WriteString(o.File)
	w.WriteString("\"")
	w.WriteString(fmt.Sprintf(", center=%t", o.Center))
	w.WriteString(fmt.Sprintf(", invert=%t", o.Invert))
	if o.Convexity != 0 {
		w.WriteString(fmt.Sprintf(", convexity=%d", o.Convexity))
	}
	w.WriteString(");\n")
}
