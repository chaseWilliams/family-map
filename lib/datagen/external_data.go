package datagen

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

type geoApiResponse struct {
	data []cityData
	metadata
}

type personData struct {
	name    string
	surname string
	gender  string
	region  string
}

func init() {
	cities := getGeoData(0)
	totalCityCount = cities.metadata.totalCount
	getNameData()
}

func getGeoData(offset int) (cities *geoApiResponse) {
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

func getNameData() {
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
