package rt

import (
	"ads/adt"
	"fmt"
)

type BinaryNode[V any] struct {
	Value V
	Left  *BinaryNode[V]
	Right *BinaryNode[V]
}

type BinaryTree[V any] struct {
	Root *BinaryNode[V]
}

func recursiveTraverseCount[V any](node *BinaryNode[V]) int {
	if node != nil {
		return 1 + recursiveTraverseCount(node.Left) + recursiveTraverseCount(node.Right)
	} else {
		return 0
	}
}

func (tree *BinaryTree[V]) RecursiveTraverseCount() int {
	node := tree.Root
	return recursiveTraverseCount(node)
}

func height[V any](node *BinaryNode[V]) int {
	if node != nil {
		l := height(node.Left)
		r := height(node.Right)
		if l > r {
			return l + 1
		} else {
			return r + 1
		}
	}
	return -1
}

func (tree *BinaryTree[V]) Height() int {
	node := tree.Root
	return height(node)
}

func recursiveTraverse[V any](node *BinaryNode[V]) {
	if node != nil {
		fmt.Println(node.Value)
		recursiveTraverse(node.Left)
		recursiveTraverse(node.Right)
	}
}

func (tree *BinaryTree[V]) RecursiveTraverse() {
	node := tree.Root
	recursiveTraverse(node)
}

func (tree *BinaryTree[V]) StackTraverse() {
	node := tree.Root
	stack := adt.ListStack[*BinaryNode[V]]{}
	stack.Push(node)
	for !stack.Empty() {
		var node = stack.Pop()
		fmt.Println(node.Value)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

func (tree *BinaryTree[V]) LevelTraverse() {
	node := tree.Root
	queue := adt.ListQueue[*BinaryNode[V]]{}
	queue.Pull(node)
	for !queue.Empty() {
		node := queue.Get()
		fmt.Println(node.Value)
		if node.Left != nil {
			queue.Pull(node.Left)
		}
		if node.Right != nil {
			queue.Pull(node.Right)
		}
	}
}
