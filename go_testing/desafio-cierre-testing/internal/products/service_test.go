package products

import (
	"errors"
	"testing"

	"example.com/internal/domain"
	"example.com/test/mocks"
	"github.com/stretchr/testify/assert"
)



func TestGetAllBySeller(t *testing.T) {

	expected := []domain.Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}

	repo := mocks.MockRepository{
		Data: []domain.Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		},
	}

	service := NewService(&repo)

	out, err := service.GetAllBySeller("FEX112AC")

	assert.Nil(t, err)
	assert.Equal(t, expected, out)
}

func TestGetAllBySellerWithError(t *testing.T){

	expected := errors.New("not found")

	repo := mocks.MockRepository{
		Data: []domain.Product{},
	}

	service := NewService(&repo)

	out, err := service.GetAllBySeller("FEX112AC")

	assert.Empty(t, out)
	assert.Equal(t, expected, err)
}