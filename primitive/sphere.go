package primitive

import (
	"bufio"
	"fmt"
)

type SphereNode struct {
	Radius float64
}

func Sphere(radius float64) *SphereNode {
	return &SphereNode{
		Radius: radius,
	}
}

func (o *SphereNode) Render(w *bufio.Writer) {
	w.WriteString(fmt.Sprintf("sphere(r=%f);\n", o.Radius))
}
