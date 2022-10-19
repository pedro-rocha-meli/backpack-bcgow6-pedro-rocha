package calculadora

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRestarDosNumeros(t *testing.T) {
	result := Restar(22, 12)
	assert.Equal(t, result, float64(10))
}

func TestDividirDosNumeros(t *testing.T){
	result, err := Dividir(10, 5)
	expected := 2
	assert.Nil(t, err)
	assert.Equal(t, result, expected)
}

