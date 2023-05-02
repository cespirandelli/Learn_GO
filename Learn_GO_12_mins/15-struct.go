package main

import (
	"fmt"
)

// criar o tipo da estrutura
type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Jake", age: 23}
	fmt.Println(p)
	// se você quiser algum campo específico, usar da seguinte forma
	fmt.Println(p.age)
}
