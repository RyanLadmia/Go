package main

import "fmt"

/*

Les maps : Ce sont des structure de donnée associant des clé et des valeurs.
Similaire en objets en JS ou des dictionnaires en Python.
On va sauvegarder une clé et une valeur associée
Les clés sont uniques mais les valeur peuvent se répéter

*/

func main() {

	superMarketPrice := map[string]int{
		"prince": 8,
		"eau":    2,
		"viande": 6,
	}

	// Pour afficher la valeur de la clé "prince"
	fmt.Println(superMarketPrice["prince"])

	fmt.Println("--------")

	// Pour afficher différentes clés et valeurs, l'ordre d'affichage des key/values est aléatoire, parfois prince en 1er parfois en dernier ...
	for key, value := range superMarketPrice {
		fmt.Println(key, value)
	}

	fmt.Println("--------")

	daysInAYear := 0

	daysInAMonth := map[string]int{
		"janvier":   31,
		"fevrier":   28,
		"mars":      31,
		"avril":     30,
		"mai":       31,
		"juin":      30,
		"juillet":   31,
		"aout":      31,
		"septembre": 30,
		"octobre":   31,
		"novembre":  30,
		"decembre":  31,
	}

	// Combien de jour dans une année
	for _, value := range daysInAMonth {
		daysInAYear = daysInAYear + value
	}
	fmt.Printf("Nombre de jours dans une année: %d jours\n", daysInAYear)

	fmt.Println("--------")

	// Combien de jour dans un mois
	for key, value := range daysInAMonth {
		fmt.Printf("Le mois de %v possède %d jours.\n", key, value) // Les mois apparaissent dans un ordre aléatoire : ex : avril, juin, decembre...
	}

	fmt.Println("--------")

	// Pour les faire apparaitre dans le bon ordre, il faut utiliser un slice :
	// SLICE : un slice est une séquence dynamique d’éléments d’un même type.
	// On peut le voir comme une vue flexible sur un tableau sous-jacent :
	// - Contrairement aux tableaux (array), la taille d’un slice n’est pas fixe.
	// - On peut ajouter, retirer, ou redimensionner facilement un slice.

	// exemple de slice :
	monthOrder := []string{
		"janvier",
		"fevrier",
		"mars",
		"avril",
		"mai",
		"juin",
		"juillet",
		"aout",
		"septembre",
		"octobre",
		"novembre",
		"decembre",
	}

	for _, month := range monthOrder { // Le _ sert à ignoré l'index (0, 1, 2...), sinon mettre i à la place (convention)
		fmt.Printf("Le mois de %v possède %v jours.\n", month, daysInAMonth[month])
	}
}
