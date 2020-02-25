package datagen

import (
	"math/rand"
	"math"
)

type eventCheck func(*familyTreeNode, *population, int)

var (
	eventGens = []eventCheck{ 
		deathCheck,
		marriageCheck,
	}
)

func evaluateYear(f *familyTreeNode, pop *population, year int) {
	for _, generator := range eventGens {
		generator(f, pop, year)
	}
}

func yearEventCheck(pop *population, check eventCheck, year int) {
	for i := 0; i < len(*pop); i++ {
		for j := 0; j < len((*pop)[i]); j++ {
			person := (*pop)[i][j]
			if !person.isDead() {
				check(&person, pop, year)
			}
		}
	}
}

func deathCheck(f *familyTreeNode, pop *population, year int) {
	if f.isDead() {
		return
	}
	prob := 0.005
	age := f.getAge(year)
	if age > 70 {
		prob = 0.05
	} 
	roll := rand.Float64() 
	maxAge := 120
	if roll <= prob || age >= maxAge {
		f.deathYear = year
	}
}

func marriageCheck(f *familyTreeNode, pop *population, year int) {
	if f.getAge(year) < 18 || f.spouse != nil {
		return
	}
	var desiredGender string
	if f.straight {
		if f.Gender == "m" {
			desiredGender = "f"
		} else {
			desiredGender = "m"
		}
	} else {
		if f.Gender == "m" {
			desiredGender = "m"
		} else {
			desiredGender = "f"
		}
	}

	suitors := make([]familyTreeNode, 10)
	// filter for people who match
	for _, person := range pop.getAlive() {
		if person.straight == f.straight &&
			person.Gender == desiredGender &&
			person.spouse == nil &&
			person.getAge(year) >= 18 &&
			math.Abs(float64(person.getAge(year) - f.getAge(year))) <= 10 &&
			!pop.areFamily(person, *f) {
			suitors = append(suitors, person)
		}
	}

	// choose person to marry
	spouse := suitors[rand.Intn(len(suitors))]

	// marry that person
	f.spouse = &spouse
	spouse.spouse = f

}