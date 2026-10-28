package main

import "fmt"

func main() {
	x := 125
	func(x int) {
		fmt.Println(x * 2)
	}(x)
}

/*Crie uma função anônima.*/
