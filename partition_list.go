/*
	Partition: Write code to partition a linked list around a value x,
	such that all nodes less than x come before all nodes greater than or equal to x.
	If x is contained within the list,
	the values of x only need to be after the elements less than x (see below).
	The partition element x can appear anywhere in the "right partition";
	it does not need to appear between the left and right partitions.

	EXAMPLE
	Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1[partition=5]
	Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
*/

package main

import (
	d "github.com/Savemaker/interviewprep/datastructures"
)

func PartitionList(list *d.LinkedList, val int) {
	head := list.Head
	var leftP, rightP, lastLeft, lastRight *d.Node

	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = nil

		if tmp.Value < val {
			if leftP == nil {
				leftP = tmp
				lastLeft = tmp
			} else {
				lastLeft.Next = tmp
				lastLeft = tmp
			}
		} else {
			if rightP == nil {
				rightP = tmp
				lastRight = tmp
			} else {
				lastRight.Next = tmp
				lastRight = tmp
			}
		}
	}
	if lastLeft != nil {
		lastLeft.Next = rightP
		list.Head = leftP
	}
}
