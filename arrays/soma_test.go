package arrays

import (
	"reflect"
	"testing"
)

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

func TestSomaTudo(t *testing.T) {
	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v, esperado %v", resultado, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {

	verificarSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}
	}

	t.Run("Soma tudo menos o primeiro elemento", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		verificarSomas(t, resultado, esperado)
	})

	t.Run("Soma passando um array vazio", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{0, 9, 10})
		esperado := []int{0, 19}

		verificarSomas(t, resultado, esperado)
	})
}
