// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package degmath

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

func Atan(x float64) float64 {
	return mgl64.RadToDeg(math.Atan(x))
}

func Atan2(y, x float64) float64 {
	return mgl64.RadToDeg(math.Atan2(y, x))
}

func Sin(x float64) float64 {
	return math.Sin(mgl64.DegToRad(x))
}

func Cos(x float64) float64 {
	return math.Cos(mgl64.DegToRad(x))
}
