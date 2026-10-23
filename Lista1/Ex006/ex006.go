package main

import "fmt"

func main() {
	x := [4]int{}
	fmt.Println("Escolha os 4 números que deseja calcular a média:")
	for i := 0; i < 4; i++ {
		fmt.Scan(&x[i])
	}
	y := 0
	for _, valor := range x {
		y += valor
	}
	media := y / 4
	fmt.Println("A média de ", x, "é:", media)
}

/*Faça um programa que calcule a média de quatro números introduzidos pelo usuário.*/
