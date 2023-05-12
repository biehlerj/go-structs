package doublylinkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tail = &Node{
	value: 400,
	next:  nil,
	prev:  nil,
}

// emptyList creates an empty Doubly Linked List
func emptyList() *List {
	return &List{
		head: nil,
		tail: nil,
	}
}

// oneNodeList creates a list with one node
func oneNodeList() *List {
	return &List{
		head: tail,
		tail: tail,
	}
}

func TestDoublyLinkedList_Empty(t *testing.T) {
	t.Run("Empty() returns true on an empty list", func(t *testing.T) {
		list := emptyList()
		assert.Equal(t, true, list.Empty())
	})
	t.Run("Empty() returns false on a non-empty list", func(t *testing.T) {
		list := oneNodeList()
		assert.Equal(t, false, list.Empty())
	})
}

func TestDoublyLinkedList_Push(t *testing.T) {
	t.Run("push a node onto an empty list", func(t *testing.T) {
		list := emptyList()
		value := 7
		list.Push(value)
		assert.Equal(t, value, list.head.value)
		assert.Equal(t, value, list.tail.value)
	})
	t.Run("push node onto non-empty list", func(t *testing.T) {
		list := oneNodeList()
		value := 1
		list.Push(value)
		assert.Equal(t, value, list.head.value)
		assert.NotEqual(t, value, list.tail.value)
	})
}

func TestDoublyLinkedList_Append(t *testing.T) {
	t.Run("append node to an empty list", func(t *testing.T) {
		list := emptyList()
		value := 6
		list.Append(value)
		assert.Equal(t, value, list.head.value)
		assert.Equal(t, value, list.tail.value)
	})
	t.Run("append a node to a non-empty list", func(t *testing.T) {
		list := oneNodeList()
		value := 33
		list.Append(value)
		assert.NotEqual(t, value, list.head.value)
		assert.Equal(t, value, list.tail.value)
	})
}

func TestDoublyLinkedList_InsertBefore(t *testing.T) {
	t.Run("insert new node before first node", func(t *testing.T) {
		list := oneNodeList()
		value := 5
		list.InsertBefore(tail, value)
		assert.Equal(t, value, list.head.value)
		assert.Equal(t, tail.value, list.tail.value)
	})
}
