package main

import "fmt"

func main() {
	x, y := 0, 0
	op := ""
	fmt.Println("Defina os valores que deseja calcular (número operador número):")
	fmt.Scan(&x, &op, &y)
	switch op {
	case "+":
		fmt.Printf("%v %v %v = %v", x, op, y, x+y)
	case "-":
		fmt.Printf("%v %v %v = %v", x, op, y, x-y)
	case "*":
		fmt.Printf("%v %v %v = %v", x, op, y, x*y)
	case "/":
		if y == 0 {
			fmt.Println("Divisão por zero não é permitida.")
		} else {
			fmt.Printf("%v %v %v = %v", x, op, y, x/y)
		}
	}
}

/*Faça um algoritmo que simule uma calculadora com as quatro operções básicas
(+,-,*,/). O algoritmo deve solicitar ao usuário a entrada de dois operandos
e a operação a ser executada, na forma de menu. Dependendo da opção escolhida,
deve ser executada a operação solicitada e escrito seu resultado. Utilize uma
variável do tipo caractere para armazenar a operação e utilize o comando caso
para escolher a operação a partir do operador.*/
