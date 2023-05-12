// Package singly_linked_list implements a Singly Linked List
package singly_linked_list

import (
	"errors"
	"fmt"
)

// A Node is a singular item in a linked list
type Node struct {
	value any
	next  *Node
}

// A List is constructed of Nodes
// A List at the minimum has a head and tail Node
// The head and tail Node can be the same Node or nil
type List struct {
	head *Node
	tail *Node
}

// Push adds a new Node at the beginning of a List
// It takes a value which can be any type
func (L *List) Push(value any) {
	newNode := &Node{
		value: value,
		next:  L.head,
	}
	L.head = newNode

	if L.tail == nil {
		L.tail = L.head
	}
}

// InsertBefore adds a Node to the List before the given Node
// It requires the Node to insert before and a value
// The value can be any type
func (L *List) InsertBefore(node *Node, value any) {
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

// InsertAfter adds a new Node after the given Node
// It requires the Node to insert after and a value
// The value can be any type
func (L *List) InsertAfter(node *Node, value any) {
	newNode := &Node{
		value: value,
		next:  node.next,
	}
	node.next = newNode

	if L.tail.value == node.value {
		L.tail = newNode
	}
}

// Append adds a new Node to the end of the List
// It requires a value which can be any type
func (L *List) Append(value any) {
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

// DeleteLast removes the last Node in the List
// It updates the tail value of the List
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

// DeleteFirst removes the first Node in the List
// It updates the head value of the List
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

// Delete removes the Node with the given value from the List
func (L *List) Delete(value any) error {
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

// Empty checks if the List has any Nodes
func (L *List) Empty() bool {
	return L.head == nil
}

// Displays the List
func (L *List) Display() {
	list := L.head

	for list != nil {
		fmt.Printf("%v ", list.value)
		list = list.next
	}
	fmt.Println()
}
