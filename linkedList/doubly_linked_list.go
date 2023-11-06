package linked_list

import "fmt"

type DLLNode[T comparable] struct {
	data T
	prev *DLLNode[T]
	next *DLLNode[T]
}

func CreateDoublyLinkedList[T comparable](array []T) *DLLNode[T] {
	// O(n)
	if len(array) == 0 {
		return nil
	}
	head := &DLLNode[T]{array[0], nil, nil}
	current := head
	for _, data := range array[1:] {
		current.next = &DLLNode[T]{data, current, nil}
		current = current.next
	}
	return head
}

func (head *DLLNode[T]) DLLAppend(value T) *DLLNode[T] {
	// O(n)
	if head == nil {
		return &DLLNode[T]{value, nil, nil}
	}
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = &DLLNode[T]{value, current, nil}
	return head
}

func (head *DLLNode[T]) DLLPrepend(value T) *DLLNode[T] {
	// O(1)
	node := &DLLNode[T]{value, nil, head}
	if head != nil {
		head.prev = node
	}
	return node
}

func (head *DLLNode[T]) DLLDelete(value T) *DLLNode[T] {
	// O(n)
	if head == nil {
		return head
	}
	if head.data == value {
		return head.next
	}
	current := head.next
	for current != nil {
		if current.data == value {
			current.prev.next = current.next
			break
		}
		current = current.next
	}
	return head
}

func (head *DLLNode[T]) DLLToString() string {
	// O(n)
	result := "("
	for head != nil {
		if head.next != nil {
			result += fmt.Sprintf("%v <-> ", head.data)
		} else {
			result += fmt.Sprintf("%v -> ", head.data)
		}
		head = head.next
	}
	result += "nil)"
	return result
}
