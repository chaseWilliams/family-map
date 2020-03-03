package external

import (
	"math"
	"math/rand"
	"sort"
)

/*
RandomCloseCity will return a random city from the database,
weighted towards those closer to the provided GPS location
*/
func RandomCloseCity(lat, long float64) City {
	cities, err := getCities()
	if err != nil {
		panic(err)
	}
	sort.Slice(cities, func(i, j int) bool {
		return haversine(lat, long, cities[i].Latitude, cities[i].Longitude) <
			haversine(lat, long, cities[j].Latitude, cities[j].Longitude)
	})
	city := cities[0]
	for _, c := range cities {
		if rand.Float64() < 0.7 {
			city = c
			break
		}
	}
	return city
}

func haversine(lat1, long1, lat2, long2 float64) float64 {
	/*
		http://www.movable-type.co.uk/scripts/latlong.html
	*/
	R := 6371e3
	lat1 = lat1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180
	long1 = long1 * math.Pi / 180
	long2 = long2 * math.Pi / 180
	deltaLat := lat2 - lat1
	deltaLong := long2 - long1

	a := math.Pow(math.Sin(deltaLat/2), 2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(deltaLong/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
