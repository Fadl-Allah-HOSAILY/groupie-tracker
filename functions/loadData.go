package functions

func LoadAllData() (AllData, error) {
    var data AllData

    err := FetchJSON("https://groupietrackers.herokuapp.com/api/artists", &data.Artists)
    if err != nil {
        return data, err
    }

    err = FetchJSON("https://groupietrackers.herokuapp.com/api/locations", &data.Locations)
    if err != nil {
        return data, err
    }

    err = FetchJSON("https://groupietrackers.herokuapp.com/api/dates", &data.Dates)
    if err != nil {
        return data, err
    }

    err = FetchJSON("https://groupietrackers.herokuapp.com/api/relation", &data.Relation)
    if err != nil {
        return data, err
    }

    return data, nil
}
