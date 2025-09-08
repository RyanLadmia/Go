package main

import "fmt"

/*

Les pointers :
Un pointer est une variable qui contient l’adresse mémoire d’une autre variable au lieu de sa valeur directe.
C'est une variable spéciale qui ne contient pas directement une valeur, mais l’adresse d’une autre variable, ce qui permet d’accéder ou de modifier cette variable originale via cette adresse.
	•	Avec un pointer, tu peux accéder à la valeur originale et la modifier depuis une fonction ou un autre endroit du programme.
	•	On déclare un pointer avec * et on obtient l’adresse d’une variable avec &.

	La différence avec Pass By Value est que pointer modifie la valeur originale.

*/

func update(number *int, value int) { // on peut aussi ajouter value int
	*number = value // et remplacer 5 par value pour la changer lors de l'appel de la fonction
}

func main() {
	number := 10
	fmt.Println(number)                               // = 10
	fmt.Println("Adresse mémoire de number", &number) // = adresse mémoire

	// Stocker l'adresse mémoire de number dans une variable (on crée un pointer)
	myPointer := &number
	fmt.Println(myPointer)
	fmt.Printf("La valeur de l'adresse mémoire %v est %d.\n", myPointer, *myPointer)

	update(myPointer, 100) // on attribue 100 à la value de la fonction update
	fmt.Printf("La valeur de l'adresse mémoire %v à changé -> %d.\n", myPointer, number)

}
