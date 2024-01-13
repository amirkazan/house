package gadgets

import (
	"fmt"
	"strings"
)

type Gadget struct {
	Type      string
	Brand     string
	Age       int16
	Condition string
}

type GadgetsSet struct {
	GadgetsSet []Gadget
}

func (g *GadgetsSet) AddGadget(gadget Gadget) {
	g.GadgetsSet = append(g.GadgetsSet, gadget)
}

func NewGadgetsSet() *GadgetsSet {
	var gadgets []Gadget

	gadgets = append(gadgets, Gadget{Type: "Smart Speaker", Brand: "Amazon", Age: 1, Condition: "Excellent"})
	gadgets = append(gadgets, Gadget{Type: "Router", Brand: "TP-Link", Age: 2, Condition: "Good"})

	return &GadgetsSet{GadgetsSet: gadgets}
}

func (gs GadgetsSet) String() string {
	var gadgetStrings []string

	for _, g := range gs.GadgetsSet {
		gadgetStrings = append(gadgetStrings, g.String())
	}

	return strings.Join(gadgetStrings, "\n")
}

func (g Gadget) String() string {
	return fmt.Sprintf("{Type: %s, Brand: %s, Age: %d, Condition: %s}",
		g.Type, g.Brand, g.Age, g.Condition)
}