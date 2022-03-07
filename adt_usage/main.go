package main

import (
	"ads/adt"
	"fmt"
)

func main() {

	//indexStack := adt.NewIndexStack(adt.NewArrayDeque(100))
	//indexStack.PushIf(3)
	//indexStack.PushIf(3)
	//indexStack.PushIf("K")
	//indexStack.PushIf("K")
	//indexStack.PushIf("kirill")
	//
	//extract(indexStack)

	// deque := adt.NewArrayDeque(1000)
	//deque := adt.ListDeque{}
	//
	//deque.Push(1)
	//deque.Push(2)
	//deque.Push(3)
	//deque.Push(4)
	//
	//fmt.Println(deque.Pop())
	//fmt.Println(deque.Pop())
	//fmt.Println(deque.Get())
	//fmt.Println(deque.Get())

	//queue := adt.NewArrayQueue(1000)
	//// queue := new(adt.ListQueue)
	//queue.Pull(12)
	//queue.Pull("Hello!")
	//queue.Pull(12)
	//fmt.Println(queue.Get())
	//fmt.Println(queue.Get())
	//fmt.Println(queue.Get())
	//queue.Pull("Hello!")
	//fmt.Println(queue.Get())
	//fmt.Println(queue.Empty())

	//uf := adt.NewUF(10)
	//uf.Unite(1, 2)
	//uf.Unite(9, 2)
	//uf.Unite(9, 3)
	//uf.Unite(3, 5)
	//uf.Unite(4, 5)
	//uf.Unite(4, 7)
	//fmt.Println("Is united:", uf.Find(1, 7))
	//
	//stack := new(adt.ArrayStack[interface{}])
	//stack.Push(1)
	//stack.Push(2)
	//stack.Pop()
	//stack.Push("Kirill!")
	//stack.Pop()
	//stack.Pop()
	//stack.Push("SPB")
	//stack.Push("SUAI")
	//pull(stack)
	//extract(stack)
}

func extract(s adt.Stack) {
	for !s.Empty() {
		fmt.Println(s.Pop())
	}
}

func pull(s adt.Stack) {
	for i := 8; i < 100; i *= 2 {
		s.Push(i)
	}
}
