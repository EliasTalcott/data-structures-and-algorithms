package linked_list

import "fmt"

type ListNode[T comparable] struct {
	data T
	next *ListNode[T]
}

func CreateLinkedList[T comparable](array []T) *ListNode[T] {
	// O(n)
	if len(array) == 0 {
		return nil
	}
	head := &ListNode[T]{array[0], nil}
	current := head
	for _, data := range array[1:] {
		current.next = &ListNode[T]{data, nil}
		current = current.next
	}
	return head
}

func (head *ListNode[T]) Append(value T) *ListNode[T] {
	// O(n)
	if head == nil {
		return &ListNode[T]{value, nil}
	}
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = &ListNode[T]{value, nil}
	return head
}

func (head *ListNode[T]) Prepend(value T) *ListNode[T] {
	// O(1)
	return &ListNode[T]{value, head}
}

func (head *ListNode[T]) Delete(value T) *ListNode[T] {
	// O(n)
	if head == nil {
		return head
	}
	if head.data == value {
		return head.next
	}
	previous := head
	current := head.next
	for current != nil {
		if current.data == value {
			previous.next = current.next
			current.next = nil
			break
		}
		previous = current
		current = current.next
	}
	return head
}

func (head *ListNode[T]) ToString() string {
	// O(n)
	result := "("
	for head != nil {
		result += fmt.Sprintf("%v -> ", head.data)
		head = head.next
	}
	result += "nil)"
	return result
}
