package main

import "fmt"

func main() {
	var num1, num2 float64
	var op string

	fmt.Println("Digite na ordem (número operador número): ")
	fmt.Scanln(&num1, &op, &num2)

	switch op {
	case "+":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, op, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, op, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, op, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Erro: Divisão por zero não é permitida.")
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f", num1, op, num2, num1/num2)
		}
	default:
		fmt.Println("Operador Inválido")
	}
}
