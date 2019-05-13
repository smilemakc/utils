package logging

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

type DebugWriter struct {
	io.Writer
}

type InfoWriter struct {
	io.Writer
}

type ErrorWriter struct {
	io.Writer
}

func (d DebugWriter) Write(p []byte) (n int, err error) {
	if _debug {
		color.Set(color.FgGreen)
		defer color.Unset()
		return os.Stdout.Write(p)
	}
	return 0, nil
}

func (i InfoWriter) Write(p []byte) (n int, err error) {
	if _debug {
		color.Set(color.FgBlue)
		defer color.Unset()
		return os.Stdout.Write(p)
	}
	return 0, nil
}

func (r ErrorWriter) Write(p []byte) (n int, err error) {
	color.Set(color.FgRed)
	defer color.Unset()
	return os.Stderr.Write(p)
}

func logInit() {
	Debug = log.New(DebugWriter{},
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(InfoWriter{},
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(ErrorWriter{},
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(ErrorWriter{},
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
