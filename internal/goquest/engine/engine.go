// Package engine is a CLI-driven engine for getting commands and advancing the game state
// continuously until the user quits.
package engine

import (
	"bufio"
	"io"
	"os"

	"github.com/bnelsonjc/goquest/internal/goquest/game"
)

// Engine contains the things needed to run a game from an interactive shell attached to an input
// stream and an output stream.
type Engine struct {
	state game.State
	in    *bufio.Reader
	out   *bufio.Writer
}

// New creates a new engine ready to operate on the given input and output streams. It will
// immediately open a buffered reader on the input stream and a buffered writer on the output
// stream.
//
// If nil is given for the input stream, a bufio.Reader is opened on stdin.
// If nil is given for the output stream, a bufio.Writer is opened on stdout.
func New(inputStream io.Reader, outputStream io.Writer) *Engine {
	if inputStream == nil {
		inputStream = os.Stdin
	}
	if outputStream == nil {
		outputStream = os.Stdout
	}

	eng := &Engine{in: bufio.NewReader(inputStream), out: bufio.NewWriter(outputStream)}

	return eng
}
