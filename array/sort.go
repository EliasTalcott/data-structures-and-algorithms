package main

func BubbleSort(array []int) ([]int) {
    // O(n^2)
    for i := len(array) - 1; i > 0; i-- {
        for j := 0; j < i; j++ {
            if array[j] > array[j + 1] {
                temp := array[j]
                array[j] = array[j + 1]
                array[j + 1] = temp
            }
        }
    }
    return array
}

