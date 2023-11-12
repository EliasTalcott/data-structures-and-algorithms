package binarySearchTree

import (
	"etalcott/datastructures/queue"
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func (root *Node[T]) Insert(value T) *Node[T] {
	// O(logn) for balanced tree, O(n) worst case
	if root == nil {
		return &Node[T]{value, nil, nil}
	} else if value < root.value {
		root.left = root.left.Insert(value)
	} else if value > root.value {
		root.right = root.right.Insert(value)
	}
	return root
}

func (root *Node[T]) TraverseLevelOrder() []T {
	// O(n)
	queue := queue.CreateQueue([]*Node[T]{root})
	var result []T

	for !queue.Empty() {
		node, _ := queue.Dequeue()
		result = append(result, node.value)
		if node.left != nil {
			queue.Enqueue(node.left)
		}
		if node.right != nil {
			queue.Enqueue(node.right)
		}
	}
	return result
}

func (root *Node[T]) TraverseInOrder() []T {
	// O(n)
	var result []T
	if root == nil {
		return result
	}
	result = append(result, root.left.TraverseInOrder()...)
	result = append(result, root.value)
	result = append(result, root.right.TraverseInOrder()...)
	return result
}

func (root *Node[T]) TraversePreOrder() []T {
	// O(n)
	var result []T
	if root == nil {
		return result
	}
	result = append(result, root.value)
	result = append(result, root.left.TraversePreOrder()...)
	result = append(result, root.right.TraversePreOrder()...)
	return result
}

func (root *Node[T]) TraversePostOrder() []T {
	// O(n)
	var result []T
	if root == nil {
		return result
	}
	result = append(result, root.left.TraversePostOrder()...)
	result = append(result, root.right.TraversePostOrder()...)
	result = append(result, root.value)
	return result
}
