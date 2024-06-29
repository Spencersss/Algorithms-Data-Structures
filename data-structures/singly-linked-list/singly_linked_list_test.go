package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessGetLengthSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(32)
	list.Append(6)

	result := list.Length()
	expected := 2

	assert.Equalf(t, expected, result, "Expected length of expected but got %d", result)
}

func TestSuccessAppendSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(25)

	expectedLength := 1
	expectedData := 25

	assert.Equalf(t, expectedLength, list.Length(), "Expected length of 1 but found %d.", list.Length())
	assert.Equalf(t, expectedData, list.Head.Data, "Expected head node data to be value of %d but found %d.", expectedData, list.Head.Data)
	assert.Nil(t, list.Head.Next, "Expected head next to be nil but found node.")
}

func TestSuccessInsertBeforeHeadSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(10)
	list.InsertBeforeHead(20)

	expectedLength := 2
	expectedData := 20

	assert.Equalf(t, expectedLength, list.Length(), "Expected length of %d but found %d.", expectedLength, list.Length())
	assert.Equalf(t, expectedData, list.Head.Data, "Expected head node data to be value of 56 but found %d.", list.Head.Data)
	assert.NotNil(t, list.Head.Next, "Expected head next to be node with value of 10.")
}

func TestSuccessInsertAtSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(11)
	list.Append(33)

	list.InsertAt(22, 1)
	addedNode := list.Head.Next
	expectedData := 22
	expectedLength := 3

	assert.Equalf(t, expectedLength, list.Length(), "Expected length of 3 but got %d.", list.Length())
	assert.Equalf(t, expectedData, addedNode.Data, "Expected node value of %d but got %d.", expectedData, addedNode.Data)
	assert.NotNil(t, addedNode.Next, "Added node should have a next node but found nil.")
	assert.Equal(t, 33, addedNode.Next.Data, "Expected node after added node to have value of %d but found %d.", 33, addedNode.Next.Data)
}

func TestSuccessPopSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(56)

	result := list.Pop()
	expectedLength := 0
	expectedData := 56

	assert.Equalf(t, expectedLength, list.Length(), "Expected length of %d but found %d.", expectedLength, list.Length())
	assert.Equalf(t, expectedData, result.Data, "Expected head node data to be value of 56 but found %d.", result.Data)
}

func TestSuccessDeleteAtSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(50)
	list.Append(100)
	list.Append(150)

	result := list.DeleteAt(2)
	expectedLength := 2
	expectedData := 150

	assert.Equalf(t, expectedLength, list.Length(), "Expected length of %d but found %d.", expectedLength, list.Length())
	assert.Equalf(t, expectedData, result.Data, "Expected returned deleted node to have value of %d but found %d.", result.Data, expectedData)
}

func TestSuccessDeleteTailSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(300)
	list.Append(515)
	list.Append(902)

	result := list.DeleteTail()
	expectedLength := 2
	expectedData := 902

	assert.Equalf(t, expectedLength, list.Length(), "Expected length of %d but found %d.", expectedLength, list.Length())
	assert.Equalf(t, expectedData, result.Data, "Expected returned deleted node to have value of %d but found %d.", result.Data, expectedData)
}
