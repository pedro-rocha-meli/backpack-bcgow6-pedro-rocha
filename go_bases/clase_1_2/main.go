package main

import "fmt"
import "strings"

func main() {

	//Ejercicio 1
	palabra := "Manzana"

	fmt.Println(len(palabra),"\n")

	characters := strings.Split(palabra, "")

	for i := 0; i < len(characters); i++ {
		fmt.Println(characters[i])
	}

	//Ejercicio 2
	edad := 30
	empleado := true
	antiguedad := 2
	sueldo := 200000

	switch {
	case edad > 22 && empleado && antiguedad > 1:
		if sueldo > 100000 {
			fmt.Println("Prestamo entregado sin tasa de interes.")
		} else {
			fmt.Println("Prestamo entregado con tasa de interes.")
		}
	default:
		fmt.Println("El cliente no cumple los requisitos para acceder al prestamo.")	
	}

	//Ejercicio 3

	numeroMes := 4

	switch numeroMes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	}

	
	//Ejercicio 4
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	
	fmt.Println(employees["Benjamin"])

	mayoresA21 := 0

	for _, element := range employees {
        if element > 21 {
			mayoresA21++
		}
    }

	fmt.Println(mayoresA21)

	employees["Jacinta"] = 18
	
	delete(employees, "Pedro")
}