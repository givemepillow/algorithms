package adt

type ArrayQueue struct {
	head, tail, size int
	array            []Node
}

func NewArrayQueue(size int) *ArrayQueue {
	return &ArrayQueue{size: size, array: make([]Node, size)}
}

func (q *ArrayQueue) Empty() bool {
	return q.head%q.size == q.tail
}

func (q *ArrayQueue) Pull(value interface{}) {
	q.array[q.tail] = Node{value: value}
	q.tail = (q.tail + 1) % q.size
}

func (q *ArrayQueue) Get() interface{} {
	result := q.array[q.head]
	q.head = (q.head + 1) % q.size
	return result.value
}
