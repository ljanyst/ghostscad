// Do not edit. Automatically generated on Sun Feb 13 16:30:48 CET 2022.

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

func (o *RotationExtrusion) SetAngle(val float64) *RotationExtrusion {
	o.Angle = val
	return o
}

func (o *RotationExtrusion) SetConvexity(val uint16) *RotationExtrusion {
	o.Convexity = val
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

func (o *Cube) Disable() *Cube {
	o.prefix.Disable()
	return o
}

func (o *Cube) ShowOnly() *Cube {
	o.prefix.ShowOnly()
	return o
}

func (o *Cube) Highlight() *Cube {
	o.prefix.Highlight()
	return o
}

func (o *Cube) Transparent() *Cube {
	o.prefix.Transparent()
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

func (o *Cylinder) Disable() *Cylinder {
	o.prefix.Disable()
	return o
}

func (o *Cylinder) ShowOnly() *Cylinder {
	o.prefix.ShowOnly()
	return o
}

func (o *Cylinder) Highlight() *Cylinder {
	o.prefix.Highlight()
	return o
}

func (o *Cylinder) Transparent() *Cylinder {
	o.prefix.Transparent()
	return o
}

func (o *Import) Disable() *Import {
	o.prefix.Disable()
	return o
}

func (o *Import) ShowOnly() *Import {
	o.prefix.ShowOnly()
	return o
}

func (o *Import) Highlight() *Import {
	o.prefix.Highlight()
	return o
}

func (o *Import) Transparent() *Import {
	o.prefix.Transparent()
	return o
}

func (o *LinearExtrusion) Add(items ...Primitive) *LinearExtrusion {
	o.Items.Add(items...)
	return o
}

func (o *LinearExtrusion) Disable() *LinearExtrusion {
	o.prefix.Disable()
	return o
}

func (o *LinearExtrusion) ShowOnly() *LinearExtrusion {
	o.prefix.ShowOnly()
	return o
}

func (o *LinearExtrusion) Highlight() *LinearExtrusion {
	o.prefix.Highlight()
	return o
}

func (o *LinearExtrusion) Transparent() *LinearExtrusion {
	o.prefix.Transparent()
	return o
}

func (o *List) Disable() *List {
	o.prefix.Disable()
	return o
}

func (o *List) ShowOnly() *List {
	o.prefix.ShowOnly()
	return o
}

func (o *List) Highlight() *List {
	o.prefix.Highlight()
	return o
}

func (o *List) Transparent() *List {
	o.prefix.Transparent()
	return o
}

func (o *Polyhedron) Disable() *Polyhedron {
	o.prefix.Disable()
	return o
}

func (o *Polyhedron) ShowOnly() *Polyhedron {
	o.prefix.ShowOnly()
	return o
}

func (o *Polyhedron) Highlight() *Polyhedron {
	o.prefix.Highlight()
	return o
}

func (o *Polyhedron) Transparent() *Polyhedron {
	o.prefix.Transparent()
	return o
}

func (o *RotationExtrusion) SetFa(val float64) *RotationExtrusion {
	o.Circular.SetFa(val)
	return o
}

func (o *RotationExtrusion) SetFs(val float64) *RotationExtrusion {
	o.Circular.SetFs(val)
	return o
}

func (o *RotationExtrusion) SetFn(val uint16) *RotationExtrusion {
	o.Circular.SetFn(val)
	return o
}

func (o *RotationExtrusion) Add(items ...Primitive) *RotationExtrusion {
	o.Items.Add(items...)
	return o
}

func (o *RotationExtrusion) Disable() *RotationExtrusion {
	o.prefix.Disable()
	return o
}

func (o *RotationExtrusion) ShowOnly() *RotationExtrusion {
	o.prefix.ShowOnly()
	return o
}

func (o *RotationExtrusion) Highlight() *RotationExtrusion {
	o.prefix.Highlight()
	return o
}

func (o *RotationExtrusion) Transparent() *RotationExtrusion {
	o.prefix.Transparent()
	return o
}

func (o *Scale) Add(items ...Primitive) *Scale {
	o.Items.Add(items...)
	return o
}

func (o *Scale) Disable() *Scale {
	o.prefix.Disable()
	return o
}

func (o *Scale) ShowOnly() *Scale {
	o.prefix.ShowOnly()
	return o
}

func (o *Scale) Highlight() *Scale {
	o.prefix.Highlight()
	return o
}

func (o *Scale) Transparent() *Scale {
	o.prefix.Transparent()
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

func (o *Sphere) Disable() *Sphere {
	o.prefix.Disable()
	return o
}

func (o *Sphere) ShowOnly() *Sphere {
	o.prefix.ShowOnly()
	return o
}

func (o *Sphere) Highlight() *Sphere {
	o.prefix.Highlight()
	return o
}

func (o *Sphere) Transparent() *Sphere {
	o.prefix.Transparent()
	return o
}

func (o *Surface) Disable() *Surface {
	o.prefix.Disable()
	return o
}

func (o *Surface) ShowOnly() *Surface {
	o.prefix.ShowOnly()
	return o
}

func (o *Surface) Highlight() *Surface {
	o.prefix.Highlight()
	return o
}

func (o *Surface) Transparent() *Surface {
	o.prefix.Transparent()
	return o
}

func (o *Translation) Add(items ...Primitive) *Translation {
	o.Items.Add(items...)
	return o
}

func (o *Translation) Disable() *Translation {
	o.prefix.Disable()
	return o
}

func (o *Translation) ShowOnly() *Translation {
	o.prefix.ShowOnly()
	return o
}

func (o *Translation) Highlight() *Translation {
	o.prefix.Highlight()
	return o
}

func (o *Translation) Transparent() *Translation {
	o.prefix.Transparent()
	return o
}
