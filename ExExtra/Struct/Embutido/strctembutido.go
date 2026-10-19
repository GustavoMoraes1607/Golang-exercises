package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}
type função struct {
	pessoa
	Função  string
	Salário int
}

func main() {
	trabalhador1 := pessoa{
		Nome:      "Carlos",
		Sobrenome: "Nogueira",
		Idade:     33,
	}
	trabalhador2 := função{
		pessoa: pessoa{
			Nome:      "Gustavo",
			Sobrenome: "Rocha",
			Idade:     25,
		},
		Função:  "Redator",
		Salário: 1200,
	}
	fmt.Println(trabalhador1)
	fmt.Println(trabalhador2)
}
