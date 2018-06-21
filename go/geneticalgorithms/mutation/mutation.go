package mutation

// type IMutation interface {
// 	Translocation(arr []int, geneIndex int, newGeneIndex int) []int;
// 	Transposition(arr []int, gene1Index int, gene2Index) []int
// 	NPointExchange(arr []int, startGeneIndex int, nPoint int)
// 	Inversion(arr []int, startGeneIndex int, endGeneIndex int)
// }

//Given an input array @arr, Translocation selects an gene at @geneIndex and moves it to a new index i.e @newgeneIndex
// It shifts everything between down or up depending on the selection order.
//  <b>RETURNS COPY</b>
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