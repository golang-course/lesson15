package models

var Phones = []Phone{
	{Id: 1, Vendor: "Xiaomi", Model: "12Lite", Price: 2000},
	{Id: 2, Vendor: "Xiaomi", Model: "Poco 4", Price: 3000},
	{Id: 3, Vendor: "Apple", Model: "iPhone 4", Price: 2500},
}

type Phone struct {
	Id     int
	Vendor string
	Model  string
	Price  int
}
