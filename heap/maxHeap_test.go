package heap

import (
	"reflect"
	"testing"
)

func TestMaxHeapInsert(t *testing.T) {
	heap := MaxHeap[int]{4, 3, 2, 1}
	heap = heap.Insert(5)

	expected := MaxHeap[int]{5, 4, 2, 1, 3}
	if !reflect.DeepEqual(heap, expected) {
		t.Errorf("Result %v not equal to expected %v", heap, expected)
	}
}

func TestMaxHeapDelete(t *testing.T) {
	heap := MaxHeap[int]{5, 4, 2, 1, 3}
	heap, value, err := heap.Delete()

	if err != nil {
		t.Error("Deleting from a heap returned an error but should not have")
	}

	expectedHeap := MaxHeap[int]{4, 3, 2, 1}
	if !reflect.DeepEqual(heap, expectedHeap) {
		t.Errorf("Result %v not equal to expected %v", heap, expectedHeap)
	}

	expectedValue := 5
	if value != expectedValue {
		t.Errorf("Deleted value %v not equal to expected %v", value, expectedValue)
	}
}

func TestMaxHeapEmptyDelete(t *testing.T) {
	heap := MaxHeap[int]{}
	_, _, err := heap.Delete()

	if err == nil {
		t.Error("Deleting from an empty heap did not return an error but should have")
	}
}
