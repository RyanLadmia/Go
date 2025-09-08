package main

import "fmt"

/*

Les fonctions anonymes : ressemble au javascript (fonctions flechées mais sans flèche ou paramètre)
	nameless fonctions/ anonymous fonctions

*/

func main() {

	w := func() {
		fmt.Println("Je suis une fonction anonyme sans return.") // On peut utiliser println sans importer fmt
	}
	w()

	fmt.Println("--------")

	z := func() string {
		println("Je suis une fonction anonyme avec return.")
		return "Ryan"
	}() // appel direct de la f°, obligatoire pour assigner la valeur
	println(z)

	fmt.Println("--------")

	x, y := func() (int, int) {
		println("Aucun paramètre. Deux return.")
		return 5, 7 // Valeurs assignées à x et y dans le scope de main
	}()
	println(x, y)

	fmt.Println("--------")

	func(a, b int) {
		println("A * A + B * B =", a*a+b*b)
	}(x, y)

}
