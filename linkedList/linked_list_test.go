package linked_list

import "testing"

func (expected *ListNode[T]) Equal(actual *ListNode[T]) bool {
	return expected.ToString() == actual.ToString()
}

func TestAppend(t *testing.T) {
	head := CreateLinkedList([]int{1, 2, 3})
	head = head.Append(4)
	expected := CreateLinkedList([]int{1, 2, 3, 4})
	if !expected.Equal(head) {
		t.Errorf("Result %v not equal to expected %v", head.ToString(), expected.ToString())
	}
}

func TestPrepend(t *testing.T) {
	head := CreateLinkedList([]int{1, 2, 3})
	head = head.Prepend(0)
	expected := CreateLinkedList([]int{0, 1, 2, 3})
	if !expected.Equal(head) {
		t.Errorf("Result %v not equal to expected %v", head.ToString(), expected.ToString())
	}
}

func TestDelete(t *testing.T) {
	head := CreateLinkedList([]int{1, 2, 3})
	head = head.Delete(3)
	expected := CreateLinkedList([]int{1, 2})
	if !expected.Equal(head) {
		t.Errorf("Result %v not equal to expected %v", head.ToString(), expected.ToString())
	}
}
