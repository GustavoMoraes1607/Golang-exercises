package main

import "fmt"

func main() {
	slcint := []int{10, 23, 46, 7}
	resultado := soma(slcint...) // Sem "..." para abrir o slice, não funcionaria.
	fmt.Println(resultado)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
