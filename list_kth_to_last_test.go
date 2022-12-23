package main

import (
	"testing"

	d "github.com/Savemaker/interviewprep/datastructures"
)

func TestKthToLast(t *testing.T) {
	testCases := []struct {
		array         []int
		kth, expected int
	}{
		{array: []int{0, 1, 2, 3}, kth: 2, expected: 1},
		{array: []int{4, 3}, kth: 0, expected: 3},
		{array: []int{1, 2, 3, 4}, kth: 3, expected: 1},
	}

	for _, testCase := range testCases {
		list := d.CreateListFromArray(testCase.array)
		res := KthToLast(list, testCase.kth)
		if res != testCase.expected {
			t.Errorf("expected %v but got %v for kth to last %v in %v", testCase.expected, res, testCase.kth, testCase.array)
		}
	}
}
