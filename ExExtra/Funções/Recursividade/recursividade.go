package main

import "fmt"

func main() {
	fmt.Println(fatorial(10))
	fmt.Println(loop(10))
}

func fatorial(x int) int { //Isto é uma função recursiva (uma funçao que chama ela, dentro de ela mesma )
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func loop(x int) int { //Isto é um loop dentro de uma função(na maioria das vezes é preferível pois usa menos memória)
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
