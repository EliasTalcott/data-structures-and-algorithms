package stack

import (
	"errors"
	"fmt"
)

type StackNode[T any] struct {
	data T
	prev *StackNode[T]
}

type Stack[T any] struct {
	head *StackNode[T]
}

func CreateStack[T any](array []T) *Stack[T] {
	// O(n)
	if len(array) == 0 {
		return &Stack[T]{nil}
	}
	head := &StackNode[T]{array[0], nil}
	prev := head
	for _, data := range array[1:] {
		head = &StackNode[T]{data, prev}
		prev = head
	}
	return &Stack[T]{head}
}

func (stack *Stack[T]) Push(data T) {
	// O(1)
	prev := stack.head
	stack.head = &StackNode[T]{data, prev}
}

func (stack *Stack[T]) Pop() (T, error) {
	// O(1)
	var data T
	if stack.head == nil {
		return data, errors.New("Stack is empty!")
	}
	data = stack.head.data
	stack.head = stack.head.prev
	return data, nil
}

func (stack *Stack[T]) ToString() string {
	// O(n)
	result := "(nil"
	for stack.head != nil {
		result += fmt.Sprintf(" <- %v", stack.head.data)
		stack.head = stack.head.prev
	}
	result += ")"
	return result
}
