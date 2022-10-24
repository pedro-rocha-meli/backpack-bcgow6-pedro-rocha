package mocks

import (
	"errors"

	"example.com/internal/domain"
)



type MockRepository struct {
	Data []domain.Product
}

func (r *MockRepository) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	var prodList []domain.Product

	for _, v := range r.Data {
		if sellerID == v.SellerID {
			prodList = append(prodList, v)
		}
	}

	if len(prodList) == 0 {
		return []domain.Product{}, errors.New("not found")
	}

	return prodList, nil
}