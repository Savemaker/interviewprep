/*
	Sum Lists: You have two numbers represented by a linked list, where each node contains a single digit.
	The digits are stored in reverse order, such that the 1 's digit is at the head of the list.
	Write a function that adds the two numbers and returns the sum as a linked list.
	EXAMPLE
	Input:(7-> 1 -> 6) + (5 -> 9 -> 2).Thatis,617 + 295. Output:2 -> 1 -> 9.Thatis,912.
*/

package main

import (
	d "github.com/Savemaker/interviewprep/datastructures"
)

func SumLists(x, y *d.LinkedList) *d.LinkedList {
	res := d.LinkedList{}
	var curNode *d.Node
	var head *d.Node

	headA, headB := x.Head, y.Head
	step := 1
	toAdd := 0

	for headA != nil || headB != nil {
		node := d.Node{}
		if headA != nil && headB != nil {
			sum := toAdd + headA.Value + headB.Value
			rem := sum % 10
			toAdd = sum / 10
			node.Value = rem
			headA = headA.Next
			headB = headB.Next
		} else if headA != nil {
			node.Value = headA.Value + toAdd
			toAdd = 0
			headA = headA.Next
		} else {
			node.Value = headB.Value + toAdd
			toAdd = 0
			headB = headB.Next
		}
		if curNode == nil {
			curNode = &node
			head = &node
		} else {
			curNode.Next = &node
			curNode = &node
		}
		step *= 10
	}
	if toAdd != 0 {
		node := d.Node{Value: +toAdd}
		if curNode == nil {
			curNode = &node
			head = &node
		} else {
			curNode.Next = &node
			curNode = &node
		}
	}
	res.Head = head
	return &res
}
