package datagen

import (
	"github.com/chaseWilliams/family-map/lib/models"
)

/*
CreateFamily will randomly create ancestors for the given person, 
up until numGenerations generations
*/
func CreateFamily(person models.Person, numGenerations int) (err error) {
	/*
	root := create a completely random person object
	should only go to depth 2 (no infinite recursion)
	for each generation, do the following (starting with root):
	1. events for each person that is not "person" (
		NOTE: simulation occurs by evaluating all regular event probabilities yearly
		locations: random city based on choosing a population of cities:
			65%: within 50 mi of last location
			20%: within 300 mi
			15%: anywhere
		static events: 
					birth,
					death (
						constraints: cannot live past 120
						prob: normal distribution around 75 years old
					)
		regular events:
					marriage (
						constraints: youngest spouse is 18 years old minimum, 
							10 year difference max, not already married
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
}