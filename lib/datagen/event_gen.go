package datagen

type eventCheck func(*familyTreeNode, int)

type familyTreeNode struct {
	models.Person
	father     *familyTreeNode
	mother     *familyTreeNode
	spouse     *familyTreeNode
	events     []models.Event
	straight   bool
	birthYear  int // this is in simulation time, not in AD
	dead       bool
	generation int
}


func (f *familyTreeNode) evaluateYear(int year) {
	for _, generator := range eventGens {
		generator(f, year)
	}
}

var (
	eventGens = []eventCheck{
		deathCheck
	}
)

func deathCheck(node *familyTreeNode, year int) {

}