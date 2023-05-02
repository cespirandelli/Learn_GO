// Se você quer que o código seja executável, inserir package main no início
package main

// importar bibliotecas/pacotes
import (
	"fmt"
)

// Incluir função MAIN
// Para declarar variáveis, seguir modelo abaixo
func main() {
	var x int // var é uma KEYWORD, x é o nome da variável, int é o type
	x = 5     // é possível declarar o valor da variável, fora de sua linha (12 neste caso)

	var y int = 7 // também pode ser declarada na mesma linha

	// caso uma variável seja declarada e não utilizada, aparecerá uma mensagem no terminal
	// exemplo: ".\hello2.go:15:6: y declared and not used"

	var sum int = x + y

	fmt.Println(sum)
	// Toda vez que não é declarado o valor, será 0
	// Todo type tem um valor 0 (zero)
	// para string, a string "zero" é vazia.
}
