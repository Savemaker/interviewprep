package main

import "testing"

func Test(t *testing.T) {
	str := "xsgsa"
	if HasUniqueChars(str) {
		t.Error("Expected false but got true")
	}
}
