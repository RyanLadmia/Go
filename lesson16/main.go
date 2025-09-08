package main

import "fmt"

/*

Receiver function :
pour associer une fonction a un type spécifique

*/

type Example struct { // On crée la structure
	a string
	b int
	c bool
}

func main() {
	myContact := newContact("Ryan", 28, map[string]string{"Fixe": "04 91 12 34 56", "Portable": "06 06 94 38 27"})

	//fmt.Println(myContact) // go run main.go contact.go
	fmt.Println(myContact.displayContact())

	mynewContact := newContact("Youssef", 33, map[string]string{"Fixe": "04 92 12 34 00", "Portable": "07 00 94 38 00"})

	//fmt.Println(mynewContact) // go run main.go contact.go
	fmt.Println(mynewContact.displayContact())

	my2newContact := newContact("Konstantin", 29, map[string]string{"Fixe": "04 04 04 04 04", "Portable": "07 07 07 07 07"})
	fmt.Println(my2newContact.displayContact())
}
