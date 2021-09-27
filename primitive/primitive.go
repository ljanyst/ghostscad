// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

//go:generate go run github.com/ljanyst/ghostscad/util/tagged_method_generator -types Cube,Cylinder,Import,LinearExtrusion,Scale,Sphere,Surface

import (
	"bufio"
)

type Primitive interface {
	Render(w *bufio.Writer)
}
