// MAPS são parecidos com DICIONÁRIOS
// a sua definição de TYPE, chama-se maps
package main

import "fmt"

func main() {
	// adicionar make, que é uma função interna de GO
	// adicionar type pra essa função
	vertices := make(map[string]int) // agora esta função irá funcionar

	// podemos definir o conjunto de chaves e valores usando "[]"
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12
	// Pegar chaves e valores de map
	fmt.Println(vertices) // Output: "map[dodecagon:12 square:3 triangle:2]"

	// Usaremos a mesma sintaxe, só que para pegar o valor de uma chave
	fmt.Println(vertices["triangle"]) // Output: "2"

	// Para deletar, usamos a função "delete"
	delete(vertices, "square")
	fmt.Println(vertices) // Output: "map[dodecagon:12 triangle:2]"
}
