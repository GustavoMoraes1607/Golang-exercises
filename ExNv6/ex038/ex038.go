package main

import (
	"fmt"
	"math"
)

type figura interface {
	area()
}

func info(f figura) {
	f.area()
}

type quadrado struct {
	lado float64
}

func (q quadrado) area() {
	total := q.lado * q.lado
	fmt.Println("A área do quadrado é:", total)
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	total := 2 * math.Pi * c.raio
	fmt.Println("A área do circulo é:", total)
}
func main() {
	q := quadrado{
		lado: 20,
	}
	c := circulo{
		raio: 20,
	}
	info(q)
	info(c)
}

/*Crie um tipo "quadrado"
Crie um tipo "círculo"
Crie um método "área" para cada tipo que calcule e retorne a área da figura
Área do círculo: 2 * π * raio
Área do quadrado: lado * lado
Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
Crie um valor de tipo "quadrado"
Crie um valor de tipo "círculo"
Use a função "info" para demonstrar a área do "quadrado"
Use a função "info" para demonstrar a área do "círculo"*/
