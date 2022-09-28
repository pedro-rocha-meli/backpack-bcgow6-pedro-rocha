package main

import (
	"errors"
	"fmt"
)

//Ejercicio 1

type SalaryError struct {
	StatusCode int
	Message    string
}

func (err *SalaryError) Error() string {
	return fmt.Sprintf("status: %d\nerror: %s", err.StatusCode, err.Message)
}

func validateSalary(salary float64) error {
	if salary < 150000 {
		return &SalaryError{
			StatusCode: 500,
			Message:    "el salario ingresado no alcanza el mínimo imponible",
		}
	}
	fmt.Println("Debe pagar impuesto")
	return nil
}

func main() {

	//Ejercicio 1
	var salary float64 = 100000
	err := validateSalary(salary)
	fmt.Println(err)

	//Ejercicio 2
	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	//Ejercicio 3
	err3 := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %f", salary)

	if salary < 150000 {
		fmt.Println(err3)
	}
}
