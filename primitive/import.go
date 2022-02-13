// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Import struct {
	ParentImpl
	File      string
	Convexity int         "optional"
	Layer     string      "optional"
	prefix    *PrefixImpl "forward:Disable,ShowOnly,Highlight,Transparent"
}

func NewImport(file string) *Import {
	return &Import{
		File:      file,
		Convexity: 0,
		Layer:     "",
		prefix:    &PrefixImpl{},
	}
}

func (o *Import) Render(w *bufio.Writer) {
	w.WriteString(o.prefix.String())
	w.WriteString("import(\"")
	w.WriteString(o.File)
	w.WriteString("\"")
	if o.Convexity != 0 {
		w.WriteString(fmt.Sprintf(", convexity=%d", o.Convexity))
	}
	if o.Layer != "" {
		w.WriteString(fmt.Sprintf(", layer=%q", o.Layer))
	}
	w.WriteString(");\n")
}
