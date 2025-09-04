package main

import "fmt"

func main() {

	// Structure de la boucle : avec for ou while
	// 1°) Initialisation de la variable, 2°) Condition d'arrêt de la boucle, 3°) Incrémentation ou décrémentation de la variable

	// Exemple de boucle :
	for i := 0; /*1°)*/ i < 3; /*2°) */ i++ /*3°) */ { // version moderne for i := range 3
		fmt.Println(i) // Affiche 0, 1, 2
	}

	fmt.Println("-------------")

	//ou de maniere simplifiée :
	x := 0

	for x < 5 {
		fmt.Println(x) //Affiche 0, 1, 2, 3, 4
		x++
	}

	fmt.Println("-------------")

	// ou avec break et continue
	y := 0
	for {
		if y > 5 {
			break
		}
		fmt.Println(y)
		y++
	}

	fmt.Println("-------------")

	z := 0
	for ; z <= 10; z++ { // le ; permet d'éviter une erreur, car la variable est déclarée à l'extérieur
		if z%2 == 1 { // affiche les nombres paires, pour les nombres impaire mettre 0
			continue
		}
		fmt.Println(z)
	}

}
