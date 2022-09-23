package main

import "fmt"

func main()  {
	x,y := 32,55
	fmt.Println(x * y)

	const HELLO string = "HELLO WORLD"

	fmt.Println(HELLO)

	var a [2]string

	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0])

	var numero int

	fmt.Println(numero)

	var myArray = [...]int{3,2,7}

	fmt.Println(myArray)


}