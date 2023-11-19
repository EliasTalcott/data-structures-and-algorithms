package heap

func leftChild(index int) int {
	// O(1)
	return index*2 + 1
}

func rightChild(index int) int {
	// O(1)
	return index*2 + 2
}

func parent(index int) int {
	// O(1)
	return (index - 1) / 2
}
