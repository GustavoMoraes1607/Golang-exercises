package main

import "fmt"

func main() {
	kml := 0.0
	fmt.Println("Quantos km/l faz o seu carro?")
	fmt.Scan(&kml)
	temp := 0.0
	fmt.Println("Quanto tempo durou a viagem? (em minutos)")
	fmt.Scan(&temp)
	kmh := 0.0
	fmt.Println("Qual a velocidade média usada no percurso? (em km/h)")
	fmt.Scan(&kmh)
	dist := kmh * (temp / 60.0)
	val := dist / kml
	fmt.Printf("A Distancia percorrida foi de cerca de: %.2fkm\n", dist)
	fmt.Printf("Foram gastos cerca de: %.2f Litros de gasolina.", val)
}

/* Escreve um programa para determinar a quantidade de litros de combustivel gastos em uma
viagem por um automóvel que faz 12km/l. Para isso, sabe-se que que o tempo gasto na viagem
é de T=35 e a média da velocidade é V= 80km/h.*/
