package main

import "fmt"

func main() {
	slry := 0.0  // Valor do salário minimo.
	aguaL := 0.0 // Quantidade de água
	fmt.Println("Insira o valor do salário mínimo:")
	fmt.Scan(&slry)
	fmt.Println("Insira a quantidade gasta no mês:")
	fmt.Scan(&aguaL)
	aguaKL := aguaL / 1000.0      // Conversão L em KL
	val := (0.02 * slry) * aguaKL // Valor da conta
	desc := val - (val * 0.15)    // Desconto de 15%
	fmt.Println("O valor total da conta de água é:", val)
	fmt.Println("Com o desconto de 15% fica:", desc)
}

/*Sabe-se que o valor de cada 1000 litros de água corresponde a 2% do salário mínimo.
Faça um programa para receber o valor do salário minimo e a quantidade de água consumida em uma residencia por mês.
O algoritmo deverá calcular e mostrar A) o valor da conta de água. B) o valor a ser pago com desconto de 15%.*/
