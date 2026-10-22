package main

import (
	"fmt"
	"math"
)

func main() {
	b := 0.0 //Cateto B
	c := 0.0 //Cateto C
	fmt.Println("Dê um valor para o cateto B:")
	fmt.Scan(&b)
	fmt.Println("Dê um valor para o cateto C:")
	fmt.Scan(&c)
	a := (b * b) + (c * c)
	a = math.Sqrt(a) //Valor Hipotenusa (a² = b² + c²)
	fmt.Println("O valor da hipotenusa é:", a)
}

/*Num triângulo retângulo, segundo Pitágoras, o quadrado da hipotenusa(a)
é igual a soma dos quadrados dos catetos (b e c). Faça um algoritmo que
recebe o valor dos catetos e imprime o valor da hipotenusa.*/
