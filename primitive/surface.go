// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Surface struct {
	File      string
	Center    bool
	Invert    bool
	Convexity int
}

func NewSurface(file string) *Surface {
	return &Surface{
		File:      file,
		Center:    true,
		Invert:    false,
		Convexity: 0,
	}
}

func (o *Surface) SetCenter(center bool) *Surface {
	o.Center = center
	return o
}

func (o *Surface) SetInvert(invert bool) *Surface {
	o.Invert = invert
	return o
}

func (o *Surface) SetConvexity(convexity int) *Surface {
	o.Convexity = convexity
	return o
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
