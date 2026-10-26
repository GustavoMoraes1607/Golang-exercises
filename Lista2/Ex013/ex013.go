package main

import "fmt"

func main() {
	nota1, nota2, nota3 := 0, 0, 0
	media := 60
	faltas := 0
	lmtfaltas := 18
	fmt.Println("Insira a nota do primeiro período:")
	fmt.Scan(&nota1)
	fmt.Println("Insira a nota do segundo período:")
	fmt.Scan(&nota2)
	fmt.Println("Insira a nota do terceiro período:")
	fmt.Scan(&nota3)
	fmt.Println("Insira quantidade de faltas:")
	fmt.Scan(&faltas)
	if nota1+nota2+nota3 >= media && faltas <= lmtfaltas {
		fmt.Println("APROVADO")
	} else if nota1+nota2+nota3 < media {
		fmt.Println("REPROVADO POR MÉDIA")
	} else {
		fmt.Println("REPROVADO POR FALTA")
	}
}

/*Faça um programa que leia as três notas e as faltas de um aluno e imprima
sua situação. ("APROVADO","REPROVADO POR FALTA" ou "REPROVADO POR MÉDIA").
Observação - A média para aprovação é 60 e o limite de faltas é 18.*/
