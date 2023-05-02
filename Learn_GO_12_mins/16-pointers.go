package main

import (
	"fmt"
)

func main() {
	i := 7
	inc(i)
	fmt.Println(i)
	// podemos pegar o endereço de memória de uma variável conforme abaixo
	fmt.Println(&i)
}

// vamos criar uma função que incrementa uma variável

func inc(x int) {
	x++ // por não retornar nada, desta forma fica inútil este código
	// no código 17 iremos utilizar ponteiros para fazer funcionar
}
