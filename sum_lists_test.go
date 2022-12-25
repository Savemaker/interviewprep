package main

import (
	"testing"

	d "github.com/Savemaker/interviewprep/datastructures"
)

func TestSumLists(t *testing.T) {
	testCases := []struct {
		x, y     []int
		expected string
	}{
		{x: []int{7, 1, 6}, y: []int{5, 9, 2}, expected: "219"},
		{x: []int{9, 5}, y: []int{1}, expected: "06"},
		{x: []int{1}, y: []int{9, 5}, expected: "06"},
		{x: []int{}, y: []int{}, expected: ""},
		{x: []int{8, 8, 8}, y: []int{9, 9, 9}, expected: "7881"},
	}
	for _, testCase := range testCases {
		x := d.CreateListFromArray(testCase.x)
		y := d.CreateListFromArray(testCase.y)
		result := SumLists(x, y)
		resStr := result.ConvertListToString()
		if resStr != testCase.expected {
			t.Errorf("expected %v but got %v for sum of %v and %v", testCase.expected, resStr, testCase.x, testCase.y)
		}
	}
}
