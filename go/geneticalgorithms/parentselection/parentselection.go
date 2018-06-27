package parentselection

type IParentSelection interface {
	tournamentSelection(n int)
	tournamentSelectionRandomN()
	random()
	ranking()
	elitist(chromosomes []int) int
	elitistN(n int)
	weakest()
	weakestN()
	highestDiversity(func())
	lowestDiversity(func())
	balanced(arr []int) 
	balancedN(n int)
}

