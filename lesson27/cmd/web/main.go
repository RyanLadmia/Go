package main

/*

Variables globales :
On créer des variables globales dans config.go pour les utiliser dans toute l'application (ici handlers.go et main.go)
(Port et TemplateCache)

*/

import (
	"fmt"
	"net/http"
	"project-go/lesson27/config"
	"project-go/lesson27/internal/handlers"
)

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()
	if err != nil {
		panic(err)
	}

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8081"

	handlers.CreateTemplate(&appConfig)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("(http://localhost:8081) - Server started on port", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}

// Maintenant que les fichiers et les dossiers sont structurés, pour lancer l'application :
// 		go run ./cmd/web
