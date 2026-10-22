package main

import "fmt"

func main() {
	cons := 12.0
	dist := 80.0 * (35.0 / 60.0)
	val := dist / cons
	fmt.Println(dist, val)
}

/* Escreve um programa para determinar a quantidade de litros de combustivel gastos em uma
viagem por um automóvel que faz 12km/l. Para isso, sabe-se que que o tempo gasto na viagem
é de T=35 e a média da velocidade é V= 80km/h.*/
