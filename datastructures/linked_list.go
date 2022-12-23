package datastructures

import (
	"strconv"
	"strings"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Next  *Node
	Value int
}

func (l *LinkedList) Insert(val int) {
	element := Node{Next: nil, Value: val}
	head := l.Head
	if head == nil {
		l.Head = &element
	} else {
		for head.Next != nil {
			head = head.Next
		}
		head.Next = &element
	}
}

func CreateListFromArray(array []int) *LinkedList {
	list := LinkedList{}
	for _, value := range array {
		list.Insert(value)
	}
	return &list
}

func (l LinkedList) ConvertListToString() string {
	head := l.Head
	builder := strings.Builder{}
	for head != nil {
		builder.WriteString(strconv.Itoa(head.Value))
		head = head.Next
	}
	return builder.String()
}
