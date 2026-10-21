package main

import "fmt"

func main() {
	x := 231
	func(x int) {
		fmt.Println(x, "* 752 é:")
		fmt.Println(x * 752)
	}(x)
}
