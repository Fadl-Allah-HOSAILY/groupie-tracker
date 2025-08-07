package functions

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var buff bytes.Buffer

func HandlerArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandlerError(w, "Error, method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.Split(r.URL.Path, "/artist/")[1]
	index, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if index < 1 || 52 < index {
		HandlerError(w, "Error, Bad Request", http.StatusBadRequest)
		return
	}

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

	err = tmpl.Execute(&buff, data)
	if err != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
