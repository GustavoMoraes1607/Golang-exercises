package main

import "fmt"

func main() {
	Hobby := map[string][]string{
		"Gonçalo_João":    {"Pesca", "Corrida"},
		"Almeida_Roberto": {"Futebol", "Video-games"},
		"Freitas_Eduardo": {"Trilha", "Ciclismo"},
	}
	for k, v := range Hobby {
		fmt.Println(k)
		for i, val := range v {
			fmt.Println("\t", i, val)
		}
	}
}

//Crie um map com key tipo string e value tipo []string.
//Key deve conter nomes no formato sobrenome_nome
//Value deve conter os hobbies favoritos da pessoa
//Demonstre todos esses valores e seus índices.
