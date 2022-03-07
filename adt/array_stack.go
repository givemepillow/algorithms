package adt

type ArrayStack[T any] struct {
	number int
	array  [1000]T
}

func (s *ArrayStack[T]) Empty() bool {
	return s.number == 0
}

func (s *ArrayStack[T]) Pop() T {
	s.number -= 1
	return s.array[s.number]
}

func (s *ArrayStack[T]) Push(value T) {
	s.array[s.number] = value
	s.number += 1
}

func (s *ArrayStack[T]) Size() int {
	return s.number
}
