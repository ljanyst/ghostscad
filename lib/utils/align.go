// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package utils

import (
	. "github.com/ljanyst/ghostscad/primitive"
)

// Returns a transform aligning Anchor b with anchor a
func Align(a *Anchor, b *Anchor) *Transform {
	trA := a.Transform()
	trB := b.Transform()
	trA.Add(trB.Inverse())
	return trA
}
