package main

import "fmt"

func main() {
	result := sum(2, 3)
	fmt.Println(result)
}

// Até agora utilizamos tudo dentro da função main, mas podemos criar outras
// para retornar um valor, é necessário especificar o tipo após fechar "()"
func sum(x int, y int) int {
	return x + y
}
