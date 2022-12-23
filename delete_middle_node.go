package main

import (
	d "github.com/Savemaker/interviewprep/datastructures"
)

func DeleteMiddleNode(node *d.Node) {
	if node == nil || node.Next == nil {
		return
	}
	next := node.Next
	node.Value = next.Value
	node.Next = next.Next
}
