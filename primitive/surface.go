// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Surface struct {
	ParentImpl
	File      string
	Center    bool   "optional"
	Invert    bool   "optional"
	Convexity int    "optional"
	prefix    string "prefix"
}

func NewSurface(file string) *Surface {
	return &Surface{
		File:   file,
		Center: true,
		Invert: false,
	}
}

func (o *Surface) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
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
