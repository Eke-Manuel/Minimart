package main

import (
	"fmt"
)

type item struct {
	name  string
	price float64
}

type section map[string]item

var sections = []section{groceries, apparel, accessories, gadgets}

var groceries = section{
	"A1": item{"Milk", 7.66},
	"A2": item{"Cheese", 3.70},
	"A3": item{"Chicken", 6.00},
	"A4": item{"Olive oil", 14.99},
	"A5": item{"Bread", 2.30},
}

var apparel = section{
	"M1": item{"T-shirt", 27.99},
	"M2": item{"Jeans", 73.70},
	"M3": item{"Belt", 16.00},
	"M4": item{"Tuxedo", 299.99},
	"M5": item{"Head warmer", 3.00},
}

var accessories = section{
	"C1": item{"Rolex", 5999.99},
	"C2": item{"Silver pendant", 93.70},
	"C3": item{"Earrings", 138.00},
	"C4": item{"Patek Phillipe", 20899.99},
	"C5": item{"Necklace", 170.00},
}

var gadgets = section{
	"D1": item{"Earbuds", 120.00},
	"D2": item{"MI 11", 699.00},
	"D3": item{"Dell XPS 15", 1200.00},
	"D4": item{"WiFi router", 299.99},
	"D5": item{"Google Pixel 6", 800.00},
}

func displaySections() {
	itemsList := fmt.Sprintf("%-6v %-25v %v \n \n", "ID", "ITEMS", "PRICES")
	for i := 0; i < len(sections); i++ {
		for k, v := range sections[i] {
			itemsList += formatItemString(k, v.name, v.price)
		}
	}
	fmt.Println(itemsList)
}

func formatItemString(id string, name string, price float64) string {
	str := fmt.Sprintf("%-6v %-25v $%0.2f \n", id, name, price)
	return str
}
