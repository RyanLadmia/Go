package main

import (
	"fmt"
	"project-go/lesson18/footypes"
)

/*

Les modules :
En Go, un module est un ensemble de packages versionné, qui sert à organiser et gérer le code ainsi que ses dépendances.

1.	Module = projet ou bibliothèque Go
	•	Contient un fichier go.mod à sa racine.
	•	Un module peut contenir plusieurs packages.
2.	Fichier go.mod
	•	Définit le nom du module et les versions des dépendances utilisées.
3.	But principal
	•	Gérer les packages et leurs versions.
	•	Faciliter l’importation du code entre différents projets.

Un module Go est un projet avec un go.mod qui contient un ou plusieurs packages et qui gère les dépendances de manière versionnée.


Pour lancer la création d'un module la commande est:
	go mod init github.com/ryanladmia/monmodule
Ici github.com est l'endroit où on stocke le module, avec le nom utilisateur suivi du nom du module.
Cela va créer un fichier go.mod avec les infos du module.

*/

func main() {
	fmt.Println("test")

	var fooVar footypes.Foo
	fooVar.TypeInt = 18
	fooVar.TypeString = "Ryan"
	fooVar.TypeBoolean = true

	fmt.Println(fooVar)
	fmt.Println(fooVar.TypeInt)
}
