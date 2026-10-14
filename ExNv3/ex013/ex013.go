package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for y := 1; y <= 3; y++ {
			fmt.Printf("\t%U '%c'\n", x, x)
		}
	}
}

// Objetivo: Colocar na tela o Unicode e o caracter de todas as
// letras maiúsculas do alfabeto, repetindo cada letra 3 vezes.
