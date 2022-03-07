package adt

type ArrayDeque[T any] struct {
	size, head, tail int
	array            []Node[T]
}

func (q *ArrayDeque[T]) Empty() bool {
	return q.head%q.size == q.tail
}

func (q *ArrayDeque[T]) Pull(value T) {
	q.array[q.tail] = Node[T]{value: value}
	q.tail = (q.tail + 1) % q.size
}

func (q *ArrayDeque[T]) Get() T {
	result := q.array[q.head]
	q.head = (q.head + 1) % q.size
	return result.value
}

func (q *ArrayDeque[T]) Push(value T) {
	q.array[q.tail] = Node[T]{value: value}
	q.tail = (q.tail + 1) % q.size
}

func (q *ArrayDeque[T]) Pop() T {
	q.tail = (q.tail - 1) % q.size
	return q.array[q.tail].value
}
