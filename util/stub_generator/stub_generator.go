// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var stub = `
package main

import (
        // "github.com/ljanyst/ghostscad/lib/shapes"
        "github.com/ljanyst/ghostscad/sys"

        // . "github.com/go-gl/mathgl/mgl64"
        . "github.com/ljanyst/ghostscad/primitive"
)

func main() {
        sys.SetFn(120)
        sys.RenderOne(NewSphere(10))
}
`

func main() {
	fileName := flag.String("file-name", "main.go", "name of the output file to put the code in")
	flag.Parse()

	file, err := os.OpenFile(*fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("[!] Error while opening file: %s", err)
		os.Exit(1)
	}

	n, err := file.WriteString(stub)
	if n != len(stub) {
		file.Close()
		fmt.Printf("[!] Managed to writy %d chars out of %d", n, len(stub))
		os.Exit(1)
	}

	if err != nil {
		file.Close()
		fmt.Printf("[!] Error while writing: %s", err)
		os.Exit(1)
	}
	file.Sync()
	file.Close()

	cmd := exec.Command("go", "fmt", *fileName)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		fmt.Printf("[!] Error while formatting: %s", err)
		os.Exit(1)
	}
}
