// Package game implements things related to the game and engine.
package game

import (
	"bufio"
	"fmt"
	"strings"
)

// GO place
// GET thing
// USE thing
// TALK to thing
// QUIT the game

// Command is a valid command received from a game prompt.
type Command struct {

	// Verb is the canonical name of the command being invoked, such as "MOVE", "GET", "USE", or
	// "QUIT". Some verbs may have shorthand forms which are typed differently, for instance "GO"
	// could be typed instead of "MOVE", or "NORTH" instead of "MOVE NORTH", etc, and for all those
	// cases they would result in a Command with a verb of GO.
	Verb string

	// Instrument is what is doing the action, for instance in "GET WATER WITH CUP", or
	// alternatively "USE WATER WITH CUP", "CUP" would be identified as the instrument. The exact
	// meaning depends on the verb.
	Instrument string

	// Recipient is the thing receiving the action, for instance "GET CUP" or "TALK TO MAN", the
	// recipient would be "CUP" and "MAN" respectively. For MOVE commands, this can also be a
	// direction.
	Recipient string
}

// GetCommand is the fundamental unit of obtaining input from the user in an interactive fashion.
// It prompts the user for an input and attempts to parse it as a valid command, returning that
// command if it is successful. If it is not, error output is printed to the ostream and the user
// is prompted until they enter a valid command.
//
// Note that this function does not check if the command is executable, only that a Command can be
// parsed from the user input.
func GetCommand(istream bufio.Reader, ostream bufio.Writer) (Command, error) {
	var cmd Command

	// do IO
	if _, err := ostream.WriteString("Enter command\n"); err != nil {
		return cmd, fmt.Errorf("could not write output: %w", err)
	}
	if err := ostream.Flush(); err != nil {
		return cmd, fmt.Errorf("could not flush output: %w", err)
	}
	input, err := istream.ReadString('\n')
	if err != nil {
		return cmd, fmt.Errorf("could not get input: %w", err)
	}

	// now attempt to parse the input

	return cmd, nil
}

// ParseCommand parses a command from the given text. If it cannot, a non-nil error is returned.
//
// If an empty string or a string composed only of whitespace is passed in, nil error is
// returned and a zero value for Command will be returned.
func ParseCommand(toParse string) (Command, error) {
	var parsedCmd Command

	// make entire input upper case to make matching easy
	normalizedCase := strings.ToUpper(toParse)

	// now tokenize our string, collapsing all whitespace
	tokens := strings.Fields(normalizedCase)

	// some simple sanity checking, make sure we at least have a command
	if len(tokens) < 1 {
		return parsedCmd, nil
	}

	// set verb as the first word here, we'll update it to synonyms as needed
	parsedCmd.Verb = tokens[0]

	// next, do simple matching on our main keywords based on the first word
	cae 
		parsedCmd.Verb = "HELP"
	case "GO":
		fallthrough
	case ""
	}
}
