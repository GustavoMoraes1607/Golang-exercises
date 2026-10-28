package main

import "fmt"

func txt1() {
	fmt.Println("Aqui estamos")
}
func txt2(f func()) {
	f()
	fmt.Println("E paramos aqui.")
}

func main() {
	txt2(txt1)
}

/*Callback: passe uma função como argumento a outra função.*/
