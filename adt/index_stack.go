package adt

type IndexStack struct {
	table map[interface{}]bool
	Stack
}

func NewIndexStack(baseStack Stack) *IndexStack {
	return &IndexStack{
		table: make(map[interface{}]bool),
		Stack: baseStack,
	}
}

func (is *IndexStack) PushIf(value interface{}) {
	if _, ok := is.table[value]; ok == false {
		is.table[value] = true
		is.Push(value)
	}
}
