package ordenamiento

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestOrdenar(t *testing.T) {
	
	result := Ordenar([]int{1,5,4,3,7,8,34,9,10})
	expected := []int{1,3,4,5,7,8,9,10,34}
	assert.Equal(t, expected, result)
}