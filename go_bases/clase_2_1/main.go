package main

import (
	"errors"
	"fmt"
)

// Ejercicio 1

const (
	MAX_TAX = 0.27
	MIN_TAX = 0.17
)

func calculateTax(salary float64) float64 {

	switch {
	case salary > 50000:
		if salary > 150000 {
			return salary * MAX_TAX
		}
		return salary * MIN_TAX
	}

	return 0
}

// Ejercicio 2

func avg(num ...float64) (result float64, err error) {

	for _, number := range num {
		if number < 0 {
			err = errors.New("negative values aren´t accepted")
		}
		result += number
	}
	result /= float64(len(num))

	return
}

// Ejercicio 3
const (
	CATEGORY_A = "A"
	CATEGORY_B = "B"
	CATEGORY_C = "C"
)

func calculateSalary(minutesWorked float64, categoryType string) (salary float64) {

	hoursWorked := minutesWorked / 60

	switch categoryType {
	case "A":
		salary = (hoursWorked * 3000)
		salary += salary * 0.5
	case "B":
		salary = (hoursWorked * 1500)
		salary += salary * 0.2
	case "C":
		salary = hoursWorked * 1000
	}

	return
}

//Ejercicio 4

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calcualteMinimum(values ...float64) (min float64) {
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return
}

func calculateMaximum(values ...float64) (max float64) {
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return
}

func calculateAverage(values ...float64) (avg float64) {
	for _, num := range values {
		avg += num
	}
	avg /= float64(len(values))

	return
}

func operation(operator string) func(...float64) float64 {
	switch operator {
	case "minimum":
		return calcualteMinimum
	case "maximum":
		return calculateMaximum
	case "average":
		return calculateAverage
	}
	return nil
}

//Ejercicio 5

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func totalDogFood(numberOfDogs ...int) (total float64) {
	return float64(len(numberOfDogs) * 10)
}
func totalCatFood(numberOfCats ...int) (total float64) {
	return float64(len(numberOfCats) * 5)
}
func totalHamsterFood(numberOfHamsters ...int) (total float64) {
	return float64(len(numberOfHamsters)) * 0.25
}
func totalSpiderFood(numberOfSpiders ...int) (total float64) {
	return float64(len(numberOfSpiders)) * 0.15
}

func animal(animal string) (func(...int) float64, error) {

	switch animal {
	case "dog":
		return totalDogFood, nil
	case "cat":
		return totalCatFood, nil
	case "hamster":
		return totalHamsterFood, nil
	case "spider":
		return totalSpiderFood, nil
	}

	return nil, errors.New("this animal isn´t in the vet shop")
}

func main() {
	fmt.Println(calculateTax(60000))
	fmt.Println(avg(32, 52, 12, 88, 3.23))
	fmt.Println(calculateSalary(60, CATEGORY_A))

	var calcUsed func(...float64) float64 = operation(maximum)
	fmt.Println(calcUsed(54, 23, 12, 55))

}
