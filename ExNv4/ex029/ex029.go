package main

import "fmt"

func main() {
	Hobby := map[string][]string{
		"Gonçalo_João":    {"Pesca", "Corrida"},
		"Almeida_Roberto": {"Futebol", "Video-games"},
		"Freitas_Eduardo": {"Trilha", "Ciclismo"},
	}
	Hobby["Vieira_Ana"] = []string{"Desenhar", "Tiro ao alvo"}
	delete(Hobby, "Freitas_Eduardo")
	for k, v := range Hobby {
		fmt.Println(k)
		for i, val := range v {
			fmt.Println("\t", i, val)
		}
	}
}

//Utilizando o exercício anterior, adicione uma nova entrada ao map e remova uma antiga.
//Demonstre o map inteiro utilizando range.
