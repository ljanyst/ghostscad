package primitive

import (
	"bufio"
)

type Primitive interface {
	Render(w *bufio.Writer)
}
