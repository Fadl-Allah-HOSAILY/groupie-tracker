package functions

import (
	"net/http"
	"os"
	"strings"
)

// handler static folder
func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandlerError(w, "Method not allowed !", http.StatusMethodNotAllowed)
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/static") {
		HandlerError(w, "Statut not found !", http.StatusNotFound)
		return
	} else {
		infos, err := os.Stat(r.URL.Path[1:])
		if err != nil {
			HandlerError(w, "Statut not found !", http.StatusNotFound)
			return
		} else if infos.IsDir() {
			HandlerError(w, "Access Forbidden !", http.StatusForbidden)
			return
		} else {
			http.ServeFile(w, r, r.URL.Path[1:])
		}
	}
}