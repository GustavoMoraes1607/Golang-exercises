package main

import "fmt"

func main() {
	y := plus()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
}
func plus() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

/*Demonstre o funcionamento de um closure.
Ou seja: crie uma função que retorna outra
função, onde esta outra função faz uso de
uma variável alem de seu scope.*/
