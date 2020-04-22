package arrays

import "testing"

func TestArrays(t *testing.T) {

	t.Run("soma todos os numeros de um array de inteiros com tamanho variado", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		soma := Soma(numeros)

		esperado := 45

		if soma != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", soma, esperado, numeros)
		}
	})

}
