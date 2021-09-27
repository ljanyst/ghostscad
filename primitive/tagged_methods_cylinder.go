// Do not edit. Automatically generated on Mon Sep 27 15:08:09 CEST 2021.

package primitive

import ()

func (o *Cylinder) SetRBottom(val float64) *Cylinder {
	o.RBottom = val
	return o
}

func (o *Cylinder) SetCenter(val bool) *Cylinder {
	o.Center = val
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
