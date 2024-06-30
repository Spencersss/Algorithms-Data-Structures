package stack

type Stack struct {
	items []int
}

// New creates a new stack with an initial capacity of n int values.
func New(n int) *Stack {
	return &Stack{
		items: make([]int, n),
	}
}

func (s *Stack) Length() int {
	return len(s.items)
}

func (s *Stack) Pop() int {
	poppedItem := s.items[0]
	s.items = s.items[1:]

	return poppedItem
}

func (s *Stack) Add(value int) {
	s.items = append([]int{value}, s.items...)
}

func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}
