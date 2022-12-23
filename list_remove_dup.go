/*
	Remove dups. Write code to remove duplicates from an unsorted linked list.
	FOLLOW UP
	How would you solve this problem if a temporary buffer is not allowed?

	Solution: it could be done easily with the hashmap,
	but I will solve it with 2 pointers in n^2 for a practice
*/

package main

import (
	d "github.com/Savemaker/interviewprep/datastructures"
)

func RemoveDuplicates(list *d.LinkedList) {
	i := list.Head

	for i != nil {
		j := i
		for j.Next != nil {
			if i.Value == j.Next.Value {
				j.Next = j.Next.Next
			} else {
				j = j.Next
			}
		}
		i = i.Next
	}
}
