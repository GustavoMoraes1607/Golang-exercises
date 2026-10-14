package main

import "fmt"

func main() {
	x := 0

	for x < 10 {
		x++
		fmt.Print(x, ": ")
		if x < 10 {
			fmt.Println("X é menor que 10")
		} else {
			fmt.Println("X é igual a 10")
			break
		}
	}
	fmt.Println("O Looping termina aqui!")
	fmt.Println("Reboot")
}
