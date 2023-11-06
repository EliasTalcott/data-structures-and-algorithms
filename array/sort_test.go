package array

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	unsortedArray := []int{3, 5, 2, 1, 4}
	sortedArray := BubbleSort(unsortedArray)

	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(expected, sortedArray) {
		t.Errorf("Result %v not equal to expected %v", sortedArray, expected)
	}
}

func TestQuickSort(t *testing.T) {
	array := []int{3, 5, 2, 1, 4}
	QuickSort(array)

	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(expected, array) {
		t.Errorf("Result %v not equal to expected %v", array, expected)
	}
}
