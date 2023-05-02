package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// Até agora utilizamos tudo dentro da função main, mas podemos criar outras
// para retornar um valor, é necessário especificar o tipo após fechar "()"
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Unfefined for negative numbers !")
	}
	return math.Sqrt(x), nil
}
