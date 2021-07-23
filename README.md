
GhostSCAD
=========

GhostSCAD is a piece of software that makes it easy to create CAD models in
using [Go][golang] and compile them to the [OpenSCAD][openscad] language. It
allows you to use the full power of a real programming language and still use
the rendering capabilities of OpenSCAD.

Example
-------

GhostSCAD aims to minimize the boilerplate but, since it uses a general purpose
programming language, there is still some. Here's an example rendering a sphere
of radius 10 and producing the appropriate OpenSCAD model in the `out.scad`
file:

```golang
package main

import (
	. "github.com/ljanyst/ghostscad/shape"
	"github.com/ljanyst/ghostscad/sys"
)

func main() {
	sys.RenderAndExit(Sphere(10))
}
```

There's a whole bunch of other examples in this code repository. The author
leaves it as an exercise to the reader to find them. :)

Rendering Automation
--------------------

One of the cool features of OpenSCAD is automatic re-rendering triggered by
changes to the underlying SCAD files. Since GhostSCAD leverages the Go
programming language, it also makes sense to use it in a fully-fledged code
editor. Here's an Elisp function to compile and run the current go buffer in
[Emacs][emacs]:

```elisp
(defun go-run-this-file ()
  "go run"
  (interactive)
  (save-buffer)
  (shell-command (format "go run %s" (buffer-file-name))))
```

You can then bind it to a key combination in your `go-mode` hook like this:

```elisp
(local-set-key (kbd "C-c C-r") 'go-run-this-file)
```

After you set this up, hitting "CTRL-C CTRL-R" compbination runs the code in the
current go buffer and overwrites the `out.scad` file with the most recent
version of your model. If you have OpenSCAD rendering this file, it will redraw
your model automagically.

Happy hacking!

[golang]: https://golang.org/
[openscad]: https://openscad.org/
[emacs]: https://www.gnu.org/software/emacs/
