package main

import "fmt"

func main() {
    array := []int{1, 2, 3, 4, 5}
    fmt.Println("\nSorted array for searching")
    fmt.Println("--------------------------")
    fmt.Println(array)

    fmt.Println("\nLinear search")
    fmt.Println("-------------")
    fmt.Println("Contains 0: ", LinearSearch(array, 0))
    fmt.Println("Contains 1: ", LinearSearch(array, 1))
    fmt.Println("Contains 5: ", LinearSearch(array, 5))
    fmt.Println("Contains 6: ", LinearSearch(array, 6))

    fmt.Println("\nBinary search")
    fmt.Println("-------------")
    fmt.Println("Contains 0: ", BinarySearch(array, 0))
    fmt.Println("Contains 1: ", BinarySearch(array, 1))
    fmt.Println("Contains 5: ", BinarySearch(array, 5))
    fmt.Println("Contains 6: ", BinarySearch(array, 6))

    initialArray := [10]int{5, 7, 1, 0, 20, 100, 30, -8, 3, 12}
    fmt.Println("\nUnsorted array for sorting")
    fmt.Println("--------------------------")
    fmt.Println(initialArray)

    fmt.Println("\nBubble sort")
    fmt.Println("-----------")
    unsortedArray := initialArray
    fmt.Println(BubbleSort(unsortedArray[:]))

    fmt.Println("")
}

