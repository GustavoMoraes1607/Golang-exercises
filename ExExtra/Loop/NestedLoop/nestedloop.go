package main

import "fmt"

func main() {
	for horas := 0; horas <= 23; horas++ {
		fmt.Println("\nHora: ", horas)
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Print(" ", minutos)
		}
	}
	for meses := 1; meses <= 12; meses++ {
		fmt.Println("\nMes: ", meses)
		for dias := 1; dias <= 30; dias++ {
			fmt.Print(" ", dias)
		}
	}
}
