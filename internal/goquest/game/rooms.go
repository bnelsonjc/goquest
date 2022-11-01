package game

import "fmt"

const (
	// StartLabel is the label of the room that is the starting point of the game.
	StartLabel = "YOUR_ROOM"
)

var (

	// AllRooms is loaded with definitions for all rooms in the game, marked by the labels to them.
	AllRooms map[string]Room
)

func init() {
	for _, r := range roomDefs {
		if _, ok := AllRooms[r.Label]; ok {
			panic(fmt.Sprintf("duplicate room label %q in roomDefs", r.Label))
		}

		// sanity check that egress aliases are not duplicated
		seenAliases := map[string]bool{}
		for _, eg := range r.Exits {
			for _, alias := range eg.Aliases {
				if _, ok := seenAliases[alias]; ok {
					errMsg := "duplicate egress alias %q in room %q in roomDefs"
					panic(fmt.Sprintf(errMsg, alias, r.Label))
				}
			}
		}

		AllRooms[r.Label] = r
	}

	// TODO: after all this is done, ensure that all room egresses are valid existing labels
}

var (

	// roomDefs is where AllRooms is loaded from.
	roomDefs []Room = []Room{
		Room{
			Label: "YOUR_ROOM",
			Name:  "your bedroom",
			Description: "You are standing in your bedroom. Perhaps you are a young person who" +
				" only now, on your 13th birthday, will receive a name. You have a pretty window" +
				" the corner overlooking the world outside, and a door leading to your bathroom" +
				" to the east.",
			Exits: []Egress{
				Egress{
					DestLabel:     "BATHROOM",
					Description:   "your bathroom door",
					Aliases:       []string{"BATHROOM", "TOILET", "DOOR", "EAST"},
					TravelMessage: "You go through the door and enter the bathroom.",
				},
			},
		},
		Room{
			Label: "BATHROOM",
			Name:  "your ensuite bathroom",
			Description: "You are in the bathroom attached to your bedroom. There's a toilet," +
				" pristine due to your constant efforts to keep it clean, next to a sink and" +
				" bathtub. You enjoy having your amenities not be a complete mess. The currently" +
				" closed door to the west leads back to your bedroom.",
			Exits: []Egress{
				Egress{
					DestLabel:     "YOUR_ROOM",
					Description:   "the door",
					Aliases:       []string{"BEDROOM", "ROOM", "DOOR", "WEST"},
					TravelMessage: "You head back into the bedroom.",
				},
			},
		},
	}
)
