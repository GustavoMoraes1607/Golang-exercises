package main

import "fmt"

func main() {
	slice := []string{"banana", "maça", "jaca", "pêssego"}
	slice = append(slice, "melancia", "mamão", "morango", "uva")
	for indice, valor := range slice {
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}
	for _, valor := range slice {
		fmt.Printf("um dos valores deste slice é: %v.\n", valor)
	}
	for indice, _ := range slice {
		fmt.Printf("um dos indices deste slice é: %v.\n", indice)
	}
	//
	// Exemplo 2 for range.
	//
	slice2 := []int{20, 21, 22, 23, 1, 13}
	total := 0
	for _, valor := range slice2 {
		// "+=" é o mesmo que "total = total + valor"
		total += valor
	}
	fmt.Println("O valor total é", total)
}
