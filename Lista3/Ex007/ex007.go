package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Escolha um número para determinar se é ou não um número primo.")
	fmt.Scan(&x)
	for y := 2; y < x; y++ {
		if x%y == 0 {
			fmt.Printf("%v não é um número primo.", x)
			return
		}
	}
	fmt.Printf("%v é um número primo.", x)
}

/*Faça um programa que leia um número inteiro positivo e determine se este é primo
ou não. Por definição, um número é primo quando é divisível somente por si próprio
e por 1.*/
