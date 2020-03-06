package family

import (
	"github.com/chaseWilliams/family-map/lib/database"
	"testing"
)

func TestIsFamily(t *testing.T) {
	database.StartTestingSession(t)
	person := RandomPerson(0)

	grandfather := RandomPerson(0)
	grandmother := RandomPerson(0)

	father := RandomPerson(0)
	mother := RandomPerson(0)
	mother.HaveParents(grandfather, grandmother)
	father.Marry(mother, 0)

	person.HaveParents(father, mother)

	sibling := RandomPerson(0)
	sibling.HaveParents(father, mother)

	halfSibling := RandomPerson(0)
	halfSibling.HaveParents(father, RandomPerson(0))

	uncle := RandomPerson(0) // mother's brother
	aunt := RandomPerson(0)  // mother's brother's spouse
	uncle.HaveParents(grandfather, grandmother)
	uncle.Marry(aunt, 0)

	cousin := RandomPerson(0)
	cousin.HaveParents(uncle, aunt)

	greatGrandfather := RandomPerson(0)
	greatGrandmother := RandomPerson(0)
	greatUncle := RandomPerson(0)

	grandfather.HaveParents(greatGrandfather, greatGrandmother)
	greatUncle.HaveParents(greatGrandfather, greatGrandmother)

	firstCousinOnceRemoved := RandomPerson(0)
	firstCousinOnceRemoved.HaveParents(greatUncle, RandomPerson(0))

	secondCousin := RandomPerson(0)
	secondCousin.HaveParents(firstCousinOnceRemoved, RandomPerson(0))

	spouse := RandomPerson(0)
	person.Marry(spouse, 0)

	spouseGrandFather := RandomPerson(0)
	spouseMother := RandomPerson(0)
	spouseMother.HaveParents(spouseGrandFather, RandomPerson(0))
	spouse.HaveParents(RandomPerson(0), spouseMother)

	spouseUncle := RandomPerson(0)
	spouseUncle.HaveParents(spouseGrandFather, RandomPerson(0))

	spouseCousin := RandomPerson(0)
	spouseCousin.HaveParents(spouseUncle, RandomPerson(0))

	tests := []struct {
		in       *Person
		relation string
		out      bool
	}{
		{grandfather, "grandfather", true},
		{grandmother, "grandmother", true},
		{father, "father", true},
		{mother, "mother", true},
		{sibling, "sibling", true},
		{halfSibling, "halfSibling", true},
		{uncle, "uncle", true},
		{aunt, "aunt", true},
		{cousin, "cousin", true},
		{greatGrandfather, "greatGrandfather", true},
		{greatGrandmother, "greatGrandmother", true},
		{greatUncle, "greatUncle", true},
		{firstCousinOnceRemoved, "firstCousinOnceRemoved", false},
		{secondCousin, "secondCousin", false},
		{spouse, "spouse", true},
		{spouseGrandFather, "spouseGrandFather", true},
		{spouseMother, "spouseMother", true},
		{spouseUncle, "spouseUncle", true},
		{spouseCousin, "spouseCousin", false},
		{RandomPerson(0), "randomPerson", false},
	}

	for _, test := range tests {
		t.Logf("%s: %s", test.relation, test.in.model.FirstName)
	}

	for _, test := range tests {
		check := AreFamily(person, test.in)
		if check != test.out {
			t.Errorf(
				"relationship comparing person -> %s got %v, should be %v",
				test.relation,
				check,
				test.out,
			)
		}
		check = AreFamily(test.in, person)
		if check != test.out {
			t.Errorf(
				"relationship comparing %s -> person got %v, should be %v",
				test.relation,
				check,
				test.out,
			)
		}
	}
}
