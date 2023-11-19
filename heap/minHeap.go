package heap

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] []T

func (heap MinHeap[T]) Insert(value T) MinHeap[T] {
	// O(logn)
	heap = append(heap, value)
	heap.heapifyUp(len(heap) - 1)
	return heap
}

func (heap MinHeap[T]) Delete() (MinHeap[T], T, error) {
	// O(logn)
	var value T
	if len(heap) == 0 {
		return heap, value, errors.New("Cannot delete from an empty heap")
	}

	value = heap[0]
	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1]
	heap.heapifyDown(0)
	return heap, value, nil
}

func (heap MinHeap[T]) heapifyDown(index int) {
	// O(logn)
	length := len(heap)
	value := heap[index]
	leftIndex := leftChild(index)
	rightIndex := rightChild(index)

	if index >= length || leftIndex >= length {
		return
	}

	leftValue := heap[leftIndex]
	if rightIndex < length {
		rightValue := heap[rightIndex]
		if leftValue <= rightValue && leftValue < value {
			heap[index] = leftValue
			heap[leftIndex] = value
			heap.heapifyDown(leftIndex)
		} else if leftValue > rightValue && rightValue < value {
			heap[index] = rightValue
			heap[rightIndex] = value
			heap.heapifyDown(rightIndex)
		}
	} else {
		if leftValue < value {
			heap[index] = leftValue
			heap[leftIndex] = value
		}
	}
}

func (heap MinHeap[T]) heapifyUp(index int) {
	// O(logn)
	if index == 0 {
		return
	}
	value := heap[index]
	parentIndex := parent(index)
	parentValue := heap[parentIndex]
	if value < parentValue {
		heap[parentIndex] = value
		heap[index] = parentValue
		heap.heapifyUp(parentIndex)
	}
}
