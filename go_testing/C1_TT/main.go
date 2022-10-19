package main

import (
	"fmt"

	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_testing/C1_TT/ordenamiento"
)

func main()  {
	
	a := []int{54, 32, 32, 12, 66, 86, 31, 33}

	result := ordenamiento.Ordenar(a)

	fmt.Println(result)
}