package primitive

import (
	"bufio"
	"fmt"
)

type LinearExtrusion struct {
	Height    float64
	Center    bool
	Convexity uint16
	Twist     uint16
	Slices    uint16
	Scale     float64
	Fn        uint16
	Items     *List
}

func NewLinearExtrusion(height float64, items ...Primitive) *LinearExtrusion {
	return &LinearExtrusion{
		Height:    height,
		Center:    true,
		Convexity: 10,
		Slices:    20,
		Scale:     1.0,
		Fn:        16,
		Items:     NewList(items...),
	}
}

func (o *LinearExtrusion) SetCenter(center bool) *LinearExtrusion {
	o.Center = center
	return o
}

func (o *LinearExtrusion) SetConvexity(convexity uint16) *LinearExtrusion {
	o.Convexity = convexity
	return o
}

func (o *LinearExtrusion) SetTwist(twist uint16) *LinearExtrusion {
	o.Twist = twist
	return o
}

func (o *LinearExtrusion) SetSlices(slices uint16) *LinearExtrusion {
	o.Slices = slices
	return o
}

func (o *LinearExtrusion) SetScale(scale float64) *LinearExtrusion {
	o.Scale = scale
	return o
}

func (o *LinearExtrusion) SetFn(fn uint16) *LinearExtrusion {
	o.Fn = fn
	return o
}

func (o *LinearExtrusion) Add(items ...Primitive) *LinearExtrusion {
	o.Items.Add(items...)
	return o
}

func (o *LinearExtrusion) Render(w *bufio.Writer) {
	w.WriteString("linear_extrude(")
	w.WriteString(fmt.Sprintf("height=%f", o.Height))
	w.WriteString(fmt.Sprintf(", center=%t", o.Center))
	w.WriteString(fmt.Sprintf(", convexity=%d", o.Convexity))
	w.WriteString(fmt.Sprintf(", twist=%d", o.Twist))
	w.WriteString(fmt.Sprintf(", slices=%d", o.Slices))
	w.WriteString(fmt.Sprintf(", scale=%f", o.Scale))
	w.WriteString(fmt.Sprintf(", $fn=%d", o.Fn))
	w.WriteString(")")
	o.Items.Render(w)
}
