package main

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
}

func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{}
}

func (l *SinglyLinkedList) Length() int {
	size := 0
	curNode := l.Head
	for curNode != nil {
		size++
		curNode = curNode.Next
	}

	return size
}

func (l *SinglyLinkedList) Append(value int) {
	newNode := &Node{
		Data: value,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = newNode
	} else {
		curNode := l.Head
		for curNode.Next != nil {
			curNode = curNode.Next
		}

		curNode.Next = newNode
	}
}

func (l *SinglyLinkedList) InsertBeforeHead(value int) {
	newNode := &Node{
		Data: value,
		Next: nil,
	}

	if l.Head != nil {
		newNode.Next = l.Head
	}

	l.Head = newNode
}
