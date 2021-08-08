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
	Circular
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

func (o *Cylinder) SetFa(val float64) *Cylinder {
	o.Circular.SetFa(val)
	return o
}

func (o *Cylinder) SetFs(val float64) *Cylinder {
	o.Circular.SetFs(val)
	return o
}

func (o *Cylinder) SetFn(val uint16) *Cylinder {
	o.Circular.SetFn(val)
	return o
}

func (o *Cylinder) Render(w *bufio.Writer) {
	w.WriteString(
		fmt.Sprintf(
			"cylinder(h=%f, r1=%f, r2=%f, center=%t%s);\n",
			o.H, o.RBottom, o.RTop, o.Center, o.Circular.String(),
		),
	)
}
