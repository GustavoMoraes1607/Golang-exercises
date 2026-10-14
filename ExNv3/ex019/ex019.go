package main

import "fmt"

func main() {
	Quiz := 2
	switch Quiz {
	case 1:
		fmt.Println("Resposta errada, tente novamente!")
	case 2:
		fmt.Println("Parabéns, a sua resposta está correta")
	case 3:
		fmt.Println("Resposta errada, tente novamente!")
	}
}

//Crie um programa que utilize a declaração SWITCH
//sem um switch statement especificado.
