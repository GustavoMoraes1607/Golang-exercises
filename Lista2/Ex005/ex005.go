package main

import "fmt"

func main() {
	x, y, z := 0, 0, 0
	eq := 0
	fmt.Println("Insira 3 valores:")
	fmt.Scan(&x, &y, &z)
	if x == y && y == z {
		eq += 3
	} else if x == y || x == z || y == z {
		eq += 2
	} else {
		eq = 0
	}
	fmt.Printf("Destes números, %v são iguais.", eq)
}

/* Faça um programa que leia 3 números e determine quantos são iguais.*/
