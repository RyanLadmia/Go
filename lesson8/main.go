package main

import (
	"errors"
	"fmt"
)

/*

Les fonctions :
	Des morceaux de code réutilisable

*/

func sayHello(name string) {
	fmt.Printf("Bonjour à tous! Je m'appelle %v\n", name)
	fmt.Println("---------")
}

// Version plus complexe : func calculatePerimetre (lo int, la int) Mais comme lo et la sont tous les deux des int on peut utiliser la version en-dessous
func calculatePerimetre(lo, la int) int { // (lié au return qui doit renvoyer un type int)
	return 2 * (lo + la)
}

// Renvoyer une erreur si l'utilisateur ne met pas le bon paramètre
func sayBye(name string) (string, error) { // on renvoi soit une string soit une erreur donc on met des parenthèses
	if name == "" {
		return "", errors.New("\berreur : vous n'avez pas mis de nom") // pas de majuscule ni de ponctuation à la fin, \b reviens en arriere pour supprimer l'espace
	}
	byeMessage := fmt.Sprintf("%v s'en va, Bonne soirée!\n", name) // Printf affiche directement le texte et Sprintf retourne le texte formaté
	return byeMessage, nil
}

func main() {

	sayHello("Ryan")

	rl := calculatePerimetre(5, 2) // On stocke la valeur de la fonction dans la variable pour l'appeller
	fmt.Printf("Le périmètre de mon rectangle est : %v.\n", rl)
	fmt.Println("---------")

	msg, err := sayBye("") // on stocke les valeurs dans deux variables pour eviter le <nil>
	if err != nil {        // signifie qu'il y a eu une erreur
		fmt.Println(err) // on affiche le message d'erreur
	} else {
		fmt.Print(msg) // on affiche le message bonne soirée
	}
}
