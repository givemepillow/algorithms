package adt

type ArrayStack struct {
	number int
	array  [1000]interface{}
}

func (s *ArrayStack) Empty() bool {
	return s.number == 0
}

func (s *ArrayStack) Pop() interface{} {
	s.number -= 1
	return s.array[s.number]
}

func (s *ArrayStack) Push(value interface{}) {
	s.array[s.number] = value
	s.number += 1
}

func (s *ArrayStack) Size() int {
	return s.number
}
