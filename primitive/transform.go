// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type transformElement struct {
	name   string
	vector Vec3
	angle  Vec3
}

type Transform struct {
	ParentImpl
	Items      *List "forward:Add"
	transforms []transformElement
	prefix     string "prefix"
}

func NewTranslation(vector Vec3, items ...Primitive) *Transform {
	ret := &Transform{
		Items: NewList(),
		transforms: []transformElement{
			{"translate", vector, Vec3{}},
		},
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func NewRotation(angle Vec3, items ...Primitive) *Transform {
	ret := &Transform{
		Items: NewList(),
		transforms: []transformElement{
			{"rotate", Vec3{}, Vec3{0, 0, angle[2]}},
			{"rotate", Vec3{}, Vec3{0, angle[1], 0}},
			{"rotate", Vec3{}, Vec3{angle[0], 0, 0}},
		},
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func NewRotationByAxis(angle float64, vector Vec3, items ...Primitive) *Transform {
	ret := &Transform{
		Items: NewList(),
		transforms: []transformElement{
			{"rotateAngle", vector, Vec3{angle, 0, 0}},
		},
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *Transform) Inverse() *Transform {
	ret := &Transform{
		Items:      NewList(),
		transforms: make([]transformElement, len(o.transforms)),
	}
	ret.Items.SetParent(ret)

	copy(ret.transforms, o.transforms)

	for i := 0; i < len(ret.transforms); i++ {
		if ret.transforms[i].name == "translate" {
			ret.transforms[i].vector = ret.transforms[i].vector.Mul(-1)
		} else {
			ret.transforms[i].angle = ret.transforms[i].angle.Mul(-1)
		}
	}

	for i, j := 0, len(ret.transforms)-1; i < j; i, j = i+1, j-1 {
		ret.transforms[i], ret.transforms[j] = ret.transforms[j], ret.transforms[i]
	}

	return ret
}

func (o *Transform) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	for _, tf := range o.transforms {
		switch tf.name {
		case "translate":
			w.WriteString(
				fmt.Sprintf("translate([%f, %f, %f]) ", tf.vector[0], tf.vector[1], tf.vector[2]),
			)
		case "rotate":
			w.WriteString(
				fmt.Sprintf("rotate([%f, %f, %f]) ", tf.angle[0], tf.angle[1], tf.angle[2]),
			)
		case "rotateAngle":
			w.WriteString(
				fmt.Sprintf("rotate(a = %f, v = [%f, %f, %f]) ",
					tf.angle[0], tf.vector[0], tf.vector[1], tf.vector[2]),
			)
		}
	}
	o.Items.Render(w)
}
