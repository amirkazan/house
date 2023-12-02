package rooms

type Room struct {
	Name string
	Length float32
	Width float32
	Square float32
}

type Rooms struct {
	kitchen Room
	bathroom Room
	bedroom Room
	hall Room
}

func addRoom() Rooms {
	return Rooms{
		kitchen: Room{Name: "Kitchen", Length: "4.0", Width: "4.0", Square: "16.0"},
		bathroom: Room{Name: "Bathroom", Length: "4.0", Width: "2.0", Square: "8.0"},
		bedroom: Room{Name: "Bedroom", Length: "3.0", Width: "4.0", Square: "12.0"},
		hall: Room{Name: "Hall", Length: "5.0", Width: "6.0", Square: "30.0"}
	}
}
