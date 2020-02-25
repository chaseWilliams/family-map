package datagen

import (
	"github.com/chaseWilliams/family-map/lib/models"
	"math/rand"
)


func randomPerson(year int) (person familyTreeNode) {
	first, last := randomName()
	return familyTreeNode {
		Person: models.Person {
			PersonID:  randomID(),
			FirstName: first,
			LastName:  last,
			Gender:    randomGender(),
		},
		straight:  randomIsStraight(),
		events:    make([]models.Event, 3),
		deathYear: -1,
		birthYear: year,
	}
}

func randomID() string {
	lower := 48
	upper := 90
	bytes := make([]byte, 8)
	for i := range bytes {
		bytes[i] = byte(lower + rand.Intn(upper - lower + 1))
	}
	return string(bytes)
}

func randomName() (string, string) {
	index := rand.Intn(500)
	return nameData[index].name, nameData[index].surname
}

func randomGender() string {
	if rand.Intn(2) == 0 {
		return "f"
	}
	return "m"
}

func randomIsStraight() bool {
	if rand.Float64() <= 0.04 {
		return false
	}
	return true
}