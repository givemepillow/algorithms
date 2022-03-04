package rt

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

type BinaryTree struct {
	Root *BinaryNode
}

func recursiveTraverseCount(node *BinaryNode) int {
	if node != nil {
		return 1 + recursiveTraverseCount(node.Left) + recursiveTraverseCount(node.Right)
	} else {
		return 0
	}
}

func (tree *BinaryTree) RecursiveTraverseCount() int {
	node := tree.Root
	return recursiveTraverseCount(node)
}
