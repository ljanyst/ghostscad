package primitive

import (
	"bufio"
	"fmt"
)

type Sphere struct {
	Radius float64
}

func NewSphere(radius float64) *Sphere {
	return &Sphere{
		Radius: radius,
	}
}

func (o *Sphere) Render(w *bufio.Writer) {
	w.WriteString(fmt.Sprintf("sphere(r=%f);\n", o.Radius))
}
