package main

type Node struct {
	Data int
	Next *Node
}

func NewNode(value int) *Node {
	return &Node{
		Data: value,
		Next: nil,
	}
}

type SinglyLinkedList struct {
	Head *Node
}

func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{
		Head: nil,
	}
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
	newNode := NewNode(value)

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
	newNode := NewNode(value)

	if l.Head != nil {
		newNode.Next = l.Head
	}

	l.Head = newNode
}

// InsertAt will insert a node with the provided value before the existing node at index value n.
// If there is an error inserting the new node, it will simply be a no-op and not perform the insertion.
func (l *SinglyLinkedList) InsertAt(value int, n int) {
	if l.Head == nil || n == 0 {
		l.InsertBeforeHead(value)
	} else {
		var lastNode *Node
		var curNode *Node = l.Head
		newNode := NewNode(value)

		curIndex := 0
		for curIndex < n {
			lastNode = curNode
			if curNode == nil {
				return
			}
			curNode = curNode.Next
			curIndex++
		}

		lastNode.Next = newNode
		newNode.Next = curNode
	}
}

// Pop deletes and returns the head node from the list.
func (l *SinglyLinkedList) Pop() *Node {
	if l.Head != nil {
		newHead := l.Head.Next
		l.Head.Next = nil
		oldHead := l.Head
		l.Head = newHead

		return oldHead
	}

	return nil
}
