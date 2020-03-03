package external

import (
	"encoding/json"
	"fmt"
	"github.com/chaseWilliams/family-map/lib/database"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	geoURL string = "http://geodb-free-service.wirefreethought.com/v1/geo/cities?hateoasMode=off&countryIds=US&sort=-population&limit=10&types=CITY&offset="
)

type City struct {
	ID          int     `json:"id" db:"id"`
	CityType    string  `json:"type" db:"city_type"`
	WikiDataID  string  `json:"wikiDataId" db:"wiki_data_id"`
	City        string  `json:"city" db:"city"`
	Name        string  `json:"name" db:"name"`
	Country     string  `json:"country" db:"country"`
	CountryCode string  `json:"countryCode" db:"country_code"`
	Region      string  `json:"region" db:"region"`
	RegionCode  string  `json:"regionCode" db:"region_code"`
	Latitude    float64 `json:"latitude" db:"latitude"`
	Longitude   float64 `json:"longitude" db:"longitude"`
}

type metadata struct {
	CurrentOffset int `json:"currentOffset"`
	TotalCount    int `json:"totalCount"`
}

type geoApiResponse struct {
	Data []City `json:"data"`
	metadata
}

func SetupCityData() {
	start := time.Now()
	cities := make([]City, 0)
	err := database.StartSession()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 200; i += 10 {
		resp, err := getGeoData(i)
		if err != nil {
			panic(err)
		}
		cities = append(cities, resp.Data...)
	}
	for _, city := range cities {
		err := city.Save()
		if err != nil {
			panic(err)
		}
	}
	database.GetTransaction().Commit()
	fmt.Printf("took %s seconds\n", time.Now().Sub(start))
}

func getGeoData(offset int) (cities *geoApiResponse, err error) {
	resp, err := http.Get(geoURL + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("got status code %d, %s", resp.StatusCode, b)
	}

	cities = new(geoApiResponse)
	err = json.NewDecoder(resp.Body).Decode(cities)
	if err != nil {
		return nil, err
	}
	return
}

/*
Save will take this City model and create it in the database
*/
func (c City) Save() (err error) {
	tx := database.GetTransaction()
	_, err = tx.Exec(
		`INSERT INTO Cities
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		c.ID,
		c.CityType,
		c.WikiDataID,
		c.City,
		c.Name,
		c.Country,
		c.CountryCode,
		c.Region,
		c.RegionCode,
		c.Latitude,
		c.Longitude,
	)
	return
}

func getCities() (cities []City, err error) {
	tx := database.GetTransaction()
	cities = []City{}
	err = tx.Select(&cities, "SELECT * FROM Cities")
	return
}
