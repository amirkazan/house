package rooms

import (
	"house/myhouse/structs/furniture"
	"house/myhouse/structs/gadgets"
	"fmt"
	"strings"
)

type Room struct {
	Name      string
	Length    float32
	Width     float32
	Square    float32
	Furniture furniture.FurnitureSet
	Gadgets   gadgets.GadgetsSet
}

type Rooms struct {
	Kitchen  Room
	Bathroom Room
	Bedroom  Room
	Hall     Room
}

func AddRoom() Rooms {
	kitchenFurnitureSet := furniture.NewFurnitureSet()
	kitchenFurnitureSet.AddFurniture(furniture.Furniture{Type: "Couch", Age: 15, Condition: "Bad", Material: "Leather", Function: "Chilling"})

	bathroomFurnitureSet := furniture.NewFurnitureSet()
	bathroomFurnitureSet.AddFurniture(furniture.Furniture{Type: "Toilet", Age: 2, Condition: "Good", Material: "Ceramic", Function: "Sitting"})

	bedroomFurnitureSet := furniture.NewFurnitureSet()
	bedroomFurnitureSet.AddFurniture(furniture.Furniture{Type: "Bed", Age: 5, Condition: "Okay", Material: "Wood", Function: "Sleeping"})
	bedroomFurnitureSet.AddFurniture(furniture.Furniture{Type: "Wardrobe", Age: 3, Condition: "Pretty good", Material: "Wood", Function: "Storage"})

	hallFurnitureSet := furniture.NewFurnitureSet()
	hallFurnitureSet.AddFurniture(furniture.Furniture{Type: "Table", Age: 3, Condition: "Pretty good", Material: "Iron&glass", Function: "Eating"})
	hallFurnitureSet.AddFurniture(furniture.Furniture{Type: "Chair", Age: 3, Condition: "Pretty good", Material: "Iron&glass", Function: "Sitting"})

	kitchenGadgetSet := gadgets.NewGadgetsSet()
	kitchenGadgetSet.AddGadget(gadgets.Gadget{Type: "Fridge", Brand: "Samsung", Age: 3, Condition: "Good"})

	bathroomGadgetSet := gadgets.NewGadgetsSet()
	bathroomGadgetSet.AddGadget(gadgets.Gadget{Type: "Washer", Brand: "LG", Age: 2, Condition: "Perfect"})

	bedroomGadgetSet := gadgets.NewGadgetsSet()
	bedroomGadgetSet.AddGadget(gadgets.Gadget{Type: "Laptop", Brand: "Asus", Condition: "Perfect"})
	bedroomGadgetSet.AddGadget(gadgets.Gadget{Type: "Phone", Brand: "Apple", Age: 2, Condition: "Perfect"})

	hallGadgetSet := gadgets.NewGadgetsSet()
	hallGadgetSet.AddGadget(gadgets.Gadget{Type: "TV", Brand: "Sony", Age: 3, Condition: "Excellent"})

	return Rooms{
		Kitchen:  Room{Name: "Kitchen", Length: 4.0, Width: 4.0, Square: 16.0, Furniture: *kitchenFurnitureSet, Gadgets: *kitchenGadgetSet},
		Bathroom: Room{Name: "Bathroom", Length: 4.0, Width: 2.0, Square: 8.0, Furniture: *bathroomFurnitureSet, Gadgets: *bathroomGadgetSet},
		Bedroom:  Room{Name: "Bedroom", Length: 3.0, Width: 4.0, Square: 12.0, Furniture: *bedroomFurnitureSet, Gadgets: *bedroomGadgetSet},
		Hall:     Room{Name: "Hall", Length: 5.0, Width: 6.0, Square: 30.0, Furniture: *hallFurnitureSet, Gadgets: *hallGadgetSet},
	}
}

func (r Rooms) String() string {
	roomStrings := []string{
		r.Kitchen.String(),
		r.Bathroom.String(),
		r.Bedroom.String(),
		r.Hall.String(),
	}

	return strings.Join(roomStrings, "\n")
}

func (r Room) String() string {
	return fmt.Sprintf("Name: %s\nLength: %f\nWidth: %f\nSquare: %f\nFurniture: %s\nGadgets: %s\n",
		r.Name, r.Length, r.Width, r.Square, r.Furniture.String(), r.Gadgets.String())
}