package main

type Product struct {
	Name, Catergory string
	Price           float64
}

var Kayak = Product{
	Name:      "Kayak",
	Catergory: "Watersports",
	Price:     275,
}

var Products = []Product{
	{"Kayak", "Watersports", 275},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}
