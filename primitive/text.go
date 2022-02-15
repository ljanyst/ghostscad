// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type Text struct {
	ParentImpl
	Text      string
	Size      uint16 "optional"
	Font      string "optional"
	Halign    string "optional"
	Valign    string "optional"
	Spacing   uint16 "optional"
	Direction string "optional"
	Language  string "optional"
	Script    string "optional"
	Fn        uint16 "optional"
	prefix    string "prefix"
}

func NewText(text string) *Text {
	return &Text{
		Text:      text,
		Size:      10,
		Halign:    "left",
		Valign:    "baseline",
		Spacing:   1,
		Direction: "ltr",
		Language:  "en",
		Script:    "latin",
	}
}

func (o *Text) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString("text(")
	w.WriteString(fmt.Sprintf("%q", o.Text))
	w.WriteString(fmt.Sprintf(", size=%d", o.Size))
	w.WriteString(fmt.Sprintf(", font=%q", o.Font))
	w.WriteString(fmt.Sprintf(", halign=%q", o.Halign))
	w.WriteString(fmt.Sprintf(", valign=%q", o.Valign))
	w.WriteString(fmt.Sprintf(", spacing=%d", o.Spacing))
	w.WriteString(fmt.Sprintf(", direction=%q", o.Direction))
	w.WriteString(fmt.Sprintf(", language=%q", o.Language))
	w.WriteString(fmt.Sprintf(", script=%q", o.Script))
	w.WriteString(fmt.Sprintf(", $fn=%d", o.Fn))
	w.WriteString(");")
}
