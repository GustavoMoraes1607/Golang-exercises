package main

import "fmt"

func main() {
	slice := [][]string{
		{
			"Gustavo",
			"Rocha",
			"Video-games",
		},
		{
			"Jamille",
			"Araújo",
			"Desenhar",
		},
		{
			"João",
			"Elias",
			"Filmes e Séries",
		},
	}
	for _, v := range slice {
		fmt.Println(v[0], v[1])
		fmt.Println("\t", "Nome:", v[0])
		fmt.Println("\t", "Sobrenome:", v[1])
		fmt.Println("\t", "Hobby:", v[2])
	}
}

//Crie uma slice contendo slices de strings ([][]string).
//Atribua valores a este slice multi-dimensional da seguinte maneira:
//"Nome", "Sobrenome", "Hobby favorito"
//Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
