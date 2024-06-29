package queue

type Queue struct {
	Items []int
}

func (q *Queue) Enqueue(value int) {
	q.Items = append(q.Items, value)
}

func (q *Queue) Dequeue() int {
	item := q.Items[0]
	q.Items = q.Items[1:]

	return item
}

func (q *Queue) Length() int {
	return len(q.Items)
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}
