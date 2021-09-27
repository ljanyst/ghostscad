// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bytes"
	"fmt"
)

type Circular struct {
	Fa    float64
	Fs    float64
	Fn    uint16
	FaSet bool
	FsSet bool
	FnSet bool
}

func (o *Circular) SetFa(val float64) {
	o.Fa = val
	o.FaSet = true
}

func (o *Circular) SetFs(val float64) {
	o.Fs = val
	o.FsSet = true
}

func (o *Circular) SetFn(val uint16) {
	o.Fn = val
	o.FnSet = true
}

func (o *Circular) String() string {
	var buffer bytes.Buffer

	if o.FaSet {
		buffer.WriteString(fmt.Sprintf(", $fa=%f", o.Fa))
	}
	if o.FsSet {
		buffer.WriteString(fmt.Sprintf(", $fs=%f", o.Fs))
	}
	if o.FnSet {
		buffer.WriteString(fmt.Sprintf(", $fn=%d", o.Fn))
	}

	return buffer.String()
}
