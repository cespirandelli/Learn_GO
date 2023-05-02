package main

import (
	"fmt"
)

// Podemos utilizar o FOR loop para "criar" um WHILE
func main() {
	// deletar a declaração de variável da função "i := 0"; deletar incremento "i++"
	// declarar variável fora do FOR loop
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

}
