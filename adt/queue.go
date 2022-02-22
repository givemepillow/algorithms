package adt

type Queue interface {
	Get() interface{}
	Pull(value interface{})
	Empty() bool
}
