package main

import (
	"fmt"
	"house/structs/pets"
	"house/structs/gadgets"
	"house/structs/family"
	"house/structs/furniture"
	"house/structs/rooms"
)

type House struct {
	Square float32
	Family family.Family
	Pets pets.Pets
	Rooms rooms.Rooms
	Gadgets gadgets.Gadgets
	Furniture furniture.FurnitureSet
}

func main() {
	house := House{
		Square: 66.0,
		Family: family.addFamily(),
		Pets: pets.addPet(),
		Rooms: rooms.addRoom(),
		Gadgets: gadgets.addGadget(),
		Furniture: furniture.addFurnitureSet()
	}

	fmt.Println(house)
}