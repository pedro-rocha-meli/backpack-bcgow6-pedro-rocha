package file

import (
	"fmt"
	"hackaton/internal/service"
	"os"

	"github.com/gocarina/gocsv"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tickets := []service.Ticket{}

	if err := gocsv.UnmarshalFile(file, &tickets); err != nil {
		panic(err)
	}

	return tickets, nil
}

func (f *File) Write(t service.Ticket) (error) {

	ticket := fmt.Sprintf(
		"\n%d,%s,%s,%s,%s,%d",
		t.Id,
		t.Names,
		t.Email,
		t.Destination,
		t.Date,
		t.Price,
	)
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.WriteString(ticket); err != nil {
		return err
	}

	return nil
}
