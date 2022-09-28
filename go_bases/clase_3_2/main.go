package main

import (
	"fmt"
)

// Ejercicio 1
type User struct {
	Name     string
	SurName  string
	Age      byte
	Email    string
	Password string
	Products []Product
}

//Ejercicio 2
type Product struct {
	Name string
	Price float64
	Quantity int
}

func (user *User) ChangeName(name string, surname string) {
	user.Name = name
	user.SurName = surname
}

func (user *User) ChangeAge(age byte) {
	user.Age = age
}

func (user *User) ChangeEmail(email string) {
	user.Email = email
}

func (user *User) ChangePassword(password string) {
	user.Password = password
}

func NewProduct(name string, price float64) Product {
	return Product{Name: name, Price: price}
}

func (user *User)AddProduct(product *Product){
	user.Products = append(user.Products, *product)
}

func main() {

	//Ejercicio 1
	user := User{
		Name:     "Rodrigo",
		SurName:  "Barrios",
		Age:      22,
		Email:    "rodrigo@gmail.com",
		Password: "password1",
	}

	user.ChangeName("Roberto", "Bauberri")
	user.ChangeAge(32)
	user.ChangeEmail("roberto@gmail.com")
	user.ChangePassword("password2")

	fmt.Println(user)

}
