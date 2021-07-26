package primitive

import (
	"bufio"
)

type List struct {
	Items []Primitive
}

func NewList(items ...Primitive) *List {
	return &List{
		Items: items,
	}
}

func (o *List) Add(items ...Primitive) *List {
	o.Items = append(o.Items, items...)
	return o
}

func (o *List) Render(w *bufio.Writer) {
	w.WriteString("{\n")
	for _, item := range o.Items {
		item.Render(w)
	}
	w.WriteString("}\n")
}
