package main

/*

Templates HTLM

*/

import (
	"fmt"
	"net/http"
)

const port = ":8081"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("(http://localhost:8081) - Server started on port", port)
	http.ListenAndServe(port, nil)
}

// Pour lancer tous les fichier dans la racine ici (main.go et handlers.go):
// 		go run .
