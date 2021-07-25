package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Cube struct {
	Dims   Vec3
	Center bool
}

func NewCube(dims Vec3) *Cube {
	return &Cube{
		Dims:   dims,
		Center: true,
	}
}

func (o *Cube) SetCenter(center bool) *Cube {
	o.Center = center
	return o
}

func (o *Cube) Render(w *bufio.Writer) {
	w.WriteString(
		fmt.Sprintf("cube([%f, %f, %f], center=%t);\n", o.Dims[0], o.Dims[1], o.Dims[2], o.Center),
	)
}
