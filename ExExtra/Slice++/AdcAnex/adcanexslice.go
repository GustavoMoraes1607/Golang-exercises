package main

import "fmt"

func main() {
	onetofive := []int{1, 2, 3, 4, 5}
	tentofourteen := []int{10, 11, 12, 13, 14}
	fmt.Println("'onetofive' Antes da adição com append:", onetofive)
	fmt.Println("'tentofourteen' Antes da junção com append:", tentofourteen)
	onetofive = append(onetofive, 6, 7, 8, 9)
	fmt.Println("'onetofive' Depois da adição com append:", onetofive)
	onetofive = append(onetofive, tentofourteen...)
	fmt.Println("'onetofive' Depois da adição e junção com 'tentofourteen' usando append:", onetofive)
}
