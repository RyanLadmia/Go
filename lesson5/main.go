package main

import (
	"fmt"
	"math/rand"
)

/*

Les conditions : if , else et else if

	if condition (=true)
	alors le else est executé (mais il n'est pas obligatoire)
	et else if (condition intermédiaire)

*/

func main() {

	fmt.Println(rand.Int()) // Affiche un nombre aléatoire entre 0 et 2 147 483 647 (2^31 - 1)

	if x := rand.Int(); x%2 == 0 { // Si le reste de la division par 2 est égal à 0, avec x généré aléatoirement
		fmt.Println(x, "est un nombre pair")
	} else {
		fmt.Println(x, "est un nombre impair")
	}

	// Ou on peut utiliser une fonction qui existe déjà
	y := rand.Int() % 2 // resultat soit 1 soit 0

	if y == 0 {
		fmt.Println(y, " est un nombre pair")
	} else {
		fmt.Println(y, " est un nombre impair")
	}

	age := 17

	if age > 18 {
		fmt.Println("Je suis majeur")
	} else if age == 18 {
		fmt.Println("Je viens d'être majeur")
	} else {
		fmt.Println("Je suis mineur")
	}
}
