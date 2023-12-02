package furniture

type Furniture struct{
	Type string
	Age int16
	Condition string
	Material string
	Function string
}

type FurnitureSet struct{
	FurnitureSet []Furniture
}

func addFurnitureSet FurnitureSet{
	var furn []Furniture
	furn = append(furn, Furniture{Type: "Couch", Age: 15 , Condition: "Bad", Material: "Leather", Function: "Chilling"})
	furn = append(furn, Furniture{Type: "Bed", Age: 5, Condition: "Okay", Material: "Wood", Function: "Sleeping"})
	furn = append(furn, Furniture{Type: "Table", Age: 3, Condition: "Pretty good", Material: "Iron&glass", Function: "Eating"})
	furn = append(furn, Furniture{Type: "Chair", Age: 3, Condition: "Pretty good", Material: "Iron&glass", Function: "Sitting"})

	return FurnitureSet{FurnitureSet: furn}
}