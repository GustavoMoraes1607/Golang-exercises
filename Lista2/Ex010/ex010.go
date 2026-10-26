package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Qual ano gostaria de testar?")
	fmt.Scan(&x)
	if (x%400) == 0 || (x%4) == 0 && (x%100) != 0 {
		fmt.Println(x, "Este é um ano bissexto.")
	} else {
		fmt.Println("Este não é um ano bissexto")
	}
}

/*Faça um programa capaz de identificar se um número é um ano bissexto.
Considere que para o ano ser bissexto basta que seja divisível por 400.
Caso contrário, precisa ser divisível por 4 e não ser divisível por 100.
Faça uma condição composta que englobe todas as regras para a definição
do ano bissexto.*/
