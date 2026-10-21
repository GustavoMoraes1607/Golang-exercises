package main

import "fmt"

func main() {
	defer fmt.Println("Com defer, veio primeiro.")
	fmt.Println("Sem defer, veio depois.")
}
