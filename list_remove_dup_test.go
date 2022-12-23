package main

import (
	"testing"

	d "github.com/Savemaker/interviewprep/datastructures"
)

func TestRemoveDupFromList(t *testing.T) {
	type testCases struct {
		arrayToConvert []int
		expected       string
	}

	for _, testCase := range []testCases{
		{arrayToConvert: []int{1, 1, 1}, expected: "1"},
		{arrayToConvert: []int{-1, 0, 2, 3, -1}, expected: "-1023"},
		{arrayToConvert: []int{-1, -1, 2, 3, -1}, expected: "-123"},
		{arrayToConvert: []int{10, -1, 20, -1, -1}, expected: "10-120"},
		{arrayToConvert: []int{-1, -1, -1}, expected: "-1"},
	} {
		list := d.CreateListFromArray(testCase.arrayToConvert)
		RemoveDuplicates(list)
		res := list.ConvertListToString()
		if res != testCase.expected {
			t.Errorf("expected %v but got %v for %v", testCase.expected, res, testCase.arrayToConvert)
		}
	}
}
