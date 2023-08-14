package avl_tree

import (
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	leftChild  *Node[T]
	rightChild *Node[T]
	Data       T
	height     int
}

func NewNode[T constraints.Ordered](data T) *Node[T] {
	return &Node[T]{
		height: 1,
		Data:   data,
	}
}

// AVLTree is an Adelson-Velsky and Landis self-balancing binary search tree.
type AVLTree[T constraints.Ordered] struct {
	Root *Node[T]
}

func NewAVLTree[T constraints.Ordered]() AVLTree[T] {
	return AVLTree[T]{}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func height[T constraints.Ordered](node *Node[T]) int {
	if node == nil {
		return 0
	}

	return node.height
}

// balanceFactor shows the difference between the left and right
// subtrees of Node n. A difference of more than one would mean
// the balance() function is used to balance the tree
func (n *Node[T]) balanceFactor() int {
	if n == nil {
		return 0
	}

	return height(n.leftChild) - height(n.rightChild)
}

func rotateRight[T constraints.Ordered](node *Node[T]) *Node[T] {
	x := node.leftChild
	tmp := x.rightChild

	x.rightChild = node
	node.leftChild = tmp

	node.height = max(
		height(node.leftChild),
		height(node.rightChild),
	) + 1
	x.height = max(
		height(x.leftChild),
		height(x.rightChild),
	) + 1
	return x
}

func rotateLeft[T constraints.Ordered](node *Node[T]) *Node[T] {
	y := node.rightChild
	tmp := y.leftChild

	y.leftChild = node
	node.rightChild = tmp

	node.height = max(
		height(node.leftChild),
		height(node.rightChild),
	) + 1

	y.height = max(
		height(y.leftChild),
		height(y.rightChild),
	) + 1

	return y
}

func insert[T constraints.Ordered](node *Node[T], data T) *Node[T] {
	if node == nil {
		node = NewNode[T](data)
		return node
	}

	if data < node.Data {
		node.leftChild = insert(node.leftChild, data)
	} else if data > node.Data {
		node.rightChild = insert(node.rightChild, data)
	} else {
		return node
	}

	node.height = 1 + max(height(node.leftChild), height(node.rightChild))

	bf := node.balanceFactor()

	if bf > 1 && data < node.leftChild.Data {
		return rotateRight(node)
	}

	if bf < -1 && data > node.rightChild.Data {
		return rotateLeft(node)
	}

	if bf > 1 && data > node.leftChild.Data {
		node.leftChild = rotateLeft(node.leftChild)
		return rotateRight(node)
	}

	if bf < -1 && data < node.rightChild.Data {
		node.rightChild = rotateRight(node.rightChild)
		return rotateLeft(node)
	}

	return node
}

// Add an item to the AVL Tree
func (t *AVLTree[T]) Add(data T) {
	t.Root = insert(t.Root, data)
}

func traverseInOrder[T constraints.Ordered](node *Node[T], data *[]T) {
	if node != nil {
		traverseInOrder(node.leftChild, data)
		*data = append(*data, node.Data)
		traverseInOrder(node.rightChild, data)
	}
}

// InOrder tree traversal from Root
func (t *AVLTree[T]) InOrder() []T {
	data := []T{}
	traverseInOrder(t.Root, &data)
	return data
}
