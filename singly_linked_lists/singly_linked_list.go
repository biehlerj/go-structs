package singly_linked_list

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Push(value interface{}) {
	newNode := &Node{
		value: value,
		next:  L.head,
	}
	L.head = newNode

	if L.tail == nil {
		L.tail = L.head
	}
}

func (L *List) InsertBefore(node *Node, value interface{}) {
	if L.Empty() || L.head == node {
		L.Push(value)
	} else {
		var prev *Node
		curr := L.head

		for curr != nil && curr != node {
			prev = curr
			curr = curr.next
		}

		if curr != nil {
			prev.next = &Node{
				value: value,
				next:  curr,
			}
		} else {
			L.Append(value)
		}
	}
}

func (L *List) InsertAfter(node *Node, value interface{}) {
	newNode := &Node{
		value: value,
		next:  node.next,
	}
	node.next = newNode

	if L.tail.value == node.value {
		L.tail = newNode
	}
}

func (L *List) Append(value interface{}) {
	newNode := &Node{
		value: value,
		next:  nil,
	}
	if L.tail == nil {
		L.head = newNode
		L.tail = newNode
	} else {
		L.tail.next = newNode
		L.tail = newNode
	}
}

func (L *List) DeleteLast() error {
	if L.Empty() {
		return errors.New("can't delete node from empty list")
	}

	// If there is only one node remove it
	if L.head.value == L.tail.value {
		L.head = nil
		L.tail = nil
		return nil
	}

	currentHead := L.head
	for currentHead.next.next != nil {
		currentHead = currentHead.next
	}

	L.tail = currentHead
	currentHead.next = nil
	return nil
}

func (L *List) DeleteFirst() error {
	if L.Empty() {
		return errors.New("can't delete node from empty list")
	}

	L.head = L.head.next
	if L.head == nil {
		L.tail = nil
	}
	return nil
}

func (L *List) Delete(value interface{}) error {
	if L.Empty() {
		return errors.New("can't delete node from empty list")
	}

	curr := L.head

	for curr.next != nil {
		if curr.next.value == value {
			// what we are deleting
			// curr.next -> curr.next.next
			curr.next = curr.next.next
			return nil
		}
		curr = curr.next
	}
	return errors.New("can't find given node")
}

func (L *List) Empty() bool {
	return L.head == nil
}

func (L *List) Display() {
	list := L.head

	for list != nil {
		fmt.Printf("%v ", list.value)
		list = list.next
	}
	fmt.Println()
}
