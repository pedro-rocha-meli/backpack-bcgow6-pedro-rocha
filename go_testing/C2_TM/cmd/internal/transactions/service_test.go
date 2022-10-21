package transactions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/*Ejercicio 1 - Service/Repo/Db Update()

Diseñar un test que pruebe en la capa service, el método o función Update(). Para lograrlo se deberá:
	1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
	2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
	3. Para dar el test como OK debe validarse que al invocar el método del Service Update(),  retorne el producto con
	mismo Id y los datos actualizados. Validar también que  Read() del Store haya sido ejecutado durante el test.
*/

type MockStorage struct {
	data       []Transaction
	readWasCalled  bool
	writeWasCalled bool
}

func (m *MockStorage) Read(data interface{}) error {
	m.writeWasCalled = true
	a := data.(*[]Transaction)
	*a = m.data
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	m.readWasCalled = true
	a := data.([]Transaction)
	m.data = a
	return nil
}

func TestServiceUpdate(t *testing.T) {
	transcation := Transaction{
		Id:       1,
		Code:     "AMX-809",
		Currency: "ETH",
		Amount:   0.05432,
		Sender:   "newTrader",
		Date:     "12-03-2022",
	}
	
	updated := Transaction{
		Id:       1,
		Code:     "AMX-809",
		Currency: "BTC",
		Amount:   0.3241,
		Sender:   "otherTrader",
		Date:     "02-05-2022",
	}

	db := []Transaction{
		transcation,
	}
	mockedStorage := MockStorage{
		data: db,
	}

	repo := NewRepository(&mockedStorage)
	service := NewService(repo)

	result, err := service.Update(
		updated.Id, 
		updated.Code, 
		updated.Currency, 
		updated.Amount, 
		updated.Sender, 
		updated.Date,
	)
	
	assert.Nil(t, err)
	assert.Equal(t, result, updated)
	assert.True(t, mockedStorage.readWasCalled)
}

/*
Diseñar un test que pruebe en la capa service, el método o función Delete(). Se debe probar la correcta eliminación de un producto, y el error cuando el producto no existe. 
Para lograrlo puede:
Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
Ejecutar el test con dos id’s de producto distintos, siendo uno de ellos un id inexistente en el Mock de Storage.
Para dar el test como OK debe validarse que efectivamente el producto borrado ya no exista en Storage luego del Delete(). También que cuando se intenta borrar un producto  inexistente, se debe obtener el error correspondiente.
*/

func TestServiceDelete(t *testing.T) {
	transaction1 := Transaction{
		Id:       1,
		Code:     "AMX-809",
		Currency: "ETH",
		Amount:   0.05432,
		Sender:   "newTrader",
		Date:     "12-03-2022",
		}
	transaction2 := Transaction{
		Id:       2,
		Code:     "AMX-809",
		Currency: "BTC",
		Amount:   0.3241,
		Sender:   "otherTrader",
		Date:     "02-05-2022",
		}

	db := []Transaction{transaction1, transaction2}

	mockedStorage := MockStorage{data: db}
	repo := NewRepository(&mockedStorage)
	service := NewService(repo)

	err := service.Delete(2)
	assert.Nil(t, err)
	assert.True(t, mockedStorage.readWasCalled)
	assert.True(t, mockedStorage.writeWasCalled)
	assert.Equal(t, mockedStorage.data, []Transaction{transaction1})

	err = service.Delete(2)
	assert.NotNil(t, err)
	assert.True(t, mockedStorage.readWasCalled)
	assert.True(t, mockedStorage.writeWasCalled)
	assert.Equal(t, mockedStorage.data, []Transaction{transaction1})
}