package main

/*

Structure des projets avec Go :

*/

import (
	"fmt"
	"net/http"
	"project-go/lesson25/internal/handlers"
)

const port = ":8081"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("(http://localhost:8081) - Server started on port", port)
	http.ListenAndServe(port, nil)
}

// Maintenant que les fichiers et les dossiers sont structurés, pour lancer l'application :
// 		go run ./cmd/web
