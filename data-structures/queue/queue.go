package queue

type Queue[T any] struct {
	Items []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.Items = append(q.Items, value)
}

func (q *Queue[T]) EnqueueAll(values ...T) {
	for _, value := range values {
		q.Enqueue(value)
	}
}

func (q *Queue[T]) Dequeue() T {
	item := q.Items[0]
	q.Items = q.Items[1:]

	return item
}

func (q *Queue[T]) Length() int {
	return len(q.Items)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Length() == 0
}
