package main

import "fmt"

func main() {
	x, y, z := 0, 0, 0
	fmt.Println("Insira 3 valores para serem ordenados:")
	fmt.Scan(&x, &y, &z)
	if x >= y && y >= z {
		fmt.Println(z, y, x)
	} else if x <= y && y <= z {
		fmt.Println(x, y, z)
	} else if y >= x && x >= z {
		fmt.Println(z, x, y)
	} else if y <= x && x <= z {
		fmt.Println(y, x, z)
	} else if y >= z && z >= x {
		fmt.Println(x, z, y)
	} else if y <= z && z >= y {
		fmt.Println(y, z, x)
	}
}

/*Implemente um programa para ordenar 3 números inteiros.*/
