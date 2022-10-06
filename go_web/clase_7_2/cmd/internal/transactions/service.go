package transactions

type Service interface {
	GetAll() ([]Transaction, error)
	Store(code string, currency string, amount float64, sender string, date string) (Transaction, error)
	Update(id int, code string, curreny string, amount float64, sender string, date string) (Transaction, error)
	UpdateCode(id int, code string) (Transaction, error)
	Delete(id int) error
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
	return s.repository.GetAll()
}

func (s *service) Store(code string, currency string, amount float64, sender string, date string) (Transaction, error) {
	id, err := s.repository.LastID()
	if err != nil {
		return Transaction{}, err
	}
	id++
	return s.repository.Store(id, code, currency, amount, sender, date)
}

func (s *service) Update(id int, code string, currency string, amount float64, sender string, date string) (Transaction, error) {
	return s.repository.Update(id, code, currency, amount, sender, date)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateCode(id int, code string) (Transaction, error) {
	return s.repository.UpdateCode(id, code)
}