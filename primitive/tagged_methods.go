// Do not edit. Automatically generated on Thu Feb 10 20:06:35 CET 2022.

package primitive

import ()

func (o *Cube) SetCenter(val bool) *Cube {
	o.Center = val
	return o
}

func (o *Cylinder) SetRBottom(val float64) *Cylinder {
	o.RBottom = val
	return o
}

func (o *Cylinder) SetCenter(val bool) *Cylinder {
	o.Center = val
	return o
}

func (o *Import) SetConvexity(val int) *Import {
	o.Convexity = val
	return o
}

func (o *Import) SetLayer(val string) *Import {
	o.Layer = val
	return o
}

func (o *LinearExtrusion) SetCenter(val bool) *LinearExtrusion {
	o.Center = val
	return o
}

func (o *LinearExtrusion) SetConvexity(val uint16) *LinearExtrusion {
	o.Convexity = val
	return o
}

func (o *LinearExtrusion) SetTwist(val uint16) *LinearExtrusion {
	o.Twist = val
	return o
}

func (o *LinearExtrusion) SetSlices(val uint16) *LinearExtrusion {
	o.Slices = val
	return o
}

func (o *LinearExtrusion) SetScale(val float64) *LinearExtrusion {
	o.Scale = val
	return o
}

func (o *LinearExtrusion) SetFn(val uint16) *LinearExtrusion {
	o.Fn = val
	return o
}

func (o *Surface) SetCenter(val bool) *Surface {
	o.Center = val
	return o
}

func (o *Surface) SetInvert(val bool) *Surface {
	o.Invert = val
	return o
}

func (o *Surface) SetConvexity(val int) *Surface {
	o.Convexity = val
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

func (o *LinearExtrusion) Add(items ...Primitive) *LinearExtrusion {
	o.Items.Add(items...)
	return o
}

func (o *Scale) Add(items ...Primitive) *Scale {
	o.Items.Add(items...)
	return o
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
