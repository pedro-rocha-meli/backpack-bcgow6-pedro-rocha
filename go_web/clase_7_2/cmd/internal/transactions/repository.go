package transactions

type Transaction struct {
	Id       int     `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Sender   string  `json:"sender"`
	Date     string  `json:"date"`
}

var transactions []Transaction
var lastId int

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code string, currency string, amount float64, sender string, date string) (Transaction, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaction, error)  {
	return transactions, nil
}

func (r *repository) Store(id int, code string, currency string, amount float64, sender string, date string) (Transaction, error){
	t := Transaction{
		Id: id, 
		Code: code, 
		Currency: currency, 
		Amount: amount, 
		Sender: sender, 
		Date: date,
	}
	transactions = append(transactions, t)
	lastId = t.Id
	return t, nil
}

func (r *repository) LastId() (int, error){
	return lastId, nil
}

