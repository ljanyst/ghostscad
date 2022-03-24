package main

import (
	"github.com/ljanyst/ghostscad/lib/shapes"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/ljanyst/ghostscad/primitive"
)

func main() {
	sector := shapes.NewSector(20, 45, 135).SetFn(72).Build()
	arc := shapes.NewArc(25, 45, 290).SetWidth(2).SetFn(72).Build()
	sys.RenderMultiple([]sys.Shape{
		{"sector", sector, sys.None},
		{"arc", arc, sys.None},
		{"sector-and-arc", NewList(sector, arc), sys.Default},
	})
}
