package list

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

// DeleteAt finds the existing node in the list at the provided n index and attempts to remove it from the list
// If the node does not exist, a no-op will occur and the list will remain as is. If a node is removed, it will be returned
// else a nil will be returned.
func (l *SinglyLinkedList) DeleteAt(n int) *Node {
	if l.Head != nil || n != 0 {
		var lastNode *Node
		var curNode *Node = l.Head

		curIndex := 0
		for curIndex < n {
			lastNode = curNode
			if curNode == nil {
				return nil
			}
			curNode = curNode.Next
			curIndex++
		}

		curNode.Next = nil
		lastNode.Next = curNode.Next

		return curNode
	}

	return nil
}

// DeleteTail removes the last node in the list and returns the node upon removal. If not found, the function will return nil
func (l *SinglyLinkedList) DeleteTail() *Node {
	if l.Head != nil {
		var lastNode *Node
		var curNode *Node = l.Head
		for curNode.Next != nil {
			lastNode = curNode
			curNode = curNode.Next
		}

		curNode.Next = nil
		lastNode.Next = nil

		return curNode
	}

	return nil
}

// Find will search the list for the first node it comes across containing the provided data.
// If found it will return the index of the node with the value. If not found, this will return -1.
func (l *SinglyLinkedList) Find(value int) int {
	curIndex := 0
	curNode := l.Head
	for curNode != nil {
		if curNode.Data == value {
			return curIndex
		}

		curIndex++
		curNode = curNode.Next
	}

	return -1
}

// FindAt will search the list for a node at the provided n index.
// If found it will return a pointer to the node. If not found, this will return nil.
func (l *SinglyLinkedList) FindAt(n int) *Node {
	curIndex := 0
	curNode := l.Head
	for curNode != nil {
		if curIndex == n {
			return curNode
		}

		curIndex++
		curNode = curNode.Next
	}

	return nil
}
