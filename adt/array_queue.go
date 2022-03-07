package adt

type ArrayQueue[T any] struct {
	head, tail, size int
	array            []Node[T]
}

func (q *ArrayQueue[T]) Empty() bool {
	return q.head%q.size == q.tail
}

func (q *ArrayQueue[T]) Pull(value T) {
	q.array[q.tail] = Node[T]{value: value}
	q.tail = (q.tail + 1) % q.size
}

func (q *ArrayQueue[T]) Get() T {
	result := q.array[q.head]
	q.head = (q.head + 1) % q.size
	return result.value
}
