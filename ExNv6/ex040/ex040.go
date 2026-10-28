package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Função dentro de variável.")
	}
	x()
}

/*Atribua uma função a uma variável.
Chame essa função.*/
