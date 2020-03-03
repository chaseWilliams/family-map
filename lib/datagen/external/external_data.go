package external

import (
	"encoding/json"
	"net/http"
)

const (
	nameURL = "https://uinames.com/api/?amount=500&region=United%20States"
)

var (
	totalCityCount int
	nameData       []personData
)



type personData struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

/*
func init() {
	cities := GetGeoData(0)
	totalCityCount = cities.metadata.TotalCount
	setNameData()
}
*/


func setNameData() {
	nameData = make([]personData, 500)
	resp, err := http.Get(nameURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&nameData)
	if err != nil {
		panic(err)
	}
}

/*
GetNameData returns a slice of personData from
uinames.com/api
*/
func GetNameData() []personData {
	return nameData
}

/*
GetTotalCities returns the total number of cities that can be indexed
*/
func GetTotalCities() int {
	return totalCityCount
}
