package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

//Começando com a slice: x:=[]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//Adicione o valor 52 ao final da slice;
//Agora adicione os valores 53,54 e 55 utilizando somente uma declaração;
//Demonstre a slice x.
//Agora anexe a ela a seguinte slice:y := []int{56, 57, 58, 59, 60}
//Demonstre novamente a slice x.
