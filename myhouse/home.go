package home

import (
	"house/myhouse/structs/family"
	"house/myhouse/structs/pets"
	"house/myhouse/structs/rooms"
	"fmt"
)

type House struct {
	Square float32
	Family family.Family
	Pets pets.Pets
	Rooms rooms.Rooms
}

func AddHouse() House {
	return House{
		Square: 66.0,
		Family: family.AddFamily(),
		Pets: pets.AddPet(),
		Rooms: rooms.AddRoom(),
	}
}

func (h House) String() string {
	return fmt.Sprintf("___My House___\nSquare: %f\n\n___Family___\n%s\n___Pets___\n%s\n___Rooms(w/ furniture&gadgets)___\n%s",
		h.Square, h.Family.String(), h.Pets.String(), h.Rooms.String())
}