package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var itemsList []item

func main() {
	displaySections()

	r := bufio.NewReader(os.Stdin)
	BillName, _ := getInput("Please enter your name:", r)
	BillName = strings.ToUpper(BillName)
	

	fmt.Println("Add items to cart by entering the item's ID. Hit enter after each, and finally, when you're done ")

	items := retrieveItems()

	newBill := createBill(BillName, items)
	printBill(&newBill)
	saveBill(&newBill)
}

func retrieveItems() []item {
	reader := bufio.NewReader(os.Stdin)
	id, _ := getInput("Enter item ID:  ", reader)
	var idStart string
	if len(id) > 0 {
		idStart = string(id[0])
	} else { idStart = "" }
	
	switch idStart {
		//TODO: Handle outliers
		case "A":
			itemsList = append(itemsList, groceries[id])
			retrieveItems()
		case "M":
			itemsList = append(itemsList, apparel[id])
			retrieveItems()
		case "C":
			itemsList = append(itemsList, accessories[id])
			retrieveItems()
		case "D":
			itemsList = append(itemsList, gadgets[id])
			retrieveItems()
		case "":
			fmt.Println("\nTHANKS FOR STOPPING BY")
		default:
			fmt.Println("Invalid ID")
			retrieveItems()
	}
	return itemsList
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)
	return input, error
}
