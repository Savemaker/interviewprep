package main

import "testing"

func TestOneAway(t *testing.T) {
	var testCases = []struct {
		a, b string
		res  bool
	}{
		{a: "pale", b: "ple", res: true},
		{a: "pales", b: "pale", res: true},
		{a: "pale", b: "bale", res: true},
		{a: "pale", b: "bake", res: false},
		{a: "apple", b: "bpple", res: true},
		{a: "pplea", b: "apple", res: false},
		{a: "aaaaple", b: "aaaple", res: true},
		{a: "aaaple", b: "aaaaple", res: true},
		{a: "", b: "", res: true},
		{a: "richard", b: "richrd", res: true},
		{a: "richard", b: "grichrd", res: false},
		{a: "michard", b: "richrd", res: false},
	}

	for _, v := range testCases {
		if OneAway(v.a, v.b) != v.res {
			t.Errorf("expected %v for %v and %v", v.res, v.a, v.b)
		}
	}
}
