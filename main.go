package main

import (
	"fmt"
	"log"

	"groupieTracker/functions"
)

func main() {
	var artists []functions.Artist
	err := functions.FetchJSON("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		log.Fatal("Erreur artists:", err)
	}

	var locations functions.LocationsResponse
	err = functions.FetchJSON("https://groupietrackers.herokuapp.com/api/locations", &locations)
	if err != nil {
		log.Fatal("Erreur locations:", err)
	}

	var dates functions.DatesResponse
	err = functions.FetchJSON("https://groupietrackers.herokuapp.com/api/dates", &dates)
	if err != nil {
		log.Fatal("Erreur dates:", err)
	}

	var relation functions.RelationResponse
	err = functions.FetchJSON("https://groupietrackers.herokuapp.com/api/relation", &relation)
	if err != nil {
		log.Fatal("Erreur relation:", err)
	}

	fmt.Println("Nom du 1er artiste :", artists[0].Name)
	fmt.Println("Lieu du 1er artiste :", locations.Index[0].Locations)
	fmt.Println("Dates du 1er artiste :", dates.Index[0].Dates)
	fmt.Println("Dates+lieux relation :", relation.Index[0].DatesLocations)
}
