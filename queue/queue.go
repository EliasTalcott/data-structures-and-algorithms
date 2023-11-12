package queue

import (
	"errors"
	"fmt"
)

type QueueNode[T any] struct {
	data T
	next *QueueNode[T]
}

type Queue[T any] struct {
	head *QueueNode[T]
	tail *QueueNode[T]
}

func CreateQueue[T any](array []T) *Queue[T] {
	// O(n)
	if len(array) == 0 {
		return &Queue[T]{nil, nil}
	}
	head := &QueueNode[T]{array[0], nil}
	current := head
	for _, data := range array[1:] {
		current.next = &QueueNode[T]{data, nil}
		current = current.next
	}
	return &Queue[T]{head, current}
}

func (queue *Queue[T]) Enqueue(data T) {
	// O(1)
	new_node := &QueueNode[T]{data, nil}
	if queue.tail == nil {
		queue.head = new_node
		queue.tail = new_node
		return
	}
	queue.tail.next = new_node
	queue.tail = new_node
}

func (queue *Queue[T]) Dequeue() (T, error) {
	// O(1)
	var data T
	if queue.head == nil {
		return data, errors.New("Queue is empty!")
	}
	data = queue.head.data
	queue.head = queue.head.next
	if queue.head == nil {
		queue.tail = nil
	}
	return data, nil
}

func (queue *Queue[T]) Empty() bool {
	return queue.head == nil
}

func (queue *Queue[T]) ToString() string {
	// O(n)
	result := "("
	for queue.head != nil {
		result += fmt.Sprintf("%v ->", queue.head.data)
		queue.head = queue.head.next
	}
	result += "nil)"
	return result
}
