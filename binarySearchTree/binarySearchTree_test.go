package binarySearchTree

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	var root *Node[int]
	root = root.Insert(2).
		Insert(4).
		Insert(1).
		Insert(3).
		Insert(5).
		Insert(0)

	// 2 children
	root = root.Delete(4)
	tree := root.TraversePreOrder()
	expected := []int{2, 1, 0, 3, 5}
	if !reflect.DeepEqual(expected, tree) {
		t.Errorf("Result %v not equal to expected %v", tree, expected)
	}

	// 1 child
	root = root.Delete(1)
	tree = root.TraversePreOrder()
	expected = []int{2, 0, 3, 5}
	if !reflect.DeepEqual(expected, tree) {
		t.Errorf("Result %v not equal to expected %v", tree, expected)
	}

	// Leaf node
	root = root.Delete(0)
	tree = root.TraversePreOrder()
	expected = []int{2, 3, 5}
	if !reflect.DeepEqual(expected, tree) {
		t.Errorf("Result %v not equal to expected %v", tree, expected)
	}
}

func TestEqual(t *testing.T) {
	var tree1, tree2 *Node[int]
	tree1 = tree1.Insert(3).
		Insert(2).
		Insert(4).
		Insert(1).
		Insert(5)

	tree2 = tree2.Insert(3).
		Insert(1).
		Insert(5).
		Insert(2).
		Insert(4)

	if !tree1.Equal(tree1) {
		t.Error("Tree 1 compared to tree 1 was not equal when it should have been")
	}
	if tree1.Equal(tree2) {
		t.Error("Tree 1 compared to tree 2 was equal when it should not have been")
	}
}

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
