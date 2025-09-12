package handlers // Le nom du package dépend du dossier

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"project-go/lesson28/config"
	"project-go/lesson28/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	names := make(map[string]string)
	names["owner"] = "Ryan"

	renderTemplate(w, "home", &models.TemplateData{
		StringData: names,
	})
}
func About(w http.ResponseWriter, r *http.Request) {
	names := make(map[string]string)
	names["owner"] = "Ryan"
	age := make(map[string]int)
	age["owner"] = 28

	renderTemplate(w, "about", &models.TemplateData{
		StringData: names,
		IntData:    age,
	})
}

var appConfig *config.Config

func CreateTemplate(app *config.Config) {
	appConfig = app
}

func renderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
	templateCache := appConfig.TemplateCache

	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		http.Error(w, "Le template n'existe pas.", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, td)
	buffer.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) { // Pour utiliser les templates
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}

		cache[name] = tmpl
	}

	return cache, nil
}
