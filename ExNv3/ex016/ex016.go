package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		y := x % 4
		fmt.Printf("O resto de %v dividido por 4 é %v\n", x, y)
	}
}

//Demonstre o resto da divisão por 4 de todos
//os números de 10 a 100.
