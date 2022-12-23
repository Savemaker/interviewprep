package main

import (
	"testing"

	data "github.com/Savemaker/interviewprep/datastructures"
)

func TestDeleteMiddleNode(t *testing.T) {
	e := &data.Node{Value: 5, Next: nil}
	d := &data.Node{Value: 4, Next: e}
	c := &data.Node{Value: 3, Next: d}
	b := &data.Node{Value: 2, Next: c}
	a := &data.Node{Value: 1, Next: b}

	DeleteMiddleNode(b)

	size := 0

	for a != nil {
		size++
		a = a.Next
	}
	if size != 4 {
		t.Errorf("expected size 4 but got %v", size)
	}
}
