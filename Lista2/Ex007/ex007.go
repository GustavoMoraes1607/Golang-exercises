package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Dê um valor para X:")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Println("X é par")
	} else {
		fmt.Println("X é impar")
	}
}

/*Faça um programa capaz de identificar se um número é par ou impar.*/
