package main

import "fmt"

func main() {
	a := 0
	b := 0
	fmt.Println("Insira o primeiro número:")
	fmt.Scan(&a)
	fmt.Println("Insira o segundo número:")
	fmt.Scan(&b)
	c := a
	for d := 1; d < b; d++ {
		a += c
	}
	fmt.Println(a)
}

/*Faça um programa que calcule a multiplicação de 2 números inteiros sem utilizar
o operador "*". Em vez disso, utilize apenas o operador de adição.*/
