package pets

type Pet struct {
	Spec string
	Name string
	Age int16
	Color string
}

type Pets struct {
	Pets []Pet
}

func addPet() Pets{
	var pets []Pet

	pets = append(pets, Pet{Spec: "Penguin", Name: "Estriper", Age: 18, Color: "Black&white"})

	return Pets{Pets: pets}
}