package main

import (
	"testing"

	d "github.com/Savemaker/interviewprep/datastructures"
)

func TestPartitionList(t *testing.T) {
	for _, testCase := range []struct {
		array    []int
		value    int
		expected string
	}{
		{array: []int{}, expected: "", value: 5},
		{array: []int{3, 5, 8, 5, 10, 2, 1}, expected: "32158510", value: 5},
		{array: []int{6, 5, 5, 3}, expected: "3655", value: 5},
		{array: []int{-1, 0, -99, 100}, expected: "-10-99100", value: 200},
	} {
		list := d.CreateListFromArray(testCase.array)
		PartitionList(list, testCase.value)
		res := list.ConvertListToString()
		if res != testCase.expected {
			t.Errorf("expected %v but got %v for %v with value %v", testCase.expected, res, testCase.array, testCase.value)
		}
	}
}
