package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice":   27,
		"Gustavo": 25,
		"José":    36,
	}
	fmt.Println(ages)

	for key, value := range ages { //Range faz com que passe por todos os valores declarados
		fmt.Println(key, value)
	}
	delete(ages, "Alice") //Deletando um item do 'map'
	fmt.Println(ages)
}
