package queue

import (
	"errors"
	"fmt"
)

type QueueNode struct {
	data int
	next *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func CreateQueue(array []int) *Queue {
	// O(n)
	if len(array) == 0 {
		return &Queue{nil, nil}
	}
	head := &QueueNode{array[0], nil}
	current := head
	for _, data := range array[1:] {
		current.next = &QueueNode{data, nil}
		current = current.next
	}
	return &Queue{head, current}
}

func (queue *Queue) Enqueue(data int) {
	// O(1)
	new_node := &QueueNode{data, nil}
	if queue.tail == nil {
		queue.head = new_node
		queue.tail = new_node
		return
	}
	queue.tail.next = new_node
	queue.tail = new_node
}

func (queue *Queue) Dequeue() (int, error) {
	// O(1)
	var data int
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

func (queue *Queue) ToString() string {
	// O(n)
	result := "("
	for queue.head != nil {
		result += fmt.Sprintf("%v ->", queue.head.data)
		queue.head = queue.head.next
	}
	result += "nil)"
	return result
}
