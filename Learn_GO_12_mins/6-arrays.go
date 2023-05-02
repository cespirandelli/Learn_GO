package main

import "fmt"

func main() {
	a := [5]int{5, 4, 3, 2, 1} // temos um problema com array com tamanho
	// a = append(a, 14) -> resulta em:
	// "# command-line-arguments
	//.\6-arrays2.go:7:13: first argument to append must be a slice; have a (variable of type [5]int)"

	// não dá para alterar o tamanho de um array
	// isso faz parte do type dele
	// para resolver isso, podemos usar slices
	fmt.Println(a, "Array 'finito'")

	b := []int{5, 4, 3, 2, 1}
	// para isso, é necessário remover o contador de elementos "5" de dentro do "[]"
	b = append(b, 13) // desta forma incluímos ao final do array o int 13, sem restrições
	// append não altera o slice original, ele retorna um novo na operação

	fmt.Println(b, "this is a slice of ints")
}
