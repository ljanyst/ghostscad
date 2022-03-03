// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package sys

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"

	. "github.com/ljanyst/ghostscad/primitive"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	out        = flag.String("out", "", "output file")
	logFile    = flag.String("log-file", "", "output file for diagnostics")
	logLevel   = flag.String("log-level", "Info", "verbosity of the diagnostic information")
	listShapes = flag.Bool("list-shapes", false, "list the available shapes")
	shapeSel   = flag.String("shape", "", "shape to render if not default")
	stl        = flag.Bool("stl", false, "produce an STL file")
	all        = flag.Bool("all", false, "procrss all shapes")
)

func init() {
	flag.Parse()

	// Set up logging
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	if *logFile != "" {
		f, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	level := log.InfoLevel
	if *logLevel != "" {
		l, err := log.ParseLevel(*logLevel)
		if err != nil {
			log.Fatal(err)
		}
		level = l
	}
	log.SetLevel(level)
}

func computeOutputName(shapeName string) string {
	if *out != "" {
		return *out
	}

	if *stl {
		return fmt.Sprintf("%s.stl", shapeName)
	}
	return fmt.Sprintf("%s.scad", shapeName)
}

func render(shape Primitive, fileName string) {
	var f *os.File
	var err error
	if *stl {
		f, err = os.CreateTemp("", "tmpfile-")
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()
		defer os.Remove(f.Name())
	} else {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	writer := bufio.NewWriter(f)
	renderGlobals(writer)
	shape.Render(writer)
	writer.Flush()

	if *stl {
		err = exec.Command("openscad", f.Name(), "-o", fileName).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func RenderOne(shape Primitive) {
	if *listShapes {
		fmt.Println("main (default)")
		return
	}

	if *shapeSel != "" && *shapeSel != "main" {
		log.Fatal("No such shape: %q\n", *shapeSel)
	}

	render(shape, computeOutputName("main"))
}

func RenderMultiple(shapes map[string]Primitive, dflt string) {
	if *listShapes {
		for shape, _ := range shapes {
			fmt.Printf("%s", shape)
			if shape == dflt {
				fmt.Printf(" (default)")
			}
			fmt.Printf("\n")
		}
		return
	}

	if *all {
		for shapeName, shape := range shapes {
			fileName := computeOutputName(shapeName)
			log.Infof("Processing %s...", fileName)
			render(shape, fileName)
		}
		log.Info("Done")
	} else {
		shapeName := *shapeSel
		if shapeName == "" {
			shapeName = dflt
		}

		shape, ok := shapes[shapeName]
		if !ok {
			log.Fatal("No such shape: %q\n", shapeName)
		}
		render(shape, computeOutputName(shapeName))
	}
}
