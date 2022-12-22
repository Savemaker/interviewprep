package main

import "testing"

func TestBasicStringCompression(t *testing.T) {
	testCases := []struct {
		input, output string
	}{
		{input: "aabcccccaaa", output: "a2b1c5a3"},
		{input: "abcdefg", output: "abcdefg"},
		{input: "", output: ""},
	}

	for _, v := range testCases {
		actual := CompressString(v.input)
		if actual != v.output {
			t.Errorf("exptected %v but got %v for %v", v.output, actual, v.input)
		}
	}

}
