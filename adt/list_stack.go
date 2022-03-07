package adt

type ListStack[T any] struct {
	item *Node[T]
}

func (s *ListStack[T]) Push(value T) {
	s.item = &Node[T]{value: value, prev: s.item}
}

func (s *ListStack[T]) Pop() T {
	result := s.item.value
	s.item = s.item.prev
	return result
}

func (s *ListStack[T]) Empty() bool {
	return s.item == nil
}
