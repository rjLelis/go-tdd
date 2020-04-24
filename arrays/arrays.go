package arrays

func Soma(numeros []int) (soma int) {
	for _, numero := range numeros {
		soma += numero
	}

	return
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}

	return
}

func SomaTodoOResto(numerosParaSomar ...[]int) (somas []int) {
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}
	return
}
