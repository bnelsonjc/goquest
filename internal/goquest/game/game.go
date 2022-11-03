// Package game implements things related to the game and engine.
package game

import (
	"bufio"
	"fmt"
	"strings"
)

// Item is an object that can be picked up. It contains a unique label, a description, and aliases
// that it can be referred to by. All aliases SHOULD be unique in case an item is dropped with
// another, but as long as at least ONE alias is present, we can handle the ambiguous case by asking
// player to restate.
type Item struct {

	// Label is a name for the item and canonical way to index it programmatically. It should be
	// upper case and MUST be unique within all labels of the world.
	Label string

	// Description is what is shown when the player LOOKs at the item.
	Description string

	// Aliases are all of the strings that can be used to refer to the item. It must have at least
	// one string that is unique amongst the labels in the world it is in. It does not include Label
	// by default, this must be explicitly given.
	Aliases []string
}

func (item Item) String() string {
	return fmt.Sprintf("Item(%q, (%s))", item.Label, strings.Join(item.Aliases, ", "))
}

// Copy returns a deeply-copied Item.
func (item Item) Copy() Item {
	iCopy := Item{
		Label:       item.Label,
		Description: item.Description,
		Aliases:     make([]string, len(item.Aliases)),
	}

	copy(iCopy.Aliases, item.Aliases)

	return iCopy
}

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

func (egress Egress) String() string {
	return fmt.Sprintf("Egress(%q -> %s)", egress.Aliases, egress.DestLabel)
}

// Copy returns a deeply-copied Egress.
func (egress Egress) Copy() Egress {
	eCopy := Egress{
		DestLabel:     egress.DestLabel,
		Description:   egress.Description,
		TravelMessage: egress.TravelMessage,
		Aliases:       make([]string, len(egress.Aliases)),
	}

	copy(eCopy.Aliases, egress.Aliases)

	return eCopy
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

	// Items is the items on the ground. This can be changed over time.
	Items []Item
}

// Copy returns a deeply-copied Room.
func (room Room) Copy() Room {
	rCopy := Room{
		Label:       room.Label,
		Name:        room.Name,
		Description: room.Description,
		Exits:       make([]Egress, len(room.Exits)),
		Items:       make([]Item, len(room.Items)),
	}

	for i := range room.Exits {
		rCopy.Exits[i] = room.Exits[i].Copy()
	}

	for i := range room.Items {
		rCopy.Items[i] = room.Items[i].Copy()
	}

	return rCopy
}

func (room Room) String() string {
	var exits []string
	for _, eg := range room.Exits {
		exits = append(exits, eg.String())
	}
	exitsStr := strings.Join(exits, ", ")

	return fmt.Sprintf("Room<%s %q EXITS: %s>", room.Label, room.Name, exitsStr)
}

// GetEgressByAlias returns the egress from the room that is represented by the given alias. If no
// Egress has that alias, the returned egress is nil.
func (room Room) GetEgressByAlias(alias string) *Egress {
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

// GetItemByAlias returns the item from the room that is represented by the given alias. If no Item
// has that alias, the returned item is nil.
func (room Room) GetItemByAlias(alias string) *Item {
	var foundItem *Item

	for _, it := range room.Items {
		for _, al := range it.Aliases {
			if al == alias {
				foundItem = &it
				break
			}
		}
	}

	return foundItem
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
		} else if cmd.Verb != "" {
			gotValidCommand = true
		}
	}

	return cmd, nil
}
