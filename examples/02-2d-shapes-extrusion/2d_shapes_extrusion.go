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

func newRotation() Primitive {
	return NewTranslation(
		Vec3{-20, 0, -25},
		NewScale(
			Vec3{2, 2, 10},
			NewRotationExtrusion(
				NewPolygon([]Vec2{{0, 0}, {2, 1}, {1, 2}, {1, 3}, {3, 4}, {0, 5}}),
			).SetFn(200),
		))
}

func newComplexPolygon() Primitive {
	points := []Vec2{}
	paths := [][]int{}

	// main
	points = append(points, []Vec2{{0, 0}, {100, 0}, {130, 50}, {30, 50}}...)
	paths = append(paths, []int{1, 0, 3, 2})

	// hole 1
	points = append(points, []Vec2{{20, 20}, {40, 20}, {30, 30}}...)
	paths = append(paths, []int{4, 5, 6})

	// hole 2
	points = append(points, []Vec2{{50, 20}, {60, 20}, {40, 30}}...)
	paths = append(paths, []int{7, 8, 9})

	// hole 3
	points = append(points, []Vec2{{65, 10}, {80, 10}, {80, 40}, {65, 40}}...)
	paths = append(paths, []int{10, 11, 12, 13})

	// hole 4
	points = append(points, []Vec2{{98, 10}, {115, 40}, {85, 40}, {85, 10}}...)
	paths = append(paths, []int{14, 15, 16, 17})

	return NewTranslation(
		Vec3{20, 0, 0},
		NewRotation(
			Vec3{0, 0, 90},
			NewLinearExtrusion(
				50,
				NewScale(
					Vec3{0.25, 0.25, 1},
					NewTranslation(
						Vec3{-75, -25, 0},
						NewPolygon(points).SetPaths(paths),
					)))))
}

func newScrew() Primitive {
	return NewTranslation(
		Vec3{-40, 0, 0},
		NewLinearExtrusion(
			50,
			NewSquare(Vec2{10, 10})).SetTwist(180).SetConvexity(25).SetFn(96).SetSlices(250),
	)
}

func newText() Primitive {
	return NewTranslation(
		Vec3{40, -5, 0},
		NewLinearExtrusion(
			50,
			NewText("A")))
}

func main() {
	sys.SetFn(360)
	sys.RenderAndExit(NewList(
		newSerpentine(),
		newRotation(),
		newComplexPolygon(),
		newScrew(),
		newText(),
	))
}
