package iteracao

import "strings"

func Repetir(caractere string, loops int) string {
	if loops < 0 {
		loops = 1
	}
	return strings.Repeat(caractere, loops)
}
