package mocks

import "example.com/internal/domain"

type MockService struct{
	Repo []domain.Product
}

func (s *MockService) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	return s.Repo, nil
}
