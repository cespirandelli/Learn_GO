package main

import (
	"fmt"
)

func main() {
	i := 7
	inc(&i) //2º Direcionamos o ponteiro usando "&"
	fmt.Println(i)
	fmt.Println(&i)
}

// 1º Podemos aceitar um apontador prefixando o tipe com um " * "
func inc(x *int) {
	// 3º Devemos desreferenciar o ponteiro utilizando outro " * "
	*x++ // Se não fosse feito isso, iríamos incrementar o endereço da memória da variável
}
