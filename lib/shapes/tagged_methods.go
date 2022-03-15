// Do not edit. Automatically generated on Tue Mar 15 20:17:33 CET 2022.

package shapes

import ()

func (o *Arc) SetWidth(val float64) *Arc {
	o.Width = val
	return o
}

func (o *Arc) SetFn(val uint16) *Arc {
	o.Fn = val
	return o
}

func (o *Bezier) SetFn(val uint16) *Bezier {
	o.Fn = val
	return o
}

func (o *Graph) SetThickness(val float64) *Graph {
	o.Thickness = val
	return o
}

func (o *Graph) SetConvexity(val int) *Graph {
	o.Convexity = val
	return o
}

func (o *Polyline) SetRound(val bool) *Polyline {
	o.Round = val
	return o
}

func (o *Polyline3d) SetFn(val uint16) *Polyline3d {
	o.Fn = val
	return o
}

func (o *Ring) SetStartAngle(val float64) *Ring {
	o.StartAngle = val
	return o
}

func (o *Ring) SetRingAngle(val float64) *Ring {
	o.RingAngle = val
	return o
}

func (o *Sector) SetFn(val uint16) *Sector {
	o.Fn = val
	return o
}

func (o *SmootedCube) SetCenter(val bool) *SmootedCube {
	o.Center = val
	return o
}
