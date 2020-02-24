package datagen

import (
	"encoding/json"
	"net/http"
	"fmt"
)

const (
	geoURL = "http://geodb-free-service.wirefreethought.com/v1/"
)

var (
	totalCityCount int
)

type cityData struct {
	id          int
	cityType    string `json:"type"`
	wikiDataID  string `json:"wikiDataId"`
	city        string
	name        string
	country     string
	countryCode string
	region      string
	regionCode  string
	latitude    float64
	longitude   float64
}

type metadata struct {
	currentOffset int
	totalCount    int
}

type apiResponse struct {
	data []cityData
	metadata
}

func init() {
	cities := getGeoData(0)
	totalCityCount = cities.metadata.totalCount
}

func getGeoData(offset int) (cities *apiResponse) {
	resp, err := http.Get(fmt.Sprintf(
		"%sgeo/cities?hateoasMode=off&minPopulation=50000?offset=%v",
		geoURL,
		offset,
	))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	cities = new(apiResponse)
	err = json.NewDecoder(resp.Body).Decode(cities)
	if err != nil {
		panic(err)
	}
	return
}

func getTotalCityCount() int {
	return totalCityCount
}
