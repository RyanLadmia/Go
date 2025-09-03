package main

import "fmt"

/*
Les opérateurs :

	arithmétiques : +, -, *, /, % (modulo pour calculer le reste d'une division), ++, --
	relationnels : ==, !=, <, <=, >, >=
	logiques : && (et), || (ou), ! (non)

Les opérateurs arithmétiques sont utilisés pour effectuer des opérations mathématiques sur des nombres.
Les opérateurs relationnels sont utilisés pour comparer des valeurs et renvoyer un booléen (true ou false).
Les opérateurs logiques sont utilisés pour combiner des expressions booléennes.
*/

func main() {

	var (
		x int
		y int
	)

	x = 7
	y = 3

	// Arithmétiques
	fmt.Println(x + y) // addition
	fmt.Println(x - y) // soustraction
	fmt.Println(x / y) // division
	fmt.Println(x * y) // multiplication
	fmt.Println(x % y) // modulo (reste de la division)

	fmt.Println("\n------------------\n")

	// Relationnels
	fmt.Println(x == y) // égal à
	fmt.Println(x != y) // différent de
	fmt.Println(x < y)  // inférieur à
	fmt.Println(x <= y) // inférieur ou égal à
	fmt.Println(x > y)  // supérieur à
	fmt.Println(x >= y) // supérieur ou égal à

	fmt.Println("\n------------------\n")

	// Logiques
	fmt.Println(x > y && x != y) // les deux sont vrais = true
	fmt.Println(x < y || x == y) // un des deux ou les deux sont vrai = true, si les deux sont faux = false

}
