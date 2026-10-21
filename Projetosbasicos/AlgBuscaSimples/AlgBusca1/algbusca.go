package main

import "fmt"

func main() {
	x := [3]int{}
	fmt.Scanln(&x[0], &x[1], &x[2])
	y := 0
	fmt.Scanln(&y)
	iteracoes := 0
	for i, v := range x {
		iteracoes++
		if v == y {
			fmt.Println("O valor:", v, "se encontra no indice:", i)
			break
		}
	}
	fmt.Println(iteracoes)
}
