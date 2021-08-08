package primitive

import (
	"bufio"
	"fmt"
)

type Sphere struct {
	Radius float64
	Circular
}

func NewSphere(radius float64) *Sphere {
	return &Sphere{
		Radius: radius,
	}
}

func (o *Sphere) SetFa(val float64) *Sphere {
	o.Circular.SetFa(val)
	return o
}

func (o *Sphere) SetFs(val float64) *Sphere {
	o.Circular.SetFs(val)
	return o
}

func (o *Sphere) SetFn(val uint16) *Sphere {
	o.Circular.SetFn(val)
	return o
}

func (o *Sphere) Render(w *bufio.Writer) {

	w.WriteString(fmt.Sprintf("sphere(r=%f%s);\n",
		o.Radius,
		o.Circular.String(),
	))
}
