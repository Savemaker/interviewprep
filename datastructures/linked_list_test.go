package datastructures

import "testing"

func TestLinkedList(t *testing.T) {
	for desc, fn := range map[string]func(t *testing.T){
		"test list creation": testListCanBeCreated,
	} {
		t.Run(desc, func(t *testing.T) {
			fn(t)
		})
	}
}

func testListCanBeCreated(t *testing.T) {
	testCases := []struct {
		array    []int
		expected string
	}{
		{array: []int{1, 2, 3}, expected: "123"},
		{array: []int{1, 1, 1}, expected: "111"},
		{array: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: "0123456789"},
	}

	for _, testCase := range testCases {
		list := CreateListFromArray(testCase.array)
		res := list.ConvertListToString()
		if res != testCase.expected {
			t.Errorf("expected %v but got %v for %v", testCase.expected, res, testCase.array)
		}
	}
}
