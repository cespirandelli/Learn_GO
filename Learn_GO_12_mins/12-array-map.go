package main

import (
	"fmt"
)

func main() {

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	// conseguimos fazer a mesma coisa que no exercício anterior só que utilizando map
	// iterar sobre cada elemento de um array ou um slice usando range

	for key, value := range m {
		// ao invés de chamar o índice e valor, usamos chave e valor
		fmt.Println("key", key, "value", value)
	}

}
