package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func main() {
	mapa := make(map[string]pessoa)
	mapa["Rocha"] = pessoa{
		Nome:      "Gustavo",
		Sobrenome: "Rocha",
		Sabores:   []string{"Strawberry", "Pineapple", "Kinder"},
	}
	mapa["Nogueira"] = pessoa{
		Nome:      "Jaqueline",
		Sobrenome: "Nogueira",
		Sabores:   []string{"Vanilla", "Blueberry", "Mango"},
	}
	for _, v := range mapa {
		fmt.Println("My name is", v.Nome, v.Sobrenome, "and my fav icecream flavors are:")
		for _, vv := range v.Sabores {
			fmt.Println("-", vv)
		}
	}

}

//Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
//Demonstre os valores do map utilizando range.
//Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
