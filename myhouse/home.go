package home

import (
	"house/myhouse/structs/family"
	"house/myhouse/structs/pets"
	"house/myhouse/structs/rooms"
	"house/myhouse/structs/gadgets"
	"house/myhouse/structs/furniture"
)

type House struct {
	Square float32
	Family family.Family
	Pets pets.Pets
	Rooms rooms.Rooms
	Gadgets gadgets.Gadgets
	Furniture furniture.FurnitureSet
}

func AddHouse() House {
	return House{
		Square: 66.0,
		Family: family.AddFamily(),
		Pets: pets.AddPet(),
		Rooms: rooms.AddRoom(),
		Gadgets: gadgets.AddGadget(),
		Furniture: furniture.AddFurnitureSet(),
	}
}