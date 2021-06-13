package keyfob

import (
	"fmt"
	"io"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

const (
	keyfobHeaderPrefix = "keyfob"
)

// printToStream prints a generic message to a stream (stdout/stderror) in color.
func printToStream(stream io.Writer, msg interface{}) {
	_, err := fmt.Fprint(stream, fmt.Sprintf("%v\n", msg))
	if err != nil {
		panic(err)
	}
}

// printToStreamWithColor prints a message after wrapping it in ANSI color codes.
func printToStreamWithColor(stream io.Writer, color text.Color, msg interface{}) {
	_, err := fmt.Fprintf(stream, color.Sprintf("%v\n", msg))
	if err != nil {
		panic(err)
	}
}

// OutputHeading prints a header to stdout.
func (keyfob Keyfob) OutputHeading(msg interface{}) {
	printToStreamWithColor(keyfob.Out, text.FgHiCyan, msg)
}

// Output prints a normal message to stdout.
func (keyfob Keyfob) Output(msg interface{}) {
	printToStream(keyfob.Out, msg)
}

// Error pritns an error to stderr and exits with error code 1.
func (keyfob Keyfob) Error(msg interface{}) {
	printToStreamWithColor(keyfob.Err, text.FgHiRed, fmt.Sprintf("Error: %v\n", msg))
	if !keyfob.NoExitCode {
		os.Exit(1)
	}
}
