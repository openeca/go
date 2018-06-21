package mutation

import (
	"fmt"
	"testing"
)

type TestMessage struct {
	result  bool
	msg string
}

func TestTranslocation(t *testing.T) {
	arr := []int{1,2,3,4,5}
	translocation := Translocation(arr, 1, 3)
	var response TestMessage = ValidateSlices([]int{1,3,4,2,5}, translocation);
	
	if !response.result {
		t.Error(response.msg)
	}	
}	

func TestInversion(t *testing.T) {
	arr := []int{1,2,3,4,5}
	inversion := Inversion(arr,1,3) 
	var response TestMessage = ValidateSlices([]int{1,3,2,4,5}, inversion);
	
	if !response.result {
		t.Error(response.msg)
	}
}

func TestTranspositionTable(t *testing.T) {
	var tests = []struct {
		input []int
		expected []int
	} {
		{ []int{1,2,3,4,5}, Transposition([]int{1,4,3,2,5},1,3) },
		{ []int{1,2}, Transposition([]int{2,1},0,1) },
		{ []int{1}, Transposition([]int{1},0,0) },
	}

	for _, test := range tests {
		if response := ValidateSlices( test.input, test.expected); !response.result {
			t.Error(response.msg)
		}
	}
}

func ValidateSlices(s1 []int, s2 []int) TestMessage {
	if len(s1) != len(s2) {
		panic("Slices of different length")
	}
	for v := range s1 {
		if s1[v] != s2[v] {
			return TestMessage{false, fmt.Sprintf("%s%d",  "Difference at index: ", v )}
		}
	}
	return TestMessage{true, "OK"}
}