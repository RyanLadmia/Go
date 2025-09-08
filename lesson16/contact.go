package main

import "fmt"

type contact struct { // déclaration d'un type struct nommé contact avec 3 champs
	name    string
	age     int
	numbers map[string]string
}

func newContact(name string, age int, numbers map[string]string) contact { // retourne une valeur de type contact
	c := contact{ // on stocke le type dans une variale
		name:    name,
		age:     age,
		numbers: numbers,
	}
	return c
}

func (c contact) displayContact() string { // Fonction qui retourne une string
	display := fmt.Sprintf("Contact: %v, (%v ans)\n", c.name, c.age)
	display += "---------------\n" // concatène des ---- pour la lisibilité

	for key, value := range c.numbers { // on parcours la map c.number avec range
		display += fmt.Sprintf("%v, %v\n", key+":", value)
	}

	return display
}
