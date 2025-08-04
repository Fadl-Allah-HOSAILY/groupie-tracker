package functions

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	gb "groupieTracker/global"
)

func HandlerArtist(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/artist/")[1]
	_, err := strconv.Atoi(id)
	if err != nil {
		HandlerError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	fmt.Println(gb.AllArtists)

	// tmpl, err := template.ParseFiles("templates/artist.html")
	// if err != nil {
	// 	HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
	// 	return
	// }

	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	HandlerError(w, "Erreur d'affichage du template", http.StatusInternalServerError)
	// 	return
	// }
}
