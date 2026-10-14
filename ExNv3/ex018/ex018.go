package main

import "fmt"

func main() {
	nota := 6
	if nota < 6 {
		fmt.Println("Reprovado, estude mais!")
	} else if nota > 6 {
		fmt.Println("Parabéns, você foi aprovado!")
	} else {
		fmt.Println("Em cima da média, podemos melhorar!")
	}
}

//Utilizando a solução anterior, adicione
//as opções else e else if.
