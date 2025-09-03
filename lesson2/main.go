package main

import "fmt"

func main() {
	var (
		name  string
		age   int
		numb  int
		state bool
		/*
		   Les variables peuvent être déclarées avec le mot clé var, suivi du nom de la variable, du type et éventuellement d'une valeur initiale.
		   Les variables peuvent aussi être déclarées avec l'opérateur :=, qui infère le type de la variable à partir de la valeur initiale.//
		*/
	)

	name = "Ryan"
	age = 28
	numb = 1
	state = true

	fmt.Printf("Je suis %v!\n", name)
	fmt.Printf("J'ai %v ans!\n", age)
	fmt.Printf("Je suis le n° %v!\n", numb)
	fmt.Printf("And it's %v!\n", state)
	fmt.Printf("Bonjour! Je m'appele %v, j'ai %v ans, je suis le n° %v! And it's %v!\n", name, age, numb, state)
}
