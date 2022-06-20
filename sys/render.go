// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package sys

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/ljanyst/ghostscad/primitive"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	out         = flag.String("out", "", "output file")
	logFile     = flag.String("log-file", "", "output file for diagnostics")
	logLevel    = flag.String("log-level", "Info", "verbosity of the diagnostic information")
	listShapes  = flag.Bool("list-shapes", false, "list the available shapes")
	shapeSel    = flag.String("shape", "", "shape to render if not default")
	stl         = flag.Bool("stl", false, "produce an STL file")
	all         = flag.Bool("all", false, "process all shapes")
	initialized = false
)

// Parse command-line flags and initialize the logging system
func Initialize() {
	if initialized {
		return
	} else {
		initialized = true
	}

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

//go:generate go run golang.org/x/tools/cmd/stringer -type ShapeFlags render.go

type ShapeFlags uint32

const (
	// No special treatment
	None ShapeFlags = 0

	// This is the default shape. The first encountered default shape will be
	// treated as such
	Default ShapeFlags = 1 << iota

	// Skip this shape when processing the shapes in bulk (the -all option)
	SkipInBulk
)

type Shape struct {
	Name      string
	Primitive Primitive
	Flags     ShapeFlags
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

func flagsToString(flags ShapeFlags) string {
	if flags == None {
		return ""
	}

	b := strings.Builder{}
	b.WriteString(" ")
	f := Default
	prev := false

	for i := 0; i < 32; i++ {
		if (flags & (f << i)) != 0 {
			if !prev {
				b.WriteString("(")
				prev = true
			} else {
				b.WriteString(", ")
			}
			b.WriteString((f << i).String())
		}
	}

	if prev {
		b.WriteString(")")
	}
	return b.String()
}

// Render a single shape in a way depending on commandline parameters
func RenderOne(shape Primitive) {
	Initialize()

	if *listShapes {
		fmt.Println("main (default)")
		return
	}

	if *shapeSel != "" && *shapeSel != "main" {
		log.Fatalf("No such shape: %q\n", *shapeSel)
	}

	render(shape, computeOutputName("main"))
}

// Render the shapes from the list in a way depending on commandline parameters
func RenderMultiple(shapes []Shape) {
	Initialize()

	if len(shapes) == 0 {
		log.Fatal("No defined shapes\n")
	}

	shapeMap := make(map[string]Shape)
	dflt := ""

	for _, shape := range shapes {
		shapeMap[shape.Name] = shape

		if *listShapes {
			fmt.Printf("%s", shape.Name)
			fmt.Printf("%s\n", flagsToString(shape.Flags))
		}

		if (shape.Flags&Default) != 0 && dflt == "" {
			dflt = shape.Name
		}
	}

	if *listShapes {
		return
	}

	if dflt == "" {
		log.Fatalf("Default shape not specifid. There needs to be one.\n")
	}

	if *all {
		for _, shape := range shapes {
			if (shape.Flags & SkipInBulk) == 0 {
				fileName := computeOutputName(shape.Name)
				log.Infof("Processing %s...", fileName)
				render(shape.Primitive, fileName)
			}
		}
		log.Info("Done")
	} else {
		shapeName := *shapeSel
		if shapeName == "" {
			shapeName = dflt
		}

		shape, ok := shapeMap[shapeName]
		if !ok {
			log.Fatalf("No such shape: %q\n", shapeName)
		}
		render(shape.Primitive, computeOutputName(shapeName))
	}
}
