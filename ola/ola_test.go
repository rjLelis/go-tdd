package main

import "testing"

func TestOla(t *testing.T) {

	verificarMensagemCorrenta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Renato", "")
		esperado := "Olá, Renato"

		verificarMensagemCorrenta(t, resultado, esperado)
	})

	t.Run("diz Olá, mundo quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificarMensagemCorrenta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("renato", "espanhol")
		esperado := "Hola, renato"
		verificarMensagemCorrenta(t, resultado, esperado)
	})

	t.Run("em frances", func(t *testing.T) {
		resultado := Ola("Renato", "frances")
		esperado := "Bonjour, Renato"
		verificarMensagemCorrenta(t, resultado, esperado)
	})

	t.Run("em ingles", func(t *testing.T) {
		resultado := Ola("Renato", "ingles")
		esperado := "Hello, Renato"
		verificarMensagemCorrenta(t, resultado, esperado)
	})
}
