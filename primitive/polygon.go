// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Polygon struct {
	ParentImpl
	Points    []Vec2
	Paths     [][]int "optional"
	Convexity int     "optional"
	prefix    string  "prefix"
}

func NewPolygon(points []Vec2) *Polygon {
	return &Polygon{
		Points: points,
	}
}

func (o *Polygon) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString("polygon(points=[")
	for i, pt := range o.Points {
		w.WriteString(fmt.Sprintf("[%f, %f]", pt[0], pt[1]))
		if i < len(o.Points)-1 {
			w.WriteString(", ")
		}
	}
	w.WriteString("])")
	if o.Paths != nil && len(o.Paths) > 0 {
		w.WriteString(", paths=[")
		for i, path := range o.Paths {
			w.WriteString("[")
			for j, pt := range path {
				w.WriteString(fmt.Sprintf("%d", pt))
				if j < len(path)-1 {
					w.WriteString(", ")
				}
			}
			w.WriteString("]")
			if i < len(o.Paths)-1 {
				w.WriteString(", ")
			}
		}
		w.WriteString("]")
	}
	if o.Convexity != 0 {
		w.WriteString(fmt.Sprintf(", convexity=%d", o.Convexity))
	}
	w.WriteString(");\n")
}
