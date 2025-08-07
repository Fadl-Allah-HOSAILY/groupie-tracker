package functions

import (
	"bytes"
	"html/template"
	"net/http"
)

var buf bytes.Buffer

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandlerError(w, "Error, Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		HandlerError(w, "Error, method not allowed", http.StatusMethodNotAllowed)
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

	err = tmpl.Execute(&buf, AllArtists)
	if err != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, AllArtists)
}
