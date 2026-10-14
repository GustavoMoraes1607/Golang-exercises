package main

import "fmt"

func main() {
	switch wrktd := "Gustavo"; wrktd {
	default:
		fmt.Println("No one is working today")
	case "Gustavo":
		fmt.Println("Gustavo's working today")
	case "Carlos":
		fmt.Println("Carlos's working today")
		fallthrough
	case "Gabriela":
		fmt.Println("Whenever Carlos goes, Gabriela goes too.")
	}
}
