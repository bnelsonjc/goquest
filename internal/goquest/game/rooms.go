package game

const (
	// StartLabel is the label of the room that is the starting point of the game.
	StartLabel = "YOUR_ROOM"
)

var (

	// DefaultRooms is the main set of room definitions, which will likely eventually be moved to
	// data.
	DefaultRooms []Room = []Room{
		{
			Label: "YOUR_ROOM",
			Name:  "your bedroom",
			Description: "You are standing in your bedroom. Perhaps you are a young person who" +
				" only now, on your 13th birthday, will receive a name. You have a pretty window" +
				" in the corner overlooking the world outside. There's a door leading to your" +
				" bathroom. to the east, and a door leading to the hall to the south.",
			Exits: []Egress{
				{
					DestLabel:     "BATHROOM",
					Description:   "your bathroom door",
					Aliases:       []string{"BATHROOM", "TOILET", "DOOR", "EAST"},
					TravelMessage: "You go through the door and enter the bathroom.",
				},
				{
					DestLabel:     "HALLWAY",
					Description:   "the door to the hall",
					Aliases:       []string{"HALLWAY", "HALL", "OUT", "SOUTH"},
					TravelMessage: "You shut the door behind you as you go into the hall.",
				},
			},
		},
		{
			Label: "BATHROOM",
			Name:  "your ensuite bathroom",
			Description: "You are in the bathroom attached to your bedroom. There's a toilet," +
				" pristine due to your constant efforts to keep it clean, next to a sink and" +
				" bathtub. You enjoy having your amenities not be a complete mess. The currently" +
				" closed door to the west leads back to your bedroom.",
			Exits: []Egress{
				{
					DestLabel:     "YOUR_ROOM",
					Description:   "the door",
					Aliases:       []string{"BEDROOM", "ROOM", "DOOR", "WEST"},
					TravelMessage: "You head back into the bedroom.",
				},
			},
		},
		{
			Label: "HALLWAY",
			Name:  "the main hallway in your house",
			Description: "This is the main hallway in your house that connects the bedrooms with" +
				" the living room and kitchen. There's a doorway at the end, but it seems boarded" +
				" up, and you're pretty sure that's because it represents the limits of the game" +
				" you're in. I guess you're stuck here. In your home. Housetrapped.\n" +
				" Anyways, there's also your bedroom door at the north end.",
			Exits: []Egress{
				{
					DestLabel:     "YOUR_ROOM",
					Description:   "the door to your bedroom",
					Aliases:       []string{"BEDROOM", "ROOM", "NORTH"},
					TravelMessage: "You step into your bedroom, closing the door behind you for privacy.",
				},
			},
		},
	}
)
