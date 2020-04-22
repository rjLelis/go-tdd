package arrays

func Soma(numeros []int) (soma int) {
	for _, numero := range numeros {
		soma += numero
	}

	return
}
