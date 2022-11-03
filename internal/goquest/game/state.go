package game

import (
	"bufio"
	"fmt"
	"strings"
)

// State is the game's entire state.
type State struct {
	// World is all rooms that exist and their current state.
	World map[string]*Room

	// CurrentRoom is the room that the player is in.
	CurrentRoom *Room

	// Inventory is the objects that the player currently has.
	Inventory []string
}

// New creates a new State and loads the list of rooms into it. It performs basic sanity checks
// to ensure that a valid world is being passed in and normalizes them as needed.
//
// startingRoom is the label of the room to start with.
func New(world []Room, startingRoom string) (State, error) {
	gs := State{}

	for _, r := range world {
		if _, ok := gs.World[r.Label]; ok {
			return gs, fmt.Errorf("duplicate room label %q in roomDefs", r.Label)
		}

		// sanity check that egress aliases are not duplicated
		seenAliases := map[string]bool{}
		for _, eg := range r.Exits {
			for _, alias := range eg.Aliases {
				if _, ok := seenAliases[alias]; ok {
					errMsg := "duplicate egress alias %q in room %q in roomDefs"
					return gs, fmt.Errorf(errMsg, alias, r.Label)
				}
			}
		}

		roomCopy := r.Copy()
		gs.World[r.Label] = &roomCopy
	}

	// TODO: after all this is done, ensure that all room egresses are valid existing labels

	// now set the current room
	var startExists bool
	gs.CurrentRoom, startExists = gs.World[startingRoom]
	if !startExists {
		return gs, fmt.Errorf("starting room with label %q does not exist in passed-in rooms", startingRoom)
	}

	return gs, nil
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
		egress := gs.CurrentRoom.GetEgress(cmd.Recipient)
		if egress == nil {
			return fmt.Errorf("%q isn't a place you can go from here", cmd.Recipient)
		}

		gs.CurrentRoom = gs.World[egress.DestLabel]

		output = egress.TravelMessage
	case "EXITS":
		exitTable := ""

		for _, eg := range gs.CurrentRoom.Exits {
			exitTable += strings.Join(eg.Aliases, "/")
			exitTable += " -> "
			exitTable += eg.Description
			exitTable += "\n"
		}

		output = exitTable
	case "LOOK":
		if cmd.Recipient != "" {
			return fmt.Errorf("I can't LOOK at particular things yet")
		}

		output = gs.CurrentRoom.Description
	case "DEBUG":
		if cmd.Recipient == "ROOM" {
			output = gs.CurrentRoom.String()
		} else {
			return fmt.Errorf("I don't know how to debug %q", cmd.Recipient)
		}
	case "HELP":
		output = "Here are the commands you can use (WIP commands do not yet work fully):\n"
		output += "HELP       - show this help\n"
		output += "DROP/PUT   - put down an object in the room [WIP]\n"
		output += "DEBUG ROOM - print info on the current room\n"
		output += "EXITS      - show the names of all exits from the room\n"
		output += "GO/MOVE    - go to another room via one of the exits\n"
		output += "LOOK       - show the description of the room\n"
		output += "QUIT/EXIT  - end the game\n"
		output += "TAKE/GET   - pick up an object in the room [WIP]\n"
		output += "TALK/SPEAK - talk to someone/something in the room [WIP]\n"
		output += "USE        - use an object in your inventory [WIP]\n"
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
