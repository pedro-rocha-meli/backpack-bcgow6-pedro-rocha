package transactions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	data []Transaction
}

func (s StubStore) Read(data interface{}) (err error) {
	a := data.(*[]Transaction)
	*a = s.data
	return
}
func (s StubStore) Write(data interface{}) (err error) {
	a := data.([]Transaction)
	s.data = append(s.data, a...)
	return
}

func TestGetAll(t *testing.T) {

	//Arrange
	expectedTransactions := []Transaction{
		{
			Id:       1,
			Code:     "AW-054",
			Currency: "ETH",
			Amount:   0.05432,
			Sender:   "newTrader",
			Date:     "12-03-2022",
		},
	}

	myStub := StubStore{data: expectedTransactions}
	repo := NewRepository(myStub)

	result, err := repo.GetAll()

	assert.Nil(t, err)

	assert.Equal(t, expectedTransactions, result)
}
