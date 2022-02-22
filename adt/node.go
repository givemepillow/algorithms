package adt

type Node struct {
	value      interface{}
	prev, next *Node
}
