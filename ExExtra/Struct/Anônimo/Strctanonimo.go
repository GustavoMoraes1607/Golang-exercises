package main

import "fmt"

func main() {
	x := struct {
		Nome  string
		Idade int
	}{
		Nome:  "Mario",
		Idade: 37,
	}
	fmt.Println(x)
}
