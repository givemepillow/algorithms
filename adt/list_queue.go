package adt

type ListQueue struct {
	head *Node
	tail *Node
}

func (q *ListQueue) Empty() bool {
	return q.head == nil
}

func (q *ListQueue) Pull(value interface{}) {
	if q.head == nil {
		q.head = &Node{value: value}
		q.tail = q.head
	} else {
		q.tail.next = &Node{value: value}
		q.tail = q.tail.next
	}
}

func (q *ListQueue) Get() interface{} {
	value := q.head.value
	q.head = q.head.next
	return value
}
