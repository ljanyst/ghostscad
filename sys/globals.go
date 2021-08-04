package sys

import (
	"bufio"
	"fmt"
)

var fa float64 = 2.0
var fs float64 = 2.0
var fn uint16 = 0

func init() {
	// According to the docs:
	// https://en.wikibooks.org/wiki/OpenSCAD_User_Manual/Other_Language_Features#.24fa.2C_.24fs_and_.24fn
	fa = 12.0
	fs = 2.0
	fn = 0
}

// Minimum angle for a fragment. Ignored if number of fragments setting is non-zero.
func SetFa(val float64) {
	if val < 0.01 {
		val = 0.01
	}
	fa = val
}

// Minimum size of a fragment. Ignored if number of fragments setting is non-zero.
func SetFs(val float64) {
	if val < 0.01 {
		val = 0.01
	}
	fa = val

}

// Number of fragments for the full circle. If zero other fragment settings apply.
func SetFn(val uint16) {
	fn = val
}

func renderGlobals(w *bufio.Writer) {
	w.WriteString(fmt.Sprintf("$fa=%f;\n", fa))
	w.WriteString(fmt.Sprintf("$fs=%f;\n", fs))
	w.WriteString(fmt.Sprintf("$fn=%d;\n", fn))
}
