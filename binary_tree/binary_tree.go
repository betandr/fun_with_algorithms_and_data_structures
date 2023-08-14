package binary_tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (n *Node[T]) GreaterThan(node *Node[T]) bool {
	return n.Value > node.Value
}

func (n *Node[T]) LessThan(node *Node[T]) bool {
	return n.Value < node.Value
}

type BinaryTree[T constraints.Ordered] struct {
	Root    *Node[T]
	Counter int
}

func new[T constraints.Ordered]() BinaryTree[T] {
	return BinaryTree[T]{}
}

// addTo inserts the item at the correct position given it's
// either less than or greater than its parent.
func (bt *BinaryTree[T]) addTo(node *Node[T], val T) {
	nd := Node[T]{
		Value: val,
	}

	if nd.LessThan(node) {
		if node.Left == nil {
			node.Left = &nd
		} else {
			bt.addTo(node.Left, val)
		}
	} else {
		if node.Right == nil {
			node.Right = &nd
		} else {
			bt.addTo(node.Right, val)
		}
	}

}

// Add an item to the binary tree.
func (bt *BinaryTree[T]) Add(o T) {
	if bt.Root == nil {
		bt.Root = &Node[T]{
			Value: o,
		}
	} else {
		bt.addTo(bt.Root, o)
	}
	bt.Counter++
}

func (bt *BinaryTree[T]) Contains(val T) bool {
	curr, _ := bt.findFromParent(bt.Root, val)
	return curr != nil
}

func (bt *BinaryTree[T]) Remove(val T) error {
	node, parent := bt.findFromParent(bt.Root, val)

	if node == nil {
		return fmt.Errorf("%v does not exist in tree", val)
	}

	// Case 1: Node has no right child so node's left replaces node
	if node.Right == nil {
		if parent == nil {
			bt.Root = node.Left
		} else {
			if parent.GreaterThan(node) {
				parent.Left = node.Left
			} else if parent.LessThan(node) {
				parent.Right = node.Left
			}
		}

		// Case 2: Node's right child has no left child so node's right child replaces node
	} else if node.Right.Left == nil {
		node.Right.Left = node.Left
		if parent == nil {
			bt.Root = node.Right
		} else {
			if parent.GreaterThan(node) {
				parent.Left = node.Right
			} else if parent.LessThan(node) {
				parent.Right = node.Right
			}
		}

		// Case 3 Node's right child has a left child so node's right child's left most child replaces node
	} else {
		leftMost := node.Right.Left
		leftMostParent := node.Right

		for leftMost.Left != nil {
			leftMostParent = leftMost
			leftMost = leftMost.Left
		}

		leftMostParent.Left = leftMost.Right

		leftMost.Left = node.Left
		leftMost.Right = node.Right

		if parent == nil {
			bt.Root = leftMost
		} else {
			if parent.GreaterThan(node) {
				parent.Left = leftMost
			} else if parent.LessThan(node) {
				parent.Right = leftMost
			}
		}
	}

	bt.Counter--

	return nil
}

func (bt *BinaryTree[T]) findFromParent(node *Node[T], val T) (*Node[T], *Node[T]) {
	var parent *Node[T]
	find := Node[T]{Value: val}
	curr := bt.Root

	for curr != nil {
		if find.LessThan(curr) {
			parent = curr
			curr = curr.Left
		} else if find.GreaterThan(curr) {
			parent = curr
			curr = curr.Right
		} else {
			break
		}
	}

	return curr, parent
}

func (bt *BinaryTree[T]) postOrderTraversal(node *Node[T], vals *[]T) {
	var zero *Node[T]

	if node != zero {
		bt.postOrderTraversal(node.Left, vals)
		bt.postOrderTraversal(node.Right, vals)
		*vals = append(*vals, node.Value)
	}
}

// ValuesPostOrder returns the values from the tree after performing an
// postorder traversal:
// 1. Traverse the left subtree
// 2. Traverse the right subtree
// 3. Obtain the value
func (bt *BinaryTree[T]) ValuesPostOrder(node *Node[T]) []T {
	values := []T{}
	bt.postOrderTraversal(node, &values)
	return values
}

func (bt *BinaryTree[T]) preOrderTraversal(node *Node[T], vals *[]T) {
	var zero *Node[T]

	if node != zero {
		*vals = append(*vals, node.Value)
		bt.preOrderTraversal(node.Left, vals)
		bt.preOrderTraversal(node.Right, vals)
	}
}

// ValuesPreOrder returns the values from the tree after performing an
// preorder traversal:
// 1. Obtain the value
// 2. Traverse the left subtree
// 3. Traverse the right subtree
func (bt *BinaryTree[T]) ValuesPreOrder(node *Node[T]) []T {
	values := []T{}
	bt.preOrderTraversal(node, &values)
	return values
}

func (bt *BinaryTree[T]) inOrderTraversal(node *Node[T], vals *[]T) {
	// var zero *Node[T]

	if node != nil {
		bt.inOrderTraversal(node.Left, vals)
		*vals = append(*vals, node.Value)
		bt.inOrderTraversal(node.Right, vals)
	}
}

// ValuesInOrder returns the values from the tree after performing an
// inorder traversal:
// 1. Traverse the left subtree
// 2. Obtain the value
// 3. Traverse the right subtree
func (bt *BinaryTree[T]) ValuesInOrder(node *Node[T]) []T {
	values := []T{}
	bt.inOrderTraversal(node, &values)
	return values
}
