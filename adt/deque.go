package adt

type Deque interface {
	Get() interface{}
	Put(value interface{})
	Push(value interface{})
	Pop() interface{}
	Empty() bool
}
