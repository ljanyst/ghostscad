package main

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
	"github.com/ljanyst/ghostscad/sys"
)

func newSerpentine() Primitive {
	return NewLinearExtrusion(
		50,
		NewTranslation(
			Vec3{4.5, 0, 0},
			NewCircle(1)),
	).SetFn(96).SetTwist(720).SetSlices(100)
}

func main() {
	sys.SetFn(360)
	sys.RenderAndExit(newSerpentine())
}
