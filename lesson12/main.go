package main

import "fmt"

/*

La fonction fmt.Print vs print :
	print est amené à disparaître dans le futur
	il est recommandé d'utiliser fmt.Print

*/

func main() {

	fmt.Printf("Salut à tous!\n")

	x, y := 10, 11

	fmt.Printf("Mon nombre (défaut) est : %v\n", x)  // écriture du nombre par défaut
	fmt.Printf("Mon nombre (base 2)  est : %b\n", x) // écriture binaire
	fmt.Printf("Mon nombre (base 8) est : %O\n", x)  // écriture octale, ici la majuscule renvoi le prefixe (la lettre)
	fmt.Printf("Mon nombre (base 10) est : %d\n", x) // ecriture décimale par défqaut
	fmt.Printf("Mon nombre (base 16) est : %X\n", x) // écriture hexadécimal  ici la majuscule

	fmt.Println("---------")

	fmt.Printf("La valeur de x est égale à la valeure de y -> %t\n", x > y) // ici %t pour un booléen

	// Il existe d'autre % pour différents types de formatage.

}
