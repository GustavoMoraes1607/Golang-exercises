package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Escolha o número para calcular o fatorial: ")
	fmt.Scan(&x)
	for y := x - 1; y > 0; y-- {
		x *= y
	}
	if x <= 0 {
		fmt.Println("O valor escolhido é inválido.")
	} else {
		fmt.Println("O resultado é:", x)
	}

}

/*Faça um programa que calcule o fatorial de um número. Se o número for menor do
que zero, então o programa deverá informar ao usuário que o valor é inválido.*/
