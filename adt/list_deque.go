package adt

type ListDeque struct {
	head *Node
	tail *Node
}

func (q *ListDeque) Empty() bool {
	return q.head == nil || q.tail == nil
}

func (q *ListDeque) Pull(value interface{}) {
	q.Push(value)
}

func (q *ListDeque) Get() interface{} {
	value := q.head.value
	q.head = q.head.next

	return value
}

func (q *ListDeque) Push(value interface{}) {
	if q.head == nil {
		q.head = &Node{value: value}
		q.tail = q.head
	} else {
		newItem := &Node{value: value}
		q.tail.next = newItem
		newItem.prev = q.tail
		q.tail = newItem
	}
}

func (q *ListDeque) Pop() interface{} {
	value := q.tail.value
	q.tail = q.tail.prev
	return value
}
