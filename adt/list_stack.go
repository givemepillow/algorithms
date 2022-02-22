package adt

type ListStack struct {
	item *Node
}

func (s *ListStack) Push(value interface{}) {
	s.item = &Node{value: value, prev: s.item}
}

func (s *ListStack) Pop() interface{} {
	result := s.item.value
	s.item = s.item.prev
	return result
}

func (s *ListStack) Empty() bool {
	return s.item == nil
}
