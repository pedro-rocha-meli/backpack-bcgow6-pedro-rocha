package transactions

import (
	"fmt"

	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_web/clase_7_2/cmd/pkg/store"
)

type Transaction struct {
	Id       int     `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Sender   string  `json:"sender"`
	Date     string  `json:"date"`
}

var ID int //ID used in func LastID()

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code string, currency string, amount float64, sender string, date string) (Transaction, error)
	Update(id int, code string, currency string, amount float64, sender string, date string) (Transaction, error)
	Delete(id int) error
	LastID() (int, error)
	UpdateCode(id int, code string) (Transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Transaction, error) {
	var transactions []Transaction
	r.db.Read(&transactions)
	return transactions, nil
}

func (r *repository) Store(id int, code string, currency string, amount float64, sender string, date string) (Transaction, error) {

	var transactions []Transaction
	r.db.Read(&transactions)

	t := Transaction{
		Id:       id,
		Code:     code,
		Currency: currency,
		Amount:   amount,
		Sender:   sender,
		Date:     date,
	}
	transactions = append(transactions, t)

	if err := r.db.Write(transactions); err != nil {
		return Transaction{}, err
	}
	return t, nil
}

func (r *repository) Update(id int, code string, currency string, amount float64, sender string, date string) (Transaction, error) {

	var transactions []Transaction
	r.db.Read(&transactions)
	updated := false

	t := Transaction{
		Id:       id,
		Code:     code,
		Currency: currency,
		Amount:   amount,
		Sender:   sender,
		Date:     date,
	}

	for i := range transactions {
		if transactions[i].Id == id {
			t.Id = id
			transactions[i] = t
			updated = true
		}
	}

	if !updated {
		return Transaction{}, fmt.Errorf("Transaction with id: %d not found", id)
	}

	if err := r.db.Write(transactions); err != nil {
		return Transaction{}, err
	}

	return t, nil
}

func (r *repository) Delete(id int) error {

	var transactions []Transaction
	r.db.Read(&transactions)
	updated := false

	for i := range transactions {
		if transactions[i].Id == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			updated = true
		}
	}

	if !updated {
		return fmt.Errorf("Transaction with id: %d not found", id)
	}

	if err := r.db.Write(transactions); err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateCode(id int, code string) (Transaction, error) {

	var transactions []Transaction
	r.db.Read(&transactions)

	var t Transaction
	updated := false

	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Code = code
			updated = true
			t = transactions[i]
		}
	}

	if !updated {
		return Transaction{}, fmt.Errorf("Transaction with id: %d not found", id)
	}

	if err := r.db.Write(transactions); err != nil {
		return Transaction{}, err
	}

	return t, nil
}

func (r *repository) LastID() (int, error) {

	var transactions []Transaction
	if err := r.db.Read(&transactions); err != nil {
		return 0, err
	}
	if len(transactions) == 0 {
		return 0, nil
	}

	return transactions[len(transactions)-1].Id, nil
}
