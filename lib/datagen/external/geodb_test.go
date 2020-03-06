package external

import (
	"github.com/chaseWilliams/family-map/lib/database"
	"testing"
)

func BenchmarkGetCities(b *testing.B) {
	database.StartTestingSession(b)
	for i := 0; i < b.N; i++ {
		_, err := getCities()
		if err != nil {
			b.Error(err)
		}
	}
}
