package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func main() {
	pessoa1 := pessoa{
		Nome:      "Gustavo",
		Sobrenome: "Rocha",
		Sabores:   []string{"Strawberry", "Pineapple", "Kinder"},
	}
	pessoa2 := pessoa{
		Nome:      "Jaqueline",
		Sobrenome: "Nogueira",
		Sabores:   []string{"Vanilla", "Blueberry", "Mango"},
	}
	fmt.Println(pessoa1.Nome, pessoa1.Sobrenome)
	for _, v := range pessoa1.Sabores {
		fmt.Println("-", v)
	}
	fmt.Println(pessoa2.Nome, pessoa2.Sobrenome)
	for _, v := range pessoa2.Sabores {
		fmt.Println("-", v)
	}
}

//Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//Nome, Sobrenome, Sabores favoritos de sorvete.
//Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
