package main

import "fmt"

func main() {
	num := 0
	fmt.Println("Insira um número de 4 dígitos: ")
	fmt.Scan(&num)
	a := num / 1000
	b := num % 1000
	c := b / 100
	d := b % 100
	e := d / 10
	f := d % 10
	fmt.Printf("%v%v%v%v", f, e, c, a)
}

/*Faça um programa que leia um número inteiro de 4 dígitos e escreva-o invertido.
por exemplo, se o número lido for 2548, o resultado será 8452. Dica: utilize o
comando % que retorna o resto da divisão entre 2 números inteiros.*/
