package array

import "cmp"

func LinearSearch[T cmp.Ordered](haystack []T, needle T) bool {
	// O(n)
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}

func BinarySearch[T cmp.Ordered](haystack []T, needle T) bool {
	// O(log(n))
	low := 0
	high := len(haystack) - 1
	for {
		midpoint := low + (high-low)/2
		value := haystack[midpoint]
		switch {
		case value == needle:
			return true
		case low >= high:
			return false
		case value < needle:
			low = midpoint + 1
		case value > needle:
			high = midpoint - 1
		}
	}
}
