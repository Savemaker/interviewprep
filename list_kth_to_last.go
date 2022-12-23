package main

import (
	d "github.com/Savemaker/interviewprep/datastructures"
)

func KthToLast(list *d.LinkedList, kth int) int {
	a, b := list.Head, list.Head
	for i := 0; i <= kth; i++ {
		if a != nil {
			a = a.Next
		} else {
			return -1
		}
	}

	for a != nil {
		a = a.Next
		b = b.Next
	}
	return b.Value
}
