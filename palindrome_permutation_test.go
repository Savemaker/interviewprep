package main

import "testing"

func TestIsPermutationOfPalindrome(t *testing.T) {
	var testCases = []struct {
		input  string
		output bool
	}{
		{input: "Tact Coa", output: true},
		{input: "actocta", output: true},
		{input: "abba", output: true},
		{input: "aaaaacaaaaa", output: true},
		{input: "aabbabcaaaaa", output: false},
		{input: "aabbabcaaaaa", output: false},
		{input: "ppabba", output: true},
	}

	for _, v := range testCases {
		out := IsPermutationOfPalindrome(v.input)
		if out != v.output {
			t.Errorf("expected %v for %v", v.output, v.input)
		}
	}
}
