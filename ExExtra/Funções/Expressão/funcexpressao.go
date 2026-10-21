package main

import "fmt"

func main() {
	x := 231
	y := func(x int) int {
		return x * 752
	}
	fmt.Println(x, "* 752 é:", y(x))
}
