// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Custom struct {
	ParentImpl
	Code   string
	prefix string "prefix"
}

func NewCustom(code string) *Custom {
	return &Custom{
		Code: code,
	}
}

func (o *Custom) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString(o.Code)
}
