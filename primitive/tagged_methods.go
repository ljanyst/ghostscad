// Do not edit. Automatically generated on Wed Feb 16 10:48:38 CET 2022.

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

func (o *Offset) SetR(val float64) *Offset {
	o.R = val
	return o
}

func (o *Offset) SetDelta(val float64) *Offset {
	o.Delta = val
	return o
}

func (o *Offset) SetChamfer(val bool) *Offset {
	o.Chamfer = val
	return o
}

func (o *Polygon) SetPaths(val [][]int) *Polygon {
	o.Paths = val
	return o
}

func (o *Polygon) SetConvexity(val int) *Polygon {
	o.Convexity = val
	return o
}

func (o *Polyhedron) SetConvexity(val int) *Polyhedron {
	o.Convexity = val
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

func (o *Square) SetCenter(val bool) *Square {
	o.Center = val
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

func (o *Text) SetSize(val uint16) *Text {
	o.Size = val
	return o
}

func (o *Text) SetFont(val string) *Text {
	o.Font = val
	return o
}

func (o *Text) SetHalign(val string) *Text {
	o.Halign = val
	return o
}

func (o *Text) SetValign(val string) *Text {
	o.Valign = val
	return o
}

func (o *Text) SetSpacing(val uint16) *Text {
	o.Spacing = val
	return o
}

func (o *Text) SetDirection(val string) *Text {
	o.Direction = val
	return o
}

func (o *Text) SetLanguage(val string) *Text {
	o.Language = val
	return o
}

func (o *Text) SetScript(val string) *Text {
	o.Script = val
	return o
}

func (o *Text) SetFn(val uint16) *Text {
	o.Fn = val
	return o
}

func (o *Boolean) Add(items ...Primitive) *Boolean {
	o.Items.Add(items...)
	return o
}

func (o *Color) Add(items ...Primitive) *Color {
	o.Items.Add(items...)
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

func (o *Offset) Add(items ...Primitive) *Offset {
	o.Items.Add(items...)
	return o
}

func (o *Render) Add(items ...Primitive) *Render {
	o.Items.Add(items...)
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

func (o *Transform) Add(items ...Primitive) *Transform {
	o.Items.Add(items...)
	return o
}

func (o *Boolean) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Boolean) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Boolean) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Boolean) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Boolean) Prefix() string {
	return o.prefix
}

func (o *Color) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Color) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Color) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Color) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Color) Prefix() string {
	return o.prefix
}

func (o *Cube) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Cube) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Cube) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Cube) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Cube) Prefix() string {
	return o.prefix
}

func (o *Custom) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Custom) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Custom) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Custom) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Custom) Prefix() string {
	return o.prefix
}

func (o *Cylinder) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Cylinder) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Cylinder) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Cylinder) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Cylinder) Prefix() string {
	return o.prefix
}

func (o *Import) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Import) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Import) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Import) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Import) Prefix() string {
	return o.prefix
}

func (o *LinearExtrusion) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *LinearExtrusion) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *LinearExtrusion) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *LinearExtrusion) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *LinearExtrusion) Prefix() string {
	return o.prefix
}

func (o *List) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *List) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *List) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *List) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *List) Prefix() string {
	return o.prefix
}

func (o *Offset) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Offset) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Offset) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Offset) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Offset) Prefix() string {
	return o.prefix
}

func (o *Polygon) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Polygon) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Polygon) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Polygon) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Polygon) Prefix() string {
	return o.prefix
}

func (o *Polyhedron) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Polyhedron) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Polyhedron) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Polyhedron) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Polyhedron) Prefix() string {
	return o.prefix
}

func (o *Render) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Render) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Render) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Render) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Render) Prefix() string {
	return o.prefix
}

func (o *RotationExtrusion) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *RotationExtrusion) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *RotationExtrusion) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *RotationExtrusion) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *RotationExtrusion) Prefix() string {
	return o.prefix
}

func (o *Scale) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Scale) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Scale) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Scale) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Scale) Prefix() string {
	return o.prefix
}

func (o *Sphere) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Sphere) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Sphere) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Sphere) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Sphere) Prefix() string {
	return o.prefix
}

func (o *Square) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Square) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Square) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Square) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Square) Prefix() string {
	return o.prefix
}

func (o *Surface) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Surface) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Surface) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Surface) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Surface) Prefix() string {
	return o.prefix
}

func (o *Text) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Text) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Text) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Text) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Text) Prefix() string {
	return o.prefix
}

func (o *Transform) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Transform) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Transform) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Transform) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Transform) Prefix() string {
	return o.prefix
}
