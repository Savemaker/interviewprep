package main

import "testing"

func TestUrlify(t *testing.T) {
	var testCases = []struct {
		input    []rune
		len      int
		expected []rune
	}{
		{
			input:    []rune{'M', 'r', ' ', 'J', 'o', 'h', 'n', ' ', 'S', 'm', 'i', 't', 'h', ' ', ' ', ' ', ' '},
			len:      13,
			expected: []rune{'M', 'r', '%', '2', '0', 'J', 'o', 'h', 'n', '%', '2', '0', 'S', 'm', 'i', 't', 'h'},
		},
		{
			input:    []rune{'S', 'm', 'i', 't', 'h', ' ', ' ', ' '},
			len:      6,
			expected: []rune{'S', 'm', 'i', 't', 'h', '%', '2', '0'},
		},
		{
			input:    []rune{' ', 'S', 'm', 'i', 't', 'h', ' ', ' '},
			len:      6,
			expected: []rune{'%', '2', '0', 'S', 'm', 'i', 't', 'h'},
		},
		{
			input:    []rune{'S', ' ', 'm', ' ', 'i', ' ', 't', ' ', 'h', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			len:      9,
			expected: []rune{'S', '%', '2', '0', 'm', '%', '2', '0', 'i', '%', '2', '0', 't', '%', '2', '0', 'h'},
		},
	}

	for _, v := range testCases {
		res := string(Urlify(v.input, v.len))
		if res != string(v.expected) {
			t.Errorf("Expected=%v but got %v", string(v.expected), res)
		}
	}
}
