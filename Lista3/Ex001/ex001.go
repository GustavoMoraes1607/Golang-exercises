package main

import "fmt"

func main() {
	x := [10]int{}
	soma := 0
	fmt.Println("Digite os 10 números que deseja inserir na array X: ")
	fmt.Scan(&x[0], &x[1], &x[2], &x[3], &x[4], &x[5], &x[6], &x[7], &x[8], &x[9])
	for _, v := range x {
		soma += v
	}
	media := soma / 10
	fmt.Println("A soma dos valores na é:", soma)
	fmt.Println("A média dos valores da array é:", media)
}

/*Faça um programa que leia 10 números digitados pelo usuário e exiba a soma e
a média desses valores.*/
