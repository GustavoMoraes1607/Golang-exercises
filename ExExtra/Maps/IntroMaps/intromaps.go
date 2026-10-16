package main

import "fmt"

func main() {
	contactlist := map[string]int{
		"Gustavo": 949491352,
		"Ariel":   673935856,
		"Iagor":   934567245,
		"João":    354750356,
	}
	fmt.Println("Esta é a lista completa de contatos:", contactlist)
	fmt.Println("Este é o número de telefone do Gustavo:", contactlist["Gustavo"])

	contactlist["gopher"] = 456727895
	fmt.Println("Esta é a lista de contatos após adicionar o 'Gopher':", contactlist)

	if contact, ok := contactlist["Ghost"]; !ok { // ", ok" vai me dizer 'true' caso encontre esse valor na lista
		fmt.Println("Este contato não existe.") // e o "!ok" vai me dizer 'false' caso não encontre. (comma ok idiom).
	} else {
		fmt.Println(contact)
	}
}
