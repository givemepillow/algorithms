package adt

type ArrayDeque struct {
	size, head, tail int
	array            []Node
}

func NewArrayDeque(size int) *ArrayDeque {
	return &ArrayDeque{size: size, array: make([]Node, size)}
}

func (q *ArrayDeque) Empty() bool {
	return q.head%q.size == q.tail
}

func (q *ArrayDeque) Pull(value interface{}) {
	q.array[q.tail] = Node{value: value}
	q.tail = (q.tail + 1) % q.size
}

func (q *ArrayDeque) Get() interface{} {
	result := q.array[q.head]
	q.head = (q.head + 1) % q.size
	return result.value
}

func (q *ArrayDeque) Push(value interface{}) {
	q.array[q.tail] = Node{value: value}
	q.tail = (q.tail + 1) % q.size
}

func (q *ArrayDeque) Pop() interface{} {
	q.tail = (q.tail - 1) % q.size
	return q.array[q.tail].value
}
