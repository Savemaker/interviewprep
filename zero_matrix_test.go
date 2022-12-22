package main

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	testCases := []struct {
		input, output [][]int
	}{
		{
			input: [][]int{
				{1, 0, 2},
				{4, 5, 0},
			},
			output: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			input: [][]int{
				{1, 0, 2},
				{4, 5, 6},
			},
			output: [][]int{
				{0, 0, 0},
				{4, 0, 6},
			},
		},
		{
			input: [][]int{
				{1, 0, 2, 3},
				{4, 5, 6, 5},
				{0, 5, 6, 0},
			},
			output: [][]int{
				{0, 0, 0, 0},
				{0, 0, 6, 0},
				{0, 0, 0, 0},
			},
		},
	}

	for _, v := range testCases {
		res := ZeroMatrix(v.input)
		if !reflect.DeepEqual(res, v.output) {
			t.Errorf("expected %v but got %v for %v", v.output, res, v.input)
		}
	}
}
