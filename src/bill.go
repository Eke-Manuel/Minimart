package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type bill struct {
	owner           string
	id              uint32
	items           []item
	totalPrice      float64
	transactionTime time.Time
}

func createBill(billName string, items []item) bill {
	var total float64
    for _, v := range items {
		total += v.price
	}
	newBill := bill{
		owner:           billName,
		id:              rand.Uint32(),
		items:           items,
		totalPrice:      total,
		transactionTime: time.Now(),
	}
	return newBill
}

func formatBillString(b *bill) string {
	billString := fmt.Sprintf("\n %25v -----%v's Bill-----\nBill ID: %v \n", "", b.owner, b.id)
	billString += fmt.Sprintf("%-25v %v \n \n", "ITEMS", "PRICES")
	for _, v := range b.items {
		billString += fmt.Sprintf("%-25v $%0.2f \n", v.name, v.price)
	}
	billString += fmt.Sprintf("\n%-25v $%0.2f \n", "TOTAL", b.totalPrice)
	billString += fmt.Sprintf("%-25v %v \n \n", "Transaction time:", b.transactionTime)

	return billString
}

func printBill(b *bill) {
	formattedBill := formatBillString(b)
	fmt.Println(formattedBill)
}

func saveBill(b *bill) {
	billString := formatBillString(b)
	path := fmt.Sprintf("bills/%v's BILL", b.owner)
	err := os.WriteFile(path, []byte(billString), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill has been saved to bills folder")
}
