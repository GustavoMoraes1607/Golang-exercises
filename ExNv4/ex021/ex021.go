package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	for i, v := range array {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", array)

}

//Usando uma literal composta: Crie um array que suporte 5 valores int.
//Atribua valores aos seus indices.
//Utilize o range para demonstrar todos os valores do array.
//Utilize format printing para demonstrar o tipo do array.
