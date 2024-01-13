package furniture

import (
	"fmt"
	"strings"
)

type Furniture struct {
	Type      string
	Age       int16
	Condition string
	Material  string
	Function  string
}

type FurnitureSet struct {
	FurnitureSet []Furniture
}

func NewFurnitureSet() *FurnitureSet {
	return &FurnitureSet{}
}

func (f *FurnitureSet) AddFurniture(item Furniture) {
	f.FurnitureSet = append(f.FurnitureSet, item)
}

func (fs FurnitureSet) String() string {
	var furnitureStrings []string

	for _, f := range fs.FurnitureSet {
		furnitureStrings = append(furnitureStrings, f.String())
	}

	return strings.Join(furnitureStrings, "\n")
}

func (f Furniture) String() string {
	return fmt.Sprintf("{Type: %s, Age: %d, Condition: %s, Material: %s, Function: %s}",
		f.Type, f.Age, f.Condition, f.Material, f.Function)
}