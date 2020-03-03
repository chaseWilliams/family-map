package external

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	geoURL  = "http://geodb-free-service.wirefreethought.com/v1/"
	nameURL = "https://uinames.com/api/?amount=500"
)

var (
	totalCityCount int
	nameData       []personData
)

type cityData struct {
	ID          int     `json:"id"`
	CityType    string  `json:"type"`
	WikiDataID  string  `json:"wikiDataId"`
	City        string  `json:"city"`
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionCode  string  `json:"regionCode"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type metadata struct {
	CurrentOffset int `json:"currentOffset"`
	TotalCount    int `json:"totalCount"`
}

type geoApiResponse struct {
	Data []cityData `json:"data"`
	metadata
}

type personData struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

func init() {
	cities := GetGeoData(0)
	totalCityCount = cities.metadata.TotalCount
	setNameData()
}

func GetGeoData(offset int) (cities *geoApiResponse) {
	resp, err := http.Get(fmt.Sprintf(
		"%sgeo/cities?hateoasMode=off&minPopulation=50000?offset=%v",
		geoURL,
		offset,
	))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	cities = new(geoApiResponse)
	err = json.NewDecoder(resp.Body).Decode(cities)
	if err != nil {
		panic(err)
	}
	return
}

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
