package main

import "fmt"

func main() {
	born := 2001
	today := 2026

	for {
		if born > today {
			break
		}
		fmt.Println("Ano: ", born)
		born++
	}
}

//Criar um loop utilizando a sintaxe: for { }
//Utiliza-la para imprimir os anos desde o ano
//em que nasci e interromper o loop com o break.
