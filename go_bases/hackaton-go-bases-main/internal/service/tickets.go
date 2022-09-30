package service

import (
	"errors"
	"fmt"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id          int    `csv:"Id"`
	Names       string `csv:"Names"`
	Email       string `csv:"Email"`
	Destination string `csv:"Destination"`
	Date        string `csv:"Date"`
	Price       int    `csv:"Price"`
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	for _, tick := range b.Tickets {
		if tick.Id == t.Id {
			err := fmt.Sprint("can't create new ticket... ID: ", tick.Id, "is already in use")
			return Ticket{}, errors.New(err)
		}
	}

	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {

	for i := range b.Tickets {
		if b.Tickets[i].Id == id {
			return b.Tickets[i], nil
		}
	}

	err := fmt.Sprint("the ticket with ID:", id, "doesn't exist")
	return Ticket{}, errors.New(err)
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {

	for i := range b.Tickets {
		if b.Tickets[i].Id == id {
			b.Tickets[i] = t
			return t, nil
		}
	}

	err := fmt.Sprint("the ticket with ID:", id, "doesn't exist")
	return Ticket{}, errors.New(err)
}

func (b *bookings) Delete(id int) (int, error) {
	
	for i := range b.Tickets {
		if b.Tickets[i].Id == id {
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)
			fmt.Println("Deleted the ticket with id:", id)
			return id, nil
		}
	}
	err := fmt.Sprint("the ticket with ID:", id, "doesn't exist")
	return 0, errors.New(err)
}
