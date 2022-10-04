package transactions

type Service interface {
	GetAll() ([]Transaction, error)
	Store(id int, code string, currency string, amount float64, sender string, date string) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Transaction, error) {
	transactions, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *service) Store(id int, code string, currency string, amount float64, sender string, date string) (Transaction, error) {
	lastId, err := s.repository.LastId()

	if err != nil {
		return Transaction{}, err
	}

	lastId++

	transaction, err := s.repository.Store(id, code, currency, amount, sender, date)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}
