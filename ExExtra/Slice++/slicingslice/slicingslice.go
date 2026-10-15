package main

import "fmt"

func main() {
	//                  0.			 1.			  2.			3.		   4.
	sabores := []string{"Mozzarela", "Pepperoni", "Margherita", "Diavola", "4Formaggi"}
	fatia := sabores[:] // [:] engloba todos os elementos do slice
	fmt.Println("Estes são todos os sabores no menu:", fatia)
	for x := 0; x < len(sabores); x++ { //len(sabores) quer dizer que o loop vai percorrer todos os elementos do slice
		fmt.Println(sabores[x])
	}
	sabores = append(sabores[:2], sabores[4:]...) //A função append() é usada para adicionar elementos a um slice.
	fmt.Println("Estes são os sabores no menu depois do corte:", sabores)
	for x := 0; x < len(sabores); x++ {
		fmt.Println(sabores[x])
	}
}
