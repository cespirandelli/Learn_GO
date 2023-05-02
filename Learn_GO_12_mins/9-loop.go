package main

import (
	"fmt"
)

// O único loop em GO é o "FOR loop"
func main() {
	// inicializar com a chamada reduzida (shorthand syntax)
	// i é variável; contagem inicia no "0"; condição enquanto i for menor que 5; adicionar 1 ao final
	// "fechar" função com "{}"
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
