package adt

type Node struct {
	value      interface{}
	prev, next *Node
}

type ListStack struct {
	item *Node
}

func (s *ListStack) Push(value interface{}) {
	if s.item == nil {
		s.item = new(Node)
		s.item.value = value
	} else {
		newItem := new(Node)
		newItem.value = value
		newItem.prev = s.item
		s.item.next = newItem
		s.item = newItem
	}
}

func (s *ListStack) Pop() interface{} {
	result := s.item.value
	if s.item.prev != nil {
		s.item = s.item.prev
		s.item.next = nil
	} else {
		s.item = nil
	}
	return result
}

func (s *ListStack) Empty() bool {
	return s.item == nil
}
