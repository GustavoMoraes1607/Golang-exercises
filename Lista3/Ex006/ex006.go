package main

import "fmt"

func main() {
	n1 := 0
	n2 := 0
	fmt.Println("Insira um valor para o primeiro número:")
	fmt.Scan(&n1)
	fmt.Println("Insira um valor para o segundo número:")
	fmt.Scan(&n2)
	if n1 > n2 {
		fmt.Println("O primeiro valor tem de ser menor que o segundo valor.")
		return
	}
	x := 0
	for y := n1 + 1; y < n2; y++ {
		if y%2 == 1 {
			x += y
		}
	}
	fmt.Printf("A soma dos números impares do intervalo entre %v e %v é: %v", n1, n2, x)
}

/*Altere o programa anterior de tal maneira que quando o usuário digitar um intervalo
inválido (n1>n2), o programa irá solicitar novos valores para n1 e n2*/
