package binarySearchTree

import (
	"reflect"
	"testing"
)

func TestTraverseLevelOrder(t *testing.T) {
	var root *Node[int]
	root = root.Insert(3).
		Insert(5).
		Insert(1).
		Insert(4).
		Insert(2)

	levelOrderTraversal := root.TraverseLevelOrder()
	expected := []int{3, 1, 5, 2, 4}
	if !reflect.DeepEqual(expected, levelOrderTraversal) {
		t.Errorf("Result %v not equal to expected %v", levelOrderTraversal, expected)
	}
}

func TestTraverseInOrder(t *testing.T) {
	var root *Node[int]
	root = root.Insert(3).
		Insert(5).
		Insert(1).
		Insert(4).
		Insert(2)

	inOrderTraversal := root.TraverseInOrder()
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(expected, inOrderTraversal) {
		t.Errorf("Result %v not equal to expected %v", inOrderTraversal, expected)
	}
}

func TestTraversePreOrder(t *testing.T) {
	var root *Node[int]
	root = root.Insert(3).
		Insert(5).
		Insert(1).
		Insert(4).
		Insert(2)

	preOrderTraversal := root.TraversePreOrder()
	expected := []int{3, 1, 2, 5, 4}
	if !reflect.DeepEqual(expected, preOrderTraversal) {
		t.Errorf("Result %v not equal to expected %v", preOrderTraversal, expected)
	}
}

func TestTraversePostOrder(t *testing.T) {
	var root *Node[int]
	root = root.Insert(3).
		Insert(5).
		Insert(1).
		Insert(4).
		Insert(2)

	postOrderTraversal := root.TraversePostOrder()
	expected := []int{2, 1, 4, 5, 3}
	if !reflect.DeepEqual(expected, postOrderTraversal) {
		t.Errorf("Result %v not equal to expected %v", postOrderTraversal, expected)
	}
}
