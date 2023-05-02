package main

import "fmt"

func main() {
	x := 7
	// Podemos omitir a declaração de variável
	// GO irá detectar quando utilizamos o ":" na chamada da função.

	if x > 6 {
		fmt.Println("Mais que 6")
	} else if x < 2 { // existe "else if" e "else"

	} else {

	}
}
