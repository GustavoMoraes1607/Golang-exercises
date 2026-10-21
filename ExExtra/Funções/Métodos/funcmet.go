package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (x pessoa) oichat() {
	fmt.Println(x.nome, "diz: Oi chat!")
}

func main() {
	Gustavo := pessoa{"Gustavo", 25}
	Gustavo.oichat()
}
