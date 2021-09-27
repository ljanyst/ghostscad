// Do not edit. Automatically generated on Mon Sep 27 15:08:09 CEST 2021.

package primitive

import ()

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

func (o *LinearExtrusion) Add(items ...Primitive) *LinearExtrusion {
	o.Items.Add(items...)
	return o
}
