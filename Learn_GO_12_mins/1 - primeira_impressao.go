// Adicionar "package main" no topo para que crie
// um executável único
// O que é chamado de "Standalone executable"
package main

// É importante ressaltar que GO tem sua biblioteca padrão
// Disponível em https://pkg.go.dev/std
// Nela iremos encontrar as bibliotecas que iremos utilizar
// ao longo do código.

// Math, é utilizado para matemática, FMT é para formatação
// NET é para rede, etc.
import "fmt"

// Declarar função principal que é onde seu programa
// irá começar a rodar.

func main() {
	fmt.Println("It's me, hi!")
	fmt.Println("Let's learn GO")
}

// Após escrito esse bloco de código, iremos rodar o comando
// "go build hello.go" no terminal
// Isso irá compilar o código fonte e suas dependências
// em um executável binário que irá aparecer à esquerda,
// na aba de arquivos como "hello.exe"

// podemos também linkar com pacotes remotos
// "remote packages" no github da seguinte forma
// "go mod init foo/bar", isso irá criar um novo módulo "foo/bar"
// depois irá adicionar os requerimentos do módulo
// Isso irá criar um arquivo "go.mod" que permitirá trackear
// as dependências
