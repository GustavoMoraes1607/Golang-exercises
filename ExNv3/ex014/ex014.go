package main

import "fmt"

func main() {
	born := 2001
	today := 2026

	for born <= today {
		fmt.Println("Ano: ", born)
		born++
	}
}

//Criar um loop utilizando a sintaxe: for condition { }
//Utiliza-la para imprimir os anos desde o ano em que nasci.
