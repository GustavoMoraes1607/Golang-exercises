package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Dec\tBin\thex")
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}
