package family

import (
	"fmt"
	"github.com/chaseWilliams/family-map/lib/models"
)

/*
Person is a family construct that represents a singular person
in a family tree
*/
type Person struct {
	model         models.Person
	events        []models.Event
	father        *Person
	mother        *Person
	spouses       []*Person
	children      []*Person
	straight      bool
	marriageYears []int
	divorceYears  []int
	married       bool
	birthYear     int // this is in simulation time, not in AD
	deathYear     int
}

// TODO
// create events
// instead of divorce on death, widow instead

/*
Dies will appropriately set the Person as dead at given year
*/
func (f *Person) Dies(year int) {
	f.deathYear = year
	if f.married {
		f.Divorce(year)
	}
}

/*
Marry will set appropriate fields for Person now being married.
*/
func (f *Person) Marry(spouse *Person, year int) {
	if f.married {
		panic(fmt.Sprintf("Person is already married\n%v", *f))
	}
	f.spouses = append(f.spouses, spouse)
	f.marriageYears = append(f.marriageYears, year)
	f.married = true

	spouse.spouses = append(spouse.spouses, f)
	spouse.marriageYears = append(spouse.marriageYears, year)
	spouse.married = true
}

/*
Divorce will set appropriate fields for Person getting divorced.
*/
func (f *Person) Divorce(year int) {
	if !f.married {
		panic(fmt.Sprintf("Person is not married\n%v", *f))
	}
	
	spouse, err := f.CurrSpouse()
	if err != nil {
		panic(err)
	}

	f.divorceYears = append(f.divorceYears, year)
	f.married = false

	spouse.divorceYears = append(spouse.divorceYears, year)
	spouse.married = false
}

func (f *Person) HaveChild(child *Person, year int) {
	f.children = append(f.children, child)
}

func (f *Person) HaveParents(father *Person, mother *Person) {
	f.father = father
	f.mother = mother
}

/*
IsDead returns if the person is dead or not
*/
func (f Person) IsDead() bool {
	return f.deathYear != -1
}

/*
IsMarried returns if the person is married or not
*/
func (f Person) IsMarried() bool {
	return f.married
}

/*
IsStraight returns if the person is straight or not
*/
func (f Person) IsStraight() bool {
	return f.straight
}

/*
Gender will return the person's gender
*/
func (f Person) Gender() string {
	return f.model.Gender
}

/*
Age returns the person's year, given a year
*/
func (f Person) Age(year int) int {
	return year - f.birthYear
}

/*
DeathYear returns the person's DeathYear, which is -1 if
the person isn't dead yet
*/
func (f Person) DeathYear() int {
	return f.deathYear
}

/*
Spouses returns a slice of the person's spouses in
chronological order
*/
func (f Person) Spouses() []*Person {
	return f.spouses
}

/*
CurrSpouse will return the person's current spouse,
and return an error if person isn't married
*/
func (f Person) CurrSpouse() (spouse *Person, err error) {
	if !f.married {
		return nil, fmt.Errorf("person is not married:\n%v", f)
	}
	return f.spouses[len(f.spouses)-1], nil
}

func (f Person) Children() (m map[*Person][]*Person) {
	m = make(map[*Person][]*Person)
	for _, spouse := range f.spouses {
		for _, child := range f.children {
			if child.father == spouse || child.mother == spouse {
				if _, ok := m[spouse]; !ok {
					m[spouse] = make([]*Person, 0)
				}
				m[spouse] = append(m[spouse], child)
			}
		}
	}
	return
}

/*
MarriageYears will return the person's marriage years
*/
func (f Person) MarriageYears() []int {
	return f.marriageYears
}

/*
DivorceYears returns the person's divorce years
*/
func (f Person) DivorceYears() []int {
	return f.divorceYears
}

/*
Generation is what the name implies, and represented by a slice
of Person pointers, in order to keep everything mutable.
*/
type Generation []*Person

/*
AllDead will returns whether or not everyone in the generation
is dead
*/
func (g Generation) AllDead() bool {
	if g == nil || len(g) == 0 {
		return true
	}
	for _, p := range g {
		if !p.IsDead() {
			return false
		}
	}
	return true
}

/*
Population is what the name implies, and represented by a slice
of Generations
*/
type Population []Generation

/*
GetAlive returns a slice of all alive people, irrespective of their generation
*/
func (pop *Population) GetAlive() []*Person {
	people := make([]*Person, 0)
	for _, gen := range *pop {
		for _, p := range gen {
			if !p.IsDead() {
				people = append(people, p)
			}
		}
	}
	return people
}

/*
AreFamily returns whether or not the two people are family members
*/
func (pop Population) AreFamily(a Person, b Person) bool {
	return false
}

/*
CreatePopulation will create a Population with one generation
that has the specified people in it.
*/
func CreatePopulation(numPeople int) Population {
	pop := make(Population, 1)
	pop[0] = make(Generation, 0, numPeople)
	// create the born people, born at year 0
	for i := 0; i < numPeople; i++ {
		pop[0] = append(pop[0], RandomPerson(0))
	}
	return pop
}
