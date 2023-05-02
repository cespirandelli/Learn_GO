package main

import "fmt"

func main() {
	var a [5]int //variável não declarada, recebe zeros
	// Isso irá criar um array que conterá 5 inteiros
	// Todas as variáveis serão inicializadas com seu valor de tipo 0
	fmt.Println(a, "não declarado")

	a[2] = 99
	// Desta forma indicamos que queremos que o terceiro elemento do array
	// tem que ser 99
	fmt.Println(a, "atribuição do terceiro elemento") // Output: [0 0 99 0 0], quatro zeros e 99 na terceira posição

	a[2] = 7
	// Se declarado após o print, não será exibido no terminal e irá atribuir da mesma forma
	fmt.Println(a, "mudança do terceiro elemento") // Output: [0 0 7 0 0], quatro zeros

	// Vamos diminuir, ou simplificar exatamente igual acima, só que declarando variáveis

	b := [5]int{5, 4, 3, 2, 1} // tiramos o var e colocamos ":" para inferir o type
	// e "{}" para declarar elementos do array
	fmt.Println(b, "variável declarada")
}
