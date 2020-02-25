package datagen

import (
	"github.com/chaseWilliams/family-map/lib/models"
	"math"
)

/*
TODO
implement all eventChecks
in generateAncestors:
	initial population should all be clustered somewhere
	after population is completely created, choose someone in the numGen generation
	to be the user. Then, all family members of the user are saved to the database.
In marriageCheck, make sure to include corner case where there is no one available to marry
*/

/*
======================
	FAMILY MEMBERS
======================
Family members are thus defined:
	1. are parents of X
	2. share at least one parent with X (siblings)
	3. share parents who share parents (cousins)
	4. was previously or is a spouse of X
	5. rules 1 and 2 can be recursively descended, and all persons from those rules are
		family members of X. 
	6. all previously defined people's spouses are also family, thus extending rule 4,
		but this is not recursive
	7. rule 4 can be immediately entered for X, and then extended with all rules but 3
	
	That is, the grandparents of X are family, X's grandparent's sibling is family, 
	the cousin of X is family, X's grandparent's sibling's (ex)-spouse is family,
	X's spouse's ex-spouse is family, X's spouse's sibling's spouse is family,
	X's spouse's grandparent's sibling's spouse is family, and X's cousin's spouse is family.

	However, X's second cousin of X is not family, X's first cousin once-removed is not family,
	and X's spouse's cousin is not family.

	Notably missing are nephews, nieces, and extensions of that, but this simulation assumes 
	that the user has no children, extended from the inference that the user has no events but birth.

	The point of the simulation is simply to create ancestor data, so events specific to and 
	succeeding the user are possible (and are actually generated), but ultimately ignored.

	Hopefully that covers everything.
*/

/*
===============================
HIGH LEVEL SIMULATION ALGORITHM
===============================

root := create a completely random person object
should only go to depth 2 (no infinite recursion)
for each generation, do the following (starting with root):
1. events for each person  (
	NOTE: simulation occurs by evaluating all regular event probabilities yearly
	locations: random city based on choosing a population of cities:
		65%: within 50 mi of last location
		20%: within 300 mi
		15%: anywhere
	static events:
				birth
	regular events:
				death (
					constraints: cannot live past 120
					prob: normal distribution around 75 years old
				)
				marriage (
					constraints: youngest spouse is 18 years old minimum,
						10 year difference max, not already married,
						not family members
					prob: 40% - (age * 0.7) + 18
					prob of same sex marriage: 5%
					side effect: create spouse and their family tree
				),
				graduation (
					constraints: ages 22-25 only
					prob: 90% * (1 / (father didn't go + mother didn't go + 1))
					location: guaranteed to choose in a pool of 500 mi, doesn't update
						last location for future events
				),
				divorced (
					constraints: must be married
					prob: 0.05%
				),
				widowed (
					constraints: spouse dies
				),
				changed career (
					constraints: ages 18+
					prob: 0.5% * (years since last career change)
				),
				had a child (
					constraints: female parent cannot be older than 50, must be during
						a heterosexual marriage
					prob: (80% - (avg age of parents) * 1.2) * ( 1 / (num children * 0.5 + 1))
				)
)
2. create child persons
	a. if no child persons created and not on last generation, retry step one
	b. else, go to next generation (do step one for children)
*/


func generateAncestors(child familyTreeNode, numGens int) {
	initPopulation := int(math.Pow(2, float64(numGens)) * 10) // starting population size
	pop := make(population, numGens)
	// set first population
	pop[0] = make(generation, initPopulation)
	for i := 0; i < initPopulation; i++ {
		pop[0][i] = randomPerson(0)
	}

	year := 0
	SimulationLoop:
		for true {
			yearEventCheck(&pop, evaluateYear, year)
			year++
			// when everyone in the numGens generation and previous generations are dead,
			// we stop the simulation
			for i := 0; i < numGens; i++ {	
				if !pop[i].allDead() {
					continue SimulationLoop
				}
			}
			break
		}
}

/*
CreateFamily will randomly create ancestors for the given person,
up until numGens generations
*/
func CreateFamily(person models.Person, numGens int) {
	generateAncestors(randomPerson(0), numGens)
	return
}
