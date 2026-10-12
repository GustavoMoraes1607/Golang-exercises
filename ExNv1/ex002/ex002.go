package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x := 10
	y := "Isto é uma String"
	z := 10 > 3
	fmt.Println(x, y, z)
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}
