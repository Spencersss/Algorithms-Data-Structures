package stack

type Stack[T any] struct {
	items []T
}

// New creates a new stack with an initial capacity of n int values.
func New[T any](n int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, n),
	}
}

func (s *Stack[T]) Length() int {
	return len(s.items)
}

func (s *Stack[T]) Pop() T {
	poppedItem := s.items[0]
	s.items = s.items[1:]

	return poppedItem
}

func (s *Stack[T]) Add(value T) {
	s.items = append([]T{value}, s.items...)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack[T]) Top() T {
	return s.items[0]
}
