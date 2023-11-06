package array

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](array []T) []T {
	// O(n^2)
	for i := len(array) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
	return array
}

func QuickSort[T constraints.Ordered](array []T) {
	// O(nlogn) average case to O(n^2) worst case
	quickSort(array, 0, len(array)-1)
}

func quickSort[T constraints.Ordered](array []T, lo int, hi int) {
	if lo >= hi {
		return
	}
	pivotIdx := partition(array, lo, hi)
	quickSort(array, lo, pivotIdx-1)
	quickSort(array, pivotIdx+1, hi)
}

func partition[T constraints.Ordered](array []T, lo int, hi int) int {
	pivot := array[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if array[i] < pivot {
			idx++
			temp := array[i]
			array[i] = array[idx]
			array[idx] = temp
		}
	}
	idx++
	array[hi] = array[idx]
	array[idx] = pivot
	return idx
}
