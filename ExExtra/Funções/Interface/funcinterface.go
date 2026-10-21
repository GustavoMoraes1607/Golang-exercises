package main

import (
	"fmt"
	"math"
)

type forma interface { // INTERFACE
	Area() float64
}

type circulo struct { // STRUCT CIRCULO
	raio float64
}

func (c circulo) Area() float64 { // CALCULO ÁREA CIRCULO
	return math.Pi * c.raio * c.raio
}

type retangulo struct { // STRUCT RETANGULO
	largura float64
	altura  float64
}

func (r retangulo) Area() float64 { // CALCULO ÁREA RETANGULO
	return r.largura * r.altura
}

func printArea(s forma) { // REPRESENTEÇÃO EM TEXTO DO RESULTADO DA INTERFACE.
	fmt.Println("Area:", s.Area())
}

func main() {
	c := circulo{raio: 3}                 // DECLARAÇÃO DE VALORES PARA O TIPO CIRCULO
	r := retangulo{largura: 5, altura: 4} // DECLARAÇÃO DE VALORES PARA O TIPO RETANGULO

	printArea(c) // "PRINTAREA" PARA PUXAR A FUNC PRINTAREA. "(C)" PARA USAR OS VALORES DA DECLARAÇÃO C
	printArea(r) // "PRINTAREA" PARA PUXAR A FUNC PRINTAREA. "(R)" PARA USAR OS VALORES DA DECLARAÇÃO R
}
