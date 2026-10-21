package main

import "fmt"

func main() {
	saudação, txt1, resultado, txt2, quantidade := soma("noite", 23, 44, 2, 7, 18)
	fmt.Println(saudação, "\n", "\t", txt1, resultado, txt2, quantidade)
}

func soma(x string, y ...int) (string, string, int, string, int) {
	saudação := ""
	if x == "manhã" {
		saudação = "Oi, bom dia!"
	} else if x == "tarde" {
		saudação = "Oi, boa tarde!"
	} else if x == "noite" {
		saudação = "Oi, boa noite!"
	} else {
		saudação = "Olá"
	}
	soma := 0
	for _, v := range y {
		soma += v
	}
	return saudação, "O resultado é:", soma, "quantidade de valores assossiados:", len(y)
}
