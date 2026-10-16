package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice[:3])
	fmt.Println(slice[1:7])
	fmt.Println(slice[4:])
	fmt.Println(slice[4:len(slice)])

}

//Utilizando o exercício anterior, use slicing para demonstrar os valores:
//Do primeiro ao terceiro item do slice
//Do segundo ao sétimo item do slice.
//Do quinto ao último item do slice.
//Desafio: Obtenha o mesmo resultado acima utilizando a função len() para determinar o penultimo item
