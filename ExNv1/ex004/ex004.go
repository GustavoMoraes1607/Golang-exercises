package main

import "fmt"

type goplayground int

var x goplayground

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
