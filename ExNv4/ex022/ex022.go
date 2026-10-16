package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", slice)
}

//Utilizando uma literal composta: Crie uma Slice Int.
//Atribua 10 valores a ela.
//Utilize range para demonstrar todos esses valores.
//Utilize format printing para demonstrar seu tipo.
