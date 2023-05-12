package singly_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList_Push(t *testing.T) {
	t.Run("push node onto empty list", func(t *testing.T) {
		list := emptyList()
		value := 7
		list.Push(value)
		assert.Equal(t, list.head.value, value)
		assert.Equal(t, list.tail.value, value)
	})
	t.Run("push node on non-empty list", func(t *testing.T) {
		list := oneNodeList()
		value := 1
		list.Push(value)
		assert.Equal(t, list.head.value, value)
		assert.NotEqual(t, list.tail.value, value)
	})
}

func TestSinglyLinkedList_DeleteFirst(t *testing.T) {
	t.Run("remove first node from empty list", func(t *testing.T) {
		list := emptyList()
		err := list.DeleteFirst()
		assert.Error(t, err)
	})
	t.Run("remove first node from one node list", func(t *testing.T) {
		list := oneNodeList()
		err := list.DeleteFirst()
		assert.NoError(t, err)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})
	t.Run("remove first node from multi node list", func(t *testing.T) {
		list := multiNodeList()
		err := list.DeleteFirst()
		assert.NoError(t, err)
		assert.NotNil(t, list.head)
		assert.Equal(t, list.head.value, second.value)
	})
}

func TestSinglyLinkedList_Append(t *testing.T) {
	t.Run("append node to empty list", func(t *testing.T) {
		list := emptyList()
		value := 6
		list.Append(value)
		assert.Equal(t, list.head.value, value)
		assert.Equal(t, list.tail.value, value)
	})
	t.Run("append node to non-empty list", func(t *testing.T) {
		list := oneNodeList()
		value := 33
		list.Append(value)
		assert.NotEqual(t, list.head.value, value)
		assert.Equal(t, list.tail.value, value)
	})
}

func TestSinglyLinkedList_DeleteLast(t *testing.T) {
	t.Run("delete node from end of empty list", func(t *testing.T) {
		list := emptyList()
		err := list.DeleteLast()
		assert.Error(t, err)
	})
	t.Run("delete node from one node list", func(t *testing.T) {
		list := oneNodeList()
		err := list.DeleteLast()
		assert.NoError(t, err)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})
	t.Run("delete node from multi node list", func(t *testing.T) {
		list := multiNodeList()
		err := list.DeleteLast()
		assert.NoError(t, err)
		assert.NotNil(t, list.head)
		assert.Equal(t, list.tail.value, tail.value)
	})
}

func TestSinglyLinkedList_InsertBefore(t *testing.T) {
	t.Run("insert new node before first node", func(t *testing.T) {
		list := oneNodeList()
		value := 5
		list.InsertBefore(tail, value)
		assert.Equal(t, list.head.value, value)
		assert.Equal(t, list.tail.value, tail.value)
	})
	t.Run("insert before node on list with multiple nodes", func(t *testing.T) {
		list := multiNodeList()
		value := 9
		list.InsertBefore(third, value)
		assert.Equal(t, second.next.value, value)
	})
	t.Run("insert before non-existent node", func(t *testing.T) {
		list := multiNodeList()
		value := 10
		list.InsertBefore(nil, value)
		assert.Equal(t, tail.next.value, value)
	})
}

func TestSinglyLinkedList_InsertAfter(t *testing.T) {
	t.Run("insert after node which is head and tail of list", func(t *testing.T) {
		list := oneNodeList()
		value := 8
		list.InsertAfter(tail, value)
		assert.Equal(t, list.head.value, tail.value)
		assert.Equal(t, list.tail.value, value)
	})
	t.Run("insert after node in middle of multi node list", func(t *testing.T) {
		list := multiNodeList()
		value := 8
		list.InsertAfter(second, value)
		assert.Equal(t, second.next.value, value)
	})
}

func TestSinglyLinkedList_Delete(t *testing.T) {
	t.Run("delete from empty list", func(t *testing.T) {
		list := emptyList()
		err := list.Delete(69)
		assert.Error(t, err)
	})
	t.Run("delete from list", func(t *testing.T) {
		list := multiNodeList()
		err := list.Delete(8)
		assert.Nil(t, err)
	})
	t.Run("delete non-existent node", func(t *testing.T) {
		list := multiNodeList()
		err := list.Delete(69)
		assert.Error(t, err)
	})
}

func TestSinglyLinkedList_Display(t *testing.T) {
	t.Run("displaying multi node linked list", func(t *testing.T) {
		list := multiNodeList()
		list.Display()
	})
	t.Run("displaying one node linked list", func(t *testing.T) {
		list := oneNodeList()
		list.Display()
	})
	t.Run("displaying empty linked list", func(t *testing.T) {
		list := emptyList()
		list.Display()
	})
}

func TestSinglyLinkedList_Empty(t *testing.T) {
	t.Run("Empty() returns true on an empty list", func(t *testing.T) {
		list := emptyList()
		assert.Equal(t, true, list.Empty())
	})
	t.Run("Empty() returns false on a non-empty list", func(t *testing.T) {
		list := oneNodeList()
		assert.Equal(t, false, list.Empty())
	})
}

var head = &Node{
	value: 100,
	next:  second,
}

var second = &Node{
	value: 200,
	next:  third,
}

var third = &Node{
	value: 300,
	next:  tail,
}

var tail = &Node{
	value: 400,
	next:  nil,
}

func emptyList() *List {
	return &List{
		head: nil,
		tail: nil,
	}
}

func oneNodeList() *List {
	return &List{
		head: tail,
		tail: tail,
	}
}

func multiNodeList() *List {
	return &List{
		head: head,
		tail: tail,
	}
}
