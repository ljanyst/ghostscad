// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Polyhedron struct {
	ParentImpl
	Points    []Vec3
	Faces     []Vec3
	Convexity int         "optional"
	prefix    *PrefixImpl "forward:Disable,ShowOnly,Highlight,Transparent"
}

func NewPolyhedron(points []Vec3, faces []Vec3) *Polyhedron {
	return &Polyhedron{
		Points:    points,
		Faces:     faces,
		Convexity: 0,
		prefix:    &PrefixImpl{},
	}
}

func (o *Polyhedron) Render(w *bufio.Writer) {
	w.WriteString(o.prefix.String())
	w.WriteString("polyhedron(points=[")
	for i, pt := range o.Points {
		w.WriteString(fmt.Sprintf("[%f, %f, %f]", pt[0], pt[1], pt[2]))
		if i < len(o.Points)-1 {
			w.WriteString(", ")
		}
	}
	w.WriteString("], faces=[")
	for i, fc := range o.Faces {
		w.WriteString(fmt.Sprintf("[%f, %f, %f]", fc[0], fc[1], fc[2]))
		if i < len(o.Faces)-1 {
			w.WriteString(", ")
		}
	}
	w.WriteString("]")
	if o.Convexity != 0 {
		w.WriteString(fmt.Sprintf(", convexity=%d", o.Convexity))
	}
	w.WriteString(");\n")
}
