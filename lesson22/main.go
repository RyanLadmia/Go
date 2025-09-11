package main

import "fmt"

/*

Intro aux generics : version 1.18 nécessaire
Un generic est un mécanisme qui permet d’écrire du code réutilisable sans fixer à l’avance le type de données qu’il manipule.
	En gros, ça permet de créer une fonction, une structure, une classe ou une interface qui s’adapte à plusieurs types tout en gardant la sécurité de typage.
Cela permet de passer des types dans des paramètres de fonctions
Permettre à nos types d'être des paramètres
3 Principes :
	- Type parameters : Un type parameter est une variable de type qu’on utilise dans une fonction ou une structure.
	- Type inferences : permet de retirer les [types] dans le fmt.Prinln si un seul type est dans le fonction definie min
						Souvent, Go peut deviner le type tout seul → pas besoin de l’indiquer.
	- Type sets : 		Un type set est un ensemble de types qu’un paramètre peut accepter.
						C’est ce qu’on met dans les constraints.

*/

type Number interface { // Va avec le type set
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | ~float64
}

func main() {
	// Type parameters
	fmt.Println(min[float32](0.5, 0.7)) // On veut mettre des parametre autre que des float
	fmt.Println(min[int32](2, 1))       // On modifie la type en int32

	// Type inferences
	fmt.Println(mum(0.5, 0.7)) // La fonction va automatiquement savoir que les parametres sont des float64, pas besoin de mettre le type
	// Exemple 2:
	type f float64 // le type de la variable foo est f, avec un sous type float64
	var foo f = 3.14
	fmt.Println(mum(foo, 0.7))

	// Type set
	fmt.Println(max(0.7, foo))

}

// func min(x, y float32) float32 {     Fonction de base non générique sans type T
// 	return x
// }

// Tpe parameters
func min[T int32 | float32](x, y T) T { // On rend la fonction générique en implémentant un type T
	return x
}

// Type inference
func mum[T ~float64](x, y T) T { // Ici in accepte uniquement le type float64 direct (exemple 2, on doit rajouter le ~ car le type float64 est un sous type)
	return x
}

// Type any
// func max[T any](x, y T) T { // any = accepte n'importe quel type
// 	if x < y { // Il y a une erreur car il existe des type avec lesquels on ne peut pas faire cette opération (nbr complexes)
// 		return x
// 	}
// 	return y
// }

// Type set va être utilse pour l'exemple précédent : accepter tous les types de l'opération
func max[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}
