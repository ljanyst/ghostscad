package main

import (
	. "github.com/ljanyst/ghostscad/shape"
	"github.com/ljanyst/ghostscad/sys"
)

func main() {
	sys.RenderAndExit(Sphere(10))
}
