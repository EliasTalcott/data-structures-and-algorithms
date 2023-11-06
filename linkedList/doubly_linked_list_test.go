package linked_list

import "testing"

func (expected *DLLNode[T]) DLLEqual(actual *DLLNode[T]) bool {
	return expected.DLLToString() == actual.DLLToString()
}

func TestDLLAppend(t *testing.T) {
	head := CreateDoublyLinkedList([]int{1, 2, 3})
	head = head.DLLAppend(4)
	expected := CreateDoublyLinkedList([]int{1, 2, 3, 4})
	if !expected.DLLEqual(head) {
		t.Errorf("Result %v not equal to expected %v", head.DLLToString(), expected.DLLToString())
	}
}

func TestDLLPrepend(t *testing.T) {
	head := CreateDoublyLinkedList([]int{1, 2, 3})
	head = head.DLLPrepend(0)
	expected := CreateDoublyLinkedList([]int{0, 1, 2, 3})
	if !expected.DLLEqual(head) {
		t.Errorf("Result %v not equal to expected %v", head.DLLToString(), expected.DLLToString())
	}
}

func TestDLLDelete(t *testing.T) {
	head := CreateDoublyLinkedList([]int{1, 2, 3})
	head = head.DLLDelete(3)
	expected := CreateDoublyLinkedList([]int{1, 2})
	if !expected.DLLEqual(head) {
		t.Errorf("Result %v not equal to expected %v", head.DLLToString(), expected.DLLToString())
	}
}
