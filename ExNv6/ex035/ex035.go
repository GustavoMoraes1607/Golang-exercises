package main

import "fmt"

func main() {
	si1 := []int{1, 3, 5, 7, 9}
	fmt.Println(soma1(si1...))
	si2 := []int{2, 4, 6, 8, 10}
	fmt.Println(soma2(si2))
}

func soma1(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func soma2(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

/*Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
Passe um valor do tipo slice of int como argumento para a função;
Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
Passe um valor do tipo slice of int como argumento para a função.*/
