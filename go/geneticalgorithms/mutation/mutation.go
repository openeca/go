package mutation

// type IMutation interface {
// 	Translocation(arr []int, geneIndex int, newGeneIndex int) []int;
// 	Transposition(arr []int, gene1Index int, gene2Index) []int
// 	NPointExchange(arr []int, startGeneIndex int, nPoint int)
// 	Inversion(arr []int, startGeneIndex int, endGeneIndex int)
// }

//Translocation selects a gene at @geneIndex and moves it to a new index i.e @newgeneIndex.
// Transolcation shifts everything between down or up depending on the selection order.
// RETURNS A COPY
func Translocation(arr []int, geneIndex int, newGeneIndex int) []int {
	var gene int = arr[geneIndex]
	var finArr []int = make([]int, len(arr))
	for i:= range arr {
		if i < geneIndex || i > newGeneIndex {
			finArr[i] = arr[i]
		}else if i == geneIndex && i < (len(arr)-1) {
			finArr[i] = arr[i+1]
		}else if i > geneIndex && i < newGeneIndex {
			finArr[i] = arr[i+1]
		}else if i == newGeneIndex {
			finArr[i] = gene
		}
	}
	return finArr;
}

//Transposition takes two elements at positions gene1Index and gene2Index and swaps them.
//<RETURNS A COPY>
func Transposition(arr []int, gene1Index int, gene2Index int) []int {
	gene1 := arr[gene1Index];
	gene2 := arr[gene2Index];
	var finArr []int = make([]int, len(arr))

	for i := range arr {
		if i == gene1Index {
			finArr[i] = gene2
		}else if i == gene2Index {
			finArr[i] = gene1
		}else {
			finArr[i] = arr[i];
		}
	}
	return finArr;
}

//NPointExchange is a specialized variation of an Inversion, where you define an nPoint which is an integer whose value must be less than the len(arr) and whose value added to the startGeneIndex must be less than len(arr). The NPoint is simply the number of elements to invert after the startGenendex. Given arr := [1,2,3,4,5], startGeneIndex=1 and NPoint=3, the result would be [1,3,2,4,5]
//<RETURNS A COPY>
func NPointExchange(arr []int, startGeneIndex int, nPoint int) []int {
	if(startGeneIndex >= len(arr) ) {
		panic("startGeneIndex (param1)  cannot be greater than or equal to the array (param 0) length")
	}
	if(nPoint+startGeneIndex > len(arr)) {
		panic("The sum of startGeneIndex (param1) nPoint(param2)  cannot be greater than  the array (param 0) length")
	}
	if(nPoint < 0) {
		panic("The nPoint (param2) cannot be less than 0")
	}
	if(nPoint > 0 && nPoint <=1) {
		return arr
	}
	endGeneIndex := startGeneIndex + nPoint;
	invRange := endGeneIndex - startGeneIndex;
	invGenes := make([]int, invRange, invRange)
	var finArr []int = make([]int, len(arr))

	for i := 0; i <  invRange; i++ {
		invGenes[i] = arr[(endGeneIndex-1)-i]
	}

	for i:= range arr {
		if(i < startGeneIndex) {
			finArr[i] = arr[i]
		}else if(i < endGeneIndex) {
			finArr[i] = invGenes[(i - startGeneIndex)]
		}else {
			finArr[i] = arr[i]
		}
	}
	return finArr;
}

//An inversion inverts a section of an array starting from the startGeneIndex and terminating at endGeneIndex. (Not including the endGeneIndex). Given arr := [1,2,3,4,5], startGeneIndex=1 and endGeneIndex=3, the result would be [1,3,2,4,5]
//<RETURNS A COPY>
func Inversion(arr []int, startGeneIndex int, endGeneIndex int) []int {
	invRange := endGeneIndex - startGeneIndex;
	invGenes := make([]int, invRange, invRange)
	var finArr []int = make([]int, len(arr))

	for i := 0; i <  invRange; i++ {
		invGenes[i] = arr[(endGeneIndex-1)-i]
	}

	for i:= range arr {
		if(i < startGeneIndex) {
			finArr[i] = arr[i]
		}else if(i < endGeneIndex) {
			finArr[i] = invGenes[(i - startGeneIndex)]
		}else {
			finArr[i] = arr[i]
		}
	}
	return finArr;
}

//Performs a deepcopy of the array, returning a copy of the array
func DeepCopy(arr []int) []int {
	var finArr []int = make([]int,len(arr), len(arr))
	for i:= range arr {
		finArr[i] = arr[i]
	}
	return finArr
}