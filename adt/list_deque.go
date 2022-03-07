package adt

type ListDeque[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (q *ListDeque[T]) Empty() bool {
	return q.head == nil || q.tail == nil
}

func (q *ListDeque[T]) Pull(value T) {
	q.Push(value)
}

func (q *ListDeque[T]) Get() T {
	value := q.head.value
	q.head = q.head.next

	return value
}

func (q *ListDeque[T]) Push(value T) {
	if q.head == nil {
		q.head = &Node[T]{value: value}
		q.tail = q.head
	} else {
		newItem := &Node[T]{value: value}
		q.tail.next = newItem
		newItem.prev = q.tail
		q.tail = newItem
	}
}

func (q *ListDeque[T]) Pop() T {
	value := q.tail.value
	q.tail = q.tail.prev
	return value
}
