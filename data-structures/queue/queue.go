package queue

type Queue[T any] struct {
	items []T
}

func New[T any](n int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, n),
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

func (q *Queue[T]) EnqueueAll(values ...T) {
	for _, value := range values {
		q.Enqueue(value)
	}
}

func (q *Queue[T]) Dequeue() T {
	item := q.items[0]
	q.items = q.items[1:]

	return item
}

func (q *Queue[T]) Length() int {
	return len(q.items)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue[T]) Front() T {
	return q.items[0]
}
