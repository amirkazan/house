package gadgets

type TV struct {
	Type string
	Brand string
	Age int16
	Condition string
}

type Laptop struct {
	Type string
	Brand string
	OS string
	Condition string
}

type Phone struct {
	Type string
	Brand string
	OS string
	Age int16
	Condition string
}

type Gadgets struct {
	tv TV
	laptop Laptop
	phone Phone
}

func AddGadget() Gadgets {
	return Gadgets{
		tv: TV{Type: "TV", Brand: "Samsung", Age: 3, Condition: "Good"},
		laptop: Laptop{Type: "Laptop", Brand: "Asus", OS: "Windows", Condition: "Perfect"},
		phone: Phone{Type: "Phone", Brand: "Apple", OS: "IOS", Age: 2, Condition: "Perfect"},
	}
}