package main

import "fmt"

func main() {
	a := 0
	b := 0
	fmt.Println("Digite um valor para A:")
	fmt.Scan(&a)
	fmt.Println("Digite um valor para B:")
	fmt.Scan(&b)
	a, b = b, a
	fmt.Println("A:", a, "B:", b)
}

/*Faça um programa que recebe dois valores nas variáveis A e B e,
em seguida, troca o conteúdo destas variáveis. Por exemplo, se o
usuário digitar A=5 e B=3, o programa deverá trocar os valores de
tal maneira que A seja igual a 3 e B igual a 5.*/
