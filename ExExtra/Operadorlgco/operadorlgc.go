package main

import "fmt"

func main() {
	x := 1
	if x%2 == 0 && x%3 == 0 { // && é o operador lógico "E"
		fmt.Println("É multiplo de 2 e 3")
	} else if x%2 == 0 || x%3 == 0 { // || é o operador lógico "OU"
		fmt.Println("É multiplo de 2 ou 3")
	} else {
		fmt.Println("Não é multiplo de 2 nem de 3")
	}
}
