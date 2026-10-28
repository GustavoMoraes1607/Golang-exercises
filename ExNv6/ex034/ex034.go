package main

import "fmt"

func main() {
	fmt.Println(rti())
	fmt.Println(rtis())

}

func rti() int {
	return 10
}

func rtis() (int, string) {
	return 101, "String"
}

/*Crie uma função que retorne um int.
Crie outra função que retorne um int e uma string.
Chame as duas funções.
Demonstre seus resultados.*/
