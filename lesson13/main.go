package main

import "fmt"

/*

Pass By Value :
On fait une copie en mémoire de la valeur du paramètre qui est transmis.
1°) on stocke number = 10
2°) Dans la fonction update A, number = 5 est copier dans la mémoire de la valeur de parametre sans supprimer la 1ere

Les données sont donc stockées dans différents endroits dans la mémoire.
En pass by value, une copie de la variable est stockée en mémoire, alors qu’un pointer permet à la fonction de modifier directement l’original en mémoire.

Exemple d'utilisation dans un e-commerce :
- Quand tu veux juste calculer ou simuler quelque chose (remise, taxe, prix final…) → pass by value est parfait.
- Quand tu veux modifier le panier, stock ou données partagées → tu utilises un pointer.

*/

func updateA(number int) int {
	number = 5
	return number
}

func updateB(item map[string]int) {
	item["bonbon"] = 4
	item["PS5"] = 500

}

func main() {
	// Pour les types string int bool float array
	number := 10
	number = updateA(number) // Pour copie la valeur de number dans la memoire en la reassignant en appelant updateA
	fmt.Println(number)

	// Pour les maps, fonctions
	groceryList := map[string]int{
		"prince": 5,
		"viande": 10,
	}
	updateB(groceryList) // On copie la valeur de map en ajoutant un item
	fmt.Println(groceryList)
}
