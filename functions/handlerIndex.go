package functions

import (
	"html/template"
	"net/http"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandlerError(w, "Error, Page Not Found", http.StatusNotFound)
		return
	}
	AllArtists, err := LoadAllData()
	if err != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, AllArtists)
	if err != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return
	}
}
