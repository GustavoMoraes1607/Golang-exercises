package main

import "fmt"

func main() {
	x := 0
	y := 0
	fmt.Println("Escolha um número para X:")
	fmt.Scan(&x)
	fmt.Println("Escolha um número para Y:")
	fmt.Scan(&y)
	if x > y {
		fmt.Printf("X(%v) é maior que Y(%v)", x, y)
	} else {
		fmt.Printf("X(%v) é menor que Y(%v)", x, y)
	}
}

/* Faça um programa que leia dois números inteiros e determine qual deles é maior.
Considere que os dois números serão diferentes.*/
