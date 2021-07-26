package primitive

import (
	"bufio"
	"fmt"
)

type Import struct {
	File      string
	Convexity int
	Layer     string
}

func NewImport(file string) *Import {
	return &Import{
		File:      file,
		Convexity: 0,
		Layer:     "",
	}
}

func (o *Import) SetConvexity(convexity int) *Import {
	o.Convexity = convexity
	return o
}

func (o *Import) SetLayer(layer string) *Import {
	o.Layer = layer
	return o
}

func (o *Import) Render(w *bufio.Writer) {
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
