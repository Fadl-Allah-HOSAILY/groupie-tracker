package functions

type Artist struct {
    ID           int      `json:"id"`
    Image        string   `json:"image"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
}

type LocationData struct {
    ID        int      `json:"id"`
    Locations []string `json:"locations"`
}

type LocationsResponse struct {
    Index []LocationData `json:"index"`
}


type DateData struct {
    ID    int      `json:"id"`
    Dates []string `json:"dates"`
}
type DatesResponse struct {
    Index []DateData `json:"index"`
}

type RelationData struct {
    ID             int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}
type RelationResponse struct {
    Index []RelationData `json:"index"`
}

type AllData struct {
    Artists   []Artist
    Locations LocationsResponse
    Dates     DatesResponse
    Relation  RelationResponse
}