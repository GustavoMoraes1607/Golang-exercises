package main

import "fmt"

func main() {
	x, y := 0, 0
	fmt.Println("Insira a coordenada X:")
	fmt.Scan(&x)
	fmt.Println("Insira a coordenada Y:")
	fmt.Scan(&y)
	if x == 0 && y == 0 {
		fmt.Println("No ponto de origem.")
	} else if y == 0 {
		fmt.Println("O Ponto está sobre o eixo X.")
	} else if x == 0 {
		fmt.Println("O ponto está sobre o eixo Y.")
	} else if x > 0 && y > 0 {
		fmt.Println("O ponto está no 1º quadrante.")
	} else if x < 0 && y > 0 {
		fmt.Println("O ponto está no 2º quadrante.")
	} else if x < 0 && y < 0 {
		fmt.Println("O ponto está no 3º quadrante.")
	} else {
		fmt.Println("O ponto está no 4º quadrante.")
	}
}

/*Implemente um programa que solicita ao usuário as coordenadas X e Y (variáveis inteiras)
de um ponto e determina sua posição no plano de acordo com as opções:
1. Ponto na origem
2. Ponto sobre o eixo "X"
3. Ponto sobre o eixo "Y"
4. Ponto no 1º quadrante
5. Ponto no 2º quadrante
6. Ponto no 3º quadrante
7. Ponto no 4º quadrante
Assim que o programa encontrar uma das opções, ele não deve testar as outras. Por exemplo,
caso o ponto esteja no primeiro quadrante, ele não deve testar as opções seguintes. Teste
seu programa digitando valores de X e Y que cubram todas as opções acima (de 1 a 7). Tente
implementar seu programa utilizando até 6 comandos "if".*/
