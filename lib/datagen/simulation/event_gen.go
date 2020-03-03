package simulation

import (
	"github.com/chaseWilliams/family-map/lib/datagen/family"
	"gonum.org/v1/gonum/stat"
	"math"
	"math/rand"
)

type eventCheck func(*family.Person, *family.Population, int)

var (
	eventGens = []eventCheck{
		deathCheck,
		marriageCheck,
		divorceCheck,
		babyCheck,
	}
)

func personYearCheck(f *family.Person, pop *family.Population, year int) {
	for _, generator := range eventGens {
		generator(f, pop, year)
	}
}

func populationYearCheck(pop *family.Population, year int) {
	for i := 0; i < len(*pop); i++ {
		for j := 0; j < len((*pop)[i]); j++ {
			person := (*pop)[i][j]
			if !person.IsDead() {
				personYearCheck(person, pop, year)
			}
		}
	}
}

func deathCheck(f *family.Person, pop *family.Population, year int) {
	if f.IsDead() {
		return
	}
	prob := 0.005
	age := f.Age(year)
	if age > 70 {
		prob = 0.05
	}
	roll := rand.Float64()
	maxAge := 120
	if roll <= prob || age >= maxAge {
		f.Dies(year)
	}
}

func divorceCheck(f *family.Person, pop *family.Population, year int) {
	if f.IsDead() || !f.IsMarried() {
		return
	}

	prob := 0.007
	roll := rand.Float64()
	if roll > prob {
		return
	}

	f.Divorce(year)
}

func marriageCheck(f *family.Person, pop *family.Population, year int) {
	// constraints
	if f.IsDead() || f.Age(year) < 18 || f.IsMarried() {
		return
	}

	// roll probability check
	prob := 0.4 - (float64(f.Age(year)) / 100)
	roll := rand.Float64()
	if roll > prob {
		return
	}

	// go about choosing people to marry
	var desiredGender string
	if f.IsStraight() {
		if f.Gender() == "m" {
			desiredGender = "f"
		} else {
			desiredGender = "m"
		}
	} else {
		if f.Gender() == "m" {
			desiredGender = "m"
		} else {
			desiredGender = "f"
		}
	}

	suitors := make([]*family.Person, 0)
	possibleSuitors := pop.GetAlive()

	// filter for people who match
	for _, person := range possibleSuitors {
		if person.IsStraight() == f.IsStraight() &&
			person.Gender() == desiredGender &&
			!person.IsMarried() &&
			person.Age(year) >= 18 &&
			math.Abs(float64(person.Age(year)-f.Age(year))) <= 10 &&
			!pop.AreFamily(*person, *f) &&
			person != f {
			suitors = append(suitors, person)
		}
	}

	if len(suitors) == 0 {
		return // no suitors
	}

	// choose person to marry
	spouse := suitors[rand.Intn(len(suitors))]

	// marry that person
	spouse.Marry(f, year)
}

// TODO put child into population
func babyCheck(f *family.Person, pop *family.Population, year int) {
	if f.IsDead() || !f.IsMarried() || !f.IsStraight() {
		return
	}

	spouse, err := f.CurrSpouse()
	if err != nil {
		panic(err)
	}

	// check that the wife is younger than 50
	var wife *family.Person
	var husband *family.Person
	if f.Gender() == "f" {
		wife = f
		husband = spouse
	} else {
		wife = spouse
		husband = f
	}

	if wife.Age(year) >= 50 {
		return
	}

	// roll
	children := f.Children()
	numChildrenWithSpouse := float64(len(children[spouse]))
	avgAge := stat.Mean([]float64{float64(f.Age(year)), float64(spouse.Age(year))}, nil)
	// prob: (80% - (avg age of parents) * 1.2) * ( 1 / (num children * 0.5 + 1))
	prob := (0.8 - (avgAge * 2.5 / 100)) * (1 / (numChildrenWithSpouse*0.25 + 1))
	roll := rand.Float64()
	if roll > prob {
		return
	}

	child := family.RandomPerson(year)
	child.HaveParents(husband, wife)
	husband.HaveChild(child, year)
	wife.HaveChild(child, year)
	pop.AddPerson(child)
}
