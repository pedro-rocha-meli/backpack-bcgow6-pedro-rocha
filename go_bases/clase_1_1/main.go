package main

import "fmt"

// Ejercicio 1
var nombre string = "Pedro"
var direccion string = "Buenos Aires, Tigre"

var temperatura float64 = 23.4
var humedad float64 = 28.9
var presion float64 = 12.3


func main()  {
	// Ejercicio 1
	fmt.Println("Hola mi nombre es " + nombre)
	fmt.Println("Soy de " + direccion)

	// Ejercicio 2
	fmt.Println("La temperatura de hoy es: ", temperatura)
	fmt.Println("La humedad es de: ", humedad)		
	fmt.Println("La presion es de: ", presion)

	//2- Declararia estas variables como float64


	// Ejercicio 3
	// 1- El nombre de la variable no puede arrancar con un numero
	// 2- Esta declaracion es valida
	// 3- El tipo de dato esta no puede declararse previo al nombre de la variable
	// 4- El nombre de la variable no puede arrancar con un numero
	// 5- Esta declaracion es valida
	// 6- El nombre de la variable no puede contener espacios
	// 7- Esta declaracion es valida

	// Ejercicio 4
	// var apellido string = "Gomez"
	// var edad int = 35
	// var sueldo float64 = 45857.90
	// var nombre2 string = "Julián"
	// booleann := false
}

