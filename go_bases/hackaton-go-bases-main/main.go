package main

import (
	"fmt"
	"hackaton/internal/file"
	"hackaton/internal/service"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("An error ocurred during the excecution of the program")
		}
	}()

	
	var tickets []service.Ticket
	
	// Funcion para obtener tickets del archivo csv
	f := file.File{
		Path: "./tickets.csv",
	}
	
	data, err := f.Read()
	if err != nil {
		panic(err)
	}
	
	tickets = append(tickets, data...)
	
	fmt.Println("Showing first 5 tickets...")
	for i := 0; i < 5; i++ {
		fmt.Println(tickets[i])
	}
	
	newTicket := service.Ticket{
		Id: 1001,
	    Names: "Mateo Salvatto",
	    Email: "mateo.salvatto@gmail.com",
	    Destination: "EEUU",
	    Date: "0:30,834",
	    Price: 2000,
	}
	
	var bookings service.Bookings = service.NewBookings(tickets)

	_, err = bookings.Create(newTicket)
	if err != nil {
		panic(err)
	}
	fmt.Println("\ncreated new ticket successfully...")

	readTicket, err := bookings.Read(1001)
	if err != nil {
	    panic(err)
	}

	fmt.Println("\nreading new ticket...")
	fmt.Println(readTicket)

	readTicket.Destination = "Luxemburgo"
	_, err = bookings.Update(1001, readTicket)
	if err != nil {
	    panic(err)
	}

	fmt.Println("\nupdating new ticket...")
	fmt.Println(readTicket)

	fmt.Println()
	
	_, err = bookings.Delete(1001)
	if err != nil {
	    panic(err)
	}
}
