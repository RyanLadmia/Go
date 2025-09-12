package main

/*

Création d'une application web
	Lancement d'un serveur en définissant le port et le message dans le terminal

*/

import (
	"fmt"
	"net/http"
)

const port = ":8081"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenue dans ma page GO !") // Fprint permet d'écrire dans un writter
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Conacter moi à : GO GO GO!")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)

	fmt.Println("(http://localhost:8081) - Server started on port", port) // A afficher dans le terminal lors du lancement du serveur.
	http.ListenAndServe(port, nil)
}

// Lancer le serveur avec :
//	go run main.go

// Arrêter le processus avec :
// 	control + C
