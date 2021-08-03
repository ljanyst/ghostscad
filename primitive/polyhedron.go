package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Polyhedron struct {
	Points []Vec3
	Faces  []Vec3
}

func NewPolyhedron(points []Vec3, faces []Vec3) *Polyhedron {
	return &Polyhedron{
		Points: points,
		Faces:  faces,
	}
}

func (o *Polyhedron) Render(w *bufio.Writer) {
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

	w.WriteString("]);\n")
}
