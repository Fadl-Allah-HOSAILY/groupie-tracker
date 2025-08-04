package functions

import (
	"html/template"
	"net/http"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	data, err := LoadAllData()
	if err != nil {
		http.Error(w, "Erreur lors du chargement des données", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erreur d'affichage du template", http.StatusInternalServerError)
		return
	}
}
