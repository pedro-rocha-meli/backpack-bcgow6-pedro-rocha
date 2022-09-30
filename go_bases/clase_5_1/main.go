package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		fmt.Println("Ejecucion finalizada")
	}()

	file, err := os.ReadFile("./archivo.txt")

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println(string(file))

}