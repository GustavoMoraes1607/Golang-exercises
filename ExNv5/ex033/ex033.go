package main

import "fmt"

func main() {
	contato := struct {
		nome    string
		idade   int
		corfav  []string
		idiomas map[string]string
	}{
		nome:   "Gustavo",
		idade:  25,
		corfav: []string{"Preto", "Roxo"},
		idiomas: map[string]string{
			"Nativo":  "pt/br",
			"Segundo": "eng",
		},
	}
	fmt.Println(contato)
}

//Crie e use um struct anônimo.
//Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
