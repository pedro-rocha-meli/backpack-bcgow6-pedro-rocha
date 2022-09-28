package main

import(
	"fmt"
)

//Ejercicio 1
type Alumno struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	DNI      int `json:"dni"`
	Fecha    string `json:"fecha"`
}

func (a Alumno) detalle() string {
	return fmt.Sprintf("\nNombre: %s \nApellido: %s\nDNI: %d\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
} 

//Ejercicio 2

type Matrix struct{
	Values []float64
	Width int
	Height int
}

func (matrix *Matrix) setValues(values []float64) {
	matrix.Values = values
}

func (matrix *Matrix) printMatrix()  {
	counter := 0
	for i := 0; i < matrix.Height; i++ {
		fmt.Println("")
		for i := 0; i < matrix.Width; i++ {
			fmt.Print(matrix.Values[counter], " ")
			counter++
		}
	}
}

//Ejercicio 3

type Store struct{
	Products []Products
}

type Products struct{
	Name string
	Type string
	Price float64
}

func main() {

	//Ejercicio 1
	alumno1 := Alumno{
		Nombre: "Marcos",
		Apellido: "Knowles",
		DNI: 32948204,
		Fecha: "12-03-2022",
	}

	fmt.Println(alumno1.detalle())

	matrix := &Matrix{
		Height: 3,
		Width: 4,
	}

	matrix.setValues([]float64{43, 12, 543, 75, 32, 12, 95, 23, 832, 12, 32, 11})

	matrix.printMatrix()
}
