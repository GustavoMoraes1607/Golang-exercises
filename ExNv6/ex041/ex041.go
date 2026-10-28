package main

import "fmt"

func main() {
	x := ret()
	x()
}

func ret() func() {
	return func() {
		fmt.Println("Uma função que retorna uma função para dar valor a uma variável")
	}
}

/*Crie uma função que retorna uma função.
Atribua a função retornada a uma variável.
Chame a função retornada.*/
