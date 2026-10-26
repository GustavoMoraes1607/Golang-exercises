package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Escolha um valor para X.")
	fmt.Scan(&x)
	if x == 1 {
		fmt.Println("X é igual a 1")
	} else if x == 5 {
		fmt.Println("X é igual a 5")
	} else if x == 10 {
		fmt.Println("X é igual a 10")
	} else {
		fmt.Println("Valor inválido.")
	}
}

/*Faça um programa capaz de identificar se um número é igual a 1, 5 ou 10.
Caso não seja igual a nenhum desses valores, retornar a mensagem "Valor inválido".*/
