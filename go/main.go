package main

import (
	"github.com/openeca/go/geneticalgorithms/mutation"
	"fmt"
)

func main() {
	arr := []int{1}
	translocation := mutation.Translocation(arr, 1, 1)
	printSlice(translocation)
}

func printSlice(sl []int ) {
	fmt.Print(sl)
}