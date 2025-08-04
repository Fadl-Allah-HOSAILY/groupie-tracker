package functions

import (
    "encoding/json"
    "net/http"
    "fmt"
)

// Fonction pour récupérer et décoder du JSON
func FetchJSON(url string, target interface{}) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("erreur requête HTTP: %w", err)
    }
    defer resp.Body.Close()

    err = json.NewDecoder(resp.Body).Decode(target)
    if err != nil {
        return fmt.Errorf("erreur décodage JSON: %w", err)
    }

    return nil
}
