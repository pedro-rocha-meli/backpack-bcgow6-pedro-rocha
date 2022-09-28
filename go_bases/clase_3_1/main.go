package main

import (
	"fmt"
	"os"
)

type Product struct{
	ID string
	Price float64
	Quantity int
}

func (p Product) details() string {
	return fmt.Sprintf("ID: %s\nPrice: %f\nQuantity: %d", p.ID, p.Price, p.Quantity)
}

func main() {

	//Ejercicio 1
	product1 := Product{
		ID: "AX-893",
		Price: 320.21,
		Quantity: 921,
	}
	product2 := Product{
		ID: "AX-4223",
		Price: 120.21,
		Quantity: 21,
	}

	os.Create("products.csv")

	f, err := os.OpenFile("products.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }

	if _, err := f.Write([]byte(product1.details())); err != nil {
        panic(err)
    }
	if _, err := f.Write([]byte(product2.details())); err != nil {
        panic(err)
    }


	//Ejercicio 2

	data, err := os.ReadFile("./products.csv")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}