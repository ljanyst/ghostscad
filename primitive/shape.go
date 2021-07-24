package shape

import (
	"bufio"
)

type Shape interface {
	Render(w *bufio.Writer)
}
