package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	t.Run("concatena 10 'a's sequencialmente", func(t *testing.T) {
		repeticoes := Repetir("a", 10)
		esperado := "aaaaaaaaaa"

		if repeticoes != esperado {
			t.Errorf("resultado %s, esperado %s", repeticoes, esperado)
		}
	})

	t.Run("Retorna 'a' se o numero de loops for menor que zero", func(t *testing.T) {
		repeticoes := Repetir("a", -2)
		esperado := "a"

		if repeticoes != esperado {
			t.Errorf("resultado %s, esperado %s", repeticoes, esperado)
		}
	})
}

func ExampleRepetir() {
	repeticao := Repetir("a", 15)
	fmt.Println(repeticao)
	// Output: aaaaaaaaaaaaaaa
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 10)
	}
}
