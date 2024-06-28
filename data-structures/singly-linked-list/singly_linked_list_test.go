package main

import "testing"

func TestSuccessAppendSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(25)

	expectedLength := 1
	expectedData := 25

	if expectedLength != list.Length() {
		t.Errorf("Expected length of %d but got %d.", expectedLength, list.Length())
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

	expectedLength := 2
	expectedData := 20

	if expectedLength != list.Length() {
		t.Errorf("Expected length of %d but got %d.", expectedLength, list.Length())
	}

	if list.Head.Data != expectedData {
		t.Errorf("Expected head node data to be value of %d but got %d.", expectedData, list.Head.Data)
	}

	if list.Head.Next == nil {
		t.Error("Expected head next to be node with value of 10.")
	}
}

func TestSuccessPopSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(56)

	result := list.Pop()
	expectedLength := 0
	expectedData := 56

	if expectedLength != list.Length() {
		t.Errorf("Expected length of %d but got %d.", expectedLength, list.Length())
	}

	if result.Data != expectedData {
		t.Errorf("Expected head node data to be value of %d but got %d.", expectedData, list.Head.Data)
	}
}
