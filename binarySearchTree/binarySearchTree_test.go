package binarySearchTree

import (
	"reflect"
	"testing"
)

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
