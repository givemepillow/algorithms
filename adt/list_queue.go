package adt

type ListQueue[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (q *ListQueue[T]) Empty() bool {
	return q.head == nil
}

func (q *ListQueue[T]) Pull(value T) {
	if q.head == nil {
		q.head = &Node[T]{value: value}
		q.tail = q.head
	} else {
		q.tail.next = &Node[T]{value: value}
		q.tail = q.tail.next
	}
}

func (q *ListQueue[T]) Get() T {
	value := q.head.value
	q.head = q.head.next
	return value
}
