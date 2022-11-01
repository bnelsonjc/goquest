package game

import (
	"bufio"
	"fmt"
)

// State is the game's entire state.
type State struct {
	// Room is the current room that the player is in.
	Room Room

	// Inventory is the objects that the player currently has.
	Inventory []string
}

// Advance advances the game state based on the given command. If there is a problem executing the
// command, it is given in the error output and the game state is not advanced. If it is, the
// result of the command is written to the provided output stream.
//
// Invalid commands will be returned as non-nil errors as opposed to writing directly to the IO
// stream; the caller can decide whether to do this themself.
//
// Note that for this, QUIT is not considered a valid command is it would be on a controlling engine
// to end the game state based on that.
//
// TODO: differentiate syntax errors from io errors
func (gs *State) Advance(cmd Command, ostream *bufio.Writer) error {
	var output string

	switch cmd.Verb {
	case "QUIT":
		return fmt.Errorf("I can't QUIT; I'm not being executed by a quitable engine")
	case "GO":
		egress := gs.Room.GetEgress(cmd.Recipient)
		if egress == nil {
			return fmt.Errorf("%q isn't a place you can go from here", cmd.Recipient)
		}

		newRoom := AllRooms[egress.DestLabel]
		gs.Room = newRoom

		output = egress.TravelMessage
	case "LOOK":
		if cmd.Recipient != "" {
			return fmt.Errorf("I can't LOOK at particular things yet")
		}

		output = gs.Room.Description
	case "DEBUG":
		if cmd.Recipient == "ROOM" {
			output = gs.Room.String()
		} else {
			return fmt.Errorf("I don't know how to debug %q", cmd.Recipient)
		}
	default:
		return fmt.Errorf("I don't know how to %q", cmd.Verb)
	}

	// IO to give output:
	if _, err := ostream.WriteString(output + "\n\n"); err != nil {
		return fmt.Errorf("could not write output: %w", err)
	}
	if err := ostream.Flush(); err != nil {
		return fmt.Errorf("could not flush output: %w", err)
	}

	return nil
}
