package adt

type Stack interface {
	Push(value interface{})
	Pop() interface{}
	Empty() bool
}
