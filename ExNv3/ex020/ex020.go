package main

import "fmt"

func main() {
	switch esportefavorito := "Futebol"; esportefavorito {
	case "Futebol":
		fmt.Println("Meu esporte favorito é Futebol.")
	case "Vôlei":
		fmt.Println("Meu esporte favorito é Vôlei.")
	case "Natação":
		fmt.Println("Meu esporte favorito é Natação")
	}
}

//Crie um programa que utilize a declaração SWITCH, onde o
//switch statemente seja uma váriavel do tipo string e com
//o identificador "esportefavorito".
