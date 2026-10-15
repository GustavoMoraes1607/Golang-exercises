package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array:", array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", slice)

	slice2 := append(slice, 6)
	fmt.Println("slice2:", slice2)

	fmt.Println("slice[3]:", slice[3])
	slice[3] = 282116
	fmt.Println("slice[3] após alteração de valor:", slice[3])
	fmt.Println("Representação slice com novo valor:", slice)
}
