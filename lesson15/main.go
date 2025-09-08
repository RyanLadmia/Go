package main

import "fmt"

/*

Création des types et découverte des struct :

Les structures permettent de générer des types en mixant int string bool(...) via des struct
Les valeurs ne sont pas assignés dans les struct. On les assigne dans les fonctions.

*/

type Example struct { // On crée la structure
	a string
	b int
	c bool
}

func main() {
	myExample := Example{ // On assigne les valeurs, on ne mets pas de valeur elles sont rempli par defaut 0 et false
		a: "Ryan",
		b: 28,
		c: true,
	}
	myExample2 := Example{"Youssef", 33, false} // Avec cette syntaxe, on doit assigner toutes les valeurs a toutes les clé sinon erreur. On privilegie la 1ere approche
	fmt.Println(myExample)
	fmt.Println(myExample2)

	fmt.Println("---------")

	myContact := newContact("Ryan", 28)

	fmt.Println(myContact) // go run main.go contact.go

}
