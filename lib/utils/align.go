// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package utils

import (
	. "github.com/ljanyst/ghostscad/primitive"
)

// Returns a transform aligning Anchor b with Anchor a
func Align(a *Anchor, b *Anchor) *Transform {
	trA := a.Transform()
	trB := b.Transform()
	trA.Append(trB.Inverse())
	return trA
}

// Returns a transform aligning origin with Anchor a
func AlignOrigin(a *Anchor) *Transform {
	return a.Transform()
}

// Returns a transform aligning Anchor a with origin
func AlignHere(a *Anchor) *Transform {
	return a.Transform().Inverse()
}
