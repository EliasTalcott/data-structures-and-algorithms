package binarySearchTree

import (
	"cmp"
	"etalcott/datastructures/queue"
)

type Node[T cmp.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func (root *Node[T]) Delete(value T) *Node[T] {
	// O(h)
	if root == nil {
		return nil
	} else if value < root.value {
		root.left = root.left.Delete(value)
		return root
	} else if value > root.value {
		root.right = root.right.Delete(value)
		return root
	} else {
		if root.left != nil && root.right != nil {
			newRoot := root.left
			if newRoot.right == nil {
				newRoot.right = root.right
				return newRoot
			}
			var prev *Node[T]
			for newRoot.right != nil {
				prev = newRoot
				newRoot = newRoot.right
			}
			prev.right = newRoot.left
			newRoot.left = root.left
			newRoot.right = root.right
			return newRoot
		} else if root.left != nil {
			return root.left
		} else if root.right != nil {
			return root.right
		} else {
			return nil
		}
	}
}

func (tree1 *Node[T]) Equal(tree2 *Node[T]) bool {
	// O(n)
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil {
		return false
	}
	if tree1.value != tree2.value {
		return false
	}
	return tree1.left.Equal(tree2.left) && tree1.right.Equal(tree2.right)
}

func (root *Node[T]) Insert(value T) *Node[T] {
	// O(h)
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
