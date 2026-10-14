package main

import "fmt"

func main() {
	if x := 5; x > 100 {
		fmt.Println("X é maior que 100")
	} else if x < 10 {
		fmt.Println("X é menor do que 10")
	} else {
		fmt.Println("X é igual a 10")
	}
}
