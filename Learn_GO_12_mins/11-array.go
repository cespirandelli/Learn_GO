package main

import (
	"fmt"
)

func main() {

	arr := []string{"a", "b", "c"}

	// iterar sobre cada elemento de um array ou um slice usando range
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

}
