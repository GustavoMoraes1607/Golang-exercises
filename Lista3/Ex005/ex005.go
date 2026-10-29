package main

import "fmt"

func main() {
	n1 := 0
	n2 := 0
	fmt.Println("Insira um valor para o primeiro número:")
	fmt.Scan(&n1)
	fmt.Println("Insira um valor para o segundo número:")
	fmt.Scan(&n2)
	x := 0
	for y := n1 + 1; y < n2; y++ {
		if y%2 == 1 {
			x += y
		}
	}
	fmt.Printf("A soma dos números impares do intervalo entre %v e %v é: %v", n1, n2, x)
}

/*Faça um programa que solicite ao usuário dois números inteiros diferentes "n1" e "n2"
e calcule a soma de todos os números Ímpares dentro do intervalo definido por
[n1,...,n2]. Considere que n1 é sempre menor do que n2.*/
