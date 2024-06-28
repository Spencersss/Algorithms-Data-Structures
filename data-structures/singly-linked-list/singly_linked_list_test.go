package main

import "testing"

func TestSuccessAppendSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()

	list.Append(25)

	resultLength := list.Length()
	expectedLength := 1
	expectedData := 25

	if expectedLength != resultLength {
		t.Errorf("Expected length of %d but got %d.", expectedLength, resultLength)
	}

	if list.Head.Data != expectedData {
		t.Errorf("Expected head node data to be value of %d but got %d.", expectedData, list.Head.Data)
	}

	if list.Head.Next != nil {
		t.Error("Expected head next to be nil but found node.")
	}
}

func TestSuccessGetLengthSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(32)
	list.Append(6)

	result := list.Length()
	expected := 2

	if expected != result {
		t.Errorf("Expected length of %d but got %d", expected, result)
	}
}

func TestSuccessInsertBeforeHeadSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(10)
	list.InsertBeforeHead(20)

	resultLength := list.Length()
	expectedLength := 2
	expectedData := 20

	if expectedLength != resultLength {
		t.Errorf("Expected length of %d but got %d.", expectedLength, resultLength)
	}

	if list.Head.Data != expectedData {
		t.Errorf("Expected head node data to be value of %d but got %d.", expectedData, list.Head.Data)
	}

	if list.Head.Next == nil {
		t.Error("Expected head next to be node with value of 10.")
	}
}
