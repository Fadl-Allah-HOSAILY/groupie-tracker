package functions

import (
	"net/http"

	gb "groupieTracker/global"
)

func DetailsData(id string, w http.ResponseWriter) (gb.AllData, error) {
	var artist gb.Artist
	var location gb.Location
	var date gb.Date
	var relation gb.Relations
	var detail gb.AllData

	errArtist := FetchJSON("https://groupietrackers.herokuapp.com/api/artists/"+id, &artist)
	if errArtist != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return detail, errArtist
	}
	errLocation := FetchJSON("https://groupietrackers.herokuapp.com/api/locations/"+id, &location)
	if errLocation != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return detail, errLocation
	}
	errDate := FetchJSON("https://groupietrackers.herokuapp.com/api/dates/"+id, &date)
	if errDate != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return detail, errDate
	}
	errRelation := FetchJSON("https://groupietrackers.herokuapp.com/api/relation/"+id, &relation)
	if errRelation != nil {
		HandlerError(w, "Error, internal server error", http.StatusInternalServerError)
		return detail, errLocation
	}

	detail = gb.AllData{
		Id:             artist.Id,
		Image:          artist.Image,
		Name:           artist.Name,
		Members:        artist.Members,
		CreationDate:   artist.CreationDate,
		FirstAlbum:     artist.FirstAlbum,
		Locations:      location,
		Dates:          date,
		DatesLocations: relation,
	}
	return detail, nil
}
