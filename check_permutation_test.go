package main

import "testing"

func TestPerm(t *testing.T) {
	var permTest = []struct {
		a, b   string
		isPerm bool
	}{
		{a: "abcd", b: "bdca", isPerm: true},
		{a: "dca", b: "acd", isPerm: true},
		{a: "ahm", b: "", isPerm: false},
		{a: "anna", b: "anna", isPerm: true},
	}

	for _, v := range permTest {
		actual := IsPermutation(v.a, v.b)
		if actual != v.isPerm {
			t.Errorf("expected %v but got %v for pair {%v %v)", v.isPerm, actual, v.a, v.b)
		}
	}
}
