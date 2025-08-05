package functions

import (
	"html/template"
	"net/http"
	"strings"
)

func HandlerArtist(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/artist/")[1]
	data, err := DetailsData(id, w)
	if err != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		HandlerError(w, "Erreur d'affichage du template", http.StatusInternalServerError)
		return
	}
}
