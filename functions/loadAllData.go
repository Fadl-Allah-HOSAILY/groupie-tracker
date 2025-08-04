package functions


import gb "groupieTracker/global"

func LoadAllData() ([]gb.Artist, error) {

	var artists []gb.Artist

	err := FetchJSON("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
