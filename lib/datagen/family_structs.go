package datagen

import (
	"github.com/chaseWilliams/family-map/lib/models"
)

type familyTreeNode struct {
	models.Person
	father     *familyTreeNode
	mother     *familyTreeNode
	spouse     *familyTreeNode
	events     []models.Event
	straight   bool
	birthYear  int // this is in simulation time, not in AD
	deathYear int
}

func (f familyTreeNode) isDead() bool {
	return f.deathYear != -1
}

func (f familyTreeNode) getAge(year int) int {
	return year - f.birthYear
}

type generation []familyTreeNode

func (g generation) allDead() bool {
	if g == nil {
		return true
	}
	for _, person := range g {
		if !person.isDead() {
			return false
		}
	}
	return true
}

type population []generation

func (p population) getAlive() []familyTreeNode {
	people := make([]familyTreeNode, 10)
	for _, gen := range p {
		for _, person := range gen {
			if !person.isDead() {
				people = append(people, person)
			}
		}
	}
	return people
}

func (p population) areFamily(a familyTreeNode, b familyTreeNode) bool {
	return false
}
