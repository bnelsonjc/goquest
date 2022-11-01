// Package game implements things related to the game and engine.
package game

import (
	"bufio"
	"fmt"
)

// Egress is an egress point from a room. It contains both a description and the label it points to.
type Egress struct {
	// DestLabel is the label of the room this egress goes to.
	DestLabel string

	// Description is the long description of the egress point.
	Description string

	// TravelMessage is the message shown when the player uses this egress point.
	TravelMessage string

	// Aliases is the list of aliases that the user can give to travel via this egress. Note that
	// the label is not included in this list by default to prevent spoilerific room names.
	Aliases []string
}

// Room is a scene in the game. It contains a series of exits that lead to other rooms and a
// description. They also contain a list of the interactables at game start (or will in the future).
type Room struct {
	// Label is how the room is referred to in the game. It must be unique from all other Rooms.
	Label string

	// Name is used in short descriptions (prior to LOOK).
	Name string

	// Description is what is returned when LOOK is given with no arguments.
	Description string

	// Exits is a list of room labels and ways to describe them, pointing to other rooms in the
	// game.
	Exits []Egress
}

// GetEgress returns the egress from the room that is represented by the given alias. If no Egress
// has that alias, the returned egress is nil.
func (room Room) GetEgress(alias string) *Egress {
	var foundEgress *Egress

	for _, eg := range room.Exits {
		for _, al := range eg.Aliases {
			if al == alias {
				foundEgress = &eg
				break
			}
		}
	}

	return foundEgress
}

// GetCommand is the fundamental unit of obtaining input from the user in an interactive fashion.
// It prompts the user for an input and attempts to parse it as a valid command, returning that
// command if it is successful. If it is not, error output is printed to the ostream and the user
// is prompted until they enter a valid command.
//
// Note that this function does not check if the command is executable, only that a Command can be
// parsed from the user input.
func GetCommand(istream *bufio.Reader, ostream *bufio.Writer) (Command, error) {
	var cmd Command
	gotValidCommand := false

	if _, err := ostream.WriteString("Enter command\n"); err != nil {
		return cmd, fmt.Errorf("could not write output: %w", err)
	}
	if err := ostream.Flush(); err != nil {
		return cmd, fmt.Errorf("could not flush output: %w", err)
	}

	for !gotValidCommand {
		// IO to get input:
		if _, err := ostream.WriteString("> "); err != nil {
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
		cmd, err = ParseCommand(input)
		if err != nil {
			errMsg := fmt.Sprintf("%v\nTry HELP for valid commands\n", err.Error())
			// IO to report error and prompt user to try again
			if _, err := ostream.WriteString(errMsg); err != nil {
				return cmd, fmt.Errorf("could not write output: %w", err)
			}
			if err := ostream.Flush(); err != nil {
				return cmd, fmt.Errorf("could not flush output: %w", err)
			}
		} else {
			gotValidCommand = true
		}
	}

	return cmd, nil
}
