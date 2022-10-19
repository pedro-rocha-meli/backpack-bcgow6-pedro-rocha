package calculadora

import "errors"

func Restar(num1, num2 float64) float64 {
	return num1 - num2
}

func Dividir(num, den int) (result int, err error) {

	if den == 0 {
		err = errors.New("el denominador no puede ser 0")
	}
	result = num / den
	return
}
