package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Renato")
	esperado := "Olá, Renato"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
