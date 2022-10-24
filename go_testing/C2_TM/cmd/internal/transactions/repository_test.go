package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen.
Comprobar que GetAll() retorne la información exactamente igual a la esperada.
Para esto:
    1. Dentro de la carpeta /internal/products, crear un archivo repository_test.go con el test diseñado.
*/

type StubStore struct {
	data []Transaction
}

func (s *StubStore) Read(data interface{}) (err error) {
	castedData := data.(*[]Transaction)
	*castedData = s.data
	return
}
func (s *StubStore) Write(data interface{}) (err error) {
	castedData := data.([]Transaction)
	s.data = append(s.data, castedData...)
	return
}

func TestGetAll(t *testing.T) {

	//Arrange.
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

	//Act.
	myStub := StubStore{data: expectedTransactions}
	repo := NewRepository(&myStub)

	result, err := repo.GetAll()

	//Assert.
	assert.Nil(t, err)

	assert.Equal(t, expectedTransactions, result)
}

/*
Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
    1. Crear un mock de Storage, dicho mock debe contener en su data un producto específico cuyo nombre puede ser “Before Update”.
    2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través de un boolean como se observó en la clase.
    3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya sido ejecutado durante el test.
*/

type MockStore struct {
	data           []Transaction
	readWasCalled  bool
	writeWasCalled bool
}

func (s *MockStore) Read(data interface{}) (err error) {
	s.readWasCalled = true
	castedData := data.(*[]Transaction)
	*castedData = s.data
	return
}
func (s *MockStore) Write(data interface{}) (err error) {
	s.writeWasCalled = true
	castedData := data.([]Transaction)
	s.data = append(s.data, castedData...)
	return
}

func TestUpdateCode(t *testing.T) {
	data := []Transaction{
		{
			Id:       1,
			Code:     "Before Update",
			Currency: "ETH",
			Amount:   0.05432,
			Sender:   "newTrader",
			Date:     "12-03-2022",
		},
	}

	expected := data[0]

	updatedCode := "Updated Name"

	expected.Code = updatedCode

	db := &MockStore{data: data, readWasCalled: false}

	repo := NewRepository(db)

	out, err := repo.UpdateCode(1, updatedCode)

	assert.Nil(t, err)
	assert.Equal(t, expected, out)
	assert.True(t, db.readWasCalled)
}

/*
C3 TM - Ejercicio 3

Para mejorar la calidad del código se debe aumentar el coverage del proyecto Go - Web, seguir los siguientes pasos:

Después de identificar las partes del código que no tienen cobertura, se debe elegir una o más partes del código donde sea razonable aumentar la cobertura.
Aumentar la cobertura de las partes elegidas.
Comparar el code coverage y el coverage report actual contra el anterior.
*/

func TestStore(t *testing.T) {

	data := []Transaction{}

	newTransaction := Transaction{
		Id:       1,
		Code:     "ABX-832",
		Currency: "SOL",
		Amount:   21.453,
		Sender:   "newTrader",
		Date:     "15-03-2022",
	}

	db := &MockStore{data: data, readWasCalled: false, writeWasCalled: false}

	repo := NewRepository(db)

	out, err := repo.Store(
		newTransaction.Id,
		newTransaction.Code,
		newTransaction.Currency,
		newTransaction.Amount,
		newTransaction.Sender,
		newTransaction.Date,
	)

	assert.Nil(t, err)
	assert.True(t, db.writeWasCalled)
	assert.Equal(t, out, newTransaction)
	assert.Equal(t, newTransaction, db.data[0])
}

func TestLastID(t *testing.T){

	data := []Transaction{
		{
			Id:       5,
			Code:     "Before Update",
			Currency: "ETH",
			Amount:   0.05432,
			Sender:   "newTrader",
			Date:     "12-03-2022",
		},
	}

	db := &MockStore{data: data, readWasCalled: false, writeWasCalled: false}

	expected := 5

	repo := NewRepository(db)

	out, err := repo.LastID()

	assert.Nil(t, err)
	assert.Equal(t, expected, out)
}