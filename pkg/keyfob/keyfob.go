package keyfob

import (
	"io"
	"os"
)

type Config struct {
	NoExitCode bool
}

// Keyfob is the logic/orchestrator.
type Keyfob struct {
	*Config

	// Allow swapping out stdout/stderr for testing.
	Out io.Writer
	Err io.Writer
}

// New returns a new instance of Keyfob.
func New() *Keyfob {
	config := Config{}

	return &Keyfob{
		Config: &config,
		Out: os.Stdout,
		Err: os.Stdin,	
	}
}
