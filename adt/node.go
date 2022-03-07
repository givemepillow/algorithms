package adt

type Node[T any] struct {
	value      T
	prev, next *Node[T]
}
