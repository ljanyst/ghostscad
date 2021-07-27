package primitive

import (
	"bufio"
	"fmt"
)

type Cylinder struct {
	H       float64
	RTop    float64
	RBottom float64
	Center  bool
}

func NewCylinder(h, r float64) *Cylinder {
	return &Cylinder{
		H:       h,
		RTop:    r,
		RBottom: r,
		Center:  true,
	}
}

func (o *Cylinder) SetCenter(center bool) *Cylinder {
	o.Center = center
	return o
}

func (o *Cylinder) SetRBottom(rBottom float64) *Cylinder {
	o.RBottom = rBottom
	return o
}

func (o *Cylinder) Render(w *bufio.Writer) {
	w.WriteString(
		fmt.Sprintf("cylinder(h=%f, r1=%f, r2=%f, center=%t);\n", o.H, o.RBottom, o.RTop, o.Center),
	)
}
