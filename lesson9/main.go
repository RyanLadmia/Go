package main

import "fmt"

/*

Le scope et création d'un deuxième fichier

Qu'est-ce qu'un Scope :
endoit du code ou la variable, la fonction va être reconnu, définit laporté des variables
- global (déclaré à l'extérieur de toute fonction comme main et des autres)
- local (déclaré dans la fontion ou le bloc)

*/

var x int // variable globale, si x n'est pas globale elle n'est pas prise en compte dans hello.go

func main() {
	x = 5 // assigne une valeur à la variable globale (modifie une variable existante)
	fmt.Println(x)
	f()
	showY()
}

func f() {
	x := 10 // assigne une valeur à la variable local ("créer" une nouvelle variable x locale [portée dans la fonction f()])
	fmt.Println(x)
}

// Pour build les deux fichier dans un même fichier exe :
// go build main.go hello.go

// Pour afficher les deux fichiers faire :
// go run main.go hello.go
