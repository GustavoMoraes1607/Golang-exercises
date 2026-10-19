package main

import "fmt"

type candidato struct {
	Nome      string
	Sobrenome string
	Idade     int
	Fumante   bool
}

func main() {
	candidato1 := candidato{
		Nome:      "Gustavo",
		Sobrenome: "Rocha",
		Idade:     25,
		Fumante:   false,
	}
	candidato2 := candidato{
		Nome:      "Joana",
		Sobrenome: "Correia",
		Idade:     36,
		Fumante:   true,
	}
	fmt.Println(candidato1)
	fmt.Println(candidato2)
}
