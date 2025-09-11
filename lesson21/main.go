package main

import (
	"errors"
	"fmt"
)

/*

Tester les codes et fonctions :
Méthode n°1 : écrire la fonction avant le test
Méthode n°2 : écrire le test avant la fonction

*/

func main() {
	result, err := divide(10.0, 2.5)
	if err != nil { // Nil = pas d'erreur
		panic(err) // permet de stopper l'execution du code
	}
	fmt.Println(result)
}

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("impossible de diviser par zéro")
	}
	return a / b, nil
}
