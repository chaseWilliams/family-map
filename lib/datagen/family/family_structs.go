package family

import (
	"database/sql"
	"fmt"
	"github.com/chaseWilliams/family-map/lib/datagen/external"
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
	"math/rand"
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

func (f *Person) createEvent(name string, year int) models.Event {
	var city external.City
	if len(f.events) == 0 {
		city = external.RandomCity()
	} else {
		recentEvent := f.events[len(f.events)-1]
		city = external.RandomCloseCity(recentEvent.Latitude, recentEvent.Longitude)
	}
	event := models.Event{
		EventID:   util.RandomID(),
		PersonID:  f.model.PersonID,
		Latitude:  city.Latitude,
		Longitude: city.Longitude,
		Country:   city.Country,
		City:      city.City,
		EventType: name,
		Year:      year,
	}
	f.events = append(f.events, event)
	return event
}

func (f *Person) createMirrorEvent(event models.Event) {
	event.EventID = util.RandomID()
	event.PersonID = f.model.PersonID
	f.events = append(f.events, event)
}

/*
NumEvents returns the person's number of events
*/
func (f Person) NumEvents() int {
	return len(f.events)
}

/*
Save will persist the person and their events in the database
*/
func (f *Person) Save(username string) (err error) {
	f.model.Username = username
	err = f.model.Save()
	if err != nil {
		return
	}
	for _, event := range f.events {
		event.Username = username
		err = event.Save()
		if err != nil {
			return
		}
	}
	return
}

/*
Dies will appropriately set the Person as dead at given year
*/
func (f *Person) Dies(year int) {
	f.deathYear = year
	f.createEvent("DEATH", year)
}

/*
Born will set the person's birth year and create the birth event
*/
func (f *Person) Born(year int) {
	f.birthYear = year
	f.createEvent("BIRTH", year)
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
	f.model.SpouseID = sql.NullString{spouse.model.PersonID, true}
	event := f.createEvent("MARRIAGE", year)

	spouse.spouses = append(spouse.spouses, f)
	spouse.marriageYears = append(spouse.marriageYears, year)
	spouse.married = true
	spouse.model.SpouseID = sql.NullString{f.model.PersonID, true}
	spouse.createMirrorEvent(event)
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
	f.model.SpouseID = sql.NullString{"", false}
	event := f.createEvent("DIVORCE", year)

	spouse.divorceYears = append(spouse.divorceYears, year)
	spouse.married = false
	spouse.model.SpouseID = sql.NullString{"", false}
	spouse.createMirrorEvent(event)
}

/*
HaveChild will edit the person's children and add the newborn event
*/
func (f *Person) HaveChild(child *Person, year int) {
	spouse, err := f.CurrSpouse()
	if err != nil {
		panic(err)
	}
	f.children = append(f.children, child)
	spouse.children = append(spouse.children, child)

	event := f.createEvent("NEWBORN", year)
	spouse.createMirrorEvent(event)
}

/*
HaveParents will set the parents of the person
*/
func (f *Person) HaveParents(father *Person, mother *Person) {
	f.father = father
	f.mother = mother
	f.model.FatherID = sql.NullString{father.model.PersonID, true}
	f.model.MotherID = sql.NullString{mother.model.PersonID, true}
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

/*
Children will return a map of all spouses -> slice of children
*/
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
AddPerson will add the person to the proper generation in the population
*/
func (pop *Population) AddPerson(f *Person) {
	generation := 0
	for i, gen := range *pop {
		for _, p := range gen {
			if (p == f.mother || p == f.father) && i >= generation {
				generation = i + 1
			}
		}
	}
	if generation >= len(*pop) {
		// should only need to add one more generation
		*pop = append(*pop, make(Generation, 0))
	}
	(*pop)[generation] = append((*pop)[generation], f)
}

/*
AreFamily returns whether or not the two people are family members
*/
func AreFamily(a *Person, b *Person) bool {
	if a == b {
		return false
	}
	// are parents or siblings x removed
	if recursiveAreParentsOrSiblings(a, b) {
		return true
	}

	// are spouses
	for _, spouse := range a.spouses {
		if spouse == b {
			return true
		}
		if recursiveAreParentsOrSiblings(spouse, b) {
			return true
		}
	}
	for _, spouse := range b.spouses {
		if spouse == a {
			return true
		}
		if recursiveAreParentsOrSiblings(spouse, a) {
			return true
		}
	}

	// are cousins
	for _, parentA := range []*Person{a.mother, a.father} {
		if parentA == nil {
			continue
		}
		for _, parentB := range []*Person{b.mother, b.father} {
			if parentB == nil {
				continue
			}
			if areSiblingsOrParents(parentA, parentB) {
				return true
			}
		}
	}

	return false
}

func recursiveAreParentsOrSiblings(a *Person, b *Person) bool {
	// recursion up the tree
	if goUpTree(a, b) || goUpTree(b, a) {
		return true
	}
	// recusion down the tree
	if goDownTree(a, b) || goDownTree(b, a) {
		return true
	}
	return false
}

func areSiblingsOrParents(a *Person, b *Person) bool {
	// is either one parents of the other
	if a.mother == b ||
		a.father == b ||
		b.mother == a ||
		b.father == a {
		return true
	}
	// are siblings
	if (a.mother == b.mother && a.mother != nil) ||
		(a.father == b.father && a.father != nil) {
		return true
	}
	return false
}

func goUpTree(a *Person, b *Person) bool {
	if a == nil || b == nil {
		return false
	}
	if areSiblingsOrParents(a, b) {
		return true
	}
	if goUpTree(a.father, b) ||
		goUpTree(a.mother, b) {
		return true
	}
	return false
}

func goDownTree(a *Person, b *Person) bool {
	if a == nil || b == nil {
		return false
	}
	if areSiblingsOrParents(a, b) {
		return true
	}
	for _, child := range a.children {
		if goDownTree(child, b) {
			return true
		}
	}
	return false
}

/*
RandomFamily returns a random family, determined by a random person at generation
numGen and all family members of that person. The person will be the first
person in the returned slice of people.
*/
func (pop Population) RandomFamily(personModel models.Person, numGen int) []*Person {
	genIndex := numGen - 1
	if len(pop) < genIndex {
		panic(fmt.Sprintf(
			"Population has %d generations, cannot create family at %d generation",
			len(pop),
			numGen,
		))
	}
	generation := pop[genIndex]
	person := generation[rand.Intn(len(generation))]
	// overrides the person's attributes in favor of what's provided
	person.model.PersonID = personModel.PersonID
	person.model.Username = personModel.Username
	person.model.FirstName = personModel.FirstName
	person.model.LastName = personModel.LastName
	person.model.Gender = personModel.Gender
	for i := range person.events {
		person.events[i].PersonID = personModel.PersonID
	}

	familyMembers := []*Person{person}
	for _, gen := range pop {
		for _, stranger := range gen {
			if AreFamily(person, stranger) {
				familyMembers = append(familyMembers, stranger)
			}
		}
	}
	return familyMembers
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
