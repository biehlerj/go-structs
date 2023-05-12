// Package doublylinkedlist implements a Doubly Linked List
package doublylinkedlists

// A Node is a singular item in a linked list
type Node struct {
	value      any
	next, prev *Node
}

// A List is constructed of Nodes
// A List at the minimum has a head and tail Node
// The head and tail Node can be the same Node or nil
type List struct {
	head, tail *Node
}

// empty checks if th eList has any Nodes
func (L *List) Empty() bool {
	return L.head == nil
}

// Push adds a new Node at the beginning of a List
// It takes a value which can be any type
func (L *List) Push(value any) {
	newNode := &Node{
		value: value,
		next:  L.head,
	}

	if L.head == nil {
		L.head = newNode
		L.tail = newNode
	}
	L.head.prev = newNode
	L.head = newNode
}

// Append adds a new Node to the end of the List
// It requires a value which can be any type
func (L *List) Append(value any) {
	newNode := &Node{
		value: value,
		next:  nil,
		prev:  L.tail,
	}

	if L.head == nil {
		L.tail = newNode
		L.head = newNode
	}
	L.tail.next = newNode
	L.tail = newNode
}

// InsertBefore adds a Node to the List before the given Node
// It requires the Node to insert before and a value
// The value can be any type
func (L *List) InsertBefore(node *Node, value any) {
	if L.Empty() || L.head == node {
		L.Push(value)
	}

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
			prev:  prev,
		}
	}

	L.Append(value)
}
