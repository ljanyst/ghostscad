package sys

import (
	"bufio"
	"flag"
	"os"

	. "github.com/ljanyst/ghostscad/primitive"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func RenderAndExit(shape Shape) {
	out := flag.String("out", "out.scad", "output OpenSCAD file")
	logFile := flag.String("log-file", "", "output file for diagnostics")
	logLevel := flag.String("log-level", "Info", "verbosity of the diagnostic information")
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

	// Open the output file
	f, err := os.OpenFile(*out, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(f)
	shape.Render(writer)
	writer.Flush()
}
