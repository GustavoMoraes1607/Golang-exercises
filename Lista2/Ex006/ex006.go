package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Escolha um valor para X.")
	fmt.Scan(&x)
	switch x {
	default:
		fmt.Println("Valor inválido")
	case 1:
		fmt.Println("X é igual a 1")
	case 5:
		fmt.Println("X é igual a 5")
	case 10:
		fmt.Println("X é igual a 10")
	}
}

/*Faça um programa capaz de identificar se um número é igual a 1, 5 ou 10.
Caso não seja igual a nenhum desses valores, retornar a mensagem "Valor inválido".*/
