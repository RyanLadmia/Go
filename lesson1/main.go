package main // Module principal de go

import "fmt" // fmt = format, Importation du package fmt pour les fonctions d'entrée/sortie (comme afficher le texte dans la console), ou formater du texte

func main() { // point d'entrée du programme, l'exécution commence ici
	fmt.Printf("Hello, World!\n") // Affiche le message "Hello, World!" dans la console
} // on peut utiliser fmt.Pintln, sans utiliser \n pour le même résultat

/*

A la racine du dossier du projet, exécuter la commande :
go mod init project-go (project-go est le nom du projet)

Ensuite, exécuter la commande :
go build ./main.go (build le projet en exécutable)

 Puis exécuter l'exécutable : ./main (main est le nom de l'exécutable)
go run main.go (main.go est le fichier principal du projet)

Cela fais apparaitre le message "Hello, World!" dans le terminal.

fmt.Print = affiche le texte tel quel sans retour à la ligne
fmt.Println = affiche le texte avec retour à la ligne
fmt.Printf = affiche le texte avec formatage (comme %s pour afficher un string, %d pour afficher un entier, %f pour afficher un float, %t pour afficher un booléen, etc.)

*/
