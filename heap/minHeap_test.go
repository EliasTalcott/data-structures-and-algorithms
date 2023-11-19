package heap

import (
	"reflect"
	"testing"
)

func TestMinHeapInsert(t *testing.T) {
	heap := MinHeap[int]{2, 3, 4, 5}
	heap = heap.Insert(1)

	expected := MinHeap[int]{1, 2, 4, 5, 3}
	if !reflect.DeepEqual(heap, expected) {
		t.Errorf("Result %v not equal to expected %v", heap, expected)
	}
}

func TestMinHeapDelete(t *testing.T) {
	heap := MinHeap[int]{1, 2, 4, 5, 3}
	heap, value, err := heap.Delete()

	if err != nil {
		t.Error("Deleting from a heap returned an error but should not have")
	}

	expectedHeap := MinHeap[int]{2, 3, 4, 5}
	if !reflect.DeepEqual(heap, expectedHeap) {
		t.Errorf("Result %v not equal to expected %v", heap, expectedHeap)
	}

	expectedValue := 1
	if value != expectedValue {
		t.Errorf("Deleted value %v not equal to expected %v", value, expectedValue)
	}
}

func TestMinHeapEmptyDelete(t *testing.T) {
	heap := MinHeap[int]{}
	_, _, err := heap.Delete()

	if err == nil {
		t.Error("Deleting from an empty heap did not return an error but should have")
	}
}
