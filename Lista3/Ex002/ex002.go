package main

import "fmt"

func main() {
	size := 0
	fmt.Println("Quantos valores serão inseridos?")
	fmt.Scan(&size)
	fmt.Printf("Insira os %v valores: ", size)
	x := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&x[i])
	}
	soma := 0
	for _, v := range x {
		soma += v
	}
	media := soma / size
	fmt.Println("A soma dos valores na é:", soma)
	fmt.Println("A média dos valores da array é:", media)
}

/*Altere o programa anterior, de tal maneira que o usuário informe a quantidade
de números que serão digitados (ou seja, o valor "10" não deve ser fixo no programa).*/
