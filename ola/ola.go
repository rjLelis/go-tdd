package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const ingles = "ingles"

const prefixoOlaFrances = "Bonjour, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaPortugues = "Olá, "
const prefixoOlaIngles = "Hello, "

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case ingles:
		prefixo = prefixoOlaIngles
	default:
		prefixo = prefixoOlaPortugues
	}

	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
