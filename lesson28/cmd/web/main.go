package main

/*

Envoyer des données dans les pages :
Utilisation du design pattern "MVC" : Model View Controller

*/

import (
	"fmt"
	"net/http"
	"project-go/lesson28/config"
	"project-go/lesson28/internal/handlers"
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
