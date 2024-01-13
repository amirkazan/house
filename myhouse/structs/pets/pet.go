package pets

import (
	"fmt"
	"strings"
)

type Pet struct {
	Spec string
	Name string
	Age int16
	Color string
}

type Pets struct {
	Pets []Pet
}

func AddPet() Pets{
	var pets []Pet

	pets = append(pets, Pet{Spec: "Penguin", Name: "Estriper", Age: 18, Color: "Black&white"})

	return Pets{Pets: pets}
}

func (p Pet) String() string {
	return fmt.Sprintf("Spec: %s, Name: %s, Age: %d, Color: %s",
		p.Spec, p.Name, p.Age, p.Color)
}

func (p Pets) String() string {
	var petsStrings []string

	for _, pet := range p.Pets {
		petsStrings = append(petsStrings, pet.String())
	}

	return "\n" + strings.Join(petsStrings, "\n") + "\n"
}