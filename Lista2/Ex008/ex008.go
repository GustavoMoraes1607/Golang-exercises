package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Dê um valor para X:")
	fmt.Scan(&x)
	if x > 0 {
		fmt.Println("X é positivo")
	} else if x < 0 {
		fmt.Println("X é negativo")
	} else {
		fmt.Println("X é igual a 0")
	}
}
