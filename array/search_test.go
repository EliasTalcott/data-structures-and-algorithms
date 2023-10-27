package array

import "testing"

func TestLinearSearch(t *testing.T) {
    array := []int{1, 3, 5, 7, 9}

    if LinearSearch(array, 0) == true {
        t.Error("Array does not contain 0, but was found by linear search")
    }
    if LinearSearch(array, 1) == false {
        t.Error("Array contains 1, but was not found by linear search")
    }
    if LinearSearch(array, 9) == false {
        t.Error("Array contains 9, but was not found by linear search")
    }
    if LinearSearch(array, 10) == true {
        t.Error("Array does not contain 10, but was found by linear search")
    }
}

func TestBinarySearch(t *testing.T) {
    array := []int{1, 3, 5, 7, 9}

    if BinarySearch(array, 0) == true {
        t.Error("Array does not contain 0, but was found by binary search")
    }
    if BinarySearch(array, 1) == false {
        t.Error("Array contains 1, but was not found by binary search")
    }
    if BinarySearch(array, 9) == false {
        t.Error("Array contains 9, but was not found by binary search")
    }
    if BinarySearch(array, 10) == true {
        t.Error("Array does not contain 10, but was found by binary search")
    }
}
