package family

import (
	"github.com/chaseWilliams/family-map/lib/datagen/external"
	"github.com/chaseWilliams/family-map/lib/models"
	"math/rand"
)

/*
RandomPerson returns a pointer to newly constructed random family.Person
who is born in the provided year
*/
func RandomPerson(year int) (p *Person) {
	first, last := randomName()
	p = &Person{
		model: models.Person{
			PersonID:  randomID(),
			FirstName: first,
			LastName:  last,
			Gender:    randomGender(),
		},
		straight:      randomIsStraight(),
		events:        make([]models.Event, 0),
		spouses:       make([]*Person, 0),
		children:      make([]*Person, 0),
		marriageYears: make([]int, 0),
		divorceYears:  make([]int, 0),
		deathYear:     -1,
		birthYear:     year,
		married:       false,
	}

	return
}

func randomID() string {
	lower := 48
	upper := 90
	bytes := make([]byte, 8)
	for i := range bytes {
		bytes[i] = byte(lower + rand.Intn(upper-lower+1))
	}
	return string(bytes)
}

func randomName() (string, string) {
	index := rand.Intn(500)
	nameData := external.GetNameData()
	return nameData[index].Name, nameData[index].Surname
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
